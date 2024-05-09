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

// checks if the HarvesterhciIoV1beta1VirtualMachineRestoreStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineRestoreStatus{}

// HarvesterhciIoV1beta1VirtualMachineRestoreStatus VirtualMachineRestoreStatus is the spec for a VirtualMachineRestore resource
type HarvesterhciIoV1beta1VirtualMachineRestoreStatus struct {
	Complete *bool `json:"complete,omitempty"`
	Conditions []HarvesterhciIoV1beta1Condition `json:"conditions,omitempty"`
	DeletedVolumes []string `json:"deletedVolumes,omitempty"`
	Progress *int32 `json:"progress,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	RestoreTime *string `json:"restoreTime,omitempty"`
	Restores []HarvesterhciIoV1beta1VolumeRestore `json:"restores,omitempty"`
	TargetUID *string `json:"targetUID,omitempty"`
}

// NewHarvesterhciIoV1beta1VirtualMachineRestoreStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineRestoreStatus() *HarvesterhciIoV1beta1VirtualMachineRestoreStatus {
	this := HarvesterhciIoV1beta1VirtualMachineRestoreStatus{}
	var restoreTime string = ""
	this.RestoreTime = &restoreTime
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineRestoreStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineRestoreStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineRestoreStatus {
	this := HarvesterhciIoV1beta1VirtualMachineRestoreStatus{}
	var restoreTime string = ""
	this.RestoreTime = &restoreTime
	return &this
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetComplete(v bool) {
	o.Complete = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetConditions() []HarvesterhciIoV1beta1Condition {
	if o == nil || IsNil(o.Conditions) {
		var ret []HarvesterhciIoV1beta1Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetConditionsOk() ([]HarvesterhciIoV1beta1Condition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []HarvesterhciIoV1beta1Condition and assigns it to the Conditions field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetConditions(v []HarvesterhciIoV1beta1Condition) {
	o.Conditions = v
}

// GetDeletedVolumes returns the DeletedVolumes field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetDeletedVolumes() []string {
	if o == nil || IsNil(o.DeletedVolumes) {
		var ret []string
		return ret
	}
	return o.DeletedVolumes
}

// GetDeletedVolumesOk returns a tuple with the DeletedVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetDeletedVolumesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeletedVolumes) {
		return nil, false
	}
	return o.DeletedVolumes, true
}

// HasDeletedVolumes returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasDeletedVolumes() bool {
	if o != nil && !IsNil(o.DeletedVolumes) {
		return true
	}

	return false
}

// SetDeletedVolumes gets a reference to the given []string and assigns it to the DeletedVolumes field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetDeletedVolumes(v []string) {
	o.DeletedVolumes = v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetProgress(v int32) {
	o.Progress = &v
}

// GetRestoreTime returns the RestoreTime field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoreTime() string {
	if o == nil || IsNil(o.RestoreTime) {
		var ret string
		return ret
	}
	return *o.RestoreTime
}

// GetRestoreTimeOk returns a tuple with the RestoreTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoreTimeOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreTime) {
		return nil, false
	}
	return o.RestoreTime, true
}

// HasRestoreTime returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasRestoreTime() bool {
	if o != nil && !IsNil(o.RestoreTime) {
		return true
	}

	return false
}

// SetRestoreTime gets a reference to the given string and assigns it to the RestoreTime field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetRestoreTime(v string) {
	o.RestoreTime = &v
}

// GetRestores returns the Restores field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestores() []HarvesterhciIoV1beta1VolumeRestore {
	if o == nil || IsNil(o.Restores) {
		var ret []HarvesterhciIoV1beta1VolumeRestore
		return ret
	}
	return o.Restores
}

// GetRestoresOk returns a tuple with the Restores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetRestoresOk() ([]HarvesterhciIoV1beta1VolumeRestore, bool) {
	if o == nil || IsNil(o.Restores) {
		return nil, false
	}
	return o.Restores, true
}

// HasRestores returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasRestores() bool {
	if o != nil && !IsNil(o.Restores) {
		return true
	}

	return false
}

// SetRestores gets a reference to the given []HarvesterhciIoV1beta1VolumeRestore and assigns it to the Restores field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetRestores(v []HarvesterhciIoV1beta1VolumeRestore) {
	o.Restores = v
}

// GetTargetUID returns the TargetUID field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetTargetUID() string {
	if o == nil || IsNil(o.TargetUID) {
		var ret string
		return ret
	}
	return *o.TargetUID
}

// GetTargetUIDOk returns a tuple with the TargetUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) GetTargetUIDOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUID) {
		return nil, false
	}
	return o.TargetUID, true
}

// HasTargetUID returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) HasTargetUID() bool {
	if o != nil && !IsNil(o.TargetUID) {
		return true
	}

	return false
}

// SetTargetUID gets a reference to the given string and assigns it to the TargetUID field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) SetTargetUID(v string) {
	o.TargetUID = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineRestoreStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineRestoreStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.DeletedVolumes) {
		toSerialize["deletedVolumes"] = o.DeletedVolumes
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.RestoreTime) {
		toSerialize["restoreTime"] = o.RestoreTime
	}
	if !IsNil(o.Restores) {
		toSerialize["restores"] = o.Restores
	}
	if !IsNil(o.TargetUID) {
		toSerialize["targetUID"] = o.TargetUID
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus struct {
	value *HarvesterhciIoV1beta1VirtualMachineRestoreStatus
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) Get() *HarvesterhciIoV1beta1VirtualMachineRestoreStatus {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) Set(val *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus(val *HarvesterhciIoV1beta1VirtualMachineRestoreStatus) *NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus {
	return &NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestoreStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


