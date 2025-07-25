
---
page_title: "awscc_ec2_snapshot_block_public_access Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess
---

# awscc_ec2_snapshot_block_public_access (Resource)

Resource Type definition for AWS::EC2::SnapshotBlockPublicAccess

## Example Usage

### Block EBS Snapshot Public Access

Configures EC2 EBS snapshot block public access settings to prevent all public sharing of EBS snapshots at the account level.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Create EBS snapshot block public access
resource "awscc_ec2_snapshot_block_public_access" "example" {
  state = "block-all-sharing" # block all sharing of EBS snapshots
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `state` (String) The state of EBS Snapshot Block Public Access.

### Read-Only

- `account_id` (String) The identifier for the specified AWS account.
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_snapshot_block_public_access.example
  id = "account_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_snapshot_block_public_access.example "account_id"
```
