---
page_title: "awscc_resourceexplorer2_index Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::ResourceExplorer2::Index Resource Type
---

# awscc_resourceexplorer2_index (Resource)

Definition of AWS::ResourceExplorer2::Index Resource Type

## Example Usage

```terraform
resource "awscc_resourceexplorer2_index" "example" {
  type = "LOCAL"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String)

### Optional

- `tags` (Map of String)

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.
- `index_state` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_resourceexplorer2_index.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_resourceexplorer2_index.example "arn"
```