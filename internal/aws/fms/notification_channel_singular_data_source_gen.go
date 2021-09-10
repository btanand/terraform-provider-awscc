// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package fms

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_fms_notification_channel", notificationChannelDataSourceType)
}

// notificationChannelDataSourceType returns the Terraform awscc_fms_notification_channel data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::FMS::NotificationChannel resource type.
func notificationChannelDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"sns_role_name": {
			// Property: SnsRoleName
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::FMS::NotificationChannel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::FMS::NotificationChannel").WithTerraformTypeName("awscc_fms_notification_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"sns_role_name": "SnsRoleName",
		"sns_topic_arn": "SnsTopicArn",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_fms_notification_channel", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}