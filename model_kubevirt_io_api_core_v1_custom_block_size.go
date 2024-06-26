/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoApiCoreV1CustomBlockSize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1CustomBlockSize{}

// KubevirtIoApiCoreV1CustomBlockSize CustomBlockSize represents the desired logical and physical block size for a VM disk.
type KubevirtIoApiCoreV1CustomBlockSize struct {
	Logical int32 `json:"logical"`
	Physical int32 `json:"physical"`
}

type _KubevirtIoApiCoreV1CustomBlockSize KubevirtIoApiCoreV1CustomBlockSize

// NewKubevirtIoApiCoreV1CustomBlockSize instantiates a new KubevirtIoApiCoreV1CustomBlockSize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1CustomBlockSize(logical int32, physical int32) *KubevirtIoApiCoreV1CustomBlockSize {
	this := KubevirtIoApiCoreV1CustomBlockSize{}
	this.Logical = logical
	this.Physical = physical
	return &this
}

// NewKubevirtIoApiCoreV1CustomBlockSizeWithDefaults instantiates a new KubevirtIoApiCoreV1CustomBlockSize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1CustomBlockSizeWithDefaults() *KubevirtIoApiCoreV1CustomBlockSize {
	this := KubevirtIoApiCoreV1CustomBlockSize{}
	var logical int32 = 0
	this.Logical = logical
	var physical int32 = 0
	this.Physical = physical
	return &this
}

// GetLogical returns the Logical field value
func (o *KubevirtIoApiCoreV1CustomBlockSize) GetLogical() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Logical
}

// GetLogicalOk returns a tuple with the Logical field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CustomBlockSize) GetLogicalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Logical, true
}

// SetLogical sets field value
func (o *KubevirtIoApiCoreV1CustomBlockSize) SetLogical(v int32) {
	o.Logical = v
}

// GetPhysical returns the Physical field value
func (o *KubevirtIoApiCoreV1CustomBlockSize) GetPhysical() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Physical
}

// GetPhysicalOk returns a tuple with the Physical field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CustomBlockSize) GetPhysicalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Physical, true
}

// SetPhysical sets field value
func (o *KubevirtIoApiCoreV1CustomBlockSize) SetPhysical(v int32) {
	o.Physical = v
}

func (o KubevirtIoApiCoreV1CustomBlockSize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1CustomBlockSize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["logical"] = o.Logical
	toSerialize["physical"] = o.Physical
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1CustomBlockSize) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"logical",
		"physical",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubevirtIoApiCoreV1CustomBlockSize := _KubevirtIoApiCoreV1CustomBlockSize{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1CustomBlockSize)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1CustomBlockSize(varKubevirtIoApiCoreV1CustomBlockSize)

	return err
}

type NullableKubevirtIoApiCoreV1CustomBlockSize struct {
	value *KubevirtIoApiCoreV1CustomBlockSize
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1CustomBlockSize) Get() *KubevirtIoApiCoreV1CustomBlockSize {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1CustomBlockSize) Set(val *KubevirtIoApiCoreV1CustomBlockSize) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1CustomBlockSize) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1CustomBlockSize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1CustomBlockSize(val *KubevirtIoApiCoreV1CustomBlockSize) *NullableKubevirtIoApiCoreV1CustomBlockSize {
	return &NullableKubevirtIoApiCoreV1CustomBlockSize{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1CustomBlockSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1CustomBlockSize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


