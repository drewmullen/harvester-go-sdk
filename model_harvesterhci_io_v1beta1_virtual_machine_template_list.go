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

// checks if the HarvesterhciIoV1beta1VirtualMachineTemplateList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineTemplateList{}

// HarvesterhciIoV1beta1VirtualMachineTemplateList struct for HarvesterhciIoV1beta1VirtualMachineTemplateList
type HarvesterhciIoV1beta1VirtualMachineTemplateList struct {
	ApiVersion string `json:"apiVersion"`
	Items []HarvesterhciIoV1beta1VirtualMachineTemplate `json:"items"`
	Kind string `json:"kind"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _HarvesterhciIoV1beta1VirtualMachineTemplateList HarvesterhciIoV1beta1VirtualMachineTemplateList

// NewHarvesterhciIoV1beta1VirtualMachineTemplateList instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineTemplateList(apiVersion string, items []HarvesterhciIoV1beta1VirtualMachineTemplate, kind string, metadata K8sIoV1ListMeta) *HarvesterhciIoV1beta1VirtualMachineTemplateList {
	this := HarvesterhciIoV1beta1VirtualMachineTemplateList{}
	this.ApiVersion = apiVersion
	this.Items = items
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineTemplateListWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineTemplateListWithDefaults() *HarvesterhciIoV1beta1VirtualMachineTemplateList {
	this := HarvesterhciIoV1beta1VirtualMachineTemplateList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetItems returns the Items field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetItems() []HarvesterhciIoV1beta1VirtualMachineTemplate {
	if o == nil {
		var ret []HarvesterhciIoV1beta1VirtualMachineTemplate
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetItemsOk() ([]HarvesterhciIoV1beta1VirtualMachineTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) SetItems(v []HarvesterhciIoV1beta1VirtualMachineTemplate) {
	o.Items = v
}

// GetKind returns the Kind field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o HarvesterhciIoV1beta1VirtualMachineTemplateList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineTemplateList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["items"] = o.Items
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1VirtualMachineTemplateList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"items",
		"kind",
		"metadata",
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

	varHarvesterhciIoV1beta1VirtualMachineTemplateList := _HarvesterhciIoV1beta1VirtualMachineTemplateList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1VirtualMachineTemplateList)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1VirtualMachineTemplateList(varHarvesterhciIoV1beta1VirtualMachineTemplateList)

	return err
}

type NullableHarvesterhciIoV1beta1VirtualMachineTemplateList struct {
	value *HarvesterhciIoV1beta1VirtualMachineTemplateList
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) Get() *HarvesterhciIoV1beta1VirtualMachineTemplateList {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) Set(val *HarvesterhciIoV1beta1VirtualMachineTemplateList) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineTemplateList(val *HarvesterhciIoV1beta1VirtualMachineTemplateList) *NullableHarvesterhciIoV1beta1VirtualMachineTemplateList {
	return &NullableHarvesterhciIoV1beta1VirtualMachineTemplateList{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplateList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


