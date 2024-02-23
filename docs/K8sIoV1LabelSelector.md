# K8sIoV1LabelSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchExpressions** | Pointer to [**[]K8sIoV1LabelSelectorRequirement**](K8sIoV1LabelSelectorRequirement.md) |  | [optional] 
**MatchLabels** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewK8sIoV1LabelSelector

`func NewK8sIoV1LabelSelector() *K8sIoV1LabelSelector`

NewK8sIoV1LabelSelector instantiates a new K8sIoV1LabelSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1LabelSelectorWithDefaults

`func NewK8sIoV1LabelSelectorWithDefaults() *K8sIoV1LabelSelector`

NewK8sIoV1LabelSelectorWithDefaults instantiates a new K8sIoV1LabelSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchExpressions

`func (o *K8sIoV1LabelSelector) GetMatchExpressions() []K8sIoV1LabelSelectorRequirement`

GetMatchExpressions returns the MatchExpressions field if non-nil, zero value otherwise.

### GetMatchExpressionsOk

`func (o *K8sIoV1LabelSelector) GetMatchExpressionsOk() (*[]K8sIoV1LabelSelectorRequirement, bool)`

GetMatchExpressionsOk returns a tuple with the MatchExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchExpressions

`func (o *K8sIoV1LabelSelector) SetMatchExpressions(v []K8sIoV1LabelSelectorRequirement)`

SetMatchExpressions sets MatchExpressions field to given value.

### HasMatchExpressions

`func (o *K8sIoV1LabelSelector) HasMatchExpressions() bool`

HasMatchExpressions returns a boolean if a field has been set.

### GetMatchLabels

`func (o *K8sIoV1LabelSelector) GetMatchLabels() map[string]string`

GetMatchLabels returns the MatchLabels field if non-nil, zero value otherwise.

### GetMatchLabelsOk

`func (o *K8sIoV1LabelSelector) GetMatchLabelsOk() (*map[string]string, bool)`

GetMatchLabelsOk returns a tuple with the MatchLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLabels

`func (o *K8sIoV1LabelSelector) SetMatchLabels(v map[string]string)`

SetMatchLabels sets MatchLabels field to given value.

### HasMatchLabels

`func (o *K8sIoV1LabelSelector) HasMatchLabels() bool`

HasMatchLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


