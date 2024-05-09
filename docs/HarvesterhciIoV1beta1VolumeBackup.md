# HarvesterhciIoV1beta1VolumeBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **string** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] [default to ""]
**CsiDriverName** | **string** |  | [default to ""]
**Error** | Pointer to [**HarvesterhciIoV1beta1Error**](HarvesterhciIoV1beta1Error.md) |  | [optional] 
**LonghornBackupName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PersistentVolumeClaim** | [**HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec**](HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec.md) |  | [default to {}]
**Progress** | Pointer to **int32** |  | [optional] 
**ReadyToUse** | Pointer to **bool** |  | [optional] 
**VolumeName** | **string** |  | [default to ""]
**VolumeSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VolumeBackup

`func NewHarvesterhciIoV1beta1VolumeBackup(csiDriverName string, persistentVolumeClaim HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec, volumeName string, ) *HarvesterhciIoV1beta1VolumeBackup`

NewHarvesterhciIoV1beta1VolumeBackup instantiates a new HarvesterhciIoV1beta1VolumeBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VolumeBackupWithDefaults

`func NewHarvesterhciIoV1beta1VolumeBackupWithDefaults() *HarvesterhciIoV1beta1VolumeBackup`

NewHarvesterhciIoV1beta1VolumeBackupWithDefaults instantiates a new HarvesterhciIoV1beta1VolumeBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCsiDriverName

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetCsiDriverName() string`

GetCsiDriverName returns the CsiDriverName field if non-nil, zero value otherwise.

### GetCsiDriverNameOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetCsiDriverNameOk() (*string, bool)`

GetCsiDriverNameOk returns a tuple with the CsiDriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiDriverName

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetCsiDriverName(v string)`

SetCsiDriverName sets CsiDriverName field to given value.


### GetError

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetError() HarvesterhciIoV1beta1Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetErrorOk() (*HarvesterhciIoV1beta1Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetError(v HarvesterhciIoV1beta1Error)`

SetError sets Error field to given value.

### HasError

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLonghornBackupName

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetLonghornBackupName() string`

GetLonghornBackupName returns the LonghornBackupName field if non-nil, zero value otherwise.

### GetLonghornBackupNameOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetLonghornBackupNameOk() (*string, bool)`

GetLonghornBackupNameOk returns a tuple with the LonghornBackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLonghornBackupName

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetLonghornBackupName(v string)`

SetLonghornBackupName sets LonghornBackupName field to given value.

### HasLonghornBackupName

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasLonghornBackupName() bool`

HasLonghornBackupName returns a boolean if a field has been set.

### GetName

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetPersistentVolumeClaim() HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetPersistentVolumeClaimOk() (*HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetPersistentVolumeClaim(v HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.


### GetProgress

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetReadyToUse

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetReadyToUse() bool`

GetReadyToUse returns the ReadyToUse field if non-nil, zero value otherwise.

### GetReadyToUseOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetReadyToUseOk() (*bool, bool)`

GetReadyToUseOk returns a tuple with the ReadyToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToUse

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetReadyToUse(v bool)`

SetReadyToUse sets ReadyToUse field to given value.

### HasReadyToUse

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasReadyToUse() bool`

HasReadyToUse returns a boolean if a field has been set.

### GetVolumeName

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.


### GetVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *HarvesterhciIoV1beta1VolumeBackup) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeBackup) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeBackup) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


