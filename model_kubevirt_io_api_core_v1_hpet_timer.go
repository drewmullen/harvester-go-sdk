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

// checks if the KubevirtIoApiCoreV1HPETTimer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1HPETTimer{}

// KubevirtIoApiCoreV1HPETTimer struct for KubevirtIoApiCoreV1HPETTimer
type KubevirtIoApiCoreV1HPETTimer struct {
	Present *bool `json:"present,omitempty"`
	TickPolicy *string `json:"tickPolicy,omitempty"`
}

// NewKubevirtIoApiCoreV1HPETTimer instantiates a new KubevirtIoApiCoreV1HPETTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1HPETTimer() *KubevirtIoApiCoreV1HPETTimer {
	this := KubevirtIoApiCoreV1HPETTimer{}
	return &this
}

// NewKubevirtIoApiCoreV1HPETTimerWithDefaults instantiates a new KubevirtIoApiCoreV1HPETTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1HPETTimerWithDefaults() *KubevirtIoApiCoreV1HPETTimer {
	this := KubevirtIoApiCoreV1HPETTimer{}
	return &this
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1HPETTimer) GetPresent() bool {
	if o == nil || IsNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HPETTimer) GetPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Present) {
		return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1HPETTimer) HasPresent() bool {
	if o != nil && !IsNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *KubevirtIoApiCoreV1HPETTimer) SetPresent(v bool) {
	o.Present = &v
}

// GetTickPolicy returns the TickPolicy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1HPETTimer) GetTickPolicy() string {
	if o == nil || IsNil(o.TickPolicy) {
		var ret string
		return ret
	}
	return *o.TickPolicy
}

// GetTickPolicyOk returns a tuple with the TickPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HPETTimer) GetTickPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TickPolicy) {
		return nil, false
	}
	return o.TickPolicy, true
}

// HasTickPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1HPETTimer) HasTickPolicy() bool {
	if o != nil && !IsNil(o.TickPolicy) {
		return true
	}

	return false
}

// SetTickPolicy gets a reference to the given string and assigns it to the TickPolicy field.
func (o *KubevirtIoApiCoreV1HPETTimer) SetTickPolicy(v string) {
	o.TickPolicy = &v
}

func (o KubevirtIoApiCoreV1HPETTimer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1HPETTimer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Present) {
		toSerialize["present"] = o.Present
	}
	if !IsNil(o.TickPolicy) {
		toSerialize["tickPolicy"] = o.TickPolicy
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1HPETTimer struct {
	value *KubevirtIoApiCoreV1HPETTimer
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1HPETTimer) Get() *KubevirtIoApiCoreV1HPETTimer {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1HPETTimer) Set(val *KubevirtIoApiCoreV1HPETTimer) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1HPETTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1HPETTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1HPETTimer(val *KubevirtIoApiCoreV1HPETTimer) *NullableKubevirtIoApiCoreV1HPETTimer {
	return &NullableKubevirtIoApiCoreV1HPETTimer{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1HPETTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1HPETTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


