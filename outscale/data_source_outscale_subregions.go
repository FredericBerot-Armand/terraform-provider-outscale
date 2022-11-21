package outscale

import (
	"context"
	"log"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/spf13/cast"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func dataSourceSubregions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubregionsRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			// Computed values.
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subregions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"location_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subregion_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSubregionsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*Client).OSCAPI

	filters, filtersOk := d.GetOk("filter")

	filtersReq := &oscgo.FiltersSubregion{}
	if filtersOk {
		filtersReq = buildDataSourceSubregionsFilters(filters.(*schema.Set))
	}

	req := oscgo.ReadSubregionsRequest{Filters: filtersReq}

	var resp oscgo.ReadSubregionsResponse
	var err error
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		rp, httpResp, err := conn.SubregionApi.ReadSubregions(context.Background()).ReadSubregionsRequest(req).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp.StatusCode, err)
		}
		resp = rp
		return nil
	})

	if err != nil {
		return err
	}

	subregions := resp.GetSubregions()

	return resourceDataAttrSetter(d, func(set AttributeSetter) error {
		d.SetId(resource.UniqueId())

		subs := make([]map[string]interface{}, len(subregions))
		for i, subregion := range subregions {
			subs[i] = map[string]interface{}{
				"location_code":  subregion.GetLocationCode(),
				"subregion_name": subregion.GetSubregionName(),
				"region_name":    subregion.GetRegionName(),
				"state":          subregion.GetState(),
			}
		}

		return set("subregions", subs)
	})
}

func buildDataSourceSubregionsFilters(set *schema.Set) *oscgo.FiltersSubregion {
	filters := &oscgo.FiltersSubregion{}
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		var filterValues []string

		for _, e := range m["values"].([]interface{}) {
			filterValues = append(filterValues, cast.ToString(e))
		}

		switch name := m["name"].(string); name {
		case "subregion_names":
			filters.SetSubregionNames(filterValues)
		default:
			log.Printf("[Debug] Unknown Filter Name: %s.", name)
		}
	}
	return filters
}
