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

// checks if the HarvesterhciIoV1beta1UpgradeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1UpgradeList{}

// HarvesterhciIoV1beta1UpgradeList struct for HarvesterhciIoV1beta1UpgradeList
type HarvesterhciIoV1beta1UpgradeList struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []HarvesterhciIoV1beta1Upgrade `json:"items"`
	Kind *string `json:"kind,omitempty"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _HarvesterhciIoV1beta1UpgradeList HarvesterhciIoV1beta1UpgradeList

// NewHarvesterhciIoV1beta1UpgradeList instantiates a new HarvesterhciIoV1beta1UpgradeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1UpgradeList(items []HarvesterhciIoV1beta1Upgrade, metadata K8sIoV1ListMeta) *HarvesterhciIoV1beta1UpgradeList {
	this := HarvesterhciIoV1beta1UpgradeList{}
	this.Items = items
	this.Metadata = metadata
	return &this
}

// NewHarvesterhciIoV1beta1UpgradeListWithDefaults instantiates a new HarvesterhciIoV1beta1UpgradeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1UpgradeListWithDefaults() *HarvesterhciIoV1beta1UpgradeList {
	this := HarvesterhciIoV1beta1UpgradeList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *HarvesterhciIoV1beta1UpgradeList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *HarvesterhciIoV1beta1UpgradeList) GetItems() []HarvesterhciIoV1beta1Upgrade {
	if o == nil {
		var ret []HarvesterhciIoV1beta1Upgrade
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) GetItemsOk() ([]HarvesterhciIoV1beta1Upgrade, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *HarvesterhciIoV1beta1UpgradeList) SetItems(v []HarvesterhciIoV1beta1Upgrade) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *HarvesterhciIoV1beta1UpgradeList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *HarvesterhciIoV1beta1UpgradeList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *HarvesterhciIoV1beta1UpgradeList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o HarvesterhciIoV1beta1UpgradeList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1UpgradeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["items"] = o.Items
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1UpgradeList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
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

	varHarvesterhciIoV1beta1UpgradeList := _HarvesterhciIoV1beta1UpgradeList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1UpgradeList)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1UpgradeList(varHarvesterhciIoV1beta1UpgradeList)

	return err
}

type NullableHarvesterhciIoV1beta1UpgradeList struct {
	value *HarvesterhciIoV1beta1UpgradeList
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1UpgradeList) Get() *HarvesterhciIoV1beta1UpgradeList {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1UpgradeList) Set(val *HarvesterhciIoV1beta1UpgradeList) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1UpgradeList) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1UpgradeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1UpgradeList(val *HarvesterhciIoV1beta1UpgradeList) *NullableHarvesterhciIoV1beta1UpgradeList {
	return &NullableHarvesterhciIoV1beta1UpgradeList{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1UpgradeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1UpgradeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


