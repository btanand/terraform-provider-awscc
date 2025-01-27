// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediaconnect_bridge", bridgeResource)
}

// bridgeResource returns the Terraform awscc_mediaconnect_bridge resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::Bridge resource.
func bridgeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BridgeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the bridge.",
		//	  "type": "string"
		//	}
		"bridge_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the bridge.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BridgeState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATING",
		//	    "STANDBY",
		//	    "STARTING",
		//	    "DEPLOYING",
		//	    "ACTIVE",
		//	    "STOPPING",
		//	    "DELETING",
		//	    "DELETED",
		//	    "START_FAILED",
		//	    "START_PENDING",
		//	    "UPDATING"
		//	  ],
		//	  "type": "string"
		//	}
		"bridge_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EgressGatewayBridge
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "MaxBitrate": {
		//	      "description": "The maximum expected bitrate of the egress bridge.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaxBitrate"
		//	  ],
		//	  "type": "object"
		//	}
		"egress_gateway_bridge": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxBitrate
				"max_bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum expected bitrate of the egress bridge.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngressGatewayBridge
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "MaxBitrate": {
		//	      "description": "The maximum expected bitrate of the ingress bridge.",
		//	      "type": "integer"
		//	    },
		//	    "MaxOutputs": {
		//	      "description": "The maximum number of outputs on the ingress bridge.",
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "MaxBitrate",
		//	    "MaxOutputs"
		//	  ],
		//	  "type": "object"
		//	}
		"ingress_gateway_bridge": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxBitrate
				"max_bitrate": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum expected bitrate of the ingress bridge.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxOutputs
				"max_outputs": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of outputs on the ingress bridge.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the bridge.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the bridge.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Outputs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The outputs on this bridge.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The output of the bridge.",
		//	    "properties": {
		//	      "NetworkOutput": {
		//	        "additionalProperties": false,
		//	        "description": "The output of the bridge. A network output is delivered to your premises.",
		//	        "properties": {
		//	          "IpAddress": {
		//	            "description": "The network output IP Address.",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The network output name.",
		//	            "type": "string"
		//	          },
		//	          "NetworkName": {
		//	            "description": "The network output's gateway network name.",
		//	            "type": "string"
		//	          },
		//	          "Port": {
		//	            "description": "The network output port.",
		//	            "type": "integer"
		//	          },
		//	          "Protocol": {
		//	            "description": "The network output protocol.",
		//	            "enum": [
		//	              "rtp-fec",
		//	              "rtp",
		//	              "udp"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Ttl": {
		//	            "description": "The network output TTL.",
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Protocol",
		//	          "IpAddress",
		//	          "Port",
		//	          "NetworkName",
		//	          "Ttl"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"outputs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: NetworkOutput
					"network_output": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: IpAddress
							"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network output IP Address.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network output name.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: NetworkName
							"network_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network output's gateway network name.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Port
							"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The network output port.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Protocol
							"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network output protocol.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"rtp-fec",
										"rtp",
										"udp",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Ttl
							"ttl": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The network output TTL.",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The output of the bridge. A network output is delivered to your premises.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The outputs on this bridge.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 2),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlacementArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The placement Amazon Resource Number (ARN) of the bridge.",
		//	  "type": "string"
		//	}
		"placement_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The placement Amazon Resource Number (ARN) of the bridge.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceFailoverConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The settings for source failover.",
		//	  "properties": {
		//	    "FailoverMode": {
		//	      "description": "The type of failover you choose for this flow. FAILOVER allows switching between different streams.",
		//	      "enum": [
		//	        "FAILOVER"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "SourcePriority": {
		//	      "additionalProperties": false,
		//	      "description": "The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.",
		//	      "properties": {
		//	        "PrimarySource": {
		//	          "description": "The name of the source you choose as the primary source for this flow.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "State": {
		//	      "enum": [
		//	        "ENABLED",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "FailoverMode"
		//	  ],
		//	  "type": "object"
		//	}
		"source_failover_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FailoverMode
				"failover_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of failover you choose for this flow. FAILOVER allows switching between different streams.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"FAILOVER",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: SourcePriority
				"source_priority": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PrimarySource
						"primary_source": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the source you choose as the primary source for this flow.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: State
				"state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ENABLED",
							"DISABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The settings for source failover.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Sources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The sources on this bridge.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The bridge's source.",
		//	    "properties": {
		//	      "FlowSource": {
		//	        "additionalProperties": false,
		//	        "description": "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
		//	        "properties": {
		//	          "FlowArn": {
		//	            "description": "The ARN of the cloud flow used as a source of this bridge.",
		//	            "type": "string"
		//	          },
		//	          "FlowVpcInterfaceAttachment": {
		//	            "additionalProperties": false,
		//	            "description": "The name of the VPC interface attachment to use for this source.",
		//	            "properties": {
		//	              "VpcInterfaceName": {
		//	                "description": "The name of the VPC interface to use for this resource.",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the flow source.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "FlowArn"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "NetworkSource": {
		//	        "additionalProperties": false,
		//	        "description": "The source of the bridge. A network source originates at your premises.",
		//	        "properties": {
		//	          "MulticastIp": {
		//	            "description": "The network source multicast IP.",
		//	            "type": "string"
		//	          },
		//	          "Name": {
		//	            "description": "The name of the network source.",
		//	            "type": "string"
		//	          },
		//	          "NetworkName": {
		//	            "description": "The network source's gateway network name.",
		//	            "type": "string"
		//	          },
		//	          "Port": {
		//	            "description": "The network source port.",
		//	            "type": "integer"
		//	          },
		//	          "Protocol": {
		//	            "description": "The network source protocol.",
		//	            "enum": [
		//	              "rtp-fec",
		//	              "rtp",
		//	              "udp"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Protocol",
		//	          "MulticastIp",
		//	          "Port",
		//	          "NetworkName"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"sources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FlowSource
					"flow_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: FlowArn
							"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ARN of the cloud flow used as a source of this bridge.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: FlowVpcInterfaceAttachment
							"flow_vpc_interface_attachment": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: VpcInterfaceName
									"vpc_interface_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The name of the VPC interface to use for this resource.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The name of the VPC interface attachment to use for this source.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the flow source.",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: NetworkSource
					"network_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MulticastIp
							"multicast_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network source multicast IP.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the network source.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: NetworkName
							"network_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network source's gateway network name.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Port
							"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The network source port.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Protocol
							"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The network source protocol.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"rtp-fec",
										"rtp",
										"udp",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The source of the bridge. A network source originates at your premises.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The sources on this bridge.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 2),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::Bridge",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::Bridge").WithTerraformTypeName("awscc_mediaconnect_bridge")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bridge_arn":                    "BridgeArn",
		"bridge_state":                  "BridgeState",
		"egress_gateway_bridge":         "EgressGatewayBridge",
		"failover_mode":                 "FailoverMode",
		"flow_arn":                      "FlowArn",
		"flow_source":                   "FlowSource",
		"flow_vpc_interface_attachment": "FlowVpcInterfaceAttachment",
		"ingress_gateway_bridge":        "IngressGatewayBridge",
		"ip_address":                    "IpAddress",
		"max_bitrate":                   "MaxBitrate",
		"max_outputs":                   "MaxOutputs",
		"multicast_ip":                  "MulticastIp",
		"name":                          "Name",
		"network_name":                  "NetworkName",
		"network_output":                "NetworkOutput",
		"network_source":                "NetworkSource",
		"outputs":                       "Outputs",
		"placement_arn":                 "PlacementArn",
		"port":                          "Port",
		"primary_source":                "PrimarySource",
		"protocol":                      "Protocol",
		"source_failover_config":        "SourceFailoverConfig",
		"source_priority":               "SourcePriority",
		"sources":                       "Sources",
		"state":                         "State",
		"ttl":                           "Ttl",
		"vpc_interface_name":            "VpcInterfaceName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
