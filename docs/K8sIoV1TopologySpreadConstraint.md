# K8sIoV1TopologySpreadConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) |  | [optional] 
**MatchLabelKeys** | Pointer to **[]string** |  | [optional] 
**MaxSkew** | **int32** |  | [default to 0]
**MinDomains** | Pointer to **int32** |  | [optional] 
**NodeAffinityPolicy** | Pointer to **string** |  | [optional] 
**NodeTaintsPolicy** | Pointer to **string** |  | [optional] 
**TopologyKey** | **string** |  | [default to ""]
**WhenUnsatisfiable** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1TopologySpreadConstraint

`func NewK8sIoV1TopologySpreadConstraint(maxSkew int32, topologyKey string, whenUnsatisfiable string, ) *K8sIoV1TopologySpreadConstraint`

NewK8sIoV1TopologySpreadConstraint instantiates a new K8sIoV1TopologySpreadConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1TopologySpreadConstraintWithDefaults

`func NewK8sIoV1TopologySpreadConstraintWithDefaults() *K8sIoV1TopologySpreadConstraint`

NewK8sIoV1TopologySpreadConstraintWithDefaults instantiates a new K8sIoV1TopologySpreadConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelector

`func (o *K8sIoV1TopologySpreadConstraint) GetLabelSelector() K8sIoV1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *K8sIoV1TopologySpreadConstraint) GetLabelSelectorOk() (*K8sIoV1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *K8sIoV1TopologySpreadConstraint) SetLabelSelector(v K8sIoV1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *K8sIoV1TopologySpreadConstraint) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetMatchLabelKeys

`func (o *K8sIoV1TopologySpreadConstraint) GetMatchLabelKeys() []string`

GetMatchLabelKeys returns the MatchLabelKeys field if non-nil, zero value otherwise.

### GetMatchLabelKeysOk

`func (o *K8sIoV1TopologySpreadConstraint) GetMatchLabelKeysOk() (*[]string, bool)`

GetMatchLabelKeysOk returns a tuple with the MatchLabelKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLabelKeys

`func (o *K8sIoV1TopologySpreadConstraint) SetMatchLabelKeys(v []string)`

SetMatchLabelKeys sets MatchLabelKeys field to given value.

### HasMatchLabelKeys

`func (o *K8sIoV1TopologySpreadConstraint) HasMatchLabelKeys() bool`

HasMatchLabelKeys returns a boolean if a field has been set.

### GetMaxSkew

`func (o *K8sIoV1TopologySpreadConstraint) GetMaxSkew() int32`

GetMaxSkew returns the MaxSkew field if non-nil, zero value otherwise.

### GetMaxSkewOk

`func (o *K8sIoV1TopologySpreadConstraint) GetMaxSkewOk() (*int32, bool)`

GetMaxSkewOk returns a tuple with the MaxSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkew

`func (o *K8sIoV1TopologySpreadConstraint) SetMaxSkew(v int32)`

SetMaxSkew sets MaxSkew field to given value.


### GetMinDomains

`func (o *K8sIoV1TopologySpreadConstraint) GetMinDomains() int32`

GetMinDomains returns the MinDomains field if non-nil, zero value otherwise.

### GetMinDomainsOk

`func (o *K8sIoV1TopologySpreadConstraint) GetMinDomainsOk() (*int32, bool)`

GetMinDomainsOk returns a tuple with the MinDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDomains

`func (o *K8sIoV1TopologySpreadConstraint) SetMinDomains(v int32)`

SetMinDomains sets MinDomains field to given value.

### HasMinDomains

`func (o *K8sIoV1TopologySpreadConstraint) HasMinDomains() bool`

HasMinDomains returns a boolean if a field has been set.

### GetNodeAffinityPolicy

`func (o *K8sIoV1TopologySpreadConstraint) GetNodeAffinityPolicy() string`

GetNodeAffinityPolicy returns the NodeAffinityPolicy field if non-nil, zero value otherwise.

### GetNodeAffinityPolicyOk

`func (o *K8sIoV1TopologySpreadConstraint) GetNodeAffinityPolicyOk() (*string, bool)`

GetNodeAffinityPolicyOk returns a tuple with the NodeAffinityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinityPolicy

`func (o *K8sIoV1TopologySpreadConstraint) SetNodeAffinityPolicy(v string)`

SetNodeAffinityPolicy sets NodeAffinityPolicy field to given value.

### HasNodeAffinityPolicy

`func (o *K8sIoV1TopologySpreadConstraint) HasNodeAffinityPolicy() bool`

HasNodeAffinityPolicy returns a boolean if a field has been set.

### GetNodeTaintsPolicy

`func (o *K8sIoV1TopologySpreadConstraint) GetNodeTaintsPolicy() string`

GetNodeTaintsPolicy returns the NodeTaintsPolicy field if non-nil, zero value otherwise.

### GetNodeTaintsPolicyOk

`func (o *K8sIoV1TopologySpreadConstraint) GetNodeTaintsPolicyOk() (*string, bool)`

GetNodeTaintsPolicyOk returns a tuple with the NodeTaintsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeTaintsPolicy

`func (o *K8sIoV1TopologySpreadConstraint) SetNodeTaintsPolicy(v string)`

SetNodeTaintsPolicy sets NodeTaintsPolicy field to given value.

### HasNodeTaintsPolicy

`func (o *K8sIoV1TopologySpreadConstraint) HasNodeTaintsPolicy() bool`

HasNodeTaintsPolicy returns a boolean if a field has been set.

### GetTopologyKey

`func (o *K8sIoV1TopologySpreadConstraint) GetTopologyKey() string`

GetTopologyKey returns the TopologyKey field if non-nil, zero value otherwise.

### GetTopologyKeyOk

`func (o *K8sIoV1TopologySpreadConstraint) GetTopologyKeyOk() (*string, bool)`

GetTopologyKeyOk returns a tuple with the TopologyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKey

`func (o *K8sIoV1TopologySpreadConstraint) SetTopologyKey(v string)`

SetTopologyKey sets TopologyKey field to given value.


### GetWhenUnsatisfiable

`func (o *K8sIoV1TopologySpreadConstraint) GetWhenUnsatisfiable() string`

GetWhenUnsatisfiable returns the WhenUnsatisfiable field if non-nil, zero value otherwise.

### GetWhenUnsatisfiableOk

`func (o *K8sIoV1TopologySpreadConstraint) GetWhenUnsatisfiableOk() (*string, bool)`

GetWhenUnsatisfiableOk returns a tuple with the WhenUnsatisfiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenUnsatisfiable

`func (o *K8sIoV1TopologySpreadConstraint) SetWhenUnsatisfiable(v string)`

SetWhenUnsatisfiable sets WhenUnsatisfiable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


