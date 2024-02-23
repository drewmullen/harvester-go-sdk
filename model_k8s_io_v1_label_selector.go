/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the K8sIoV1LabelSelector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1LabelSelector{}

// K8sIoV1LabelSelector struct for K8sIoV1LabelSelector
type K8sIoV1LabelSelector struct {
	MatchExpressions []K8sIoV1LabelSelectorRequirement `json:"matchExpressions,omitempty"`
	MatchLabels *map[string]string `json:"matchLabels,omitempty"`
}

// NewK8sIoV1LabelSelector instantiates a new K8sIoV1LabelSelector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1LabelSelector() *K8sIoV1LabelSelector {
	this := K8sIoV1LabelSelector{}
	return &this
}

// NewK8sIoV1LabelSelectorWithDefaults instantiates a new K8sIoV1LabelSelector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1LabelSelectorWithDefaults() *K8sIoV1LabelSelector {
	this := K8sIoV1LabelSelector{}
	return &this
}

// GetMatchExpressions returns the MatchExpressions field value if set, zero value otherwise.
func (o *K8sIoV1LabelSelector) GetMatchExpressions() []K8sIoV1LabelSelectorRequirement {
	if o == nil || IsNil(o.MatchExpressions) {
		var ret []K8sIoV1LabelSelectorRequirement
		return ret
	}
	return o.MatchExpressions
}

// GetMatchExpressionsOk returns a tuple with the MatchExpressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1LabelSelector) GetMatchExpressionsOk() ([]K8sIoV1LabelSelectorRequirement, bool) {
	if o == nil || IsNil(o.MatchExpressions) {
		return nil, false
	}
	return o.MatchExpressions, true
}

// HasMatchExpressions returns a boolean if a field has been set.
func (o *K8sIoV1LabelSelector) HasMatchExpressions() bool {
	if o != nil && !IsNil(o.MatchExpressions) {
		return true
	}

	return false
}

// SetMatchExpressions gets a reference to the given []K8sIoV1LabelSelectorRequirement and assigns it to the MatchExpressions field.
func (o *K8sIoV1LabelSelector) SetMatchExpressions(v []K8sIoV1LabelSelectorRequirement) {
	o.MatchExpressions = v
}

// GetMatchLabels returns the MatchLabels field value if set, zero value otherwise.
func (o *K8sIoV1LabelSelector) GetMatchLabels() map[string]string {
	if o == nil || IsNil(o.MatchLabels) {
		var ret map[string]string
		return ret
	}
	return *o.MatchLabels
}

// GetMatchLabelsOk returns a tuple with the MatchLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1LabelSelector) GetMatchLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MatchLabels) {
		return nil, false
	}
	return o.MatchLabels, true
}

// HasMatchLabels returns a boolean if a field has been set.
func (o *K8sIoV1LabelSelector) HasMatchLabels() bool {
	if o != nil && !IsNil(o.MatchLabels) {
		return true
	}

	return false
}

// SetMatchLabels gets a reference to the given map[string]string and assigns it to the MatchLabels field.
func (o *K8sIoV1LabelSelector) SetMatchLabels(v map[string]string) {
	o.MatchLabels = &v
}

func (o K8sIoV1LabelSelector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1LabelSelector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchExpressions) {
		toSerialize["matchExpressions"] = o.MatchExpressions
	}
	if !IsNil(o.MatchLabels) {
		toSerialize["matchLabels"] = o.MatchLabels
	}
	return toSerialize, nil
}

type NullableK8sIoV1LabelSelector struct {
	value *K8sIoV1LabelSelector
	isSet bool
}

func (v NullableK8sIoV1LabelSelector) Get() *K8sIoV1LabelSelector {
	return v.value
}

func (v *NullableK8sIoV1LabelSelector) Set(val *K8sIoV1LabelSelector) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1LabelSelector) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1LabelSelector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1LabelSelector(val *K8sIoV1LabelSelector) *NullableK8sIoV1LabelSelector {
	return &NullableK8sIoV1LabelSelector{value: val, isSet: true}
}

func (v NullableK8sIoV1LabelSelector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1LabelSelector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


