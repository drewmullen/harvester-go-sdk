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

// checks if the K8sIoV1NodeSelectorTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1NodeSelectorTerm{}

// K8sIoV1NodeSelectorTerm struct for K8sIoV1NodeSelectorTerm
type K8sIoV1NodeSelectorTerm struct {
	MatchExpressions []K8sIoV1NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchFields []K8sIoV1NodeSelectorRequirement `json:"matchFields,omitempty"`
}

// NewK8sIoV1NodeSelectorTerm instantiates a new K8sIoV1NodeSelectorTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1NodeSelectorTerm() *K8sIoV1NodeSelectorTerm {
	this := K8sIoV1NodeSelectorTerm{}
	return &this
}

// NewK8sIoV1NodeSelectorTermWithDefaults instantiates a new K8sIoV1NodeSelectorTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1NodeSelectorTermWithDefaults() *K8sIoV1NodeSelectorTerm {
	this := K8sIoV1NodeSelectorTerm{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *K8sIoV1NodeSelectorTerm) GetMatchExpressions() []K8sIoV1NodeSelectorRequirement {
	if o == nil || IsNil(o.MatchExpressions) {
		var ret []K8sIoV1NodeSelectorRequirement
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1NodeSelectorTerm) GetMatchExpressionsOk() ([]K8sIoV1NodeSelectorRequirement, bool) {
	if o == nil || IsNil(o.MatchExpressions) {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *K8sIoV1NodeSelectorTerm) HasMatchExpressions() bool {
	if o != nil && !IsNil(o.MatchExpressions) {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []K8sIoV1NodeSelectorRequirement and assigns it to the MatchExpressions field.
func (o *K8sIoV1NodeSelectorTerm) SetMatchExpressions(v []K8sIoV1NodeSelectorRequirement) {
	o.MatchExpressions = v
}

// GetMatchFields returns the MatchFields field value if set, zero value otherwise.
func (o *K8sIoV1NodeSelectorTerm) GetMatchFields() []K8sIoV1NodeSelectorRequirement {
	if o == nil || IsNil(o.MatchFields) {
		var ret []K8sIoV1NodeSelectorRequirement
		return ret
	}
	return o.MatchFields
}

// GetMatchFieldsOk returns a tuple with the MatchFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1NodeSelectorTerm) GetMatchFieldsOk() ([]K8sIoV1NodeSelectorRequirement, bool) {
	if o == nil || IsNil(o.MatchFields) {
		return nil, false
	}
	return o.MatchFields, true
}

// HasMatchFields returns a boolean if a field has been set.
func (o *K8sIoV1NodeSelectorTerm) HasMatchFields() bool {
	if o != nil && !IsNil(o.MatchFields) {
		return true
	}

	return false
}

// SetMatchFields gets a reference to the given []K8sIoV1NodeSelectorRequirement and assigns it to the MatchFields field.
func (o *K8sIoV1NodeSelectorTerm) SetMatchFields(v []K8sIoV1NodeSelectorRequirement) {
	o.MatchFields = v
}

func (o K8sIoV1NodeSelectorTerm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1NodeSelectorTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchExpressions) {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if !IsNil(o.MatchFields) {
		toSerialize["matchFields"] = o.MatchFields
	}
	return toSerialize, nil
}

type NullableK8sIoV1NodeSelectorTerm struct {
	value *K8sIoV1NodeSelectorTerm
	isSet bool
}

func (v NullableK8sIoV1NodeSelectorTerm) Get() *K8sIoV1NodeSelectorTerm {
	return v.value
}

func (v *NullableK8sIoV1NodeSelectorTerm) Set(val *K8sIoV1NodeSelectorTerm) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1NodeSelectorTerm) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1NodeSelectorTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1NodeSelectorTerm(val *K8sIoV1NodeSelectorTerm) *NullableK8sIoV1NodeSelectorTerm {
	return &NullableK8sIoV1NodeSelectorTerm{value: val, isSet: true}
}

func (v NullableK8sIoV1NodeSelectorTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1NodeSelectorTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


