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

// checks if the KubevirtIoApiCoreV1LiveUpdateFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1LiveUpdateFeatures{}

// KubevirtIoApiCoreV1LiveUpdateFeatures struct for KubevirtIoApiCoreV1LiveUpdateFeatures
type KubevirtIoApiCoreV1LiveUpdateFeatures struct {
	// Affinity allows live updating the virtual machines node affinity
	Affinity map[string]interface{} `json:"affinity,omitempty"`
	// LiveUpdateCPU holds hotplug configuration for the CPU resource. Empty struct indicates that default will be used for maxSockets. Default is specified on cluster level. Absence of the struct means opt-out from CPU hotplug functionality.
	Cpu *KubevirtIoApiCoreV1LiveUpdateCPU `json:"cpu,omitempty"`
	// MemoryLiveUpdateConfiguration defines the live update memory features for the VirtualMachine
	Memory *KubevirtIoApiCoreV1LiveUpdateMemory `json:"memory,omitempty"`
}

// NewKubevirtIoApiCoreV1LiveUpdateFeatures instantiates a new KubevirtIoApiCoreV1LiveUpdateFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1LiveUpdateFeatures() *KubevirtIoApiCoreV1LiveUpdateFeatures {
	this := KubevirtIoApiCoreV1LiveUpdateFeatures{}
	return &this
}

// NewKubevirtIoApiCoreV1LiveUpdateFeaturesWithDefaults instantiates a new KubevirtIoApiCoreV1LiveUpdateFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1LiveUpdateFeaturesWithDefaults() *KubevirtIoApiCoreV1LiveUpdateFeatures {
	this := KubevirtIoApiCoreV1LiveUpdateFeatures{}
	return &this
}

// GetAffinity returns the Affinity field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetAffinity() map[string]interface{} {
	if o == nil || IsNil(o.Affinity) {
		var ret map[string]interface{}
		return ret
	}
	return o.Affinity
}

// GetAffinityOk returns a tuple with the Affinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetAffinityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Affinity) {
		return map[string]interface{}{}, false
	}
	return o.Affinity, true
}

// HasAffinity returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasAffinity() bool {
	if o != nil && !IsNil(o.Affinity) {
		return true
	}

	return false
}

// SetAffinity gets a reference to the given map[string]interface{} and assigns it to the Affinity field.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetAffinity(v map[string]interface{}) {
	o.Affinity = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetCpu() KubevirtIoApiCoreV1LiveUpdateCPU {
	if o == nil || IsNil(o.Cpu) {
		var ret KubevirtIoApiCoreV1LiveUpdateCPU
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetCpuOk() (*KubevirtIoApiCoreV1LiveUpdateCPU, bool) {
	if o == nil || IsNil(o.Cpu) {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasCpu() bool {
	if o != nil && !IsNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given KubevirtIoApiCoreV1LiveUpdateCPU and assigns it to the Cpu field.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetCpu(v KubevirtIoApiCoreV1LiveUpdateCPU) {
	o.Cpu = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetMemory() KubevirtIoApiCoreV1LiveUpdateMemory {
	if o == nil || IsNil(o.Memory) {
		var ret KubevirtIoApiCoreV1LiveUpdateMemory
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetMemoryOk() (*KubevirtIoApiCoreV1LiveUpdateMemory, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given KubevirtIoApiCoreV1LiveUpdateMemory and assigns it to the Memory field.
func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetMemory(v KubevirtIoApiCoreV1LiveUpdateMemory) {
	o.Memory = &v
}

func (o KubevirtIoApiCoreV1LiveUpdateFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1LiveUpdateFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Affinity) {
		toSerialize["affinity"] = o.Affinity
	}
	if !IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1LiveUpdateFeatures struct {
	value *KubevirtIoApiCoreV1LiveUpdateFeatures
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateFeatures) Get() *KubevirtIoApiCoreV1LiveUpdateFeatures {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateFeatures) Set(val *KubevirtIoApiCoreV1LiveUpdateFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1LiveUpdateFeatures(val *KubevirtIoApiCoreV1LiveUpdateFeatures) *NullableKubevirtIoApiCoreV1LiveUpdateFeatures {
	return &NullableKubevirtIoApiCoreV1LiveUpdateFeatures{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


