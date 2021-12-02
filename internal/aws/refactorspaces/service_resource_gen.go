// Code generated by generators/resource/main.go; DO NOT EDIT.

package refactorspaces

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_refactorspaces_service", serviceResourceType)
}

// serviceResourceType returns the Terraform awscc_refactorspaces_service resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::RefactorSpaces::Service resource type.
func serviceResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_identifier": {
			// Property: ApplicationIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(14, 14),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// Description is a write-only property.
		},
		"endpoint_type": {
			// Property: EndpointType
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "LAMBDA",
			//     "URL"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"LAMBDA",
					"URL",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// EndpointType is a write-only property.
		},
		"environment_identifier": {
			// Property: EnvironmentIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(14, 14),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"lambda_endpoint": {
			// Property: LambdaEndpoint
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Arn": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Arn"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// LambdaEndpoint is a write-only property.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(3, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// Name is a write-only property.
		},
		"service_identifier": {
			// Property: ServiceIdentifier
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 14,
			//   "minLength": 14,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A label for tagging Environment resource",
			//     "properties": {
			//       "Key": {
			//         "description": "A string used to identify this tag",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "A string containing the value for the tag",
			//         "maxLength": 256,
			//         "minLength": 0,
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
			Description: "Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A string used to identify this tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "A string containing the value for the tag",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
		"url_endpoint": {
			// Property: UrlEndpoint
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "HealthUrl": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Url": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Url"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"health_url": {
						// Property: HealthUrl
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
					"url": {
						// Property: Url
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 2048),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// UrlEndpoint is a write-only property.
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 21,
			//   "minLength": 12,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(12, 21),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// VpcId is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Definition of AWS::RefactorSpaces::Service Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::RefactorSpaces::Service").WithTerraformTypeName("awscc_refactorspaces_service")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_identifier": "ApplicationIdentifier",
		"arn":                    "Arn",
		"description":            "Description",
		"endpoint_type":          "EndpointType",
		"environment_identifier": "EnvironmentIdentifier",
		"health_url":             "HealthUrl",
		"key":                    "Key",
		"lambda_endpoint":        "LambdaEndpoint",
		"name":                   "Name",
		"service_identifier":     "ServiceIdentifier",
		"tags":                   "Tags",
		"url":                    "Url",
		"url_endpoint":           "UrlEndpoint",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Description",
		"/properties/EndpointType",
		"/properties/LambdaEndpoint",
		"/properties/Name",
		"/properties/UrlEndpoint",
		"/properties/VpcId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
