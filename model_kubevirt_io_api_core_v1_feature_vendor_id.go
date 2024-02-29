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

// checks if the KubevirtIoApiCoreV1FeatureVendorID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1FeatureVendorID{}

// KubevirtIoApiCoreV1FeatureVendorID struct for KubevirtIoApiCoreV1FeatureVendorID
type KubevirtIoApiCoreV1FeatureVendorID struct {
	Enabled *bool `json:"enabled,omitempty"`
	Vendorid *string `json:"vendorid,omitempty"`
}

// NewKubevirtIoApiCoreV1FeatureVendorID instantiates a new KubevirtIoApiCoreV1FeatureVendorID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1FeatureVendorID() *KubevirtIoApiCoreV1FeatureVendorID {
	this := KubevirtIoApiCoreV1FeatureVendorID{}
	return &this
}

// NewKubevirtIoApiCoreV1FeatureVendorIDWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureVendorID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1FeatureVendorIDWithDefaults() *KubevirtIoApiCoreV1FeatureVendorID {
	this := KubevirtIoApiCoreV1FeatureVendorID{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureVendorID) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureVendorID) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureVendorID) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *KubevirtIoApiCoreV1FeatureVendorID) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVendorid returns the Vendorid field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureVendorID) GetVendorid() string {
	if o == nil || IsNil(o.Vendorid) {
		var ret string
		return ret
	}
	return *o.Vendorid
}

// GetVendoridOk returns a tuple with the Vendorid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureVendorID) GetVendoridOk() (*string, bool) {
	if o == nil || IsNil(o.Vendorid) {
		return nil, false
	}
	return o.Vendorid, true
}

// HasVendorid returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureVendorID) HasVendorid() bool {
	if o != nil && !IsNil(o.Vendorid) {
		return true
	}

	return false
}

// SetVendorid gets a reference to the given string and assigns it to the Vendorid field.
func (o *KubevirtIoApiCoreV1FeatureVendorID) SetVendorid(v string) {
	o.Vendorid = &v
}

func (o KubevirtIoApiCoreV1FeatureVendorID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1FeatureVendorID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Vendorid) {
		toSerialize["vendorid"] = o.Vendorid
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1FeatureVendorID struct {
	value *KubevirtIoApiCoreV1FeatureVendorID
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1FeatureVendorID) Get() *KubevirtIoApiCoreV1FeatureVendorID {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1FeatureVendorID) Set(val *KubevirtIoApiCoreV1FeatureVendorID) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1FeatureVendorID) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1FeatureVendorID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1FeatureVendorID(val *KubevirtIoApiCoreV1FeatureVendorID) *NullableKubevirtIoApiCoreV1FeatureVendorID {
	return &NullableKubevirtIoApiCoreV1FeatureVendorID{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1FeatureVendorID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1FeatureVendorID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


