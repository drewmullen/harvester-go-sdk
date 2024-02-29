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

// checks if the KubevirtIoApiCoreV1CPUFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1CPUFeature{}

// KubevirtIoApiCoreV1CPUFeature struct for KubevirtIoApiCoreV1CPUFeature
type KubevirtIoApiCoreV1CPUFeature struct {
	Name string `json:"name"`
	Policy *string `json:"policy,omitempty"`
}

type _KubevirtIoApiCoreV1CPUFeature KubevirtIoApiCoreV1CPUFeature

// NewKubevirtIoApiCoreV1CPUFeature instantiates a new KubevirtIoApiCoreV1CPUFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1CPUFeature(name string) *KubevirtIoApiCoreV1CPUFeature {
	this := KubevirtIoApiCoreV1CPUFeature{}
	this.Name = name
	return &this
}

// NewKubevirtIoApiCoreV1CPUFeatureWithDefaults instantiates a new KubevirtIoApiCoreV1CPUFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1CPUFeatureWithDefaults() *KubevirtIoApiCoreV1CPUFeature {
	this := KubevirtIoApiCoreV1CPUFeature{}
	var name string = ""
	this.Name = name
	return &this
}

// GetName returns the Name field value
func (o *KubevirtIoApiCoreV1CPUFeature) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CPUFeature) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoApiCoreV1CPUFeature) SetName(v string) {
	o.Name = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CPUFeature) GetPolicy() string {
	if o == nil || IsNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CPUFeature) GetPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CPUFeature) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *KubevirtIoApiCoreV1CPUFeature) SetPolicy(v string) {
	o.Policy = &v
}

func (o KubevirtIoApiCoreV1CPUFeature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1CPUFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1CPUFeature) UnmarshalJSON(data []byte) (err error) {
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

	varKubevirtIoApiCoreV1CPUFeature := _KubevirtIoApiCoreV1CPUFeature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1CPUFeature)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1CPUFeature(varKubevirtIoApiCoreV1CPUFeature)

	return err
}

type NullableKubevirtIoApiCoreV1CPUFeature struct {
	value *KubevirtIoApiCoreV1CPUFeature
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1CPUFeature) Get() *KubevirtIoApiCoreV1CPUFeature {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1CPUFeature) Set(val *KubevirtIoApiCoreV1CPUFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1CPUFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1CPUFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1CPUFeature(val *KubevirtIoApiCoreV1CPUFeature) *NullableKubevirtIoApiCoreV1CPUFeature {
	return &NullableKubevirtIoApiCoreV1CPUFeature{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1CPUFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1CPUFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


