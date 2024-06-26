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

// checks if the KubevirtIoApiCoreV1PreferenceMatcher type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1PreferenceMatcher{}

// KubevirtIoApiCoreV1PreferenceMatcher PreferenceMatcher references a set of preference that is used to fill fields in the VMI template.
type KubevirtIoApiCoreV1PreferenceMatcher struct {
	// InferFromVolume lists the name of a volume that should be used to infer or discover the preference to be used through known annotations on the underlying resource. Once applied to the PreferenceMatcher this field is removed.
	InferFromVolume *string `json:"inferFromVolume,omitempty"`
	// InferFromVolumeFailurePolicy controls what should happen on failure when preference the instancetype. Allowed values are: \"RejectInferFromVolumeFailure\" and \"IgnoreInferFromVolumeFailure\". If not specified, \"RejectInferFromVolumeFailure\" is used by default.
	InferFromVolumeFailurePolicy *string `json:"inferFromVolumeFailurePolicy,omitempty"`
	// Kind specifies which preference resource is referenced. Allowed values are: \"VirtualMachinePreference\" and \"VirtualMachineClusterPreference\". If not specified, \"VirtualMachineClusterPreference\" is used by default.
	Kind *string `json:"kind,omitempty"`
	// Name is the name of the VirtualMachinePreference or VirtualMachineClusterPreference
	Name *string `json:"name,omitempty"`
	// RevisionName specifies a ControllerRevision containing a specific copy of the VirtualMachinePreference or VirtualMachineClusterPreference to be used. This is initially captured the first time the instancetype is applied to the VirtualMachineInstance.
	RevisionName *string `json:"revisionName,omitempty"`
}

// NewKubevirtIoApiCoreV1PreferenceMatcher instantiates a new KubevirtIoApiCoreV1PreferenceMatcher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1PreferenceMatcher() *KubevirtIoApiCoreV1PreferenceMatcher {
	this := KubevirtIoApiCoreV1PreferenceMatcher{}
	return &this
}

// NewKubevirtIoApiCoreV1PreferenceMatcherWithDefaults instantiates a new KubevirtIoApiCoreV1PreferenceMatcher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1PreferenceMatcherWithDefaults() *KubevirtIoApiCoreV1PreferenceMatcher {
	this := KubevirtIoApiCoreV1PreferenceMatcher{}
	return &this
}

// GetInferFromVolume returns the InferFromVolume field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetInferFromVolume() string {
	if o == nil || IsNil(o.InferFromVolume) {
		var ret string
		return ret
	}
	return *o.InferFromVolume
}

// GetInferFromVolumeOk returns a tuple with the InferFromVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetInferFromVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.InferFromVolume) {
		return nil, false
	}
	return o.InferFromVolume, true
}

// HasInferFromVolume returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) HasInferFromVolume() bool {
	if o != nil && !IsNil(o.InferFromVolume) {
		return true
	}

	return false
}

// SetInferFromVolume gets a reference to the given string and assigns it to the InferFromVolume field.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) SetInferFromVolume(v string) {
	o.InferFromVolume = &v
}

// GetInferFromVolumeFailurePolicy returns the InferFromVolumeFailurePolicy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetInferFromVolumeFailurePolicy() string {
	if o == nil || IsNil(o.InferFromVolumeFailurePolicy) {
		var ret string
		return ret
	}
	return *o.InferFromVolumeFailurePolicy
}

// GetInferFromVolumeFailurePolicyOk returns a tuple with the InferFromVolumeFailurePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetInferFromVolumeFailurePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.InferFromVolumeFailurePolicy) {
		return nil, false
	}
	return o.InferFromVolumeFailurePolicy, true
}

// HasInferFromVolumeFailurePolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) HasInferFromVolumeFailurePolicy() bool {
	if o != nil && !IsNil(o.InferFromVolumeFailurePolicy) {
		return true
	}

	return false
}

// SetInferFromVolumeFailurePolicy gets a reference to the given string and assigns it to the InferFromVolumeFailurePolicy field.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) SetInferFromVolumeFailurePolicy(v string) {
	o.InferFromVolumeFailurePolicy = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) SetName(v string) {
	o.Name = &v
}

// GetRevisionName returns the RevisionName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetRevisionName() string {
	if o == nil || IsNil(o.RevisionName) {
		var ret string
		return ret
	}
	return *o.RevisionName
}

// GetRevisionNameOk returns a tuple with the RevisionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) GetRevisionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RevisionName) {
		return nil, false
	}
	return o.RevisionName, true
}

// HasRevisionName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) HasRevisionName() bool {
	if o != nil && !IsNil(o.RevisionName) {
		return true
	}

	return false
}

// SetRevisionName gets a reference to the given string and assigns it to the RevisionName field.
func (o *KubevirtIoApiCoreV1PreferenceMatcher) SetRevisionName(v string) {
	o.RevisionName = &v
}

func (o KubevirtIoApiCoreV1PreferenceMatcher) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1PreferenceMatcher) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InferFromVolume) {
		toSerialize["inferFromVolume"] = o.InferFromVolume
	}
	if !IsNil(o.InferFromVolumeFailurePolicy) {
		toSerialize["inferFromVolumeFailurePolicy"] = o.InferFromVolumeFailurePolicy
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RevisionName) {
		toSerialize["revisionName"] = o.RevisionName
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1PreferenceMatcher struct {
	value *KubevirtIoApiCoreV1PreferenceMatcher
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1PreferenceMatcher) Get() *KubevirtIoApiCoreV1PreferenceMatcher {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1PreferenceMatcher) Set(val *KubevirtIoApiCoreV1PreferenceMatcher) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1PreferenceMatcher) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1PreferenceMatcher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1PreferenceMatcher(val *KubevirtIoApiCoreV1PreferenceMatcher) *NullableKubevirtIoApiCoreV1PreferenceMatcher {
	return &NullableKubevirtIoApiCoreV1PreferenceMatcher{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1PreferenceMatcher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1PreferenceMatcher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


