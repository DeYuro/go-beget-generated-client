# StructuresOperatingSystemRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **int32** | Количество ядер процессора, шт | [optional] 
**DiskSize** | Pointer to **int32** | Размер диска, мб | [optional] 
**Memory** | Pointer to **int32** | Объем оперативной памяти, мб | [optional] 

## Methods

### NewStructuresOperatingSystemRequirements

`func NewStructuresOperatingSystemRequirements() *StructuresOperatingSystemRequirements`

NewStructuresOperatingSystemRequirements instantiates a new StructuresOperatingSystemRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresOperatingSystemRequirementsWithDefaults

`func NewStructuresOperatingSystemRequirementsWithDefaults() *StructuresOperatingSystemRequirements`

NewStructuresOperatingSystemRequirementsWithDefaults instantiates a new StructuresOperatingSystemRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *StructuresOperatingSystemRequirements) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *StructuresOperatingSystemRequirements) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *StructuresOperatingSystemRequirements) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *StructuresOperatingSystemRequirements) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetDiskSize

`func (o *StructuresOperatingSystemRequirements) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *StructuresOperatingSystemRequirements) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *StructuresOperatingSystemRequirements) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *StructuresOperatingSystemRequirements) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMemory

`func (o *StructuresOperatingSystemRequirements) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *StructuresOperatingSystemRequirements) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *StructuresOperatingSystemRequirements) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *StructuresOperatingSystemRequirements) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


