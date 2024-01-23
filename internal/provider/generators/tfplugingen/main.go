// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -cfschema <CF-type-schema-file> <generated-tfplugingen-spec-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	tfPluginGenSpecFilename := args[0]
}
