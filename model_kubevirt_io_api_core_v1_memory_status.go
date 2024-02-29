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

// checks if the KubevirtIoApiCoreV1MemoryStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1MemoryStatus{}

// KubevirtIoApiCoreV1MemoryStatus struct for KubevirtIoApiCoreV1MemoryStatus
type KubevirtIoApiCoreV1MemoryStatus struct {
	GuestAtBoot *string `json:"guestAtBoot,omitempty"`
	GuestCurrent *string `json:"guestCurrent,omitempty"`
	GuestRequested *string `json:"guestRequested,omitempty"`
}

// NewKubevirtIoApiCoreV1MemoryStatus instantiates a new KubevirtIoApiCoreV1MemoryStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1MemoryStatus() *KubevirtIoApiCoreV1MemoryStatus {
	this := KubevirtIoApiCoreV1MemoryStatus{}
	return &this
}

// NewKubevirtIoApiCoreV1MemoryStatusWithDefaults instantiates a new KubevirtIoApiCoreV1MemoryStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1MemoryStatusWithDefaults() *KubevirtIoApiCoreV1MemoryStatus {
	this := KubevirtIoApiCoreV1MemoryStatus{}
	return &this
}

// GetGuestAtBoot returns the GuestAtBoot field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestAtBoot() string {
	if o == nil || IsNil(o.GuestAtBoot) {
		var ret string
		return ret
	}
	return *o.GuestAtBoot
}

// GetGuestAtBootOk returns a tuple with the GuestAtBoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestAtBootOk() (*string, bool) {
	if o == nil || IsNil(o.GuestAtBoot) {
		return nil, false
	}
	return o.GuestAtBoot, true
}

// HasGuestAtBoot returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) HasGuestAtBoot() bool {
	if o != nil && !IsNil(o.GuestAtBoot) {
		return true
	}

	return false
}

// SetGuestAtBoot gets a reference to the given string and assigns it to the GuestAtBoot field.
func (o *KubevirtIoApiCoreV1MemoryStatus) SetGuestAtBoot(v string) {
	o.GuestAtBoot = &v
}

// GetGuestCurrent returns the GuestCurrent field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestCurrent() string {
	if o == nil || IsNil(o.GuestCurrent) {
		var ret string
		return ret
	}
	return *o.GuestCurrent
}

// GetGuestCurrentOk returns a tuple with the GuestCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestCurrentOk() (*string, bool) {
	if o == nil || IsNil(o.GuestCurrent) {
		return nil, false
	}
	return o.GuestCurrent, true
}

// HasGuestCurrent returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) HasGuestCurrent() bool {
	if o != nil && !IsNil(o.GuestCurrent) {
		return true
	}

	return false
}

// SetGuestCurrent gets a reference to the given string and assigns it to the GuestCurrent field.
func (o *KubevirtIoApiCoreV1MemoryStatus) SetGuestCurrent(v string) {
	o.GuestCurrent = &v
}

// GetGuestRequested returns the GuestRequested field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestRequested() string {
	if o == nil || IsNil(o.GuestRequested) {
		var ret string
		return ret
	}
	return *o.GuestRequested
}

// GetGuestRequestedOk returns a tuple with the GuestRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) GetGuestRequestedOk() (*string, bool) {
	if o == nil || IsNil(o.GuestRequested) {
		return nil, false
	}
	return o.GuestRequested, true
}

// HasGuestRequested returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1MemoryStatus) HasGuestRequested() bool {
	if o != nil && !IsNil(o.GuestRequested) {
		return true
	}

	return false
}

// SetGuestRequested gets a reference to the given string and assigns it to the GuestRequested field.
func (o *KubevirtIoApiCoreV1MemoryStatus) SetGuestRequested(v string) {
	o.GuestRequested = &v
}

func (o KubevirtIoApiCoreV1MemoryStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1MemoryStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestAtBoot) {
		toSerialize["guestAtBoot"] = o.GuestAtBoot
	}
	if !IsNil(o.GuestCurrent) {
		toSerialize["guestCurrent"] = o.GuestCurrent
	}
	if !IsNil(o.GuestRequested) {
		toSerialize["guestRequested"] = o.GuestRequested
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1MemoryStatus struct {
	value *KubevirtIoApiCoreV1MemoryStatus
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1MemoryStatus) Get() *KubevirtIoApiCoreV1MemoryStatus {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1MemoryStatus) Set(val *KubevirtIoApiCoreV1MemoryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1MemoryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1MemoryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1MemoryStatus(val *KubevirtIoApiCoreV1MemoryStatus) *NullableKubevirtIoApiCoreV1MemoryStatus {
	return &NullableKubevirtIoApiCoreV1MemoryStatus{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1MemoryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1MemoryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


