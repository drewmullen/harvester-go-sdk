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

// checks if the KubevirtIoApiCoreV1FeatureState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1FeatureState{}

// KubevirtIoApiCoreV1FeatureState struct for KubevirtIoApiCoreV1FeatureState
type KubevirtIoApiCoreV1FeatureState struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// NewKubevirtIoApiCoreV1FeatureState instantiates a new KubevirtIoApiCoreV1FeatureState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1FeatureState() *KubevirtIoApiCoreV1FeatureState {
	this := KubevirtIoApiCoreV1FeatureState{}
	return &this
}

// NewKubevirtIoApiCoreV1FeatureStateWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1FeatureStateWithDefaults() *KubevirtIoApiCoreV1FeatureState {
	this := KubevirtIoApiCoreV1FeatureState{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureState) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureState) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureState) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KubevirtIoApiCoreV1FeatureState) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o KubevirtIoApiCoreV1FeatureState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1FeatureState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1FeatureState struct {
	value *KubevirtIoApiCoreV1FeatureState
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1FeatureState) Get() *KubevirtIoApiCoreV1FeatureState {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1FeatureState) Set(val *KubevirtIoApiCoreV1FeatureState) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1FeatureState) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1FeatureState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1FeatureState(val *KubevirtIoApiCoreV1FeatureState) *NullableKubevirtIoApiCoreV1FeatureState {
	return &NullableKubevirtIoApiCoreV1FeatureState{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1FeatureState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1FeatureState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


