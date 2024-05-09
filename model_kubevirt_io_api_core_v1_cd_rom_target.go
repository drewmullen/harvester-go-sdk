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

// checks if the KubevirtIoApiCoreV1CDRomTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1CDRomTarget{}

// KubevirtIoApiCoreV1CDRomTarget struct for KubevirtIoApiCoreV1CDRomTarget
type KubevirtIoApiCoreV1CDRomTarget struct {
	// Bus indicates the type of disk device to emulate. supported values: virtio, sata, scsi.
	Bus *string `json:"bus,omitempty"`
	// ReadOnly. Defaults to true.
	Readonly *bool `json:"readonly,omitempty"`
	// Tray indicates if the tray of the device is open or closed. Allowed values are \"open\" and \"closed\". Defaults to closed.
	Tray *string `json:"tray,omitempty"`
}

// NewKubevirtIoApiCoreV1CDRomTarget instantiates a new KubevirtIoApiCoreV1CDRomTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1CDRomTarget() *KubevirtIoApiCoreV1CDRomTarget {
	this := KubevirtIoApiCoreV1CDRomTarget{}
	return &this
}

// NewKubevirtIoApiCoreV1CDRomTargetWithDefaults instantiates a new KubevirtIoApiCoreV1CDRomTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1CDRomTargetWithDefaults() *KubevirtIoApiCoreV1CDRomTarget {
	this := KubevirtIoApiCoreV1CDRomTarget{}
	return &this
}

// GetBus returns the Bus field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetBus() string {
	if o == nil || IsNil(o.Bus) {
		var ret string
		return ret
	}
	return *o.Bus
}

// GetBusOk returns a tuple with the Bus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetBusOk() (*string, bool) {
	if o == nil || IsNil(o.Bus) {
		return nil, false
	}
	return o.Bus, true
}

// HasBus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) HasBus() bool {
	if o != nil && !IsNil(o.Bus) {
		return true
	}

	return false
}

// SetBus gets a reference to the given string and assigns it to the Bus field.
func (o *KubevirtIoApiCoreV1CDRomTarget) SetBus(v string) {
	o.Bus = &v
}

// GetReadonly returns the Readonly field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetReadonly() bool {
	if o == nil || IsNil(o.Readonly) {
		var ret bool
		return ret
	}
	return *o.Readonly
}

// GetReadonlyOk returns a tuple with the Readonly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetReadonlyOk() (*bool, bool) {
	if o == nil || IsNil(o.Readonly) {
		return nil, false
	}
	return o.Readonly, true
}

// HasReadonly returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) HasReadonly() bool {
	if o != nil && !IsNil(o.Readonly) {
		return true
	}

	return false
}

// SetReadonly gets a reference to the given bool and assigns it to the Readonly field.
func (o *KubevirtIoApiCoreV1CDRomTarget) SetReadonly(v bool) {
	o.Readonly = &v
}

// GetTray returns the Tray field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetTray() string {
	if o == nil || IsNil(o.Tray) {
		var ret string
		return ret
	}
	return *o.Tray
}

// GetTrayOk returns a tuple with the Tray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) GetTrayOk() (*string, bool) {
	if o == nil || IsNil(o.Tray) {
		return nil, false
	}
	return o.Tray, true
}

// HasTray returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1CDRomTarget) HasTray() bool {
	if o != nil && !IsNil(o.Tray) {
		return true
	}

	return false
}

// SetTray gets a reference to the given string and assigns it to the Tray field.
func (o *KubevirtIoApiCoreV1CDRomTarget) SetTray(v string) {
	o.Tray = &v
}

func (o KubevirtIoApiCoreV1CDRomTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1CDRomTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bus) {
		toSerialize["bus"] = o.Bus
	}
	if !IsNil(o.Readonly) {
		toSerialize["readonly"] = o.Readonly
	}
	if !IsNil(o.Tray) {
		toSerialize["tray"] = o.Tray
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1CDRomTarget struct {
	value *KubevirtIoApiCoreV1CDRomTarget
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1CDRomTarget) Get() *KubevirtIoApiCoreV1CDRomTarget {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1CDRomTarget) Set(val *KubevirtIoApiCoreV1CDRomTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1CDRomTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1CDRomTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1CDRomTarget(val *KubevirtIoApiCoreV1CDRomTarget) *NullableKubevirtIoApiCoreV1CDRomTarget {
	return &NullableKubevirtIoApiCoreV1CDRomTarget{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1CDRomTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1CDRomTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


