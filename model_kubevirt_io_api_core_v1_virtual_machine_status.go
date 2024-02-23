/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1VirtualMachineStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineStatus{}

// KubevirtIoApiCoreV1VirtualMachineStatus struct for KubevirtIoApiCoreV1VirtualMachineStatus
type KubevirtIoApiCoreV1VirtualMachineStatus struct {
	Conditions []KubevirtIoApiCoreV1VirtualMachineCondition `json:"conditions,omitempty"`
	Created *bool `json:"created,omitempty"`
	DesiredGeneration *int64 `json:"desiredGeneration,omitempty"`
	MemoryDumpRequest *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest `json:"memoryDumpRequest,omitempty"`
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
	PrintableStatus *string `json:"printableStatus,omitempty"`
	Ready *bool `json:"ready,omitempty"`
	RestoreInProgress *string `json:"restoreInProgress,omitempty"`
	SnapshotInProgress *string `json:"snapshotInProgress,omitempty"`
	StartFailure *KubevirtIoApiCoreV1VirtualMachineStartFailure `json:"startFailure,omitempty"`
	StateChangeRequests []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest `json:"stateChangeRequests,omitempty"`
	VolumeRequests []KubevirtIoApiCoreV1VirtualMachineVolumeRequest `json:"volumeRequests,omitempty"`
	VolumeSnapshotStatuses []KubevirtIoApiCoreV1VolumeSnapshotStatus `json:"volumeSnapshotStatuses,omitempty"`
}

// NewKubevirtIoApiCoreV1VirtualMachineStatus instantiates a new KubevirtIoApiCoreV1VirtualMachineStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineStatus() *KubevirtIoApiCoreV1VirtualMachineStatus {
	this := KubevirtIoApiCoreV1VirtualMachineStatus{}
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineStatusWithDefaults() *KubevirtIoApiCoreV1VirtualMachineStatus {
	this := KubevirtIoApiCoreV1VirtualMachineStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetConditions() []KubevirtIoApiCoreV1VirtualMachineCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []KubevirtIoApiCoreV1VirtualMachineCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetConditionsOk() ([]KubevirtIoApiCoreV1VirtualMachineCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []KubevirtIoApiCoreV1VirtualMachineCondition and assigns it to the Conditions field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetConditions(v []KubevirtIoApiCoreV1VirtualMachineCondition) {
	o.Conditions = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetCreated() bool {
	if o == nil || IsNil(o.Created) {
		var ret bool
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetCreatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given bool and assigns it to the Created field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetCreated(v bool) {
	o.Created = &v
}

// GetDesiredGeneration returns the DesiredGeneration field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetDesiredGeneration() int64 {
	if o == nil || IsNil(o.DesiredGeneration) {
		var ret int64
		return ret
	}
	return *o.DesiredGeneration
}

// GetDesiredGenerationOk returns a tuple with the DesiredGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetDesiredGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.DesiredGeneration) {
		return nil, false
	}
	return o.DesiredGeneration, true
}

// HasDesiredGeneration returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasDesiredGeneration() bool {
	if o != nil && !IsNil(o.DesiredGeneration) {
		return true
	}

	return false
}

// SetDesiredGeneration gets a reference to the given int64 and assigns it to the DesiredGeneration field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetDesiredGeneration(v int64) {
	o.DesiredGeneration = &v
}

// GetMemoryDumpRequest returns the MemoryDumpRequest field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetMemoryDumpRequest() KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest {
	if o == nil || IsNil(o.MemoryDumpRequest) {
		var ret KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest
		return ret
	}
	return *o.MemoryDumpRequest
}

// GetMemoryDumpRequestOk returns a tuple with the MemoryDumpRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetMemoryDumpRequestOk() (*KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest, bool) {
	if o == nil || IsNil(o.MemoryDumpRequest) {
		return nil, false
	}
	return o.MemoryDumpRequest, true
}

// HasMemoryDumpRequest returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasMemoryDumpRequest() bool {
	if o != nil && !IsNil(o.MemoryDumpRequest) {
		return true
	}

	return false
}

// SetMemoryDumpRequest gets a reference to the given KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest and assigns it to the MemoryDumpRequest field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetMemoryDumpRequest(v KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) {
	o.MemoryDumpRequest = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetObservedGeneration() int64 {
	if o == nil || IsNil(o.ObservedGeneration) {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.ObservedGeneration) {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasObservedGeneration() bool {
	if o != nil && !IsNil(o.ObservedGeneration) {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetPrintableStatus returns the PrintableStatus field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetPrintableStatus() string {
	if o == nil || IsNil(o.PrintableStatus) {
		var ret string
		return ret
	}
	return *o.PrintableStatus
}

// GetPrintableStatusOk returns a tuple with the PrintableStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetPrintableStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PrintableStatus) {
		return nil, false
	}
	return o.PrintableStatus, true
}

// HasPrintableStatus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasPrintableStatus() bool {
	if o != nil && !IsNil(o.PrintableStatus) {
		return true
	}

	return false
}

// SetPrintableStatus gets a reference to the given string and assigns it to the PrintableStatus field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetPrintableStatus(v string) {
	o.PrintableStatus = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetReady(v bool) {
	o.Ready = &v
}

// GetRestoreInProgress returns the RestoreInProgress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetRestoreInProgress() string {
	if o == nil || IsNil(o.RestoreInProgress) {
		var ret string
		return ret
	}
	return *o.RestoreInProgress
}

// GetRestoreInProgressOk returns a tuple with the RestoreInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetRestoreInProgressOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreInProgress) {
		return nil, false
	}
	return o.RestoreInProgress, true
}

// HasRestoreInProgress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasRestoreInProgress() bool {
	if o != nil && !IsNil(o.RestoreInProgress) {
		return true
	}

	return false
}

// SetRestoreInProgress gets a reference to the given string and assigns it to the RestoreInProgress field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetRestoreInProgress(v string) {
	o.RestoreInProgress = &v
}

// GetSnapshotInProgress returns the SnapshotInProgress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetSnapshotInProgress() string {
	if o == nil || IsNil(o.SnapshotInProgress) {
		var ret string
		return ret
	}
	return *o.SnapshotInProgress
}

// GetSnapshotInProgressOk returns a tuple with the SnapshotInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetSnapshotInProgressOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotInProgress) {
		return nil, false
	}
	return o.SnapshotInProgress, true
}

// HasSnapshotInProgress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasSnapshotInProgress() bool {
	if o != nil && !IsNil(o.SnapshotInProgress) {
		return true
	}

	return false
}

// SetSnapshotInProgress gets a reference to the given string and assigns it to the SnapshotInProgress field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetSnapshotInProgress(v string) {
	o.SnapshotInProgress = &v
}

// GetStartFailure returns the StartFailure field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStartFailure() KubevirtIoApiCoreV1VirtualMachineStartFailure {
	if o == nil || IsNil(o.StartFailure) {
		var ret KubevirtIoApiCoreV1VirtualMachineStartFailure
		return ret
	}
	return *o.StartFailure
}

// GetStartFailureOk returns a tuple with the StartFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStartFailureOk() (*KubevirtIoApiCoreV1VirtualMachineStartFailure, bool) {
	if o == nil || IsNil(o.StartFailure) {
		return nil, false
	}
	return o.StartFailure, true
}

// HasStartFailure returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasStartFailure() bool {
	if o != nil && !IsNil(o.StartFailure) {
		return true
	}

	return false
}

// SetStartFailure gets a reference to the given KubevirtIoApiCoreV1VirtualMachineStartFailure and assigns it to the StartFailure field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetStartFailure(v KubevirtIoApiCoreV1VirtualMachineStartFailure) {
	o.StartFailure = &v
}

// GetStateChangeRequests returns the StateChangeRequests field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStateChangeRequests() []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest {
	if o == nil || IsNil(o.StateChangeRequests) {
		var ret []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest
		return ret
	}
	return o.StateChangeRequests
}

// GetStateChangeRequestsOk returns a tuple with the StateChangeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStateChangeRequestsOk() ([]KubevirtIoApiCoreV1VirtualMachineStateChangeRequest, bool) {
	if o == nil || IsNil(o.StateChangeRequests) {
		return nil, false
	}
	return o.StateChangeRequests, true
}

// HasStateChangeRequests returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasStateChangeRequests() bool {
	if o != nil && !IsNil(o.StateChangeRequests) {
		return true
	}

	return false
}

// SetStateChangeRequests gets a reference to the given []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest and assigns it to the StateChangeRequests field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetStateChangeRequests(v []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) {
	o.StateChangeRequests = v
}

// GetVolumeRequests returns the VolumeRequests field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeRequests() []KubevirtIoApiCoreV1VirtualMachineVolumeRequest {
	if o == nil || IsNil(o.VolumeRequests) {
		var ret []KubevirtIoApiCoreV1VirtualMachineVolumeRequest
		return ret
	}
	return o.VolumeRequests
}

// GetVolumeRequestsOk returns a tuple with the VolumeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeRequestsOk() ([]KubevirtIoApiCoreV1VirtualMachineVolumeRequest, bool) {
	if o == nil || IsNil(o.VolumeRequests) {
		return nil, false
	}
	return o.VolumeRequests, true
}

// HasVolumeRequests returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasVolumeRequests() bool {
	if o != nil && !IsNil(o.VolumeRequests) {
		return true
	}

	return false
}

// SetVolumeRequests gets a reference to the given []KubevirtIoApiCoreV1VirtualMachineVolumeRequest and assigns it to the VolumeRequests field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetVolumeRequests(v []KubevirtIoApiCoreV1VirtualMachineVolumeRequest) {
	o.VolumeRequests = v
}

// GetVolumeSnapshotStatuses returns the VolumeSnapshotStatuses field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeSnapshotStatuses() []KubevirtIoApiCoreV1VolumeSnapshotStatus {
	if o == nil || IsNil(o.VolumeSnapshotStatuses) {
		var ret []KubevirtIoApiCoreV1VolumeSnapshotStatus
		return ret
	}
	return o.VolumeSnapshotStatuses
}

// GetVolumeSnapshotStatusesOk returns a tuple with the VolumeSnapshotStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeSnapshotStatusesOk() ([]KubevirtIoApiCoreV1VolumeSnapshotStatus, bool) {
	if o == nil || IsNil(o.VolumeSnapshotStatuses) {
		return nil, false
	}
	return o.VolumeSnapshotStatuses, true
}

// HasVolumeSnapshotStatuses returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasVolumeSnapshotStatuses() bool {
	if o != nil && !IsNil(o.VolumeSnapshotStatuses) {
		return true
	}

	return false
}

// SetVolumeSnapshotStatuses gets a reference to the given []KubevirtIoApiCoreV1VolumeSnapshotStatus and assigns it to the VolumeSnapshotStatuses field.
func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetVolumeSnapshotStatuses(v []KubevirtIoApiCoreV1VolumeSnapshotStatus) {
	o.VolumeSnapshotStatuses = v
}

func (o KubevirtIoApiCoreV1VirtualMachineStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DesiredGeneration) {
		toSerialize["desiredGeneration"] = o.DesiredGeneration
	}
	if !IsNil(o.MemoryDumpRequest) {
		toSerialize["memoryDumpRequest"] = o.MemoryDumpRequest
	}
	if !IsNil(o.ObservedGeneration) {
		toSerialize["observedGeneration"] = o.ObservedGeneration
	}
	if !IsNil(o.PrintableStatus) {
		toSerialize["printableStatus"] = o.PrintableStatus
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.RestoreInProgress) {
		toSerialize["restoreInProgress"] = o.RestoreInProgress
	}
	if !IsNil(o.SnapshotInProgress) {
		toSerialize["snapshotInProgress"] = o.SnapshotInProgress
	}
	if !IsNil(o.StartFailure) {
		toSerialize["startFailure"] = o.StartFailure
	}
	if !IsNil(o.StateChangeRequests) {
		toSerialize["stateChangeRequests"] = o.StateChangeRequests
	}
	if !IsNil(o.VolumeRequests) {
		toSerialize["volumeRequests"] = o.VolumeRequests
	}
	if !IsNil(o.VolumeSnapshotStatuses) {
		toSerialize["volumeSnapshotStatuses"] = o.VolumeSnapshotStatuses
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1VirtualMachineStatus struct {
	value *KubevirtIoApiCoreV1VirtualMachineStatus
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineStatus) Get() *KubevirtIoApiCoreV1VirtualMachineStatus {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineStatus) Set(val *KubevirtIoApiCoreV1VirtualMachineStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineStatus(val *KubevirtIoApiCoreV1VirtualMachineStatus) *NullableKubevirtIoApiCoreV1VirtualMachineStatus {
	return &NullableKubevirtIoApiCoreV1VirtualMachineStatus{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


