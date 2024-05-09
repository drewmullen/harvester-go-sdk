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

// checks if the KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod{}

// KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod UserPasswordAccessCredentialPropagationMethod represents the method used to inject a user passwords into the vm guest. Only one of its members may be specified.
type KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod struct {
	// QemuGuestAgentAccessCredentailPropagation means passwords are dynamically injected into the vm at runtime via the qemu guest agent. This feature requires the qemu guest agent to be running within the guest.
	QemuGuestAgent map[string]interface{} `json:"qemuGuestAgent,omitempty"`
}

// NewKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod instantiates a new KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod() *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod {
	this := KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod{}
	return &this
}

// NewKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethodWithDefaults instantiates a new KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethodWithDefaults() *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod {
	this := KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod{}
	return &this
}

// GetQemuGuestAgent returns the QemuGuestAgent field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) GetQemuGuestAgent() map[string]interface{} {
	if o == nil || IsNil(o.QemuGuestAgent) {
		var ret map[string]interface{}
		return ret
	}
	return o.QemuGuestAgent
}

// GetQemuGuestAgentOk returns a tuple with the QemuGuestAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) GetQemuGuestAgentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.QemuGuestAgent) {
		return map[string]interface{}{}, false
	}
	return o.QemuGuestAgent, true
}

// HasQemuGuestAgent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) HasQemuGuestAgent() bool {
	if o != nil && !IsNil(o.QemuGuestAgent) {
		return true
	}

	return false
}

// SetQemuGuestAgent gets a reference to the given map[string]interface{} and assigns it to the QemuGuestAgent field.
func (o *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) SetQemuGuestAgent(v map[string]interface{}) {
	o.QemuGuestAgent = v
}

func (o KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QemuGuestAgent) {
		toSerialize["qemuGuestAgent"] = o.QemuGuestAgent
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod struct {
	value *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) Get() *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) Set(val *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod(val *KubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) *NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod {
	return &NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1UserPasswordAccessCredentialPropagationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


