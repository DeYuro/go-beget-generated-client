# StructuresOperatingSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Идентификатор ОС | [optional] 
**DistroName** | Pointer to **string** | Системное имя дистрибутива (например, для группировки) | [optional] 
**DistroDisplayName** | Pointer to **string** | Отображаемое имя дистрибутива | [optional] 
**Version** | Pointer to **string** | Версия дистрибутива (отображаемое значение) | [optional] 
**Active** | Pointer to **bool** | Доступна ли ОС для выбора пользователю (визуально) | [optional] 
**Requirements** | Pointer to [**StructuresOperatingSystemRequirements**](StructuresOperatingSystemRequirements.md) |  | [optional] 

## Methods

### NewStructuresOperatingSystem

`func NewStructuresOperatingSystem() *StructuresOperatingSystem`

NewStructuresOperatingSystem instantiates a new StructuresOperatingSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresOperatingSystemWithDefaults

`func NewStructuresOperatingSystemWithDefaults() *StructuresOperatingSystem`

NewStructuresOperatingSystemWithDefaults instantiates a new StructuresOperatingSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructuresOperatingSystem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructuresOperatingSystem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructuresOperatingSystem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StructuresOperatingSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDistroName

`func (o *StructuresOperatingSystem) GetDistroName() string`

GetDistroName returns the DistroName field if non-nil, zero value otherwise.

### GetDistroNameOk

`func (o *StructuresOperatingSystem) GetDistroNameOk() (*string, bool)`

GetDistroNameOk returns a tuple with the DistroName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroName

`func (o *StructuresOperatingSystem) SetDistroName(v string)`

SetDistroName sets DistroName field to given value.

### HasDistroName

`func (o *StructuresOperatingSystem) HasDistroName() bool`

HasDistroName returns a boolean if a field has been set.

### GetDistroDisplayName

`func (o *StructuresOperatingSystem) GetDistroDisplayName() string`

GetDistroDisplayName returns the DistroDisplayName field if non-nil, zero value otherwise.

### GetDistroDisplayNameOk

`func (o *StructuresOperatingSystem) GetDistroDisplayNameOk() (*string, bool)`

GetDistroDisplayNameOk returns a tuple with the DistroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroDisplayName

`func (o *StructuresOperatingSystem) SetDistroDisplayName(v string)`

SetDistroDisplayName sets DistroDisplayName field to given value.

### HasDistroDisplayName

`func (o *StructuresOperatingSystem) HasDistroDisplayName() bool`

HasDistroDisplayName returns a boolean if a field has been set.

### GetVersion

`func (o *StructuresOperatingSystem) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StructuresOperatingSystem) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StructuresOperatingSystem) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StructuresOperatingSystem) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetActive

`func (o *StructuresOperatingSystem) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *StructuresOperatingSystem) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *StructuresOperatingSystem) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *StructuresOperatingSystem) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRequirements

`func (o *StructuresOperatingSystem) GetRequirements() StructuresOperatingSystemRequirements`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *StructuresOperatingSystem) GetRequirementsOk() (*StructuresOperatingSystemRequirements, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *StructuresOperatingSystem) SetRequirements(v StructuresOperatingSystemRequirements)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *StructuresOperatingSystem) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


