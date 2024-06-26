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

// checks if the NetworkHarvesterhciIoV1beta1NodeNetworkList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkHarvesterhciIoV1beta1NodeNetworkList{}

// NetworkHarvesterhciIoV1beta1NodeNetworkList NodeNetworkList is a list of NodeNetwork resources
type NetworkHarvesterhciIoV1beta1NodeNetworkList struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []NetworkHarvesterhciIoV1beta1NodeNetwork `json:"items"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _NetworkHarvesterhciIoV1beta1NodeNetworkList NetworkHarvesterhciIoV1beta1NodeNetworkList

// NewNetworkHarvesterhciIoV1beta1NodeNetworkList instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHarvesterhciIoV1beta1NodeNetworkList(items []NetworkHarvesterhciIoV1beta1NodeNetwork, metadata K8sIoV1ListMeta) *NetworkHarvesterhciIoV1beta1NodeNetworkList {
	this := NetworkHarvesterhciIoV1beta1NodeNetworkList{}
	this.Items = items
	this.Metadata = metadata
	return &this
}

// NewNetworkHarvesterhciIoV1beta1NodeNetworkListWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHarvesterhciIoV1beta1NodeNetworkListWithDefaults() *NetworkHarvesterhciIoV1beta1NodeNetworkList {
	this := NetworkHarvesterhciIoV1beta1NodeNetworkList{}
	var metadata K8sIoV1ListMeta
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetItems() []NetworkHarvesterhciIoV1beta1NodeNetwork {
	if o == nil {
		var ret []NetworkHarvesterhciIoV1beta1NodeNetwork
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetItemsOk() ([]NetworkHarvesterhciIoV1beta1NodeNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetItems(v []NetworkHarvesterhciIoV1beta1NodeNetwork) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o NetworkHarvesterhciIoV1beta1NodeNetworkList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkHarvesterhciIoV1beta1NodeNetworkList) ToMap() (map[string]interface{}, error) {
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

func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) UnmarshalJSON(data []byte) (err error) {
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

	varNetworkHarvesterhciIoV1beta1NodeNetworkList := _NetworkHarvesterhciIoV1beta1NodeNetworkList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNetworkHarvesterhciIoV1beta1NodeNetworkList)

	if err != nil {
		return err
	}

	*o = NetworkHarvesterhciIoV1beta1NodeNetworkList(varNetworkHarvesterhciIoV1beta1NodeNetworkList)

	return err
}

type NullableNetworkHarvesterhciIoV1beta1NodeNetworkList struct {
	value *NetworkHarvesterhciIoV1beta1NodeNetworkList
	isSet bool
}

func (v NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) Get() *NetworkHarvesterhciIoV1beta1NodeNetworkList {
	return v.value
}

func (v *NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) Set(val *NetworkHarvesterhciIoV1beta1NodeNetworkList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHarvesterhciIoV1beta1NodeNetworkList(val *NetworkHarvesterhciIoV1beta1NodeNetworkList) *NullableNetworkHarvesterhciIoV1beta1NodeNetworkList {
	return &NullableNetworkHarvesterhciIoV1beta1NodeNetworkList{value: val, isSet: true}
}

func (v NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHarvesterhciIoV1beta1NodeNetworkList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


