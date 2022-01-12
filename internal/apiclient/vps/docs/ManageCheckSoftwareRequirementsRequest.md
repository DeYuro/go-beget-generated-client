# ManageCheckSoftwareRequirementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingSystemId** | Pointer to **int32** | Идентификатор ОС | [optional] 
**Info** | Pointer to [**ManageSoftwareInstallInfo**](ManageSoftwareInstallInfo.md) |  | [optional] 

## Methods

### NewManageCheckSoftwareRequirementsRequest

`func NewManageCheckSoftwareRequirementsRequest() *ManageCheckSoftwareRequirementsRequest`

NewManageCheckSoftwareRequirementsRequest instantiates a new ManageCheckSoftwareRequirementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManageCheckSoftwareRequirementsRequestWithDefaults

`func NewManageCheckSoftwareRequirementsRequestWithDefaults() *ManageCheckSoftwareRequirementsRequest`

NewManageCheckSoftwareRequirementsRequestWithDefaults instantiates a new ManageCheckSoftwareRequirementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingSystemId

`func (o *ManageCheckSoftwareRequirementsRequest) GetOperatingSystemId() int32`

GetOperatingSystemId returns the OperatingSystemId field if non-nil, zero value otherwise.

### GetOperatingSystemIdOk

`func (o *ManageCheckSoftwareRequirementsRequest) GetOperatingSystemIdOk() (*int32, bool)`

GetOperatingSystemIdOk returns a tuple with the OperatingSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemId

`func (o *ManageCheckSoftwareRequirementsRequest) SetOperatingSystemId(v int32)`

SetOperatingSystemId sets OperatingSystemId field to given value.

### HasOperatingSystemId

`func (o *ManageCheckSoftwareRequirementsRequest) HasOperatingSystemId() bool`

HasOperatingSystemId returns a boolean if a field has been set.

### GetInfo

`func (o *ManageCheckSoftwareRequirementsRequest) GetInfo() ManageSoftwareInstallInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManageCheckSoftwareRequirementsRequest) GetInfoOk() (*ManageSoftwareInstallInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManageCheckSoftwareRequirementsRequest) SetInfo(v ManageSoftwareInstallInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManageCheckSoftwareRequirementsRequest) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


