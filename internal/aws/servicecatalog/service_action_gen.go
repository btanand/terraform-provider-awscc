// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog

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
	registry.AddResourceTypeFactory("awscc_servicecatalog_service_action", serviceActionResourceType)
}

// serviceActionResourceType returns the Terraform awscc_servicecatalog_service_action resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalog::ServiceAction resource type.
func serviceActionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accept_language": {
			// Property: AcceptLanguage
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "en",
			//     "jp",
			//     "zh"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"definition": {
			// Property: Definition
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 4096,
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
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 1000)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"definition_type": {
			// Property: DefinitionType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "SSM_AUTOMATION"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Computed:   true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
			Required:   true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::ServiceCatalog::ServiceAction",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::ServiceAction").WithTerraformTypeName("awscc_servicecatalog_service_action")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accept_language": "AcceptLanguage",
		"definition":      "Definition",
		"definition_type": "DefinitionType",
		"description":     "Description",
		"id":              "Id",
		"key":             "Key",
		"name":            "Name",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_servicecatalog_service_action", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
