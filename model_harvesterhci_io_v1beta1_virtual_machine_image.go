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

// checks if the HarvesterhciIoV1beta1VirtualMachineImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineImage{}

// HarvesterhciIoV1beta1VirtualMachineImage struct for HarvesterhciIoV1beta1VirtualMachineImage
type HarvesterhciIoV1beta1VirtualMachineImage struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec HarvesterhciIoV1beta1VirtualMachineImageSpec `json:"spec"`
	Status *HarvesterhciIoV1beta1VirtualMachineImageStatus `json:"status,omitempty"`
}

type _HarvesterhciIoV1beta1VirtualMachineImage HarvesterhciIoV1beta1VirtualMachineImage

// NewHarvesterhciIoV1beta1VirtualMachineImage instantiates a new HarvesterhciIoV1beta1VirtualMachineImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineImage(spec HarvesterhciIoV1beta1VirtualMachineImageSpec) *HarvesterhciIoV1beta1VirtualMachineImage {
	this := HarvesterhciIoV1beta1VirtualMachineImage{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	this.Spec = spec
	var status HarvesterhciIoV1beta1VirtualMachineImageStatus
	this.Status = &status
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineImageWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineImageWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImage {
	this := HarvesterhciIoV1beta1VirtualMachineImage{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec HarvesterhciIoV1beta1VirtualMachineImageSpec
	this.Spec = spec
	var status HarvesterhciIoV1beta1VirtualMachineImageStatus
	this.Status = &status
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetSpec() HarvesterhciIoV1beta1VirtualMachineImageSpec {
	if o == nil {
		var ret HarvesterhciIoV1beta1VirtualMachineImageSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineImageSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetSpec(v HarvesterhciIoV1beta1VirtualMachineImageSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetStatus() HarvesterhciIoV1beta1VirtualMachineImageStatus {
	if o == nil || IsNil(o.Status) {
		var ret HarvesterhciIoV1beta1VirtualMachineImageStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineImageStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given HarvesterhciIoV1beta1VirtualMachineImageStatus and assigns it to the Status field.
func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetStatus(v HarvesterhciIoV1beta1VirtualMachineImageStatus) {
	o.Status = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["spec"] = o.Spec
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *HarvesterhciIoV1beta1VirtualMachineImage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"spec",
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

	varHarvesterhciIoV1beta1VirtualMachineImage := _HarvesterhciIoV1beta1VirtualMachineImage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1VirtualMachineImage)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1VirtualMachineImage(varHarvesterhciIoV1beta1VirtualMachineImage)

	return err
}

type NullableHarvesterhciIoV1beta1VirtualMachineImage struct {
	value *HarvesterhciIoV1beta1VirtualMachineImage
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImage) Get() *HarvesterhciIoV1beta1VirtualMachineImage {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImage) Set(val *HarvesterhciIoV1beta1VirtualMachineImage) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImage) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineImage(val *HarvesterhciIoV1beta1VirtualMachineImage) *NullableHarvesterhciIoV1beta1VirtualMachineImage {
	return &NullableHarvesterhciIoV1beta1VirtualMachineImage{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


