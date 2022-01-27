// Copyright 2020 The Go Cloud Development Kit Authors
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

package awssecretsmanager_test

import (
	"context"
	"log"

	awsv2cfg "github.com/aws/aws-sdk-go-v2/config"
	secretsmanagerv2 "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/hy9be/gocloud/runtimevar"
	"github.com/hy9be/gocloud/runtimevar/awssecretsmanager"
)

func ExampleOpenVariable() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.

	// Establish an AWS session.
	// See https://docs.aws.amazon.com/sdk-for-go/api/aws/session/ for more info.
	sess, err := session.NewSession(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a *runtimevar.Variable that watches the variable.
	// `secret-variable-name` must be a friendly name of the secret, NOT the Amazon Resource Name (ARN).
	v, err := awssecretsmanager.OpenVariable(sess, "secret-variable-name", runtimevar.StringDecoder, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()
}

func ExampleOpenVariableV2() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.

	// Establish a AWS V2 Config.
	// See https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/ for more info.
	ctx := context.Background()
	cfg, err := awsv2cfg.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Construct a *runtimevar.Variable that watches the variable.
	// `secret-variable-name` must be a friendly name of the secret, NOT the Amazon Resource Name (ARN).
	clientV2 := secretsmanagerv2.NewFromConfig(cfg)
	v, err := awssecretsmanager.OpenVariableV2(clientV2, "secret-variable-name", runtimevar.StringDecoder, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()
}

func Example_openVariableFromURL() {
	// PRAGMA: This example is used on gocloud.dev; PRAGMA comments adjust how it is shown and can be ignored.
	// PRAGMA: On gocloud.dev, add a blank import: _ "github.com/hy9be/gocloud/runtimevar/awssecretsmanager"
	// PRAGMA: On gocloud.dev, hide lines until the next blank line.
	ctx := context.Background()

	// runtimevar.OpenVariable creates a *runtimevar.Variable from a URL.
	// `secret-variable-name` must be a friendly name of the secret, NOT the Amazon Resource Name (ARN).
	v, err := runtimevar.OpenVariable(ctx, "awssecretsmanager://secret-variable-name?region=us-east-2&decoder=string")
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()

	// Use "awssdk=v1" or "v2" to force a specific AWS SDK version.
	vUsingV2, err := runtimevar.OpenVariable(ctx, "awssecretsmanager://secret-variable-name?region=us-east-2&decoder=string&awssdk=v2")
	if err != nil {
		log.Fatal(err)
	}
	defer vUsingV2.Close()
}
