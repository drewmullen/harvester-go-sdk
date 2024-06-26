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

// checks if the K8sIoV1TypedLocalObjectReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1TypedLocalObjectReference{}

// K8sIoV1TypedLocalObjectReference TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.
type K8sIoV1TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.
	ApiGroup *string `json:"apiGroup,omitempty"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind"`
	// Name is the name of resource being referenced
	Name string `json:"name"`
}

type _K8sIoV1TypedLocalObjectReference K8sIoV1TypedLocalObjectReference

// NewK8sIoV1TypedLocalObjectReference instantiates a new K8sIoV1TypedLocalObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1TypedLocalObjectReference(kind string, name string) *K8sIoV1TypedLocalObjectReference {
	this := K8sIoV1TypedLocalObjectReference{}
	this.Kind = kind
	this.Name = name
	return &this
}

// NewK8sIoV1TypedLocalObjectReferenceWithDefaults instantiates a new K8sIoV1TypedLocalObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1TypedLocalObjectReferenceWithDefaults() *K8sIoV1TypedLocalObjectReference {
	this := K8sIoV1TypedLocalObjectReference{}
	var kind string = ""
	this.Kind = kind
	var name string = ""
	this.Name = name
	return &this
}

// GetApiGroup returns the ApiGroup field value if set, zero value otherwise.
func (o *K8sIoV1TypedLocalObjectReference) GetApiGroup() string {
	if o == nil || IsNil(o.ApiGroup) {
		var ret string
		return ret
	}
	return *o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1TypedLocalObjectReference) GetApiGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ApiGroup) {
		return nil, false
	}
	return o.ApiGroup, true
}

// HasApiGroup returns a boolean if a field has been set.
func (o *K8sIoV1TypedLocalObjectReference) HasApiGroup() bool {
	if o != nil && !IsNil(o.ApiGroup) {
		return true
	}

	return false
}

// SetApiGroup gets a reference to the given string and assigns it to the ApiGroup field.
func (o *K8sIoV1TypedLocalObjectReference) SetApiGroup(v string) {
	o.ApiGroup = &v
}

// GetKind returns the Kind field value
func (o *K8sIoV1TypedLocalObjectReference) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1TypedLocalObjectReference) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *K8sIoV1TypedLocalObjectReference) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *K8sIoV1TypedLocalObjectReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1TypedLocalObjectReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *K8sIoV1TypedLocalObjectReference) SetName(v string) {
	o.Name = v
}

func (o K8sIoV1TypedLocalObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1TypedLocalObjectReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiGroup) {
		toSerialize["apiGroup"] = o.ApiGroup
	}
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *K8sIoV1TypedLocalObjectReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
		"name",
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

	varK8sIoV1TypedLocalObjectReference := _K8sIoV1TypedLocalObjectReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sIoV1TypedLocalObjectReference)

	if err != nil {
		return err
	}

	*o = K8sIoV1TypedLocalObjectReference(varK8sIoV1TypedLocalObjectReference)

	return err
}

type NullableK8sIoV1TypedLocalObjectReference struct {
	value *K8sIoV1TypedLocalObjectReference
	isSet bool
}

func (v NullableK8sIoV1TypedLocalObjectReference) Get() *K8sIoV1TypedLocalObjectReference {
	return v.value
}

func (v *NullableK8sIoV1TypedLocalObjectReference) Set(val *K8sIoV1TypedLocalObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1TypedLocalObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1TypedLocalObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1TypedLocalObjectReference(val *K8sIoV1TypedLocalObjectReference) *NullableK8sIoV1TypedLocalObjectReference {
	return &NullableK8sIoV1TypedLocalObjectReference{value: val, isSet: true}
}

func (v NullableK8sIoV1TypedLocalObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1TypedLocalObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


