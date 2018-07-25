package outscale

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider

var testAccProvider *schema.Provider

//var isOAPI bool

func init() {
	testAccProvider = Provider().(*schema.Provider)

	testAccProviders = map[string]terraform.ResourceProvider{
		"outscale": testAccProvider,
	}

}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
}
