// Code generated by generators/resource/main.go; DO NOT EDIT.

package servicecatalog_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSServiceCatalogCloudFormationProvisionedProduct_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ServiceCatalog::CloudFormationProvisionedProduct", "aws_servicecatalog_cloud_formation_provisioned_product", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
	})
}

func TestAccAWSServiceCatalogCloudFormationProvisionedProduct_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ServiceCatalog::CloudFormationProvisionedProduct", "aws_servicecatalog_cloud_formation_provisioned_product", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}