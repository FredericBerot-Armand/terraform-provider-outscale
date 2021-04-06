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

// CreateApiAccessRuleResponse struct for CreateApiAccessRuleResponse
type CreateApiAccessRuleResponse struct {
	ApiAccessRule   *ApiAccessRule   `json:"ApiAccessRule,omitempty"`
	ResponseContext *ResponseContext `json:"ResponseContext,omitempty"`
}

// NewCreateApiAccessRuleResponse instantiates a new CreateApiAccessRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateApiAccessRuleResponse() *CreateApiAccessRuleResponse {
	this := CreateApiAccessRuleResponse{}
	return &this
}

// NewCreateApiAccessRuleResponseWithDefaults instantiates a new CreateApiAccessRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateApiAccessRuleResponseWithDefaults() *CreateApiAccessRuleResponse {
	this := CreateApiAccessRuleResponse{}
	return &this
}

// GetApiAccessRule returns the ApiAccessRule field value if set, zero value otherwise.
func (o *CreateApiAccessRuleResponse) GetApiAccessRule() ApiAccessRule {
	if o == nil || o.ApiAccessRule == nil {
		var ret ApiAccessRule
		return ret
	}
	return *o.ApiAccessRule
}

// GetApiAccessRuleOk returns a tuple with the ApiAccessRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleResponse) GetApiAccessRuleOk() (*ApiAccessRule, bool) {
	if o == nil || o.ApiAccessRule == nil {
		return nil, false
	}
	return o.ApiAccessRule, true
}

// HasApiAccessRule returns a boolean if a field has been set.
func (o *CreateApiAccessRuleResponse) HasApiAccessRule() bool {
	if o != nil && o.ApiAccessRule != nil {
		return true
	}

	return false
}

// SetApiAccessRule gets a reference to the given ApiAccessRule and assigns it to the ApiAccessRule field.
func (o *CreateApiAccessRuleResponse) SetApiAccessRule(v ApiAccessRule) {
	o.ApiAccessRule = &v
}

// GetResponseContext returns the ResponseContext field value if set, zero value otherwise.
func (o *CreateApiAccessRuleResponse) GetResponseContext() ResponseContext {
	if o == nil || o.ResponseContext == nil {
		var ret ResponseContext
		return ret
	}
	return *o.ResponseContext
}

// GetResponseContextOk returns a tuple with the ResponseContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateApiAccessRuleResponse) GetResponseContextOk() (*ResponseContext, bool) {
	if o == nil || o.ResponseContext == nil {
		return nil, false
	}
	return o.ResponseContext, true
}

// HasResponseContext returns a boolean if a field has been set.
func (o *CreateApiAccessRuleResponse) HasResponseContext() bool {
	if o != nil && o.ResponseContext != nil {
		return true
	}

	return false
}

// SetResponseContext gets a reference to the given ResponseContext and assigns it to the ResponseContext field.
func (o *CreateApiAccessRuleResponse) SetResponseContext(v ResponseContext) {
	o.ResponseContext = &v
}

func (o CreateApiAccessRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiAccessRule != nil {
		toSerialize["ApiAccessRule"] = o.ApiAccessRule
	}
	if o.ResponseContext != nil {
		toSerialize["ResponseContext"] = o.ResponseContext
	}
	return json.Marshal(toSerialize)
}

type NullableCreateApiAccessRuleResponse struct {
	value *CreateApiAccessRuleResponse
	isSet bool
}

func (v NullableCreateApiAccessRuleResponse) Get() *CreateApiAccessRuleResponse {
	return v.value
}

func (v *NullableCreateApiAccessRuleResponse) Set(val *CreateApiAccessRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateApiAccessRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateApiAccessRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateApiAccessRuleResponse(val *CreateApiAccessRuleResponse) *NullableCreateApiAccessRuleResponse {
	return &NullableCreateApiAccessRuleResponse{value: val, isSet: true}
}

func (v NullableCreateApiAccessRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateApiAccessRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
