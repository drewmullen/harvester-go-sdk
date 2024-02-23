# HarvesterhciIoV1beta1VirtualMachineImageStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedUrl** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] [default to 0]
**LastFailedTime** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**StorageClassName** | Pointer to **string** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineImageStatus

`func NewHarvesterhciIoV1beta1VirtualMachineImageStatus() *HarvesterhciIoV1beta1VirtualMachineImageStatus`

NewHarvesterhciIoV1beta1VirtualMachineImageStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineImageStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineImageStatusWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineImageStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImageStatus`

NewHarvesterhciIoV1beta1VirtualMachineImageStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImageStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetAppliedUrl() string`

GetAppliedUrl returns the AppliedUrl field if non-nil, zero value otherwise.

### GetAppliedUrlOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetAppliedUrlOk() (*string, bool)`

GetAppliedUrlOk returns a tuple with the AppliedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetAppliedUrl(v string)`

SetAppliedUrl sets AppliedUrl field to given value.

### HasAppliedUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasAppliedUrl() bool`

HasAppliedUrl returns a boolean if a field has been set.

### GetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFailed

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetLastFailedTime

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetLastFailedTime() string`

GetLastFailedTime returns the LastFailedTime field if non-nil, zero value otherwise.

### GetLastFailedTimeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetLastFailedTimeOk() (*string, bool)`

GetLastFailedTimeOk returns a tuple with the LastFailedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedTime

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetLastFailedTime(v string)`

SetLastFailedTime sets LastFailedTime field to given value.

### HasLastFailedTime

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasLastFailedTime() bool`

HasLastFailedTime returns a boolean if a field has been set.

### GetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSize

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageClassName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


