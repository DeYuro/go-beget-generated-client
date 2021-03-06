/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StatisticGetProcessListResponseProcessListProcessInfo struct for StatisticGetProcessListResponseProcessListProcessInfo
type StatisticGetProcessListResponseProcessListProcessInfo struct {
	Pid *int32 `json:"pid,omitempty"`
	Nice *int32 `json:"nice,omitempty"`
	Virt *string `json:"virt,omitempty"`
	Rss *string `json:"rss,omitempty"`
	State *string `json:"state,omitempty"`
	CpuPercent *float32 `json:"cpu_percent,omitempty"`
	MemPercent *float32 `json:"mem_percent,omitempty"`
	TimeRunning *float32 `json:"time_running,omitempty"`
	Name *string `json:"name,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewStatisticGetProcessListResponseProcessListProcessInfo instantiates a new StatisticGetProcessListResponseProcessListProcessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatisticGetProcessListResponseProcessListProcessInfo() *StatisticGetProcessListResponseProcessListProcessInfo {
	this := StatisticGetProcessListResponseProcessListProcessInfo{}
	return &this
}

// NewStatisticGetProcessListResponseProcessListProcessInfoWithDefaults instantiates a new StatisticGetProcessListResponseProcessListProcessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatisticGetProcessListResponseProcessListProcessInfoWithDefaults() *StatisticGetProcessListResponseProcessListProcessInfo {
	this := StatisticGetProcessListResponseProcessListProcessInfo{}
	return &this
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetPid() int32 {
	if o == nil || o.Pid == nil {
		var ret int32
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetPidOk() (*int32, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given int32 and assigns it to the Pid field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetPid(v int32) {
	o.Pid = &v
}

// GetNice returns the Nice field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNice() int32 {
	if o == nil || o.Nice == nil {
		var ret int32
		return ret
	}
	return *o.Nice
}

// GetNiceOk returns a tuple with the Nice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNiceOk() (*int32, bool) {
	if o == nil || o.Nice == nil {
		return nil, false
	}
	return o.Nice, true
}

// HasNice returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasNice() bool {
	if o != nil && o.Nice != nil {
		return true
	}

	return false
}

// SetNice gets a reference to the given int32 and assigns it to the Nice field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetNice(v int32) {
	o.Nice = &v
}

// GetVirt returns the Virt field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetVirt() string {
	if o == nil || o.Virt == nil {
		var ret string
		return ret
	}
	return *o.Virt
}

// GetVirtOk returns a tuple with the Virt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetVirtOk() (*string, bool) {
	if o == nil || o.Virt == nil {
		return nil, false
	}
	return o.Virt, true
}

// HasVirt returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasVirt() bool {
	if o != nil && o.Virt != nil {
		return true
	}

	return false
}

// SetVirt gets a reference to the given string and assigns it to the Virt field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetVirt(v string) {
	o.Virt = &v
}

// GetRss returns the Rss field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetRss() string {
	if o == nil || o.Rss == nil {
		var ret string
		return ret
	}
	return *o.Rss
}

// GetRssOk returns a tuple with the Rss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetRssOk() (*string, bool) {
	if o == nil || o.Rss == nil {
		return nil, false
	}
	return o.Rss, true
}

// HasRss returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasRss() bool {
	if o != nil && o.Rss != nil {
		return true
	}

	return false
}

// SetRss gets a reference to the given string and assigns it to the Rss field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetRss(v string) {
	o.Rss = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetState(v string) {
	o.State = &v
}

// GetCpuPercent returns the CpuPercent field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetCpuPercent() float32 {
	if o == nil || o.CpuPercent == nil {
		var ret float32
		return ret
	}
	return *o.CpuPercent
}

// GetCpuPercentOk returns a tuple with the CpuPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetCpuPercentOk() (*float32, bool) {
	if o == nil || o.CpuPercent == nil {
		return nil, false
	}
	return o.CpuPercent, true
}

// HasCpuPercent returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasCpuPercent() bool {
	if o != nil && o.CpuPercent != nil {
		return true
	}

	return false
}

// SetCpuPercent gets a reference to the given float32 and assigns it to the CpuPercent field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetCpuPercent(v float32) {
	o.CpuPercent = &v
}

// GetMemPercent returns the MemPercent field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetMemPercent() float32 {
	if o == nil || o.MemPercent == nil {
		var ret float32
		return ret
	}
	return *o.MemPercent
}

// GetMemPercentOk returns a tuple with the MemPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetMemPercentOk() (*float32, bool) {
	if o == nil || o.MemPercent == nil {
		return nil, false
	}
	return o.MemPercent, true
}

// HasMemPercent returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasMemPercent() bool {
	if o != nil && o.MemPercent != nil {
		return true
	}

	return false
}

// SetMemPercent gets a reference to the given float32 and assigns it to the MemPercent field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetMemPercent(v float32) {
	o.MemPercent = &v
}

// GetTimeRunning returns the TimeRunning field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetTimeRunning() float32 {
	if o == nil || o.TimeRunning == nil {
		var ret float32
		return ret
	}
	return *o.TimeRunning
}

// GetTimeRunningOk returns a tuple with the TimeRunning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetTimeRunningOk() (*float32, bool) {
	if o == nil || o.TimeRunning == nil {
		return nil, false
	}
	return o.TimeRunning, true
}

// HasTimeRunning returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasTimeRunning() bool {
	if o != nil && o.TimeRunning != nil {
		return true
	}

	return false
}

// SetTimeRunning gets a reference to the given float32 and assigns it to the TimeRunning field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetTimeRunning(v float32) {
	o.TimeRunning = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetName(v string) {
	o.Name = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *StatisticGetProcessListResponseProcessListProcessInfo) SetUser(v string) {
	o.User = &v
}

func (o StatisticGetProcessListResponseProcessListProcessInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pid != nil {
		toSerialize["pid"] = o.Pid
	}
	if o.Nice != nil {
		toSerialize["nice"] = o.Nice
	}
	if o.Virt != nil {
		toSerialize["virt"] = o.Virt
	}
	if o.Rss != nil {
		toSerialize["rss"] = o.Rss
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CpuPercent != nil {
		toSerialize["cpu_percent"] = o.CpuPercent
	}
	if o.MemPercent != nil {
		toSerialize["mem_percent"] = o.MemPercent
	}
	if o.TimeRunning != nil {
		toSerialize["time_running"] = o.TimeRunning
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableStatisticGetProcessListResponseProcessListProcessInfo struct {
	value *StatisticGetProcessListResponseProcessListProcessInfo
	isSet bool
}

func (v NullableStatisticGetProcessListResponseProcessListProcessInfo) Get() *StatisticGetProcessListResponseProcessListProcessInfo {
	return v.value
}

func (v *NullableStatisticGetProcessListResponseProcessListProcessInfo) Set(val *StatisticGetProcessListResponseProcessListProcessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStatisticGetProcessListResponseProcessListProcessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStatisticGetProcessListResponseProcessListProcessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatisticGetProcessListResponseProcessListProcessInfo(val *StatisticGetProcessListResponseProcessListProcessInfo) *NullableStatisticGetProcessListResponseProcessListProcessInfo {
	return &NullableStatisticGetProcessListResponseProcessListProcessInfo{value: val, isSet: true}
}

func (v NullableStatisticGetProcessListResponseProcessListProcessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatisticGetProcessListResponseProcessListProcessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


