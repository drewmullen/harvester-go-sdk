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

// checks if the KubevirtIoApiCoreV1PITTimer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1PITTimer{}

// KubevirtIoApiCoreV1PITTimer struct for KubevirtIoApiCoreV1PITTimer
type KubevirtIoApiCoreV1PITTimer struct {
	// Enabled set to false makes sure that the machine type or a preset can't add the timer. Defaults to true.
	Present *bool `json:"present,omitempty"`
	// TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest. One of \"delay\", \"catchup\", \"discard\".
	TickPolicy *string `json:"tickPolicy,omitempty"`
}

// NewKubevirtIoApiCoreV1PITTimer instantiates a new KubevirtIoApiCoreV1PITTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1PITTimer() *KubevirtIoApiCoreV1PITTimer {
	this := KubevirtIoApiCoreV1PITTimer{}
	return &this
}

// NewKubevirtIoApiCoreV1PITTimerWithDefaults instantiates a new KubevirtIoApiCoreV1PITTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1PITTimerWithDefaults() *KubevirtIoApiCoreV1PITTimer {
	this := KubevirtIoApiCoreV1PITTimer{}
	return &this
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PITTimer) GetPresent() bool {
	if o == nil || IsNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PITTimer) GetPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Present) {
		return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PITTimer) HasPresent() bool {
	if o != nil && !IsNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *KubevirtIoApiCoreV1PITTimer) SetPresent(v bool) {
	o.Present = &v
}

// GetTickPolicy returns the TickPolicy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PITTimer) GetTickPolicy() string {
	if o == nil || IsNil(o.TickPolicy) {
		var ret string
		return ret
	}
	return *o.TickPolicy
}

// GetTickPolicyOk returns a tuple with the TickPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PITTimer) GetTickPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TickPolicy) {
		return nil, false
	}
	return o.TickPolicy, true
}

// HasTickPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PITTimer) HasTickPolicy() bool {
	if o != nil && !IsNil(o.TickPolicy) {
		return true
	}

	return false
}

// SetTickPolicy gets a reference to the given string and assigns it to the TickPolicy field.
func (o *KubevirtIoApiCoreV1PITTimer) SetTickPolicy(v string) {
	o.TickPolicy = &v
}

func (o KubevirtIoApiCoreV1PITTimer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1PITTimer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Present) {
		toSerialize["present"] = o.Present
	}
	if !IsNil(o.TickPolicy) {
		toSerialize["tickPolicy"] = o.TickPolicy
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1PITTimer struct {
	value *KubevirtIoApiCoreV1PITTimer
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1PITTimer) Get() *KubevirtIoApiCoreV1PITTimer {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1PITTimer) Set(val *KubevirtIoApiCoreV1PITTimer) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1PITTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1PITTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1PITTimer(val *KubevirtIoApiCoreV1PITTimer) *NullableKubevirtIoApiCoreV1PITTimer {
	return &NullableKubevirtIoApiCoreV1PITTimer{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1PITTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1PITTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


