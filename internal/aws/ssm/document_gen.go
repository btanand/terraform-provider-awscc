// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

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
	registry.AddResourceTypeFactory("awscc_ssm_document", documentResourceType)
}

// documentResourceType returns the Terraform awscc_ssm_document resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSM::Document resource type.
func documentResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"attachments": {
			// Property: Attachments
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key and value pairs that describe attachments to a version of a document.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The key of a key-value pair that identifies the location of an attachment to a document.",
			//         "enum": [
			//           "SourceUrl",
			//           "S3FileUrl",
			//           "AttachmentReference"
			//         ],
			//         "type": "string"
			//       },
			//       "Name": {
			//         "description": "The name of the document attachment file.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Values": {
			//         "description": "The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.",
			//         "items": {
			//           "maxLength": 100000,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "maxItems": 1,
			//         "minItems": 1,
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 20,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "A list of key and value pairs that describe attachments to a version of a document.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key of a key-value pair that identifies the location of an attachment to a document.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						Description: "The name of the document attachment file.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Optional:    true,
					},
					"values": {
						// Property: Values
						Description: "The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.",
						Type:        types.ListType{ElemType: types.StringType},
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 1),
						},
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 20,
				},
			),
			Optional: true,
			Computed: true,
			// Attachments is a force-new attribute.
		},
		"content": {
			// Property: Content
			// CloudFormation resource type schema:
			// {
			//   "description": "The content for the Systems Manager document in JSON, YAML or String format.",
			//   "type": "string"
			// }
			Description: "The content for the Systems Manager document in JSON, YAML or String format.",
			Type:        types.StringType,
			Required:    true,
			// Content is a force-new attribute.
		},
		"document_format": {
			// Property: DocumentFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.",
			//   "enum": [
			//     "YAML",
			//     "JSON",
			//     "TEXT"
			//   ],
			//   "type": "string"
			// }
			Description: "Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// DocumentFormat is a force-new attribute.
		},
		"document_type": {
			// Property: DocumentType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of document to create.",
			//   "enum": [
			//     "ApplicationConfiguration",
			//     "ApplicationConfigurationSchema",
			//     "Automation",
			//     "Automation.ChangeTemplate",
			//     "ChangeCalendar",
			//     "CloudFormation",
			//     "Command",
			//     "DeploymentStrategy",
			//     "Package",
			//     "Policy",
			//     "ProblemAnalysis",
			//     "ProblemAnalysisTemplate",
			//     "Session"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of document to create.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// DocumentType is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the Systems Manager document.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the Systems Manager document.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"requires": {
			// Property: Requires
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "description": "The name of the required SSM document. The name can be an Amazon Resource Name (ARN).",
			//         "maxLength": 200,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Version": {
			//         "description": "The document version required by the current document.",
			//         "maxLength": 8,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of the required SSM document. The name can be an Amazon Resource Name (ARN).",
						Type:        types.StringType,
						Optional:    true,
					},
					"version": {
						// Property: Version
						Description: "The document version required by the current document.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Optional: true,
			Computed: true,
			// Requires is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "The name of the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value of the tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 1000,
			//   "type": "array"
			// }
			Description: "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The name of the tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Optional:    true,
					},
					"value": {
						// Property: Value
						Description: "The value of the tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
						Optional:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 1000,
				},
			),
			Optional: true,
		},
		"target_type": {
			// Property: TargetType
			// CloudFormation resource type schema:
			// {
			//   "description": "Specify a target type to define the kinds of resources the document can run on.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Specify a target type to define the kinds of resources the document can run on.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// TargetType is a force-new attribute.
		},
		"version_name": {
			// Property: VersionName
			// CloudFormation resource type schema:
			// {
			//   "description": "An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// VersionName is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The AWS::SSM::Document resource is an SSM document in AWS Systems Manager that defines the actions that Systems Manager performs, which can be used to set up and run commands on your instances.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::Document").WithTerraformTypeName("awscc_ssm_document")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachments":     "Attachments",
		"content":         "Content",
		"document_format": "DocumentFormat",
		"document_type":   "DocumentType",
		"key":             "Key",
		"name":            "Name",
		"requires":        "Requires",
		"tags":            "Tags",
		"target_type":     "TargetType",
		"value":           "Value",
		"values":          "Values",
		"version":         "Version",
		"version_name":    "VersionName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_ssm_document", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
