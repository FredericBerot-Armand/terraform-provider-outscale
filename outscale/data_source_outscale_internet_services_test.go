package outscale

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccInternetServicesDatasource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccInternetServicesDatasourceConfig,
			},
		},
	})
}

const testAccInternetServicesDatasourceConfig = `
	resource "outscale_internet_service" "gateway" {}

	data "outscale_internet_services" "outscale_internet_services" {
		filter {
			name = "internet_service_id"
			values = ["${outscale_internet_service.gateway.id}"]
		}
	}
`
