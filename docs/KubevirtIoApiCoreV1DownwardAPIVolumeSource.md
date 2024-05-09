# KubevirtIoApiCoreV1DownwardAPIVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]K8sIoV1DownwardAPIVolumeFile**](K8sIoV1DownwardAPIVolumeFile.md) | Fields is a list of downward API volume file | [optional] 
**VolumeLabel** | Pointer to **string** | The volume label of the resulting disk inside the VMI. Different bootstrapping mechanisms require different values. Typical values are \&quot;cidata\&quot; (cloud-init), \&quot;config-2\&quot; (cloud-init) or \&quot;OEMDRV\&quot; (kickstart). | [optional] 

## Methods

### NewKubevirtIoApiCoreV1DownwardAPIVolumeSource

`func NewKubevirtIoApiCoreV1DownwardAPIVolumeSource() *KubevirtIoApiCoreV1DownwardAPIVolumeSource`

NewKubevirtIoApiCoreV1DownwardAPIVolumeSource instantiates a new KubevirtIoApiCoreV1DownwardAPIVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DownwardAPIVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1DownwardAPIVolumeSourceWithDefaults() *KubevirtIoApiCoreV1DownwardAPIVolumeSource`

NewKubevirtIoApiCoreV1DownwardAPIVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1DownwardAPIVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetFields() []K8sIoV1DownwardAPIVolumeFile`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetFieldsOk() (*[]K8sIoV1DownwardAPIVolumeFile, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) SetFields(v []K8sIoV1DownwardAPIVolumeFile)`

SetFields sets Fields field to given value.

### HasFields

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetVolumeLabel

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetVolumeLabel() string`

GetVolumeLabel returns the VolumeLabel field if non-nil, zero value otherwise.

### GetVolumeLabelOk

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) GetVolumeLabelOk() (*string, bool)`

GetVolumeLabelOk returns a tuple with the VolumeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLabel

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) SetVolumeLabel(v string)`

SetVolumeLabel sets VolumeLabel field to given value.

### HasVolumeLabel

`func (o *KubevirtIoApiCoreV1DownwardAPIVolumeSource) HasVolumeLabel() bool`

HasVolumeLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


