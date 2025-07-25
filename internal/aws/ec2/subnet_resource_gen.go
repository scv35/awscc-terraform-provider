// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_subnet", subnetResource)
}

// subnetResource returns the Terraform awscc_ec2_subnet resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::Subnet resource.
func subnetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssignIpv6AddressOnCreation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
		//	  "type": "boolean"
		//	}
		"assign_ipv_6_address_on_creation": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ``false``.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone of the subnet.\n If you update this property, you must also update the ``CidrBlock`` property.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone of the subnet.\n If you update this property, you must also update the ``CidrBlock`` property.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZoneId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AZ ID of the subnet.",
		//	  "type": "string"
		//	}
		"availability_zone_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AZ ID of the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BlockPublicAccessStates
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "",
		//	  "properties": {
		//	    "InternetGatewayBlockMode": {
		//	      "description": "The mode of VPC BPA. Options here are off, block-bidirectional, block-ingress ",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"block_public_access_states": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: InternetGatewayBlockMode
				"internet_gateway_block_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The mode of VPC BPA. Options here are off, block-bidirectional, block-ingress ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv4 CIDR block assigned to the subnet.\n If you update this property, we create a new subnet, and then delete the existing one.",
		//	  "type": "string"
		//	}
		"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv4 CIDR block assigned to the subnet.\n If you update this property, we create a new subnet, and then delete the existing one.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableDns64
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.\n  You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ``0.0.0.0/0`` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.",
		//	  "type": "boolean"
		//	}
		"enable_dns_64": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.\n  You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ``0.0.0.0/0`` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableLniAtDeviceIndex
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).",
		//	  "type": "integer"
		//	}
		"enable_lni_at_device_index": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Indicates the device position for local network interfaces in this subnet. For example, ``1`` indicates local network interfaces in this subnet are the secondary network interface (eth1).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// EnableLniAtDeviceIndex is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Ipv4IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv4 IPAM pool ID for the subnet.",
		//	  "type": "string"
		//	}
		"ipv_4_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An IPv4 IPAM pool ID for the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Ipv4IpamPoolId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Ipv4NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv4 netmask length for the subnet.",
		//	  "type": "integer"
		//	}
		"ipv_4_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "An IPv4 netmask length for the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Ipv4NetmaskLength is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv6 CIDR block.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
		//	  "type": "string"
		//	}
		"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv6 CIDR block.\n If you specify ``AssignIpv6AddressOnCreation``, you must also specify an IPv6 CIDR block.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlocks
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"ipv_6_cidr_blocks": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ipv6IpamPoolId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv6 IPAM pool ID for the subnet.",
		//	  "type": "string"
		//	}
		"ipv_6_ipam_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An IPv6 IPAM pool ID for the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Ipv6IpamPoolId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Ipv6Native
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.",
		//	  "type": "boolean"
		//	}
		"ipv_6_native": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ipv6NetmaskLength
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IPv6 netmask length for the subnet.",
		//	  "type": "integer"
		//	}
		"ipv_6_netmask_length": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "An IPv6 netmask length for the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Ipv6NetmaskLength is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: MapPublicIpOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.\n AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
		//	  "type": "boolean"
		//	}
		"map_public_ip_on_launch": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ``false``.\n AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkAclAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"network_acl_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Outpost.",
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Outpost.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateDnsNameOptionsOnLaunch
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.\n Available options:\n  +  EnableResourceNameDnsAAAARecord (true | false)\n  +  EnableResourceNameDnsARecord (true | false)\n  +  HostnameType (ip-name | resource-name)",
		//	  "properties": {
		//	    "EnableResourceNameDnsAAAARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "EnableResourceNameDnsARecord": {
		//	      "type": "boolean"
		//	    },
		//	    "HostnameType": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"private_dns_name_options_on_launch": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EnableResourceNameDnsAAAARecord
				"enable_resource_name_dns_aaaa_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EnableResourceNameDnsARecord
				"enable_resource_name_dns_a_record": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HostnameType
				"hostname_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.\n Available options:\n  +  EnableResourceNameDnsAAAARecord (true | false)\n  +  EnableResourceNameDnsARecord (true | false)\n  +  HostnameType (ip-name | resource-name)",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the subnet.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Resource tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the subnet.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC the subnet is in.\n If you update this property, you must also update the ``CidrBlock`` property.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC the subnet is in.\n If you update this property, you must also update the ``CidrBlock`` property.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Specifies a subnet for the specified VPC.\n For an IPv4 only subnet, specify an IPv4 CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6 CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6 CIDR block.\n For more information, see [Subnets for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the *Amazon VPC User Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Subnet").WithTerraformTypeName("awscc_ec2_subnet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_ipv_6_address_on_creation":     "AssignIpv6AddressOnCreation",
		"availability_zone":                    "AvailabilityZone",
		"availability_zone_id":                 "AvailabilityZoneId",
		"block_public_access_states":           "BlockPublicAccessStates",
		"cidr_block":                           "CidrBlock",
		"enable_dns_64":                        "EnableDns64",
		"enable_lni_at_device_index":           "EnableLniAtDeviceIndex",
		"enable_resource_name_dns_a_record":    "EnableResourceNameDnsARecord",
		"enable_resource_name_dns_aaaa_record": "EnableResourceNameDnsAAAARecord",
		"hostname_type":                        "HostnameType",
		"internet_gateway_block_mode":          "InternetGatewayBlockMode",
		"ipv_4_ipam_pool_id":                   "Ipv4IpamPoolId",
		"ipv_4_netmask_length":                 "Ipv4NetmaskLength",
		"ipv_6_cidr_block":                     "Ipv6CidrBlock",
		"ipv_6_cidr_blocks":                    "Ipv6CidrBlocks",
		"ipv_6_ipam_pool_id":                   "Ipv6IpamPoolId",
		"ipv_6_native":                         "Ipv6Native",
		"ipv_6_netmask_length":                 "Ipv6NetmaskLength",
		"key":                                  "Key",
		"map_public_ip_on_launch":              "MapPublicIpOnLaunch",
		"network_acl_association_id":           "NetworkAclAssociationId",
		"outpost_arn":                          "OutpostArn",
		"private_dns_name_options_on_launch":   "PrivateDnsNameOptionsOnLaunch",
		"subnet_id":                            "SubnetId",
		"tags":                                 "Tags",
		"value":                                "Value",
		"vpc_id":                               "VpcId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EnableLniAtDeviceIndex",
		"/properties/Ipv4IpamPoolId",
		"/properties/Ipv4NetmaskLength",
		"/properties/Ipv6IpamPoolId",
		"/properties/Ipv6NetmaskLength",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
