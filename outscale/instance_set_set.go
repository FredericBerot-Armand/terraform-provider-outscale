package outscale

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	oscgo "github.com/marinsalinas/osc-sdk-go"
	"github.com/outscale/osc-go/oapi"
	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
)

func resourceInstancSetHash(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})

	if m["ami_launch_index"] != nil {
		buf.WriteString(fmt.Sprintf("%d-", m["ami_launch_index"].(int)))
	}

	if m["architecture"] != nil {
		buf.WriteString(fmt.Sprintf("%s-", m["architecture"].(string)))
	}
	if m["ip_address"] != nil {
		buf.WriteString(fmt.Sprintf("%s-", m["architecture"].(string)))
	}

	return hashcode.String(buf.String())

}

func getPrivateIPAddressSet(privateIPs []*fcu.InstancePrivateIpAddress) []map[string]interface{} {
	res := []map[string]interface{}{}
	if privateIPs != nil {
		for _, p := range privateIPs {
			inter := make(map[string]interface{})
			assoc := make(map[string]interface{})

			if p.Association != nil {
				assoc["ip_owner_id"] = *p.Association.IpOwnerId
				assoc["public_dns_name"] = *p.Association.PublicDnsName
				assoc["public_ip"] = *p.Association.PublicIp
			}

			inter["association"] = assoc
			inter["private_dns_name"] = *p.Primary
			inter["private_ip_address"] = *p.PrivateIpAddress

		}
	}
	return res
}

func getPlacement(placement *fcu.Placement) map[string]interface{} {
	res := map[string]interface{}{}

	if placement != nil {
		if placement.Affinity != nil {
			res["affinity"] = *placement.Affinity
		}
		res["availability_zone"] = *placement.AvailabilityZone
		res["group_name"] = *placement.GroupName
		if placement.HostId != nil {
			res["host_id"] = *placement.HostId
		}
		res["tenancy"] = *placement.Tenancy
	}

	return res
}

func getProductCodes(codes []*fcu.ProductCode) []map[string]interface{} {
	var res []map[string]interface{}

	if len(codes) > 0 {
		res = make([]map[string]interface{}, len(codes))
		for _, c := range codes {
			code := map[string]interface{}{}

			code["product_code"] = *c.ProductCode
			code["type"] = *c.Type

			res = append(res, code)
		}
	} else {
		res = make([]map[string]interface{}, 0)
	}

	return res
}

func getStateReason(reason *fcu.StateReason) map[string]interface{} {
	res := map[string]interface{}{}
	if reason != nil {
		res["code"] = reason.Code
		res["message"] = reason.Message
	}
	return res
}

func getTagSet(tags []*fcu.Tag) []map[string]interface{} {
	res := []map[string]interface{}{}

	if tags != nil {
		for _, t := range tags {
			tag := map[string]interface{}{}

			tag["key"] = *t.Key
			tag["value"] = *t.Value

			res = append(res, tag)
		}
	}

	return res
}

func getTagDescriptionSet(tags []*fcu.TagDescription) []map[string]interface{} {
	res := []map[string]interface{}{}

	if tags != nil {
		for _, t := range tags {
			tag := map[string]interface{}{}

			tag["key"] = *t.Key
			tag["value"] = *t.Value
			tag["resourceId"] = *t.ResourceId
			tag["resourceType"] = *t.ResourceType

			res = append(res, tag)
		}
	}

	return res
}

func flattenEBS(ebs *fcu.EbsInstanceBlockDevice) map[string]interface{} {

	res := map[string]interface{}{
		"delete_on_termination": fmt.Sprintf("%t", *ebs.DeleteOnTermination),
		"status":                *ebs.Status,
		"volume_id":             *ebs.VolumeId,
	}

	return res
}

func getBlockDeviceMapping(blockDeviceMappings []*fcu.InstanceBlockDeviceMapping) []map[string]interface{} {
	var blockDeviceMapping []map[string]interface{}

	if len(blockDeviceMappings) > 0 {
		blockDeviceMapping = make([]map[string]interface{}, len(blockDeviceMappings))
		for _, mapping := range blockDeviceMappings {
			r := map[string]interface{}{}
			r["device_name"] = *mapping.DeviceName

			e := map[string]interface{}{}
			e["delete_on_termination"] = *mapping.Ebs.DeleteOnTermination
			e["status"] = *mapping.Ebs.Status
			e["volume_id"] = *mapping.Ebs.VolumeId
			r["ebs"] = e

			blockDeviceMapping = append(blockDeviceMapping, r)
		}
	} else {
		blockDeviceMapping = make([]map[string]interface{}, 0)
	}

	return blockDeviceMapping
}

func getGroupSet(groupSet []*fcu.GroupIdentifier) []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, g := range groupSet {

		r := map[string]interface{}{
			"group_id":   *g.GroupId,
			"group_name": *g.GroupName,
		}
		res = append(res, r)
	}

	return res
}

func getOAPISecurityGroups(groups []oscgo.SecurityGroupLight) (SecurityGroup []map[string]interface{}, SecurityGroupIds []string) {
	for _, g := range groups {
		SecurityGroup = append(SecurityGroup, map[string]interface{}{
			"security_group_id":   g.GetSecurityGroupId(),
			"security_group_name": g.GetSecurityGroupName(),
		})
		SecurityGroupIds = append(SecurityGroupIds, g.GetSecurityGroupId())
	}
	return
}

func getOAPILinkNicLight(l oscgo.LinkNicLight) map[string]interface{} {
	return map[string]interface{}{
		"delete_on_vm_deletion": strconv.FormatBool(l.GetDeleteOnVmDeletion()),
		"device_number":         strconv.FormatInt(int64(l.GetDeviceNumber()), 10),
		"link_nic_id":           l.GetLinkNicId(),
		"state":                 l.GetState(),
	}
}

func getOAPILinkNic(l oapi.LinkNic) map[string]interface{} {
	return map[string]interface{}{
		"delete_on_vm_deletion": strconv.FormatBool(aws.BoolValue(l.DeleteOnVmDeletion)),
		"device_number":         strconv.FormatInt(l.DeviceNumber, 10),
		"link_nic_id":           l.LinkNicId,
		"state":                 l.State,
		"vm_account_id":         l.VmAccountId,
		"vm_id":                 l.VmId,
	}
}

func getOAPILinkPublicIPLight(l oscgo.LinkPublicIpLightForVm) *schema.Set {
	res := &schema.Set{
		F: func(v interface{}) int {
			var buf bytes.Buffer
			m := v.(map[string]interface{})
			buf.WriteString(fmt.Sprintf("%s-", m["public_ip"].(string)))
			buf.WriteString(fmt.Sprintf("%s-", m["public_ip_account_id"].(string)))
			return hashcode.String(buf.String())
		},
	}

	res.Add(map[string]interface{}{
		"public_dns_name":      l.GetPublicDnsName(),
		"public_ip":            l.GetPublicIp(),
		"public_ip_account_id": l.GetPublicIpAccountId(),
	})
	return res
}

func getOAPILinkPublicIP(l oapi.LinkPublicIp) map[string]interface{} {
	return map[string]interface{}{
		"link_public_ip_id":    l.LinkPublicIpId,
		"public_dns_name":      l.PublicDnsName,
		"public_ip":            l.PublicIp,
		"public_ip_account_id": l.PublicIpAccountId,
		"public_ip_id":         l.PublicIpId,
	}
}

func getOAPIPrivateIPsLight(privateIPs []oscgo.PrivateIpLightForVm) *schema.Set {
	res := &schema.Set{
		F: func(v interface{}) int {
			var buf bytes.Buffer
			m := v.(map[string]interface{})
			buf.WriteString(fmt.Sprintf("%s-", m["private_ip"].(string)))
			buf.WriteString(fmt.Sprintf("%s-", m["private_dns_name"].(string)))
			return hashcode.String(buf.String())
		},
	}

	for _, p := range privateIPs {
		r := map[string]interface{}{
			"is_primary":       p.GetIsPrimary(),
			"link_public_ip":   getOAPILinkPublicIPLight(p.GetLinkPublicIp()),
			"private_dns_name": p.GetPrivateDnsName(),
			"private_ip":       p.GetPrivateIp(),
		}
		res.Add(r)
	}
	return res
}

func getOAPIPrivateIPs(privateIPs []oapi.PrivateIp) (res []map[string]interface{}) {
	for _, p := range privateIPs {
		res = append(res, map[string]interface{}{
			"is_primary":       p.IsPrimary,
			"link_public_ip":   getOAPILinkPublicIP(p.LinkPublicIp),
			"private_dns_name": p.PrivateDnsName,
			"private_ip":       p.PrivateIp,
		})
	}
	return
}

func getOAPIVMNetworkInterfaceLightSet(nics []oscgo.NicLight) (res []map[string]interface{}) {
	if nics != nil {
		for _, nic := range nics {
			securityGroups, securityGroupIds := getOAPISecurityGroups(nic.GetSecurityGroups())

			res = append(res, map[string]interface{}{
				"delete_on_vm_deletion":  *nic.GetLinkNic().DeleteOnVmDeletion, // Workaround.
				"account_id":             nic.GetAccountId(),
				"description":            nic.GetDescription(),
				"is_source_dest_checked": nic.GetIsSourceDestChecked(),
				"link_nic":               getOAPILinkNicLight(nic.GetLinkNic()),
				"link_public_ip":         getOAPILinkPublicIPLight(nic.GetLinkPublicIp()),
				"mac_address":            nic.GetMacAddress(),
				"net_id":                 nic.GetNetId(),
				"nic_id":                 nic.GetNicId(),
				"private_dns_name":       nic.GetPrivateDnsName(),
				"private_ips":            getOAPIPrivateIPsLight(nic.GetPrivateIps()),
				"security_groups":        securityGroups,
				"security_group_ids":     securityGroupIds,
				"state":                  nic.GetState(),
				"subnet_id":              nic.GetSubnetId(),
			})
		}
	}
	return
}

func getOAPIVMNetworkInterfaceSet(nics []oapi.Nic) (res []map[string]interface{}) {
	if nics != nil {
		for _, nic := range nics {
			//securityGroups, _ := getOAPISecurityGroups(nic.SecurityGroups)

			res = append(res, map[string]interface{}{
				"account_id":             nic.AccountId,
				"description":            nic.Description,
				"is_source_dest_checked": nic.IsSourceDestChecked,
				"link_nic":               getOAPILinkNic(nic.LinkNic),
				"link_public_ip":         getOAPILinkPublicIP(nic.LinkPublicIp),
				"mac_address":            nic.MacAddress,
				"net_id":                 nic.NetId,
				"nic_id":                 nic.NicId,
				"private_dns_name":       nic.PrivateDnsName,
				"private_ips":            getOAPIPrivateIPs(nic.PrivateIps),
				//"security_groups":        securityGroups,
				"state":          nic.State,
				"subnet_id":      nic.SubnetId,
				"subregion_name": nic.SubregionName,
				//"tags":           getOapiTagSet(nic.Tags),
			})
		}
	}
	return
}
