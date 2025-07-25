---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_pcaconnectorscep_connector Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a Connector that allows certificate issuance through Simple Certificate Enrollment Protocol (SCEP)
---

# awscc_pcaconnectorscep_connector (Resource)

Represents a Connector that allows certificate issuance through Simple Certificate Enrollment Protocol (SCEP)



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `certificate_authority_arn` (String)

### Optional

- `mobile_device_management` (Attributes) (see [below for nested schema](#nestedatt--mobile_device_management))
- `tags` (Map of String)

### Read-Only

- `connector_arn` (String)
- `endpoint` (String)
- `id` (String) Uniquely identifies the resource.
- `open_id_configuration` (Attributes) (see [below for nested schema](#nestedatt--open_id_configuration))
- `type` (String)

<a id="nestedatt--mobile_device_management"></a>
### Nested Schema for `mobile_device_management`

Optional:

- `intune` (Attributes) (see [below for nested schema](#nestedatt--mobile_device_management--intune))

<a id="nestedatt--mobile_device_management--intune"></a>
### Nested Schema for `mobile_device_management.intune`

Optional:

- `azure_application_id` (String)
- `domain` (String)



<a id="nestedatt--open_id_configuration"></a>
### Nested Schema for `open_id_configuration`

Read-Only:

- `audience` (String)
- `issuer` (String)
- `subject` (String)

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_pcaconnectorscep_connector.example
  id = "connector_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_pcaconnectorscep_connector.example "connector_arn"
```
