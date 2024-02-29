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

// checks if the KubevirtIoApiCoreV1DownwardAPIVolumeSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1DownwardAPIVolumeSource{}

// KubevirtIoApiCoreV1DownwardAPIVolumeSource struct for KubevirtIoApiCoreV1DownwardAPIVolumeSource
type KubevirtIoApiCoreV1DownwardAPIVolumeSource struct {
	Fields []K8sIoV1DownwardAPIVolumeFile `json:"fields,omitempty"`
	VolumeLabel *string `json:"volumeLabel,omitempty"`
}

// NewKubevirtIoApiCoreV1DownwardAPIVolumeSource instantiates a new KubevirtIoApiCoreV1DownwardAPIVolumeSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1DownwardAPIVolumeSource() *KubevirtIoApiCoreV1DownwardAPIVolumeSource {
	this := KubevirtIoApiCoreV1DownwardAPIVolumeSource{}
	return &this
}

// NewKubevirtIoApiCoreV1DownwardAPIVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1DownwardAPIVolumeSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1DownwardAPIVolumeSourceWithDefaults() *KubevirtIoApiCoreV1DownwardAPIVolumeSource {
	this := KubevirtIoApiCoreV1DownwardAPIVolumeSource{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetFields() []K8sIoV1DownwardAPIVolumeFile {
	if o == nil || IsNil(o.Fields) {
		var ret []K8sIoV1DownwardAPIVolumeFile
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetFieldsOk() ([]K8sIoV1DownwardAPIVolumeFile, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []K8sIoV1DownwardAPIVolumeFile and assigns it to the Fields field.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) SetFields(v []K8sIoV1DownwardAPIVolumeFile) {
	o.Fields = v
}

// GetVolumeLabel returns the VolumeLabel field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetVolumeLabel() string {
	if o == nil || IsNil(o.VolumeLabel) {
		var ret string
		return ret
	}
	return *o.VolumeLabel
}

// GetVolumeLabelOk returns a tuple with the VolumeLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetVolumeLabelOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeLabel) {
		return nil, false
	}
	return o.VolumeLabel, true
}

// HasVolumeLabel returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) HasVolumeLabel() bool {
	if o != nil && !IsNil(o.VolumeLabel) {
		return true
	}

	return false
}

// SetVolumeLabel gets a reference to the given string and assigns it to the VolumeLabel field.
func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) SetVolumeLabel(v string) {
	o.VolumeLabel = &v
}

func (o KubevirtIoApiCoreV1DownwardAPIVolumeSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1DownwardAPIVolumeSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.VolumeLabel) {
		toSerialize["volumeLabel"] = o.VolumeLabel
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource struct {
	value *KubevirtIoApiCoreV1DownwardAPIVolumeSource
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) Get() *KubevirtIoApiCoreV1DownwardAPIVolumeSource {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) Set(val *KubevirtIoApiCoreV1DownwardAPIVolumeSource) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1DownwardAPIVolumeSource(val *KubevirtIoApiCoreV1DownwardAPIVolumeSource) *NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource {
	return &NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1DownwardAPIVolumeSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


