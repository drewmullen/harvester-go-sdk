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

// checks if the KubevirtIoApiCoreV1TPMDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1TPMDevice{}

// KubevirtIoApiCoreV1TPMDevice struct for KubevirtIoApiCoreV1TPMDevice
type KubevirtIoApiCoreV1TPMDevice struct {
	// Persistent indicates the state of the TPM device should be kept accross reboots Defaults to false
	Persistent *bool `json:"persistent,omitempty"`
}

// NewKubevirtIoApiCoreV1TPMDevice instantiates a new KubevirtIoApiCoreV1TPMDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1TPMDevice() *KubevirtIoApiCoreV1TPMDevice {
	this := KubevirtIoApiCoreV1TPMDevice{}
	return &this
}

// NewKubevirtIoApiCoreV1TPMDeviceWithDefaults instantiates a new KubevirtIoApiCoreV1TPMDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1TPMDeviceWithDefaults() *KubevirtIoApiCoreV1TPMDevice {
	this := KubevirtIoApiCoreV1TPMDevice{}
	return &this
}

// GetPersistent returns the Persistent field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1TPMDevice) GetPersistent() bool {
	if o == nil || IsNil(o.Persistent) {
		var ret bool
		return ret
	}
	return *o.Persistent
}

// GetPersistentOk returns a tuple with the Persistent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1TPMDevice) GetPersistentOk() (*bool, bool) {
	if o == nil || IsNil(o.Persistent) {
		return nil, false
	}
	return o.Persistent, true
}

// HasPersistent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1TPMDevice) HasPersistent() bool {
	if o != nil && !IsNil(o.Persistent) {
		return true
	}

	return false
}

// SetPersistent gets a reference to the given bool and assigns it to the Persistent field.
func (o *KubevirtIoApiCoreV1TPMDevice) SetPersistent(v bool) {
	o.Persistent = &v
}

func (o KubevirtIoApiCoreV1TPMDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1TPMDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Persistent) {
		toSerialize["persistent"] = o.Persistent
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1TPMDevice struct {
	value *KubevirtIoApiCoreV1TPMDevice
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1TPMDevice) Get() *KubevirtIoApiCoreV1TPMDevice {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1TPMDevice) Set(val *KubevirtIoApiCoreV1TPMDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1TPMDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1TPMDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1TPMDevice(val *KubevirtIoApiCoreV1TPMDevice) *NullableKubevirtIoApiCoreV1TPMDevice {
	return &NullableKubevirtIoApiCoreV1TPMDevice{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1TPMDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1TPMDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


