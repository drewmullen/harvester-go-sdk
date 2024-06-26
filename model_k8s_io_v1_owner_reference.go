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

// checks if the K8sIoV1OwnerReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1OwnerReference{}

// K8sIoV1OwnerReference OwnerReference contains enough information to let you identify an owning object. An owning object must be in the same namespace as the dependent, or be cluster-scoped, so there is no namespace field.
type K8sIoV1OwnerReference struct {
	// API version of the referent.
	ApiVersion string `json:"apiVersion"`
	// If true, AND if the owner has the \"foregroundDeletion\" finalizer, then the owner cannot be deleted from the key-value store until this reference is removed. See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion for how the garbage collector interacts with this field and enforces the foreground deletion. Defaults to false. To set this field, a user needs \"delete\" permission of the owner, otherwise 422 (Unprocessable Entity) will be returned.
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty"`
	// If true, this reference points to the managing controller.
	Controller *bool `json:"controller,omitempty"`
	// Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind"`
	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Name string `json:"name"`
	// UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	Uid string `json:"uid"`
}

type _K8sIoV1OwnerReference K8sIoV1OwnerReference

// NewK8sIoV1OwnerReference instantiates a new K8sIoV1OwnerReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1OwnerReference(apiVersion string, kind string, name string, uid string) *K8sIoV1OwnerReference {
	this := K8sIoV1OwnerReference{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Name = name
	this.Uid = uid
	return &this
}

// NewK8sIoV1OwnerReferenceWithDefaults instantiates a new K8sIoV1OwnerReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1OwnerReferenceWithDefaults() *K8sIoV1OwnerReference {
	this := K8sIoV1OwnerReference{}
	var apiVersion string = ""
	this.ApiVersion = apiVersion
	var kind string = ""
	this.Kind = kind
	var name string = ""
	this.Name = name
	var uid string = ""
	this.Uid = uid
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *K8sIoV1OwnerReference) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *K8sIoV1OwnerReference) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetBlockOwnerDeletion returns the BlockOwnerDeletion field value if set, zero value otherwise.
func (o *K8sIoV1OwnerReference) GetBlockOwnerDeletion() bool {
	if o == nil || IsNil(o.BlockOwnerDeletion) {
		var ret bool
		return ret
	}
	return *o.BlockOwnerDeletion
}

// GetBlockOwnerDeletionOk returns a tuple with the BlockOwnerDeletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetBlockOwnerDeletionOk() (*bool, bool) {
	if o == nil || IsNil(o.BlockOwnerDeletion) {
		return nil, false
	}
	return o.BlockOwnerDeletion, true
}

// HasBlockOwnerDeletion returns a boolean if a field has been set.
func (o *K8sIoV1OwnerReference) HasBlockOwnerDeletion() bool {
	if o != nil && !IsNil(o.BlockOwnerDeletion) {
		return true
	}

	return false
}

// SetBlockOwnerDeletion gets a reference to the given bool and assigns it to the BlockOwnerDeletion field.
func (o *K8sIoV1OwnerReference) SetBlockOwnerDeletion(v bool) {
	o.BlockOwnerDeletion = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *K8sIoV1OwnerReference) GetController() bool {
	if o == nil || IsNil(o.Controller) {
		var ret bool
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetControllerOk() (*bool, bool) {
	if o == nil || IsNil(o.Controller) {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *K8sIoV1OwnerReference) HasController() bool {
	if o != nil && !IsNil(o.Controller) {
		return true
	}

	return false
}

// SetController gets a reference to the given bool and assigns it to the Controller field.
func (o *K8sIoV1OwnerReference) SetController(v bool) {
	o.Controller = &v
}

// GetKind returns the Kind field value
func (o *K8sIoV1OwnerReference) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *K8sIoV1OwnerReference) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *K8sIoV1OwnerReference) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *K8sIoV1OwnerReference) SetName(v string) {
	o.Name = v
}

// GetUid returns the Uid field value
func (o *K8sIoV1OwnerReference) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1OwnerReference) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *K8sIoV1OwnerReference) SetUid(v string) {
	o.Uid = v
}

func (o K8sIoV1OwnerReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1OwnerReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	if !IsNil(o.BlockOwnerDeletion) {
		toSerialize["blockOwnerDeletion"] = o.BlockOwnerDeletion
	}
	if !IsNil(o.Controller) {
		toSerialize["controller"] = o.Controller
	}
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	toSerialize["uid"] = o.Uid
	return toSerialize, nil
}

func (o *K8sIoV1OwnerReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"apiVersion",
		"kind",
		"name",
		"uid",
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

	varK8sIoV1OwnerReference := _K8sIoV1OwnerReference{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sIoV1OwnerReference)

	if err != nil {
		return err
	}

	*o = K8sIoV1OwnerReference(varK8sIoV1OwnerReference)

	return err
}

type NullableK8sIoV1OwnerReference struct {
	value *K8sIoV1OwnerReference
	isSet bool
}

func (v NullableK8sIoV1OwnerReference) Get() *K8sIoV1OwnerReference {
	return v.value
}

func (v *NullableK8sIoV1OwnerReference) Set(val *K8sIoV1OwnerReference) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1OwnerReference) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1OwnerReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1OwnerReference(val *K8sIoV1OwnerReference) *NullableK8sIoV1OwnerReference {
	return &NullableK8sIoV1OwnerReference{value: val, isSet: true}
}

func (v NullableK8sIoV1OwnerReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1OwnerReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


