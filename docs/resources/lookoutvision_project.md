---
page_title: "awscc_lookoutvision_project Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.
---

# awscc_lookoutvision_project (Resource)

The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.

## Example Usage

Before trying the examples, verify that your AWS region supports Amazon Lookout for Vision. Reference: [Amazon Lookout for Vision Endpoints and Quotas](https://docs.aws.amazon.com/general/latest/gr/lookoutvision.html)

### Basic

```terraform
resource "awscc_lookoutvision_project" "example" {
  project_name = "example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_name` (String) The name of the Amazon Lookout for Vision project

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_lookoutvision_project.example
  id = "project_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_lookoutvision_project.example "project_name"
```