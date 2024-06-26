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

// checks if the KubevirtIoApiCoreV1RTCTimer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1RTCTimer{}

// KubevirtIoApiCoreV1RTCTimer struct for KubevirtIoApiCoreV1RTCTimer
type KubevirtIoApiCoreV1RTCTimer struct {
	// Enabled set to false makes sure that the machine type or a preset can't add the timer. Defaults to true.
	Present *bool `json:"present,omitempty"`
	// TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest. One of \"delay\", \"catchup\".
	TickPolicy *string `json:"tickPolicy,omitempty"`
	// Track the guest or the wall clock.
	Track *string `json:"track,omitempty"`
}

// NewKubevirtIoApiCoreV1RTCTimer instantiates a new KubevirtIoApiCoreV1RTCTimer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1RTCTimer() *KubevirtIoApiCoreV1RTCTimer {
	this := KubevirtIoApiCoreV1RTCTimer{}
	return &this
}

// NewKubevirtIoApiCoreV1RTCTimerWithDefaults instantiates a new KubevirtIoApiCoreV1RTCTimer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1RTCTimerWithDefaults() *KubevirtIoApiCoreV1RTCTimer {
	this := KubevirtIoApiCoreV1RTCTimer{}
	return &this
}

// GetPresent returns the Present field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1RTCTimer) GetPresent() bool {
	if o == nil || IsNil(o.Present) {
		var ret bool
		return ret
	}
	return *o.Present
}

// GetPresentOk returns a tuple with the Present field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) GetPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.Present) {
		return nil, false
	}
	return o.Present, true
}

// HasPresent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) HasPresent() bool {
	if o != nil && !IsNil(o.Present) {
		return true
	}

	return false
}

// SetPresent gets a reference to the given bool and assigns it to the Present field.
func (o *KubevirtIoApiCoreV1RTCTimer) SetPresent(v bool) {
	o.Present = &v
}

// GetTickPolicy returns the TickPolicy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1RTCTimer) GetTickPolicy() string {
	if o == nil || IsNil(o.TickPolicy) {
		var ret string
		return ret
	}
	return *o.TickPolicy
}

// GetTickPolicyOk returns a tuple with the TickPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) GetTickPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TickPolicy) {
		return nil, false
	}
	return o.TickPolicy, true
}

// HasTickPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) HasTickPolicy() bool {
	if o != nil && !IsNil(o.TickPolicy) {
		return true
	}

	return false
}

// SetTickPolicy gets a reference to the given string and assigns it to the TickPolicy field.
func (o *KubevirtIoApiCoreV1RTCTimer) SetTickPolicy(v string) {
	o.TickPolicy = &v
}

// GetTrack returns the Track field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1RTCTimer) GetTrack() string {
	if o == nil || IsNil(o.Track) {
		var ret string
		return ret
	}
	return *o.Track
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) GetTrackOk() (*string, bool) {
	if o == nil || IsNil(o.Track) {
		return nil, false
	}
	return o.Track, true
}

// HasTrack returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1RTCTimer) HasTrack() bool {
	if o != nil && !IsNil(o.Track) {
		return true
	}

	return false
}

// SetTrack gets a reference to the given string and assigns it to the Track field.
func (o *KubevirtIoApiCoreV1RTCTimer) SetTrack(v string) {
	o.Track = &v
}

func (o KubevirtIoApiCoreV1RTCTimer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1RTCTimer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Present) {
		toSerialize["present"] = o.Present
	}
	if !IsNil(o.TickPolicy) {
		toSerialize["tickPolicy"] = o.TickPolicy
	}
	if !IsNil(o.Track) {
		toSerialize["track"] = o.Track
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1RTCTimer struct {
	value *KubevirtIoApiCoreV1RTCTimer
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1RTCTimer) Get() *KubevirtIoApiCoreV1RTCTimer {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1RTCTimer) Set(val *KubevirtIoApiCoreV1RTCTimer) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1RTCTimer) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1RTCTimer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1RTCTimer(val *KubevirtIoApiCoreV1RTCTimer) *NullableKubevirtIoApiCoreV1RTCTimer {
	return &NullableKubevirtIoApiCoreV1RTCTimer{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1RTCTimer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1RTCTimer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


