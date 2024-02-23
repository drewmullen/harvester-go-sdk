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

// checks if the K8sIoV1StatusCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1StatusCause{}

// K8sIoV1StatusCause struct for K8sIoV1StatusCause
type K8sIoV1StatusCause struct {
	Field *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// NewK8sIoV1StatusCause instantiates a new K8sIoV1StatusCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1StatusCause() *K8sIoV1StatusCause {
	this := K8sIoV1StatusCause{}
	return &this
}

// NewK8sIoV1StatusCauseWithDefaults instantiates a new K8sIoV1StatusCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1StatusCauseWithDefaults() *K8sIoV1StatusCause {
	this := K8sIoV1StatusCause{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *K8sIoV1StatusCause) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1StatusCause) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *K8sIoV1StatusCause) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *K8sIoV1StatusCause) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *K8sIoV1StatusCause) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1StatusCause) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *K8sIoV1StatusCause) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *K8sIoV1StatusCause) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *K8sIoV1StatusCause) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1StatusCause) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *K8sIoV1StatusCause) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *K8sIoV1StatusCause) SetReason(v string) {
	o.Reason = &v
}

func (o K8sIoV1StatusCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1StatusCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableK8sIoV1StatusCause struct {
	value *K8sIoV1StatusCause
	isSet bool
}

func (v NullableK8sIoV1StatusCause) Get() *K8sIoV1StatusCause {
	return v.value
}

func (v *NullableK8sIoV1StatusCause) Set(val *K8sIoV1StatusCause) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1StatusCause) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1StatusCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1StatusCause(val *K8sIoV1StatusCause) *NullableK8sIoV1StatusCause {
	return &NullableK8sIoV1StatusCause{value: val, isSet: true}
}

func (v NullableK8sIoV1StatusCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1StatusCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


