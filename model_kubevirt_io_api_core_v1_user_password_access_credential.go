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

// checks if the KubevirtIoApiCoreV1UserPasswordAccessCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1UserPasswordAccessCredential{}

// KubevirtIoApiCoreV1UserPasswordAccessCredential UserPasswordAccessCredential represents a source and propagation method for injecting user passwords into a vm guest Only one of its members may be specified.
type KubevirtIoApiCoreV1UserPasswordAccessCredential struct {
	// propagationMethod represents how the user passwords are injected into the vm guest.
	PropagationMethod KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod `json:"propagationMethod"`
	// Source represents where the user passwords are pulled from
	Source KubevirtIoApiCoreV1UserPasswordAccessCredentialSource `json:"source"`
}

type _KubevirtIoApiCoreV1UserPasswordAccessCredential KubevirtIoApiCoreV1UserPasswordAccessCredential

// NewKubevirtIoApiCoreV1UserPasswordAccessCredential instantiates a new KubevirtIoApiCoreV1UserPasswordAccessCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1UserPasswordAccessCredential(propagationMethod KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod, source KubevirtIoApiCoreV1UserPasswordAccessCredentialSource) *KubevirtIoApiCoreV1UserPasswordAccessCredential {
	this := KubevirtIoApiCoreV1UserPasswordAccessCredential{}
	this.PropagationMethod = propagationMethod
	this.Source = source
	return &this
}

// NewKubevirtIoApiCoreV1UserPasswordAccessCredentialWithDefaults instantiates a new KubevirtIoApiCoreV1UserPasswordAccessCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1UserPasswordAccessCredentialWithDefaults() *KubevirtIoApiCoreV1UserPasswordAccessCredential {
	this := KubevirtIoApiCoreV1UserPasswordAccessCredential{}
	var propagationMethod KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod
	this.PropagationMethod = propagationMethod
	var source KubevirtIoApiCoreV1UserPasswordAccessCredentialSource
	this.Source = source
	return &this
}

// GetPropagationMethod returns the PropagationMethod field value
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) GetPropagationMethod() KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod {
	if o == nil {
		var ret KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod
		return ret
	}

	return o.PropagationMethod
}

// GetPropagationMethodOk returns a tuple with the PropagationMethod field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) GetPropagationMethodOk() (*KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropagationMethod, true
}

// SetPropagationMethod sets field value
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) SetPropagationMethod(v KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) {
	o.PropagationMethod = v
}

// GetSource returns the Source field value
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) GetSource() KubevirtIoApiCoreV1UserPasswordAccessCredentialSource {
	if o == nil {
		var ret KubevirtIoApiCoreV1UserPasswordAccessCredentialSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) GetSourceOk() (*KubevirtIoApiCoreV1UserPasswordAccessCredentialSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) SetSource(v KubevirtIoApiCoreV1UserPasswordAccessCredentialSource) {
	o.Source = v
}

func (o KubevirtIoApiCoreV1UserPasswordAccessCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1UserPasswordAccessCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["propagationMethod"] = o.PropagationMethod
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1UserPasswordAccessCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"propagationMethod",
		"source",
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

	varKubevirtIoApiCoreV1UserPasswordAccessCredential := _KubevirtIoApiCoreV1UserPasswordAccessCredential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1UserPasswordAccessCredential)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1UserPasswordAccessCredential(varKubevirtIoApiCoreV1UserPasswordAccessCredential)

	return err
}

type NullableKubevirtIoApiCoreV1UserPasswordAccessCredential struct {
	value *KubevirtIoApiCoreV1UserPasswordAccessCredential
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) Get() *KubevirtIoApiCoreV1UserPasswordAccessCredential {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) Set(val *KubevirtIoApiCoreV1UserPasswordAccessCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1UserPasswordAccessCredential(val *KubevirtIoApiCoreV1UserPasswordAccessCredential) *NullableKubevirtIoApiCoreV1UserPasswordAccessCredential {
	return &NullableKubevirtIoApiCoreV1UserPasswordAccessCredential{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


