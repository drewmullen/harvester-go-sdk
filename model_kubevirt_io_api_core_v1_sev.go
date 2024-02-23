/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1SEV type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1SEV{}

// KubevirtIoApiCoreV1SEV struct for KubevirtIoApiCoreV1SEV
type KubevirtIoApiCoreV1SEV struct {
	Attestation map[string]interface{} `json:"attestation,omitempty"`
	DhCert *string `json:"dhCert,omitempty"`
	Policy *KubevirtIoApiCoreV1SEVPolicy `json:"policy,omitempty"`
	Session *string `json:"session,omitempty"`
}

// NewKubevirtIoApiCoreV1SEV instantiates a new KubevirtIoApiCoreV1SEV object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1SEV() *KubevirtIoApiCoreV1SEV {
	this := KubevirtIoApiCoreV1SEV{}
	return &this
}

// NewKubevirtIoApiCoreV1SEVWithDefaults instantiates a new KubevirtIoApiCoreV1SEV object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1SEVWithDefaults() *KubevirtIoApiCoreV1SEV {
	this := KubevirtIoApiCoreV1SEV{}
	return &this
}

// GetAttestation returns the Attestation field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1SEV) GetAttestation() map[string]interface{} {
	if o == nil || IsNil(o.Attestation) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attestation
}

// GetAttestationOk returns a tuple with the Attestation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SEV) GetAttestationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attestation) {
		return map[string]interface{}{}, false
	}
	return o.Attestation, true
}

// HasAttestation returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1SEV) HasAttestation() bool {
	if o != nil && !IsNil(o.Attestation) {
		return true
	}

	return false
}

// SetAttestation gets a reference to the given map[string]interface{} and assigns it to the Attestation field.
func (o *KubevirtIoApiCoreV1SEV) SetAttestation(v map[string]interface{}) {
	o.Attestation = v
}

// GetDhCert returns the DhCert field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1SEV) GetDhCert() string {
	if o == nil || IsNil(o.DhCert) {
		var ret string
		return ret
	}
	return *o.DhCert
}

// GetDhCertOk returns a tuple with the DhCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SEV) GetDhCertOk() (*string, bool) {
	if o == nil || IsNil(o.DhCert) {
		return nil, false
	}
	return o.DhCert, true
}

// HasDhCert returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1SEV) HasDhCert() bool {
	if o != nil && !IsNil(o.DhCert) {
		return true
	}

	return false
}

// SetDhCert gets a reference to the given string and assigns it to the DhCert field.
func (o *KubevirtIoApiCoreV1SEV) SetDhCert(v string) {
	o.DhCert = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1SEV) GetPolicy() KubevirtIoApiCoreV1SEVPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret KubevirtIoApiCoreV1SEVPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SEV) GetPolicyOk() (*KubevirtIoApiCoreV1SEVPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1SEV) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given KubevirtIoApiCoreV1SEVPolicy and assigns it to the Policy field.
func (o *KubevirtIoApiCoreV1SEV) SetPolicy(v KubevirtIoApiCoreV1SEVPolicy) {
	o.Policy = &v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1SEV) GetSession() string {
	if o == nil || IsNil(o.Session) {
		var ret string
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1SEV) GetSessionOk() (*string, bool) {
	if o == nil || IsNil(o.Session) {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1SEV) HasSession() bool {
	if o != nil && !IsNil(o.Session) {
		return true
	}

	return false
}

// SetSession gets a reference to the given string and assigns it to the Session field.
func (o *KubevirtIoApiCoreV1SEV) SetSession(v string) {
	o.Session = &v
}

func (o KubevirtIoApiCoreV1SEV) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1SEV) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attestation) {
		toSerialize["attestation"] = o.Attestation
	}
	if !IsNil(o.DhCert) {
		toSerialize["dhCert"] = o.DhCert
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !IsNil(o.Session) {
		toSerialize["session"] = o.Session
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1SEV struct {
	value *KubevirtIoApiCoreV1SEV
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1SEV) Get() *KubevirtIoApiCoreV1SEV {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1SEV) Set(val *KubevirtIoApiCoreV1SEV) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1SEV) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1SEV) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1SEV(val *KubevirtIoApiCoreV1SEV) *NullableKubevirtIoApiCoreV1SEV {
	return &NullableKubevirtIoApiCoreV1SEV{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1SEV) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1SEV) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


