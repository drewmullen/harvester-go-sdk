# K8sIoV1Affinity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAffinity** | Pointer to [**K8sIoV1NodeAffinity**](K8sIoV1NodeAffinity.md) |  | [optional] 
**PodAffinity** | Pointer to [**K8sIoV1PodAffinity**](K8sIoV1PodAffinity.md) |  | [optional] 
**PodAntiAffinity** | Pointer to [**K8sIoV1PodAntiAffinity**](K8sIoV1PodAntiAffinity.md) |  | [optional] 

## Methods

### NewK8sIoV1Affinity

`func NewK8sIoV1Affinity() *K8sIoV1Affinity`

NewK8sIoV1Affinity instantiates a new K8sIoV1Affinity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1AffinityWithDefaults

`func NewK8sIoV1AffinityWithDefaults() *K8sIoV1Affinity`

NewK8sIoV1AffinityWithDefaults instantiates a new K8sIoV1Affinity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAffinity

`func (o *K8sIoV1Affinity) GetNodeAffinity() K8sIoV1NodeAffinity`

GetNodeAffinity returns the NodeAffinity field if non-nil, zero value otherwise.

### GetNodeAffinityOk

`func (o *K8sIoV1Affinity) GetNodeAffinityOk() (*K8sIoV1NodeAffinity, bool)`

GetNodeAffinityOk returns a tuple with the NodeAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinity

`func (o *K8sIoV1Affinity) SetNodeAffinity(v K8sIoV1NodeAffinity)`

SetNodeAffinity sets NodeAffinity field to given value.

### HasNodeAffinity

`func (o *K8sIoV1Affinity) HasNodeAffinity() bool`

HasNodeAffinity returns a boolean if a field has been set.

### GetPodAffinity

`func (o *K8sIoV1Affinity) GetPodAffinity() K8sIoV1PodAffinity`

GetPodAffinity returns the PodAffinity field if non-nil, zero value otherwise.

### GetPodAffinityOk

`func (o *K8sIoV1Affinity) GetPodAffinityOk() (*K8sIoV1PodAffinity, bool)`

GetPodAffinityOk returns a tuple with the PodAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinity

`func (o *K8sIoV1Affinity) SetPodAffinity(v K8sIoV1PodAffinity)`

SetPodAffinity sets PodAffinity field to given value.

### HasPodAffinity

`func (o *K8sIoV1Affinity) HasPodAffinity() bool`

HasPodAffinity returns a boolean if a field has been set.

### GetPodAntiAffinity

`func (o *K8sIoV1Affinity) GetPodAntiAffinity() K8sIoV1PodAntiAffinity`

GetPodAntiAffinity returns the PodAntiAffinity field if non-nil, zero value otherwise.

### GetPodAntiAffinityOk

`func (o *K8sIoV1Affinity) GetPodAntiAffinityOk() (*K8sIoV1PodAntiAffinity, bool)`

GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAntiAffinity

`func (o *K8sIoV1Affinity) SetPodAntiAffinity(v K8sIoV1PodAntiAffinity)`

SetPodAntiAffinity sets PodAntiAffinity field to given value.

### HasPodAntiAffinity

`func (o *K8sIoV1Affinity) HasPodAntiAffinity() bool`

HasPodAntiAffinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


