# K8sIoV1TopologySpreadConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**K8sIoV1LabelSelector**](K8sIoV1LabelSelector.md) | LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain. | [optional] 
**MatchLabelKeys** | Pointer to **[]string** | MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don&#39;t exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. | [optional] 
**MaxSkew** | **int32** | MaxSkew describes the degree to which pods may be unevenly distributed. When &#x60;whenUnsatisfiable&#x3D;DoNotSchedule&#x60;, it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When &#x60;whenUnsatisfiable&#x3D;ScheduleAnyway&#x60;, it is used to give higher precedence to topologies that satisfy it. It&#39;s a required field. Default value is 1 and 0 is not allowed. | [default to 0]
**MinDomains** | Pointer to **int32** | MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats \&quot;global minimum\&quot; as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won&#39;t schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so \&quot;global minimum\&quot; is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default). | [optional] 
**NodeAffinityPolicy** | Pointer to **string** | NodeAffinityPolicy indicates how we will treat Pod&#39;s nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag. | [optional] 
**NodeTaintsPolicy** | Pointer to **string** | NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag. | [optional] 
**TopologyKey** | **string** | TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each &lt;key, value&gt; as a \&quot;bucket\&quot;, and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is \&quot;kubernetes.io/hostname\&quot;, each Node is a domain of that topology. And, if TopologyKey is \&quot;topology.kubernetes.io/zone\&quot;, each zone is a domain of that topology. It&#39;s a required field. | [default to ""]
**WhenUnsatisfiable** | **string** | WhenUnsatisfiable indicates how to deal with a pod if it doesn&#39;t satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered \&quot;Unsatisfiable\&quot; for an incoming pod if and only if every possible node assignment for that pod would violate \&quot;MaxSkew\&quot; on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won&#39;t make it *more* imbalanced. It&#39;s a required field.  Possible enum values:  - &#x60;\&quot;DoNotSchedule\&quot;&#x60; instructs the scheduler not to schedule the pod when constraints are not satisfied.  - &#x60;\&quot;ScheduleAnyway\&quot;&#x60; instructs the scheduler to schedule the pod even if constraints are not satisfied. | [default to ""]

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


