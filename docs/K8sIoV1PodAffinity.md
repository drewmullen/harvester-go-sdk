# K8sIoV1PodAffinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreferredDuringSchedulingIgnoredDuringExecution** | Pointer to [**[]K8sIoV1WeightedPodAffinityTerm**](K8sIoV1WeightedPodAffinityTerm.md) |  | [optional] 
**RequiredDuringSchedulingIgnoredDuringExecution** | Pointer to [**[]K8sIoV1PodAffinityTerm**](K8sIoV1PodAffinityTerm.md) |  | [optional] 

## Methods

### NewK8sIoV1PodAffinity

`func NewK8sIoV1PodAffinity() *K8sIoV1PodAffinity`

NewK8sIoV1PodAffinity instantiates a new K8sIoV1PodAffinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PodAffinityWithDefaults

`func NewK8sIoV1PodAffinityWithDefaults() *K8sIoV1PodAffinity`

NewK8sIoV1PodAffinityWithDefaults instantiates a new K8sIoV1PodAffinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecution() []K8sIoV1WeightedPodAffinityTerm`

GetPreferredDuringSchedulingIgnoredDuringExecution returns the PreferredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetPreferredDuringSchedulingIgnoredDuringExecutionOk

`func (o *K8sIoV1PodAffinity) GetPreferredDuringSchedulingIgnoredDuringExecutionOk() (*[]K8sIoV1WeightedPodAffinityTerm, bool)`

GetPreferredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the PreferredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) SetPreferredDuringSchedulingIgnoredDuringExecution(v []K8sIoV1WeightedPodAffinityTerm)`

SetPreferredDuringSchedulingIgnoredDuringExecution sets PreferredDuringSchedulingIgnoredDuringExecution field to given value.

### HasPreferredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) HasPreferredDuringSchedulingIgnoredDuringExecution() bool`

HasPreferredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.

### GetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecution() []K8sIoV1PodAffinityTerm`

GetRequiredDuringSchedulingIgnoredDuringExecution returns the RequiredDuringSchedulingIgnoredDuringExecution field if non-nil, zero value otherwise.

### GetRequiredDuringSchedulingIgnoredDuringExecutionOk

`func (o *K8sIoV1PodAffinity) GetRequiredDuringSchedulingIgnoredDuringExecutionOk() (*[]K8sIoV1PodAffinityTerm, bool)`

GetRequiredDuringSchedulingIgnoredDuringExecutionOk returns a tuple with the RequiredDuringSchedulingIgnoredDuringExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) SetRequiredDuringSchedulingIgnoredDuringExecution(v []K8sIoV1PodAffinityTerm)`

SetRequiredDuringSchedulingIgnoredDuringExecution sets RequiredDuringSchedulingIgnoredDuringExecution field to given value.

### HasRequiredDuringSchedulingIgnoredDuringExecution

`func (o *K8sIoV1PodAffinity) HasRequiredDuringSchedulingIgnoredDuringExecution() bool`

HasRequiredDuringSchedulingIgnoredDuringExecution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


