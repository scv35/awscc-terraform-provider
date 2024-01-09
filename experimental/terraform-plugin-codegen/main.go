// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-provider-awscc/experimental/terraform-plugin-codegen/internal/provider/provider_awscc_gen"
)

func main() {
	ctx := context.Background()
	schema := provider_awscc_gen.AwsccGenProviderSchema(ctx)

	fmt.Printf("%s\n", schema.Description)
}
