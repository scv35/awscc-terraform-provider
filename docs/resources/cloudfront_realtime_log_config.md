
---
page_title: "awscc_cloudfront_realtime_log_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  A real-time log configuration.
---

# awscc_cloudfront_realtime_log_config (Resource)

A real-time log configuration.

## Example Usage

### CloudFront Real-Time Log Configuration with Kinesis

Creates a CloudFront real-time log configuration that streams logs to a Kinesis stream with comprehensive field selection and 100% sampling rate, supported by necessary IAM roles and permissions.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
# Kinesis stream to receive the logs
resource "awscc_kinesis_stream" "example" {
  name                   = "cloudfront-realtime-logs"
  retention_period_hours = 24
  shard_count            = 1
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# IAM role for CloudFront to write to Kinesis
data "aws_iam_policy_document" "cloudfront_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["cloudfront.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "cloudfront_kinesis" {
  statement {
    actions = [
      "kinesis:DescribeStreamSummary",
      "kinesis:DescribeStream",
      "kinesis:PutRecord",
      "kinesis:PutRecords"
    ]
    effect    = "Allow"
    resources = [awscc_kinesis_stream.example.arn]
  }
}

resource "awscc_iam_role" "cloudfront_realtime_logs" {
  role_name                   = "cloudfront-realtime-logs"
  assume_role_policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.cloudfront_assume_role.json))
  policies = [{
    policy_document = jsonencode(jsondecode(data.aws_iam_policy_document.cloudfront_kinesis.json))
    policy_name     = "cloudfront-kinesis-access"
  }]
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}

# CloudFront real-time log config
resource "awscc_cloudfront_realtime_log_config" "example" {
  name          = "example-realtime-logs"
  sampling_rate = 100
  fields = [
    "timestamp",
    "c-ip",
    "time-to-first-byte",
    "sc-status",
    "sc-bytes",
    "cs-method",
    "cs-protocol",
    "cs-host",
    "cs-uri-stem",
    "cs-bytes",
    "x-edge-location",
    "x-edge-request-id",
    "x-host-header",
    "time-taken",
    "cs-protocol-version",
    "c-ip-version",
    "cs-user-agent",
    "cs-referer"
  ]
  end_points = [
    {
      stream_type = "Kinesis"
      kinesis_stream_config = {
        role_arn   = awscc_iam_role.cloudfront_realtime_logs.arn
        stream_arn = awscc_kinesis_stream.example.arn
      }
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `end_points` (Attributes List) Contains information about the Amazon Kinesis data stream where you are sending real-time log data for this real-time log configuration. (see [below for nested schema](#nestedatt--end_points))
- `fields` (List of String) A list of fields that are included in each real-time log record. In an API response, the fields are provided in the same order in which they are sent to the Amazon Kinesis data stream.
 For more information about fields, see [Real-time log configuration fields](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) in the *Amazon CloudFront Developer Guide*.
- `name` (String) The unique name of this real-time log configuration.
- `sampling_rate` (Number) The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. The sampling rate is an integer between 1 and 100, inclusive.

### Read-Only

- `arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--end_points"></a>
### Nested Schema for `end_points`

Required:

- `kinesis_stream_config` (Attributes) Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration. (see [below for nested schema](#nestedatt--end_points--kinesis_stream_config))
- `stream_type` (String) The type of data stream where you are sending real-time log data. The only valid value is ``Kinesis``.

<a id="nestedatt--end_points--kinesis_stream_config"></a>
### Nested Schema for `end_points.kinesis_stream_config`

Required:

- `role_arn` (String) The Amazon Resource Name (ARN) of an IAMlong (IAM) role that CloudFront can use to send real-time log data to your Kinesis data stream.
 For more information the IAM role, see [Real-time log configuration IAM role](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) in the *Amazon CloudFront Developer Guide*.
- `stream_arn` (String) The Amazon Resource Name (ARN) of the Kinesis data stream where you are sending real-time log data.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_cloudfront_realtime_log_config.example
  id = "arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_cloudfront_realtime_log_config.example "arn"
```
