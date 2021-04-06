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

// PermissionsOnResource Information about the users who have permissions for the resource.
type PermissionsOnResource struct {
	// The account ID of one or more users who have permissions for the resource.
	AccountIds *[]string `json:"AccountIds,omitempty"`
	// If true, the resource is public. If false, the resource is private.
	GlobalPermission *bool `json:"GlobalPermission,omitempty"`
}

// NewPermissionsOnResource instantiates a new PermissionsOnResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionsOnResource() *PermissionsOnResource {
	this := PermissionsOnResource{}
	return &this
}

// NewPermissionsOnResourceWithDefaults instantiates a new PermissionsOnResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionsOnResourceWithDefaults() *PermissionsOnResource {
	this := PermissionsOnResource{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *PermissionsOnResource) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnResource) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *PermissionsOnResource) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *PermissionsOnResource) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetGlobalPermission returns the GlobalPermission field value if set, zero value otherwise.
func (o *PermissionsOnResource) GetGlobalPermission() bool {
	if o == nil || o.GlobalPermission == nil {
		var ret bool
		return ret
	}
	return *o.GlobalPermission
}

// GetGlobalPermissionOk returns a tuple with the GlobalPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionsOnResource) GetGlobalPermissionOk() (*bool, bool) {
	if o == nil || o.GlobalPermission == nil {
		return nil, false
	}
	return o.GlobalPermission, true
}

// HasGlobalPermission returns a boolean if a field has been set.
func (o *PermissionsOnResource) HasGlobalPermission() bool {
	if o != nil && o.GlobalPermission != nil {
		return true
	}

	return false
}

// SetGlobalPermission gets a reference to the given bool and assigns it to the GlobalPermission field.
func (o *PermissionsOnResource) SetGlobalPermission(v bool) {
	o.GlobalPermission = &v
}

func (o PermissionsOnResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["AccountIds"] = o.AccountIds
	}
	if o.GlobalPermission != nil {
		toSerialize["GlobalPermission"] = o.GlobalPermission
	}
	return json.Marshal(toSerialize)
}

type NullablePermissionsOnResource struct {
	value *PermissionsOnResource
	isSet bool
}

func (v NullablePermissionsOnResource) Get() *PermissionsOnResource {
	return v.value
}

func (v *NullablePermissionsOnResource) Set(val *PermissionsOnResource) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionsOnResource) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionsOnResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionsOnResource(val *PermissionsOnResource) *NullablePermissionsOnResource {
	return &NullablePermissionsOnResource{value: val, isSet: true}
}

func (v NullablePermissionsOnResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionsOnResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
