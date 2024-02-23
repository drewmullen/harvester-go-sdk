/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition{}

// KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition struct for KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition
type KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition struct {
	LastProbeTime *string `json:"lastProbeTime,omitempty"`
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Status string `json:"status"`
	Type string `json:"type"`
}

type _KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition(status string, type_ string) *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition{}
	var lastProbeTime string = ""
	this.LastProbeTime = &lastProbeTime
	var lastTransitionTime string = ""
	this.LastTransitionTime = &lastTransitionTime
	this.Status = status
	this.Type = type_
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationConditionWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationConditionWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition{}
	var lastProbeTime string = ""
	this.LastProbeTime = &lastProbeTime
	var lastTransitionTime string = ""
	this.LastTransitionTime = &lastTransitionTime
	var status string = ""
	this.Status = status
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetLastProbeTime returns the LastProbeTime field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastProbeTime() string {
	if o == nil || IsNil(o.LastProbeTime) {
		var ret string
		return ret
	}
	return *o.LastProbeTime
}

// GetLastProbeTimeOk returns a tuple with the LastProbeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastProbeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastProbeTime) {
		return nil, false
	}
	return o.LastProbeTime, true
}

// HasLastProbeTime returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasLastProbeTime() bool {
	if o != nil && !IsNil(o.LastProbeTime) {
		return true
	}

	return false
}

// SetLastProbeTime gets a reference to the given string and assigns it to the LastProbeTime field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetLastProbeTime(v string) {
	o.LastProbeTime = &v
}

// GetLastTransitionTime returns the LastTransitionTime field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastTransitionTime() string {
	if o == nil || IsNil(o.LastTransitionTime) {
		var ret string
		return ret
	}
	return *o.LastTransitionTime
}

// GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastTransitionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastTransitionTime) {
		return nil, false
	}
	return o.LastTransitionTime, true
}

// HasLastTransitionTime returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasLastTransitionTime() bool {
	if o != nil && !IsNil(o.LastTransitionTime) {
		return true
	}

	return false
}

// SetLastTransitionTime gets a reference to the given string and assigns it to the LastTransitionTime field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetLastTransitionTime(v string) {
	o.LastTransitionTime = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetReason(v string) {
	o.Reason = &v
}

// GetStatus returns the Status field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetType(v string) {
	o.Type = v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastProbeTime) {
		toSerialize["lastProbeTime"] = o.LastProbeTime
	}
	if !IsNil(o.LastTransitionTime) {
		toSerialize["lastTransitionTime"] = o.LastTransitionTime
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"type",
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

	varKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition := _KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition(varKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


