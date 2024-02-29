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

// checks if the K8sIoV1HTTPHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1HTTPHeader{}

// K8sIoV1HTTPHeader struct for K8sIoV1HTTPHeader
type K8sIoV1HTTPHeader struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type _K8sIoV1HTTPHeader K8sIoV1HTTPHeader

// NewK8sIoV1HTTPHeader instantiates a new K8sIoV1HTTPHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1HTTPHeader(name string, value string) *K8sIoV1HTTPHeader {
	this := K8sIoV1HTTPHeader{}
	this.Name = name
	this.Value = value
	return &this
}

// NewK8sIoV1HTTPHeaderWithDefaults instantiates a new K8sIoV1HTTPHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1HTTPHeaderWithDefaults() *K8sIoV1HTTPHeader {
	this := K8sIoV1HTTPHeader{}
	var name string = ""
	this.Name = name
	var value string = ""
	this.Value = value
	return &this
}

// GetName returns the Name field value
func (o *K8sIoV1HTTPHeader) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1HTTPHeader) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *K8sIoV1HTTPHeader) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *K8sIoV1HTTPHeader) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1HTTPHeader) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *K8sIoV1HTTPHeader) SetValue(v string) {
	o.Value = v
}

func (o K8sIoV1HTTPHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1HTTPHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *K8sIoV1HTTPHeader) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varK8sIoV1HTTPHeader := _K8sIoV1HTTPHeader{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sIoV1HTTPHeader)

	if err != nil {
		return err
	}

	*o = K8sIoV1HTTPHeader(varK8sIoV1HTTPHeader)

	return err
}

type NullableK8sIoV1HTTPHeader struct {
	value *K8sIoV1HTTPHeader
	isSet bool
}

func (v NullableK8sIoV1HTTPHeader) Get() *K8sIoV1HTTPHeader {
	return v.value
}

func (v *NullableK8sIoV1HTTPHeader) Set(val *K8sIoV1HTTPHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1HTTPHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1HTTPHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1HTTPHeader(val *K8sIoV1HTTPHeader) *NullableK8sIoV1HTTPHeader {
	return &NullableK8sIoV1HTTPHeader{value: val, isSet: true}
}

func (v NullableK8sIoV1HTTPHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1HTTPHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


