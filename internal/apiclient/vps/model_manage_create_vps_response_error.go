/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageCreateVpsResponseError struct for ManageCreateVpsResponseError
type ManageCreateVpsResponseError struct {
	Code *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	VariableError *ManageCreateVpsResponseErrorSoftwareVariableError `json:"variable_error,omitempty"`
}

// NewManageCreateVpsResponseError instantiates a new ManageCreateVpsResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageCreateVpsResponseError() *ManageCreateVpsResponseError {
	this := ManageCreateVpsResponseError{}
	return &this
}

// NewManageCreateVpsResponseErrorWithDefaults instantiates a new ManageCreateVpsResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageCreateVpsResponseErrorWithDefaults() *ManageCreateVpsResponseError {
	this := ManageCreateVpsResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageCreateVpsResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageCreateVpsResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetVariableError returns the VariableError field value if set, zero value otherwise.
func (o *ManageCreateVpsResponseError) GetVariableError() ManageCreateVpsResponseErrorSoftwareVariableError {
	if o == nil || o.VariableError == nil {
		var ret ManageCreateVpsResponseErrorSoftwareVariableError
		return ret
	}
	return *o.VariableError
}

// GetVariableErrorOk returns a tuple with the VariableError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageCreateVpsResponseError) GetVariableErrorOk() (*ManageCreateVpsResponseErrorSoftwareVariableError, bool) {
	if o == nil || o.VariableError == nil {
		return nil, false
	}
	return o.VariableError, true
}

// HasVariableError returns a boolean if a field has been set.
func (o *ManageCreateVpsResponseError) HasVariableError() bool {
	if o != nil && o.VariableError != nil {
		return true
	}

	return false
}

// SetVariableError gets a reference to the given ManageCreateVpsResponseErrorSoftwareVariableError and assigns it to the VariableError field.
func (o *ManageCreateVpsResponseError) SetVariableError(v ManageCreateVpsResponseErrorSoftwareVariableError) {
	o.VariableError = &v
}

func (o ManageCreateVpsResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.VariableError != nil {
		toSerialize["variable_error"] = o.VariableError
	}
	return json.Marshal(toSerialize)
}

type NullableManageCreateVpsResponseError struct {
	value *ManageCreateVpsResponseError
	isSet bool
}

func (v NullableManageCreateVpsResponseError) Get() *ManageCreateVpsResponseError {
	return v.value
}

func (v *NullableManageCreateVpsResponseError) Set(val *ManageCreateVpsResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageCreateVpsResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageCreateVpsResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageCreateVpsResponseError(val *ManageCreateVpsResponseError) *NullableManageCreateVpsResponseError {
	return &NullableManageCreateVpsResponseError{value: val, isSet: true}
}

func (v NullableManageCreateVpsResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageCreateVpsResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


