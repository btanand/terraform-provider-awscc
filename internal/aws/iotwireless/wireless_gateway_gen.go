// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotwireless

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
	registry.AddResourceTypeFactory("awscc_iotwireless_wireless_gateway", wirelessGatewayResourceType)
}

// wirelessGatewayResourceType returns the Terraform awscc_iotwireless_wireless_gateway resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::WirelessGateway resource type.
func wirelessGatewayResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Arn for Wireless Gateway. Returned upon successful create.",
			//   "type": "string"
			// }
			Description: "Arn for Wireless Gateway. Returned upon successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of Wireless Gateway.",
			//   "maxLength": 2048,
			//   "type": "string"
			// }
			Description: "Description of Wireless Gateway.",
			Type:        types.StringType,
			Optional:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Id for Wireless Gateway. Returned upon successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Id for Wireless Gateway. Returned upon successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"last_uplink_received_at": {
			// Property: LastUplinkReceivedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the most recent uplink was received.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date and time when the most recent uplink was received.",
			Type:        types.StringType,
			Optional:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "GatewayEui": {
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "RfRegion": {
			//       "maxLength": 64,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "GatewayEui",
			//     "RfRegion"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"gateway_eui": {
						// Property: GatewayEui
						Type:     types.StringType,
						Required: true,
					},
					"rf_region": {
						// Property: RfRegion
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Required: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Wireless Gateway.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Name of Wireless Gateway.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the gateway.",
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
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the gateway.",
			Attributes: providertypes.SetNestedAttributes(
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
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"thing_arn": {
			// Property: ThingArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
			//   "type": "string"
			// }
			Description: "Thing Arn. Passed into Update to associate a Thing with the Wireless Gateway.",
			Type:        types.StringType,
			Optional:    true,
		},
		"thing_name": {
			// Property: ThingName
			// CloudFormation resource type schema:
			// {
			//   "description": "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			//   "type": "string"
			// }
			Description: "Thing Arn. If there is a Thing created, this can be returned with a Get call.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	schema := tfsdk.Schema{
		Description: "Create and manage wireless gateways, including LoRa gateways.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::WirelessGateway").WithTerraformTypeName("awscc_iotwireless_wireless_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"gateway_eui":             "GatewayEui",
		"id":                      "Id",
		"key":                     "Key",
		"last_uplink_received_at": "LastUplinkReceivedAt",
		"lo_ra_wan":               "LoRaWAN",
		"name":                    "Name",
		"rf_region":               "RfRegion",
		"tags":                    "Tags",
		"thing_arn":               "ThingArn",
		"thing_name":              "ThingName",
		"value":                   "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotwireless_wireless_gateway", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
