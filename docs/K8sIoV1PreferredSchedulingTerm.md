# K8sIoV1PreferredSchedulingTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preference** | [**K8sIoV1NodeSelectorTerm**](K8sIoV1NodeSelectorTerm.md) |  | 
**Weight** | **int32** |  | [default to 0]

## Methods

### NewK8sIoV1PreferredSchedulingTerm

`func NewK8sIoV1PreferredSchedulingTerm(preference K8sIoV1NodeSelectorTerm, weight int32, ) *K8sIoV1PreferredSchedulingTerm`

NewK8sIoV1PreferredSchedulingTerm instantiates a new K8sIoV1PreferredSchedulingTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PreferredSchedulingTermWithDefaults

`func NewK8sIoV1PreferredSchedulingTermWithDefaults() *K8sIoV1PreferredSchedulingTerm`

NewK8sIoV1PreferredSchedulingTermWithDefaults instantiates a new K8sIoV1PreferredSchedulingTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreference

`func (o *K8sIoV1PreferredSchedulingTerm) GetPreference() K8sIoV1NodeSelectorTerm`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *K8sIoV1PreferredSchedulingTerm) GetPreferenceOk() (*K8sIoV1NodeSelectorTerm, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *K8sIoV1PreferredSchedulingTerm) SetPreference(v K8sIoV1NodeSelectorTerm)`

SetPreference sets Preference field to given value.


### GetWeight

`func (o *K8sIoV1PreferredSchedulingTerm) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *K8sIoV1PreferredSchedulingTerm) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *K8sIoV1PreferredSchedulingTerm) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


