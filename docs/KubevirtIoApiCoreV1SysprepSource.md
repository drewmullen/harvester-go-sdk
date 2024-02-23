# KubevirtIoApiCoreV1SysprepSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMap** | Pointer to [**K8sIoV1LocalObjectReference**](K8sIoV1LocalObjectReference.md) |  | [optional] 
**Secret** | Pointer to [**K8sIoV1LocalObjectReference**](K8sIoV1LocalObjectReference.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1SysprepSource

`func NewKubevirtIoApiCoreV1SysprepSource() *KubevirtIoApiCoreV1SysprepSource`

NewKubevirtIoApiCoreV1SysprepSource instantiates a new KubevirtIoApiCoreV1SysprepSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1SysprepSourceWithDefaults

`func NewKubevirtIoApiCoreV1SysprepSourceWithDefaults() *KubevirtIoApiCoreV1SysprepSource`

NewKubevirtIoApiCoreV1SysprepSourceWithDefaults instantiates a new KubevirtIoApiCoreV1SysprepSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMap

`func (o *KubevirtIoApiCoreV1SysprepSource) GetConfigMap() K8sIoV1LocalObjectReference`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *KubevirtIoApiCoreV1SysprepSource) GetConfigMapOk() (*K8sIoV1LocalObjectReference, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *KubevirtIoApiCoreV1SysprepSource) SetConfigMap(v K8sIoV1LocalObjectReference)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *KubevirtIoApiCoreV1SysprepSource) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetSecret

`func (o *KubevirtIoApiCoreV1SysprepSource) GetSecret() K8sIoV1LocalObjectReference`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *KubevirtIoApiCoreV1SysprepSource) GetSecretOk() (*K8sIoV1LocalObjectReference, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *KubevirtIoApiCoreV1SysprepSource) SetSecret(v K8sIoV1LocalObjectReference)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *KubevirtIoApiCoreV1SysprepSource) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


