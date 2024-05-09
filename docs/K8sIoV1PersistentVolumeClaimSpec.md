# K8sIoV1PersistentVolumeClaimSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**DataSource** | Pointer to [**K8sIoV1TypedLocalObjectReference**](K8sIoV1TypedLocalObjectReference.md) | dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource. | [optional] 
**DataSourceRef** | Pointer to [**K8sIoV1TypedObjectReference**](K8sIoV1TypedObjectReference.md) | dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn&#39;t specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn&#39;t set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. * While dataSource only allows local objects, dataSourceRef allows objects   in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled. | [optional] 
**Resources** | Pointer to [**K8sIoV1ResourceRequirements**](K8sIoV1ResourceRequirements.md) | resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources | [optional] [default to {}]
**Selector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) | selector is a label query over volumes to consider for binding. | [optional] 
**StorageClassName** | Pointer to **string** | storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 | [optional] 
**VolumeMode** | Pointer to **string** | volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. | [optional] 
**VolumeName** | Pointer to **string** | volumeName is the binding reference to the PersistentVolume backing this claim. | [optional] 

## Methods

### NewK8sIoV1PersistentVolumeClaimSpec

`func NewK8sIoV1PersistentVolumeClaimSpec() *K8sIoV1PersistentVolumeClaimSpec`

NewK8sIoV1PersistentVolumeClaimSpec instantiates a new K8sIoV1PersistentVolumeClaimSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimSpecWithDefaults

`func NewK8sIoV1PersistentVolumeClaimSpecWithDefaults() *K8sIoV1PersistentVolumeClaimSpec`

NewK8sIoV1PersistentVolumeClaimSpecWithDefaults instantiates a new K8sIoV1PersistentVolumeClaimSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetDataSource

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetDataSource() K8sIoV1TypedLocalObjectReference`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetDataSourceOk() (*K8sIoV1TypedLocalObjectReference, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetDataSource(v K8sIoV1TypedLocalObjectReference)`

SetDataSource sets DataSource field to given value.

### HasDataSource

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasDataSource() bool`

HasDataSource returns a boolean if a field has been set.

### GetDataSourceRef

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetDataSourceRef() K8sIoV1TypedObjectReference`

GetDataSourceRef returns the DataSourceRef field if non-nil, zero value otherwise.

### GetDataSourceRefOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetDataSourceRefOk() (*K8sIoV1TypedObjectReference, bool)`

GetDataSourceRefOk returns a tuple with the DataSourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceRef

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetDataSourceRef(v K8sIoV1TypedObjectReference)`

SetDataSourceRef sets DataSourceRef field to given value.

### HasDataSourceRef

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasDataSourceRef() bool`

HasDataSourceRef returns a boolean if a field has been set.

### GetResources

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetResources() K8sIoV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetResourcesOk() (*K8sIoV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetResources(v K8sIoV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSelector

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetSelector() K8sIoV1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetSelectorOk() (*K8sIoV1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetSelector(v K8sIoV1LabelSelector)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStorageClassName

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.

### HasStorageClassName

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasStorageClassName() bool`

HasStorageClassName returns a boolean if a field has been set.

### GetVolumeMode

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetVolumeMode() string`

GetVolumeMode returns the VolumeMode field if non-nil, zero value otherwise.

### GetVolumeModeOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetVolumeModeOk() (*string, bool)`

GetVolumeModeOk returns a tuple with the VolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMode

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetVolumeMode(v string)`

SetVolumeMode sets VolumeMode field to given value.

### HasVolumeMode

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasVolumeMode() bool`

HasVolumeMode returns a boolean if a field has been set.

### GetVolumeName

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *K8sIoV1PersistentVolumeClaimSpec) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *K8sIoV1PersistentVolumeClaimSpec) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *K8sIoV1PersistentVolumeClaimSpec) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


