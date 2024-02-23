# K8sIoV1NodeAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredDuringSchedulingIgnoredDuringExecution** | Pointer to [**[]K8sIoV1PreferredSchedulingTerm**](K8sIoV1PreferredSchedulingTerm.md) |  | [optional] 
**RequiredDuringSchedulingIgnoredDuringExecution** | Pointer to [**K8sIoV1NodeSelector**](K8sIoV1NodeSelector.md) |  | [optional] 

## Methods

### NewK8sIoV1NodeAffinity

`func NewK8sIoV1NodeAffinity() *K8sIoV1NodeAffinity`

NewK8sIoV1NodeAffinity instantiates a new K8sIoV1NodeAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1NodeAffinityWithDefaults

`func NewK8sIoV1NodeAffinityWithDefaults() *K8sIoV1NodeAffinity`

NewK8sIoV1NodeAffinityWithDefaults instantiates a new K8sIoV1NodeAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []K8sIoV1PreferredSchedulingTerm`

GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetPreferredDuringSchedulingIgnoredDuringExecutionOk

`func (o *K8sIoV1NodeAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() (*[]K8sIoV1PreferredSchedulingTerm, bool)`

GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []K8sIoV1PreferredSchedulingTerm)`

SetPreferredDuringSchedulingIgnoredDuringExecution sets PreferredDuringSchedulingIgnoredDuringExecution field to given value.

### HasPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool`

HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.

### GetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() K8sIoV1NodeSelector`

GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetRequiredDuringSchedulingIgnoredDuringExecutionOk

`func (o *K8sIoV1NodeAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*K8sIoV1NodeSelector, bool)`

GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v K8sIoV1NodeSelector)`

SetRequiredDuringSchedulingIgnoredDuringExecution sets RequiredDuringSchedulingIgnoredDuringExecution field to given value.

### HasRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1NodeAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool`

HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


