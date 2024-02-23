# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertConfigMap** | Pointer to **string** |  | [optional] 
**ExtraHeaders** | Pointer to **[]string** |  | [optional] 
**SecretExtraHeaders** | Pointer to **[]string** |  | [optional] 
**SecretRef** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP(url string, ) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTPWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTPWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTPWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetCertConfigMap() string`

GetCertConfigMap returns the CertConfigMap field if non-nil, zero value otherwise.

### GetCertConfigMapOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetCertConfigMapOk() (*string, bool)`

GetCertConfigMapOk returns a tuple with the CertConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) SetCertConfigMap(v string)`

SetCertConfigMap sets CertConfigMap field to given value.

### HasCertConfigMap

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) HasCertConfigMap() bool`

HasCertConfigMap returns a boolean if a field has been set.

### GetExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetExtraHeaders() []string`

GetExtraHeaders returns the ExtraHeaders field if non-nil, zero value otherwise.

### GetExtraHeadersOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetExtraHeadersOk() (*[]string, bool)`

GetExtraHeadersOk returns a tuple with the ExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) SetExtraHeaders(v []string)`

SetExtraHeaders sets ExtraHeaders field to given value.

### HasExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) HasExtraHeaders() bool`

HasExtraHeaders returns a boolean if a field has been set.

### GetSecretExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetSecretExtraHeaders() []string`

GetSecretExtraHeaders returns the SecretExtraHeaders field if non-nil, zero value otherwise.

### GetSecretExtraHeadersOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetSecretExtraHeadersOk() (*[]string, bool)`

GetSecretExtraHeadersOk returns a tuple with the SecretExtraHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) SetSecretExtraHeaders(v []string)`

SetSecretExtraHeaders sets SecretExtraHeaders field to given value.

### HasSecretExtraHeaders

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) HasSecretExtraHeaders() bool`

HasSecretExtraHeaders returns a boolean if a field has been set.

### GetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetSecretRef() string`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetSecretRefOk() (*string, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) SetSecretRef(v string)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


