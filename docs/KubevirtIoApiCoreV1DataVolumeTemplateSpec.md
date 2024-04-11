# KubevirtIoApiCoreV1DataVolumeTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec.md) |  | [default to {}]
**Status** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1DataVolumeTemplateSpec

`func NewKubevirtIoApiCoreV1DataVolumeTemplateSpec(spec KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec, ) *KubevirtIoApiCoreV1DataVolumeTemplateSpec`

NewKubevirtIoApiCoreV1DataVolumeTemplateSpec instantiates a new KubevirtIoApiCoreV1DataVolumeTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DataVolumeTemplateSpecWithDefaults

`func NewKubevirtIoApiCoreV1DataVolumeTemplateSpecWithDefaults() *KubevirtIoApiCoreV1DataVolumeTemplateSpec`

NewKubevirtIoApiCoreV1DataVolumeTemplateSpecWithDefaults instantiates a new KubevirtIoApiCoreV1DataVolumeTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetSpec() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetSpecOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) SetSpec(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetStatus() map[string]interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) GetStatusOk() (*map[string]interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) SetStatus(v map[string]interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubevirtIoApiCoreV1DataVolumeTemplateSpec) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


