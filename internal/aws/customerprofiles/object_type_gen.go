// Code generated by generators/resource/main.go; DO NOT EDIT.

package customerprofiles

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
	registry.AddResourceTypeFactory("awscc_customerprofiles_object_type", objectTypeResourceType)
}

// objectTypeResourceType returns the Terraform awscc_customerprofiles_object_type resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CustomerProfiles::ObjectType resource type.
func objectTypeResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"allow_profile_creation": {
			// Property: AllowProfileCreation
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether a profile should be created when data is received.",
			//   "type": "boolean"
			// }
			Description: "Indicates whether a profile should be created when data is received.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got created.",
			//   "type": "string"
			// }
			Description: "The time of this integration got created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the profile object type.",
			//   "maxLength": 1000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Description of the profile object type.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 1000)},
			Optional:    true,
		},
		"domain_name": {
			// Property: DomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique name of the domain.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The unique name of the domain.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 64)},
			Required:    true,
			// DomainName is a force-new attribute.
		},
		"encryption_key": {
			// Property: EncryptionKey
			// CloudFormation resource type schema:
			// {
			//   "description": "The default encryption key",
			//   "maxLength": 255,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "The default encryption key",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(0, 255)},
			Optional:    true,
		},
		"expiration_days": {
			// Property: ExpirationDays
			// CloudFormation resource type schema:
			// {
			//   "description": "The default number of days until the data within the domain expires.",
			//   "type": "integer"
			// }
			Description: "The default number of days until the data within the domain expires.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"fields": {
			// Property: Fields
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of the name and ObjectType field.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ObjectTypeField": {
			//         "additionalProperties": false,
			//         "description": "Represents a field in a ProfileObjectType.",
			//         "properties": {
			//           "ContentType": {
			//             "description": "The content type of the field. Used for determining equality when searching.",
			//             "enum": [
			//               "STRING",
			//               "NUMBER",
			//               "PHONE_NUMBER",
			//               "EMAIL_ADDRESS",
			//               "NAME"
			//             ],
			//             "type": "string"
			//           },
			//           "Source": {
			//             "description": "A field of a ProfileObject. For example: _source.FirstName, where \"_source\" is a ProfileObjectType of a Zendesk user and \"FirstName\" is a field in that ObjectType.",
			//             "maxLength": 1000,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Target": {
			//             "description": "The location of the data in the standard ProfileObject model. For example: _profile.Address.PostalCode.",
			//             "maxLength": 1000,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of the name and ObjectType field.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 64)},
						Optional:   true,
					},
					"object_type_field": {
						// Property: ObjectTypeField
						Description: "Represents a field in a ProfileObjectType.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"content_type": {
									// Property: ContentType
									Description: "The content type of the field. Used for determining equality when searching.",
									Type:        types.StringType,
									Optional:    true,
								},
								"source": {
									// Property: Source
									Description: "A field of a ProfileObject. For example: _source.FirstName, where \"_source\" is a ProfileObjectType of a Zendesk user and \"FirstName\" is a field in that ObjectType.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 1000)},
									Optional:    true,
								},
								"target": {
									// Property: Target
									Description: "The location of the data in the standard ProfileObject model. For example: _profile.Address.PostalCode.",
									Type:        types.StringType,
									Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 1000)},
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"keys": {
			// Property: Keys
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of unique keys that can be used to map data to the profile.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Name": {
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ObjectTypeKeyList": {
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "description": "An object that defines the Key element of a ProfileObject. A Key is a special element that can be used to search for a customer profile.",
			//           "properties": {
			//             "FieldNames": {
			//               "description": "The reference for the key name of the fields map. ",
			//               "items": {
			//                 "maxLength": 64,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             },
			//             "StandardIdentifiers": {
			//               "description": "The types of keys that a ProfileObject can have. Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.",
			//               "items": {
			//                 "enum": [
			//                   "PROFILE",
			//                   "UNIQUE",
			//                   "SECONDARY",
			//                   "LOOKUP_ONLY",
			//                   "NEW_ONLY"
			//                 ],
			//                 "type": "string"
			//               },
			//               "type": "array"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "type": "array"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Description: "A list of unique keys that can be used to map data to the profile.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 64)},
						Optional:   true,
					},
					"object_type_key_list": {
						// Property: ObjectTypeKeyList
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"field_names": {
									// Property: FieldNames
									Description: "The reference for the key name of the fields map. ",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
								},
								"standard_identifiers": {
									// Property: StandardIdentifiers
									Description: "The types of keys that a ProfileObject can have. Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.",
									Type:        types.ListType{ElemType: types.StringType},
									Optional:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"last_updated_at": {
			// Property: LastUpdatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The time of this integration got last updated at.",
			//   "type": "string"
			// }
			Description: "The time of this integration got last updated at.",
			Type:        types.StringType,
			Computed:    true,
		},
		"object_type_name": {
			// Property: ObjectTypeName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the profile object type.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the profile object type.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 255)},
			Optional:    true,
			Computed:    true,
			// ObjectTypeName is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags (keys and values) associated with the integration.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
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
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "The tags (keys and values) associated with the integration.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Required:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(0, 256)},
						Required:   true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 0,
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"template_id": {
			// Property: TemplateId
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the object template.",
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for the object template.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 64)},
			Optional:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An ObjectType resource of Amazon Connect Customer Profiles",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CustomerProfiles::ObjectType").WithTerraformTypeName("awscc_customerprofiles_object_type")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_profile_creation": "AllowProfileCreation",
		"content_type":           "ContentType",
		"created_at":             "CreatedAt",
		"description":            "Description",
		"domain_name":            "DomainName",
		"encryption_key":         "EncryptionKey",
		"expiration_days":        "ExpirationDays",
		"field_names":            "FieldNames",
		"fields":                 "Fields",
		"key":                    "Key",
		"keys":                   "Keys",
		"last_updated_at":        "LastUpdatedAt",
		"name":                   "Name",
		"object_type_field":      "ObjectTypeField",
		"object_type_key_list":   "ObjectTypeKeyList",
		"object_type_name":       "ObjectTypeName",
		"source":                 "Source",
		"standard_identifiers":   "StandardIdentifiers",
		"tags":                   "Tags",
		"target":                 "Target",
		"template_id":            "TemplateId",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_customerprofiles_object_type", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
