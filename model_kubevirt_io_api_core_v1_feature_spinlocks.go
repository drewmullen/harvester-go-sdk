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

// checks if the KubevirtIoApiCoreV1FeatureSpinlocks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1FeatureSpinlocks{}

// KubevirtIoApiCoreV1FeatureSpinlocks struct for KubevirtIoApiCoreV1FeatureSpinlocks
type KubevirtIoApiCoreV1FeatureSpinlocks struct {
	Enabled *bool `json:"enabled,omitempty"`
	Spinlocks *int64 `json:"spinlocks,omitempty"`
}

// NewKubevirtIoApiCoreV1FeatureSpinlocks instantiates a new KubevirtIoApiCoreV1FeatureSpinlocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1FeatureSpinlocks() *KubevirtIoApiCoreV1FeatureSpinlocks {
	this := KubevirtIoApiCoreV1FeatureSpinlocks{}
	return &this
}

// NewKubevirtIoApiCoreV1FeatureSpinlocksWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureSpinlocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1FeatureSpinlocksWithDefaults() *KubevirtIoApiCoreV1FeatureSpinlocks {
	this := KubevirtIoApiCoreV1FeatureSpinlocks{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSpinlocks returns the Spinlocks field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetSpinlocks() int64 {
	if o == nil || IsNil(o.Spinlocks) {
		var ret int64
		return ret
	}
	return *o.Spinlocks
}

// GetSpinlocksOk returns a tuple with the Spinlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetSpinlocksOk() (*int64, bool) {
	if o == nil || IsNil(o.Spinlocks) {
		return nil, false
	}
	return o.Spinlocks, true
}

// HasSpinlocks returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) HasSpinlocks() bool {
	if o != nil && !IsNil(o.Spinlocks) {
		return true
	}

	return false
}

// SetSpinlocks gets a reference to the given int64 and assigns it to the Spinlocks field.
func (o *KubevirtIoApiCoreV1FeatureSpinlocks) SetSpinlocks(v int64) {
	o.Spinlocks = &v
}

func (o KubevirtIoApiCoreV1FeatureSpinlocks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1FeatureSpinlocks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Spinlocks) {
		toSerialize["spinlocks"] = o.Spinlocks
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1FeatureSpinlocks struct {
	value *KubevirtIoApiCoreV1FeatureSpinlocks
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1FeatureSpinlocks) Get() *KubevirtIoApiCoreV1FeatureSpinlocks {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1FeatureSpinlocks) Set(val *KubevirtIoApiCoreV1FeatureSpinlocks) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1FeatureSpinlocks) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1FeatureSpinlocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1FeatureSpinlocks(val *KubevirtIoApiCoreV1FeatureSpinlocks) *NullableKubevirtIoApiCoreV1FeatureSpinlocks {
	return &NullableKubevirtIoApiCoreV1FeatureSpinlocks{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1FeatureSpinlocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1FeatureSpinlocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


