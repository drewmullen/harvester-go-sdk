/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
)

// checks if the HarvesterhciIoV1beta1VirtualMachineSourceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineSourceSpec{}

// HarvesterhciIoV1beta1VirtualMachineSourceSpec struct for HarvesterhciIoV1beta1VirtualMachineSourceSpec
type HarvesterhciIoV1beta1VirtualMachineSourceSpec struct {
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec *KubevirtIoApiCoreV1VirtualMachineSpec `json:"spec,omitempty"`
}

// NewHarvesterhciIoV1beta1VirtualMachineSourceSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineSourceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineSourceSpec() *HarvesterhciIoV1beta1VirtualMachineSourceSpec {
	this := HarvesterhciIoV1beta1VirtualMachineSourceSpec{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec KubevirtIoApiCoreV1VirtualMachineSpec
	this.Spec = &spec
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineSourceSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineSourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineSourceSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineSourceSpec {
	this := HarvesterhciIoV1beta1VirtualMachineSourceSpec{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec KubevirtIoApiCoreV1VirtualMachineSpec
	this.Spec = &spec
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetSpec() KubevirtIoApiCoreV1VirtualMachineSpec {
	if o == nil || IsNil(o.Spec) {
		var ret KubevirtIoApiCoreV1VirtualMachineSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given KubevirtIoApiCoreV1VirtualMachineSpec and assigns it to the Spec field.
func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) SetSpec(v KubevirtIoApiCoreV1VirtualMachineSpec) {
	o.Spec = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineSourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineSourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec struct {
	value *HarvesterhciIoV1beta1VirtualMachineSourceSpec
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) Get() *HarvesterhciIoV1beta1VirtualMachineSourceSpec {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) Set(val *HarvesterhciIoV1beta1VirtualMachineSourceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineSourceSpec(val *HarvesterhciIoV1beta1VirtualMachineSourceSpec) *NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec {
	return &NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineSourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


