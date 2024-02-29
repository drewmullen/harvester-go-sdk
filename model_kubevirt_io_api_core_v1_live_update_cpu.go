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

// checks if the KubevirtIoApiCoreV1LiveUpdateCPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1LiveUpdateCPU{}

// KubevirtIoApiCoreV1LiveUpdateCPU struct for KubevirtIoApiCoreV1LiveUpdateCPU
type KubevirtIoApiCoreV1LiveUpdateCPU struct {
	MaxSockets *int64 `json:"maxSockets,omitempty"`
}

// NewKubevirtIoApiCoreV1LiveUpdateCPU instantiates a new KubevirtIoApiCoreV1LiveUpdateCPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1LiveUpdateCPU() *KubevirtIoApiCoreV1LiveUpdateCPU {
	this := KubevirtIoApiCoreV1LiveUpdateCPU{}
	return &this
}

// NewKubevirtIoApiCoreV1LiveUpdateCPUWithDefaults instantiates a new KubevirtIoApiCoreV1LiveUpdateCPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1LiveUpdateCPUWithDefaults() *KubevirtIoApiCoreV1LiveUpdateCPU {
	this := KubevirtIoApiCoreV1LiveUpdateCPU{}
	return &this
}

// GetMaxSockets returns the MaxSockets field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1LiveUpdateCPU) GetMaxSockets() int64 {
	if o == nil || IsNil(o.MaxSockets) {
		var ret int64
		return ret
	}
	return *o.MaxSockets
}

// GetMaxSocketsOk returns a tuple with the MaxSockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateCPU) GetMaxSocketsOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxSockets) {
		return nil, false
	}
	return o.MaxSockets, true
}

// HasMaxSockets returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateCPU) HasMaxSockets() bool {
	if o != nil && !IsNil(o.MaxSockets) {
		return true
	}

	return false
}

// SetMaxSockets gets a reference to the given int64 and assigns it to the MaxSockets field.
func (o *KubevirtIoApiCoreV1LiveUpdateCPU) SetMaxSockets(v int64) {
	o.MaxSockets = &v
}

func (o KubevirtIoApiCoreV1LiveUpdateCPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1LiveUpdateCPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxSockets) {
		toSerialize["maxSockets"] = o.MaxSockets
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1LiveUpdateCPU struct {
	value *KubevirtIoApiCoreV1LiveUpdateCPU
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateCPU) Get() *KubevirtIoApiCoreV1LiveUpdateCPU {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateCPU) Set(val *KubevirtIoApiCoreV1LiveUpdateCPU) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateCPU) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateCPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1LiveUpdateCPU(val *KubevirtIoApiCoreV1LiveUpdateCPU) *NullableKubevirtIoApiCoreV1LiveUpdateCPU {
	return &NullableKubevirtIoApiCoreV1LiveUpdateCPU{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateCPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateCPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


