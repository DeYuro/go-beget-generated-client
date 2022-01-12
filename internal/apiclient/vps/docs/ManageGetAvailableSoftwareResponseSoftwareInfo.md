# ManageGetAvailableSoftwareResponseSoftwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | Внутреннее имя (напр. redmine) | [optional] 
**DisplayName** | Pointer to **string** | Отображаемое имя (напр. Redmine) | [optional] 
**Version** | Pointer to **string** | Версия (напр. 1.0.1) | [optional] 
**Type** | Pointer to **string** | Тип ПО (панель управления, обычное ПО) | [optional] 
**Description** | Pointer to **string** | Описание ПО | [optional] 
**DescriptionVersion** | Pointer to **string** | Дополнительное описание для конкретной версии ПО | [optional] 
**ImageId** | Pointer to **[]int32** | Список ID образов, для которых возможна установка этого ПО | [optional] 
**Variable** | Pointer to **[]string** | Список дополнительных полей, которые необходимо передать для установки ПО | [optional] 
**Requirements** | Pointer to [**ManageGetAvailableSoftwareResponseSoftwareInfoRequirements**](ManageGetAvailableSoftwareResponseSoftwareInfoRequirements.md) |  | [optional] 

## Methods

### NewManageGetAvailableSoftwareResponseSoftwareInfo

`func NewManageGetAvailableSoftwareResponseSoftwareInfo() *ManageGetAvailableSoftwareResponseSoftwareInfo`

NewManageGetAvailableSoftwareResponseSoftwareInfo instantiates a new ManageGetAvailableSoftwareResponseSoftwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageGetAvailableSoftwareResponseSoftwareInfoWithDefaults

`func NewManageGetAvailableSoftwareResponseSoftwareInfoWithDefaults() *ManageGetAvailableSoftwareResponseSoftwareInfo`

NewManageGetAvailableSoftwareResponseSoftwareInfoWithDefaults instantiates a new ManageGetAvailableSoftwareResponseSoftwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetType

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionVersion() string`

GetDescriptionVersion returns the DescriptionVersion field if non-nil, zero value otherwise.

### GetDescriptionVersionOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetDescriptionVersionOk() (*string, bool)`

GetDescriptionVersionOk returns a tuple with the DescriptionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetDescriptionVersion(v string)`

SetDescriptionVersion sets DescriptionVersion field to given value.

### HasDescriptionVersion

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasDescriptionVersion() bool`

HasDescriptionVersion returns a boolean if a field has been set.

### GetImageId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetImageId() []int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetImageIdOk() (*[]int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetImageId(v []int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetVariable

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVariable() []string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetVariableOk() (*[]string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetVariable(v []string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetRequirements

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetRequirements() ManageGetAvailableSoftwareResponseSoftwareInfoRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) GetRequirementsOk() (*ManageGetAvailableSoftwareResponseSoftwareInfoRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) SetRequirements(v ManageGetAvailableSoftwareResponseSoftwareInfoRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *ManageGetAvailableSoftwareResponseSoftwareInfo) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


