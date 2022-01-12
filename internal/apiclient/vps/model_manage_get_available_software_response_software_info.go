/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ManageGetAvailableSoftwareResponseSoftwareInfo struct for ManageGetAvailableSoftwareResponseSoftwareInfo
type ManageGetAvailableSoftwareResponseSoftwareInfo struct {
	Id *int32 `json:"id,omitempty"`
	// Внутреннее имя (напр. redmine)
	Name *string `json:"name,omitempty"`
	// Отображаемое имя (напр. Redmine)
	DisplayName *string `json:"display_name,omitempty"`
	// Версия (напр. 1.0.1)
	Version *string `json:"version,omitempty"`
	// Тип ПО (панель управления, обычное ПО)
	Type *string `json:"type,omitempty"`
	// Описание ПО
	Description *string `json:"description,omitempty"`
	// Дополнительное описание для конкретной версии ПО
	DescriptionVersion *string `json:"description_version,omitempty"`
	// Список ID образов, для которых возможна установка этого ПО
	ImageId *[]int32 `json:"image_id,omitempty"`
	// Список дополнительных полей, которые необходимо передать для установки ПО
	Variable *[]string `json:"variable,omitempty"`
	Requirements *ManageGetAvailableSoftwareResponseSoftwareInfoRequirements `json:"requirements,omitempty"`
}

// NewManageGetAvailableSoftwareResponseSoftwareInfo instantiates a new ManageGetAvailableSoftwareResponseSoftwareInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManageGetAvailableSoftwareResponseSoftwareInfo() *ManageGetAvailableSoftwareResponseSoftwareInfo {
	this := ManageGetAvailableSoftwareResponseSoftwareInfo{}
	return &this
}

// NewManageGetAvailableSoftwareResponseSoftwareInfoWithDefaults instantiates a new ManageGetAvailableSoftwareResponseSoftwareInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManageGetAvailableSoftwareResponseSoftwareInfoWithDefaults() *ManageGetAvailableSoftwareResponseSoftwareInfo {
	this := ManageGetAvailableSoftwareResponseSoftwareInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetVersion(v string) {
	o.Version = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetType(v string) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDescriptionVersion returns the DescriptionVersion field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionVersion() string {
	if o == nil || o.DescriptionVersion == nil {
		var ret string
		return ret
	}
	return *o.DescriptionVersion
}

// GetDescriptionVersionOk returns a tuple with the DescriptionVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionVersionOk() (*string, bool) {
	if o == nil || o.DescriptionVersion == nil {
		return nil, false
	}
	return o.DescriptionVersion, true
}

// HasDescriptionVersion returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDescriptionVersion() bool {
	if o != nil && o.DescriptionVersion != nil {
		return true
	}

	return false
}

// SetDescriptionVersion gets a reference to the given string and assigns it to the DescriptionVersion field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDescriptionVersion(v string) {
	o.DescriptionVersion = &v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetImageId() []int32 {
	if o == nil || o.ImageId == nil {
		var ret []int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetImageIdOk() (*[]int32, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given []int32 and assigns it to the ImageId field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetImageId(v []int32) {
	o.ImageId = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVariable() []string {
	if o == nil || o.Variable == nil {
		var ret []string
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVariableOk() (*[]string, bool) {
	if o == nil || o.Variable == nil {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasVariable() bool {
	if o != nil && o.Variable != nil {
		return true
	}

	return false
}

// SetVariable gets a reference to the given []string and assigns it to the Variable field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetVariable(v []string) {
	o.Variable = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetRequirements() ManageGetAvailableSoftwareResponseSoftwareInfoRequirements {
	if o == nil || o.Requirements == nil {
		var ret ManageGetAvailableSoftwareResponseSoftwareInfoRequirements
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetRequirementsOk() (*ManageGetAvailableSoftwareResponseSoftwareInfoRequirements, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasRequirements() bool {
	if o != nil && o.Requirements != nil {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given ManageGetAvailableSoftwareResponseSoftwareInfoRequirements and assigns it to the Requirements field.
func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetRequirements(v ManageGetAvailableSoftwareResponseSoftwareInfoRequirements) {
	o.Requirements = &v
}

func (o ManageGetAvailableSoftwareResponseSoftwareInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DescriptionVersion != nil {
		toSerialize["description_version"] = o.DescriptionVersion
	}
	if o.ImageId != nil {
		toSerialize["image_id"] = o.ImageId
	}
	if o.Variable != nil {
		toSerialize["variable"] = o.Variable
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	return json.Marshal(toSerialize)
}

type NullableManageGetAvailableSoftwareResponseSoftwareInfo struct {
	value *ManageGetAvailableSoftwareResponseSoftwareInfo
	isSet bool
}

func (v NullableManageGetAvailableSoftwareResponseSoftwareInfo) Get() *ManageGetAvailableSoftwareResponseSoftwareInfo {
	return v.value
}

func (v *NullableManageGetAvailableSoftwareResponseSoftwareInfo) Set(val *ManageGetAvailableSoftwareResponseSoftwareInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableManageGetAvailableSoftwareResponseSoftwareInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableManageGetAvailableSoftwareResponseSoftwareInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageGetAvailableSoftwareResponseSoftwareInfo(val *ManageGetAvailableSoftwareResponseSoftwareInfo) *NullableManageGetAvailableSoftwareResponseSoftwareInfo {
	return &NullableManageGetAvailableSoftwareResponseSoftwareInfo{value: val, isSet: true}
}

func (v NullableManageGetAvailableSoftwareResponseSoftwareInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageGetAvailableSoftwareResponseSoftwareInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


