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

// checks if the KubevirtIoApiCoreV1Clock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1Clock{}

// KubevirtIoApiCoreV1Clock struct for KubevirtIoApiCoreV1Clock
type KubevirtIoApiCoreV1Clock struct {
	Timer *KubevirtIoApiCoreV1Timer `json:"timer,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Utc *KubevirtIoApiCoreV1ClockOffsetUTC `json:"utc,omitempty"`
}

// NewKubevirtIoApiCoreV1Clock instantiates a new KubevirtIoApiCoreV1Clock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1Clock() *KubevirtIoApiCoreV1Clock {
	this := KubevirtIoApiCoreV1Clock{}
	return &this
}

// NewKubevirtIoApiCoreV1ClockWithDefaults instantiates a new KubevirtIoApiCoreV1Clock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1ClockWithDefaults() *KubevirtIoApiCoreV1Clock {
	this := KubevirtIoApiCoreV1Clock{}
	return &this
}

// GetTimer returns the Timer field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Clock) GetTimer() KubevirtIoApiCoreV1Timer {
	if o == nil || IsNil(o.Timer) {
		var ret KubevirtIoApiCoreV1Timer
		return ret
	}
	return *o.Timer
}

// GetTimerOk returns a tuple with the Timer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Clock) GetTimerOk() (*KubevirtIoApiCoreV1Timer, bool) {
	if o == nil || IsNil(o.Timer) {
		return nil, false
	}
	return o.Timer, true
}

// HasTimer returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Clock) HasTimer() bool {
	if o != nil && !IsNil(o.Timer) {
		return true
	}

	return false
}

// SetTimer gets a reference to the given KubevirtIoApiCoreV1Timer and assigns it to the Timer field.
func (o *KubevirtIoApiCoreV1Clock) SetTimer(v KubevirtIoApiCoreV1Timer) {
	o.Timer = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Clock) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Clock) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Clock) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *KubevirtIoApiCoreV1Clock) SetTimezone(v string) {
	o.Timezone = &v
}

// GetUtc returns the Utc field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Clock) GetUtc() KubevirtIoApiCoreV1ClockOffsetUTC {
	if o == nil || IsNil(o.Utc) {
		var ret KubevirtIoApiCoreV1ClockOffsetUTC
		return ret
	}
	return *o.Utc
}

// GetUtcOk returns a tuple with the Utc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Clock) GetUtcOk() (*KubevirtIoApiCoreV1ClockOffsetUTC, bool) {
	if o == nil || IsNil(o.Utc) {
		return nil, false
	}
	return o.Utc, true
}

// HasUtc returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Clock) HasUtc() bool {
	if o != nil && !IsNil(o.Utc) {
		return true
	}

	return false
}

// SetUtc gets a reference to the given KubevirtIoApiCoreV1ClockOffsetUTC and assigns it to the Utc field.
func (o *KubevirtIoApiCoreV1Clock) SetUtc(v KubevirtIoApiCoreV1ClockOffsetUTC) {
	o.Utc = &v
}

func (o KubevirtIoApiCoreV1Clock) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1Clock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timer) {
		toSerialize["timer"] = o.Timer
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Utc) {
		toSerialize["utc"] = o.Utc
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1Clock struct {
	value *KubevirtIoApiCoreV1Clock
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1Clock) Get() *KubevirtIoApiCoreV1Clock {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1Clock) Set(val *KubevirtIoApiCoreV1Clock) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1Clock) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1Clock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1Clock(val *KubevirtIoApiCoreV1Clock) *NullableKubevirtIoApiCoreV1Clock {
	return &NullableKubevirtIoApiCoreV1Clock{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1Clock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1Clock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


