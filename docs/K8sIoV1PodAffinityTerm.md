# K8sIoV1PodAffinityTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) |  | [optional] 
**NamespaceSelector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) |  | [optional] 
**Namespaces** | Pointer to **[]string** |  | [optional] 
**TopologyKey** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1PodAffinityTerm

`func NewK8sIoV1PodAffinityTerm(topologyKey string, ) *K8sIoV1PodAffinityTerm`

NewK8sIoV1PodAffinityTerm instantiates a new K8sIoV1PodAffinityTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PodAffinityTermWithDefaults

`func NewK8sIoV1PodAffinityTermWithDefaults() *K8sIoV1PodAffinityTerm`

NewK8sIoV1PodAffinityTermWithDefaults instantiates a new K8sIoV1PodAffinityTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelector

`func (o *K8sIoV1PodAffinityTerm) GetLabelSelector() K8sIoV1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *K8sIoV1PodAffinityTerm) GetLabelSelectorOk() (*K8sIoV1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *K8sIoV1PodAffinityTerm) SetLabelSelector(v K8sIoV1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *K8sIoV1PodAffinityTerm) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetNamespaceSelector

`func (o *K8sIoV1PodAffinityTerm) GetNamespaceSelector() K8sIoV1LabelSelector`

GetNamespaceSelector returns the NamespaceSelector field if non-nil, zero value otherwise.

### GetNamespaceSelectorOk

`func (o *K8sIoV1PodAffinityTerm) GetNamespaceSelectorOk() (*K8sIoV1LabelSelector, bool)`

GetNamespaceSelectorOk returns a tuple with the NamespaceSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceSelector

`func (o *K8sIoV1PodAffinityTerm) SetNamespaceSelector(v K8sIoV1LabelSelector)`

SetNamespaceSelector sets NamespaceSelector field to given value.

### HasNamespaceSelector

`func (o *K8sIoV1PodAffinityTerm) HasNamespaceSelector() bool`

HasNamespaceSelector returns a boolean if a field has been set.

### GetNamespaces

`func (o *K8sIoV1PodAffinityTerm) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *K8sIoV1PodAffinityTerm) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *K8sIoV1PodAffinityTerm) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *K8sIoV1PodAffinityTerm) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetTopologyKey

`func (o *K8sIoV1PodAffinityTerm) GetTopologyKey() string`

GetTopologyKey returns the TopologyKey field if non-nil, zero value otherwise.

### GetTopologyKeyOk

`func (o *K8sIoV1PodAffinityTerm) GetTopologyKeyOk() (*string, bool)`

GetTopologyKeyOk returns a tuple with the TopologyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKey

`func (o *K8sIoV1PodAffinityTerm) SetTopologyKey(v string)`

SetTopologyKey sets TopologyKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


