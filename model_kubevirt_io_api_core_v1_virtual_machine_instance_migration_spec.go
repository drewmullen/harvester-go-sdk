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

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec{}

// KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec struct for KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec
type KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec struct {
	// The name of the VMI to perform the migration on. VMI must exist in the migration objects namespace
	VmiName *string `json:"vmiName,omitempty"`
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec{}
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpecWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpecWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec{}
	return &this
}

// GetVmiName returns the VmiName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) GetVmiName() string {
	if o == nil || IsNil(o.VmiName) {
		var ret string
		return ret
	}
	return *o.VmiName
}

// GetVmiNameOk returns a tuple with the VmiName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) GetVmiNameOk() (*string, bool) {
	if o == nil || IsNil(o.VmiName) {
		return nil, false
	}
	return o.VmiName, true
}

// HasVmiName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) HasVmiName() bool {
	if o != nil && !IsNil(o.VmiName) {
		return true
	}

	return false
}

// SetVmiName gets a reference to the given string and assigns it to the VmiName field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) SetVmiName(v string) {
	o.VmiName = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VmiName) {
		toSerialize["vmiName"] = o.VmiName
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


