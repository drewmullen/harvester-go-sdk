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

// checks if the KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec{}

// KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec StorageSpec defines the Storage type specification
type KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec struct {
	// AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	AccessModes []string `json:"accessModes,omitempty"`
	// This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.
	DataSource *K8sIoV1TypedLocalObjectReference `json:"dataSource,omitempty"`
	// Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.
	DataSourceRef *K8sIoV1TypedObjectReference `json:"dataSourceRef,omitempty"`
	// Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	Resources *K8sIoV1ResourceRequirements `json:"resources,omitempty"`
	// A label query over volumes to consider for binding.
	Selector *K8sIoV1LabelSelector `json:"selector,omitempty"`
	// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	StorageClassName *string `json:"storageClassName,omitempty"`
	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.
	VolumeMode *string `json:"volumeMode,omitempty"`
	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	VolumeName *string `json:"volumeName,omitempty"`
}

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec{}
	var resources K8sIoV1ResourceRequirements
	this.Resources = &resources
	return &this
}

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpecWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpecWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec{}
	var resources K8sIoV1ResourceRequirements
	this.Resources = &resources
	return &this
}

// GetAccessModes returns the AccessModes field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetAccessModes() []string {
	if o == nil || IsNil(o.AccessModes) {
		var ret []string
		return ret
	}
	return o.AccessModes
}

// GetAccessModesOk returns a tuple with the AccessModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetAccessModesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessModes) {
		return nil, false
	}
	return o.AccessModes, true
}

// HasAccessModes returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasAccessModes() bool {
	if o != nil && !IsNil(o.AccessModes) {
		return true
	}

	return false
}

// SetAccessModes gets a reference to the given []string and assigns it to the AccessModes field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetAccessModes(v []string) {
	o.AccessModes = v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSource() K8sIoV1TypedLocalObjectReference {
	if o == nil || IsNil(o.DataSource) {
		var ret K8sIoV1TypedLocalObjectReference
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceOk() (*K8sIoV1TypedLocalObjectReference, bool) {
	if o == nil || IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasDataSource() bool {
	if o != nil && !IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given K8sIoV1TypedLocalObjectReference and assigns it to the DataSource field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetDataSource(v K8sIoV1TypedLocalObjectReference) {
	o.DataSource = &v
}

// GetDataSourceRef returns the DataSourceRef field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceRef() K8sIoV1TypedObjectReference {
	if o == nil || IsNil(o.DataSourceRef) {
		var ret K8sIoV1TypedObjectReference
		return ret
	}
	return *o.DataSourceRef
}

// GetDataSourceRefOk returns a tuple with the DataSourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceRefOk() (*K8sIoV1TypedObjectReference, bool) {
	if o == nil || IsNil(o.DataSourceRef) {
		return nil, false
	}
	return o.DataSourceRef, true
}

// HasDataSourceRef returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasDataSourceRef() bool {
	if o != nil && !IsNil(o.DataSourceRef) {
		return true
	}

	return false
}

// SetDataSourceRef gets a reference to the given K8sIoV1TypedObjectReference and assigns it to the DataSourceRef field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetDataSourceRef(v K8sIoV1TypedObjectReference) {
	o.DataSourceRef = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetResources() K8sIoV1ResourceRequirements {
	if o == nil || IsNil(o.Resources) {
		var ret K8sIoV1ResourceRequirements
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetResourcesOk() (*K8sIoV1ResourceRequirements, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given K8sIoV1ResourceRequirements and assigns it to the Resources field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetResources(v K8sIoV1ResourceRequirements) {
	o.Resources = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetSelector() K8sIoV1LabelSelector {
	if o == nil || IsNil(o.Selector) {
		var ret K8sIoV1LabelSelector
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetSelectorOk() (*K8sIoV1LabelSelector, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given K8sIoV1LabelSelector and assigns it to the Selector field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetSelector(v K8sIoV1LabelSelector) {
	o.Selector = &v
}

// GetStorageClassName returns the StorageClassName field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetStorageClassName() string {
	if o == nil || IsNil(o.StorageClassName) {
		var ret string
		return ret
	}
	return *o.StorageClassName
}

// GetStorageClassNameOk returns a tuple with the StorageClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetStorageClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClassName) {
		return nil, false
	}
	return o.StorageClassName, true
}

// HasStorageClassName returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasStorageClassName() bool {
	if o != nil && !IsNil(o.StorageClassName) {
		return true
	}

	return false
}

// SetStorageClassName gets a reference to the given string and assigns it to the StorageClassName field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetStorageClassName(v string) {
	o.StorageClassName = &v
}

// GetVolumeMode returns the VolumeMode field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeMode() string {
	if o == nil || IsNil(o.VolumeMode) {
		var ret string
		return ret
	}
	return *o.VolumeMode
}

// GetVolumeModeOk returns a tuple with the VolumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeModeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeMode) {
		return nil, false
	}
	return o.VolumeMode, true
}

// HasVolumeMode returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasVolumeMode() bool {
	if o != nil && !IsNil(o.VolumeMode) {
		return true
	}

	return false
}

// SetVolumeMode gets a reference to the given string and assigns it to the VolumeMode field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetVolumeMode(v string) {
	o.VolumeMode = &v
}

// GetVolumeName returns the VolumeName field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeName() string {
	if o == nil || IsNil(o.VolumeName) {
		var ret string
		return ret
	}
	return *o.VolumeName
}

// GetVolumeNameOk returns a tuple with the VolumeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeNameOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeName) {
		return nil, false
	}
	return o.VolumeName, true
}

// HasVolumeName returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasVolumeName() bool {
	if o != nil && !IsNil(o.VolumeName) {
		return true
	}

	return false
}

// SetVolumeName gets a reference to the given string and assigns it to the VolumeName field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetVolumeName(v string) {
	o.VolumeName = &v
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessModes) {
		toSerialize["accessModes"] = o.AccessModes
	}
	if !IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	if !IsNil(o.DataSourceRef) {
		toSerialize["dataSourceRef"] = o.DataSourceRef
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !IsNil(o.StorageClassName) {
		toSerialize["storageClassName"] = o.StorageClassName
	}
	if !IsNil(o.VolumeMode) {
		toSerialize["volumeMode"] = o.VolumeMode
	}
	if !IsNil(o.VolumeName) {
		toSerialize["volumeName"] = o.VolumeName
	}
	return toSerialize, nil
}

type NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec struct {
	value *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec
	isSet bool
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) Get() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec {
	return v.value
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) Set(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec {
	return &NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec{value: val, isSet: true}
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


