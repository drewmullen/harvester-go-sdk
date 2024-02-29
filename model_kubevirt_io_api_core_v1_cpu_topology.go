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

// checks if the KubevirtIoApiCoreV1CPUTopology type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1CPUTopology{}

// KubevirtIoApiCoreV1CPUTopology struct for KubevirtIoApiCoreV1CPUTopology
type KubevirtIoApiCoreV1CPUTopology struct {
	Cores *int64 `json:"cores,omitempty"`
	Sockets *int64 `json:"sockets,omitempty"`
	Threads *int64 `json:"threads,omitempty"`
}

// NewKubevirtIoApiCoreV1CPUTopology instantiates a new KubevirtIoApiCoreV1CPUTopology object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1CPUTopology() *KubevirtIoApiCoreV1CPUTopology {
	this := KubevirtIoApiCoreV1CPUTopology{}
	return &this
}

// NewKubevirtIoApiCoreV1CPUTopologyWithDefaults instantiates a new KubevirtIoApiCoreV1CPUTopology object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1CPUTopologyWithDefaults() *KubevirtIoApiCoreV1CPUTopology {
	this := KubevirtIoApiCoreV1CPUTopology{}
	return &this
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CPUTopology) GetCores() int64 {
	if o == nil || IsNil(o.Cores) {
		var ret int64
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) GetCoresOk() (*int64, bool) {
	if o == nil || IsNil(o.Cores) {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given int64 and assigns it to the Cores field.
func (o *KubevirtIoApiCoreV1CPUTopology) SetCores(v int64) {
	o.Cores = &v
}

// GetSockets returns the Sockets field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CPUTopology) GetSockets() int64 {
	if o == nil || IsNil(o.Sockets) {
		var ret int64
		return ret
	}
	return *o.Sockets
}

// GetSocketsOk returns a tuple with the Sockets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) GetSocketsOk() (*int64, bool) {
	if o == nil || IsNil(o.Sockets) {
		return nil, false
	}
	return o.Sockets, true
}

// HasSockets returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) HasSockets() bool {
	if o != nil && !IsNil(o.Sockets) {
		return true
	}

	return false
}

// SetSockets gets a reference to the given int64 and assigns it to the Sockets field.
func (o *KubevirtIoApiCoreV1CPUTopology) SetSockets(v int64) {
	o.Sockets = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CPUTopology) GetThreads() int64 {
	if o == nil || IsNil(o.Threads) {
		var ret int64
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) GetThreadsOk() (*int64, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CPUTopology) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given int64 and assigns it to the Threads field.
func (o *KubevirtIoApiCoreV1CPUTopology) SetThreads(v int64) {
	o.Threads = &v
}

func (o KubevirtIoApiCoreV1CPUTopology) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1CPUTopology) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsNil(o.Sockets) {
		toSerialize["sockets"] = o.Sockets
	}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1CPUTopology struct {
	value *KubevirtIoApiCoreV1CPUTopology
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1CPUTopology) Get() *KubevirtIoApiCoreV1CPUTopology {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1CPUTopology) Set(val *KubevirtIoApiCoreV1CPUTopology) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1CPUTopology) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1CPUTopology) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1CPUTopology(val *KubevirtIoApiCoreV1CPUTopology) *NullableKubevirtIoApiCoreV1CPUTopology {
	return &NullableKubevirtIoApiCoreV1CPUTopology{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1CPUTopology) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1CPUTopology) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


