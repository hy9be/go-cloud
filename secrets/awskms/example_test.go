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

package awskms_test

import (
	"context"
	"log"

	awsv2cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hy9be/gocloud/secrets"
	"github.com/hy9be/gocloud/secrets/awskms"
)

func ExampleOpenKeeper() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	sess, err := session.NewSession(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get a client to use with the KMS API.
	client, err := awskms.Dial(sess)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a *secrets.Keeper.
	keeper := awskms.OpenKeeper(client, "alias/test-secrets", nil)
	defer keeper.Close()
}

func ExampleOpenKeeperV2() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.

	// Establish a AWS V2 Config.
	// See https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/ for more info.
	ctx := context.Background()
	cfg, err := awsv2cfg.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Get a client to use with the KMS API.
	client, err := awskms.DialV2(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a *secrets.Keeper.
	keeper := awskms.OpenKeeperV2(client, "alias/test-secrets", nil)
	defer keeper.Close()
}

func Example_openFromURL() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On gocloud.dev, add a blank import: _ "github.com/hy9be/gocloud/secrets/awskms"
	// PRAGMA: On gocloud.dev, hide lines until the next blank line.
	ctx := context.Background()

	// Use one of the following:

	// 1. By ID.
	keeperByID, err := secrets.OpenKeeper(ctx,
		"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1")
	if err != nil {
		log.Fatal(err)
	}
	defer keeperByID.Close()

	// 2. By alias.
	keeperByAlias, err := secrets.OpenKeeper(ctx,
		"awskms://alias/ExampleAlias?region=us-east-1")
	if err != nil {
		log.Fatal(err)
	}
	defer keeperByAlias.Close()

	// 3. By ARN.
	const arn = "arn:aws:kms:us-east-1:111122223333:key/" +
		"1234abcd-12ab-34bc-56ef-1234567890ab"
	keeperByARN, err := secrets.OpenKeeper(ctx,
		"awskms://"+arn+"?region=us-east-1")
	if err != nil {
		log.Fatal(err)
	}
	defer keeperByARN.Close()

	// Use "awssdk=v1" or "v2" to force a specific AWS SDK version.
	keeperUsingV2, err := secrets.OpenKeeper(ctx,
		"awskms://1234abcd-12ab-34cd-56ef-1234567890ab?region=us-east-1&awssdk=v2")
	if err != nil {
		log.Fatal(err)
	}
	defer keeperUsingV2.Close()
}
