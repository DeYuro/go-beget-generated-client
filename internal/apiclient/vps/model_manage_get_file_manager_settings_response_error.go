/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageGetFileManagerSettingsResponseError struct for ManageGetFileManagerSettingsResponseError
type ManageGetFileManagerSettingsResponseError struct {
	// Код ошибки
	Code *string `json:"code,omitempty"`
	// Сообщение об ошибке
	Message *string `json:"message,omitempty"`
}

// NewManageGetFileManagerSettingsResponseError instantiates a new ManageGetFileManagerSettingsResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetFileManagerSettingsResponseError() *ManageGetFileManagerSettingsResponseError {
	this := ManageGetFileManagerSettingsResponseError{}
	return &this
}

// NewManageGetFileManagerSettingsResponseErrorWithDefaults instantiates a new ManageGetFileManagerSettingsResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetFileManagerSettingsResponseErrorWithDefaults() *ManageGetFileManagerSettingsResponseError {
	this := ManageGetFileManagerSettingsResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ManageGetFileManagerSettingsResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ManageGetFileManagerSettingsResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetFileManagerSettingsResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ManageGetFileManagerSettingsResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ManageGetFileManagerSettingsResponseError) SetMessage(v string) {
	o.Message = &v
}

func (o ManageGetFileManagerSettingsResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetFileManagerSettingsResponseError struct {
	value *ManageGetFileManagerSettingsResponseError
	isSet bool
}

func (v NullableManageGetFileManagerSettingsResponseError) Get() *ManageGetFileManagerSettingsResponseError {
	return v.value
}

func (v *NullableManageGetFileManagerSettingsResponseError) Set(val *ManageGetFileManagerSettingsResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetFileManagerSettingsResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetFileManagerSettingsResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetFileManagerSettingsResponseError(val *ManageGetFileManagerSettingsResponseError) *NullableManageGetFileManagerSettingsResponseError {
	return &NullableManageGetFileManagerSettingsResponseError{value: val, isSet: true}
}

func (v NullableManageGetFileManagerSettingsResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetFileManagerSettingsResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


