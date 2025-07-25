---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_redshift_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  An example resource schema demonstrating some basic constructs and validation rules.
---

# awscc_redshift_cluster (Resource)

An example resource schema demonstrating some basic constructs and validation rules.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `cluster_type` (String) The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required
- `db_name` (String) The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.
- `master_username` (String) The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.
- `node_type` (String) The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.large | ra3.4xlarge | ra3.16xlarge

### Optional

- `allow_version_upgrade` (Boolean) Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True
- `aqua_configuration_status` (String) The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.

enabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.
disabled - Don't use AQUA.
auto - Amazon Redshift determines whether to use AQUA.
- `automated_snapshot_retention_period` (Number) The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1
- `availability_zone` (String) The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint
- `availability_zone_relocation` (Boolean) The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.
- `availability_zone_relocation_status` (String) The availability zone relocation status of the cluster
- `classic` (Boolean) A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic.
- `cluster_identifier` (String) A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
- `cluster_parameter_group_name` (String) The name of the parameter group to be associated with this cluster.
- `cluster_security_groups` (List of String) A list of security groups to be associated with this cluster.
- `cluster_subnet_group_name` (String) The name of a cluster subnet group to be associated with this cluster.
- `cluster_version` (String) The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.
- `defer_maintenance` (Boolean) A boolean indicating whether to enable the deferred maintenance window.
- `defer_maintenance_duration` (Number) An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 45 days or less.
- `defer_maintenance_end_time` (String) A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.
- `defer_maintenance_start_time` (String) A timestamp indicating the start time for the deferred maintenance window.
- `destination_region` (String) The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference
- `elastic_ip` (String) The Elastic IP (EIP) address for the cluster.
- `encrypted` (Boolean) If true, the data in the cluster is encrypted at rest.
- `endpoint` (Attributes) (see [below for nested schema](#nestedatt--endpoint))
- `enhanced_vpc_routing` (Boolean) An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.

If this option is true , enhanced VPC routing is enabled.

Default: false
- `hsm_client_certificate_identifier` (String) Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM
- `hsm_configuration_identifier` (String) Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.
- `iam_roles` (List of String) A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request
- `kms_key_id` (String) The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.
- `logging_properties` (Attributes) (see [below for nested schema](#nestedatt--logging_properties))
- `maintenance_track_name` (String) The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.
- `manage_master_password` (Boolean) A boolean indicating if the redshift cluster's admin user credentials is managed by Redshift or not. You can't use MasterUserPassword if ManageMasterPassword is true. If ManageMasterPassword is false or not set, Amazon Redshift uses MasterUserPassword for the admin user account's password.
- `manual_snapshot_retention_period` (Number) The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.

The value must be either -1 or an integer between 1 and 3,653.
- `master_password_secret_kms_key_id` (String) The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.
- `master_user_password` (String) The password associated with the master user account for the cluster that is being created. You can't use MasterUserPassword if ManageMasterPassword is true. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.
- `multi_az` (Boolean) A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.
- `namespace_resource_policy` (String) The namespace resource policy document that will be attached to a Redshift cluster.
- `number_of_nodes` (Number) The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.
- `owner_account` (String)
- `port` (Number) The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings
- `preferred_maintenance_window` (String) The weekly time range (in UTC) during which automated cluster maintenance can occur.
- `publicly_accessible` (Boolean) If true, the cluster can be accessed from a public network.
- `resource_action` (String) The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs
- `revision_target` (String) The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.
- `rotate_encryption_key` (Boolean) A boolean indicating if we want to rotate Encryption Keys.
- `snapshot_cluster_identifier` (String) The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.
- `snapshot_copy_grant_name` (String) The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.
- `snapshot_copy_manual` (Boolean) Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.
- `snapshot_copy_retention_period` (Number) The number of days to retain automated snapshots in the destination region after they are copied from the source region. 

 Default is 7. 

 Constraints: Must be at least 1 and no more than 35.
- `snapshot_identifier` (String) The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.
- `tags` (Attributes List) The list of tags for the cluster parameter group. (see [below for nested schema](#nestedatt--tags))
- `vpc_security_group_ids` (List of String) A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.

### Read-Only

- `cluster_namespace_arn` (String) The Amazon Resource Name (ARN) of the cluster namespace.
- `defer_maintenance_identifier` (String) A unique identifier for the deferred maintenance window.
- `id` (String) Uniquely identifies the resource.
- `master_password_secret_arn` (String) The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.

<a id="nestedatt--endpoint"></a>
### Nested Schema for `endpoint`

Read-Only:

- `address` (String)
- `port` (String)


<a id="nestedatt--logging_properties"></a>
### Nested Schema for `logging_properties`

Optional:

- `bucket_name` (String)
- `log_destination_type` (String)
- `log_exports` (List of String)
- `s3_key_prefix` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_redshift_cluster.example
  id = "cluster_identifier"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_redshift_cluster.example "cluster_identifier"
```
