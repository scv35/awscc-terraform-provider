
---
page_title: "awscc_servicecatalogappregistry_attribute_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.
---

# awscc_servicecatalogappregistry_attribute_group (Resource)

Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroup.

## Example Usage

### Service Catalog App Registry Attribute Group

Creates a Service Catalog AppRegistry attribute group with customizable attributes for production environment, including team and cost-center information with appropriate tags.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
resource "awscc_servicecatalogappregistry_attribute_group" "example" {
  name        = "example-attribute-group"
  description = "Example attribute group created via AWSCC provider"

  attributes = jsonencode({
    "environment" = "production"
    "team"        = "infrastructure"
    "cost-center" = "12345"
  })

  tags = [{
    key   = "ModifiedBy"
    value = "AWSCC"
    }, {
    key   = "Environment"
    value = "Example"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `attributes` (String)
- `name` (String) The name of the attribute group.

### Optional

- `description` (String) The description of the attribute group.
- `tags` (Map of String)

### Read-Only

- `arn` (String)
- `attribute_group_id` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_servicecatalogappregistry_attribute_group.example
  id = "id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_servicecatalogappregistry_attribute_group.example "id"
```
