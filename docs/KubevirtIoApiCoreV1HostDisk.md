# KubevirtIoApiCoreV1HostDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **string** |  | [optional] [default to "{}"]
**Path** | **string** |  | [default to ""]
**Shared** | Pointer to **bool** |  | [optional] 
**Type** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1HostDisk

`func NewKubevirtIoApiCoreV1HostDisk(path string, type_ string, ) *KubevirtIoApiCoreV1HostDisk`

NewKubevirtIoApiCoreV1HostDisk instantiates a new KubevirtIoApiCoreV1HostDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1HostDiskWithDefaults

`func NewKubevirtIoApiCoreV1HostDiskWithDefaults() *KubevirtIoApiCoreV1HostDisk`

NewKubevirtIoApiCoreV1HostDiskWithDefaults instantiates a new KubevirtIoApiCoreV1HostDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *KubevirtIoApiCoreV1HostDisk) GetCapacity() string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *KubevirtIoApiCoreV1HostDisk) GetCapacityOk() (*string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *KubevirtIoApiCoreV1HostDisk) SetCapacity(v string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *KubevirtIoApiCoreV1HostDisk) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetPath

`func (o *KubevirtIoApiCoreV1HostDisk) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KubevirtIoApiCoreV1HostDisk) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KubevirtIoApiCoreV1HostDisk) SetPath(v string)`

SetPath sets Path field to given value.


### GetShared

`func (o *KubevirtIoApiCoreV1HostDisk) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *KubevirtIoApiCoreV1HostDisk) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *KubevirtIoApiCoreV1HostDisk) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *KubevirtIoApiCoreV1HostDisk) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetType

`func (o *KubevirtIoApiCoreV1HostDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubevirtIoApiCoreV1HostDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubevirtIoApiCoreV1HostDisk) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


