// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run generators/tfplugingen/main.go -resource awscc_logs_log_group -cfschema ../service/cloudformation/schemas/AWS_Logs_LogGroup.json -- ../aws/xray/log_group_resource_spec.json
//go:generate tfplugingen-framework generate resources --input ../aws/xray/log_group_resource_spec.json --output ../aws/logs --package logs

package provider
