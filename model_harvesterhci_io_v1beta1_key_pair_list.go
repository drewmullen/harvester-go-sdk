/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HarvesterhciIoV1beta1KeyPairList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1KeyPairList{}

// HarvesterhciIoV1beta1KeyPairList struct for HarvesterhciIoV1beta1KeyPairList
type HarvesterhciIoV1beta1KeyPairList struct {
	ApiVersion string `json:"apiVersion"`
	Items []HarvesterhciIoV1beta1KeyPair `json:"items"`
	Kind string `json:"kind"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _HarvesterhciIoV1beta1KeyPairList HarvesterhciIoV1beta1KeyPairList

// NewHarvesterhciIoV1beta1KeyPairList instantiates a new HarvesterhciIoV1beta1KeyPairList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1KeyPairList(apiVersion string, items []HarvesterhciIoV1beta1KeyPair, kind string, metadata K8sIoV1ListMeta) *HarvesterhciIoV1beta1KeyPairList {
	this := HarvesterhciIoV1beta1KeyPairList{}
	this.ApiVersion = apiVersion
	this.Items = items
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewHarvesterhciIoV1beta1KeyPairListWithDefaults instantiates a new HarvesterhciIoV1beta1KeyPairList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1KeyPairListWithDefaults() *HarvesterhciIoV1beta1KeyPairList {
	this := HarvesterhciIoV1beta1KeyPairList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *HarvesterhciIoV1beta1KeyPairList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1KeyPairList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *HarvesterhciIoV1beta1KeyPairList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetItems returns the Items field value
func (o *HarvesterhciIoV1beta1KeyPairList) GetItems() []HarvesterhciIoV1beta1KeyPair {
	if o == nil {
		var ret []HarvesterhciIoV1beta1KeyPair
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1KeyPairList) GetItemsOk() ([]HarvesterhciIoV1beta1KeyPair, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *HarvesterhciIoV1beta1KeyPairList) SetItems(v []HarvesterhciIoV1beta1KeyPair) {
	o.Items = v
}

// GetKind returns the Kind field value
func (o *HarvesterhciIoV1beta1KeyPairList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1KeyPairList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *HarvesterhciIoV1beta1KeyPairList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *HarvesterhciIoV1beta1KeyPairList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1KeyPairList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *HarvesterhciIoV1beta1KeyPairList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o HarvesterhciIoV1beta1KeyPairList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1KeyPairList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["items"] = o.Items
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1KeyPairList) UnmarshalJSON(data []byte) (err error) {
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

	varHarvesterhciIoV1beta1KeyPairList := _HarvesterhciIoV1beta1KeyPairList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1KeyPairList)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1KeyPairList(varHarvesterhciIoV1beta1KeyPairList)

	return err
}

type NullableHarvesterhciIoV1beta1KeyPairList struct {
	value *HarvesterhciIoV1beta1KeyPairList
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1KeyPairList) Get() *HarvesterhciIoV1beta1KeyPairList {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1KeyPairList) Set(val *HarvesterhciIoV1beta1KeyPairList) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1KeyPairList) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1KeyPairList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1KeyPairList(val *HarvesterhciIoV1beta1KeyPairList) *NullableHarvesterhciIoV1beta1KeyPairList {
	return &NullableHarvesterhciIoV1beta1KeyPairList{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1KeyPairList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1KeyPairList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


