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

// checks if the KubevirtIoApiCoreV1RemoveVolumeOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1RemoveVolumeOptions{}

// KubevirtIoApiCoreV1RemoveVolumeOptions struct for KubevirtIoApiCoreV1RemoveVolumeOptions
type KubevirtIoApiCoreV1RemoveVolumeOptions struct {
	DryRun []string `json:"dryRun,omitempty"`
	Name string `json:"name"`
}

type _KubevirtIoApiCoreV1RemoveVolumeOptions KubevirtIoApiCoreV1RemoveVolumeOptions

// NewKubevirtIoApiCoreV1RemoveVolumeOptions instantiates a new KubevirtIoApiCoreV1RemoveVolumeOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1RemoveVolumeOptions(name string) *KubevirtIoApiCoreV1RemoveVolumeOptions {
	this := KubevirtIoApiCoreV1RemoveVolumeOptions{}
	this.Name = name
	return &this
}

// NewKubevirtIoApiCoreV1RemoveVolumeOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1RemoveVolumeOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1RemoveVolumeOptionsWithDefaults() *KubevirtIoApiCoreV1RemoveVolumeOptions {
	this := KubevirtIoApiCoreV1RemoveVolumeOptions{}
	var name string = ""
	this.Name = name
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetDryRun() []string {
	if o == nil || IsNil(o.DryRun) {
		var ret []string
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetDryRunOk() ([]string, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given []string and assigns it to the DryRun field.
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) SetDryRun(v []string) {
	o.DryRun = v
}

// GetName returns the Name field value
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) SetName(v string) {
	o.Name = v
}

func (o KubevirtIoApiCoreV1RemoveVolumeOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1RemoveVolumeOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varKubevirtIoApiCoreV1RemoveVolumeOptions := _KubevirtIoApiCoreV1RemoveVolumeOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1RemoveVolumeOptions)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1RemoveVolumeOptions(varKubevirtIoApiCoreV1RemoveVolumeOptions)

	return err
}

type NullableKubevirtIoApiCoreV1RemoveVolumeOptions struct {
	value *KubevirtIoApiCoreV1RemoveVolumeOptions
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1RemoveVolumeOptions) Get() *KubevirtIoApiCoreV1RemoveVolumeOptions {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1RemoveVolumeOptions) Set(val *KubevirtIoApiCoreV1RemoveVolumeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1RemoveVolumeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1RemoveVolumeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1RemoveVolumeOptions(val *KubevirtIoApiCoreV1RemoveVolumeOptions) *NullableKubevirtIoApiCoreV1RemoveVolumeOptions {
	return &NullableKubevirtIoApiCoreV1RemoveVolumeOptions{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1RemoveVolumeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1RemoveVolumeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


