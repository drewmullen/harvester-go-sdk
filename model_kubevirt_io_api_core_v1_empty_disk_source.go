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

// checks if the KubevirtIoApiCoreV1EmptyDiskSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1EmptyDiskSource{}

// KubevirtIoApiCoreV1EmptyDiskSource struct for KubevirtIoApiCoreV1EmptyDiskSource
type KubevirtIoApiCoreV1EmptyDiskSource struct {
	Capacity string `json:"capacity"`
}

type _KubevirtIoApiCoreV1EmptyDiskSource KubevirtIoApiCoreV1EmptyDiskSource

// NewKubevirtIoApiCoreV1EmptyDiskSource instantiates a new KubevirtIoApiCoreV1EmptyDiskSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1EmptyDiskSource(capacity string) *KubevirtIoApiCoreV1EmptyDiskSource {
	this := KubevirtIoApiCoreV1EmptyDiskSource{}
	this.Capacity = capacity
	return &this
}

// NewKubevirtIoApiCoreV1EmptyDiskSourceWithDefaults instantiates a new KubevirtIoApiCoreV1EmptyDiskSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1EmptyDiskSourceWithDefaults() *KubevirtIoApiCoreV1EmptyDiskSource {
	this := KubevirtIoApiCoreV1EmptyDiskSource{}
	return &this
}

// GetCapacity returns the Capacity field value
func (o *KubevirtIoApiCoreV1EmptyDiskSource) GetCapacity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1EmptyDiskSource) GetCapacityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *KubevirtIoApiCoreV1EmptyDiskSource) SetCapacity(v string) {
	o.Capacity = v
}

func (o KubevirtIoApiCoreV1EmptyDiskSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1EmptyDiskSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["capacity"] = o.Capacity
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1EmptyDiskSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"capacity",
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

	varKubevirtIoApiCoreV1EmptyDiskSource := _KubevirtIoApiCoreV1EmptyDiskSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1EmptyDiskSource)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1EmptyDiskSource(varKubevirtIoApiCoreV1EmptyDiskSource)

	return err
}

type NullableKubevirtIoApiCoreV1EmptyDiskSource struct {
	value *KubevirtIoApiCoreV1EmptyDiskSource
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1EmptyDiskSource) Get() *KubevirtIoApiCoreV1EmptyDiskSource {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1EmptyDiskSource) Set(val *KubevirtIoApiCoreV1EmptyDiskSource) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1EmptyDiskSource) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1EmptyDiskSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1EmptyDiskSource(val *KubevirtIoApiCoreV1EmptyDiskSource) *NullableKubevirtIoApiCoreV1EmptyDiskSource {
	return &NullableKubevirtIoApiCoreV1EmptyDiskSource{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1EmptyDiskSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1EmptyDiskSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


