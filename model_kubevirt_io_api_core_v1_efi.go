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

// checks if the KubevirtIoApiCoreV1EFI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1EFI{}

// KubevirtIoApiCoreV1EFI If set, EFI will be used instead of BIOS.
type KubevirtIoApiCoreV1EFI struct {
	// If set to true, Persistent will persist the EFI NVRAM across reboots. Defaults to false
	Persistent *bool `json:"persistent,omitempty"`
	// If set, SecureBoot will be enabled and the OVMF roms will be swapped for SecureBoot-enabled ones. Requires SMM to be enabled. Defaults to true
	SecureBoot *bool `json:"secureBoot,omitempty"`
}

// NewKubevirtIoApiCoreV1EFI instantiates a new KubevirtIoApiCoreV1EFI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1EFI() *KubevirtIoApiCoreV1EFI {
	this := KubevirtIoApiCoreV1EFI{}
	return &this
}

// NewKubevirtIoApiCoreV1EFIWithDefaults instantiates a new KubevirtIoApiCoreV1EFI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1EFIWithDefaults() *KubevirtIoApiCoreV1EFI {
	this := KubevirtIoApiCoreV1EFI{}
	return &this
}

// GetPersistent returns the Persistent field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1EFI) GetPersistent() bool {
	if o == nil || IsNil(o.Persistent) {
		var ret bool
		return ret
	}
	return *o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1EFI) GetPersistentOk() (*bool, bool) {
	if o == nil || IsNil(o.Persistent) {
		return nil, false
	}
	return o.Persistent, true
}

// HasPersistent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1EFI) HasPersistent() bool {
	if o != nil && !IsNil(o.Persistent) {
		return true
	}

	return false
}

// SetPersistent gets a reference to the given bool and assigns it to the Persistent field.
func (o *KubevirtIoApiCoreV1EFI) SetPersistent(v bool) {
	o.Persistent = &v
}

// GetSecureBoot returns the SecureBoot field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1EFI) GetSecureBoot() bool {
	if o == nil || IsNil(o.SecureBoot) {
		var ret bool
		return ret
	}
	return *o.SecureBoot
}

// GetSecureBootOk returns a tuple with the SecureBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1EFI) GetSecureBootOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureBoot) {
		return nil, false
	}
	return o.SecureBoot, true
}

// HasSecureBoot returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1EFI) HasSecureBoot() bool {
	if o != nil && !IsNil(o.SecureBoot) {
		return true
	}

	return false
}

// SetSecureBoot gets a reference to the given bool and assigns it to the SecureBoot field.
func (o *KubevirtIoApiCoreV1EFI) SetSecureBoot(v bool) {
	o.SecureBoot = &v
}

func (o KubevirtIoApiCoreV1EFI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1EFI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Persistent) {
		toSerialize["persistent"] = o.Persistent
	}
	if !IsNil(o.SecureBoot) {
		toSerialize["secureBoot"] = o.SecureBoot
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1EFI struct {
	value *KubevirtIoApiCoreV1EFI
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1EFI) Get() *KubevirtIoApiCoreV1EFI {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1EFI) Set(val *KubevirtIoApiCoreV1EFI) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1EFI) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1EFI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1EFI(val *KubevirtIoApiCoreV1EFI) *NullableKubevirtIoApiCoreV1EFI {
	return &NullableKubevirtIoApiCoreV1EFI{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1EFI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1EFI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


