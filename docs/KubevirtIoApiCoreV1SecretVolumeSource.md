# KubevirtIoApiCoreV1SecretVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Optional** | Pointer to **bool** | Specify whether the Secret or it&#39;s keys must be defined | [optional] 
**SecretName** | Pointer to **string** | Name of the secret in the pod&#39;s namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret | [optional] 
**VolumeLabel** | Pointer to **string** | The volume label of the resulting disk inside the VMI. Different bootstrapping mechanisms require different values. Typical values are \&quot;cidata\&quot; (cloud-init), \&quot;config-2\&quot; (cloud-init) or \&quot;OEMDRV\&quot; (kickstart). | [optional] 

## Methods

### NewKubevirtIoApiCoreV1SecretVolumeSource

`func NewKubevirtIoApiCoreV1SecretVolumeSource() *KubevirtIoApiCoreV1SecretVolumeSource`

NewKubevirtIoApiCoreV1SecretVolumeSource instantiates a new KubevirtIoApiCoreV1SecretVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1SecretVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1SecretVolumeSourceWithDefaults() *KubevirtIoApiCoreV1SecretVolumeSource`

NewKubevirtIoApiCoreV1SecretVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1SecretVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptional

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetSecretName

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.

### GetVolumeLabel

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetVolumeLabel() string`

GetVolumeLabel returns the VolumeLabel field if non-nil, zero value otherwise.

### GetVolumeLabelOk

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) GetVolumeLabelOk() (*string, bool)`

GetVolumeLabelOk returns a tuple with the VolumeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLabel

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) SetVolumeLabel(v string)`

SetVolumeLabel sets VolumeLabel field to given value.

### HasVolumeLabel

`func (o *KubevirtIoApiCoreV1SecretVolumeSource) HasVolumeLabel() bool`

HasVolumeLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


