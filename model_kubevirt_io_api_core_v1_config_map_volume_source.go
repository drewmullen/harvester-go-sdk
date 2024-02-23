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

// checks if the KubevirtIoApiCoreV1ConfigMapVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1ConfigMapVolumeSource{}

// KubevirtIoApiCoreV1ConfigMapVolumeSource struct for KubevirtIoApiCoreV1ConfigMapVolumeSource
type KubevirtIoApiCoreV1ConfigMapVolumeSource struct {
	Name *string `json:"name,omitempty"`
	Optional *bool `json:"optional,omitempty"`
	VolumeLabel *string `json:"volumeLabel,omitempty"`
}

// NewKubevirtIoApiCoreV1ConfigMapVolumeSource instantiates a new KubevirtIoApiCoreV1ConfigMapVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1ConfigMapVolumeSource() *KubevirtIoApiCoreV1ConfigMapVolumeSource {
	this := KubevirtIoApiCoreV1ConfigMapVolumeSource{}
	return &this
}

// NewKubevirtIoApiCoreV1ConfigMapVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1ConfigMapVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1ConfigMapVolumeSourceWithDefaults() *KubevirtIoApiCoreV1ConfigMapVolumeSource {
	this := KubevirtIoApiCoreV1ConfigMapVolumeSource{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetOptional() bool {
	if o == nil || IsNil(o.Optional) {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetOptionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Optional) {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasOptional() bool {
	if o != nil && !IsNil(o.Optional) {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetOptional(v bool) {
	o.Optional = &v
}

// GetVolumeLabel returns the VolumeLabel field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetVolumeLabel() string {
	if o == nil || IsNil(o.VolumeLabel) {
		var ret string
		return ret
	}
	return *o.VolumeLabel
}

// GetVolumeLabelOk returns a tuple with the VolumeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetVolumeLabelOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeLabel) {
		return nil, false
	}
	return o.VolumeLabel, true
}

// HasVolumeLabel returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasVolumeLabel() bool {
	if o != nil && !IsNil(o.VolumeLabel) {
		return true
	}

	return false
}

// SetVolumeLabel gets a reference to the given string and assigns it to the VolumeLabel field.
func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetVolumeLabel(v string) {
	o.VolumeLabel = &v
}

func (o KubevirtIoApiCoreV1ConfigMapVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1ConfigMapVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Optional) {
		toSerialize["optional"] = o.Optional
	}
	if !IsNil(o.VolumeLabel) {
		toSerialize["volumeLabel"] = o.VolumeLabel
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1ConfigMapVolumeSource struct {
	value *KubevirtIoApiCoreV1ConfigMapVolumeSource
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) Get() *KubevirtIoApiCoreV1ConfigMapVolumeSource {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) Set(val *KubevirtIoApiCoreV1ConfigMapVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1ConfigMapVolumeSource(val *KubevirtIoApiCoreV1ConfigMapVolumeSource) *NullableKubevirtIoApiCoreV1ConfigMapVolumeSource {
	return &NullableKubevirtIoApiCoreV1ConfigMapVolumeSource{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1ConfigMapVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


