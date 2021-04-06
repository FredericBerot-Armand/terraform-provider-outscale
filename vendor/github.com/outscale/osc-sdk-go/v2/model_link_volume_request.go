/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.7
 * Contact: support@outscale.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package osc

import (
	"encoding/json"
)

// LinkVolumeRequest struct for LinkVolumeRequest
type LinkVolumeRequest struct {
	// The name of the device.
	DeviceName string `json:"DeviceName"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The ID of the VM you want to attach the volume to.
	VmId string `json:"VmId"`
	// The ID of the volume you want to attach.
	VolumeId string `json:"VolumeId"`
}

// NewLinkVolumeRequest instantiates a new LinkVolumeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkVolumeRequest(deviceName string, vmId string, volumeId string) *LinkVolumeRequest {
	this := LinkVolumeRequest{}
	this.DeviceName = deviceName
	this.VmId = vmId
	this.VolumeId = volumeId
	return &this
}

// NewLinkVolumeRequestWithDefaults instantiates a new LinkVolumeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkVolumeRequestWithDefaults() *LinkVolumeRequest {
	this := LinkVolumeRequest{}
	return &this
}

// GetDeviceName returns the DeviceName field value
func (o *LinkVolumeRequest) GetDeviceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value
// and a boolean to check if the value has been set.
func (o *LinkVolumeRequest) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceName, true
}

// SetDeviceName sets field value
func (o *LinkVolumeRequest) SetDeviceName(v string) {
	o.DeviceName = v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *LinkVolumeRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkVolumeRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *LinkVolumeRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *LinkVolumeRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetVmId returns the VmId field value
func (o *LinkVolumeRequest) GetVmId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VmId
}

// GetVmIdOk returns a tuple with the VmId field value
// and a boolean to check if the value has been set.
func (o *LinkVolumeRequest) GetVmIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmId, true
}

// SetVmId sets field value
func (o *LinkVolumeRequest) SetVmId(v string) {
	o.VmId = v
}

// GetVolumeId returns the VolumeId field value
func (o *LinkVolumeRequest) GetVolumeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *LinkVolumeRequest) GetVolumeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *LinkVolumeRequest) SetVolumeId(v string) {
	o.VolumeId = v
}

func (o LinkVolumeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	if true {
		toSerialize["VmId"] = o.VmId
	}
	if true {
		toSerialize["VolumeId"] = o.VolumeId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkVolumeRequest struct {
	value *LinkVolumeRequest
	isSet bool
}

func (v NullableLinkVolumeRequest) Get() *LinkVolumeRequest {
	return v.value
}

func (v *NullableLinkVolumeRequest) Set(val *LinkVolumeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkVolumeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkVolumeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkVolumeRequest(val *LinkVolumeRequest) *NullableLinkVolumeRequest {
	return &NullableLinkVolumeRequest{value: val, isSet: true}
}

func (v NullableLinkVolumeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkVolumeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
