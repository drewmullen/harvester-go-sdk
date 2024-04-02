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

// checks if the KubevirtIoApiCoreV1VirtualMachineInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstance{}

// KubevirtIoApiCoreV1VirtualMachineInstance struct for KubevirtIoApiCoreV1VirtualMachineInstance
type KubevirtIoApiCoreV1VirtualMachineInstance struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec KubevirtIoApiCoreV1VirtualMachineInstanceSpec `json:"spec"`
	Status *KubevirtIoApiCoreV1VirtualMachineInstanceStatus `json:"status,omitempty"`
}

type _KubevirtIoApiCoreV1VirtualMachineInstance KubevirtIoApiCoreV1VirtualMachineInstance

// NewKubevirtIoApiCoreV1VirtualMachineInstance instantiates a new KubevirtIoApiCoreV1VirtualMachineInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstance(apiVersion string, kind string, spec KubevirtIoApiCoreV1VirtualMachineInstanceSpec) *KubevirtIoApiCoreV1VirtualMachineInstance {
	this := KubevirtIoApiCoreV1VirtualMachineInstance{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	var metadata K8sIoV1ObjectMeta = {}
	this.Metadata = &metadata
	this.Spec = spec
	var status KubevirtIoApiCoreV1VirtualMachineInstanceStatus = {}
	this.Status = &status
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstance {
	this := KubevirtIoApiCoreV1VirtualMachineInstance{}
	var metadata K8sIoV1ObjectMeta = {}
	this.Metadata = &metadata
	var spec KubevirtIoApiCoreV1VirtualMachineInstanceSpec = {}
	this.Spec = spec
	var status KubevirtIoApiCoreV1VirtualMachineInstanceStatus = {}
	this.Status = &status
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetSpec() KubevirtIoApiCoreV1VirtualMachineInstanceSpec {
	if o == nil {
		var ret KubevirtIoApiCoreV1VirtualMachineInstanceSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetSpec(v KubevirtIoApiCoreV1VirtualMachineInstanceSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetStatus() KubevirtIoApiCoreV1VirtualMachineInstanceStatus {
	if o == nil || IsNil(o.Status) {
		var ret KubevirtIoApiCoreV1VirtualMachineInstanceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetStatusOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given KubevirtIoApiCoreV1VirtualMachineInstanceStatus and assigns it to the Status field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetStatus(v KubevirtIoApiCoreV1VirtualMachineInstanceStatus) {
	o.Status = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["spec"] = o.Spec
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1VirtualMachineInstance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
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

	varKubevirtIoApiCoreV1VirtualMachineInstance := _KubevirtIoApiCoreV1VirtualMachineInstance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineInstance)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineInstance(varKubevirtIoApiCoreV1VirtualMachineInstance)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstance struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstance
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstance) Get() *KubevirtIoApiCoreV1VirtualMachineInstance {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstance) Set(val *KubevirtIoApiCoreV1VirtualMachineInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstance(val *KubevirtIoApiCoreV1VirtualMachineInstance) *NullableKubevirtIoApiCoreV1VirtualMachineInstance {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstance{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


