
---
page_title: "awscc_mediatailor_source_location Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::MediaTailor::SourceLocation Resource Type
---

# awscc_mediatailor_source_location (Resource)

Definition of AWS::MediaTailor::SourceLocation Resource Type

## Example Usage

### Configure MediaTailor Source Location

To create a MediaTailor source location with HTTP configuration that points to an S3 bucket as the content source, use this example which automatically adapts to your current AWS region.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Get current AWS region
data "aws_region" "current" {}

# AWS MediaTailor Source Location example
resource "awscc_mediatailor_source_location" "example" {
  source_location_name = "example-source-location"
  http_configuration = {
    base_url = "https://s3.${data.aws_region.current.name}.amazonaws.com/example-bucket/content"
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

- `http_configuration` (Attributes) <p>The HTTP configuration for the source location.</p> (see [below for nested schema](#nestedatt--http_configuration))
- `source_location_name` (String)

### Optional

- `access_configuration` (Attributes) <p>Access configuration parameters.</p> (see [below for nested schema](#nestedatt--access_configuration))
- `default_segment_delivery_configuration` (Attributes) <p>The optional configuration for a server that serves segments. Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.</p> (see [below for nested schema](#nestedatt--default_segment_delivery_configuration))
- `segment_delivery_configurations` (Attributes List) <p>A list of the segment delivery configurations associated with this resource.</p> (see [below for nested schema](#nestedatt--segment_delivery_configurations))
- `tags` (Attributes Set) The tags to assign to the source location. (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String) <p>The ARN of the source location.</p>
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--http_configuration"></a>
### Nested Schema for `http_configuration`

Required:

- `base_url` (String) <p>The base URL for the source location host server. This string must include the protocol, such as <b>https://</b>.</p>


<a id="nestedatt--access_configuration"></a>
### Nested Schema for `access_configuration`

Optional:

- `access_type` (String)
- `secrets_manager_access_token_configuration` (Attributes) <p>AWS Secrets Manager access token configuration parameters. For information about Secrets Manager access token authentication, see <a href="https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-access-configuration-access-token.html">Working with AWS Secrets Manager access token authentication</a>.</p> (see [below for nested schema](#nestedatt--access_configuration--secrets_manager_access_token_configuration))

<a id="nestedatt--access_configuration--secrets_manager_access_token_configuration"></a>
### Nested Schema for `access_configuration.secrets_manager_access_token_configuration`

Optional:

- `header_name` (String) <p>The name of the HTTP header used to supply the access token in requests to the source location.</p>
- `secret_arn` (String) <p>The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the access token.</p>
- `secret_string_key` (String) <p>The AWS Secrets Manager <a href="https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html#SecretsManager-CreateSecret-request-SecretString.html">SecretString</a> key associated with the access token. MediaTailor uses the key to look up SecretString key and value pair containing the access token.</p>



<a id="nestedatt--default_segment_delivery_configuration"></a>
### Nested Schema for `default_segment_delivery_configuration`

Optional:

- `base_url` (String) <p>The hostname of the server that will be used to serve segments. This string must include the protocol, such as <b>https://</b>.</p>


<a id="nestedatt--segment_delivery_configurations"></a>
### Nested Schema for `segment_delivery_configurations`

Optional:

- `base_url` (String) <p>The base URL of the host or path of the segment delivery server that you're using to serve segments. This is typically a content delivery network (CDN). The URL can be absolute or relative. To use an absolute URL include the protocol, such as <code>https://example.com/some/path</code>. To use a relative URL specify the relative path, such as <code>/some/path*</code>.</p>
- `name` (String) <p>A unique identifier used to distinguish between multiple segment delivery configurations in a source location.</p>


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
  to = awscc_mediatailor_source_location.example
  id = "source_location_name"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_mediatailor_source_location.example "source_location_name"
```
