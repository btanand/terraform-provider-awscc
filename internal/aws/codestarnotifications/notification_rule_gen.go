// Code generated by generators/resource/main.go; DO NOT EDIT.

package codestarnotifications

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"

	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_codestarnotifications_notification_rule", notificationRuleResourceType)
}

// notificationRuleResourceType returns the Terraform awscc_codestarnotifications_notification_rule resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeStarNotifications::NotificationRule resource type.
func notificationRuleResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"detail_type": {
			// Property: DetailType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "BASIC",
			//     "FULL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"event_type_ids": {
			// Property: EventTypeIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "maxLength": 200,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 64)},
			Required:   true,
		},
		"resource": {
			// Property: Resource
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// Resource is a force-new attribute.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "ENABLED",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"targets": {
			// Property: Targets
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "TargetAddress": {
			//         "type": "string"
			//       },
			//       "TargetType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TargetType",
			//       "TargetAddress"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 10,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"target_address": {
						// Property: TargetAddress
						Type:     types.StringType,
						Required: true,
					},
					"target_type": {
						// Property: TargetType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 10,
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::CodeStarNotifications::NotificationRule",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeStarNotifications::NotificationRule").WithTerraformTypeName("awscc_codestarnotifications_notification_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":            "Arn",
		"detail_type":    "DetailType",
		"event_type_ids": "EventTypeIds",
		"name":           "Name",
		"resource":       "Resource",
		"status":         "Status",
		"tags":           "Tags",
		"target_address": "TargetAddress",
		"target_type":    "TargetType",
		"targets":        "Targets",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_codestarnotifications_notification_rule", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
