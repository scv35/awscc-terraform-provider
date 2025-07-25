
---
page_title: "awscc_voiceid_domain Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.
---

# awscc_voiceid_domain (Resource)

The AWS::VoiceID::Domain resource specifies an Amazon VoiceID Domain.

## Example Usage

### VoiceID Domain with KMS Encryption

Creates a VoiceID domain with server-side encryption using a dedicated KMS key, which includes necessary key policy for VoiceID service integration.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Create a KMS key for VoiceID Domain
resource "awscc_kms_key" "voiceid_key" {
  description = "KMS key for VoiceID Domain"
  key_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow VoiceID Service"
        Effect = "Allow"
        Principal = {
          Service = "voiceid.amazonaws.com"
        }
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:GenerateDataKey"
        ]
        Resource = "*"
      }
    ]
  })
}

# Create VoiceID Domain
resource "awscc_voiceid_domain" "example" {
  name        = "example-voiceid-domain"
  description = "Example VoiceID Domain created with AWSCC provider"

  server_side_encryption_configuration = {
    kms_key_id = awscc_kms_key.voiceid_key.id
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `server_side_encryption_configuration` (Attributes) (see [below for nested schema](#nestedatt--server_side_encryption_configuration))

### Optional

- `description` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `domain_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--server_side_encryption_configuration"></a>
### Nested Schema for `server_side_encryption_configuration`

Required:

- `kms_key_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_voiceid_domain.example
  id = "domain_id"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_voiceid_domain.example "domain_id"
```
