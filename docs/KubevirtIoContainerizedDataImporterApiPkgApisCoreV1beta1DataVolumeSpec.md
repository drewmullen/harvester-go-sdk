# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkpoints** | Pointer to [**[]KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint.md) |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**FinalCheckpoint** | Pointer to **bool** |  | [optional] 
**Preallocation** | Pointer to **bool** |  | [optional] 
**PriorityClassName** | Pointer to **string** |  | [optional] 
**Pvc** | Pointer to [**K8sIoV1PersistentVolumeClaimSpec**](K8sIoV1PersistentVolumeClaimSpec.md) |  | [optional] 
**Source** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource.md) |  | [optional] 
**SourceRef** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef.md) |  | [optional] 
**Storage** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec.md) |  | [optional] 

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpecWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpecWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpecWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckpoints

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetCheckpoints() []KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint`

GetCheckpoints returns the Checkpoints field if non-nil, zero value otherwise.

### GetCheckpointsOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetCheckpointsOk() (*[]KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint, bool)`

GetCheckpointsOk returns a tuple with the Checkpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpoints

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetCheckpoints(v []KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint)`

SetCheckpoints sets Checkpoints field to given value.

### HasCheckpoints

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasCheckpoints() bool`

HasCheckpoints returns a boolean if a field has been set.

### GetContentType

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFinalCheckpoint

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetFinalCheckpoint() bool`

GetFinalCheckpoint returns the FinalCheckpoint field if non-nil, zero value otherwise.

### GetFinalCheckpointOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetFinalCheckpointOk() (*bool, bool)`

GetFinalCheckpointOk returns a tuple with the FinalCheckpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalCheckpoint

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetFinalCheckpoint(v bool)`

SetFinalCheckpoint sets FinalCheckpoint field to given value.

### HasFinalCheckpoint

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasFinalCheckpoint() bool`

HasFinalCheckpoint returns a boolean if a field has been set.

### GetPreallocation

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPreallocation() bool`

GetPreallocation returns the Preallocation field if non-nil, zero value otherwise.

### GetPreallocationOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPreallocationOk() (*bool, bool)`

GetPreallocationOk returns a tuple with the Preallocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreallocation

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetPreallocation(v bool)`

SetPreallocation sets Preallocation field to given value.

### HasPreallocation

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasPreallocation() bool`

HasPreallocation returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPvc() K8sIoV1PersistentVolumeClaimSpec`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetPvcOk() (*K8sIoV1PersistentVolumeClaimSpec, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetPvc(v K8sIoV1PersistentVolumeClaimSpec)`

SetPvc sets Pvc field to given value.

### HasPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasPvc() bool`

HasPvc returns a boolean if a field has been set.

### GetSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetSource() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetSourceOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetSource(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetSourceRef() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef`

GetSourceRef returns the SourceRef field if non-nil, zero value otherwise.

### GetSourceRefOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetSourceRefOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef, bool)`

GetSourceRefOk returns a tuple with the SourceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetSourceRef(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef)`

SetSourceRef sets SourceRef field to given value.

### HasSourceRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasSourceRef() bool`

HasSourceRef returns a boolean if a field has been set.

### GetStorage

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetStorage() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) GetStorageOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) SetStorage(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1StorageSpec)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


