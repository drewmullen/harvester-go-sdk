# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | Pointer to [**K8sIoV1TypedLocalObjectReference**](K8sIoV1TypedLocalObjectReference.md) | This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field. | [optional] 
**DataSourceRef** | Pointer to [**K8sIoV1TypedObjectReference**](K8sIoV1TypedObjectReference.md) | Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. | [optional] 
**Resources** | Pointer to [**K8sIoV1ResourceRequirements**](K8sIoV1ResourceRequirements.md) | Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources | [optional] [default to {}]
**Selector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) | A label query over volumes to consider for binding. | [optional] 
**StorageClassName** | Pointer to **string** | Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. | [optional] 
**VolumeName** | Pointer to **string** | VolumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpecWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpecWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpecWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetDataSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSource() K8sIoV1TypedLocalObjectReference`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceOk() (*K8sIoV1TypedLocalObjectReference, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetDataSource(v K8sIoV1TypedLocalObjectReference)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceRef() K8sIoV1TypedObjectReference`

GetDataSourceRef returns the DataSourceRef field if non-nil, zero value otherwise.

### GetDataSourceRefOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetDataSourceRefOk() (*K8sIoV1TypedObjectReference, bool)`

GetDataSourceRefOk returns a tuple with the DataSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetDataSourceRef(v K8sIoV1TypedObjectReference)`

SetDataSourceRef sets DataSourceRef field to given value.

### HasDataSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasDataSourceRef() bool`

HasDataSourceRef returns a boolean if a field has been set.

### GetResources

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetResources() K8sIoV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetResourcesOk() (*K8sIoV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetResources(v K8sIoV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSelector

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetSelector() K8sIoV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetSelectorOk() (*K8sIoV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetSelector(v K8sIoV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStorageClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetVolumeMode

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


