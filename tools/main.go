// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build tools
// +build tools

package main

import (
	_ "github.com/client9/misspell/cmd/misspell"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/hashicorp/terraform-plugin-codegen-framework/cmd/tfplugingen-framework"
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/pavius/impi/cmd/impi"
)
