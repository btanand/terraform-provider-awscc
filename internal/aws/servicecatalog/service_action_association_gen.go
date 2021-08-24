// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog

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
	registry.AddResourceTypeFactory("awscc_servicecatalog_service_action_association", serviceActionAssociationResourceType)
}

// serviceActionAssociationResourceType returns the Terraform awscc_servicecatalog_service_action_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ServiceCatalog::ServiceActionAssociation resource type.
func serviceActionAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"product_id": {
			// Property: ProductId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Required:   true,
			// ProductId is a force-new attribute.
		},
		"provisioning_artifact_id": {
			// Property: ProvisioningArtifactId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Required:   true,
			// ProvisioningArtifactId is a force-new attribute.
		},
		"service_action_id": {
			// Property: ServiceActionId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:       types.StringType,
			Validators: []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Required:   true,
			// ServiceActionId is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ServiceCatalog::ServiceActionAssociation").WithTerraformTypeName("awscc_servicecatalog_service_action_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"product_id":               "ProductId",
		"provisioning_artifact_id": "ProvisioningArtifactId",
		"service_action_id":        "ServiceActionId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_servicecatalog_service_action_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
