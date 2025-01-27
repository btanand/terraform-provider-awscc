// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_pcaconnectorad_service_principal_name", servicePrincipalNameResource)
}

// servicePrincipalNameResource returns the Terraform awscc_pcaconnectorad_service_principal_name resource.
// This Terraform resource corresponds to the CloudFormation AWS::PCAConnectorAD::ServicePrincipalName resource.
func servicePrincipalNameResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ConnectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DirectoryRegistrationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:directory-registration(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"directory_registration_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:directory-registration(\\/[\\w-]+)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
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
		Description: "Definition of AWS::PCAConnectorAD::ServicePrincipalName Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCAConnectorAD::ServicePrincipalName").WithTerraformTypeName("awscc_pcaconnectorad_service_principal_name")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"connector_arn":              "ConnectorArn",
		"directory_registration_arn": "DirectoryRegistrationArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
