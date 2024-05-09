# KubevirtIoApiCoreV1ConfigMapVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [optional] 
**Optional** | Pointer to **bool** | Specify whether the ConfigMap or it&#39;s keys must be defined | [optional] 
**VolumeLabel** | Pointer to **string** | The volume label of the resulting disk inside the VMI. Different bootstrapping mechanisms require different values. Typical values are \&quot;cidata\&quot; (cloud-init), \&quot;config-2\&quot; (cloud-init) or \&quot;OEMDRV\&quot; (kickstart). | [optional] 

## Methods

### NewKubevirtIoApiCoreV1ConfigMapVolumeSource

`func NewKubevirtIoApiCoreV1ConfigMapVolumeSource() *KubevirtIoApiCoreV1ConfigMapVolumeSource`

NewKubevirtIoApiCoreV1ConfigMapVolumeSource instantiates a new KubevirtIoApiCoreV1ConfigMapVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ConfigMapVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1ConfigMapVolumeSourceWithDefaults() *KubevirtIoApiCoreV1ConfigMapVolumeSource`

NewKubevirtIoApiCoreV1ConfigMapVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1ConfigMapVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptional

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetOptional() bool`

GetOptional returns the Optional field if non-nil, zero value otherwise.

### GetOptionalOk

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetOptionalOk() (*bool, bool)`

GetOptionalOk returns a tuple with the Optional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptional

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetOptional(v bool)`

SetOptional sets Optional field to given value.

### HasOptional

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasOptional() bool`

HasOptional returns a boolean if a field has been set.

### GetVolumeLabel

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetVolumeLabel() string`

GetVolumeLabel returns the VolumeLabel field if non-nil, zero value otherwise.

### GetVolumeLabelOk

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) GetVolumeLabelOk() (*string, bool)`

GetVolumeLabelOk returns a tuple with the VolumeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLabel

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) SetVolumeLabel(v string)`

SetVolumeLabel sets VolumeLabel field to given value.

### HasVolumeLabel

`func (o *KubevirtIoApiCoreV1ConfigMapVolumeSource) HasVolumeLabel() bool`

HasVolumeLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


