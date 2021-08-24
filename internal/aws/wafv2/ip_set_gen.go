// Code generated by generators/resource/main.go; DO NOT EDIT.

package wafv2

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
	registry.AddResourceTypeFactory("awscc_wafv2_ip_set", iPSetResourceType)
}

// iPSetResourceType returns the Terraform awscc_wafv2_ip_set resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WAFv2::IPSet resource type.
func iPSetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"addresses": {
			// Property: Addresses
			// CloudFormation resource type schema:
			// {
			//   "description": "List of IPAddresses.",
			//   "items": {
			//     "description": "IP address",
			//     "maxLength": 50,
			//     "minLength": 1,
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "List of IPAddresses.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "ARN of the WAF entity.",
			//   "type": "string"
			// }
			Description: "ARN of the WAF entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of the entity.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Description of the entity.",
			Type:        types.StringType,
			Optional:    true,
		},
		"ip_address_version": {
			// Property: IPAddressVersion
			// CloudFormation resource type schema:
			// {
			//   "description": "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			//   "enum": [
			//     "IPV4",
			//     "IPV6"
			//   ],
			//   "type": "string"
			// }
			Description: "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
			Type:        types.StringType,
			Required:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id of the IPSet",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Id of the IPSet",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the IPSet.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the IPSet.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"scope": {
			// Property: Scope
			// CloudFormation resource type schema:
			// {
			//   "description": "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			//   "enum": [
			//     "CLOUDFRONT",
			//     "REGIONAL"
			//   ],
			//   "type": "string"
			// }
			Description: "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
			Type:        types.StringType,
			Required:    true,
			// Scope is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Optional:   true,
					},
					"value": {
						// Property: Value
						Type:       types.StringType,
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(0, 256)},
						Optional:   true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::IPSet").WithTerraformTypeName("awscc_wafv2_ip_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addresses":          "Addresses",
		"arn":                "Arn",
		"description":        "Description",
		"id":                 "Id",
		"ip_address_version": "IPAddressVersion",
		"key":                "Key",
		"name":               "Name",
		"scope":              "Scope",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_wafv2_ip_set", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
