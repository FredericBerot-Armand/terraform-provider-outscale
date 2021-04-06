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

// CreateCaRequest struct for CreateCaRequest
type CreateCaRequest struct {
	// The CA in PEM format.
	CaPem string `json:"CaPem"`
	// The description of the CA.
	Description *string `json:"Description,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
}

// NewCreateCaRequest instantiates a new CreateCaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaRequest(caPem string) *CreateCaRequest {
	this := CreateCaRequest{}
	this.CaPem = caPem
	return &this
}

// NewCreateCaRequestWithDefaults instantiates a new CreateCaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaRequestWithDefaults() *CreateCaRequest {
	this := CreateCaRequest{}
	return &this
}

// GetCaPem returns the CaPem field value
func (o *CreateCaRequest) GetCaPem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CaPem
}

// GetCaPemOk returns a tuple with the CaPem field value
// and a boolean to check if the value has been set.
func (o *CreateCaRequest) GetCaPemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaPem, true
}

// SetCaPem sets field value
func (o *CreateCaRequest) SetCaPem(v string) {
	o.CaPem = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateCaRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateCaRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateCaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateCaRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaRequest) GetDryRunOk() (*bool, bool) {
	if o == nil || o.DryRun == nil {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateCaRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateCaRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

func (o CreateCaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["CaPem"] = o.CaPem
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.DryRun != nil {
		toSerialize["DryRun"] = o.DryRun
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCaRequest struct {
	value *CreateCaRequest
	isSet bool
}

func (v NullableCreateCaRequest) Get() *CreateCaRequest {
	return v.value
}

func (v *NullableCreateCaRequest) Set(val *CreateCaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaRequest(val *CreateCaRequest) *NullableCreateCaRequest {
	return &NullableCreateCaRequest{value: val, isSet: true}
}

func (v NullableCreateCaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
