package outscale

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccImagesDataSource_Instance(t *testing.T) {
	t.Parallel()
	omi := os.Getenv("OUTSCALE_IMAGEID")
	region := os.Getenv("OUTSCALE_REGION")
	imageName := fmt.Sprintf("image-test-%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckImagesDataSourceConfig(omi, "tinav4.c2r2p2", region, imageName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesDataSourceID("data.outscale_images.nat_ami"),
				),
			},
		},
	})
}

func testAccCheckImagesDataSourceID(n string) resource.TestCheckFunc {
	// Wait for IAM role
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find AMI data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("AMI data source ID not set")
		}
		return nil
	}
}

func testAccCheckImagesDataSourceConfig(omi, vmType, region, imageName string) string {
	return fmt.Sprintf(`
		resource "outscale_vm" "basic_one" {
			image_id			           = "%[1]s"
			vm_type                  = "%[2]s"
			keypair_name	           = "terraform-basic"
			placement_subregion_name = "%[3]sa"
		}

		resource "outscale_vm" "basic_two" {
			image_id			           = "%[1]s"
			vm_type                  = "%[2]s"
			keypair_name	           = "terraform-basic"
			placement_subregion_name = "%[3]sa"
		}

		resource "outscale_image" "image_one" {
			image_name = "%[4]s-one"
			vm_id = "${outscale_vm.basic_one.id}"
		}

		resource "outscale_image" "image_two" {
			image_name = "%[4]s-two"
			vm_id = "${outscale_vm.basic_two.id}"
		}

		data "outscale_images" "nat_ami" {
			filter {
				name = "image_ids"
				values = ["${outscale_image.image_one.id}", "${outscale_image.image_two.id}"]
			}
		}
	`, omi, vmType, region, imageName)
}
