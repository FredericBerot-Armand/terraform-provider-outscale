package outscale

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/openlyinc/pointy"
	oscgo "github.com/outscale/osc-sdk-go/v2"
	"github.com/spf13/cast"
	"github.com/terraform-providers/terraform-provider-outscale/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceOutscaleOAPIOutboundRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleOAPIOutboundRuleCreate,
		Read:   resourceOutscaleOAPIOutboundRuleRead,
		Delete: resourceOutscaleOAPIOutboundRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceOutscaleOAPISecurityGroupRuleImportState,
		},
		Schema: map[string]*schema.Schema{
			"flow": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Inbound", "Outbound"}, false),
			},
			"from_port_range": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"ip_protocol": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"rules", "security_group_name_to_link"},
			},
			"ip_range": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"rules": getRulesSchema(false),
			"security_group_account_id_to_link": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"security_group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group_name_to_link": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				ConflictsWith: []string{"ip_protocol", "rules"},
			},
			"to_port_range": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"security_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"net_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inbound_rules":  getRulesSchema(true),
			"outbound_rules": getRulesSchema(true),
			"request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceOutscaleOAPIOutboundRuleCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	req := oscgo.CreateSecurityGroupRuleRequest{
		Flow:            d.Get("flow").(string),
		SecurityGroupId: d.Get("security_group_id").(string),
		Rules:           expandRules(d),
	}

	if v, ok := d.GetOkExists("from_port_range"); ok {
		req.FromPortRange = pointy.Int32(cast.ToInt32(v))
	}
	if v, ok := d.GetOkExists("to_port_range"); ok {
		req.ToPortRange = pointy.Int32(cast.ToInt32(v))
	}
	if v, ok := d.GetOk("ip_protocol"); ok {
		req.IpProtocol = pointy.String(v.(string))
	}
	if v, ok := d.GetOk("ip_range"); ok {
		req.IpRange = pointy.String(v.(string))
	}
	if v, ok := d.GetOk("security_group_account_id_to_link"); ok {
		req.SecurityGroupAccountIdToLink = pointy.String(v.(string))
	}
	if v, ok := d.GetOk("security_group_name_to_link"); ok {
		req.SecurityGroupNameToLink = pointy.String(v.(string))
	}

	var err error
	var resp oscgo.CreateSecurityGroupRuleResponse
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		rp, httpResp, err := conn.SecurityGroupRuleApi.CreateSecurityGroupRule(context.Background()).CreateSecurityGroupRuleRequest(req).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp, err)
		}
		resp = rp
		return nil
	})

	if err != nil {
		return fmt.Errorf(
			"Error authorizing security group rule type: %s", utils.GetErrorResponse(err))
	}

	d.SetId(*resp.GetSecurityGroup().SecurityGroupId)

	return resourceOutscaleOAPIOutboundRuleRead(d, meta)
}

func resourceOutscaleOAPIOutboundRuleRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	sg, _, err := readSecurityGroups(conn, d.Id())
	if err != nil {
		return err
	}
	if sg == nil {
		utils.LogManuallyDeleted("SecurityGroupeRule", d.Id())
		d.SetId("")
		return nil
	}

	if err := d.Set("inbound_rules", flattenRules(sg.GetInboundRules())); err != nil {
		return fmt.Errorf("error setting `inbound_rules` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("outbound_rules", flattenRules(sg.GetOutboundRules())); err != nil {
		return fmt.Errorf("error setting `outbound_rules` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("security_group_name", sg.GetSecurityGroupName()); err != nil {
		return fmt.Errorf("error setting `security_group_name` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("net_id", sg.GetNetId()); err != nil {
		return fmt.Errorf("error setting `net_id` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}

	return nil
}

func resourceOutscaleOAPIOutboundRuleDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OSCAPI

	req := oscgo.DeleteSecurityGroupRuleRequest{
		Flow:            d.Get("flow").(string),
		SecurityGroupId: d.Get("security_group_id").(string),
		Rules:           expandRules(d),
	}

	if v, ok := d.GetOkExists("from_port_range"); ok {
		req.FromPortRange = pointy.Int32(cast.ToInt32(v))
	}
	if v, ok := d.GetOkExists("to_port_range"); ok {
		req.ToPortRange = pointy.Int32(cast.ToInt32(v))
	}
	if v, ok := d.GetOk("ip_protocol"); ok {
		req.IpProtocol = pointy.String(v.(string))
	}
	if v, ok := d.GetOk("ip_range"); ok {
		req.IpRange = pointy.String(v.(string))
	}

	err := resource.Retry(5*time.Minute, func() *resource.RetryError {
		_, httpResp, err := conn.SecurityGroupRuleApi.DeleteSecurityGroupRule(context.Background()).DeleteSecurityGroupRuleRequest(req).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp, err)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("Error revoking security group %s rules: %s", d.Id(), err)
	}

	return nil
}

func expandRules(d *schema.ResourceData) *[]oscgo.SecurityGroupRule {
	if len(d.Get("rules").([]interface{})) > 0 {
		rules := make([]oscgo.SecurityGroupRule, len(d.Get("rules").([]interface{})))

		for i, rule := range d.Get("rules").([]interface{}) {
			r := rule.(map[string]interface{})

			rules[i] = oscgo.SecurityGroupRule{
				SecurityGroupsMembers: expandSecurityGroupsMembers(r["security_groups_members"].([]interface{})),
			}

			if ipRanges := utils.InterfaceSliceToStringSlicePtr(r["ip_ranges"].([]interface{})); len(*ipRanges) > 0 {
				rules[i].IpRanges = utils.InterfaceSliceToStringSlicePtr(r["ip_ranges"].([]interface{}))
			}
			if serviceIDs := utils.InterfaceSliceToStringSlicePtr(r["service_ids"].([]interface{})); len(*serviceIDs) > 0 {
				rules[i].ServiceIds = utils.InterfaceSliceToStringSlicePtr(r["service_ids"].([]interface{}))
			}
			if v, ok := r["from_port_range"]; ok {
				rules[i].FromPortRange = pointy.Int32(cast.ToInt32(v))
			}
			if v, ok := r["ip_protocol"]; ok && v != "" {
				rules[i].IpProtocol = pointy.String(cast.ToString(v))
			}
			if v, ok := r["to_port_range"]; ok {
				rules[i].ToPortRange = pointy.Int32(cast.ToInt32(v))
			}
		}
		return &rules
	}
	return nil
}

func flattenRules(securityGroupsRules []oscgo.SecurityGroupRule) []map[string]interface{} {
	sgrs := make([]map[string]interface{}, len(securityGroupsRules))

	for i, s := range securityGroupsRules {
		sgrs[i] = map[string]interface{}{
			"from_port_range":         s.GetFromPortRange(),
			"ip_protocol":             s.GetIpProtocol(),
			"ip_ranges":               s.GetIpRanges(),
			"service_ids":             s.GetServiceIds(),
			"to_port_range":           s.GetToPortRange(),
			"security_groups_members": flattenSecurityGroupsMembers(s.GetSecurityGroupsMembers()),
		}
	}
	return sgrs
}

func flattenSecurityGroupsMembers(securityGroupMembers []oscgo.SecurityGroupsMember) []map[string]interface{} {
	sgms := make([]map[string]interface{}, len(securityGroupMembers))

	for i, s := range securityGroupMembers {
		sgms[i] = map[string]interface{}{
			"account_id":          s.GetAccountId(),
			"security_group_id":   s.GetSecurityGroupId(),
			"security_group_name": s.GetSecurityGroupName(),
		}
	}
	return sgms
}

func expandSecurityGroupsMembers(gps []interface{}) *[]oscgo.SecurityGroupsMember {
	groups := make([]oscgo.SecurityGroupsMember, len(gps))

	for i, group := range gps {
		g := group.(map[string]interface{})
		groups[i] = oscgo.SecurityGroupsMember{}

		if v, ok := g["account_id"]; ok && v != "" {
			groups[i].AccountId = pointy.String(cast.ToString(v))
		}
		if v, ok := g["security_group_id"]; ok && v != "" {
			groups[i].SecurityGroupId = pointy.String(cast.ToString(v))
		}
		if v, ok := g["security_group_name"]; ok && v != "" {
			groups[i].SecurityGroupName = pointy.String(cast.ToString(v))
		}
	}
	return &groups
}

func getRulesSchema(isForAttr bool) *schema.Schema {
	return &schema.Schema{
		Type:          schema.TypeList,
		Optional:      true,
		Computed:      isForAttr,
		ForceNew:      !isForAttr,
		ConflictsWith: []string{"ip_protocol", "security_group_name_to_link"},

		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"from_port_range": {
					Type:     schema.TypeInt,
					Optional: !isForAttr,
					ForceNew: !isForAttr,
					Computed: isForAttr,
				},
				"ip_protocol": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: !isForAttr,
					ForceNew: !isForAttr,
				},
				"ip_ranges": {
					Type:     schema.TypeList,
					Optional: !isForAttr,
					ForceNew: !isForAttr,
					Computed: isForAttr,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"service_ids": {
					Type:     schema.TypeList,
					Optional: !isForAttr,
					ForceNew: !isForAttr,
					Computed: isForAttr,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"to_port_range": {
					Type:     schema.TypeInt,
					Optional: !isForAttr,
					ForceNew: !isForAttr,
					Computed: isForAttr,
				},
				"security_groups_members": {
					Type:     schema.TypeList,
					Optional: true,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"account_id": {
								Type:     schema.TypeString,
								Optional: !isForAttr,
								ForceNew: !isForAttr,
								Computed: isForAttr,
							},
							"security_group_id": {
								Type:     schema.TypeString,
								Optional: !isForAttr,
								ForceNew: !isForAttr,
								Computed: isForAttr,
							},
							"security_group_name": {
								Type:     schema.TypeString,
								Optional: !isForAttr,
								ForceNew: !isForAttr,
								Computed: isForAttr,
							},
						},
					},
				},
			},
		},
	}
}

func resourceOutscaleOAPISecurityGroupRuleImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// example: sg-53173ec7_inbound_tcp_80_80_80.14.129.222/32
	// example: sg-53173ec7_inbound_tcp_80_80_sg-53173ec7

	conn := meta.(*OutscaleClient).OSCAPI

	parts := strings.SplitN(d.Id(), "_", 6)
	if len(parts) != 6 {
		return nil, errors.New("import format error: to import a Outscale Security Group Rule, use the format {id}_{flow}_{protocol}_{fromPort}_{toPort}_{ip range or id}")
	}

	sgID := parts[0]
	ruleType := parts[1]
	protocol := parts[2]
	fromPort := parts[3]
	toPort := parts[4]
	sources := parts[5]
	sourcesValidation := parts[5:]

	//Validations
	if !strings.EqualFold(ruleType, "inbound") && !strings.EqualFold(ruleType, "outbound") {
		return nil, errors.New("flow must be inbound or outbound")
	}

	if _, ok := sgProtocolIntegers()[protocol]; !ok {
		if _, err := strconv.Atoi(protocol); err != nil {
			return nil, errors.New("protocol must be tcp/udp/icmp/all or a number")
		}
	}

	if p1, err := strconv.Atoi(fromPort); err != nil {
		return nil, errors.New("invalid from port")
	} else if p2, err := strconv.Atoi(toPort); err != nil || p2 < p1 {
		return nil, errors.New("invalid to port")
	}

	isSGID := false
	for _, source := range sourcesValidation {
		// will be properly validated later
		if !strings.Contains(source, "sg-") && !strings.Contains(source, "pl-") && !strings.Contains(source, ":") && !strings.Contains(source, ".") {
			return nil, errors.New("source must be cidr, ipv6cidr, or a sg ID")
		}

		if strings.Contains(source, "sg-") || strings.Contains(source, "pl-") {
			isSGID = true
		}
	}

	filter := &oscgo.FiltersSecurityGroup{
		SecurityGroupIds: &[]string{sgID},
	}

	if strings.EqualFold(ruleType, "inbound") {
		filter.InboundRuleProtocols = &[]string{protocol}
		filter.InboundRuleFromPortRanges = &[]int32{cast.ToInt32(fromPort)}
		filter.InboundRuleToPortRanges = &[]int32{cast.ToInt32(toPort)}
		if isSGID {
			filter.InboundRuleSecurityGroupIds = &[]string{sources}
		} else {
			filter.InboundRuleIpRanges = &[]string{sources}
		}
	}
	if strings.EqualFold(ruleType, "outbound") {
		filter.OutboundRuleProtocols = &[]string{protocol}
		filter.OutboundRuleFromPortRanges = &[]int32{cast.ToInt32(fromPort)}
		filter.OutboundRuleToPortRanges = &[]int32{cast.ToInt32(toPort)}
		if isSGID {
			filter.OutboundRuleSecurityGroupIds = &[]string{sources}
		} else {
			filter.OutboundRuleIpRanges = &[]string{sources}
		}
	}

	sg, resp, err := readSecurityGroupsWithFilter(conn, filter)
	if err != nil {
		return nil, err
	}

	resultInbound := []oscgo.SecurityGroupRule{}
	resultOutbound := []oscgo.SecurityGroupRule{}

	var ipRange, ipProtocol, fromRange, toRange string

	if strings.EqualFold(ruleType, "inbound") {
		for _, inbound := range sg.GetInboundRules() {
			if inbound.GetIpProtocol() == protocol && inbound.GetFromPortRange() == cast.ToInt32(fromPort) && inbound.GetToPortRange() == cast.ToInt32(toPort) {
				for _, ip := range inbound.GetIpRanges() {
					if ip == sources {
						ipRange = ip
						ipProtocol = protocol
						fromRange = fromPort
						toRange = toPort
						resultInbound = append(resultInbound, inbound)
					}
				}
			}
		}
	}

	if strings.EqualFold(ruleType, "outbound") {
		for _, outbound := range sg.GetOutboundRules() {
			if outbound.GetIpProtocol() == protocol && outbound.GetFromPortRange() == cast.ToInt32(fromPort) && outbound.GetToPortRange() == cast.ToInt32(toPort) {
				for _, ip := range outbound.GetIpRanges() {
					if ip == sources {
						ipRange = ip
						ipProtocol = protocol
						fromRange = fromPort
						toRange = toPort
						resultOutbound = append(resultOutbound, outbound)
					}
				}
			}
		}
	}

	if err := d.Set("inbound_rules", flattenRules(resultInbound)); err != nil {
		return nil, fmt.Errorf("error setting `inbound_rules` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("outbound_rules", flattenRules(resultOutbound)); err != nil {
		return nil, fmt.Errorf("error setting `outbound_rules` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("security_group_name", sg.GetSecurityGroupName()); err != nil {
		return nil, fmt.Errorf("error setting `security_group_name` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("net_id", sg.GetNetId()); err != nil {
		return nil, fmt.Errorf("error setting `net_id` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if ruleType == "inbound" {
		if err := d.Set("flow", "Inbound"); err != nil {
			return nil, fmt.Errorf("error setting `flow` for Outscale Security Group Rule(%s): %s", d.Id(), err)
		}
	} else {
		if err := d.Set("flow", "Outbound"); err != nil {
			return nil, fmt.Errorf("error setting `flow` for Outscale Security Group Rule(%s): %s", d.Id(), err)
		}
	}
	if err := d.Set("from_port_range", cast.ToInt32(fromRange)); err != nil {
		return nil, fmt.Errorf("error setting `from_port_range` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("to_port_range", cast.ToInt32(toRange)); err != nil {
		return nil, fmt.Errorf("error setting `to_port_range` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("ip_protocol", ipProtocol); err != nil {
		return nil, fmt.Errorf("error setting `ip_protocol` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("ip_range", ipRange); err != nil {
		return nil, fmt.Errorf("error setting `ip_range` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("security_group_id", sg.GetSecurityGroupId()); err != nil {
		return nil, fmt.Errorf("error setting `security_group_id` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}
	if err := d.Set("request_id", resp.ResponseContext.GetRequestId()); err != nil {
		return nil, fmt.Errorf("error setting `request_id` for Outscale Security Group Rule(%s): %s", d.Id(), err)
	}

	d.SetId(sg.GetSecurityGroupId())

	return []*schema.ResourceData{d}, nil
}

func readSecurityGroupsWithFilter(client *oscgo.APIClient, filter *oscgo.FiltersSecurityGroup) (*oscgo.SecurityGroup, *oscgo.ReadSecurityGroupsResponse, error) {
	filters := oscgo.ReadSecurityGroupsRequest{
		Filters: filter,
	}

	var err error
	var resp oscgo.ReadSecurityGroupsResponse
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		rp, httpResp, err := client.SecurityGroupApi.ReadSecurityGroups(context.Background()).ReadSecurityGroupsRequest(filters).Execute()
		if err != nil {
			return utils.CheckThrottling(httpResp, err)
		}
		resp = rp
		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error reading the Outscale Security Group(%s): %s", cast.ToString(filter.GetSecurityGroupIds()[0]), err)
	}

	if len(*resp.SecurityGroups) == 0 {
		return nil, nil, fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	}

	return &resp.GetSecurityGroups()[0], &resp, nil
}

func sgProtocolIntegers() map[string]int {
	return map[string]int{
		"udp":  17,
		"tcp":  6,
		"icmp": 1,
		"all":  -1,
	}
}
