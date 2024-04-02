# K8sIoV1WeightedPodAffinityTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodAffinityTerm** | [**K8sIoV1PodAffinityTerm**](K8sIoV1PodAffinityTerm.md) |  | [default to {}]
**Weight** | **int32** |  | [default to 0]

## Methods

### NewK8sIoV1WeightedPodAffinityTerm

`func NewK8sIoV1WeightedPodAffinityTerm(podAffinityTerm K8sIoV1PodAffinityTerm, weight int32, ) *K8sIoV1WeightedPodAffinityTerm`

NewK8sIoV1WeightedPodAffinityTerm instantiates a new K8sIoV1WeightedPodAffinityTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1WeightedPodAffinityTermWithDefaults

`func NewK8sIoV1WeightedPodAffinityTermWithDefaults() *K8sIoV1WeightedPodAffinityTerm`

NewK8sIoV1WeightedPodAffinityTermWithDefaults instantiates a new K8sIoV1WeightedPodAffinityTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodAffinityTerm

`func (o *K8sIoV1WeightedPodAffinityTerm) GetPodAffinityTerm() K8sIoV1PodAffinityTerm`

GetPodAffinityTerm returns the PodAffinityTerm field if non-nil, zero value otherwise.

### GetPodAffinityTermOk

`func (o *K8sIoV1WeightedPodAffinityTerm) GetPodAffinityTermOk() (*K8sIoV1PodAffinityTerm, bool)`

GetPodAffinityTermOk returns a tuple with the PodAffinityTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinityTerm

`func (o *K8sIoV1WeightedPodAffinityTerm) SetPodAffinityTerm(v K8sIoV1PodAffinityTerm)`

SetPodAffinityTerm sets PodAffinityTerm field to given value.


### GetWeight

`func (o *K8sIoV1WeightedPodAffinityTerm) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *K8sIoV1WeightedPodAffinityTerm) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *K8sIoV1WeightedPodAffinityTerm) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


