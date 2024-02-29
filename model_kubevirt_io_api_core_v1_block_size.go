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

// checks if the KubevirtIoApiCoreV1BlockSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1BlockSize{}

// KubevirtIoApiCoreV1BlockSize struct for KubevirtIoApiCoreV1BlockSize
type KubevirtIoApiCoreV1BlockSize struct {
	Custom *KubevirtIoApiCoreV1CustomBlockSize `json:"custom,omitempty"`
	MatchVolume *KubevirtIoApiCoreV1FeatureState `json:"matchVolume,omitempty"`
}

// NewKubevirtIoApiCoreV1BlockSize instantiates a new KubevirtIoApiCoreV1BlockSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1BlockSize() *KubevirtIoApiCoreV1BlockSize {
	this := KubevirtIoApiCoreV1BlockSize{}
	return &this
}

// NewKubevirtIoApiCoreV1BlockSizeWithDefaults instantiates a new KubevirtIoApiCoreV1BlockSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1BlockSizeWithDefaults() *KubevirtIoApiCoreV1BlockSize {
	this := KubevirtIoApiCoreV1BlockSize{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1BlockSize) GetCustom() KubevirtIoApiCoreV1CustomBlockSize {
	if o == nil || IsNil(o.Custom) {
		var ret KubevirtIoApiCoreV1CustomBlockSize
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1BlockSize) GetCustomOk() (*KubevirtIoApiCoreV1CustomBlockSize, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1BlockSize) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given KubevirtIoApiCoreV1CustomBlockSize and assigns it to the Custom field.
func (o *KubevirtIoApiCoreV1BlockSize) SetCustom(v KubevirtIoApiCoreV1CustomBlockSize) {
	o.Custom = &v
}

// GetMatchVolume returns the MatchVolume field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1BlockSize) GetMatchVolume() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.MatchVolume) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.MatchVolume
}

// GetMatchVolumeOk returns a tuple with the MatchVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1BlockSize) GetMatchVolumeOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.MatchVolume) {
		return nil, false
	}
	return o.MatchVolume, true
}

// HasMatchVolume returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1BlockSize) HasMatchVolume() bool {
	if o != nil && !IsNil(o.MatchVolume) {
		return true
	}

	return false
}

// SetMatchVolume gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the MatchVolume field.
func (o *KubevirtIoApiCoreV1BlockSize) SetMatchVolume(v KubevirtIoApiCoreV1FeatureState) {
	o.MatchVolume = &v
}

func (o KubevirtIoApiCoreV1BlockSize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1BlockSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.MatchVolume) {
		toSerialize["matchVolume"] = o.MatchVolume
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1BlockSize struct {
	value *KubevirtIoApiCoreV1BlockSize
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1BlockSize) Get() *KubevirtIoApiCoreV1BlockSize {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1BlockSize) Set(val *KubevirtIoApiCoreV1BlockSize) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1BlockSize) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1BlockSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1BlockSize(val *KubevirtIoApiCoreV1BlockSize) *NullableKubevirtIoApiCoreV1BlockSize {
	return &NullableKubevirtIoApiCoreV1BlockSize{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1BlockSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1BlockSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


