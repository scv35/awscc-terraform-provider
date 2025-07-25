
---
page_title: "awscc_ssm_association Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.
---

# awscc_ssm_association (Resource)

The AWS::SSM::Association resource associates an SSM document in AWS Systems Manager with EC2 instances that contain a configuration agent to process the document.

## Example Usage

### PowerShell Script Association with S3 Output

Creates an SSM association to run PowerShell script on Production-tagged instances with scheduled execution, storing output in a secure S3 bucket that has proper access controls and bucket policies for SSM service.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Data sources for region and caller identity
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

# Create an AWS SSM Association for running AWS-RunPowerShellScript
resource "awscc_ssm_association" "example" {
  name             = "AWS-RunPowerShellScript"
  association_name = "example-powershell-association"

  targets = [{
    key    = "tag:Environment"
    values = ["Production"]
  }]

  parameters = {
    commands = ["Get-Service"]
  }

  schedule_expression = "cron(0 0 * * ? *)"

  output_location = {
    s3_location = {
      output_s3_bucket_name = awscc_s3_bucket.output.id
      output_s3_key_prefix  = "ssm-output/"
      output_s3_region      = data.aws_region.current.name
    }
  }

  max_concurrency     = "50%"
  max_errors          = "25%"
  compliance_severity = "MEDIUM"
}

# Create S3 bucket for output
resource "awscc_s3_bucket" "output" {
  bucket_name = "ssm-association-output-${data.aws_caller_identity.current.account_id}-${data.aws_region.current.name}"

  public_access_block_configuration = {
    block_public_acls       = true
    block_public_policy     = true
    ignore_public_acls      = true
    restrict_public_buckets = true
  }

  ownership_controls = {
    rules = [{
      object_ownership = "BucketOwnerPreferred"
    }]
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# Create bucket policy for SSM
data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["ssm.amazonaws.com"]
    }
    actions = [
      "s3:PutObject"
    ]
    resources = [
      "${awscc_s3_bucket.output.arn}/*"
    ]
    condition {
      test     = "StringEquals"
      variable = "aws:SourceAccount"
      values   = [data.aws_caller_identity.current.account_id]
    }
    condition {
      test     = "StringEquals"
      variable = "aws:SourceRegion"
      values   = [data.aws_region.current.name]
    }
  }
}

resource "awscc_s3_bucket_policy" "allow_ssm" {
  bucket          = awscc_s3_bucket.output.id
  policy_document = data.aws_iam_policy_document.bucket_policy.json
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the SSM document.

### Optional

- `apply_only_at_cron_interval` (Boolean)
- `association_name` (String) The name of the association.
- `automation_target_parameter_name` (String)
- `calendar_names` (List of String)
- `compliance_severity` (String)
- `document_version` (String) The version of the SSM document to associate with the target.
- `instance_id` (String) The ID of the instance that the SSM document is associated with.
- `max_concurrency` (String)
- `max_errors` (String)
- `output_location` (Attributes) (see [below for nested schema](#nestedatt--output_location))
- `parameters` (Map of List of String) Parameter values that the SSM document uses at runtime.
- `schedule_expression` (String) A Cron or Rate expression that specifies when the association is applied to the target.
- `schedule_offset` (Number)
- `sync_compliance` (String)
- `targets` (Attributes List) The targets that the SSM document sends commands to. (see [below for nested schema](#nestedatt--targets))
- `wait_for_success_timeout_seconds` (Number)

### Read-Only

- `association_id` (String) Unique identifier of the association.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--output_location"></a>
### Nested Schema for `output_location`

Optional:

- `s3_location` (Attributes) (see [below for nested schema](#nestedatt--output_location--s3_location))

<a id="nestedatt--output_location--s3_location"></a>
### Nested Schema for `output_location.s3_location`

Optional:

- `output_s3_bucket_name` (String)
- `output_s3_key_prefix` (String)
- `output_s3_region` (String)



<a id="nestedatt--targets"></a>
### Nested Schema for `targets`

Optional:

- `key` (String)
- `values` (List of String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ssm_association.example
  id = "association_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ssm_association.example "association_id"
```
