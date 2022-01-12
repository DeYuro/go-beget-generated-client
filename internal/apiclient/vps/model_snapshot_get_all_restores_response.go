/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SnapshotGetAllRestoresResponse struct for SnapshotGetAllRestoresResponse
type SnapshotGetAllRestoresResponse struct {
	Restore *[]SnapshotRestore `json:"restore,omitempty"`
}

// NewSnapshotGetAllRestoresResponse instantiates a new SnapshotGetAllRestoresResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotGetAllRestoresResponse() *SnapshotGetAllRestoresResponse {
	this := SnapshotGetAllRestoresResponse{}
	return &this
}

// NewSnapshotGetAllRestoresResponseWithDefaults instantiates a new SnapshotGetAllRestoresResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotGetAllRestoresResponseWithDefaults() *SnapshotGetAllRestoresResponse {
	this := SnapshotGetAllRestoresResponse{}
	return &this
}

// GetRestore returns the Restore field value if set, zero value otherwise.
func (o *SnapshotGetAllRestoresResponse) GetRestore() []SnapshotRestore {
	if o == nil || o.Restore == nil {
		var ret []SnapshotRestore
		return ret
	}
	return *o.Restore
}

// GetRestoreOk returns a tuple with the Restore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotGetAllRestoresResponse) GetRestoreOk() (*[]SnapshotRestore, bool) {
	if o == nil || o.Restore == nil {
		return nil, false
	}
	return o.Restore, true
}

// HasRestore returns a boolean if a field has been set.
func (o *SnapshotGetAllRestoresResponse) HasRestore() bool {
	if o != nil && o.Restore != nil {
		return true
	}

	return false
}

// SetRestore gets a reference to the given []SnapshotRestore and assigns it to the Restore field.
func (o *SnapshotGetAllRestoresResponse) SetRestore(v []SnapshotRestore) {
	o.Restore = &v
}

func (o SnapshotGetAllRestoresResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Restore != nil {
		toSerialize["restore"] = o.Restore
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotGetAllRestoresResponse struct {
	value *SnapshotGetAllRestoresResponse
	isSet bool
}

func (v NullableSnapshotGetAllRestoresResponse) Get() *SnapshotGetAllRestoresResponse {
	return v.value
}

func (v *NullableSnapshotGetAllRestoresResponse) Set(val *SnapshotGetAllRestoresResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotGetAllRestoresResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotGetAllRestoresResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotGetAllRestoresResponse(val *SnapshotGetAllRestoresResponse) *NullableSnapshotGetAllRestoresResponse {
	return &NullableSnapshotGetAllRestoresResponse{value: val, isSet: true}
}

func (v NullableSnapshotGetAllRestoresResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotGetAllRestoresResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

