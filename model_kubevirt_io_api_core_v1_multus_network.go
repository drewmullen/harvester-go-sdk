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

// checks if the KubevirtIoApiCoreV1MultusNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1MultusNetwork{}

// KubevirtIoApiCoreV1MultusNetwork Represents the multus cni network.
type KubevirtIoApiCoreV1MultusNetwork struct {
	// Select the default network and add it to the multus-cni.io/default-network annotation.
	Default *bool `json:"default,omitempty"`
	// References to a NetworkAttachmentDefinition CRD object. Format: <networkName>, <namespace>/<networkName>. If namespace is not specified, VMI namespace is assumed.
	NetworkName string `json:"networkName"`
}

type _KubevirtIoApiCoreV1MultusNetwork KubevirtIoApiCoreV1MultusNetwork

// NewKubevirtIoApiCoreV1MultusNetwork instantiates a new KubevirtIoApiCoreV1MultusNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1MultusNetwork(networkName string) *KubevirtIoApiCoreV1MultusNetwork {
	this := KubevirtIoApiCoreV1MultusNetwork{}
	this.NetworkName = networkName
	return &this
}

// NewKubevirtIoApiCoreV1MultusNetworkWithDefaults instantiates a new KubevirtIoApiCoreV1MultusNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1MultusNetworkWithDefaults() *KubevirtIoApiCoreV1MultusNetwork {
	this := KubevirtIoApiCoreV1MultusNetwork{}
	var networkName string = ""
	this.NetworkName = networkName
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1MultusNetwork) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1MultusNetwork) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1MultusNetwork) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *KubevirtIoApiCoreV1MultusNetwork) SetDefault(v bool) {
	o.Default = &v
}

// GetNetworkName returns the NetworkName field value
func (o *KubevirtIoApiCoreV1MultusNetwork) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1MultusNetwork) GetNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *KubevirtIoApiCoreV1MultusNetwork) SetNetworkName(v string) {
	o.NetworkName = v
}

func (o KubevirtIoApiCoreV1MultusNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1MultusNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["networkName"] = o.NetworkName
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1MultusNetwork) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"networkName",
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

	varKubevirtIoApiCoreV1MultusNetwork := _KubevirtIoApiCoreV1MultusNetwork{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1MultusNetwork)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1MultusNetwork(varKubevirtIoApiCoreV1MultusNetwork)

	return err
}

type NullableKubevirtIoApiCoreV1MultusNetwork struct {
	value *KubevirtIoApiCoreV1MultusNetwork
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1MultusNetwork) Get() *KubevirtIoApiCoreV1MultusNetwork {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1MultusNetwork) Set(val *KubevirtIoApiCoreV1MultusNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1MultusNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1MultusNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1MultusNetwork(val *KubevirtIoApiCoreV1MultusNetwork) *NullableKubevirtIoApiCoreV1MultusNetwork {
	return &NullableKubevirtIoApiCoreV1MultusNetwork{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1MultusNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1MultusNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


