# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecretRef** | Pointer to **string** | SecretRef provides the secret reference needed to access the GCS source | [optional] 
**Url** | **string** | URL is the url of the GCS source | [default to ""]

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS(url string, ) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCSWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCSWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCSWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) GetSecretRef() string`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) GetSecretRefOk() (*string, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) SetSecretRef(v string)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


