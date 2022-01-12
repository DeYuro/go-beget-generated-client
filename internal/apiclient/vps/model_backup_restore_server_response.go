/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BackupRestoreServerResponse struct for BackupRestoreServerResponse
type BackupRestoreServerResponse struct {
	Order *StructuresOrderInfo `json:"order,omitempty"`
	Error *BackupRestoreServerResponseError `json:"error,omitempty"`
}

// NewBackupRestoreServerResponse instantiates a new BackupRestoreServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupRestoreServerResponse() *BackupRestoreServerResponse {
	this := BackupRestoreServerResponse{}
	return &this
}

// NewBackupRestoreServerResponseWithDefaults instantiates a new BackupRestoreServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupRestoreServerResponseWithDefaults() *BackupRestoreServerResponse {
	this := BackupRestoreServerResponse{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *BackupRestoreServerResponse) GetOrder() StructuresOrderInfo {
	if o == nil || o.Order == nil {
		var ret StructuresOrderInfo
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreServerResponse) GetOrderOk() (*StructuresOrderInfo, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *BackupRestoreServerResponse) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given StructuresOrderInfo and assigns it to the Order field.
func (o *BackupRestoreServerResponse) SetOrder(v StructuresOrderInfo) {
	o.Order = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BackupRestoreServerResponse) GetError() BackupRestoreServerResponseError {
	if o == nil || o.Error == nil {
		var ret BackupRestoreServerResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupRestoreServerResponse) GetErrorOk() (*BackupRestoreServerResponseError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BackupRestoreServerResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given BackupRestoreServerResponseError and assigns it to the Error field.
func (o *BackupRestoreServerResponse) SetError(v BackupRestoreServerResponseError) {
	o.Error = &v
}

func (o BackupRestoreServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableBackupRestoreServerResponse struct {
	value *BackupRestoreServerResponse
	isSet bool
}

func (v NullableBackupRestoreServerResponse) Get() *BackupRestoreServerResponse {
	return v.value
}

func (v *NullableBackupRestoreServerResponse) Set(val *BackupRestoreServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupRestoreServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupRestoreServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupRestoreServerResponse(val *BackupRestoreServerResponse) *NullableBackupRestoreServerResponse {
	return &NullableBackupRestoreServerResponse{value: val, isSet: true}
}

func (v NullableBackupRestoreServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupRestoreServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

