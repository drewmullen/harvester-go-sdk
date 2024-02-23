# HarvesterhciIoV1beta1VirtualMachineRestoreStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complete** | Pointer to **bool** |  | [optional] 
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**DeletedVolumes** | Pointer to **[]string** |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**RestoreTime** | Pointer to **string** |  | [optional] [default to ""]
**Restores** | Pointer to [**[]HarvesterhciIoV1beta1VolumeRestore**](HarvesterhciIoV1beta1VolumeRestore.md) |  | [optional] 
**TargetUID** | Pointer to **string** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineRestoreStatus

`func NewHarvesterhciIoV1beta1VirtualMachineRestoreStatus() *HarvesterhciIoV1beta1VirtualMachineRestoreStatus`

NewHarvesterhciIoV1beta1VirtualMachineRestoreStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineRestoreStatusWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineRestoreStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineRestoreStatus`

NewHarvesterhciIoV1beta1VirtualMachineRestoreStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplete

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetDeletedVolumes

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetDeletedVolumes() []string`

GetDeletedVolumes returns the DeletedVolumes field if non-nil, zero value otherwise.

### GetDeletedVolumesOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetDeletedVolumesOk() (*[]string, bool)`

GetDeletedVolumesOk returns a tuple with the DeletedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedVolumes

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetDeletedVolumes(v []string)`

SetDeletedVolumes sets DeletedVolumes field to given value.

### HasDeletedVolumes

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasDeletedVolumes() bool`

HasDeletedVolumes returns a boolean if a field has been set.

### GetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRestoreTime

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoreTime() string`

GetRestoreTime returns the RestoreTime field if non-nil, zero value otherwise.

### GetRestoreTimeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoreTimeOk() (*string, bool)`

GetRestoreTimeOk returns a tuple with the RestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTime

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetRestoreTime(v string)`

SetRestoreTime sets RestoreTime field to given value.

### HasRestoreTime

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasRestoreTime() bool`

HasRestoreTime returns a boolean if a field has been set.

### GetRestores

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestores() []HarvesterhciIoV1beta1VolumeRestore`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### GetRestoresOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoresOk() (*[]HarvesterhciIoV1beta1VolumeRestore, bool)`

GetRestoresOk returns a tuple with the Restores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestores

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetRestores(v []HarvesterhciIoV1beta1VolumeRestore)`

SetRestores sets Restores field to given value.

### HasRestores

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasRestores() bool`

HasRestores returns a boolean if a field has been set.

### GetTargetUID

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetTargetUID() string`

GetTargetUID returns the TargetUID field if non-nil, zero value otherwise.

### GetTargetUIDOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetTargetUIDOk() (*string, bool)`

GetTargetUIDOk returns a tuple with the TargetUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUID

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetTargetUID(v string)`

SetTargetUID sets TargetUID field to given value.

### HasTargetUID

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasTargetUID() bool`

HasTargetUID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


