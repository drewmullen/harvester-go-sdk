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

// checks if the KubevirtIoApiCoreV1SoundDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1SoundDevice{}

// KubevirtIoApiCoreV1SoundDevice struct for KubevirtIoApiCoreV1SoundDevice
type KubevirtIoApiCoreV1SoundDevice struct {
	Model *string `json:"model,omitempty"`
	Name string `json:"name"`
}

type _KubevirtIoApiCoreV1SoundDevice KubevirtIoApiCoreV1SoundDevice

// NewKubevirtIoApiCoreV1SoundDevice instantiates a new KubevirtIoApiCoreV1SoundDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1SoundDevice(name string) *KubevirtIoApiCoreV1SoundDevice {
	this := KubevirtIoApiCoreV1SoundDevice{}
	this.Name = name
	return &this
}

// NewKubevirtIoApiCoreV1SoundDeviceWithDefaults instantiates a new KubevirtIoApiCoreV1SoundDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1SoundDeviceWithDefaults() *KubevirtIoApiCoreV1SoundDevice {
	this := KubevirtIoApiCoreV1SoundDevice{}
	var name string = ""
	this.Name = name
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1SoundDevice) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SoundDevice) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1SoundDevice) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *KubevirtIoApiCoreV1SoundDevice) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value
func (o *KubevirtIoApiCoreV1SoundDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SoundDevice) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoApiCoreV1SoundDevice) SetName(v string) {
	o.Name = v
}

func (o KubevirtIoApiCoreV1SoundDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1SoundDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1SoundDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varKubevirtIoApiCoreV1SoundDevice := _KubevirtIoApiCoreV1SoundDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1SoundDevice)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1SoundDevice(varKubevirtIoApiCoreV1SoundDevice)

	return err
}

type NullableKubevirtIoApiCoreV1SoundDevice struct {
	value *KubevirtIoApiCoreV1SoundDevice
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1SoundDevice) Get() *KubevirtIoApiCoreV1SoundDevice {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1SoundDevice) Set(val *KubevirtIoApiCoreV1SoundDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1SoundDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1SoundDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1SoundDevice(val *KubevirtIoApiCoreV1SoundDevice) *NullableKubevirtIoApiCoreV1SoundDevice {
	return &NullableKubevirtIoApiCoreV1SoundDevice{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1SoundDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1SoundDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


