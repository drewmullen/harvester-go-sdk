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

// checks if the KubevirtIoApiCoreV1LiveUpdateMemory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1LiveUpdateMemory{}

// KubevirtIoApiCoreV1LiveUpdateMemory struct for KubevirtIoApiCoreV1LiveUpdateMemory
type KubevirtIoApiCoreV1LiveUpdateMemory struct {
	MaxGuest *string `json:"maxGuest,omitempty"`
}

// NewKubevirtIoApiCoreV1LiveUpdateMemory instantiates a new KubevirtIoApiCoreV1LiveUpdateMemory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1LiveUpdateMemory() *KubevirtIoApiCoreV1LiveUpdateMemory {
	this := KubevirtIoApiCoreV1LiveUpdateMemory{}
	return &this
}

// NewKubevirtIoApiCoreV1LiveUpdateMemoryWithDefaults instantiates a new KubevirtIoApiCoreV1LiveUpdateMemory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1LiveUpdateMemoryWithDefaults() *KubevirtIoApiCoreV1LiveUpdateMemory {
	this := KubevirtIoApiCoreV1LiveUpdateMemory{}
	return &this
}

// GetMaxGuest returns the MaxGuest field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1LiveUpdateMemory) GetMaxGuest() string {
	if o == nil || IsNil(o.MaxGuest) {
		var ret string
		return ret
	}
	return *o.MaxGuest
}

// GetMaxGuestOk returns a tuple with the MaxGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateMemory) GetMaxGuestOk() (*string, bool) {
	if o == nil || IsNil(o.MaxGuest) {
		return nil, false
	}
	return o.MaxGuest, true
}

// HasMaxGuest returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1LiveUpdateMemory) HasMaxGuest() bool {
	if o != nil && !IsNil(o.MaxGuest) {
		return true
	}

	return false
}

// SetMaxGuest gets a reference to the given string and assigns it to the MaxGuest field.
func (o *KubevirtIoApiCoreV1LiveUpdateMemory) SetMaxGuest(v string) {
	o.MaxGuest = &v
}

func (o KubevirtIoApiCoreV1LiveUpdateMemory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1LiveUpdateMemory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxGuest) {
		toSerialize["maxGuest"] = o.MaxGuest
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1LiveUpdateMemory struct {
	value *KubevirtIoApiCoreV1LiveUpdateMemory
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateMemory) Get() *KubevirtIoApiCoreV1LiveUpdateMemory {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateMemory) Set(val *KubevirtIoApiCoreV1LiveUpdateMemory) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateMemory) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateMemory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1LiveUpdateMemory(val *KubevirtIoApiCoreV1LiveUpdateMemory) *NullableKubevirtIoApiCoreV1LiveUpdateMemory {
	return &NullableKubevirtIoApiCoreV1LiveUpdateMemory{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1LiveUpdateMemory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1LiveUpdateMemory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


