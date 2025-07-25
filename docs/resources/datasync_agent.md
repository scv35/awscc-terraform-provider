---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_datasync_agent Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::DataSync::Agent.
---

# awscc_datasync_agent (Resource)

Resource schema for AWS::DataSync::Agent.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `activation_key` (String) Activation key of the Agent.
- `agent_name` (String) The name configured for the agent. Text reference used to identify the agent in the console.
- `security_group_arns` (List of String) The ARNs of the security group used to protect your data transfer task subnets.
- `subnet_arns` (List of String) The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `vpc_endpoint_id` (String) The ID of the VPC endpoint that the agent has access to.

### Read-Only

- `agent_arn` (String) The DataSync Agent ARN.
- `endpoint_type` (String) The service endpoints that the agent will connect to.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key for an AWS resource tag.
- `value` (String) The value for an AWS resource tag.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_datasync_agent.example
  id = "agent_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_datasync_agent.example "agent_arn"
```
