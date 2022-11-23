package outscale

import (
	"context"
	"fmt"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func dataSourceLBUTags() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLBUTagsRead,

		Schema: getDataSourceSchemas(getDSLBUTagsSchema()),
	}
}

func dataSourceLBUTagsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*Client).OSCAPI

	ename, nameOk := d.GetOk("load_balancer_names")
	if !nameOk {
		return fmt.Errorf("load_balancer_names is required")
	}

	names := ename.([]interface{})

	req := oscgo.ReadLoadBalancerTagsRequest{
		LoadBalancerNames: utils.InterfaceSliceToStringSlice(names),
	}

	var resp oscgo.ReadLoadBalancerTagsResponse
	var err error
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		rp, httpResp, err := conn.LoadBalancerApi.ReadLoadBalancerTags(
			context.Background()).
			ReadLoadBalancerTagsRequest(req).Execute()

		if err != nil {
			return utils.CheckThrottling(httpResp.StatusCode, err)
		}
		resp = rp
		return nil
	})

	if err != nil {
		return err
	}

	tags := *resp.Tags
	l := len(*resp.Tags)

	ta := make([]map[string]interface{}, l)
	for k1, v1 := range tags {
		t := make(map[string]interface{})
		t["key"] = v1.Key
		t["value"] = v1.Value
		t["load_balancer_name"] = v1.LoadBalancerName
		ta[k1] = t
	}

	d.Set("tags", ta)
	d.SetId(resource.UniqueId())
	return nil
}

func getDSLBUTagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"load_balancer_names": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"load_balancer_name": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"value": {
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
	}
}
