package outscale

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func dataSourceRouteTables() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouteTablesRead,

		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"route_table_id": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_tables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"net_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"route_table_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tags": dataSourceTagsSchema(),
						"route_propagating_virtual_gateways": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_gateway_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"routes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination_ip_range": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"destination_service_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gateway_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vm_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vm_account_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"net_peering_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"net_access_point_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"creation_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nic_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nat_service_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"link_route_tables": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"main": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"route_table_to_subnet_link_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"link_route_table_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"route_table_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouteTablesRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*Client).OSCAPI
	rtbID, rtbOk := d.GetOk("route_table_id")
	filter, filterOk := d.GetOk("filter")
	if !filterOk && !rtbOk {
		return fmt.Errorf("One of route_table_id or filters must be assigned")
	}

	params := oscgo.ReadRouteTablesRequest{
		Filters: &oscgo.FiltersRouteTable{},
	}

	if rtbOk {
		i := rtbID.([]interface{})
		in := make([]string, len(i))
		for k, v := range i {
			in[k] = v.(string)
		}
		filter := oscgo.FiltersRouteTable{}
		filter.SetRouteTableIds(in)
		params.SetFilters(filter)
	}

	if filterOk {
		params.Filters = buildDataSourceRouteTableFilters(filter.(*schema.Set))
	}

	var resp oscgo.ReadRouteTablesResponse
	var err error
	err = resource.Retry(60*time.Second, func() *resource.RetryError {
		rp, httpResp, err := conn.RouteTableApi.ReadRouteTables(context.Background()).ReadRouteTablesRequest(params).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp.StatusCode, err)
		}
		resp = rp
		return nil
	})

	var errString string
	if err != nil {
		errString = err.Error()
		return fmt.Errorf("[DEBUG] Error reading Internet Services (%s)", errString)
	}

	if err != nil {
		return err
	}

	rt := resp.GetRouteTables()
	if len(rt) == 0 {
		return fmt.Errorf("your query returned no results, please change your search criteria and try again")
	}

	routeTables := make([]map[string]interface{}, len(rt))

	for k, v := range rt {
		routeTable := make(map[string]interface{})
		routeTable["route_propagating_virtual_gateways"] = setPropagatingVirtualGateways(v.GetRoutePropagatingVirtualGateways())
		routeTable["route_table_id"] = v.GetRouteTableId()
		routeTable["net_id"] = v.GetNetId()
		routeTable["tags"] = tagsToMap(v.GetTags())
		routeTable["routes"] = setRoutes(v.GetRoutes())
		routeTable["link_route_tables"] = setLinkRouteTables(v.GetLinkRouteTables())
		routeTables[k] = routeTable
	}

	d.SetId(resource.UniqueId())

	return d.Set("route_tables", routeTables)
}
