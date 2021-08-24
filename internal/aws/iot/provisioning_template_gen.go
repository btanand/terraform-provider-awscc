// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("awscc_iot_provisioning_template", provisioningTemplateResourceType)
}

// provisioningTemplateResourceType returns the Terraform awscc_iot_provisioning_template resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::ProvisioningTemplate resource type.
func provisioningTemplateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 500,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"pre_provisioning_hook": {
			// Property: PreProvisioningHook
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "PayloadVersion": {
			//       "type": "string"
			//     },
			//     "TargetArn": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"payload_version": {
						// Property: PayloadVersion
						Type:     types.StringType,
						Optional: true,
					},
					"target_arn": {
						// Property: TargetArn
						Type:     types.StringType,
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"provisioning_role_arn": {
			// Property: ProvisioningRoleArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"template_arn": {
			// Property: TemplateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"template_body": {
			// Property: TemplateBody
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"template_name": {
			// Property: TemplateName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 36)},
			Optional:   true,
			Computed:   true,
			// TemplateName is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Creates a fleet provisioning template.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::ProvisioningTemplate").WithTerraformTypeName("awscc_iot_provisioning_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":           "Description",
		"enabled":               "Enabled",
		"key":                   "Key",
		"payload_version":       "PayloadVersion",
		"pre_provisioning_hook": "PreProvisioningHook",
		"provisioning_role_arn": "ProvisioningRoleArn",
		"tags":                  "Tags",
		"target_arn":            "TargetArn",
		"template_arn":          "TemplateArn",
		"template_body":         "TemplateBody",
		"template_name":         "TemplateName",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iot_provisioning_template", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
