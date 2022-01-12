/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatisticGetDiskResponse struct for StatisticGetDiskResponse
type StatisticGetDiskResponse struct {
	DataRead *StatisticSeriesData `json:"data_read,omitempty"`
	DataWrite *StatisticSeriesData `json:"data_write,omitempty"`
}

// NewStatisticGetDiskResponse instantiates a new StatisticGetDiskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetDiskResponse() *StatisticGetDiskResponse {
	this := StatisticGetDiskResponse{}
	return &this
}

// NewStatisticGetDiskResponseWithDefaults instantiates a new StatisticGetDiskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetDiskResponseWithDefaults() *StatisticGetDiskResponse {
	this := StatisticGetDiskResponse{}
	return &this
}

// GetDataRead returns the DataRead field value if set, zero value otherwise.
func (o *StatisticGetDiskResponse) GetDataRead() StatisticSeriesData {
	if o == nil || o.DataRead == nil {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataRead
}

// GetDataReadOk returns a tuple with the DataRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetDiskResponse) GetDataReadOk() (*StatisticSeriesData, bool) {
	if o == nil || o.DataRead == nil {
		return nil, false
	}
	return o.DataRead, true
}

// HasDataRead returns a boolean if a field has been set.
func (o *StatisticGetDiskResponse) HasDataRead() bool {
	if o != nil && o.DataRead != nil {
		return true
	}

	return false
}

// SetDataRead gets a reference to the given StatisticSeriesData and assigns it to the DataRead field.
func (o *StatisticGetDiskResponse) SetDataRead(v StatisticSeriesData) {
	o.DataRead = &v
}

// GetDataWrite returns the DataWrite field value if set, zero value otherwise.
func (o *StatisticGetDiskResponse) GetDataWrite() StatisticSeriesData {
	if o == nil || o.DataWrite == nil {
		var ret StatisticSeriesData
		return ret
	}
	return *o.DataWrite
}

// GetDataWriteOk returns a tuple with the DataWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetDiskResponse) GetDataWriteOk() (*StatisticSeriesData, bool) {
	if o == nil || o.DataWrite == nil {
		return nil, false
	}
	return o.DataWrite, true
}

// HasDataWrite returns a boolean if a field has been set.
func (o *StatisticGetDiskResponse) HasDataWrite() bool {
	if o != nil && o.DataWrite != nil {
		return true
	}

	return false
}

// SetDataWrite gets a reference to the given StatisticSeriesData and assigns it to the DataWrite field.
func (o *StatisticGetDiskResponse) SetDataWrite(v StatisticSeriesData) {
	o.DataWrite = &v
}

func (o StatisticGetDiskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataRead != nil {
		toSerialize["data_read"] = o.DataRead
	}
	if o.DataWrite != nil {
		toSerialize["data_write"] = o.DataWrite
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetDiskResponse struct {
	value *StatisticGetDiskResponse
	isSet bool
}

func (v NullableStatisticGetDiskResponse) Get() *StatisticGetDiskResponse {
	return v.value
}

func (v *NullableStatisticGetDiskResponse) Set(val *StatisticGetDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetDiskResponse(val *StatisticGetDiskResponse) *NullableStatisticGetDiskResponse {
	return &NullableStatisticGetDiskResponse{value: val, isSet: true}
}

func (v NullableStatisticGetDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


