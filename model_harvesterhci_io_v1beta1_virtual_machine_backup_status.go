/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
)

// checks if the HarvesterhciIoV1beta1VirtualMachineBackupStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineBackupStatus{}

// HarvesterhciIoV1beta1VirtualMachineBackupStatus struct for HarvesterhciIoV1beta1VirtualMachineBackupStatus
type HarvesterhciIoV1beta1VirtualMachineBackupStatus struct {
	BackupTarget *HarvesterhciIoV1beta1BackupTarget `json:"backupTarget,omitempty"`
	Conditions []HarvesterhciIoV1beta1Condition `json:"conditions,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
	CsiDriverVolumeSnapshotClassNames *map[string]string `json:"csiDriverVolumeSnapshotClassNames,omitempty"`
	Error *HarvesterhciIoV1beta1Error `json:"error,omitempty"`
	Progress *int32 `json:"progress,omitempty"`
	ReadyToUse *bool `json:"readyToUse,omitempty"`
	SecretBackups []HarvesterhciIoV1beta1SecretBackup `json:"secretBackups,omitempty"`
	Source *HarvesterhciIoV1beta1VirtualMachineSourceSpec `json:"source,omitempty"`
	SourceUID *string `json:"sourceUID,omitempty"`
	VolumeBackups []HarvesterhciIoV1beta1VolumeBackup `json:"volumeBackups,omitempty"`
}

// NewHarvesterhciIoV1beta1VirtualMachineBackupStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineBackupStatus() *HarvesterhciIoV1beta1VirtualMachineBackupStatus {
	this := HarvesterhciIoV1beta1VirtualMachineBackupStatus{}
	var creationTime string = ""
	this.CreationTime = &creationTime
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineBackupStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineBackupStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineBackupStatus {
	this := HarvesterhciIoV1beta1VirtualMachineBackupStatus{}
	var creationTime string = ""
	this.CreationTime = &creationTime
	return &this
}

// GetBackupTarget returns the BackupTarget field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetBackupTarget() HarvesterhciIoV1beta1BackupTarget {
	if o == nil || IsNil(o.BackupTarget) {
		var ret HarvesterhciIoV1beta1BackupTarget
		return ret
	}
	return *o.BackupTarget
}

// GetBackupTargetOk returns a tuple with the BackupTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetBackupTargetOk() (*HarvesterhciIoV1beta1BackupTarget, bool) {
	if o == nil || IsNil(o.BackupTarget) {
		return nil, false
	}
	return o.BackupTarget, true
}

// HasBackupTarget returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasBackupTarget() bool {
	if o != nil && !IsNil(o.BackupTarget) {
		return true
	}

	return false
}

// SetBackupTarget gets a reference to the given HarvesterhciIoV1beta1BackupTarget and assigns it to the BackupTarget field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetBackupTarget(v HarvesterhciIoV1beta1BackupTarget) {
	o.BackupTarget = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetConditions() []HarvesterhciIoV1beta1Condition {
	if o == nil || IsNil(o.Conditions) {
		var ret []HarvesterhciIoV1beta1Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetConditionsOk() ([]HarvesterhciIoV1beta1Condition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []HarvesterhciIoV1beta1Condition and assigns it to the Conditions field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetConditions(v []HarvesterhciIoV1beta1Condition) {
	o.Conditions = v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCreationTime() string {
	if o == nil || IsNil(o.CreationTime) {
		var ret string
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCreationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given string and assigns it to the CreationTime field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetCreationTime(v string) {
	o.CreationTime = &v
}

// GetCsiDriverVolumeSnapshotClassNames returns the CsiDriverVolumeSnapshotClassNames field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCsiDriverVolumeSnapshotClassNames() map[string]string {
	if o == nil || IsNil(o.CsiDriverVolumeSnapshotClassNames) {
		var ret map[string]string
		return ret
	}
	return *o.CsiDriverVolumeSnapshotClassNames
}

// GetCsiDriverVolumeSnapshotClassNamesOk returns a tuple with the CsiDriverVolumeSnapshotClassNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetCsiDriverVolumeSnapshotClassNamesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CsiDriverVolumeSnapshotClassNames) {
		return nil, false
	}
	return o.CsiDriverVolumeSnapshotClassNames, true
}

// HasCsiDriverVolumeSnapshotClassNames returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasCsiDriverVolumeSnapshotClassNames() bool {
	if o != nil && !IsNil(o.CsiDriverVolumeSnapshotClassNames) {
		return true
	}

	return false
}

// SetCsiDriverVolumeSnapshotClassNames gets a reference to the given map[string]string and assigns it to the CsiDriverVolumeSnapshotClassNames field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetCsiDriverVolumeSnapshotClassNames(v map[string]string) {
	o.CsiDriverVolumeSnapshotClassNames = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetError() HarvesterhciIoV1beta1Error {
	if o == nil || IsNil(o.Error) {
		var ret HarvesterhciIoV1beta1Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetErrorOk() (*HarvesterhciIoV1beta1Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given HarvesterhciIoV1beta1Error and assigns it to the Error field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetError(v HarvesterhciIoV1beta1Error) {
	o.Error = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetProgress(v int32) {
	o.Progress = &v
}

// GetReadyToUse returns the ReadyToUse field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetReadyToUse() bool {
	if o == nil || IsNil(o.ReadyToUse) {
		var ret bool
		return ret
	}
	return *o.ReadyToUse
}

// GetReadyToUseOk returns a tuple with the ReadyToUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetReadyToUseOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadyToUse) {
		return nil, false
	}
	return o.ReadyToUse, true
}

// HasReadyToUse returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasReadyToUse() bool {
	if o != nil && !IsNil(o.ReadyToUse) {
		return true
	}

	return false
}

// SetReadyToUse gets a reference to the given bool and assigns it to the ReadyToUse field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetReadyToUse(v bool) {
	o.ReadyToUse = &v
}

// GetSecretBackups returns the SecretBackups field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSecretBackups() []HarvesterhciIoV1beta1SecretBackup {
	if o == nil || IsNil(o.SecretBackups) {
		var ret []HarvesterhciIoV1beta1SecretBackup
		return ret
	}
	return o.SecretBackups
}

// GetSecretBackupsOk returns a tuple with the SecretBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSecretBackupsOk() ([]HarvesterhciIoV1beta1SecretBackup, bool) {
	if o == nil || IsNil(o.SecretBackups) {
		return nil, false
	}
	return o.SecretBackups, true
}

// HasSecretBackups returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSecretBackups() bool {
	if o != nil && !IsNil(o.SecretBackups) {
		return true
	}

	return false
}

// SetSecretBackups gets a reference to the given []HarvesterhciIoV1beta1SecretBackup and assigns it to the SecretBackups field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSecretBackups(v []HarvesterhciIoV1beta1SecretBackup) {
	o.SecretBackups = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSource() HarvesterhciIoV1beta1VirtualMachineSourceSpec {
	if o == nil || IsNil(o.Source) {
		var ret HarvesterhciIoV1beta1VirtualMachineSourceSpec
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceOk() (*HarvesterhciIoV1beta1VirtualMachineSourceSpec, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given HarvesterhciIoV1beta1VirtualMachineSourceSpec and assigns it to the Source field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSource(v HarvesterhciIoV1beta1VirtualMachineSourceSpec) {
	o.Source = &v
}

// GetSourceUID returns the SourceUID field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceUID() string {
	if o == nil || IsNil(o.SourceUID) {
		var ret string
		return ret
	}
	return *o.SourceUID
}

// GetSourceUIDOk returns a tuple with the SourceUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetSourceUIDOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUID) {
		return nil, false
	}
	return o.SourceUID, true
}

// HasSourceUID returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasSourceUID() bool {
	if o != nil && !IsNil(o.SourceUID) {
		return true
	}

	return false
}

// SetSourceUID gets a reference to the given string and assigns it to the SourceUID field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetSourceUID(v string) {
	o.SourceUID = &v
}

// GetVolumeBackups returns the VolumeBackups field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetVolumeBackups() []HarvesterhciIoV1beta1VolumeBackup {
	if o == nil || IsNil(o.VolumeBackups) {
		var ret []HarvesterhciIoV1beta1VolumeBackup
		return ret
	}
	return o.VolumeBackups
}

// GetVolumeBackupsOk returns a tuple with the VolumeBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) GetVolumeBackupsOk() ([]HarvesterhciIoV1beta1VolumeBackup, bool) {
	if o == nil || IsNil(o.VolumeBackups) {
		return nil, false
	}
	return o.VolumeBackups, true
}

// HasVolumeBackups returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) HasVolumeBackups() bool {
	if o != nil && !IsNil(o.VolumeBackups) {
		return true
	}

	return false
}

// SetVolumeBackups gets a reference to the given []HarvesterhciIoV1beta1VolumeBackup and assigns it to the VolumeBackups field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackupStatus) SetVolumeBackups(v []HarvesterhciIoV1beta1VolumeBackup) {
	o.VolumeBackups = v
}

func (o HarvesterhciIoV1beta1VirtualMachineBackupStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineBackupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackupTarget) {
		toSerialize["backupTarget"] = o.BackupTarget
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creationTime"] = o.CreationTime
	}
	if !IsNil(o.CsiDriverVolumeSnapshotClassNames) {
		toSerialize["csiDriverVolumeSnapshotClassNames"] = o.CsiDriverVolumeSnapshotClassNames
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.ReadyToUse) {
		toSerialize["readyToUse"] = o.ReadyToUse
	}
	if !IsNil(o.SecretBackups) {
		toSerialize["secretBackups"] = o.SecretBackups
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.SourceUID) {
		toSerialize["sourceUID"] = o.SourceUID
	}
	if !IsNil(o.VolumeBackups) {
		toSerialize["volumeBackups"] = o.VolumeBackups
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus struct {
	value *HarvesterhciIoV1beta1VirtualMachineBackupStatus
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) Get() *HarvesterhciIoV1beta1VirtualMachineBackupStatus {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) Set(val *HarvesterhciIoV1beta1VirtualMachineBackupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineBackupStatus(val *HarvesterhciIoV1beta1VirtualMachineBackupStatus) *NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus {
	return &NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


