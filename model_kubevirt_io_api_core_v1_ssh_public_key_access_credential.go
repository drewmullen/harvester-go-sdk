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

// checks if the KubevirtIoApiCoreV1SSHPublicKeyAccessCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1SSHPublicKeyAccessCredential{}

// KubevirtIoApiCoreV1SSHPublicKeyAccessCredential SSHPublicKeyAccessCredential represents a source and propagation method for injecting ssh public keys into a vm guest
type KubevirtIoApiCoreV1SSHPublicKeyAccessCredential struct {
	// PropagationMethod represents how the public key is injected into the vm guest.
	PropagationMethod KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod `json:"propagationMethod"`
	// Source represents where the public keys are pulled from
	Source KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource `json:"source"`
}

type _KubevirtIoApiCoreV1SSHPublicKeyAccessCredential KubevirtIoApiCoreV1SSHPublicKeyAccessCredential

// NewKubevirtIoApiCoreV1SSHPublicKeyAccessCredential instantiates a new KubevirtIoApiCoreV1SSHPublicKeyAccessCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1SSHPublicKeyAccessCredential(propagationMethod KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod, source KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource) *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential {
	this := KubevirtIoApiCoreV1SSHPublicKeyAccessCredential{}
	this.PropagationMethod = propagationMethod
	this.Source = source
	return &this
}

// NewKubevirtIoApiCoreV1SSHPublicKeyAccessCredentialWithDefaults instantiates a new KubevirtIoApiCoreV1SSHPublicKeyAccessCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1SSHPublicKeyAccessCredentialWithDefaults() *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential {
	this := KubevirtIoApiCoreV1SSHPublicKeyAccessCredential{}
	var propagationMethod KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod
	this.PropagationMethod = propagationMethod
	var source KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource
	this.Source = source
	return &this
}

// GetPropagationMethod returns the PropagationMethod field value
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) GetPropagationMethod() KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod {
	if o == nil {
		var ret KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod
		return ret
	}

	return o.PropagationMethod
}

// GetPropagationMethodOk returns a tuple with the PropagationMethod field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) GetPropagationMethodOk() (*KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropagationMethod, true
}

// SetPropagationMethod sets field value
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) SetPropagationMethod(v KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialPropagationMethod) {
	o.PropagationMethod = v
}

// GetSource returns the Source field value
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) GetSource() KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource {
	if o == nil {
		var ret KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) GetSourceOk() (*KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) SetSource(v KubevirtIoApiCoreV1SSHPublicKeyAccessCredentialSource) {
	o.Source = v
}

func (o KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["propagationMethod"] = o.PropagationMethod
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) UnmarshalJSON(data []byte) (err error) {
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

	varKubevirtIoApiCoreV1SSHPublicKeyAccessCredential := _KubevirtIoApiCoreV1SSHPublicKeyAccessCredential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1SSHPublicKeyAccessCredential)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1SSHPublicKeyAccessCredential(varKubevirtIoApiCoreV1SSHPublicKeyAccessCredential)

	return err
}

type NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential struct {
	value *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) Get() *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) Set(val *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential(val *KubevirtIoApiCoreV1SSHPublicKeyAccessCredential) *NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential {
	return &NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1SSHPublicKeyAccessCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


