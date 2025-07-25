
---
page_title: "awscc_ec2_network_performance_metric_subscription Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::EC2::NetworkPerformanceMetricSubscription
---

# awscc_ec2_network_performance_metric_subscription (Resource)

Resource Type definition for AWS::EC2::NetworkPerformanceMetricSubscription

## Example Usage

### Network Performance Metric Subscription Between Regions

Creates a subscription to monitor network performance metrics between two AWS regions, specifically tracking aggregate latency with p50 statistics from US East 1 to US West 2.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
resource "awscc_ec2_network_performance_metric_subscription" "example" {
  source      = "us-east-1" # Source region
  destination = "us-west-2" # Destination region (must be different from source)
  metric      = "aggregate-latency"
  statistic   = "p50"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination` (String) The target Region or Availability Zone for the metric to subscribe to.
- `metric` (String) The metric type to subscribe to.
- `source` (String) The starting Region or Availability Zone for metric to subscribe to.
- `statistic` (String) The statistic to subscribe to.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_ec2_network_performance_metric_subscription.example
  id = "source|destination|metric|statistic"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_ec2_network_performance_metric_subscription.example "source|destination|metric|statistic"
```
