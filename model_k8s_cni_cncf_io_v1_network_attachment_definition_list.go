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

// checks if the K8sCniCncfIoV1NetworkAttachmentDefinitionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sCniCncfIoV1NetworkAttachmentDefinitionList{}

// K8sCniCncfIoV1NetworkAttachmentDefinitionList struct for K8sCniCncfIoV1NetworkAttachmentDefinitionList
type K8sCniCncfIoV1NetworkAttachmentDefinitionList struct {
	ApiVersion string `json:"apiVersion"`
	Items []K8sCniCncfIoV1NetworkAttachmentDefinition `json:"items"`
	Kind string `json:"kind"`
	Metadata K8sIoV1ListMeta `json:"metadata"`
}

type _K8sCniCncfIoV1NetworkAttachmentDefinitionList K8sCniCncfIoV1NetworkAttachmentDefinitionList

// NewK8sCniCncfIoV1NetworkAttachmentDefinitionList instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinitionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sCniCncfIoV1NetworkAttachmentDefinitionList(apiVersion string, items []K8sCniCncfIoV1NetworkAttachmentDefinition, kind string, metadata K8sIoV1ListMeta) *K8sCniCncfIoV1NetworkAttachmentDefinitionList {
	this := K8sCniCncfIoV1NetworkAttachmentDefinitionList{}
	this.ApiVersion = apiVersion
	this.Items = items
	this.Kind = kind
	this.Metadata = metadata
	return &this
}

// NewK8sCniCncfIoV1NetworkAttachmentDefinitionListWithDefaults instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinitionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sCniCncfIoV1NetworkAttachmentDefinitionListWithDefaults() *K8sCniCncfIoV1NetworkAttachmentDefinitionList {
	this := K8sCniCncfIoV1NetworkAttachmentDefinitionList{}
	var metadata K8sIoV1ListMeta = {}
	this.Metadata = metadata
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetItems returns the Items field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetItems() []K8sCniCncfIoV1NetworkAttachmentDefinition {
	if o == nil {
		var ret []K8sCniCncfIoV1NetworkAttachmentDefinition
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetItemsOk() ([]K8sCniCncfIoV1NetworkAttachmentDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetItems(v []K8sCniCncfIoV1NetworkAttachmentDefinition) {
	o.Items = v
}

// GetKind returns the Kind field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetMetadata() K8sIoV1ListMeta {
	if o == nil {
		var ret K8sIoV1ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetMetadataOk() (*K8sIoV1ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetMetadata(v K8sIoV1ListMeta) {
	o.Metadata = v
}

func (o K8sCniCncfIoV1NetworkAttachmentDefinitionList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sCniCncfIoV1NetworkAttachmentDefinitionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["items"] = o.Items
	toSerialize["kind"] = o.Kind
	toSerialize["metadata"] = o.Metadata
	return toSerialize, nil
}

func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) UnmarshalJSON(data []byte) (err error) {
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

	varK8sCniCncfIoV1NetworkAttachmentDefinitionList := _K8sCniCncfIoV1NetworkAttachmentDefinitionList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sCniCncfIoV1NetworkAttachmentDefinitionList)

	if err != nil {
		return err
	}

	*o = K8sCniCncfIoV1NetworkAttachmentDefinitionList(varK8sCniCncfIoV1NetworkAttachmentDefinitionList)

	return err
}

type NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList struct {
	value *K8sCniCncfIoV1NetworkAttachmentDefinitionList
	isSet bool
}

func (v NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) Get() *K8sCniCncfIoV1NetworkAttachmentDefinitionList {
	return v.value
}

func (v *NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) Set(val *K8sCniCncfIoV1NetworkAttachmentDefinitionList) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sCniCncfIoV1NetworkAttachmentDefinitionList(val *K8sCniCncfIoV1NetworkAttachmentDefinitionList) *NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList {
	return &NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList{value: val, isSet: true}
}

func (v NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sCniCncfIoV1NetworkAttachmentDefinitionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


