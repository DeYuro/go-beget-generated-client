/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatisticGetProcessListResponseProcessList struct for StatisticGetProcessListResponseProcessList
type StatisticGetProcessListResponseProcessList struct {
	Process *[]StatisticGetProcessListResponseProcessListProcessInfo `json:"process,omitempty"`
}

// NewStatisticGetProcessListResponseProcessList instantiates a new StatisticGetProcessListResponseProcessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetProcessListResponseProcessList() *StatisticGetProcessListResponseProcessList {
	this := StatisticGetProcessListResponseProcessList{}
	return &this
}

// NewStatisticGetProcessListResponseProcessListWithDefaults instantiates a new StatisticGetProcessListResponseProcessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetProcessListResponseProcessListWithDefaults() *StatisticGetProcessListResponseProcessList {
	this := StatisticGetProcessListResponseProcessList{}
	return &this
}

// GetProcess returns the Process field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessList) GetProcess() []StatisticGetProcessListResponseProcessListProcessInfo {
	if o == nil || o.Process == nil {
		var ret []StatisticGetProcessListResponseProcessListProcessInfo
		return ret
	}
	return *o.Process
}

// GetProcessOk returns a tuple with the Process field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessList) GetProcessOk() (*[]StatisticGetProcessListResponseProcessListProcessInfo, bool) {
	if o == nil || o.Process == nil {
		return nil, false
	}
	return o.Process, true
}

// HasProcess returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessList) HasProcess() bool {
	if o != nil && o.Process != nil {
		return true
	}

	return false
}

// SetProcess gets a reference to the given []StatisticGetProcessListResponseProcessListProcessInfo and assigns it to the Process field.
func (o *StatisticGetProcessListResponseProcessList) SetProcess(v []StatisticGetProcessListResponseProcessListProcessInfo) {
	o.Process = &v
}

func (o StatisticGetProcessListResponseProcessList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Process != nil {
		toSerialize["process"] = o.Process
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetProcessListResponseProcessList struct {
	value *StatisticGetProcessListResponseProcessList
	isSet bool
}

func (v NullableStatisticGetProcessListResponseProcessList) Get() *StatisticGetProcessListResponseProcessList {
	return v.value
}

func (v *NullableStatisticGetProcessListResponseProcessList) Set(val *StatisticGetProcessListResponseProcessList) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetProcessListResponseProcessList) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetProcessListResponseProcessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetProcessListResponseProcessList(val *StatisticGetProcessListResponseProcessList) *NullableStatisticGetProcessListResponseProcessList {
	return &NullableStatisticGetProcessListResponseProcessList{value: val, isSet: true}
}

func (v NullableStatisticGetProcessListResponseProcessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetProcessListResponseProcessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


