---
page_title: "awscc_sagemaker_image Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::SageMaker::Image
---

# awscc_sagemaker_image (Resource)

Resource Type definition for AWS::SageMaker::Image

## Example Usage

To create a SageMaker Image resource.
```terraform
resource "awscc_sagemaker_image" "example" {
  image_name     = "example"
  image_role_arn = awscc_iam_role.main.arn
  tags = [
    {
      key   = "Name"
      value = "example"
    },
    {
      key   = "Environment"
      value = "Development"
    },
    {
      key   = "Modified By"
      value = "AWSCC"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `image_name` (String) The name of the image.
- `image_role_arn` (String) The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.

### Optional

- `image_description` (String) A description of the image.
- `image_display_name` (String) The display name of the image.
- `tags` (Attributes List) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `image_arn` (String) The Amazon Resource Name (ARN) of the image.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_sagemaker_image.example
  id = "image_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_sagemaker_image.example "image_arn"
```
