// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_datasync_location_object_storage", locationObjectStorageResourceType)
}

// locationObjectStorageResourceType returns the Terraform awscc_datasync_location_object_storage resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource type.
func locationObjectStorageResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"access_key": {
			// Property: AccessKey
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			//   "maxLength": 200,
			//   "minLength": 8,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(8, 200)},
			Optional:    true,
		},
		"agent_arns": {
			// Property: AgentArns
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 128,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "maxItems": 4,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			Type:        types.ListType{ElemType: types.StringType},
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 4),
			},
			Required: true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the bucket on the self-managed object storage server.",
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the bucket on the self-managed object storage server.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(3, 63)},
			Required:    true,
			// BucketName is a force-new attribute.
			// BucketName is a write-only attribute.
		},
		"location_arn": {
			// Property: LocationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the location that is created.",
			//   "maxLength": 128,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the location that is created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location_uri": {
			// Property: LocationUri
			// CloudFormation resource type schema:
			// {
			//   "description": "The URL of the object storage location that was described.",
			//   "maxLength": 4356,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The URL of the object storage location that was described.",
			Type:        types.StringType,
			Computed:    true,
		},
		"secret_key": {
			// Property: SecretKey
			// CloudFormation resource type schema:
			// {
			//   "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			//   "maxLength": 200,
			//   "minLength": 8,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(8, 200)},
			Optional:    true,
			// SecretKey is a write-only attribute.
		},
		"server_hostname": {
			// Property: ServerHostname
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			//   "maxLength": 255,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			Type:        types.StringType,
			Required:    true,
			// ServerHostname is a force-new attribute.
			// ServerHostname is a write-only attribute.
		},
		"server_port": {
			// Property: ServerPort
			// CloudFormation resource type schema:
			// {
			//   "description": "The port that your self-managed server accepts inbound network traffic on.",
			//   "type": "integer"
			// }
			Description: "The port that your self-managed server accepts inbound network traffic on.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"server_protocol": {
			// Property: ServerProtocol
			// CloudFormation resource type schema:
			// {
			//   "description": "The protocol that the object storage server uses to communicate.",
			//   "enum": [
			//     "HTTPS",
			//     "HTTP"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol that the object storage server uses to communicate.",
			Type:        types.StringType,
			Optional:    true,
		},
		"subdirectory": {
			// Property: Subdirectory
			// CloudFormation resource type schema:
			// {
			//   "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
			//   "maxLength": 4096,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The subdirectory in the self-managed object storage server that is used to read data from.",
			Type:        types.StringType,
			Optional:    true,
			// Subdirectory is a write-only attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for an AWS resource tag.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key for an AWS resource tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for an AWS resource tag.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::DataSync::LocationObjectStorage.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("awscc_datasync_location_object_storage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_key":      "AccessKey",
		"agent_arns":      "AgentArns",
		"bucket_name":     "BucketName",
		"key":             "Key",
		"location_arn":    "LocationArn",
		"location_uri":    "LocationUri",
		"secret_key":      "SecretKey",
		"server_hostname": "ServerHostname",
		"server_port":     "ServerPort",
		"server_protocol": "ServerProtocol",
		"subdirectory":    "Subdirectory",
		"tags":            "Tags",
		"value":           "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SecretKey",
		"/properties/BucketName",
		"/properties/ServerHostname",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_datasync_location_object_storage", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
