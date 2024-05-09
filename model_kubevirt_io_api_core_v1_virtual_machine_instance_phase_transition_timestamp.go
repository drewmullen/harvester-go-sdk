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

// checks if the KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp{}

// KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp VirtualMachineInstancePhaseTransitionTimestamp gives a timestamp in relation to when a phase is set on a vmi
type KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp struct {
	// Phase is the status of the VirtualMachineInstance in kubernetes world. It is not the VirtualMachineInstance status, but partially correlates to it.
	Phase *string `json:"phase,omitempty"`
	// PhaseTransitionTimestamp is the timestamp of when the phase change occurred
	PhaseTransitionTimestamp *string `json:"phaseTransitionTimestamp,omitempty"`
}

// NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp instantiates a new KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp() *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp {
	this := KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp{}
	var phaseTransitionTimestamp string = "{}"
	this.PhaseTransitionTimestamp = &phaseTransitionTimestamp
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestampWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestampWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp {
	this := KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp{}
	var phaseTransitionTimestamp string = "{}"
	this.PhaseTransitionTimestamp = &phaseTransitionTimestamp
	return &this
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhase() string {
	if o == nil || IsNil(o.Phase) {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseOk() (*string, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) SetPhase(v string) {
	o.Phase = &v
}

// GetPhaseTransitionTimestamp returns the PhaseTransitionTimestamp field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseTransitionTimestamp() string {
	if o == nil || IsNil(o.PhaseTransitionTimestamp) {
		var ret string
		return ret
	}
	return *o.PhaseTransitionTimestamp
}

// GetPhaseTransitionTimestampOk returns a tuple with the PhaseTransitionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseTransitionTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.PhaseTransitionTimestamp) {
		return nil, false
	}
	return o.PhaseTransitionTimestamp, true
}

// HasPhaseTransitionTimestamp returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) HasPhaseTransitionTimestamp() bool {
	if o != nil && !IsNil(o.PhaseTransitionTimestamp) {
		return true
	}

	return false
}

// SetPhaseTransitionTimestamp gets a reference to the given string and assigns it to the PhaseTransitionTimestamp field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) SetPhaseTransitionTimestamp(v string) {
	o.PhaseTransitionTimestamp = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.PhaseTransitionTimestamp) {
		toSerialize["phaseTransitionTimestamp"] = o.PhaseTransitionTimestamp
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) Get() *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) Set(val *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp(val *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) *NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


