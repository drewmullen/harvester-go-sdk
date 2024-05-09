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

// checks if the KubevirtIoApiCoreV1DiskTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1DiskTarget{}

// KubevirtIoApiCoreV1DiskTarget struct for KubevirtIoApiCoreV1DiskTarget
type KubevirtIoApiCoreV1DiskTarget struct {
	// Bus indicates the type of disk device to emulate. supported values: virtio, sata, scsi, usb.
	Bus *string `json:"bus,omitempty"`
	// If specified, the virtual disk will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10
	PciAddress *string `json:"pciAddress,omitempty"`
	// ReadOnly. Defaults to false.
	Readonly *bool `json:"readonly,omitempty"`
}

// NewKubevirtIoApiCoreV1DiskTarget instantiates a new KubevirtIoApiCoreV1DiskTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1DiskTarget() *KubevirtIoApiCoreV1DiskTarget {
	this := KubevirtIoApiCoreV1DiskTarget{}
	return &this
}

// NewKubevirtIoApiCoreV1DiskTargetWithDefaults instantiates a new KubevirtIoApiCoreV1DiskTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1DiskTargetWithDefaults() *KubevirtIoApiCoreV1DiskTarget {
	this := KubevirtIoApiCoreV1DiskTarget{}
	return &this
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DiskTarget) GetBus() string {
	if o == nil || IsNil(o.Bus) {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) GetBusOk() (*string, bool) {
	if o == nil || IsNil(o.Bus) {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) HasBus() bool {
	if o != nil && !IsNil(o.Bus) {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *KubevirtIoApiCoreV1DiskTarget) SetBus(v string) {
	o.Bus = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DiskTarget) GetPciAddress() string {
	if o == nil || IsNil(o.PciAddress) {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) GetPciAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddress) {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) HasPciAddress() bool {
	if o != nil && !IsNil(o.PciAddress) {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *KubevirtIoApiCoreV1DiskTarget) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DiskTarget) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DiskTarget) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *KubevirtIoApiCoreV1DiskTarget) SetReadonly(v bool) {
	o.Readonly = &v
}

func (o KubevirtIoApiCoreV1DiskTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1DiskTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bus) {
		toSerialize["bus"] = o.Bus
	}
	if !IsNil(o.PciAddress) {
		toSerialize["pciAddress"] = o.PciAddress
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1DiskTarget struct {
	value *KubevirtIoApiCoreV1DiskTarget
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1DiskTarget) Get() *KubevirtIoApiCoreV1DiskTarget {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1DiskTarget) Set(val *KubevirtIoApiCoreV1DiskTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1DiskTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1DiskTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1DiskTarget(val *KubevirtIoApiCoreV1DiskTarget) *NullableKubevirtIoApiCoreV1DiskTarget {
	return &NullableKubevirtIoApiCoreV1DiskTarget{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1DiskTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1DiskTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


