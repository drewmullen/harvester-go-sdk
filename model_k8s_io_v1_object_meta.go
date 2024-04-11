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

// checks if the K8sIoV1ObjectMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1ObjectMeta{}

// K8sIoV1ObjectMeta struct for K8sIoV1ObjectMeta
type K8sIoV1ObjectMeta struct {
	Annotations *map[string]string `json:"annotations,omitempty"`
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
	DeletionGracePeriodSeconds *int64 `json:"deletionGracePeriodSeconds,omitempty"`
	DeletionTimestamp *string `json:"deletionTimestamp,omitempty"`
	Finalizers []string `json:"finalizers,omitempty"`
	GenerateName *string `json:"generateName,omitempty"`
	Generation *int64 `json:"generation,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
	ManagedFields []K8sIoV1ManagedFieldsEntry `json:"managedFields,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	OwnerReferences []K8sIoV1OwnerReference `json:"ownerReferences,omitempty"`
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	SelfLink *string `json:"selfLink,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewK8sIoV1ObjectMeta instantiates a new K8sIoV1ObjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1ObjectMeta() *K8sIoV1ObjectMeta {
	this := K8sIoV1ObjectMeta{}
	var creationTimestamp string = "{}"
	this.CreationTimestamp = &creationTimestamp
	var deletionTimestamp string = ""
	this.DeletionTimestamp = &deletionTimestamp
	return &this
}

// NewK8sIoV1ObjectMetaWithDefaults instantiates a new K8sIoV1ObjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1ObjectMetaWithDefaults() *K8sIoV1ObjectMeta {
	this := K8sIoV1ObjectMeta{}
	var creationTimestamp string = "{}"
	this.CreationTimestamp = &creationTimestamp
	var deletionTimestamp string = ""
	this.DeletionTimestamp = &deletionTimestamp
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetAnnotations() map[string]string {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]string
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetAnnotationsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Annotations) {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the Annotations field.
func (o *K8sIoV1ObjectMeta) SetAnnotations(v map[string]string) {
	o.Annotations = &v
}

// GetCreationTimestamp returns the CreationTimestamp field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetCreationTimestamp() string {
	if o == nil || IsNil(o.CreationTimestamp) {
		var ret string
		return ret
	}
	return *o.CreationTimestamp
}

// GetCreationTimestampOk returns a tuple with the CreationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetCreationTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.CreationTimestamp) {
		return nil, false
	}
	return o.CreationTimestamp, true
}

// HasCreationTimestamp returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasCreationTimestamp() bool {
	if o != nil && !IsNil(o.CreationTimestamp) {
		return true
	}

	return false
}

// SetCreationTimestamp gets a reference to the given string and assigns it to the CreationTimestamp field.
func (o *K8sIoV1ObjectMeta) SetCreationTimestamp(v string) {
	o.CreationTimestamp = &v
}

// GetDeletionGracePeriodSeconds returns the DeletionGracePeriodSeconds field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetDeletionGracePeriodSeconds() int64 {
	if o == nil || IsNil(o.DeletionGracePeriodSeconds) {
		var ret int64
		return ret
	}
	return *o.DeletionGracePeriodSeconds
}

// GetDeletionGracePeriodSecondsOk returns a tuple with the DeletionGracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetDeletionGracePeriodSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.DeletionGracePeriodSeconds) {
		return nil, false
	}
	return o.DeletionGracePeriodSeconds, true
}

// HasDeletionGracePeriodSeconds returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasDeletionGracePeriodSeconds() bool {
	if o != nil && !IsNil(o.DeletionGracePeriodSeconds) {
		return true
	}

	return false
}

// SetDeletionGracePeriodSeconds gets a reference to the given int64 and assigns it to the DeletionGracePeriodSeconds field.
func (o *K8sIoV1ObjectMeta) SetDeletionGracePeriodSeconds(v int64) {
	o.DeletionGracePeriodSeconds = &v
}

// GetDeletionTimestamp returns the DeletionTimestamp field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetDeletionTimestamp() string {
	if o == nil || IsNil(o.DeletionTimestamp) {
		var ret string
		return ret
	}
	return *o.DeletionTimestamp
}

// GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetDeletionTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.DeletionTimestamp) {
		return nil, false
	}
	return o.DeletionTimestamp, true
}

// HasDeletionTimestamp returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasDeletionTimestamp() bool {
	if o != nil && !IsNil(o.DeletionTimestamp) {
		return true
	}

	return false
}

// SetDeletionTimestamp gets a reference to the given string and assigns it to the DeletionTimestamp field.
func (o *K8sIoV1ObjectMeta) SetDeletionTimestamp(v string) {
	o.DeletionTimestamp = &v
}

// GetFinalizers returns the Finalizers field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetFinalizers() []string {
	if o == nil || IsNil(o.Finalizers) {
		var ret []string
		return ret
	}
	return o.Finalizers
}

// GetFinalizersOk returns a tuple with the Finalizers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetFinalizersOk() ([]string, bool) {
	if o == nil || IsNil(o.Finalizers) {
		return nil, false
	}
	return o.Finalizers, true
}

// HasFinalizers returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasFinalizers() bool {
	if o != nil && !IsNil(o.Finalizers) {
		return true
	}

	return false
}

// SetFinalizers gets a reference to the given []string and assigns it to the Finalizers field.
func (o *K8sIoV1ObjectMeta) SetFinalizers(v []string) {
	o.Finalizers = v
}

// GetGenerateName returns the GenerateName field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetGenerateName() string {
	if o == nil || IsNil(o.GenerateName) {
		var ret string
		return ret
	}
	return *o.GenerateName
}

// GetGenerateNameOk returns a tuple with the GenerateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetGenerateNameOk() (*string, bool) {
	if o == nil || IsNil(o.GenerateName) {
		return nil, false
	}
	return o.GenerateName, true
}

// HasGenerateName returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasGenerateName() bool {
	if o != nil && !IsNil(o.GenerateName) {
		return true
	}

	return false
}

// SetGenerateName gets a reference to the given string and assigns it to the GenerateName field.
func (o *K8sIoV1ObjectMeta) SetGenerateName(v string) {
	o.GenerateName = &v
}

// GetGeneration returns the Generation field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetGeneration() int64 {
	if o == nil || IsNil(o.Generation) {
		var ret int64
		return ret
	}
	return *o.Generation
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetGenerationOk() (*int64, bool) {
	if o == nil || IsNil(o.Generation) {
		return nil, false
	}
	return o.Generation, true
}

// HasGeneration returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasGeneration() bool {
	if o != nil && !IsNil(o.Generation) {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given int64 and assigns it to the Generation field.
func (o *K8sIoV1ObjectMeta) SetGeneration(v int64) {
	o.Generation = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *K8sIoV1ObjectMeta) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetManagedFields returns the ManagedFields field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetManagedFields() []K8sIoV1ManagedFieldsEntry {
	if o == nil || IsNil(o.ManagedFields) {
		var ret []K8sIoV1ManagedFieldsEntry
		return ret
	}
	return o.ManagedFields
}

// GetManagedFieldsOk returns a tuple with the ManagedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetManagedFieldsOk() ([]K8sIoV1ManagedFieldsEntry, bool) {
	if o == nil || IsNil(o.ManagedFields) {
		return nil, false
	}
	return o.ManagedFields, true
}

// HasManagedFields returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasManagedFields() bool {
	if o != nil && !IsNil(o.ManagedFields) {
		return true
	}

	return false
}

// SetManagedFields gets a reference to the given []K8sIoV1ManagedFieldsEntry and assigns it to the ManagedFields field.
func (o *K8sIoV1ObjectMeta) SetManagedFields(v []K8sIoV1ManagedFieldsEntry) {
	o.ManagedFields = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *K8sIoV1ObjectMeta) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *K8sIoV1ObjectMeta) SetNamespace(v string) {
	o.Namespace = &v
}

// GetOwnerReferences returns the OwnerReferences field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetOwnerReferences() []K8sIoV1OwnerReference {
	if o == nil || IsNil(o.OwnerReferences) {
		var ret []K8sIoV1OwnerReference
		return ret
	}
	return o.OwnerReferences
}

// GetOwnerReferencesOk returns a tuple with the OwnerReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetOwnerReferencesOk() ([]K8sIoV1OwnerReference, bool) {
	if o == nil || IsNil(o.OwnerReferences) {
		return nil, false
	}
	return o.OwnerReferences, true
}

// HasOwnerReferences returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasOwnerReferences() bool {
	if o != nil && !IsNil(o.OwnerReferences) {
		return true
	}

	return false
}

// SetOwnerReferences gets a reference to the given []K8sIoV1OwnerReference and assigns it to the OwnerReferences field.
func (o *K8sIoV1ObjectMeta) SetOwnerReferences(v []K8sIoV1OwnerReference) {
	o.OwnerReferences = v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *K8sIoV1ObjectMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetSelfLink returns the SelfLink field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetSelfLink() string {
	if o == nil || IsNil(o.SelfLink) {
		var ret string
		return ret
	}
	return *o.SelfLink
}

// GetSelfLinkOk returns a tuple with the SelfLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetSelfLinkOk() (*string, bool) {
	if o == nil || IsNil(o.SelfLink) {
		return nil, false
	}
	return o.SelfLink, true
}

// HasSelfLink returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasSelfLink() bool {
	if o != nil && !IsNil(o.SelfLink) {
		return true
	}

	return false
}

// SetSelfLink gets a reference to the given string and assigns it to the SelfLink field.
func (o *K8sIoV1ObjectMeta) SetSelfLink(v string) {
	o.SelfLink = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *K8sIoV1ObjectMeta) GetUid() string {
	if o == nil || IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1ObjectMeta) GetUidOk() (*string, bool) {
	if o == nil || IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *K8sIoV1ObjectMeta) HasUid() bool {
	if o != nil && !IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *K8sIoV1ObjectMeta) SetUid(v string) {
	o.Uid = &v
}

func (o K8sIoV1ObjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1ObjectMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	if !IsNil(o.CreationTimestamp) {
		toSerialize["creationTimestamp"] = o.CreationTimestamp
	}
	if !IsNil(o.DeletionGracePeriodSeconds) {
		toSerialize["deletionGracePeriodSeconds"] = o.DeletionGracePeriodSeconds
	}
	if !IsNil(o.DeletionTimestamp) {
		toSerialize["deletionTimestamp"] = o.DeletionTimestamp
	}
	if !IsNil(o.Finalizers) {
		toSerialize["finalizers"] = o.Finalizers
	}
	if !IsNil(o.GenerateName) {
		toSerialize["generateName"] = o.GenerateName
	}
	if !IsNil(o.Generation) {
		toSerialize["generation"] = o.Generation
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.ManagedFields) {
		toSerialize["managedFields"] = o.ManagedFields
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.OwnerReferences) {
		toSerialize["ownerReferences"] = o.OwnerReferences
	}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if !IsNil(o.SelfLink) {
		toSerialize["selfLink"] = o.SelfLink
	}
	if !IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	return toSerialize, nil
}

type NullableK8sIoV1ObjectMeta struct {
	value *K8sIoV1ObjectMeta
	isSet bool
}

func (v NullableK8sIoV1ObjectMeta) Get() *K8sIoV1ObjectMeta {
	return v.value
}

func (v *NullableK8sIoV1ObjectMeta) Set(val *K8sIoV1ObjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1ObjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1ObjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1ObjectMeta(val *K8sIoV1ObjectMeta) *NullableK8sIoV1ObjectMeta {
	return &NullableK8sIoV1ObjectMeta{value: val, isSet: true}
}

func (v NullableK8sIoV1ObjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1ObjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


