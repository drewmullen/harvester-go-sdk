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

// checks if the KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot{}

// KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot DataVolumeSourceSnapshot provides the parameters to create a Data Volume from an existing VolumeSnapshot
type KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot struct {
	// The name of the source VolumeSnapshot
	Name string `json:"name"`
	// The namespace of the source VolumeSnapshot
	Namespace string `json:"namespace"`
}

type _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot(name string, namespace string) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot{}
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshotWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshotWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot{}
	var name string = ""
	this.Name = name
	var namespace string = ""
	this.Namespace = namespace
	return &this
}

// GetName returns the Name field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) SetNamespace(v string) {
	o.Namespace = v
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"namespace",
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

	varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot := _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot)

	if err != nil {
		return err
	}

	*o = KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot(varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot)

	return err
}

type NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot struct {
	value *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot
	isSet bool
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) Get() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot {
	return v.value
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) Set(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot {
	return &NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot{value: val, isSet: true}
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


