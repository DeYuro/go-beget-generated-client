/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageReinstallResponseErrorSoftwareVariableError struct for ManageReinstallResponseErrorSoftwareVariableError
type ManageReinstallResponseErrorSoftwareVariableError struct {
	Id *int32 `json:"id,omitempty"`
	VariableName *[]string `json:"variable_name,omitempty"`
}

// NewManageReinstallResponseErrorSoftwareVariableError instantiates a new ManageReinstallResponseErrorSoftwareVariableError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageReinstallResponseErrorSoftwareVariableError() *ManageReinstallResponseErrorSoftwareVariableError {
	this := ManageReinstallResponseErrorSoftwareVariableError{}
	return &this
}

// NewManageReinstallResponseErrorSoftwareVariableErrorWithDefaults instantiates a new ManageReinstallResponseErrorSoftwareVariableError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageReinstallResponseErrorSoftwareVariableErrorWithDefaults() *ManageReinstallResponseErrorSoftwareVariableError {
	this := ManageReinstallResponseErrorSoftwareVariableError{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageReinstallResponseErrorSoftwareVariableError) SetId(v int32) {
	o.Id = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetVariableName() []string {
	if o == nil || o.VariableName == nil {
		var ret []string
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) GetVariableNameOk() (*[]string, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *ManageReinstallResponseErrorSoftwareVariableError) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given []string and assigns it to the VariableName field.
func (o *ManageReinstallResponseErrorSoftwareVariableError) SetVariableName(v []string) {
	o.VariableName = &v
}

func (o ManageReinstallResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.VariableName != nil {
		toSerialize["variable_name"] = o.VariableName
	}
	return json.Marshal(toSerialize)
}

type NullableManageReinstallResponseErrorSoftwareVariableError struct {
	value *ManageReinstallResponseErrorSoftwareVariableError
	isSet bool
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) Get() *ManageReinstallResponseErrorSoftwareVariableError {
	return v.value
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) Set(val *ManageReinstallResponseErrorSoftwareVariableError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageReinstallResponseErrorSoftwareVariableError(val *ManageReinstallResponseErrorSoftwareVariableError) *NullableManageReinstallResponseErrorSoftwareVariableError {
	return &NullableManageReinstallResponseErrorSoftwareVariableError{value: val, isSet: true}
}

func (v NullableManageReinstallResponseErrorSoftwareVariableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageReinstallResponseErrorSoftwareVariableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

