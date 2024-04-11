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

// checks if the NetworkHarvesterhciIoV1beta1ClusterNetworkList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkHarvesterhciIoV1beta1ClusterNetworkList{}

// NetworkHarvesterhciIoV1beta1ClusterNetworkList struct for NetworkHarvesterhciIoV1beta1ClusterNetworkList
type NetworkHarvesterhciIoV1beta1ClusterNetworkList struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []NetworkHarvesterhciIoV1beta1ClusterNetwork `json:"items"`
	Kind *string `json:"kind,omitempty"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _NetworkHarvesterhciIoV1beta1ClusterNetworkList NetworkHarvesterhciIoV1beta1ClusterNetworkList

// NewNetworkHarvesterhciIoV1beta1ClusterNetworkList instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHarvesterhciIoV1beta1ClusterNetworkList(items []NetworkHarvesterhciIoV1beta1ClusterNetwork, metadata K8sIoV1ListMeta) *NetworkHarvesterhciIoV1beta1ClusterNetworkList {
	this := NetworkHarvesterhciIoV1beta1ClusterNetworkList{}
	this.Items = items
	this.Metadata = metadata
	return &this
}

// NewNetworkHarvesterhciIoV1beta1ClusterNetworkListWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHarvesterhciIoV1beta1ClusterNetworkListWithDefaults() *NetworkHarvesterhciIoV1beta1ClusterNetworkList {
	this := NetworkHarvesterhciIoV1beta1ClusterNetworkList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetItems() []NetworkHarvesterhciIoV1beta1ClusterNetwork {
	if o == nil {
		var ret []NetworkHarvesterhciIoV1beta1ClusterNetwork
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetItemsOk() ([]NetworkHarvesterhciIoV1beta1ClusterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetItems(v []NetworkHarvesterhciIoV1beta1ClusterNetwork) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o NetworkHarvesterhciIoV1beta1ClusterNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkHarvesterhciIoV1beta1ClusterNetworkList) ToMap() (map[string]interface{}, error) {
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

func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) UnmarshalJSON(data []byte) (err error) {
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

	varNetworkHarvesterhciIoV1beta1ClusterNetworkList := _NetworkHarvesterhciIoV1beta1ClusterNetworkList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkHarvesterhciIoV1beta1ClusterNetworkList)

	if err != nil {
		return err
	}

	*o = NetworkHarvesterhciIoV1beta1ClusterNetworkList(varNetworkHarvesterhciIoV1beta1ClusterNetworkList)

	return err
}

type NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList struct {
	value *NetworkHarvesterhciIoV1beta1ClusterNetworkList
	isSet bool
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) Get() *NetworkHarvesterhciIoV1beta1ClusterNetworkList {
	return v.value
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) Set(val *NetworkHarvesterhciIoV1beta1ClusterNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHarvesterhciIoV1beta1ClusterNetworkList(val *NetworkHarvesterhciIoV1beta1ClusterNetworkList) *NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList {
	return &NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList{value: val, isSet: true}
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


