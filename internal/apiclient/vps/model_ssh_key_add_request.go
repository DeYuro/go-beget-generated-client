/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// SshKeyAddRequest struct for SshKeyAddRequest
type SshKeyAddRequest struct {
	// Имя ключа (для удобства пользователя)
	Name *string `json:"name,omitempty"`
	// Публичный ключ в формате, описанном в https://tools.ietf.org/html/rfc4253#section-6.6
	PublicKey *string `json:"public_key,omitempty"`
}

// NewSshKeyAddRequest instantiates a new SshKeyAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyAddRequest() *SshKeyAddRequest {
	this := SshKeyAddRequest{}
	return &this
}

// NewSshKeyAddRequestWithDefaults instantiates a new SshKeyAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyAddRequestWithDefaults() *SshKeyAddRequest {
	this := SshKeyAddRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SshKeyAddRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SshKeyAddRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SshKeyAddRequest) SetName(v string) {
	o.Name = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *SshKeyAddRequest) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyAddRequest) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *SshKeyAddRequest) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *SshKeyAddRequest) SetPublicKey(v string) {
	o.PublicKey = &v
}

func (o SshKeyAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	return json.Marshal(toSerialize)
}

type NullableSshKeyAddRequest struct {
	value *SshKeyAddRequest
	isSet bool
}

func (v NullableSshKeyAddRequest) Get() *SshKeyAddRequest {
	return v.value
}

func (v *NullableSshKeyAddRequest) Set(val *SshKeyAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyAddRequest(val *SshKeyAddRequest) *NullableSshKeyAddRequest {
	return &NullableSshKeyAddRequest{value: val, isSet: true}
}

func (v NullableSshKeyAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


