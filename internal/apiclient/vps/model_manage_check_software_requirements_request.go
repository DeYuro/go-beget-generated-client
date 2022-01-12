/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageCheckSoftwareRequirementsRequest struct for ManageCheckSoftwareRequirementsRequest
type ManageCheckSoftwareRequirementsRequest struct {
	// Идентификатор ОС
	OperatingSystemId *int32 `json:"operating_system_id,omitempty"`
	Info *ManageSoftwareInstallInfo `json:"info,omitempty"`
}

// NewManageCheckSoftwareRequirementsRequest instantiates a new ManageCheckSoftwareRequirementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCheckSoftwareRequirementsRequest() *ManageCheckSoftwareRequirementsRequest {
	this := ManageCheckSoftwareRequirementsRequest{}
	return &this
}

// NewManageCheckSoftwareRequirementsRequestWithDefaults instantiates a new ManageCheckSoftwareRequirementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCheckSoftwareRequirementsRequestWithDefaults() *ManageCheckSoftwareRequirementsRequest {
	this := ManageCheckSoftwareRequirementsRequest{}
	return &this
}

// GetOperatingSystemId returns the OperatingSystemId field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsRequest) GetOperatingSystemId() int32 {
	if o == nil || o.OperatingSystemId == nil {
		var ret int32
		return ret
	}
	return *o.OperatingSystemId
}

// GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsRequest) GetOperatingSystemIdOk() (*int32, bool) {
	if o == nil || o.OperatingSystemId == nil {
		return nil, false
	}
	return o.OperatingSystemId, true
}

// HasOperatingSystemId returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsRequest) HasOperatingSystemId() bool {
	if o != nil && o.OperatingSystemId != nil {
		return true
	}

	return false
}

// SetOperatingSystemId gets a reference to the given int32 and assigns it to the OperatingSystemId field.
func (o *ManageCheckSoftwareRequirementsRequest) SetOperatingSystemId(v int32) {
	o.OperatingSystemId = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ManageCheckSoftwareRequirementsRequest) GetInfo() ManageSoftwareInstallInfo {
	if o == nil || o.Info == nil {
		var ret ManageSoftwareInstallInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCheckSoftwareRequirementsRequest) GetInfoOk() (*ManageSoftwareInstallInfo, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ManageCheckSoftwareRequirementsRequest) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given ManageSoftwareInstallInfo and assigns it to the Info field.
func (o *ManageCheckSoftwareRequirementsRequest) SetInfo(v ManageSoftwareInstallInfo) {
	o.Info = &v
}

func (o ManageCheckSoftwareRequirementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OperatingSystemId != nil {
		toSerialize["operating_system_id"] = o.OperatingSystemId
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	return json.Marshal(toSerialize)
}

type NullableManageCheckSoftwareRequirementsRequest struct {
	value *ManageCheckSoftwareRequirementsRequest
	isSet bool
}

func (v NullableManageCheckSoftwareRequirementsRequest) Get() *ManageCheckSoftwareRequirementsRequest {
	return v.value
}

func (v *NullableManageCheckSoftwareRequirementsRequest) Set(val *ManageCheckSoftwareRequirementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCheckSoftwareRequirementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCheckSoftwareRequirementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCheckSoftwareRequirementsRequest(val *ManageCheckSoftwareRequirementsRequest) *NullableManageCheckSoftwareRequirementsRequest {
	return &NullableManageCheckSoftwareRequirementsRequest{value: val, isSet: true}
}

func (v NullableManageCheckSoftwareRequirementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCheckSoftwareRequirementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

