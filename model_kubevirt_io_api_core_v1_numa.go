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

// checks if the KubevirtIoApiCoreV1NUMA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1NUMA{}

// KubevirtIoApiCoreV1NUMA struct for KubevirtIoApiCoreV1NUMA
type KubevirtIoApiCoreV1NUMA struct {
	GuestMappingPassthrough map[string]interface{} `json:"guestMappingPassthrough,omitempty"`
}

// NewKubevirtIoApiCoreV1NUMA instantiates a new KubevirtIoApiCoreV1NUMA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1NUMA() *KubevirtIoApiCoreV1NUMA {
	this := KubevirtIoApiCoreV1NUMA{}
	return &this
}

// NewKubevirtIoApiCoreV1NUMAWithDefaults instantiates a new KubevirtIoApiCoreV1NUMA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1NUMAWithDefaults() *KubevirtIoApiCoreV1NUMA {
	this := KubevirtIoApiCoreV1NUMA{}
	return &this
}

// GetGuestMappingPassthrough returns the GuestMappingPassthrough field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1NUMA) GetGuestMappingPassthrough() map[string]interface{} {
	if o == nil || IsNil(o.GuestMappingPassthrough) {
		var ret map[string]interface{}
		return ret
	}
	return o.GuestMappingPassthrough
}

// GetGuestMappingPassthroughOk returns a tuple with the GuestMappingPassthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1NUMA) GetGuestMappingPassthroughOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.GuestMappingPassthrough) {
		return map[string]interface{}{}, false
	}
	return o.GuestMappingPassthrough, true
}

// HasGuestMappingPassthrough returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1NUMA) HasGuestMappingPassthrough() bool {
	if o != nil && !IsNil(o.GuestMappingPassthrough) {
		return true
	}

	return false
}

// SetGuestMappingPassthrough gets a reference to the given map[string]interface{} and assigns it to the GuestMappingPassthrough field.
func (o *KubevirtIoApiCoreV1NUMA) SetGuestMappingPassthrough(v map[string]interface{}) {
	o.GuestMappingPassthrough = v
}

func (o KubevirtIoApiCoreV1NUMA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1NUMA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestMappingPassthrough) {
		toSerialize["guestMappingPassthrough"] = o.GuestMappingPassthrough
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1NUMA struct {
	value *KubevirtIoApiCoreV1NUMA
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1NUMA) Get() *KubevirtIoApiCoreV1NUMA {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1NUMA) Set(val *KubevirtIoApiCoreV1NUMA) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1NUMA) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1NUMA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1NUMA(val *KubevirtIoApiCoreV1NUMA) *NullableKubevirtIoApiCoreV1NUMA {
	return &NullableKubevirtIoApiCoreV1NUMA{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1NUMA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1NUMA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


