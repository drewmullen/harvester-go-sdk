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

// checks if the HarvesterhciIoV1beta1VirtualMachineTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineTemplate{}

// HarvesterhciIoV1beta1VirtualMachineTemplate struct for HarvesterhciIoV1beta1VirtualMachineTemplate
type HarvesterhciIoV1beta1VirtualMachineTemplate struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec *HarvesterhciIoV1beta1VirtualMachineTemplateSpec `json:"spec,omitempty"`
	Status *HarvesterhciIoV1beta1VirtualMachineTemplateStatus `json:"status,omitempty"`
}

// NewHarvesterhciIoV1beta1VirtualMachineTemplate instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineTemplate() *HarvesterhciIoV1beta1VirtualMachineTemplate {
	this := HarvesterhciIoV1beta1VirtualMachineTemplate{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec HarvesterhciIoV1beta1VirtualMachineTemplateSpec
	this.Spec = &spec
	var status HarvesterhciIoV1beta1VirtualMachineTemplateStatus
	this.Status = &status
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineTemplateWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineTemplateWithDefaults() *HarvesterhciIoV1beta1VirtualMachineTemplate {
	this := HarvesterhciIoV1beta1VirtualMachineTemplate{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec HarvesterhciIoV1beta1VirtualMachineTemplateSpec
	this.Spec = &spec
	var status HarvesterhciIoV1beta1VirtualMachineTemplateStatus
	this.Status = &status
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetSpec() HarvesterhciIoV1beta1VirtualMachineTemplateSpec {
	if o == nil || IsNil(o.Spec) {
		var ret HarvesterhciIoV1beta1VirtualMachineTemplateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineTemplateSpec, bool) {
	if o == nil || IsNil(o.Spec) {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) HasSpec() bool {
	if o != nil && !IsNil(o.Spec) {
		return true
	}

	return false
}

// SetSpec gets a reference to the given HarvesterhciIoV1beta1VirtualMachineTemplateSpec and assigns it to the Spec field.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) SetSpec(v HarvesterhciIoV1beta1VirtualMachineTemplateSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetStatus() HarvesterhciIoV1beta1VirtualMachineTemplateStatus {
	if o == nil || IsNil(o.Status) {
		var ret HarvesterhciIoV1beta1VirtualMachineTemplateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineTemplateStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given HarvesterhciIoV1beta1VirtualMachineTemplateStatus and assigns it to the Status field.
func (o *HarvesterhciIoV1beta1VirtualMachineTemplate) SetStatus(v HarvesterhciIoV1beta1VirtualMachineTemplateStatus) {
	o.Status = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineTemplate) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Spec) {
		toSerialize["spec"] = o.Spec
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1VirtualMachineTemplate struct {
	value *HarvesterhciIoV1beta1VirtualMachineTemplate
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplate) Get() *HarvesterhciIoV1beta1VirtualMachineTemplate {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplate) Set(val *HarvesterhciIoV1beta1VirtualMachineTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineTemplate(val *HarvesterhciIoV1beta1VirtualMachineTemplate) *NullableHarvesterhciIoV1beta1VirtualMachineTemplate {
	return &NullableHarvesterhciIoV1beta1VirtualMachineTemplate{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


