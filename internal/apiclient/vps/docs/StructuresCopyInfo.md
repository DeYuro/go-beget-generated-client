# StructuresCopyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Идентификатор копии | [optional] 
**VpsId** | Pointer to **string** | Идентификатор VPS (uuid) | [optional] 
**VpsName** | Pointer to **string** | Имя VPS, с которого создавалась резервная копия | [optional] 
**Date** | Pointer to **string** | Дата создания копии в RFC3339 | [optional] 
**OperatingSystem** | Pointer to [**StructuresOperatingSystem**](StructuresOperatingSystem.md) |  | [optional] 
**Size** | Pointer to **string** | Общий размер резервной копии в байтах | [optional] 
**Configuration** | Pointer to [**StructuresCopyInfoConfiguration**](StructuresCopyInfoConfiguration.md) |  | [optional] 

## Methods

### NewStructuresCopyInfo

`func NewStructuresCopyInfo() *StructuresCopyInfo`

NewStructuresCopyInfo instantiates a new StructuresCopyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresCopyInfoWithDefaults

`func NewStructuresCopyInfoWithDefaults() *StructuresCopyInfo`

NewStructuresCopyInfoWithDefaults instantiates a new StructuresCopyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructuresCopyInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructuresCopyInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructuresCopyInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StructuresCopyInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVpsId

`func (o *StructuresCopyInfo) GetVpsId() string`

GetVpsId returns the VpsId field if non-nil, zero value otherwise.

### GetVpsIdOk

`func (o *StructuresCopyInfo) GetVpsIdOk() (*string, bool)`

GetVpsIdOk returns a tuple with the VpsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsId

`func (o *StructuresCopyInfo) SetVpsId(v string)`

SetVpsId sets VpsId field to given value.

### HasVpsId

`func (o *StructuresCopyInfo) HasVpsId() bool`

HasVpsId returns a boolean if a field has been set.

### GetVpsName

`func (o *StructuresCopyInfo) GetVpsName() string`

GetVpsName returns the VpsName field if non-nil, zero value otherwise.

### GetVpsNameOk

`func (o *StructuresCopyInfo) GetVpsNameOk() (*string, bool)`

GetVpsNameOk returns a tuple with the VpsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpsName

`func (o *StructuresCopyInfo) SetVpsName(v string)`

SetVpsName sets VpsName field to given value.

### HasVpsName

`func (o *StructuresCopyInfo) HasVpsName() bool`

HasVpsName returns a boolean if a field has been set.

### GetDate

`func (o *StructuresCopyInfo) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *StructuresCopyInfo) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *StructuresCopyInfo) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *StructuresCopyInfo) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *StructuresCopyInfo) GetOperatingSystem() StructuresOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *StructuresCopyInfo) GetOperatingSystemOk() (*StructuresOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *StructuresCopyInfo) SetOperatingSystem(v StructuresOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *StructuresCopyInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetSize

`func (o *StructuresCopyInfo) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *StructuresCopyInfo) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *StructuresCopyInfo) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *StructuresCopyInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetConfiguration

`func (o *StructuresCopyInfo) GetConfiguration() StructuresCopyInfoConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *StructuresCopyInfo) GetConfigurationOk() (*StructuresCopyInfoConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *StructuresCopyInfo) SetConfiguration(v StructuresCopyInfoConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *StructuresCopyInfo) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


