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

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceMigration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceMigration{}

// KubevirtIoApiCoreV1VirtualMachineInstanceMigration VirtualMachineInstanceMigration represents the object tracking a VMI's migration to another host in the cluster
type KubevirtIoApiCoreV1VirtualMachineInstanceMigration struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec `json:"spec"`
	Status *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus `json:"status,omitempty"`
}

type _KubevirtIoApiCoreV1VirtualMachineInstanceMigration KubevirtIoApiCoreV1VirtualMachineInstanceMigration

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigration instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigration(spec KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) *KubevirtIoApiCoreV1VirtualMachineInstanceMigration {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigration{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	this.Spec = spec
	var status KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus
	this.Status = &status
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigration {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigration{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var spec KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec
	this.Spec = spec
	var status KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus
	this.Status = &status
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetSpec() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec {
	if o == nil {
		var ret KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetSpec(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetStatus() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus {
	if o == nil || IsNil(o.Status) {
		var ret KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetStatusOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus and assigns it to the Status field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetStatus(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) {
	o.Status = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigration) ToMap() (map[string]interface{}, error) {
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

func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) UnmarshalJSON(data []byte) (err error) {
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

	varKubevirtIoApiCoreV1VirtualMachineInstanceMigration := _KubevirtIoApiCoreV1VirtualMachineInstanceMigration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineInstanceMigration)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineInstanceMigration(varKubevirtIoApiCoreV1VirtualMachineInstanceMigration)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceMigration
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceMigration {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


