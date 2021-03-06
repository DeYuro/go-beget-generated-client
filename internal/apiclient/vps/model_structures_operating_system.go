/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StructuresOperatingSystem Описывает ОС для установки на Vps (пара дистрибутив + версия)
type StructuresOperatingSystem struct {
	// Идентификатор ОС
	Id *int32 `json:"id,omitempty"`
	// Системное имя дистрибутива (например, для группировки)
	DistroName *string `json:"distro_name,omitempty"`
	// Отображаемое имя дистрибутива
	DistroDisplayName *string `json:"distro_display_name,omitempty"`
	// Версия дистрибутива (отображаемое значение)
	Version *string `json:"version,omitempty"`
	// Доступна ли ОС для выбора пользователю (визуально)
	Active *bool `json:"active,omitempty"`
	Requirements *StructuresOperatingSystemRequirements `json:"requirements,omitempty"`
}

// NewStructuresOperatingSystem instantiates a new StructuresOperatingSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresOperatingSystem() *StructuresOperatingSystem {
	this := StructuresOperatingSystem{}
	return &this
}

// NewStructuresOperatingSystemWithDefaults instantiates a new StructuresOperatingSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresOperatingSystemWithDefaults() *StructuresOperatingSystem {
	this := StructuresOperatingSystem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *StructuresOperatingSystem) SetId(v int32) {
	o.Id = &v
}

// GetDistroName returns the DistroName field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetDistroName() string {
	if o == nil || o.DistroName == nil {
		var ret string
		return ret
	}
	return *o.DistroName
}

// GetDistroNameOk returns a tuple with the DistroName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetDistroNameOk() (*string, bool) {
	if o == nil || o.DistroName == nil {
		return nil, false
	}
	return o.DistroName, true
}

// HasDistroName returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasDistroName() bool {
	if o != nil && o.DistroName != nil {
		return true
	}

	return false
}

// SetDistroName gets a reference to the given string and assigns it to the DistroName field.
func (o *StructuresOperatingSystem) SetDistroName(v string) {
	o.DistroName = &v
}

// GetDistroDisplayName returns the DistroDisplayName field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetDistroDisplayName() string {
	if o == nil || o.DistroDisplayName == nil {
		var ret string
		return ret
	}
	return *o.DistroDisplayName
}

// GetDistroDisplayNameOk returns a tuple with the DistroDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetDistroDisplayNameOk() (*string, bool) {
	if o == nil || o.DistroDisplayName == nil {
		return nil, false
	}
	return o.DistroDisplayName, true
}

// HasDistroDisplayName returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasDistroDisplayName() bool {
	if o != nil && o.DistroDisplayName != nil {
		return true
	}

	return false
}

// SetDistroDisplayName gets a reference to the given string and assigns it to the DistroDisplayName field.
func (o *StructuresOperatingSystem) SetDistroDisplayName(v string) {
	o.DistroDisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StructuresOperatingSystem) SetVersion(v string) {
	o.Version = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *StructuresOperatingSystem) SetActive(v bool) {
	o.Active = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *StructuresOperatingSystem) GetRequirements() StructuresOperatingSystemRequirements {
	if o == nil || o.Requirements == nil {
		var ret StructuresOperatingSystemRequirements
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresOperatingSystem) GetRequirementsOk() (*StructuresOperatingSystemRequirements, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *StructuresOperatingSystem) HasRequirements() bool {
	if o != nil && o.Requirements != nil {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given StructuresOperatingSystemRequirements and assigns it to the Requirements field.
func (o *StructuresOperatingSystem) SetRequirements(v StructuresOperatingSystemRequirements) {
	o.Requirements = &v
}

func (o StructuresOperatingSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DistroName != nil {
		toSerialize["distro_name"] = o.DistroName
	}
	if o.DistroDisplayName != nil {
		toSerialize["distro_display_name"] = o.DistroDisplayName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresOperatingSystem struct {
	value *StructuresOperatingSystem
	isSet bool
}

func (v NullableStructuresOperatingSystem) Get() *StructuresOperatingSystem {
	return v.value
}

func (v *NullableStructuresOperatingSystem) Set(val *StructuresOperatingSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresOperatingSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresOperatingSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresOperatingSystem(val *StructuresOperatingSystem) *NullableStructuresOperatingSystem {
	return &NullableStructuresOperatingSystem{value: val, isSet: true}
}

func (v NullableStructuresOperatingSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresOperatingSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


