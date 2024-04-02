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

// checks if the HarvesterhciIoV1beta1VirtualMachineImageList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineImageList{}

// HarvesterhciIoV1beta1VirtualMachineImageList struct for HarvesterhciIoV1beta1VirtualMachineImageList
type HarvesterhciIoV1beta1VirtualMachineImageList struct {
	ApiVersion string `json:"apiVersion"`
	Items []HarvesterhciIoV1beta1VirtualMachineImage `json:"items"`
	Kind string `json:"kind"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _HarvesterhciIoV1beta1VirtualMachineImageList HarvesterhciIoV1beta1VirtualMachineImageList

// NewHarvesterhciIoV1beta1VirtualMachineImageList instantiates a new HarvesterhciIoV1beta1VirtualMachineImageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineImageList(apiVersion string, items []HarvesterhciIoV1beta1VirtualMachineImage, kind string, metadata K8sIoV1ListMeta) *HarvesterhciIoV1beta1VirtualMachineImageList {
	this := HarvesterhciIoV1beta1VirtualMachineImageList{}
	this.ApiVersion = apiVersion
	this.Items = items
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineImageListWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineImageListWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImageList {
	this := HarvesterhciIoV1beta1VirtualMachineImageList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetItems returns the Items field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetItems() []HarvesterhciIoV1beta1VirtualMachineImage {
	if o == nil {
		var ret []HarvesterhciIoV1beta1VirtualMachineImage
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetItemsOk() ([]HarvesterhciIoV1beta1VirtualMachineImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) SetItems(v []HarvesterhciIoV1beta1VirtualMachineImage) {
	o.Items = v
}

// GetKind returns the Kind field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineImageList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o HarvesterhciIoV1beta1VirtualMachineImageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineImageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["items"] = o.Items
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1VirtualMachineImageList) UnmarshalJSON(data []byte) (err error) {
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

	varHarvesterhciIoV1beta1VirtualMachineImageList := _HarvesterhciIoV1beta1VirtualMachineImageList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1VirtualMachineImageList)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1VirtualMachineImageList(varHarvesterhciIoV1beta1VirtualMachineImageList)

	return err
}

type NullableHarvesterhciIoV1beta1VirtualMachineImageList struct {
	value *HarvesterhciIoV1beta1VirtualMachineImageList
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageList) Get() *HarvesterhciIoV1beta1VirtualMachineImageList {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageList) Set(val *HarvesterhciIoV1beta1VirtualMachineImageList) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageList) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineImageList(val *HarvesterhciIoV1beta1VirtualMachineImageList) *NullableHarvesterhciIoV1beta1VirtualMachineImageList {
	return &NullableHarvesterhciIoV1beta1VirtualMachineImageList{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


