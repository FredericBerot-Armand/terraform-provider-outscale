package outscale

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceOutscaleOAPIPublicIP(t *testing.T) {
	o := os.Getenv("OUTSCALE_OAPI")

	oapi, err := strconv.ParseBool(o)
	if err != nil {
		oapi = false
	}

	if !oapi {
		t.Skip()
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceOutscaleOAPIPublicIPConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceOutscaleOAPIPublicIPCheck("data.outscale_public_ip.by_reservation_id"),
					testAccDataSourceOutscaleOAPIPublicIPCheck("data.outscale_public_ip.by_public_ip"),
				),
			},
		},
	})
}

func testAccDataSourceOutscaleOAPIPublicIPCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		fmt.Printf("\n[DEBUG] TEST RS %s \n", s.RootModule().Resources)

		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		eipRs, ok := s.RootModule().Resources["outscale_public_ip.test"]
		if !ok {
			return fmt.Errorf("can't find outscale_public_ip.test in state")
		}

		attr := rs.Primary.Attributes

		if attr["reservation_id"] != eipRs.Primary.Attributes["reservation_id"] {
			return fmt.Errorf(
				"reservation_id is %s; want %s",
				attr["reservation_id"],
				eipRs.Primary.Attributes["reservation_id"],
			)
		}

		if attr["public_ip"] != eipRs.Primary.Attributes["public_ip"] {
			return fmt.Errorf(
				"public_ip is %s; want %s",
				attr["public_ip"],
				eipRs.Primary.Attributes["public_ip"],
			)
		}

		return nil
	}
}

const testAccDataSourceOutscaleOAPIPublicIPConfig = `
resource "outscale_public_ip" "test" {}

data "outscale_public_ip" "by_reservation_id" {
  reservation_id = "${outscale_public_ip.test.reservation_id}"
}
data "outscale_public_ip" "by_public_ip" {
  public_ip = "${outscale_public_ip.test.public_ip}"
}
`
