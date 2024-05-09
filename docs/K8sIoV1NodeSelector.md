# K8sIoV1NodeSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeSelectorTerms** | [**[]K8sIoV1NodeSelectorTerm**](K8sIoV1NodeSelectorTerm.md) | Required. A list of node selector terms. The terms are ORed. | 

## Methods

### NewK8sIoV1NodeSelector

`func NewK8sIoV1NodeSelector(nodeSelectorTerms []K8sIoV1NodeSelectorTerm, ) *K8sIoV1NodeSelector`

NewK8sIoV1NodeSelector instantiates a new K8sIoV1NodeSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1NodeSelectorWithDefaults

`func NewK8sIoV1NodeSelectorWithDefaults() *K8sIoV1NodeSelector`

NewK8sIoV1NodeSelectorWithDefaults instantiates a new K8sIoV1NodeSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeSelectorTerms

`func (o *K8sIoV1NodeSelector) GetNodeSelectorTerms() []K8sIoV1NodeSelectorTerm`

GetNodeSelectorTerms returns the NodeSelectorTerms field if non-nil, zero value otherwise.

### GetNodeSelectorTermsOk

`func (o *K8sIoV1NodeSelector) GetNodeSelectorTermsOk() (*[]K8sIoV1NodeSelectorTerm, bool)`

GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelectorTerms

`func (o *K8sIoV1NodeSelector) SetNodeSelectorTerms(v []K8sIoV1NodeSelectorTerm)`

SetNodeSelectorTerms sets NodeSelectorTerms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


