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

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceList{}

// KubevirtIoApiCoreV1VirtualMachineInstanceList VirtualMachineInstanceList is a list of VirtualMachines
type KubevirtIoApiCoreV1VirtualMachineInstanceList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []KubevirtIoApiCoreV1VirtualMachineInstance `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ListMeta `json:"metadata,omitempty"`
}

type _KubevirtIoApiCoreV1VirtualMachineInstanceList KubevirtIoApiCoreV1VirtualMachineInstanceList

// NewKubevirtIoApiCoreV1VirtualMachineInstanceList instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceList(items []KubevirtIoApiCoreV1VirtualMachineInstance) *KubevirtIoApiCoreV1VirtualMachineInstanceList {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceList{}
	this.Items = items
	var metadata K8sIoV1ListMeta
	this.Metadata = &metadata
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceListWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceListWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceList {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = &metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetItems() []KubevirtIoApiCoreV1VirtualMachineInstance {
	if o == nil {
		var ret []KubevirtIoApiCoreV1VirtualMachineInstance
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetItemsOk() ([]KubevirtIoApiCoreV1VirtualMachineInstance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetItems(v []KubevirtIoApiCoreV1VirtualMachineInstance) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetMetadata() K8sIoV1ListMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ListMeta and assigns it to the Metadata field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varKubevirtIoApiCoreV1VirtualMachineInstanceList := _KubevirtIoApiCoreV1VirtualMachineInstanceList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineInstanceList)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineInstanceList(varKubevirtIoApiCoreV1VirtualMachineInstanceList)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceList struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceList
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceList {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceList(val *KubevirtIoApiCoreV1VirtualMachineInstanceList) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceList {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceList{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


