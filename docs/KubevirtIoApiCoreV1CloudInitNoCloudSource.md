# KubevirtIoApiCoreV1CloudInitNoCloudSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkData** | Pointer to **string** | NetworkData contains NoCloud inline cloud-init networkdata. | [optional] 
**NetworkDataBase64** | Pointer to **string** | NetworkDataBase64 contains NoCloud cloud-init networkdata as a base64 encoded string. | [optional] 
**NetworkDataSecretRef** | Pointer to [**K8sIoV1LocalObjectReference**](K8sIoV1LocalObjectReference.md) | NetworkDataSecretRef references a k8s secret that contains NoCloud networkdata. | [optional] 
**SecretRef** | Pointer to [**K8sIoV1LocalObjectReference**](K8sIoV1LocalObjectReference.md) | UserDataSecretRef references a k8s secret that contains NoCloud userdata. | [optional] 
**UserData** | Pointer to **string** | UserData contains NoCloud inline cloud-init userdata. | [optional] 
**UserDataBase64** | Pointer to **string** | UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1CloudInitNoCloudSource

`func NewKubevirtIoApiCoreV1CloudInitNoCloudSource() *KubevirtIoApiCoreV1CloudInitNoCloudSource`

NewKubevirtIoApiCoreV1CloudInitNoCloudSource instantiates a new KubevirtIoApiCoreV1CloudInitNoCloudSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1CloudInitNoCloudSourceWithDefaults

`func NewKubevirtIoApiCoreV1CloudInitNoCloudSourceWithDefaults() *KubevirtIoApiCoreV1CloudInitNoCloudSource`

NewKubevirtIoApiCoreV1CloudInitNoCloudSourceWithDefaults instantiates a new KubevirtIoApiCoreV1CloudInitNoCloudSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkData() string`

GetNetworkData returns the NetworkData field if non-nil, zero value otherwise.

### GetNetworkDataOk

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkDataOk() (*string, bool)`

GetNetworkDataOk returns a tuple with the NetworkData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetNetworkData(v string)`

SetNetworkData sets NetworkData field to given value.

### HasNetworkData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasNetworkData() bool`

HasNetworkData returns a boolean if a field has been set.

### GetNetworkDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkDataBase64() string`

GetNetworkDataBase64 returns the NetworkDataBase64 field if non-nil, zero value otherwise.

### GetNetworkDataBase64Ok

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkDataBase64Ok() (*string, bool)`

GetNetworkDataBase64Ok returns a tuple with the NetworkDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetNetworkDataBase64(v string)`

SetNetworkDataBase64 sets NetworkDataBase64 field to given value.

### HasNetworkDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasNetworkDataBase64() bool`

HasNetworkDataBase64 returns a boolean if a field has been set.

### GetNetworkDataSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkDataSecretRef() K8sIoV1LocalObjectReference`

GetNetworkDataSecretRef returns the NetworkDataSecretRef field if non-nil, zero value otherwise.

### GetNetworkDataSecretRefOk

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetNetworkDataSecretRefOk() (*K8sIoV1LocalObjectReference, bool)`

GetNetworkDataSecretRefOk returns a tuple with the NetworkDataSecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDataSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetNetworkDataSecretRef(v K8sIoV1LocalObjectReference)`

SetNetworkDataSecretRef sets NetworkDataSecretRef field to given value.

### HasNetworkDataSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasNetworkDataSecretRef() bool`

HasNetworkDataSecretRef returns a boolean if a field has been set.

### GetSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetSecretRef() K8sIoV1LocalObjectReference`

GetSecretRef returns the SecretRef field if non-nil, zero value otherwise.

### GetSecretRefOk

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetSecretRefOk() (*K8sIoV1LocalObjectReference, bool)`

GetSecretRefOk returns a tuple with the SecretRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetSecretRef(v K8sIoV1LocalObjectReference)`

SetSecretRef sets SecretRef field to given value.

### HasSecretRef

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasSecretRef() bool`

HasSecretRef returns a boolean if a field has been set.

### GetUserData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetUserDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetUserDataBase64() string`

GetUserDataBase64 returns the UserDataBase64 field if non-nil, zero value otherwise.

### GetUserDataBase64Ok

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) GetUserDataBase64Ok() (*string, bool)`

GetUserDataBase64Ok returns a tuple with the UserDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) SetUserDataBase64(v string)`

SetUserDataBase64 sets UserDataBase64 field to given value.

### HasUserDataBase64

`func (o *KubevirtIoApiCoreV1CloudInitNoCloudSource) HasUserDataBase64() bool`

HasUserDataBase64 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


