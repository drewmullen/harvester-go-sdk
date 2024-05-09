# HarvesterhciIoV1beta1VirtualMachineBackupStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupTarget** | Pointer to [**HarvesterhciIoV1beta1BackupTarget**](HarvesterhciIoV1beta1BackupTarget.md) |  | [optional] 
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**CreationTime** | Pointer to **string** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] [default to ""]
**CsiDriverVolumeSnapshotClassNames** | Pointer to **map[string]string** |  | [optional] 
**Error** | Pointer to [**HarvesterhciIoV1beta1Error**](HarvesterhciIoV1beta1Error.md) |  | [optional] 
**Progress** | Pointer to **int32** |  | [optional] 
**ReadyToUse** | Pointer to **bool** |  | [optional] 
**SecretBackups** | Pointer to [**[]HarvesterhciIoV1beta1SecretBackup**](HarvesterhciIoV1beta1SecretBackup.md) |  | [optional] 
**Source** | Pointer to [**HarvesterhciIoV1beta1VirtualMachineSourceSpec**](HarvesterhciIoV1beta1VirtualMachineSourceSpec.md) | SourceSpec contains the vm spec source of the backup target | [optional] 
**SourceUID** | Pointer to **string** |  | [optional] 
**VolumeBackups** | Pointer to [**[]HarvesterhciIoV1beta1VolumeBackup**](HarvesterhciIoV1beta1VolumeBackup.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineBackupStatus

`func NewHarvesterhciIoV1beta1VirtualMachineBackupStatus() *HarvesterhciIoV1beta1VirtualMachineBackupStatus`

NewHarvesterhciIoV1beta1VirtualMachineBackupStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineBackupStatusWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineBackupStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineBackupStatus`

NewHarvesterhciIoV1beta1VirtualMachineBackupStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupTarget

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetBackupTarget() HarvesterhciIoV1beta1BackupTarget`

GetBackupTarget returns the BackupTarget field if non-nil, zero value otherwise.

### GetBackupTargetOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetBackupTargetOk() (*HarvesterhciIoV1beta1BackupTarget, bool)`

GetBackupTargetOk returns a tuple with the BackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupTarget

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetBackupTarget(v HarvesterhciIoV1beta1BackupTarget)`

SetBackupTarget sets BackupTarget field to given value.

### HasBackupTarget

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasBackupTarget() bool`

HasBackupTarget returns a boolean if a field has been set.

### GetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreationTime

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCsiDriverVolumeSnapshotClassNames

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCsiDriverVolumeSnapshotClassNames() map[string]string`

GetCsiDriverVolumeSnapshotClassNames returns the CsiDriverVolumeSnapshotClassNames field if non-nil, zero value otherwise.

### GetCsiDriverVolumeSnapshotClassNamesOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCsiDriverVolumeSnapshotClassNamesOk() (*map[string]string, bool)`

GetCsiDriverVolumeSnapshotClassNamesOk returns a tuple with the CsiDriverVolumeSnapshotClassNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsiDriverVolumeSnapshotClassNames

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetCsiDriverVolumeSnapshotClassNames(v map[string]string)`

SetCsiDriverVolumeSnapshotClassNames sets CsiDriverVolumeSnapshotClassNames field to given value.

### HasCsiDriverVolumeSnapshotClassNames

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasCsiDriverVolumeSnapshotClassNames() bool`

HasCsiDriverVolumeSnapshotClassNames returns a boolean if a field has been set.

### GetError

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetError() HarvesterhciIoV1beta1Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetErrorOk() (*HarvesterhciIoV1beta1Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetError(v HarvesterhciIoV1beta1Error)`

SetError sets Error field to given value.

### HasError

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetReadyToUse

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetReadyToUse() bool`

GetReadyToUse returns the ReadyToUse field if non-nil, zero value otherwise.

### GetReadyToUseOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetReadyToUseOk() (*bool, bool)`

GetReadyToUseOk returns a tuple with the ReadyToUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyToUse

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetReadyToUse(v bool)`

SetReadyToUse sets ReadyToUse field to given value.

### HasReadyToUse

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasReadyToUse() bool`

HasReadyToUse returns a boolean if a field has been set.

### GetSecretBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSecretBackups() []HarvesterhciIoV1beta1SecretBackup`

GetSecretBackups returns the SecretBackups field if non-nil, zero value otherwise.

### GetSecretBackupsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSecretBackupsOk() (*[]HarvesterhciIoV1beta1SecretBackup, bool)`

GetSecretBackupsOk returns a tuple with the SecretBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSecretBackups(v []HarvesterhciIoV1beta1SecretBackup)`

SetSecretBackups sets SecretBackups field to given value.

### HasSecretBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSecretBackups() bool`

HasSecretBackups returns a boolean if a field has been set.

### GetSource

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSource() HarvesterhciIoV1beta1VirtualMachineSourceSpec`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceOk() (*HarvesterhciIoV1beta1VirtualMachineSourceSpec, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSource(v HarvesterhciIoV1beta1VirtualMachineSourceSpec)`

SetSource sets Source field to given value.

### HasSource

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceUID

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceUID() string`

GetSourceUID returns the SourceUID field if non-nil, zero value otherwise.

### GetSourceUIDOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceUIDOk() (*string, bool)`

GetSourceUIDOk returns a tuple with the SourceUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUID

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSourceUID(v string)`

SetSourceUID sets SourceUID field to given value.

### HasSourceUID

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSourceUID() bool`

HasSourceUID returns a boolean if a field has been set.

### GetVolumeBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetVolumeBackups() []HarvesterhciIoV1beta1VolumeBackup`

GetVolumeBackups returns the VolumeBackups field if non-nil, zero value otherwise.

### GetVolumeBackupsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetVolumeBackupsOk() (*[]HarvesterhciIoV1beta1VolumeBackup, bool)`

GetVolumeBackupsOk returns a tuple with the VolumeBackups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetVolumeBackups(v []HarvesterhciIoV1beta1VolumeBackup)`

SetVolumeBackups sets VolumeBackups field to given value.

### HasVolumeBackups

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasVolumeBackups() bool`

HasVolumeBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


