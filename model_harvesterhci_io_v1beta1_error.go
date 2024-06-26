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

// checks if the HarvesterhciIoV1beta1Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1Error{}

// HarvesterhciIoV1beta1Error Error is the last error encountered during the snapshot/restore
type HarvesterhciIoV1beta1Error struct {
	Message *string `json:"message,omitempty"`
	// Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.
	Time *string `json:"time,omitempty"`
}

// NewHarvesterhciIoV1beta1Error instantiates a new HarvesterhciIoV1beta1Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1Error() *HarvesterhciIoV1beta1Error {
	this := HarvesterhciIoV1beta1Error{}
	var time string = ""
	this.Time = &time
	return &this
}

// NewHarvesterhciIoV1beta1ErrorWithDefaults instantiates a new HarvesterhciIoV1beta1Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1ErrorWithDefaults() *HarvesterhciIoV1beta1Error {
	this := HarvesterhciIoV1beta1Error{}
	var time string = ""
	this.Time = &time
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1Error) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1Error) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1Error) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *HarvesterhciIoV1beta1Error) SetMessage(v string) {
	o.Message = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1Error) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1Error) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1Error) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *HarvesterhciIoV1beta1Error) SetTime(v string) {
	o.Time = &v
}

func (o HarvesterhciIoV1beta1Error) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1Error) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1Error struct {
	value *HarvesterhciIoV1beta1Error
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1Error) Get() *HarvesterhciIoV1beta1Error {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1Error) Set(val *HarvesterhciIoV1beta1Error) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1Error) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1Error) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1Error(val *HarvesterhciIoV1beta1Error) *NullableHarvesterhciIoV1beta1Error {
	return &NullableHarvesterhciIoV1beta1Error{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1Error) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


