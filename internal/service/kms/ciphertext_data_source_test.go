package kms_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func TestAccDataSourceAwsKmsCiphertext_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t) },
		ErrorCheck: acctest.ErrorCheck(t, kms.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsKmsCiphertextConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_ciphertext.foo", "ciphertext_blob"),
				),
			},
		},
	})
}

func TestAccDataSourceAwsKmsCiphertext_validate(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t) },
		ErrorCheck: acctest.ErrorCheck(t, kms.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsKmsCiphertextConfig_validate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_ciphertext.foo", "ciphertext_blob"),
				),
			},
		},
	})
}

func TestAccDataSourceAwsKmsCiphertext_validate_withContext(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(t) },
		ErrorCheck: acctest.ErrorCheck(t, kms.EndpointsID),
		Providers:  acctest.Providers,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAwsKmsCiphertextConfig_validate_withContext,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.aws_kms_ciphertext.foo", "ciphertext_blob"),
				),
			},
		},
	})
}

const testAccDataSourceAwsKmsCiphertextConfig_basic = `
resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-basic"
  is_enabled  = true
}

data "aws_kms_ciphertext" "foo" {
  key_id = aws_kms_key.foo.key_id

  plaintext = "Super secret data"
}
`

const testAccDataSourceAwsKmsCiphertextConfig_validate = `
resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-validate"
  is_enabled  = true
}

data "aws_kms_ciphertext" "foo" {
  key_id = aws_kms_key.foo.key_id

  plaintext = "Super secret data"
}
`

const testAccDataSourceAwsKmsCiphertextConfig_validate_withContext = `
resource "aws_kms_key" "foo" {
  description = "tf-test-acc-data-source-aws-kms-ciphertext-validate-with-context"
  is_enabled  = true
}

data "aws_kms_ciphertext" "foo" {
  key_id = aws_kms_key.foo.key_id

  plaintext = "Super secret data"

  context = {
    name = "value"
  }
}
`