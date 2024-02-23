# HarvesterhciIoV1beta1VolumeRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LonghornEngineName** | Pointer to **string** |  | [optional] 
**PersistentVolumeClaimSpec** | Pointer to [**HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec**](HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec.md) |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**VolumeBackupName** | Pointer to **string** |  | [optional] 
**VolumeName** | Pointer to **string** |  | [optional] 
**VolumeSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VolumeRestore

`func NewHarvesterhciIoV1beta1VolumeRestore() *HarvesterhciIoV1beta1VolumeRestore`

NewHarvesterhciIoV1beta1VolumeRestore instantiates a new HarvesterhciIoV1beta1VolumeRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VolumeRestoreWithDefaults

`func NewHarvesterhciIoV1beta1VolumeRestoreWithDefaults() *HarvesterhciIoV1beta1VolumeRestore`

NewHarvesterhciIoV1beta1VolumeRestoreWithDefaults instantiates a new HarvesterhciIoV1beta1VolumeRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLonghornEngineName

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetLonghornEngineName() string`

GetLonghornEngineName returns the LonghornEngineName field if non-nil, zero value otherwise.

### GetLonghornEngineNameOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetLonghornEngineNameOk() (*string, bool)`

GetLonghornEngineNameOk returns a tuple with the LonghornEngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLonghornEngineName

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetLonghornEngineName(v string)`

SetLonghornEngineName sets LonghornEngineName field to given value.

### HasLonghornEngineName

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasLonghornEngineName() bool`

HasLonghornEngineName returns a boolean if a field has been set.

### GetPersistentVolumeClaimSpec

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetPersistentVolumeClaimSpec() HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec`

GetPersistentVolumeClaimSpec returns the PersistentVolumeClaimSpec field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimSpecOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetPersistentVolumeClaimSpecOk() (*HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec, bool)`

GetPersistentVolumeClaimSpecOk returns a tuple with the PersistentVolumeClaimSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaimSpec

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetPersistentVolumeClaimSpec(v HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec)`

SetPersistentVolumeClaimSpec sets PersistentVolumeClaimSpec field to given value.

### HasPersistentVolumeClaimSpec

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasPersistentVolumeClaimSpec() bool`

HasPersistentVolumeClaimSpec returns a boolean if a field has been set.

### GetProgress

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetVolumeBackupName

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeBackupName() string`

GetVolumeBackupName returns the VolumeBackupName field if non-nil, zero value otherwise.

### GetVolumeBackupNameOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeBackupNameOk() (*string, bool)`

GetVolumeBackupNameOk returns a tuple with the VolumeBackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBackupName

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetVolumeBackupName(v string)`

SetVolumeBackupName sets VolumeBackupName field to given value.

### HasVolumeBackupName

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasVolumeBackupName() bool`

HasVolumeBackupName returns a boolean if a field has been set.

### GetVolumeName

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.

### GetVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeSize() int64`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *HarvesterhciIoV1beta1VolumeRestore) GetVolumeSizeOk() (*int64, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeRestore) SetVolumeSize(v int64)`

SetVolumeSize sets VolumeSize field to given value.

### HasVolumeSize

`func (o *HarvesterhciIoV1beta1VolumeRestore) HasVolumeSize() bool`

HasVolumeSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


