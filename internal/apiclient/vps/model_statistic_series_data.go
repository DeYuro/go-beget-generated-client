/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatisticSeriesData struct for StatisticSeriesData
type StatisticSeriesData struct {
	Date *[]string `json:"date,omitempty"`
	Value *[]float64 `json:"value,omitempty"`
}

// NewStatisticSeriesData instantiates a new StatisticSeriesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticSeriesData() *StatisticSeriesData {
	this := StatisticSeriesData{}
	return &this
}

// NewStatisticSeriesDataWithDefaults instantiates a new StatisticSeriesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticSeriesDataWithDefaults() *StatisticSeriesData {
	this := StatisticSeriesData{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StatisticSeriesData) GetDate() []string {
	if o == nil || o.Date == nil {
		var ret []string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSeriesData) GetDateOk() (*[]string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StatisticSeriesData) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given []string and assigns it to the Date field.
func (o *StatisticSeriesData) SetDate(v []string) {
	o.Date = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StatisticSeriesData) GetValue() []float64 {
	if o == nil || o.Value == nil {
		var ret []float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticSeriesData) GetValueOk() (*[]float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StatisticSeriesData) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []float64 and assigns it to the Value field.
func (o *StatisticSeriesData) SetValue(v []float64) {
	o.Value = &v
}

func (o StatisticSeriesData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticSeriesData struct {
	value *StatisticSeriesData
	isSet bool
}

func (v NullableStatisticSeriesData) Get() *StatisticSeriesData {
	return v.value
}

func (v *NullableStatisticSeriesData) Set(val *StatisticSeriesData) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticSeriesData) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticSeriesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticSeriesData(val *StatisticSeriesData) *NullableStatisticSeriesData {
	return &NullableStatisticSeriesData{value: val, isSet: true}
}

func (v NullableStatisticSeriesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticSeriesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


