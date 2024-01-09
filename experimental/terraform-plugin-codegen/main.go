// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-provider-awscc/experimental/terraform-plugin-codegen/internal/awscc"
)

func main() {
	ctx := context.Background()
	schema := awscc.AwsccServiceThingResourceSchema(ctx)

	fmt.Printf("%s\n", schema.Description)
}
