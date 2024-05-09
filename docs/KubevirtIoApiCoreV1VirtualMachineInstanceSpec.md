# KubevirtIoApiCoreV1VirtualMachineInstanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessCredentials** | Pointer to [**[]KubevirtIoApiCoreV1AccessCredential**](KubevirtIoApiCoreV1AccessCredential.md) | Specifies a set of public keys to inject into the vm guest | [optional] 
**Affinity** | Pointer to [**K8sIoV1Affinity**](K8sIoV1Affinity.md) | If affinity is specifies, obey all the affinity rules | [optional] 
**Architecture** | Pointer to **string** | Specifies the architecture of the vm guest you are attempting to run. Defaults to the compiled architecture of the KubeVirt components | [optional] 
**DnsConfig** | Pointer to [**K8sIoV1PodDNSConfig**](K8sIoV1PodDNSConfig.md) | Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy. | [optional] 
**DnsPolicy** | Pointer to **string** | Set DNS policy for the pod. Defaults to \&quot;ClusterFirst\&quot;. Valid values are &#39;ClusterFirstWithHostNet&#39;, &#39;ClusterFirst&#39;, &#39;Default&#39; or &#39;None&#39;. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to &#39;ClusterFirstWithHostNet&#39;.  Possible enum values:  - &#x60;\&quot;ClusterFirst\&quot;&#x60; indicates that the pod should use cluster DNS first unless hostNetwork is true, if it is available, then fall back on the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;ClusterFirstWithHostNet\&quot;&#x60; indicates that the pod should use cluster DNS first, if it is available, then fall back on the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;Default\&quot;&#x60; indicates that the pod should use the default (as determined by kubelet) DNS settings.  - &#x60;\&quot;None\&quot;&#x60; indicates that the pod should use empty DNS settings. DNS parameters such as nameservers and search paths should be defined via DNSConfig. | [optional] 
**Domain** | [**KubevirtIoApiCoreV1DomainSpec**](KubevirtIoApiCoreV1DomainSpec.md) | Specification of the desired behavior of the VirtualMachineInstance on the host. | [default to {}]
**EvictionStrategy** | Pointer to **string** | EvictionStrategy can be set to \&quot;LiveMigrate\&quot; if the VirtualMachineInstance should be migrated instead of shut-off in case of a node drain. | [optional] 
**Hostname** | Pointer to **string** | Specifies the hostname of the vmi If not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly. | [optional] 
**LivenessProbe** | Pointer to [**KubevirtIoApiCoreV1Probe**](KubevirtIoApiCoreV1Probe.md) | Periodic probe of VirtualMachineInstance liveness. VirtualmachineInstances will be stopped if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 
**Networks** | Pointer to [**[]KubevirtIoApiCoreV1Network**](KubevirtIoApiCoreV1Network.md) | List of networks that can be attached to a vm&#39;s virtual interface. | [optional] 
**NodeSelector** | Pointer to **map[string]string** | NodeSelector is a selector which must be true for the vmi to fit on a node. Selector which must match a node&#39;s labels for the vmi to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ | [optional] 
**PriorityClassName** | Pointer to **string** | If specified, indicates the pod&#39;s priority. If not specified, the pod priority will be default or zero if there is no default. | [optional] 
**ReadinessProbe** | Pointer to [**KubevirtIoApiCoreV1Probe**](KubevirtIoApiCoreV1Probe.md) | Periodic probe of VirtualMachineInstance service readiness. VirtualmachineInstances will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes | [optional] 
**SchedulerName** | Pointer to **string** | If specified, the VMI will be dispatched by specified scheduler. If not specified, the VMI will be dispatched by default scheduler. | [optional] 
**StartStrategy** | Pointer to **string** | StartStrategy can be set to \&quot;Paused\&quot; if Virtual Machine should be started in paused state. | [optional] 
**Subdomain** | Pointer to **string** | If specified, the fully qualified vmi hostname will be \&quot;&lt;hostname&gt;.&lt;subdomain&gt;.&lt;pod namespace&gt;.svc.&lt;cluster domain&gt;\&quot;. If not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi, no matter if the vmi itself can pick up a hostname. | [optional] 
**TerminationGracePeriodSeconds** | Pointer to **int64** | Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated. | [optional] 
**Tolerations** | Pointer to [**[]K8sIoV1Toleration**](K8sIoV1Toleration.md) | If toleration is specified, obey all the toleration rules. | [optional] 
**TopologySpreadConstraints** | Pointer to [**[]K8sIoV1TopologySpreadConstraint**](K8sIoV1TopologySpreadConstraint.md) | TopologySpreadConstraints describes how a group of VMIs will be spread across a given topology domains. K8s scheduler will schedule VMI pods in a way which abides by the constraints. | [optional] 
**Volumes** | Pointer to [**[]KubevirtIoApiCoreV1Volume**](KubevirtIoApiCoreV1Volume.md) | List of volumes that can be mounted by disks belonging to the vmi. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceSpec

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceSpec(domain KubevirtIoApiCoreV1DomainSpec, ) *KubevirtIoApiCoreV1VirtualMachineInstanceSpec`

NewKubevirtIoApiCoreV1VirtualMachineInstanceSpec instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceSpecWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceSpecWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceSpec`

NewKubevirtIoApiCoreV1VirtualMachineInstanceSpecWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessCredentials

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAccessCredentials() []KubevirtIoApiCoreV1AccessCredential`

GetAccessCredentials returns the AccessCredentials field if non-nil, zero value otherwise.

### GetAccessCredentialsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAccessCredentialsOk() (*[]KubevirtIoApiCoreV1AccessCredential, bool)`

GetAccessCredentialsOk returns a tuple with the AccessCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessCredentials

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetAccessCredentials(v []KubevirtIoApiCoreV1AccessCredential)`

SetAccessCredentials sets AccessCredentials field to given value.

### HasAccessCredentials

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasAccessCredentials() bool`

HasAccessCredentials returns a boolean if a field has been set.

### GetAffinity

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAffinity() K8sIoV1Affinity`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAffinityOk() (*K8sIoV1Affinity, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetAffinity(v K8sIoV1Affinity)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetArchitecture

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetDnsConfig

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsConfig() K8sIoV1PodDNSConfig`

GetDnsConfig returns the DnsConfig field if non-nil, zero value otherwise.

### GetDnsConfigOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsConfigOk() (*K8sIoV1PodDNSConfig, bool)`

GetDnsConfigOk returns a tuple with the DnsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsConfig

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDnsConfig(v K8sIoV1PodDNSConfig)`

SetDnsConfig sets DnsConfig field to given value.

### HasDnsConfig

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasDnsConfig() bool`

HasDnsConfig returns a boolean if a field has been set.

### GetDnsPolicy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsPolicy() string`

GetDnsPolicy returns the DnsPolicy field if non-nil, zero value otherwise.

### GetDnsPolicyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsPolicyOk() (*string, bool)`

GetDnsPolicyOk returns a tuple with the DnsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPolicy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDnsPolicy(v string)`

SetDnsPolicy sets DnsPolicy field to given value.

### HasDnsPolicy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasDnsPolicy() bool`

HasDnsPolicy returns a boolean if a field has been set.

### GetDomain

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDomain() KubevirtIoApiCoreV1DomainSpec`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDomainOk() (*KubevirtIoApiCoreV1DomainSpec, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDomain(v KubevirtIoApiCoreV1DomainSpec)`

SetDomain sets Domain field to given value.


### GetEvictionStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetEvictionStrategy() string`

GetEvictionStrategy returns the EvictionStrategy field if non-nil, zero value otherwise.

### GetEvictionStrategyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetEvictionStrategyOk() (*string, bool)`

GetEvictionStrategyOk returns a tuple with the EvictionStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvictionStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetEvictionStrategy(v string)`

SetEvictionStrategy sets EvictionStrategy field to given value.

### HasEvictionStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasEvictionStrategy() bool`

HasEvictionStrategy returns a boolean if a field has been set.

### GetHostname

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLivenessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetLivenessProbe() KubevirtIoApiCoreV1Probe`

GetLivenessProbe returns the LivenessProbe field if non-nil, zero value otherwise.

### GetLivenessProbeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetLivenessProbeOk() (*KubevirtIoApiCoreV1Probe, bool)`

GetLivenessProbeOk returns a tuple with the LivenessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLivenessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetLivenessProbe(v KubevirtIoApiCoreV1Probe)`

SetLivenessProbe sets LivenessProbe field to given value.

### HasLivenessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasLivenessProbe() bool`

HasLivenessProbe returns a boolean if a field has been set.

### GetNetworks

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNetworks() []KubevirtIoApiCoreV1Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNetworksOk() (*[]KubevirtIoApiCoreV1Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetNetworks(v []KubevirtIoApiCoreV1Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetNodeSelector

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNodeSelector() map[string]string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNodeSelectorOk() (*map[string]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetNodeSelector(v map[string]string)`

SetNodeSelector sets NodeSelector field to given value.

### HasNodeSelector

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasNodeSelector() bool`

HasNodeSelector returns a boolean if a field has been set.

### GetPriorityClassName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.

### HasPriorityClassName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasPriorityClassName() bool`

HasPriorityClassName returns a boolean if a field has been set.

### GetReadinessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetReadinessProbe() KubevirtIoApiCoreV1Probe`

GetReadinessProbe returns the ReadinessProbe field if non-nil, zero value otherwise.

### GetReadinessProbeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetReadinessProbeOk() (*KubevirtIoApiCoreV1Probe, bool)`

GetReadinessProbeOk returns a tuple with the ReadinessProbe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadinessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetReadinessProbe(v KubevirtIoApiCoreV1Probe)`

SetReadinessProbe sets ReadinessProbe field to given value.

### HasReadinessProbe

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasReadinessProbe() bool`

HasReadinessProbe returns a boolean if a field has been set.

### GetSchedulerName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSchedulerName() string`

GetSchedulerName returns the SchedulerName field if non-nil, zero value otherwise.

### GetSchedulerNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSchedulerNameOk() (*string, bool)`

GetSchedulerNameOk returns a tuple with the SchedulerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetSchedulerName(v string)`

SetSchedulerName sets SchedulerName field to given value.

### HasSchedulerName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasSchedulerName() bool`

HasSchedulerName returns a boolean if a field has been set.

### GetStartStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetStartStrategy() string`

GetStartStrategy returns the StartStrategy field if non-nil, zero value otherwise.

### GetStartStrategyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetStartStrategyOk() (*string, bool)`

GetStartStrategyOk returns a tuple with the StartStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetStartStrategy(v string)`

SetStartStrategy sets StartStrategy field to given value.

### HasStartStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasStartStrategy() bool`

HasStartStrategy returns a boolean if a field has been set.

### GetSubdomain

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetTerminationGracePeriodSeconds

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTerminationGracePeriodSeconds() int64`

GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field if non-nil, zero value otherwise.

### GetTerminationGracePeriodSecondsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTerminationGracePeriodSecondsOk() (*int64, bool)`

GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationGracePeriodSeconds

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTerminationGracePeriodSeconds(v int64)`

SetTerminationGracePeriodSeconds sets TerminationGracePeriodSeconds field to given value.

### HasTerminationGracePeriodSeconds

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTerminationGracePeriodSeconds() bool`

HasTerminationGracePeriodSeconds returns a boolean if a field has been set.

### GetTolerations

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTolerations() []K8sIoV1Toleration`

GetTolerations returns the Tolerations field if non-nil, zero value otherwise.

### GetTolerationsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTolerationsOk() (*[]K8sIoV1Toleration, bool)`

GetTolerationsOk returns a tuple with the Tolerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerations

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTolerations(v []K8sIoV1Toleration)`

SetTolerations sets Tolerations field to given value.

### HasTolerations

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTolerations() bool`

HasTolerations returns a boolean if a field has been set.

### GetTopologySpreadConstraints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTopologySpreadConstraints() []K8sIoV1TopologySpreadConstraint`

GetTopologySpreadConstraints returns the TopologySpreadConstraints field if non-nil, zero value otherwise.

### GetTopologySpreadConstraintsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTopologySpreadConstraintsOk() (*[]K8sIoV1TopologySpreadConstraint, bool)`

GetTopologySpreadConstraintsOk returns a tuple with the TopologySpreadConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologySpreadConstraints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTopologySpreadConstraints(v []K8sIoV1TopologySpreadConstraint)`

SetTopologySpreadConstraints sets TopologySpreadConstraints field to given value.

### HasTopologySpreadConstraints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTopologySpreadConstraints() bool`

HasTopologySpreadConstraints returns a boolean if a field has been set.

### GetVolumes

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetVolumes() []KubevirtIoApiCoreV1Volume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetVolumesOk() (*[]KubevirtIoApiCoreV1Volume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetVolumes(v []KubevirtIoApiCoreV1Volume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


