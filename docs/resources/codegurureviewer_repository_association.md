---
page_title: "awscc_codegurureviewer_repository_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.
---

# awscc_codegurureviewer_repository_association (Resource)

This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.

## Example Usage

### S3 as the repository to associate CodeGuru Reviewer with.  The bucket name must start with the prefix codeguru-reviewer-*

```terraform
resource "awscc_codegurureviewer_repository_association" "example2" {
  bucket_name = var.bucket_name
  name        = var.bucket_name
  type        = "S3Bucket"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

variable "bucket_name" {
  type = string
}
```

### CodeCommit as the repository to associate CodeGuru Reviewer with. Make sure to add a lifecycle rule to ignore the tag changes once CodeGuru Reviewer is associated with CodeCommit, if the repository was created with Terraform. 

```terraform
resource "awscc_codegurureviewer_repository_association" "example" {
  name = var.repo_name
  type = "CodeCommit"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]

}

variable "repo_name" {
  type = string
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the repository to be associated.
- `type` (String) The type of repository to be associated.

### Optional

- `bucket_name` (String) The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.
- `connection_arn` (String) The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.
- `owner` (String) The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.
- `tags` (Attributes List) The tags associated with a repository association. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `association_arn` (String) The Amazon Resource Name (ARN) of the repository association.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_codegurureviewer_repository_association.example
  id = "association_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_codegurureviewer_repository_association.example "association_arn"
```