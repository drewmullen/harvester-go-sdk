# K8sIoV1NodeSelectorTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]K8sIoV1NodeSelectorRequirement**](K8sIoV1NodeSelectorRequirement.md) |  | [optional] 
**MatchFields** | Pointer to [**[]K8sIoV1NodeSelectorRequirement**](K8sIoV1NodeSelectorRequirement.md) |  | [optional] 

## Methods

### NewK8sIoV1NodeSelectorTerm

`func NewK8sIoV1NodeSelectorTerm() *K8sIoV1NodeSelectorTerm`

NewK8sIoV1NodeSelectorTerm instantiates a new K8sIoV1NodeSelectorTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1NodeSelectorTermWithDefaults

`func NewK8sIoV1NodeSelectorTermWithDefaults() *K8sIoV1NodeSelectorTerm`

NewK8sIoV1NodeSelectorTermWithDefaults instantiates a new K8sIoV1NodeSelectorTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *K8sIoV1NodeSelectorTerm) GetMatchExpressions() []K8sIoV1NodeSelectorRequirement`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *K8sIoV1NodeSelectorTerm) GetMatchExpressionsOk() (*[]K8sIoV1NodeSelectorRequirement, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *K8sIoV1NodeSelectorTerm) SetMatchExpressions(v []K8sIoV1NodeSelectorRequirement)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *K8sIoV1NodeSelectorTerm) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.

### GetMatchFields

`func (o *K8sIoV1NodeSelectorTerm) GetMatchFields() []K8sIoV1NodeSelectorRequirement`

GetMatchFields returns the MatchFields field if non-nil, zero value otherwise.

### GetMatchFieldsOk

`func (o *K8sIoV1NodeSelectorTerm) GetMatchFieldsOk() (*[]K8sIoV1NodeSelectorRequirement, bool)`

GetMatchFieldsOk returns a tuple with the MatchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchFields

`func (o *K8sIoV1NodeSelectorTerm) SetMatchFields(v []K8sIoV1NodeSelectorRequirement)`

SetMatchFields sets MatchFields field to given value.

### HasMatchFields

`func (o *K8sIoV1NodeSelectorTerm) HasMatchFields() bool`

HasMatchFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


