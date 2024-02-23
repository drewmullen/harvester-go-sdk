# KubevirtIoApiCoreV1PersistentVolumeClaimInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** |  | [optional] 
**Capacity** | Pointer to **map[string]string** |  | [optional] 
**FilesystemOverhead** | Pointer to **string** |  | [optional] 
**Preallocated** | Pointer to **bool** |  | [optional] 
**Requests** | Pointer to **map[string]string** |  | [optional] 
**VolumeMode** | Pointer to **string** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1PersistentVolumeClaimInfo

`func NewKubevirtIoApiCoreV1PersistentVolumeClaimInfo() *KubevirtIoApiCoreV1PersistentVolumeClaimInfo`

NewKubevirtIoApiCoreV1PersistentVolumeClaimInfo instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1PersistentVolumeClaimInfoWithDefaults

`func NewKubevirtIoApiCoreV1PersistentVolumeClaimInfoWithDefaults() *KubevirtIoApiCoreV1PersistentVolumeClaimInfo`

NewKubevirtIoApiCoreV1PersistentVolumeClaimInfoWithDefaults instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetCapacity

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetCapacity() map[string]string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetCapacityOk() (*map[string]string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetCapacity(v map[string]string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetFilesystemOverhead

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetFilesystemOverhead() string`

GetFilesystemOverhead returns the FilesystemOverhead field if non-nil, zero value otherwise.

### GetFilesystemOverheadOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetFilesystemOverheadOk() (*string, bool)`

GetFilesystemOverheadOk returns a tuple with the FilesystemOverhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemOverhead

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetFilesystemOverhead(v string)`

SetFilesystemOverhead sets FilesystemOverhead field to given value.

### HasFilesystemOverhead

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasFilesystemOverhead() bool`

HasFilesystemOverhead returns a boolean if a field has been set.

### GetPreallocated

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetPreallocated() bool`

GetPreallocated returns the Preallocated field if non-nil, zero value otherwise.

### GetPreallocatedOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetPreallocatedOk() (*bool, bool)`

GetPreallocatedOk returns a tuple with the Preallocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocated

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetPreallocated(v bool)`

SetPreallocated sets Preallocated field to given value.

### HasPreallocated

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasPreallocated() bool`

HasPreallocated returns a boolean if a field has been set.

### GetRequests

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetRequests() map[string]string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetRequestsOk() (*map[string]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetRequests(v map[string]string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetVolumeMode

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


