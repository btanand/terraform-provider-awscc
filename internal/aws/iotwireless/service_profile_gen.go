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
	registry.AddResourceTypeFactory("awscc_iotwireless_service_profile", serviceProfileResourceType)
}

// serviceProfileResourceType returns the Terraform awscc_iotwireless_service_profile resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTWireless::ServiceProfile resource type.
func serviceProfileResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Service profile Arn. Returned after successful create.",
			//   "type": "string"
			// }
			Description: "Service profile Arn. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Service profile Id. Returned after successful create.",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Service profile Id. Returned after successful create.",
			Type:        types.StringType,
			Computed:    true,
		},
		"lo_ra_wan": {
			// Property: LoRaWAN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AddGwMetadata": {
			//       "type": "boolean"
			//     },
			//     "ChannelMask": {
			//       "type": "string"
			//     },
			//     "DevStatusReqFreq": {
			//       "type": "integer"
			//     },
			//     "DlBucketSize": {
			//       "type": "integer"
			//     },
			//     "DlRate": {
			//       "type": "integer"
			//     },
			//     "DlRatePolicy": {
			//       "type": "string"
			//     },
			//     "DrMax": {
			//       "type": "integer"
			//     },
			//     "DrMin": {
			//       "type": "integer"
			//     },
			//     "HrAllowed": {
			//       "type": "boolean"
			//     },
			//     "MinGwDiversity": {
			//       "type": "integer"
			//     },
			//     "NwkGeoLoc": {
			//       "type": "boolean"
			//     },
			//     "PrAllowed": {
			//       "type": "boolean"
			//     },
			//     "RaAllowed": {
			//       "type": "boolean"
			//     },
			//     "ReportDevStatusBattery": {
			//       "type": "boolean"
			//     },
			//     "ReportDevStatusMargin": {
			//       "type": "boolean"
			//     },
			//     "TargetPer": {
			//       "type": "integer"
			//     },
			//     "UlBucketSize": {
			//       "type": "integer"
			//     },
			//     "UlRate": {
			//       "type": "integer"
			//     },
			//     "UlRatePolicy": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"add_gw_metadata": {
						// Property: AddGwMetadata
						Type:     types.BoolType,
						Optional: true,
					},
					"channel_mask": {
						// Property: ChannelMask
						Type:     types.StringType,
						Computed: true,
					},
					"dev_status_req_freq": {
						// Property: DevStatusReqFreq
						Type:     types.NumberType,
						Computed: true,
					},
					"dl_bucket_size": {
						// Property: DlBucketSize
						Type:     types.NumberType,
						Computed: true,
					},
					"dl_rate": {
						// Property: DlRate
						Type:     types.NumberType,
						Computed: true,
					},
					"dl_rate_policy": {
						// Property: DlRatePolicy
						Type:     types.StringType,
						Computed: true,
					},
					"dr_max": {
						// Property: DrMax
						Type:     types.NumberType,
						Computed: true,
					},
					"dr_min": {
						// Property: DrMin
						Type:     types.NumberType,
						Computed: true,
					},
					"hr_allowed": {
						// Property: HrAllowed
						Type:     types.BoolType,
						Computed: true,
					},
					"min_gw_diversity": {
						// Property: MinGwDiversity
						Type:     types.NumberType,
						Computed: true,
					},
					"nwk_geo_loc": {
						// Property: NwkGeoLoc
						Type:     types.BoolType,
						Computed: true,
					},
					"pr_allowed": {
						// Property: PrAllowed
						Type:     types.BoolType,
						Computed: true,
					},
					"ra_allowed": {
						// Property: RaAllowed
						Type:     types.BoolType,
						Computed: true,
					},
					"report_dev_status_battery": {
						// Property: ReportDevStatusBattery
						Type:     types.BoolType,
						Computed: true,
					},
					"report_dev_status_margin": {
						// Property: ReportDevStatusMargin
						Type:     types.BoolType,
						Computed: true,
					},
					"target_per": {
						// Property: TargetPer
						Type:     types.NumberType,
						Computed: true,
					},
					"ul_bucket_size": {
						// Property: UlBucketSize
						Type:     types.NumberType,
						Computed: true,
					},
					"ul_rate": {
						// Property: UlRate
						Type:     types.NumberType,
						Computed: true,
					},
					"ul_rate_policy": {
						// Property: UlRatePolicy
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of service profile",
			//   "maxLength": 256,
			//   "type": "string"
			// }
			Description: "Name of service profile",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the service profile.",
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
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of key-value pairs that contain metadata for the service profile.",
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
						Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 256)},
						Optional:   true,
					},
				},
				providertypes.SetNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
	}

	schema := tfsdk.Schema{
		Description: "An example resource schema demonstrating some basic constructs and validation rules.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTWireless::ServiceProfile").WithTerraformTypeName("awscc_iotwireless_service_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_gw_metadata":           "AddGwMetadata",
		"arn":                       "Arn",
		"channel_mask":              "ChannelMask",
		"dev_status_req_freq":       "DevStatusReqFreq",
		"dl_bucket_size":            "DlBucketSize",
		"dl_rate":                   "DlRate",
		"dl_rate_policy":            "DlRatePolicy",
		"dr_max":                    "DrMax",
		"dr_min":                    "DrMin",
		"hr_allowed":                "HrAllowed",
		"id":                        "Id",
		"key":                       "Key",
		"lo_ra_wan":                 "LoRaWAN",
		"min_gw_diversity":          "MinGwDiversity",
		"name":                      "Name",
		"nwk_geo_loc":               "NwkGeoLoc",
		"pr_allowed":                "PrAllowed",
		"ra_allowed":                "RaAllowed",
		"report_dev_status_battery": "ReportDevStatusBattery",
		"report_dev_status_margin":  "ReportDevStatusMargin",
		"tags":                      "Tags",
		"target_per":                "TargetPer",
		"ul_bucket_size":            "UlBucketSize",
		"ul_rate":                   "UlRate",
		"ul_rate_policy":            "UlRatePolicy",
		"value":                     "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotwireless_service_profile", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
