// Code generated by generators/resource/main.go; DO NOT EDIT.

package codegurureviewer

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
	registry.AddResourceTypeFactory("awscc_codegurureviewer_repository_association", repositoryAssociationResourceType)
}

// repositoryAssociationResourceType returns the Terraform awscc_codegurureviewer_repository_association resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::CodeGuruReviewer::RepositoryAssociation resource type.
func repositoryAssociationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"association_arn": {
			// Property: AssociationArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the repository association.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the repository association.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(0, 256)},
			Computed:    true,
		},
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of the S3 bucket associated with an associated S3 repository. It must start with `codeguru-reviewer-`.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(3, 63)},
			Optional:    true,
			Computed:    true,
			// BucketName is a force-new attribute.
		},
		"connection_arn": {
			// Property: ConnectionArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(0, 256)},
			Optional:    true,
			Computed:    true,
			// ConnectionArn is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the repository to be associated.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name of the repository to be associated.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Required:    true,
			// Name is a force-new attribute.
		},
		"owner": {
			// Property: Owner
			// CloudFormation resource type schema:
			// {
			//   "description": "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The owner of the repository. For a Bitbucket repository, this is the username for the account that owns the repository.",
			Type:        types.StringType,
			Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 100)},
			Optional:    true,
			Computed:    true,
			// Owner is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "The tags associated with a repository association.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
			//         "maxLength": 256,
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
			//   "maxItems": 50,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "The tags associated with a repository association.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(1, 128)},
						Required:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. The allowed characters across services are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.",
						Type:        types.StringType,
						Validators:  []tfsdk.AttributeValidator{validate.StringLenBetween(0, 256)},
						Required:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
			Computed: true,
			// Tags is a force-new attribute.
		},
		"type": {
			// Property: Type
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of repository to be associated.",
			//   "enum": [
			//     "CodeCommit",
			//     "Bitbucket",
			//     "GitHubEnterpriseServer",
			//     "S3Bucket"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of repository to be associated.",
			Type:        types.StringType,
			Required:    true,
			// Type is a force-new attribute.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "This resource schema represents the RepositoryAssociation resource in the Amazon CodeGuru Reviewer service.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeGuruReviewer::RepositoryAssociation").WithTerraformTypeName("awscc_codegurureviewer_repository_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_arn": "AssociationArn",
		"bucket_name":     "BucketName",
		"connection_arn":  "ConnectionArn",
		"key":             "Key",
		"name":            "Name",
		"owner":           "Owner",
		"tags":            "Tags",
		"type":            "Type",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_codegurureviewer_repository_association", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
