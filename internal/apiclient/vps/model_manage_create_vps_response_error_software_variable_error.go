/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageCreateVpsResponseErrorSoftwareVariableError struct for ManageCreateVpsResponseErrorSoftwareVariableError
type ManageCreateVpsResponseErrorSoftwareVariableError struct {
	Id *int32 `json:"id,omitempty"`
	VariableName *[]string `json:"variable_name,omitempty"`
}

// NewManageCreateVpsResponseErrorSoftwareVariableError instantiates a new ManageCreateVpsResponseErrorSoftwareVariableError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCreateVpsResponseErrorSoftwareVariableError() *ManageCreateVpsResponseErrorSoftwareVariableError {
	this := ManageCreateVpsResponseErrorSoftwareVariableError{}
	return &this
}

// NewManageCreateVpsResponseErrorSoftwareVariableErrorWithDefaults instantiates a new ManageCreateVpsResponseErrorSoftwareVariableError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCreateVpsResponseErrorSoftwareVariableErrorWithDefaults() *ManageCreateVpsResponseErrorSoftwareVariableError {
	this := ManageCreateVpsResponseErrorSoftwareVariableError{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) SetId(v int32) {
	o.Id = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetVariableName() []string {
	if o == nil || o.VariableName == nil {
		var ret []string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) GetVariableNameOk() (*[]string, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given []string and assigns it to the VariableName field.
func (o *ManageCreateVpsResponseErrorSoftwareVariableError) SetVariableName(v []string) {
	o.VariableName = &v
}

func (o ManageCreateVpsResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.VariableName != nil {
		toSerialize["variable_name"] = o.VariableName
	}
	return json.Marshal(toSerialize)
}

type NullableManageCreateVpsResponseErrorSoftwareVariableError struct {
	value *ManageCreateVpsResponseErrorSoftwareVariableError
	isSet bool
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) Get() *ManageCreateVpsResponseErrorSoftwareVariableError {
	return v.value
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) Set(val *ManageCreateVpsResponseErrorSoftwareVariableError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCreateVpsResponseErrorSoftwareVariableError(val *ManageCreateVpsResponseErrorSoftwareVariableError) *NullableManageCreateVpsResponseErrorSoftwareVariableError {
	return &NullableManageCreateVpsResponseErrorSoftwareVariableError{value: val, isSet: true}
}

func (v NullableManageCreateVpsResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCreateVpsResponseErrorSoftwareVariableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


