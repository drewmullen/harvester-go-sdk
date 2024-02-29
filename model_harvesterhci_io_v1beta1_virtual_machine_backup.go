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

// checks if the HarvesterhciIoV1beta1VirtualMachineBackup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineBackup{}

// HarvesterhciIoV1beta1VirtualMachineBackup struct for HarvesterhciIoV1beta1VirtualMachineBackup
type HarvesterhciIoV1beta1VirtualMachineBackup struct {
	ApiVersion string `json:"apiVersion"`
	Kind string `json:"kind"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Spec HarvesterhciIoV1beta1VirtualMachineBackupSpec `json:"spec"`
	Status *HarvesterhciIoV1beta1VirtualMachineBackupStatus `json:"status,omitempty"`
}

type _HarvesterhciIoV1beta1VirtualMachineBackup HarvesterhciIoV1beta1VirtualMachineBackup

// NewHarvesterhciIoV1beta1VirtualMachineBackup instantiates a new HarvesterhciIoV1beta1VirtualMachineBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineBackup(apiVersion string, kind string, spec HarvesterhciIoV1beta1VirtualMachineBackupSpec) *HarvesterhciIoV1beta1VirtualMachineBackup {
	this := HarvesterhciIoV1beta1VirtualMachineBackup{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Spec = spec
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineBackupWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineBackupWithDefaults() *HarvesterhciIoV1beta1VirtualMachineBackup {
	this := HarvesterhciIoV1beta1VirtualMachineBackup{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetSpec() HarvesterhciIoV1beta1VirtualMachineBackupSpec {
	if o == nil {
		var ret HarvesterhciIoV1beta1VirtualMachineBackupSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineBackupSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetSpec(v HarvesterhciIoV1beta1VirtualMachineBackupSpec) {
	o.Spec = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetStatus() HarvesterhciIoV1beta1VirtualMachineBackupStatus {
	if o == nil || IsNil(o.Status) {
		var ret HarvesterhciIoV1beta1VirtualMachineBackupStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineBackupStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given HarvesterhciIoV1beta1VirtualMachineBackupStatus and assigns it to the Status field.
func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetStatus(v HarvesterhciIoV1beta1VirtualMachineBackupStatus) {
	o.Status = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineBackup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineBackup) ToMap() (map[string]interface{}, error) {
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

func (o *HarvesterhciIoV1beta1VirtualMachineBackup) UnmarshalJSON(data []byte) (err error) {
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

	varHarvesterhciIoV1beta1VirtualMachineBackup := _HarvesterhciIoV1beta1VirtualMachineBackup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHarvesterhciIoV1beta1VirtualMachineBackup)

	if err != nil {
		return err
	}

	*o = HarvesterhciIoV1beta1VirtualMachineBackup(varHarvesterhciIoV1beta1VirtualMachineBackup)

	return err
}

type NullableHarvesterhciIoV1beta1VirtualMachineBackup struct {
	value *HarvesterhciIoV1beta1VirtualMachineBackup
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackup) Get() *HarvesterhciIoV1beta1VirtualMachineBackup {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackup) Set(val *HarvesterhciIoV1beta1VirtualMachineBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineBackup(val *HarvesterhciIoV1beta1VirtualMachineBackup) *NullableHarvesterhciIoV1beta1VirtualMachineBackup {
	return &NullableHarvesterhciIoV1beta1VirtualMachineBackup{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


