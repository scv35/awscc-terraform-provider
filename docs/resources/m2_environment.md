
---
page_title: "awscc_m2_environment Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Represents a runtime environment that can run migrated mainframe applications.
---

# awscc_m2_environment (Resource)

Represents a runtime environment that can run migrated mainframe applications.

## Example Usage

### High-Availability M2 Environment Setup

Creates a Microfocus M2 environment with high availability configuration across two subnets, utilizing m5.xlarge instances and custom KMS encryption, supported by necessary networking components including VPC, subnets, and security groups.

~> This example is generated by LLM using Amazon Bedrock and validated using terraform validate, apply and destroy. While we strive for accuracy and quality, please note that the information provided may not be entirely error-free or up-to-date. We recommend independently verifying the content.

```terraform
data "aws_region" "current" {}

# VPC and Network Configuration
resource "awscc_ec2_vpc" "m2_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true
  tags = [{
    key   = "Name"
    value = "m2-environment-vpc"
  }]
}

resource "awscc_ec2_internet_gateway" "m2_igw" {
  tags = [{
    key   = "Name"
    value = "m2-environment-igw"
  }]
}

resource "awscc_ec2_vpc_gateway_attachment" "m2_igw_attachment" {
  internet_gateway_id = awscc_ec2_internet_gateway.m2_igw.id
  vpc_id              = awscc_ec2_vpc.m2_vpc.id
}

resource "awscc_ec2_subnet" "m2_subnet_1" {
  vpc_id            = awscc_ec2_vpc.m2_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "${data.aws_region.current.name}a"
  tags = [{
    key   = "Name"
    value = "m2-environment-subnet-1"
  }]
}

resource "awscc_ec2_subnet" "m2_subnet_2" {
  vpc_id            = awscc_ec2_vpc.m2_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "${data.aws_region.current.name}b"
  tags = [{
    key   = "Name"
    value = "m2-environment-subnet-2"
  }]
}

# Security Group
resource "awscc_ec2_security_group" "m2_sg" {
  group_description = "Security group for M2 Environment"
  vpc_id            = awscc_ec2_vpc.m2_vpc.id
  security_group_egress = [{
    ip_protocol = "-1"
    from_port   = 0
    to_port     = 0
    cidr_ip     = "0.0.0.0/0"
  }]
  group_name = "m2-environment-sg"
  tags = [{
    key   = "Name"
    value = "m2-environment-sg"
  }]
}

# KMS Key
resource "awscc_kms_key" "m2_key" {
  description            = "KMS key for M2 Environment"
  pending_window_in_days = 7
  tags = [{
    key   = "Name"
    value = "m2-environment-key"
  }]
}

# M2 Environment
resource "awscc_m2_environment" "example" {
  name          = "my-m2-environment"
  engine_type   = "microfocus"
  instance_type = "m5.xlarge"
  description   = "Example M2 Environment for mainframe applications"

  high_availability_config = {
    desired_capacity = 2
  }

  subnet_ids         = [awscc_ec2_subnet.m2_subnet_1.id, awscc_ec2_subnet.m2_subnet_2.id]
  security_group_ids = [awscc_ec2_security_group.m2_sg.id]
  kms_key_id         = awscc_kms_key.m2_key.arn

  publicly_accessible = false
  network_type        = "dual"

  preferred_maintenance_window = "sun:23:00-mon:01:30"

  tags = [{
    key   = "Environment"
    value = "test"
    }, {
    key   = "Modified_By"
    value = "AWSCC"
  }]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `engine_type` (String) The target platform for the environment.
- `instance_type` (String) The type of instance underlying the environment.
- `name` (String) The name of the environment.

### Optional

- `description` (String) The description of the environment.
- `engine_version` (String) The version of the runtime engine for the environment.
- `high_availability_config` (Attributes) Defines the details of a high availability configuration. (see [below for nested schema](#nestedatt--high_availability_config))
- `kms_key_id` (String) The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting environment-related resources.
- `network_type` (String)
- `preferred_maintenance_window` (String) Configures a desired maintenance window for the environment. If you do not provide a value, a random system-generated value will be assigned.
- `publicly_accessible` (Boolean) Specifies whether the environment is publicly accessible.
- `security_group_ids` (List of String) The list of security groups for the VPC associated with this environment.
- `storage_configurations` (Attributes List) The storage configurations defined for the runtime environment. (see [below for nested schema](#nestedatt--storage_configurations))
- `subnet_ids` (List of String) The unique identifiers of the subnets assigned to this runtime environment.
- `tags` (Map of String) Tags associated to this environment.

### Read-Only

- `environment_arn` (String) The Amazon Resource Name (ARN) of the runtime environment.
- `environment_id` (String) The unique identifier of the environment.
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--high_availability_config"></a>
### Nested Schema for `high_availability_config`

Optional:

- `desired_capacity` (Number)


<a id="nestedatt--storage_configurations"></a>
### Nested Schema for `storage_configurations`

Optional:

- `efs` (Attributes) Defines the storage configuration for an Amazon EFS file system. (see [below for nested schema](#nestedatt--storage_configurations--efs))
- `fsx` (Attributes) Defines the storage configuration for an Amazon FSx file system. (see [below for nested schema](#nestedatt--storage_configurations--fsx))

<a id="nestedatt--storage_configurations--efs"></a>
### Nested Schema for `storage_configurations.efs`

Optional:

- `file_system_id` (String) The file system identifier.
- `mount_point` (String) The mount point for the file system.


<a id="nestedatt--storage_configurations--fsx"></a>
### Nested Schema for `storage_configurations.fsx`

Optional:

- `file_system_id` (String) The file system identifier.
- `mount_point` (String) The mount point for the file system.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_m2_environment.example
  id = "environment_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_m2_environment.example "environment_arn"
```
