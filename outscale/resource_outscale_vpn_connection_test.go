package outscale

import (
	"context"
	"fmt"
	"testing"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOutscaleVPNConnection_basic(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_vpn_connection.foo"

	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(1, 255))
	sharedKey := fmt.Sprintf("shared-key-%d", acctest.RandIntRange(1, 10000))
	sharedKeyUpdated := fmt.Sprintf("shared-key-%d", acctest.RandIntRange(1, 10000))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP, sharedKey, "169.254.254.0/30"),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "vgw_telemetries.#"),
					resource.TestCheckResourceAttrSet(resourceName, "vpn_options.#"),

					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.pre_shared_key", sharedKey),
					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.tunnel_inside_ip_range", "169.254.254.0/30"),
					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
				),
			},
			{
				Config: testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP, sharedKeyUpdated, "169.254.254.128/30"),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "vgw_telemetries.#"),
					resource.TestCheckResourceAttrSet(resourceName, "vpn_options.#"),

					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.pre_shared_key", sharedKey),
					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.tunnel_inside_ip_range", "169.254.254.128/30"),
					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
				),
			},
			{
				Config: testAccOutscaleVPNConnectionConfig(publicIP, sharedKey, false),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "static_routes_only"),
					resource.TestCheckResourceAttrSet(resourceName, "vgw_telemetries.#"),
					resource.TestCheckResourceAttrSet(resourceName, "vpn_options.#"),

					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.pre_shared_key", sharedKey),
					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
					resource.TestCheckResourceAttr(resourceName, "static_routes_only", "false"),
				),
			},
			{
				Config: testAccOutscaleVPNConnectionConfig(publicIP, sharedKeyUpdated, true),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "static_routes_only"),
					resource.TestCheckResourceAttrSet(resourceName, "vgw_telemetries.#"),
					resource.TestCheckResourceAttrSet(resourceName, "vpn_options.#"),

					//resource.TestCheckResourceAttr(resourceName, "vpn_options.0.pre_shared_key", sharedKeyUpdated),
					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
					resource.TestCheckResourceAttr(resourceName, "static_routes_only", "true"),
				),
			},
		},
	})
}

func TestAccOutscaleVPNConnection_withoutStaticRoutes(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_vpn_connection.foo"
	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(0, 255))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "outscale_vpn_connection.foo",
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP, "key", "169.254.254.128/30"),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),

					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
				),
			},
		},
	})
}

func TestAccOutscaleVPNConnection_withTags(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_vpn_connection.foo"

	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(1, 255))
	value := fmt.Sprintf("testacc-%s", acctest.RandString(5))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfigWithTags(publicIP, value),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "tags.#"),

					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "1"),
					//resource.TestCheckResourceAttr(resourceName, "tags.0.key", "Name"),
					//resource.TestCheckResourceAttr(resourceName, "tags.0.value", value),
				),
			},
		},
	})
}

func TestAccOutscaleVPNConnection_importBasic(t *testing.T) {
	t.Parallel()
	resourceName := "outscale_vpn_connection.foo"

	publicIP := fmt.Sprintf("172.0.0.%d", acctest.RandIntRange(1, 255))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: resourceName,
		Providers:     testAccProviders,
		CheckDestroy:  testAccOutscaleVPNConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleVPNConnectionConfig(publicIP, "key", true),
				Check: resource.ComposeTestCheckFunc(
					testAccOutscaleVPNConnectionExists(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "client_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "virtual_gateway_id"),
					resource.TestCheckResourceAttrSet(resourceName, "connection_type"),
					resource.TestCheckResourceAttrSet(resourceName, "static_routes_only"),
					resource.TestCheckResourceAttr(resourceName, "connection_type", "ipsec.1"),
					resource.TestCheckResourceAttr(resourceName, "static_routes_only", "true"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"request_id"},
			},
		},
	})
}

func testAccOutscaleVPNConnectionExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN Connection ID is set")
		}

		filter := oscgo.ReadVpnConnectionsRequest{
			Filters: &oscgo.FiltersVpnConnection{
				VpnConnectionIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.VpnConnectionApi.ReadVpnConnections(context.Background()).ReadVpnConnectionsRequest(filter).Execute()
		if err != nil || len(resp.GetVpnConnections()) < 1 {
			return fmt.Errorf("Outscale VPN Connection not found (%s)", rs.Primary.ID)
		}
		return nil
	}
}

func testAccOutscaleVPNConnectionDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*OutscaleClient).OSCAPI
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "outscale_vpn_connection" {
			continue
		}

		filter := oscgo.ReadVpnConnectionsRequest{
			Filters: &oscgo.FiltersVpnConnection{
				VpnConnectionIds: &[]string{rs.Primary.ID},
			},
		}

		resp, _, err := conn.VpnConnectionApi.ReadVpnConnections(context.Background()).ReadVpnConnectionsRequest(filter).Execute()
		if err != nil ||
			len(resp.GetVpnConnections()) > 0 && resp.GetVpnConnections()[0].GetState() != "deleted" {
			return fmt.Errorf("Outscale VPN Connection still exists (%s): %s", rs.Primary.ID, err)
		}
	}
	return nil
}

func testAccOutscaleVPNConnectionConfig(publicIP, sharedKey string, staticRoutesOnly bool) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "virtual_gateway" {
			connection_type = "ipsec.1"
		}

		resource "outscale_client_gateway" "customer_gateway" {
			bgp_asn         = 3
			public_ip       = "%s"
			connection_type = "ipsec.1"
		}

		resource "outscale_vpn_connection" "foo" {
			client_gateway_id  = "${outscale_client_gateway.customer_gateway.id}"
			virtual_gateway_id = "${outscale_virtual_gateway.virtual_gateway.id}"
			connection_type    = "ipsec.1"
			static_routes_only = "%t"

			vpn_options  {
				pre_shared_key                = "%s"
			}
		}
	`, publicIP, staticRoutesOnly, sharedKey)
}

func testAccOutscaleVPNConnectionConfigWithoutStaticRoutes(publicIP, sharedKey, ipRange string) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "virtual_gateway" {
			connection_type = "ipsec.1"
		}

		resource "outscale_client_gateway" "customer_gateway" {
			bgp_asn         = 3
			public_ip       = "%s"
			connection_type = "ipsec.1"
		}

		resource "outscale_vpn_connection" "foo" {
			client_gateway_id  = "${outscale_client_gateway.customer_gateway.id}"
			virtual_gateway_id = "${outscale_virtual_gateway.virtual_gateway.id}"
			connection_type    = "ipsec.1"

			vpn_options  {
				pre_shared_key                = "%s"
				tunnel_inside_ip_range        = "%s"
			}
		}
	`, publicIP, sharedKey, ipRange)
}

func testAccOutscaleVPNConnectionConfigWithTags(publicIP, value string) string {
	return fmt.Sprintf(`
		resource "outscale_virtual_gateway" "virtual_gateway" {
			connection_type = "ipsec.1"
		}

		resource "outscale_client_gateway" "customer_gateway" {
			bgp_asn         = 3
			public_ip       = "%s"
			connection_type = "ipsec.1"
		}

		resource "outscale_vpn_connection" "foo" {
			client_gateway_id  = "${outscale_client_gateway.customer_gateway.id}"
			virtual_gateway_id = "${outscale_virtual_gateway.virtual_gateway.id}"
			connection_type    = "ipsec.1"
			static_routes_only = true


			tags {
				key   = "Name"
				value = "%s"
			}
		}
	`, publicIP, value)
}
