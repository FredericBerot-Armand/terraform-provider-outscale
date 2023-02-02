package outscale

import (
	"context"
	"fmt"
	"net/http"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/terraform-providers/terraform-provider-outscale/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceOutscaleOAPiKeyPairsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI
	req := oscgo.ReadKeypairsRequest{}

	if filters, filtersOk := d.GetOk("filter"); filtersOk {
		req.SetFilters(buildOutscaleOAPIKeyPairsDataSourceFilters(filters.(*schema.Set)))
	}

	var resp oscgo.ReadKeypairsResponse
	var statusCode int
	err := resource.Retry(120*time.Second, func() *resource.RetryError {
		var err error
		rp, httpResp, err := conn.KeypairApi.ReadKeypairs(context.Background()).ReadKeypairsRequest(req).Execute()

		if err != nil {
			return utils.CheckThrottling(httpResp, err)
		}
		resp = rp
		statusCode = httpResp.StatusCode
		return nil
	})

	var errString string

	if err != nil {
		if statusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		errString = err.Error()

		return fmt.Errorf("Error retrieving OAPIKeyPair: %s", errString)
	}

	result_len := len(resp.GetKeypairs())
	if result_len == 0 {
		return fmt.Errorf("your query returned no results, please change your search criteria and try again")
	}

	d.SetId(resource.UniqueId())

	keypairs := make([]map[string]interface{}, result_len)
	for k, v := range resp.GetKeypairs() {
		keypair := make(map[string]interface{})
		if v.GetKeypairName() != "" {
			keypair["keypair_name"] = v.GetKeypairName()
		}
		if v.GetKeypairFingerprint() != "" {
			keypair["keypair_fingerprint"] = v.GetKeypairFingerprint()
		}
		keypairs[k] = keypair
	}

	return d.Set("keypairs", keypairs)
}

func datasourceOutscaleOAPIKeyPairs() *schema.Resource {
	return &schema.Resource{
		Read: datasourceOutscaleOAPiKeyPairsRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			// Attributes
			"keypairs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypair_fingerprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"keypair_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
