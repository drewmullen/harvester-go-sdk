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

// checks if the KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef{}

// KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef struct for KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef
type KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

type _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef(kind string, name string) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef{}
	this.Kind = kind
	this.Name = name
	return &this
}

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRefWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRefWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef{}
	var kind string = ""
	this.Kind = kind
	var name string = ""
	this.Name = name
	return &this
}

// GetKind returns the Kind field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetNamespace(v string) {
	o.Namespace = &v
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return toSerialize, nil
}

func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
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

	varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef := _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef)

	if err != nil {
		return err
	}

	*o = KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef(varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef)

	return err
}

type NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef struct {
	value *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef
	isSet bool
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) Get() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef {
	return v.value
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) Set(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef {
	return &NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef{value: val, isSet: true}
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


