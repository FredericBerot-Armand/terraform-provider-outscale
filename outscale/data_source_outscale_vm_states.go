package outscale

import (
	"context"
	"errors"
	"fmt"
	"time"

	oscgo "github.com/outscale/osc-sdk-go/v2"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-outscale/utils"
)

func dataSourceVMStates() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceVMStatesRead,
		Schema: getVMStatesDataSourceSchema(),
	}
}

func getVMStatesDataSourceSchema() map[string]*schema.Schema {
	wholeSchema := map[string]*schema.Schema{
		"filter": dataSourceFiltersSchema(),
		"vm_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"vm_states": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: getVMStateAttrsSchema(),
			},
		},
		"request_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return wholeSchema
}

func dataSourceVMStatesRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*Client).OSCAPI

	filters, filtersOk := d.GetOk("filter")
	instanceIds, instanceIdsOk := d.GetOk("vm_ids")

	if !instanceIdsOk && !filtersOk {
		return errors.New("vm_id or filter must be set")
	}

	params := oscgo.ReadVmsStateRequest{}
	if filtersOk {
		params.SetFilters(buildDataSourceVMStateFilters(filters.(*schema.Set)))
	}
	if instanceIdsOk {
		filter := oscgo.FiltersVmsState{}
		filter.SetVmIds(utils.InterfaceSliceToStringSlice(instanceIds.([]interface{})))
		params.SetFilters(filter)
	}

	var resp oscgo.ReadVmsStateResponse
	var err error

	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		rp, httpResp, err := conn.VmApi.ReadVmsState(context.Background()).ReadVmsStateRequest(params).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp.StatusCode, err)
		}
		resp = rp
		return nil
	})

	if err != nil {
		return err
	}

	filteredStates := resp.GetVmStates()[:]

	if len(filteredStates) < 1 {
		return fmt.Errorf("your query returned no results, please change your search criteria and try again")
	}

	return statusDescriptionVMStatesAttributes(d, filteredStates)
}

func statusDescriptionVMStatesAttributes(d *schema.ResourceData, status []oscgo.VmStates) error {
	d.SetId(resource.UniqueId())

	states := make([]map[string]interface{}, len(status))

	for k, v := range status {
		state := make(map[string]interface{})

		setterFunc := func(key string, value interface{}) error {
			state[key] = value
			return nil
		}

		if err := statusDescriptionVMStateAttributes(setterFunc, &v); err != nil {
			return err
		}

		states[k] = state
	}

	return d.Set("vm_states", states)
}
