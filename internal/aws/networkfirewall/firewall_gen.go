// Code generated by generators/resource/main.go; DO NOT EDIT.

package networkfirewall

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
	registry.AddResourceTypeFactory("awscc_networkfirewall_firewall", firewallResourceType)
}

// firewallResourceType returns the Terraform awscc_networkfirewall_firewall resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::NetworkFirewall::Firewall resource type.
func firewallResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"delete_protection": {
			// Property: DeleteProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"endpoint_ids": {
			// Property: EndpointIds
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "description": "An endpoint Id.",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Computed: true,
		},
		"firewall_arn": {
			// Property: FirewallArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
			Computed:    true,
		},
		"firewall_id": {
			// Property: FirewallId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 36,
			//   "minLength": 36,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(36, 36)},
			Computed:   true,
		},
		"firewall_name": {
			// Property: FirewallName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
			Required:   true,
			// FirewallName is a force-new attribute.
		},
		"firewall_policy_arn": {
			// Property: FirewallPolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "A resource ARN.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A resource ARN.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
			Required:    true,
		},
		"firewall_policy_change_protection": {
			// Property: FirewallPolicyChangeProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"subnet_change_protection": {
			// Property: SubnetChangeProtection
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
		},
		"subnet_mappings": {
			// Property: SubnetMappings
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "SubnetId": {
			//         "description": "A SubnetId.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "SubnetId"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"subnet_id": {
						// Property: SubnetId
						Description: "A SubnetId.",
						Type:        types.StringType,
						Required:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 255,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: providertypes.SetNestedAttributes(
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
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(0, 255)},
						Required:   true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"vpc_id": {
			// Property: VpcId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
			Required:   true,
			// VpcId is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource type definition for AWS::NetworkFirewall::Firewall",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkFirewall::Firewall").WithTerraformTypeName("awscc_networkfirewall_firewall")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"delete_protection":                 "DeleteProtection",
		"description":                       "Description",
		"endpoint_ids":                      "EndpointIds",
		"firewall_arn":                      "FirewallArn",
		"firewall_id":                       "FirewallId",
		"firewall_name":                     "FirewallName",
		"firewall_policy_arn":               "FirewallPolicyArn",
		"firewall_policy_change_protection": "FirewallPolicyChangeProtection",
		"key":                               "Key",
		"subnet_change_protection":          "SubnetChangeProtection",
		"subnet_id":                         "SubnetId",
		"subnet_mappings":                   "SubnetMappings",
		"tags":                              "Tags",
		"value":                             "Value",
		"vpc_id":                            "VpcId",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_networkfirewall_firewall", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
