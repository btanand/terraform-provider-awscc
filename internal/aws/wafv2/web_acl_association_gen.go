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
	registry.AddResourceTypeFactory("awscc_wafv2_web_acl_association", webACLAssociationResourceType)
}

// webACLAssociationResourceType returns the Terraform awscc_wafv2_web_acl_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WAFv2::WebACLAssociation resource type.
func webACLAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"resource_arn": {
			// Property: ResourceArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(20, 2048)},
			Required:   true,
			// ResourceArn is a force-new attribute.
		},
		"web_acl_arn": {
			// Property: WebACLArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(20, 2048)},
			Required:   true,
			// WebACLArn is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Associates WebACL to Application Load Balancer, CloudFront or API Gateway.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WAFv2::WebACLAssociation").WithTerraformTypeName("awscc_wafv2_web_acl_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"resource_arn": "ResourceArn",
		"web_acl_arn":  "WebACLArn",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_wafv2_web_acl_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
