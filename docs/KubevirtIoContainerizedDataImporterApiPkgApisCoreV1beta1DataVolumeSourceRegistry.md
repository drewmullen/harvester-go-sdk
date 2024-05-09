# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertConfigMap** | Pointer to **string** | CertConfigMap provides a reference to the Registry certs | [optional] 
**ImageStream** | Pointer to **string** | ImageStream is the name of image stream for import | [optional] 
**PullMethod** | Pointer to **string** | PullMethod can be either \&quot;pod\&quot; (default import), or \&quot;node\&quot; (node docker cache based import) | [optional] 
**SecretRef** | Pointer to **string** | SecretRef provides the secret reference needed to access the Registry source | [optional] 
**Url** | Pointer to **string** | URL is the url of the registry source (starting with the scheme: docker, oci-archive) | [optional] 

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistryWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistryWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistryWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetCertConfigMap() string`

GetCertConfigMap returns the CertConfigMap field if non-nil, zero value otherwise.

### GetCertConfigMapOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetCertConfigMapOk() (*string, bool)`

GetCertConfigMapOk returns a tuple with the CertConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) SetCertConfigMap(v string)`

SetCertConfigMap sets CertConfigMap field to given value.

### HasCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) HasCertConfigMap() bool`

HasCertConfigMap returns a boolean if a field has been set.

### GetImageStream

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetImageStream() string`

GetImageStream returns the ImageStream field if non-nil, zero value otherwise.

### GetImageStreamOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetImageStreamOk() (*string, bool)`

GetImageStreamOk returns a tuple with the ImageStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStream

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) SetImageStream(v string)`

SetImageStream sets ImageStream field to given value.

### HasImageStream

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) HasImageStream() bool`

HasImageStream returns a boolean if a field has been set.

### GetPullMethod

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetPullMethod() string`

GetPullMethod returns the PullMethod field if non-nil, zero value otherwise.

### GetPullMethodOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetPullMethodOk() (*string, bool)`

GetPullMethodOk returns a tuple with the PullMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullMethod

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) SetPullMethod(v string)`

SetPullMethod sets PullMethod field to given value.

### HasPullMethod

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) HasPullMethod() bool`

HasPullMethod returns a boolean if a field has been set.

### GetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetSecretRef() string`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetSecretRefOk() (*string, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) SetSecretRef(v string)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


