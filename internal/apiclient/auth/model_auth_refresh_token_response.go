/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthRefreshTokenResponse struct for AuthRefreshTokenResponse
type AuthRefreshTokenResponse struct {
	// oneof{response}
	Token *string `json:"token,omitempty"`
	// oneof{response}
	Error *string `json:"error,omitempty"`
}

// NewAuthRefreshTokenResponse instantiates a new AuthRefreshTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRefreshTokenResponse() *AuthRefreshTokenResponse {
	this := AuthRefreshTokenResponse{}
	return &this
}

// NewAuthRefreshTokenResponseWithDefaults instantiates a new AuthRefreshTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRefreshTokenResponseWithDefaults() *AuthRefreshTokenResponse {
	this := AuthRefreshTokenResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthRefreshTokenResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRefreshTokenResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthRefreshTokenResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthRefreshTokenResponse) SetToken(v string) {
	o.Token = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *AuthRefreshTokenResponse) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRefreshTokenResponse) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *AuthRefreshTokenResponse) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *AuthRefreshTokenResponse) SetError(v string) {
	o.Error = &v
}

func (o AuthRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableAuthRefreshTokenResponse struct {
	value *AuthRefreshTokenResponse
	isSet bool
}

func (v NullableAuthRefreshTokenResponse) Get() *AuthRefreshTokenResponse {
	return v.value
}

func (v *NullableAuthRefreshTokenResponse) Set(val *AuthRefreshTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRefreshTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRefreshTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRefreshTokenResponse(val *AuthRefreshTokenResponse) *NullableAuthRefreshTokenResponse {
	return &NullableAuthRefreshTokenResponse{value: val, isSet: true}
}

func (v NullableAuthRefreshTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRefreshTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

