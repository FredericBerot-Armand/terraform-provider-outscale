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

// LinkFlexibleGpuResponse struct for LinkFlexibleGpuResponse
type LinkFlexibleGpuResponse struct {
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewLinkFlexibleGpuResponse instantiates a new LinkFlexibleGpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkFlexibleGpuResponse() *LinkFlexibleGpuResponse {
	this := LinkFlexibleGpuResponse{}
	return &this
}

// NewLinkFlexibleGpuResponseWithDefaults instantiates a new LinkFlexibleGpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkFlexibleGpuResponseWithDefaults() *LinkFlexibleGpuResponse {
	this := LinkFlexibleGpuResponse{}
	return &this
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *LinkFlexibleGpuResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkFlexibleGpuResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *LinkFlexibleGpuResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *LinkFlexibleGpuResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o LinkFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableLinkFlexibleGpuResponse struct {
	value *LinkFlexibleGpuResponse
	isSet bool
}

func (v NullableLinkFlexibleGpuResponse) Get() *LinkFlexibleGpuResponse {
	return v.value
}

func (v *NullableLinkFlexibleGpuResponse) Set(val *LinkFlexibleGpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkFlexibleGpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkFlexibleGpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkFlexibleGpuResponse(val *LinkFlexibleGpuResponse) *NullableLinkFlexibleGpuResponse {
	return &NullableLinkFlexibleGpuResponse{value: val, isSet: true}
}

func (v NullableLinkFlexibleGpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkFlexibleGpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
