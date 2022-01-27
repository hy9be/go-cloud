// Copyright 2019 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blob_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hy9be/gocloud/blob"
	"github.com/hy9be/gocloud/blob/memblob"
	"github.com/hy9be/gocloud/gcerrors"
	"github.com/hy9be/gocloud/internal/oc"
	"github.com/hy9be/gocloud/internal/testing/octest"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

func TestOpenCensus(t *testing.T) {
	ctx := context.Background()
	te := octest.NewTestExporter(blob.OpenCensusViews)
	defer te.Unregister()

	bytes := []byte("foo")
	b := memblob.OpenBucket(nil)
	defer b.Close()
	if err := b.WriteAll(ctx, "key", bytes, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := b.ReadAll(ctx, "key"); err != nil {
		t.Fatal(err)
	}
	if _, err := b.Attributes(ctx, "key"); err != nil {
		t.Fatal(err)
	}
	if _, _, err := b.ListPage(ctx, blob.FirstPageToken, 3, nil); err != nil {
		t.Fatal(err)
	}
	if err := b.Delete(ctx, "key"); err != nil {
		t.Fatal(err)
	}
	if _, err := b.ReadAll(ctx, "noSuchKey"); err == nil {
		t.Fatal("got nil, want error")
	}

	const driver = "github.com/hy9be/gocloud/blob/memblob"

	diff := octest.Diff(te.Spans(), te.Counts(), "github.com/hy9be/gocloud/blob", driver, []octest.Call{
		{Method: "NewWriter", Code: gcerrors.OK},
		{Method: "NewRangeReader", Code: gcerrors.OK},
		{Method: "Attributes", Code: gcerrors.OK},
		{Method: "ListPage", Code: gcerrors.OK},
		{Method: "Delete", Code: gcerrors.OK},
		{Method: "NewRangeReader", Code: gcerrors.NotFound},
	})
	if diff != "" {
		t.Error(diff)
	}

	// Find and verify the bytes read/written metrics.
	var sawRead, sawWritten bool
	tags := []tag.Tag{{Key: oc.ProviderKey, Value: driver}}
	for !sawRead || !sawWritten {
		data := <-te.Stats
		switch data.View.Name {
		case "github.com/hy9be/gocloud/blob/bytes_read":
			if sawRead {
				continue
			}
			sawRead = true
		case "github.com/hy9be/gocloud/blob/bytes_written":
			if sawWritten {
				continue
			}
			sawWritten = true
		default:
			continue
		}
		if diff := cmp.Diff(data.Rows[0].Tags, tags, cmp.AllowUnexported(tag.Key{})); diff != "" {
			t.Errorf("tags for %s: %s", data.View.Name, diff)
			continue
		}
		sd, ok := data.Rows[0].Data.(*view.SumData)
		if !ok {
			t.Errorf("%s: data is %T, want SumData", data.View.Name, data.Rows[0].Data)
			continue
		}
		if got := int(sd.Value); got < len(bytes) {
			t.Errorf("%s: got %d, want at least %d", data.View.Name, got, len(bytes))
		}
	}
}
