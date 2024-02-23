/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HarvesterhciIoV1beta1KeyPairSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1KeyPairSpec{}

// HarvesterhciIoV1beta1KeyPairSpec struct for HarvesterhciIoV1beta1KeyPairSpec
type HarvesterhciIoV1beta1KeyPairSpec struct {
	PublicKey string `json:"publicKey"`
}

type _HarvesterhciIoV1beta1KeyPairSpec HarvesterhciIoV1beta1KeyPairSpec

// NewHarvesterhciIoV1beta1KeyPairSpec instantiates a new HarvesterhciIoV1beta1KeyPairSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1KeyPairSpec(publicKey string) *HarvesterhciIoV1beta1KeyPairSpec {
	this := HarvesterhciIoV1beta1KeyPairSpec{}
	this.PublicKey = publicKey
	return &this
}

// NewHarvesterhciIoV1beta1KeyPairSpecWithDefaults instantiates a new HarvesterhciIoV1beta1KeyPairSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1KeyPairSpecWithDefaults() *HarvesterhciIoV1beta1KeyPairSpec {
	this := HarvesterhciIoV1beta1KeyPairSpec{}
	var publicKey string = ""
	this.PublicKey = publicKey
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *HarvesterhciIoV1beta1KeyPairSpec) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1KeyPairSpec) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *HarvesterhciIoV1beta1KeyPairSpec) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o HarvesterhciIoV1beta1KeyPairSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1KeyPairSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["publicKey"] = o.PublicKey
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1KeyPairSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"publicKey",
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

	varHarvesterhciIoV1beta1KeyPairSpec := _HarvesterhciIoV1beta1KeyPairSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1KeyPairSpec)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1KeyPairSpec(varHarvesterhciIoV1beta1KeyPairSpec)

	return err
}

type NullableHarvesterhciIoV1beta1KeyPairSpec struct {
	value *HarvesterhciIoV1beta1KeyPairSpec
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1KeyPairSpec) Get() *HarvesterhciIoV1beta1KeyPairSpec {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1KeyPairSpec) Set(val *HarvesterhciIoV1beta1KeyPairSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1KeyPairSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1KeyPairSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1KeyPairSpec(val *HarvesterhciIoV1beta1KeyPairSpec) *NullableHarvesterhciIoV1beta1KeyPairSpec {
	return &NullableHarvesterhciIoV1beta1KeyPairSpec{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1KeyPairSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1KeyPairSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


