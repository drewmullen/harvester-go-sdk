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

// checks if the KubevirtIoApiCoreV1VirtualMachineList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineList{}

// KubevirtIoApiCoreV1VirtualMachineList struct for KubevirtIoApiCoreV1VirtualMachineList
type KubevirtIoApiCoreV1VirtualMachineList struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []KubevirtIoApiCoreV1VirtualMachine `json:"items"`
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ListMeta `json:"metadata,omitempty"`
}

type _KubevirtIoApiCoreV1VirtualMachineList KubevirtIoApiCoreV1VirtualMachineList

// NewKubevirtIoApiCoreV1VirtualMachineList instantiates a new KubevirtIoApiCoreV1VirtualMachineList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineList(items []KubevirtIoApiCoreV1VirtualMachine) *KubevirtIoApiCoreV1VirtualMachineList {
	this := KubevirtIoApiCoreV1VirtualMachineList{}
	this.Items = items
	var metadata K8sIoV1ListMeta
	this.Metadata = &metadata
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineListWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineListWithDefaults() *KubevirtIoApiCoreV1VirtualMachineList {
	this := KubevirtIoApiCoreV1VirtualMachineList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = &metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *KubevirtIoApiCoreV1VirtualMachineList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetItems() []KubevirtIoApiCoreV1VirtualMachine {
	if o == nil {
		var ret []KubevirtIoApiCoreV1VirtualMachine
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetItemsOk() ([]KubevirtIoApiCoreV1VirtualMachine, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineList) SetItems(v []KubevirtIoApiCoreV1VirtualMachine) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KubevirtIoApiCoreV1VirtualMachineList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetMetadata() K8sIoV1ListMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ListMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineList) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ListMeta and assigns it to the Metadata field.
func (o *KubevirtIoApiCoreV1VirtualMachineList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineList) ToMap() (map[string]interface{}, error) {
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

func (o *KubevirtIoApiCoreV1VirtualMachineList) UnmarshalJSON(data []byte) (err error) {
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

	varKubevirtIoApiCoreV1VirtualMachineList := _KubevirtIoApiCoreV1VirtualMachineList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineList)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineList(varKubevirtIoApiCoreV1VirtualMachineList)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineList struct {
	value *KubevirtIoApiCoreV1VirtualMachineList
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineList) Get() *KubevirtIoApiCoreV1VirtualMachineList {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineList) Set(val *KubevirtIoApiCoreV1VirtualMachineList) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineList) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineList(val *KubevirtIoApiCoreV1VirtualMachineList) *NullableKubevirtIoApiCoreV1VirtualMachineList {
	return &NullableKubevirtIoApiCoreV1VirtualMachineList{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


