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

// checks if the K8sIoV1Preconditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1Preconditions{}

// K8sIoV1Preconditions Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.
type K8sIoV1Preconditions struct {
	// Specifies the target ResourceVersion
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	// Specifies the target UID.
	Uid *string `json:"uid,omitempty"`
}

// NewK8sIoV1Preconditions instantiates a new K8sIoV1Preconditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1Preconditions() *K8sIoV1Preconditions {
	this := K8sIoV1Preconditions{}
	return &this
}

// NewK8sIoV1PreconditionsWithDefaults instantiates a new K8sIoV1Preconditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1PreconditionsWithDefaults() *K8sIoV1Preconditions {
	this := K8sIoV1Preconditions{}
	return &this
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *K8sIoV1Preconditions) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1Preconditions) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *K8sIoV1Preconditions) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *K8sIoV1Preconditions) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *K8sIoV1Preconditions) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1Preconditions) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *K8sIoV1Preconditions) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *K8sIoV1Preconditions) SetUid(v string) {
	o.Uid = &v
}

func (o K8sIoV1Preconditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1Preconditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableK8sIoV1Preconditions struct {
	value *K8sIoV1Preconditions
	isSet bool
}

func (v NullableK8sIoV1Preconditions) Get() *K8sIoV1Preconditions {
	return v.value
}

func (v *NullableK8sIoV1Preconditions) Set(val *K8sIoV1Preconditions) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1Preconditions) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1Preconditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1Preconditions(val *K8sIoV1Preconditions) *NullableK8sIoV1Preconditions {
	return &NullableK8sIoV1Preconditions{value: val, isSet: true}
}

func (v NullableK8sIoV1Preconditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1Preconditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


