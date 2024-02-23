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

// checks if the K8sIoV1TopologySpreadConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1TopologySpreadConstraint{}

// K8sIoV1TopologySpreadConstraint struct for K8sIoV1TopologySpreadConstraint
type K8sIoV1TopologySpreadConstraint struct {
	LabelSelector *K8sIoV1LabelSelector `json:"labelSelector,omitempty"`
	MatchLabelKeys []string `json:"matchLabelKeys,omitempty"`
	MaxSkew int32 `json:"maxSkew"`
	MinDomains *int32 `json:"minDomains,omitempty"`
	NodeAffinityPolicy *string `json:"nodeAffinityPolicy,omitempty"`
	NodeTaintsPolicy *string `json:"nodeTaintsPolicy,omitempty"`
	TopologyKey string `json:"topologyKey"`
	WhenUnsatisfiable string `json:"whenUnsatisfiable"`
}

type _K8sIoV1TopologySpreadConstraint K8sIoV1TopologySpreadConstraint

// NewK8sIoV1TopologySpreadConstraint instantiates a new K8sIoV1TopologySpreadConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1TopologySpreadConstraint(maxSkew int32, topologyKey string, whenUnsatisfiable string) *K8sIoV1TopologySpreadConstraint {
	this := K8sIoV1TopologySpreadConstraint{}
	this.MaxSkew = maxSkew
	this.TopologyKey = topologyKey
	this.WhenUnsatisfiable = whenUnsatisfiable
	return &this
}

// NewK8sIoV1TopologySpreadConstraintWithDefaults instantiates a new K8sIoV1TopologySpreadConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1TopologySpreadConstraintWithDefaults() *K8sIoV1TopologySpreadConstraint {
	this := K8sIoV1TopologySpreadConstraint{}
	var maxSkew int32 = 0
	this.MaxSkew = maxSkew
	var topologyKey string = ""
	this.TopologyKey = topologyKey
	var whenUnsatisfiable string = ""
	this.WhenUnsatisfiable = whenUnsatisfiable
	return &this
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *K8sIoV1TopologySpreadConstraint) GetLabelSelector() K8sIoV1LabelSelector {
	if o == nil || IsNil(o.LabelSelector) {
		var ret K8sIoV1LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetLabelSelectorOk() (*K8sIoV1LabelSelector, bool) {
	if o == nil || IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *K8sIoV1TopologySpreadConstraint) HasLabelSelector() bool {
	if o != nil && !IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given K8sIoV1LabelSelector and assigns it to the LabelSelector field.
func (o *K8sIoV1TopologySpreadConstraint) SetLabelSelector(v K8sIoV1LabelSelector) {
	o.LabelSelector = &v
}

// GetMatchLabelKeys returns the MatchLabelKeys field value if set, zero value otherwise.
func (o *K8sIoV1TopologySpreadConstraint) GetMatchLabelKeys() []string {
	if o == nil || IsNil(o.MatchLabelKeys) {
		var ret []string
		return ret
	}
	return o.MatchLabelKeys
}

// GetMatchLabelKeysOk returns a tuple with the MatchLabelKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetMatchLabelKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.MatchLabelKeys) {
		return nil, false
	}
	return o.MatchLabelKeys, true
}

// HasMatchLabelKeys returns a boolean if a field has been set.
func (o *K8sIoV1TopologySpreadConstraint) HasMatchLabelKeys() bool {
	if o != nil && !IsNil(o.MatchLabelKeys) {
		return true
	}

	return false
}

// SetMatchLabelKeys gets a reference to the given []string and assigns it to the MatchLabelKeys field.
func (o *K8sIoV1TopologySpreadConstraint) SetMatchLabelKeys(v []string) {
	o.MatchLabelKeys = v
}

// GetMaxSkew returns the MaxSkew field value
func (o *K8sIoV1TopologySpreadConstraint) GetMaxSkew() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxSkew
}

// GetMaxSkewOk returns a tuple with the MaxSkew field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetMaxSkewOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSkew, true
}

// SetMaxSkew sets field value
func (o *K8sIoV1TopologySpreadConstraint) SetMaxSkew(v int32) {
	o.MaxSkew = v
}

// GetMinDomains returns the MinDomains field value if set, zero value otherwise.
func (o *K8sIoV1TopologySpreadConstraint) GetMinDomains() int32 {
	if o == nil || IsNil(o.MinDomains) {
		var ret int32
		return ret
	}
	return *o.MinDomains
}

// GetMinDomainsOk returns a tuple with the MinDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetMinDomainsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinDomains) {
		return nil, false
	}
	return o.MinDomains, true
}

// HasMinDomains returns a boolean if a field has been set.
func (o *K8sIoV1TopologySpreadConstraint) HasMinDomains() bool {
	if o != nil && !IsNil(o.MinDomains) {
		return true
	}

	return false
}

// SetMinDomains gets a reference to the given int32 and assigns it to the MinDomains field.
func (o *K8sIoV1TopologySpreadConstraint) SetMinDomains(v int32) {
	o.MinDomains = &v
}

// GetNodeAffinityPolicy returns the NodeAffinityPolicy field value if set, zero value otherwise.
func (o *K8sIoV1TopologySpreadConstraint) GetNodeAffinityPolicy() string {
	if o == nil || IsNil(o.NodeAffinityPolicy) {
		var ret string
		return ret
	}
	return *o.NodeAffinityPolicy
}

// GetNodeAffinityPolicyOk returns a tuple with the NodeAffinityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetNodeAffinityPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.NodeAffinityPolicy) {
		return nil, false
	}
	return o.NodeAffinityPolicy, true
}

// HasNodeAffinityPolicy returns a boolean if a field has been set.
func (o *K8sIoV1TopologySpreadConstraint) HasNodeAffinityPolicy() bool {
	if o != nil && !IsNil(o.NodeAffinityPolicy) {
		return true
	}

	return false
}

// SetNodeAffinityPolicy gets a reference to the given string and assigns it to the NodeAffinityPolicy field.
func (o *K8sIoV1TopologySpreadConstraint) SetNodeAffinityPolicy(v string) {
	o.NodeAffinityPolicy = &v
}

// GetNodeTaintsPolicy returns the NodeTaintsPolicy field value if set, zero value otherwise.
func (o *K8sIoV1TopologySpreadConstraint) GetNodeTaintsPolicy() string {
	if o == nil || IsNil(o.NodeTaintsPolicy) {
		var ret string
		return ret
	}
	return *o.NodeTaintsPolicy
}

// GetNodeTaintsPolicyOk returns a tuple with the NodeTaintsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetNodeTaintsPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.NodeTaintsPolicy) {
		return nil, false
	}
	return o.NodeTaintsPolicy, true
}

// HasNodeTaintsPolicy returns a boolean if a field has been set.
func (o *K8sIoV1TopologySpreadConstraint) HasNodeTaintsPolicy() bool {
	if o != nil && !IsNil(o.NodeTaintsPolicy) {
		return true
	}

	return false
}

// SetNodeTaintsPolicy gets a reference to the given string and assigns it to the NodeTaintsPolicy field.
func (o *K8sIoV1TopologySpreadConstraint) SetNodeTaintsPolicy(v string) {
	o.NodeTaintsPolicy = &v
}

// GetTopologyKey returns the TopologyKey field value
func (o *K8sIoV1TopologySpreadConstraint) GetTopologyKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TopologyKey
}

// GetTopologyKeyOk returns a tuple with the TopologyKey field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetTopologyKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopologyKey, true
}

// SetTopologyKey sets field value
func (o *K8sIoV1TopologySpreadConstraint) SetTopologyKey(v string) {
	o.TopologyKey = v
}

// GetWhenUnsatisfiable returns the WhenUnsatisfiable field value
func (o *K8sIoV1TopologySpreadConstraint) GetWhenUnsatisfiable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WhenUnsatisfiable
}

// GetWhenUnsatisfiableOk returns a tuple with the WhenUnsatisfiable field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1TopologySpreadConstraint) GetWhenUnsatisfiableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhenUnsatisfiable, true
}

// SetWhenUnsatisfiable sets field value
func (o *K8sIoV1TopologySpreadConstraint) SetWhenUnsatisfiable(v string) {
	o.WhenUnsatisfiable = v
}

func (o K8sIoV1TopologySpreadConstraint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1TopologySpreadConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	if !IsNil(o.MatchLabelKeys) {
		toSerialize["matchLabelKeys"] = o.MatchLabelKeys
	}
	toSerialize["maxSkew"] = o.MaxSkew
	if !IsNil(o.MinDomains) {
		toSerialize["minDomains"] = o.MinDomains
	}
	if !IsNil(o.NodeAffinityPolicy) {
		toSerialize["nodeAffinityPolicy"] = o.NodeAffinityPolicy
	}
	if !IsNil(o.NodeTaintsPolicy) {
		toSerialize["nodeTaintsPolicy"] = o.NodeTaintsPolicy
	}
	toSerialize["topologyKey"] = o.TopologyKey
	toSerialize["whenUnsatisfiable"] = o.WhenUnsatisfiable
	return toSerialize, nil
}

func (o *K8sIoV1TopologySpreadConstraint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"maxSkew",
		"topologyKey",
		"whenUnsatisfiable",
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

	varK8sIoV1TopologySpreadConstraint := _K8sIoV1TopologySpreadConstraint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sIoV1TopologySpreadConstraint)

	if err != nil {
		return err
	}

	*o = K8sIoV1TopologySpreadConstraint(varK8sIoV1TopologySpreadConstraint)

	return err
}

type NullableK8sIoV1TopologySpreadConstraint struct {
	value *K8sIoV1TopologySpreadConstraint
	isSet bool
}

func (v NullableK8sIoV1TopologySpreadConstraint) Get() *K8sIoV1TopologySpreadConstraint {
	return v.value
}

func (v *NullableK8sIoV1TopologySpreadConstraint) Set(val *K8sIoV1TopologySpreadConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1TopologySpreadConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1TopologySpreadConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1TopologySpreadConstraint(val *K8sIoV1TopologySpreadConstraint) *NullableK8sIoV1TopologySpreadConstraint {
	return &NullableK8sIoV1TopologySpreadConstraint{value: val, isSet: true}
}

func (v NullableK8sIoV1TopologySpreadConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1TopologySpreadConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


