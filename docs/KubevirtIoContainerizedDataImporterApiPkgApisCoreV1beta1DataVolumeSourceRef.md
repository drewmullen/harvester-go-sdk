# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** | The kind of the source reference, currently only \&quot;DataSource\&quot; is supported | [default to ""]
**Name** | **string** | The name of the source reference | [default to ""]
**Namespace** | Pointer to **string** | The namespace of the source reference, defaults to the DataVolume namespace | [optional] 

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef(kind string, name string, ) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRefWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRefWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRefWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


