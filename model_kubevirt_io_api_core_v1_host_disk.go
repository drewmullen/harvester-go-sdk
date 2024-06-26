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

// checks if the KubevirtIoApiCoreV1HostDisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1HostDisk{}

// KubevirtIoApiCoreV1HostDisk Represents a disk created on the cluster level
type KubevirtIoApiCoreV1HostDisk struct {
	// Capacity of the sparse disk
	Capacity *string `json:"capacity,omitempty"`
	// The path to HostDisk image located on the cluster
	Path string `json:"path"`
	// Shared indicate whether the path is shared between nodes
	Shared *bool `json:"shared,omitempty"`
	// Contains information if disk.img exists or should be created allowed options are 'Disk' and 'DiskOrCreate'
	Type string `json:"type"`
}

type _KubevirtIoApiCoreV1HostDisk KubevirtIoApiCoreV1HostDisk

// NewKubevirtIoApiCoreV1HostDisk instantiates a new KubevirtIoApiCoreV1HostDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1HostDisk(path string, type_ string) *KubevirtIoApiCoreV1HostDisk {
	this := KubevirtIoApiCoreV1HostDisk{}
	var capacity string = "{}"
	this.Capacity = &capacity
	this.Path = path
	this.Type = type_
	return &this
}

// NewKubevirtIoApiCoreV1HostDiskWithDefaults instantiates a new KubevirtIoApiCoreV1HostDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1HostDiskWithDefaults() *KubevirtIoApiCoreV1HostDisk {
	this := KubevirtIoApiCoreV1HostDisk{}
	var capacity string = "{}"
	this.Capacity = &capacity
	var path string = ""
	this.Path = path
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1HostDisk) GetCapacity() string {
	if o == nil || IsNil(o.Capacity) {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HostDisk) GetCapacityOk() (*string, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1HostDisk) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *KubevirtIoApiCoreV1HostDisk) SetCapacity(v string) {
	o.Capacity = &v
}

// GetPath returns the Path field value
func (o *KubevirtIoApiCoreV1HostDisk) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HostDisk) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *KubevirtIoApiCoreV1HostDisk) SetPath(v string) {
	o.Path = v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1HostDisk) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HostDisk) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1HostDisk) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *KubevirtIoApiCoreV1HostDisk) SetShared(v bool) {
	o.Shared = &v
}

// GetType returns the Type field value
func (o *KubevirtIoApiCoreV1HostDisk) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1HostDisk) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KubevirtIoApiCoreV1HostDisk) SetType(v string) {
	o.Type = v
}

func (o KubevirtIoApiCoreV1HostDisk) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1HostDisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1HostDisk) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
		"type",
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

	varKubevirtIoApiCoreV1HostDisk := _KubevirtIoApiCoreV1HostDisk{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1HostDisk)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1HostDisk(varKubevirtIoApiCoreV1HostDisk)

	return err
}

type NullableKubevirtIoApiCoreV1HostDisk struct {
	value *KubevirtIoApiCoreV1HostDisk
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1HostDisk) Get() *KubevirtIoApiCoreV1HostDisk {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1HostDisk) Set(val *KubevirtIoApiCoreV1HostDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1HostDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1HostDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1HostDisk(val *KubevirtIoApiCoreV1HostDisk) *NullableKubevirtIoApiCoreV1HostDisk {
	return &NullableKubevirtIoApiCoreV1HostDisk{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1HostDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1HostDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


