// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build ignore
// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-codegen-spec/provider"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	"github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-codegen-spec/spec"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/common"
	"github.com/hashicorp/terraform-provider-awscc/internal/provider/generators/shared"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-tfplugingen-spec-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	g := NewGenerator()

	tfPluginGenSpecFilename := args[0]

	if err := g.Generate(tfPluginGenSpecFilename); err != nil {
		g.Fatalf("error generating Terraform Provider Code Specification: %s", err)
	}
}

type Generator struct {
	*common.Generator
	cfTypeSchemaFile string
	tfResourceType   string
}

func NewGenerator() *Generator {
	return &Generator{
		Generator:        common.NewGenerator(),
		cfTypeSchemaFile: *cfTypeSchemaFile,
		tfResourceType:   *tfResourceType,
	}
}

func (g *Generator) Generate(specFilename string) error {
	g.Infof("generating Terraform Provider Code Specification for %[1]q from %[2]q into %[3]q", g.tfResourceType, g.cfTypeSchemaFile, specFilename)

	res, err := shared.NewResource(g.tfResourceType, g.cfTypeSchemaFile)

	if err != nil {
		return fmt.Errorf("reading CloudFormation resource schema for %s: %w", g.tfResourceType, err)
	}

	specResource, err := g.generateProviderCodeSpec(res)

	if err != nil {
		return err
	}

	spec := &spec.Specification{
		Provider: &provider.Provider{
			Name: "awscc",
		},
		Resources: []resource.Resource{*specResource},
		Version:   spec.Version0_1,
	}

	d := g.NewUnformattedFileDestination(specFilename)

	if err := d.CreateDirectories(); err != nil {
		return err
	}

	if err := d.WriteJSON(spec); err != nil {
		return err
	}

	if err := d.Write(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateProviderCodeSpec(res *shared.Resource) (*resource.Resource, error) {
	schema := &resource.Schema{
		Attributes: []resource.Attribute{
			{
				Name: "floop",
				Bool: &resource.BoolAttribute{
					ComputedOptionalRequired: schema.Required,
				},
			},
		},
	}

	return &resource.Resource{
		Name:   res.TfType,
		Schema: schema,
	}, nil
}
