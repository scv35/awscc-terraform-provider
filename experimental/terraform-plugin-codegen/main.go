// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-provider-awscc/experimental/terraform-plugin-codegen/internal/schemas"
)

func main() {
	ctx := context.Background()
	schema := schemas.AwsccServiceThingResourceSchema(ctx)

	fmt.Printf("%s\n", schema.Description)
}
