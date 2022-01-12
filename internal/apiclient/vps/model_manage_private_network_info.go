/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManagePrivateNetworkInfo struct for ManagePrivateNetworkInfo
type ManagePrivateNetworkInfo struct {
	// ID сети
	Id *string `json:"id,omitempty"`
	// желаемый ip-адрес. Если пустая строка, то выберется случайный
	Address *string `json:"address,omitempty"`
}

// NewManagePrivateNetworkInfo instantiates a new ManagePrivateNetworkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagePrivateNetworkInfo() *ManagePrivateNetworkInfo {
	this := ManagePrivateNetworkInfo{}
	return &this
}

// NewManagePrivateNetworkInfoWithDefaults instantiates a new ManagePrivateNetworkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagePrivateNetworkInfoWithDefaults() *ManagePrivateNetworkInfo {
	this := ManagePrivateNetworkInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManagePrivateNetworkInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagePrivateNetworkInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManagePrivateNetworkInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManagePrivateNetworkInfo) SetId(v string) {
	o.Id = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ManagePrivateNetworkInfo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagePrivateNetworkInfo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ManagePrivateNetworkInfo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ManagePrivateNetworkInfo) SetAddress(v string) {
	o.Address = &v
}

func (o ManagePrivateNetworkInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableManagePrivateNetworkInfo struct {
	value *ManagePrivateNetworkInfo
	isSet bool
}

func (v NullableManagePrivateNetworkInfo) Get() *ManagePrivateNetworkInfo {
	return v.value
}

func (v *NullableManagePrivateNetworkInfo) Set(val *ManagePrivateNetworkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManagePrivateNetworkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManagePrivateNetworkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagePrivateNetworkInfo(val *ManagePrivateNetworkInfo) *NullableManagePrivateNetworkInfo {
	return &NullableManagePrivateNetworkInfo{value: val, isSet: true}
}

func (v NullableManagePrivateNetworkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagePrivateNetworkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


