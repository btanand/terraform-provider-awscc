// Code generated by generators/resource/main.go; DO NOT EDIT.

package lookoutmetrics

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
	registry.AddResourceTypeFactory("awscc_lookoutmetrics_anomaly_detector", anomalyDetectorResourceType)
}

// anomalyDetectorResourceType returns the Terraform awscc_lookoutmetrics_anomaly_detector resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::LookoutMetrics::AnomalyDetector resource type.
func anomalyDetectorResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"anomaly_detector_config": {
			// Property: AnomalyDetectorConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AnomalyDetectorFrequency": {
			//       "description": "Frequency of anomaly detection",
			//       "enum": [
			//         "PT5M",
			//         "PT10M",
			//         "PT1H",
			//         "P1D"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AnomalyDetectorFrequency"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"anomaly_detector_frequency": {
						// Property: AnomalyDetectorFrequency
						Description: "Frequency of anomaly detection",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PT5M",
								"PT10M",
								"PT1H",
								"P1D",
							}),
						},
					},
				},
			),
			Required: true,
		},
		"anomaly_detector_description": {
			// Property: AnomalyDetectorDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the AnomalyDetector.",
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A description for the AnomalyDetector.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 256),
			},
		},
		"anomaly_detector_name": {
			// Property: AnomalyDetectorName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for the Amazon Lookout for Metrics Anomaly Detector",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name for the Amazon Lookout for Metrics Anomaly Detector",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(), // AnomalyDetectorName is a force-new property.
			},
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 256,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"kms_key_arn": {
			// Property: KmsKeyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "KMS key used to encrypt the AnomalyDetector data",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "KMS key used to encrypt the AnomalyDetector data",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
		},
		"metric_set_list": {
			// Property: MetricSetList
			// CloudFormation resource type schema:
			// {
			//   "description": "List of metric sets for anomaly detection",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DimensionList": {
			//         "description": "Dimensions for this MetricSet.",
			//         "insertionOrder": false,
			//         "items": {
			//           "description": "Name of a column in the data.",
			//           "maxLength": 63,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         },
			//         "minItems": 0,
			//         "type": "array"
			//       },
			//       "MetricList": {
			//         "description": "Metrics captured by this MetricSet.",
			//         "insertionOrder": false,
			//         "items": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "AggregationFunction": {
			//               "description": "Operator used to aggregate metric values",
			//               "enum": [
			//                 "AVG",
			//                 "SUM"
			//               ],
			//               "type": "string"
			//             },
			//             "MetricName": {
			//               "description": "Name of a column in the data.",
			//               "maxLength": 63,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             },
			//             "Namespace": {
			//               "maxLength": 255,
			//               "minLength": 1,
			//               "pattern": "",
			//               "type": "string"
			//             }
			//           },
			//           "required": [
			//             "MetricName",
			//             "AggregationFunction"
			//           ],
			//           "type": "object"
			//         },
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "MetricSetDescription": {
			//         "description": "A description for the MetricSet.",
			//         "maxLength": 256,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "MetricSetFrequency": {
			//         "description": "A frequency period to aggregate the data",
			//         "enum": [
			//           "PT5M",
			//           "PT10M",
			//           "PT1H",
			//           "P1D"
			//         ],
			//         "type": "string"
			//       },
			//       "MetricSetName": {
			//         "description": "The name of the MetricSet.",
			//         "maxLength": 63,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "MetricSource": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AppFlowConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "FlowName": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "RoleArn",
			//               "FlowName"
			//             ],
			//             "type": "object"
			//           },
			//           "CloudwatchConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "RoleArn"
			//             ],
			//             "type": "object"
			//           },
			//           "RDSSourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DBInstanceIdentifier": {
			//                 "maxLength": 63,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseHost": {
			//                 "maxLength": 253,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseName": {
			//                 "maxLength": 64,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabasePort": {
			//                 "maximum": 65535,
			//                 "minimum": 1,
			//                 "type": "integer"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "SecretManagerArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "TableName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "VpcConfiguration": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "SecurityGroupIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "minLength": 1,
			//                       "pattern": "",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   },
			//                   "SubnetIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "pattern": "",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   }
			//                 },
			//                 "required": [
			//                   "SubnetIdList",
			//                   "SecurityGroupIdList"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "required": [
			//               "DBInstanceIdentifier",
			//               "DatabaseHost",
			//               "DatabasePort",
			//               "SecretManagerArn",
			//               "DatabaseName",
			//               "TableName",
			//               "RoleArn",
			//               "VpcConfiguration"
			//             ],
			//             "type": "object"
			//           },
			//           "RedshiftSourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "ClusterIdentifier": {
			//                 "maxLength": 63,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseHost": {
			//                 "maxLength": 253,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabaseName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "DatabasePort": {
			//                 "maximum": 65535,
			//                 "minimum": 1,
			//                 "type": "integer"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "SecretManagerArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "TableName": {
			//                 "maxLength": 100,
			//                 "minLength": 1,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "VpcConfiguration": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "SecurityGroupIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "minLength": 1,
			//                       "pattern": "",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   },
			//                   "SubnetIdList": {
			//                     "items": {
			//                       "maxLength": 255,
			//                       "pattern": "",
			//                       "type": "string"
			//                     },
			//                     "type": "array"
			//                   }
			//                 },
			//                 "required": [
			//                   "SubnetIdList",
			//                   "SecurityGroupIdList"
			//                 ],
			//                 "type": "object"
			//               }
			//             },
			//             "required": [
			//               "ClusterIdentifier",
			//               "DatabaseHost",
			//               "DatabasePort",
			//               "SecretManagerArn",
			//               "DatabaseName",
			//               "TableName",
			//               "RoleArn",
			//               "VpcConfiguration"
			//             ],
			//             "type": "object"
			//           },
			//           "S3SourceConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "FileFormatDescriptor": {
			//                 "additionalProperties": false,
			//                 "properties": {
			//                   "CsvFormatDescriptor": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Charset": {
			//                         "maxLength": 63,
			//                         "pattern": "",
			//                         "type": "string"
			//                       },
			//                       "ContainsHeader": {
			//                         "type": "boolean"
			//                       },
			//                       "Delimiter": {
			//                         "maxLength": 1,
			//                         "pattern": "",
			//                         "type": "string"
			//                       },
			//                       "FileCompression": {
			//                         "enum": [
			//                           "NONE",
			//                           "GZIP"
			//                         ],
			//                         "type": "string"
			//                       },
			//                       "HeaderList": {
			//                         "items": {
			//                           "description": "Name of a column in the data.",
			//                           "maxLength": 63,
			//                           "minLength": 1,
			//                           "pattern": "",
			//                           "type": "string"
			//                         },
			//                         "type": "array"
			//                       },
			//                       "QuoteSymbol": {
			//                         "maxLength": 1,
			//                         "pattern": "",
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "JsonFormatDescriptor": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "Charset": {
			//                         "maxLength": 63,
			//                         "pattern": "",
			//                         "type": "string"
			//                       },
			//                       "FileCompression": {
			//                         "enum": [
			//                           "NONE",
			//                           "GZIP"
			//                         ],
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   }
			//                 },
			//                 "type": "object"
			//               },
			//               "HistoricalDataPathList": {
			//                 "items": {
			//                   "maxLength": 1024,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "maxItems": 1,
			//                 "minItems": 1,
			//                 "type": "array"
			//               },
			//               "RoleArn": {
			//                 "maxLength": 256,
			//                 "pattern": "",
			//                 "type": "string"
			//               },
			//               "TemplatedPathList": {
			//                 "items": {
			//                   "maxLength": 1024,
			//                   "pattern": "",
			//                   "type": "string"
			//                 },
			//                 "maxItems": 1,
			//                 "minItems": 1,
			//                 "type": "array"
			//               }
			//             },
			//             "required": [
			//               "RoleArn",
			//               "FileFormatDescriptor"
			//             ],
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Offset": {
			//         "description": "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
			//         "maximum": 432000,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "TimestampColumn": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ColumnFormat": {
			//             "description": "A timestamp format for the timestamps in the dataset",
			//             "maxLength": 63,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "ColumnName": {
			//             "description": "Name of a column in the data.",
			//             "maxLength": 63,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Timezone": {
			//         "maxLength": 60,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MetricSetName",
			//       "MetricList",
			//       "MetricSource"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 1,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "List of metric sets for anomaly detection",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"dimension_list": {
						// Property: DimensionList
						Description: "Dimensions for this MetricSet.",
						Type:        types.ListType{ElemType: types.StringType},
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtLeast(0),
						},
					},
					"metric_list": {
						// Property: MetricList
						Description: "Metrics captured by this MetricSet.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"aggregation_function": {
									// Property: AggregationFunction
									Description: "Operator used to aggregate metric values",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"AVG",
											"SUM",
										}),
									},
								},
								"metric_name": {
									// Property: MetricName
									Description: "Name of a column in the data.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 63),
									},
								},
								"namespace": {
									// Property: Namespace
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 255),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{
								MinItems: 1,
							},
						),
						Required: true,
					},
					"metric_set_description": {
						// Property: MetricSetDescription
						Description: "A description for the MetricSet.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
					"metric_set_frequency": {
						// Property: MetricSetFrequency
						Description: "A frequency period to aggregate the data",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PT5M",
								"PT10M",
								"PT1H",
								"P1D",
							}),
						},
					},
					"metric_set_name": {
						// Property: MetricSetName
						Description: "The name of the MetricSet.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 63),
						},
					},
					"metric_source": {
						// Property: MetricSource
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"app_flow_config": {
									// Property: AppFlowConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"flow_name": {
												// Property: FlowName
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
										},
									),
									Optional: true,
								},
								"cloudwatch_config": {
									// Property: CloudwatchConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
										},
									),
									Optional: true,
								},
								"rds_source_config": {
									// Property: RDSSourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"db_instance_identifier": {
												// Property: DBInstanceIdentifier
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 63),
												},
											},
											"database_host": {
												// Property: DatabaseHost
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 253),
												},
											},
											"database_name": {
												// Property: DatabaseName
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 64),
												},
											},
											"database_port": {
												// Property: DatabasePort
												Type:     types.NumberType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.IntBetween(1, 65535),
												},
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"secret_manager_arn": {
												// Property: SecretManagerArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"table_name": {
												// Property: TableName
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
											"vpc_configuration": {
												// Property: VpcConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"security_group_id_list": {
															// Property: SecurityGroupIdList
															Type:     types.ListType{ElemType: types.StringType},
															Required: true,
														},
														"subnet_id_list": {
															// Property: SubnetIdList
															Type:     types.ListType{ElemType: types.StringType},
															Required: true,
														},
													},
												),
												Required: true,
											},
										},
									),
									Optional: true,
								},
								"redshift_source_config": {
									// Property: RedshiftSourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cluster_identifier": {
												// Property: ClusterIdentifier
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 63),
												},
											},
											"database_host": {
												// Property: DatabaseHost
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 253),
												},
											},
											"database_name": {
												// Property: DatabaseName
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
											"database_port": {
												// Property: DatabasePort
												Type:     types.NumberType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.IntBetween(1, 65535),
												},
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"secret_manager_arn": {
												// Property: SecretManagerArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"table_name": {
												// Property: TableName
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(1, 100),
												},
											},
											"vpc_configuration": {
												// Property: VpcConfiguration
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"security_group_id_list": {
															// Property: SecurityGroupIdList
															Type:     types.ListType{ElemType: types.StringType},
															Required: true,
														},
														"subnet_id_list": {
															// Property: SubnetIdList
															Type:     types.ListType{ElemType: types.StringType},
															Required: true,
														},
													},
												),
												Required: true,
											},
										},
									),
									Optional: true,
								},
								"s3_source_config": {
									// Property: S3SourceConfig
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"file_format_descriptor": {
												// Property: FileFormatDescriptor
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"csv_format_descriptor": {
															// Property: CsvFormatDescriptor
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"charset": {
																		// Property: Charset
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringLenBetween(0, 63),
																		},
																	},
																	"contains_header": {
																		// Property: ContainsHeader
																		Type:     types.BoolType,
																		Optional: true,
																	},
																	"delimiter": {
																		// Property: Delimiter
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringLenBetween(0, 1),
																		},
																	},
																	"file_compression": {
																		// Property: FileCompression
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringInSlice([]string{
																				"NONE",
																				"GZIP",
																			}),
																		},
																	},
																	"header_list": {
																		// Property: HeaderList
																		Type:     types.ListType{ElemType: types.StringType},
																		Optional: true,
																	},
																	"quote_symbol": {
																		// Property: QuoteSymbol
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringLenBetween(0, 1),
																		},
																	},
																},
															),
															Optional: true,
														},
														"json_format_descriptor": {
															// Property: JsonFormatDescriptor
															Attributes: tfsdk.SingleNestedAttributes(
																map[string]tfsdk.Attribute{
																	"charset": {
																		// Property: Charset
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringLenBetween(0, 63),
																		},
																	},
																	"file_compression": {
																		// Property: FileCompression
																		Type:     types.StringType,
																		Optional: true,
																		Validators: []tfsdk.AttributeValidator{
																			validate.StringInSlice([]string{
																				"NONE",
																				"GZIP",
																			}),
																		},
																	},
																},
															),
															Optional: true,
														},
													},
												),
												Required: true,
											},
											"historical_data_path_list": {
												// Property: HistoricalDataPathList
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenBetween(1, 1),
												},
											},
											"role_arn": {
												// Property: RoleArn
												Type:     types.StringType,
												Required: true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringLenBetween(0, 256),
												},
											},
											"templated_path_list": {
												// Property: TemplatedPathList
												Type:     types.ListType{ElemType: types.StringType},
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenBetween(1, 1),
												},
											},
										},
									),
									Optional: true,
								},
							},
						),
						Required: true,
					},
					"offset": {
						// Property: Offset
						Description: "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
						Type:        types.NumberType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 432000),
						},
					},
					"timestamp_column": {
						// Property: TimestampColumn
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"column_format": {
									// Property: ColumnFormat
									Description: "A timestamp format for the timestamps in the dataset",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(0, 63),
									},
								},
								"column_name": {
									// Property: ColumnName
									Description: "Name of a column in the data.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 63),
									},
								},
							},
						),
						Optional: true,
					},
					"timezone": {
						// Property: Timezone
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 60),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 1,
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "An Amazon Lookout for Metrics Detector",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::LookoutMetrics::AnomalyDetector").WithTerraformTypeName("awscc_lookoutmetrics_anomaly_detector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregation_function":         "AggregationFunction",
		"anomaly_detector_config":      "AnomalyDetectorConfig",
		"anomaly_detector_description": "AnomalyDetectorDescription",
		"anomaly_detector_frequency":   "AnomalyDetectorFrequency",
		"anomaly_detector_name":        "AnomalyDetectorName",
		"app_flow_config":              "AppFlowConfig",
		"arn":                          "Arn",
		"charset":                      "Charset",
		"cloudwatch_config":            "CloudwatchConfig",
		"cluster_identifier":           "ClusterIdentifier",
		"column_format":                "ColumnFormat",
		"column_name":                  "ColumnName",
		"contains_header":              "ContainsHeader",
		"csv_format_descriptor":        "CsvFormatDescriptor",
		"database_host":                "DatabaseHost",
		"database_name":                "DatabaseName",
		"database_port":                "DatabasePort",
		"db_instance_identifier":       "DBInstanceIdentifier",
		"delimiter":                    "Delimiter",
		"dimension_list":               "DimensionList",
		"file_compression":             "FileCompression",
		"file_format_descriptor":       "FileFormatDescriptor",
		"flow_name":                    "FlowName",
		"header_list":                  "HeaderList",
		"historical_data_path_list":    "HistoricalDataPathList",
		"json_format_descriptor":       "JsonFormatDescriptor",
		"kms_key_arn":                  "KmsKeyArn",
		"metric_list":                  "MetricList",
		"metric_name":                  "MetricName",
		"metric_set_description":       "MetricSetDescription",
		"metric_set_frequency":         "MetricSetFrequency",
		"metric_set_list":              "MetricSetList",
		"metric_set_name":              "MetricSetName",
		"metric_source":                "MetricSource",
		"namespace":                    "Namespace",
		"offset":                       "Offset",
		"quote_symbol":                 "QuoteSymbol",
		"rds_source_config":            "RDSSourceConfig",
		"redshift_source_config":       "RedshiftSourceConfig",
		"role_arn":                     "RoleArn",
		"s3_source_config":             "S3SourceConfig",
		"secret_manager_arn":           "SecretManagerArn",
		"security_group_id_list":       "SecurityGroupIdList",
		"subnet_id_list":               "SubnetIdList",
		"table_name":                   "TableName",
		"templated_path_list":          "TemplatedPathList",
		"timestamp_column":             "TimestampColumn",
		"timezone":                     "Timezone",
		"vpc_configuration":            "VpcConfiguration",
	})

	opts = opts.WithCreateTimeoutInMinutes(15).WithDeleteTimeoutInMinutes(15)

	opts = opts.WithUpdateTimeoutInMinutes(15)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_lookoutmetrics_anomaly_detector", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}