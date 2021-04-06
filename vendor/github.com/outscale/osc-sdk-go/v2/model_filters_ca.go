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

// FiltersCa One or more filters.
type FiltersCa struct {
	// The fingerprints of the CAs.
	CaFingerprints *[]string `json:"CaFingerprints,omitempty"`
	// The IDs of the CAs.
	CaIds *[]string `json:"CaIds,omitempty"`
	// The descriptions of the CAs.
	Descriptions *[]string `json:"Descriptions,omitempty"`
}

// NewFiltersCa instantiates a new FiltersCa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiltersCa() *FiltersCa {
	this := FiltersCa{}
	return &this
}

// NewFiltersCaWithDefaults instantiates a new FiltersCa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersCaWithDefaults() *FiltersCa {
	this := FiltersCa{}
	return &this
}

// GetCaFingerprints returns the CaFingerprints field value if set, zero value otherwise.
func (o *FiltersCa) GetCaFingerprints() []string {
	if o == nil || o.CaFingerprints == nil {
		var ret []string
		return ret
	}
	return *o.CaFingerprints
}

// GetCaFingerprintsOk returns a tuple with the CaFingerprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCa) GetCaFingerprintsOk() (*[]string, bool) {
	if o == nil || o.CaFingerprints == nil {
		return nil, false
	}
	return o.CaFingerprints, true
}

// HasCaFingerprints returns a boolean if a field has been set.
func (o *FiltersCa) HasCaFingerprints() bool {
	if o != nil && o.CaFingerprints != nil {
		return true
	}

	return false
}

// SetCaFingerprints gets a reference to the given []string and assigns it to the CaFingerprints field.
func (o *FiltersCa) SetCaFingerprints(v []string) {
	o.CaFingerprints = &v
}

// GetCaIds returns the CaIds field value if set, zero value otherwise.
func (o *FiltersCa) GetCaIds() []string {
	if o == nil || o.CaIds == nil {
		var ret []string
		return ret
	}
	return *o.CaIds
}

// GetCaIdsOk returns a tuple with the CaIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCa) GetCaIdsOk() (*[]string, bool) {
	if o == nil || o.CaIds == nil {
		return nil, false
	}
	return o.CaIds, true
}

// HasCaIds returns a boolean if a field has been set.
func (o *FiltersCa) HasCaIds() bool {
	if o != nil && o.CaIds != nil {
		return true
	}

	return false
}

// SetCaIds gets a reference to the given []string and assigns it to the CaIds field.
func (o *FiltersCa) SetCaIds(v []string) {
	o.CaIds = &v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *FiltersCa) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return *o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiltersCa) GetDescriptionsOk() (*[]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *FiltersCa) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *FiltersCa) SetDescriptions(v []string) {
	o.Descriptions = &v
}

func (o FiltersCa) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaFingerprints != nil {
		toSerialize["CaFingerprints"] = o.CaFingerprints
	}
	if o.CaIds != nil {
		toSerialize["CaIds"] = o.CaIds
	}
	if o.Descriptions != nil {
		toSerialize["Descriptions"] = o.Descriptions
	}
	return json.Marshal(toSerialize)
}

type NullableFiltersCa struct {
	value *FiltersCa
	isSet bool
}

func (v NullableFiltersCa) Get() *FiltersCa {
	return v.value
}

func (v *NullableFiltersCa) Set(val *FiltersCa) {
	v.value = val
	v.isSet = true
}

func (v NullableFiltersCa) IsSet() bool {
	return v.isSet
}

func (v *NullableFiltersCa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiltersCa(val *FiltersCa) *NullableFiltersCa {
	return &NullableFiltersCa{value: val, isSet: true}
}

func (v NullableFiltersCa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiltersCa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
