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

// checks if the HarvesterhciIoV1beta1VirtualMachineRestore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineRestore{}

// HarvesterhciIoV1beta1VirtualMachineRestore struct for HarvesterhciIoV1beta1VirtualMachineRestore
type HarvesterhciIoV1beta1VirtualMachineRestore struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec HarvesterhciIoV1beta1VirtualMachineRestoreSpec `json:"spec"`
	Status *HarvesterhciIoV1beta1VirtualMachineRestoreStatus `json:"status,omitempty"`
}

type _HarvesterhciIoV1beta1VirtualMachineRestore HarvesterhciIoV1beta1VirtualMachineRestore

// NewHarvesterhciIoV1beta1VirtualMachineRestore instantiates a new HarvesterhciIoV1beta1VirtualMachineRestore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineRestore(spec HarvesterhciIoV1beta1VirtualMachineRestoreSpec) *HarvesterhciIoV1beta1VirtualMachineRestore {
	this := HarvesterhciIoV1beta1VirtualMachineRestore{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	this.Spec = spec
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineRestoreWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineRestore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineRestoreWithDefaults() *HarvesterhciIoV1beta1VirtualMachineRestore {
	this := HarvesterhciIoV1beta1VirtualMachineRestore{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec HarvesterhciIoV1beta1VirtualMachineRestoreSpec
	this.Spec = spec
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetSpec() HarvesterhciIoV1beta1VirtualMachineRestoreSpec {
	if o == nil {
		var ret HarvesterhciIoV1beta1VirtualMachineRestoreSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineRestoreSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetSpec(v HarvesterhciIoV1beta1VirtualMachineRestoreSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetStatus() HarvesterhciIoV1beta1VirtualMachineRestoreStatus {
	if o == nil || IsNil(o.Status) {
		var ret HarvesterhciIoV1beta1VirtualMachineRestoreStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineRestoreStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given HarvesterhciIoV1beta1VirtualMachineRestoreStatus and assigns it to the Status field.
func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetStatus(v HarvesterhciIoV1beta1VirtualMachineRestoreStatus) {
	o.Status = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineRestore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineRestore) ToMap() (map[string]interface{}, error) {
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

func (o *HarvesterhciIoV1beta1VirtualMachineRestore) UnmarshalJSON(data []byte) (err error) {
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

	varHarvesterhciIoV1beta1VirtualMachineRestore := _HarvesterhciIoV1beta1VirtualMachineRestore{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1VirtualMachineRestore)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1VirtualMachineRestore(varHarvesterhciIoV1beta1VirtualMachineRestore)

	return err
}

type NullableHarvesterhciIoV1beta1VirtualMachineRestore struct {
	value *HarvesterhciIoV1beta1VirtualMachineRestore
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestore) Get() *HarvesterhciIoV1beta1VirtualMachineRestore {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestore) Set(val *HarvesterhciIoV1beta1VirtualMachineRestore) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestore) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineRestore(val *HarvesterhciIoV1beta1VirtualMachineRestore) *NullableHarvesterhciIoV1beta1VirtualMachineRestore {
	return &NullableHarvesterhciIoV1beta1VirtualMachineRestore{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineRestore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineRestore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


