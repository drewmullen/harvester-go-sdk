/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceSpec type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceSpec{}

// KubevirtIoApiCoreV1VirtualMachineInstanceSpec struct for KubevirtIoApiCoreV1VirtualMachineInstanceSpec
type KubevirtIoApiCoreV1VirtualMachineInstanceSpec struct {
	AccessCredentials []KubevirtIoApiCoreV1AccessCredential `json:"accessCredentials,omitempty"`
	Affinity *K8sIoV1Affinity `json:"affinity,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	DnsConfig *K8sIoV1PodDNSConfig `json:"dnsConfig,omitempty"`
	DnsPolicy *string `json:"dnsPolicy,omitempty"`
	Domain KubevirtIoApiCoreV1DomainSpec `json:"domain"`
	EvictionStrategy *string `json:"evictionStrategy,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	LivenessProbe *KubevirtIoApiCoreV1Probe `json:"livenessProbe,omitempty"`
	Networks []KubevirtIoApiCoreV1Network `json:"networks,omitempty"`
	NodeSelector *map[string]string `json:"nodeSelector,omitempty"`
	PriorityClassName *string `json:"priorityClassName,omitempty"`
	ReadinessProbe *KubevirtIoApiCoreV1Probe `json:"readinessProbe,omitempty"`
	SchedulerName *string `json:"schedulerName,omitempty"`
	StartStrategy *string `json:"startStrategy,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
	Tolerations []K8sIoV1Toleration `json:"tolerations,omitempty"`
	TopologySpreadConstraints []K8sIoV1TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty"`
	Volumes []KubevirtIoApiCoreV1Volume `json:"volumes,omitempty"`
}

type _KubevirtIoApiCoreV1VirtualMachineInstanceSpec KubevirtIoApiCoreV1VirtualMachineInstanceSpec

// NewKubevirtIoApiCoreV1VirtualMachineInstanceSpec instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceSpec(domain KubevirtIoApiCoreV1DomainSpec) *KubevirtIoApiCoreV1VirtualMachineInstanceSpec {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceSpec{}
	this.Domain = domain
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceSpecWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceSpecWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceSpec {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceSpec{}
	var domain KubevirtIoApiCoreV1DomainSpec = {}
	this.Domain = domain
	return &this
}

// GetAccessCredentials returns the AccessCredentials field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAccessCredentials() []KubevirtIoApiCoreV1AccessCredential {
	if o == nil || IsNil(o.AccessCredentials) {
		var ret []KubevirtIoApiCoreV1AccessCredential
		return ret
	}
	return o.AccessCredentials
}

// GetAccessCredentialsOk returns a tuple with the AccessCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAccessCredentialsOk() ([]KubevirtIoApiCoreV1AccessCredential, bool) {
	if o == nil || IsNil(o.AccessCredentials) {
		return nil, false
	}
	return o.AccessCredentials, true
}

// HasAccessCredentials returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasAccessCredentials() bool {
	if o != nil && !IsNil(o.AccessCredentials) {
		return true
	}

	return false
}

// SetAccessCredentials gets a reference to the given []KubevirtIoApiCoreV1AccessCredential and assigns it to the AccessCredentials field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetAccessCredentials(v []KubevirtIoApiCoreV1AccessCredential) {
	o.AccessCredentials = v
}

// GetAffinity returns the Affinity field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAffinity() K8sIoV1Affinity {
	if o == nil || IsNil(o.Affinity) {
		var ret K8sIoV1Affinity
		return ret
	}
	return *o.Affinity
}

// GetAffinityOk returns a tuple with the Affinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetAffinityOk() (*K8sIoV1Affinity, bool) {
	if o == nil || IsNil(o.Affinity) {
		return nil, false
	}
	return o.Affinity, true
}

// HasAffinity returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasAffinity() bool {
	if o != nil && !IsNil(o.Affinity) {
		return true
	}

	return false
}

// SetAffinity gets a reference to the given K8sIoV1Affinity and assigns it to the Affinity field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetAffinity(v K8sIoV1Affinity) {
	o.Affinity = &v
}

// GetArchitecture returns the Architecture field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetArchitecture() string {
	if o == nil || IsNil(o.Architecture) {
		var ret string
		return ret
	}
	return *o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetArchitectureOk() (*string, bool) {
	if o == nil || IsNil(o.Architecture) {
		return nil, false
	}
	return o.Architecture, true
}

// HasArchitecture returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasArchitecture() bool {
	if o != nil && !IsNil(o.Architecture) {
		return true
	}

	return false
}

// SetArchitecture gets a reference to the given string and assigns it to the Architecture field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetArchitecture(v string) {
	o.Architecture = &v
}

// GetDnsConfig returns the DnsConfig field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsConfig() K8sIoV1PodDNSConfig {
	if o == nil || IsNil(o.DnsConfig) {
		var ret K8sIoV1PodDNSConfig
		return ret
	}
	return *o.DnsConfig
}

// GetDnsConfigOk returns a tuple with the DnsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsConfigOk() (*K8sIoV1PodDNSConfig, bool) {
	if o == nil || IsNil(o.DnsConfig) {
		return nil, false
	}
	return o.DnsConfig, true
}

// HasDnsConfig returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasDnsConfig() bool {
	if o != nil && !IsNil(o.DnsConfig) {
		return true
	}

	return false
}

// SetDnsConfig gets a reference to the given K8sIoV1PodDNSConfig and assigns it to the DnsConfig field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDnsConfig(v K8sIoV1PodDNSConfig) {
	o.DnsConfig = &v
}

// GetDnsPolicy returns the DnsPolicy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsPolicy() string {
	if o == nil || IsNil(o.DnsPolicy) {
		var ret string
		return ret
	}
	return *o.DnsPolicy
}

// GetDnsPolicyOk returns a tuple with the DnsPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDnsPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.DnsPolicy) {
		return nil, false
	}
	return o.DnsPolicy, true
}

// HasDnsPolicy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasDnsPolicy() bool {
	if o != nil && !IsNil(o.DnsPolicy) {
		return true
	}

	return false
}

// SetDnsPolicy gets a reference to the given string and assigns it to the DnsPolicy field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDnsPolicy(v string) {
	o.DnsPolicy = &v
}

// GetDomain returns the Domain field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDomain() KubevirtIoApiCoreV1DomainSpec {
	if o == nil {
		var ret KubevirtIoApiCoreV1DomainSpec
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetDomainOk() (*KubevirtIoApiCoreV1DomainSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetDomain(v KubevirtIoApiCoreV1DomainSpec) {
	o.Domain = v
}

// GetEvictionStrategy returns the EvictionStrategy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetEvictionStrategy() string {
	if o == nil || IsNil(o.EvictionStrategy) {
		var ret string
		return ret
	}
	return *o.EvictionStrategy
}

// GetEvictionStrategyOk returns a tuple with the EvictionStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetEvictionStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.EvictionStrategy) {
		return nil, false
	}
	return o.EvictionStrategy, true
}

// HasEvictionStrategy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasEvictionStrategy() bool {
	if o != nil && !IsNil(o.EvictionStrategy) {
		return true
	}

	return false
}

// SetEvictionStrategy gets a reference to the given string and assigns it to the EvictionStrategy field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetEvictionStrategy(v string) {
	o.EvictionStrategy = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetHostname(v string) {
	o.Hostname = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetLivenessProbe() KubevirtIoApiCoreV1Probe {
	if o == nil || IsNil(o.LivenessProbe) {
		var ret KubevirtIoApiCoreV1Probe
		return ret
	}
	return *o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetLivenessProbeOk() (*KubevirtIoApiCoreV1Probe, bool) {
	if o == nil || IsNil(o.LivenessProbe) {
		return nil, false
	}
	return o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasLivenessProbe() bool {
	if o != nil && !IsNil(o.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given KubevirtIoApiCoreV1Probe and assigns it to the LivenessProbe field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetLivenessProbe(v KubevirtIoApiCoreV1Probe) {
	o.LivenessProbe = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNetworks() []KubevirtIoApiCoreV1Network {
	if o == nil || IsNil(o.Networks) {
		var ret []KubevirtIoApiCoreV1Network
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNetworksOk() ([]KubevirtIoApiCoreV1Network, bool) {
	if o == nil || IsNil(o.Networks) {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasNetworks() bool {
	if o != nil && !IsNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []KubevirtIoApiCoreV1Network and assigns it to the Networks field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetNetworks(v []KubevirtIoApiCoreV1Network) {
	o.Networks = v
}

// GetNodeSelector returns the NodeSelector field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNodeSelector() map[string]string {
	if o == nil || IsNil(o.NodeSelector) {
		var ret map[string]string
		return ret
	}
	return *o.NodeSelector
}

// GetNodeSelectorOk returns a tuple with the NodeSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetNodeSelectorOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.NodeSelector) {
		return nil, false
	}
	return o.NodeSelector, true
}

// HasNodeSelector returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasNodeSelector() bool {
	if o != nil && !IsNil(o.NodeSelector) {
		return true
	}

	return false
}

// SetNodeSelector gets a reference to the given map[string]string and assigns it to the NodeSelector field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetNodeSelector(v map[string]string) {
	o.NodeSelector = &v
}

// GetPriorityClassName returns the PriorityClassName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetPriorityClassName() string {
	if o == nil || IsNil(o.PriorityClassName) {
		var ret string
		return ret
	}
	return *o.PriorityClassName
}

// GetPriorityClassNameOk returns a tuple with the PriorityClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetPriorityClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.PriorityClassName) {
		return nil, false
	}
	return o.PriorityClassName, true
}

// HasPriorityClassName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasPriorityClassName() bool {
	if o != nil && !IsNil(o.PriorityClassName) {
		return true
	}

	return false
}

// SetPriorityClassName gets a reference to the given string and assigns it to the PriorityClassName field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetPriorityClassName(v string) {
	o.PriorityClassName = &v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetReadinessProbe() KubevirtIoApiCoreV1Probe {
	if o == nil || IsNil(o.ReadinessProbe) {
		var ret KubevirtIoApiCoreV1Probe
		return ret
	}
	return *o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetReadinessProbeOk() (*KubevirtIoApiCoreV1Probe, bool) {
	if o == nil || IsNil(o.ReadinessProbe) {
		return nil, false
	}
	return o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasReadinessProbe() bool {
	if o != nil && !IsNil(o.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given KubevirtIoApiCoreV1Probe and assigns it to the ReadinessProbe field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetReadinessProbe(v KubevirtIoApiCoreV1Probe) {
	o.ReadinessProbe = &v
}

// GetSchedulerName returns the SchedulerName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSchedulerName() string {
	if o == nil || IsNil(o.SchedulerName) {
		var ret string
		return ret
	}
	return *o.SchedulerName
}

// GetSchedulerNameOk returns a tuple with the SchedulerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSchedulerNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchedulerName) {
		return nil, false
	}
	return o.SchedulerName, true
}

// HasSchedulerName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasSchedulerName() bool {
	if o != nil && !IsNil(o.SchedulerName) {
		return true
	}

	return false
}

// SetSchedulerName gets a reference to the given string and assigns it to the SchedulerName field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetSchedulerName(v string) {
	o.SchedulerName = &v
}

// GetStartStrategy returns the StartStrategy field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetStartStrategy() string {
	if o == nil || IsNil(o.StartStrategy) {
		var ret string
		return ret
	}
	return *o.StartStrategy
}

// GetStartStrategyOk returns a tuple with the StartStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetStartStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.StartStrategy) {
		return nil, false
	}
	return o.StartStrategy, true
}

// HasStartStrategy returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasStartStrategy() bool {
	if o != nil && !IsNil(o.StartStrategy) {
		return true
	}

	return false
}

// SetStartStrategy gets a reference to the given string and assigns it to the StartStrategy field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetStartStrategy(v string) {
	o.StartStrategy = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetTerminationGracePeriodSeconds returns the TerminationGracePeriodSeconds field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTerminationGracePeriodSeconds() int64 {
	if o == nil || IsNil(o.TerminationGracePeriodSeconds) {
		var ret int64
		return ret
	}
	return *o.TerminationGracePeriodSeconds
}

// GetTerminationGracePeriodSecondsOk returns a tuple with the TerminationGracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTerminationGracePeriodSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.TerminationGracePeriodSeconds) {
		return nil, false
	}
	return o.TerminationGracePeriodSeconds, true
}

// HasTerminationGracePeriodSeconds returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTerminationGracePeriodSeconds() bool {
	if o != nil && !IsNil(o.TerminationGracePeriodSeconds) {
		return true
	}

	return false
}

// SetTerminationGracePeriodSeconds gets a reference to the given int64 and assigns it to the TerminationGracePeriodSeconds field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTerminationGracePeriodSeconds(v int64) {
	o.TerminationGracePeriodSeconds = &v
}

// GetTolerations returns the Tolerations field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTolerations() []K8sIoV1Toleration {
	if o == nil || IsNil(o.Tolerations) {
		var ret []K8sIoV1Toleration
		return ret
	}
	return o.Tolerations
}

// GetTolerationsOk returns a tuple with the Tolerations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTolerationsOk() ([]K8sIoV1Toleration, bool) {
	if o == nil || IsNil(o.Tolerations) {
		return nil, false
	}
	return o.Tolerations, true
}

// HasTolerations returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTolerations() bool {
	if o != nil && !IsNil(o.Tolerations) {
		return true
	}

	return false
}

// SetTolerations gets a reference to the given []K8sIoV1Toleration and assigns it to the Tolerations field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTolerations(v []K8sIoV1Toleration) {
	o.Tolerations = v
}

// GetTopologySpreadConstraints returns the TopologySpreadConstraints field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTopologySpreadConstraints() []K8sIoV1TopologySpreadConstraint {
	if o == nil || IsNil(o.TopologySpreadConstraints) {
		var ret []K8sIoV1TopologySpreadConstraint
		return ret
	}
	return o.TopologySpreadConstraints
}

// GetTopologySpreadConstraintsOk returns a tuple with the TopologySpreadConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetTopologySpreadConstraintsOk() ([]K8sIoV1TopologySpreadConstraint, bool) {
	if o == nil || IsNil(o.TopologySpreadConstraints) {
		return nil, false
	}
	return o.TopologySpreadConstraints, true
}

// HasTopologySpreadConstraints returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasTopologySpreadConstraints() bool {
	if o != nil && !IsNil(o.TopologySpreadConstraints) {
		return true
	}

	return false
}

// SetTopologySpreadConstraints gets a reference to the given []K8sIoV1TopologySpreadConstraint and assigns it to the TopologySpreadConstraints field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetTopologySpreadConstraints(v []K8sIoV1TopologySpreadConstraint) {
	o.TopologySpreadConstraints = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetVolumes() []KubevirtIoApiCoreV1Volume {
	if o == nil || IsNil(o.Volumes) {
		var ret []KubevirtIoApiCoreV1Volume
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) GetVolumesOk() ([]KubevirtIoApiCoreV1Volume, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []KubevirtIoApiCoreV1Volume and assigns it to the Volumes field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) SetVolumes(v []KubevirtIoApiCoreV1Volume) {
	o.Volumes = v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceSpec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessCredentials) {
		toSerialize["accessCredentials"] = o.AccessCredentials
	}
	if !IsNil(o.Affinity) {
		toSerialize["affinity"] = o.Affinity
	}
	if !IsNil(o.Architecture) {
		toSerialize["architecture"] = o.Architecture
	}
	if !IsNil(o.DnsConfig) {
		toSerialize["dnsConfig"] = o.DnsConfig
	}
	if !IsNil(o.DnsPolicy) {
		toSerialize["dnsPolicy"] = o.DnsPolicy
	}
	toSerialize["domain"] = o.Domain
	if !IsNil(o.EvictionStrategy) {
		toSerialize["evictionStrategy"] = o.EvictionStrategy
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if !IsNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !IsNil(o.NodeSelector) {
		toSerialize["nodeSelector"] = o.NodeSelector
	}
	if !IsNil(o.PriorityClassName) {
		toSerialize["priorityClassName"] = o.PriorityClassName
	}
	if !IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if !IsNil(o.SchedulerName) {
		toSerialize["schedulerName"] = o.SchedulerName
	}
	if !IsNil(o.StartStrategy) {
		toSerialize["startStrategy"] = o.StartStrategy
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.TerminationGracePeriodSeconds) {
		toSerialize["terminationGracePeriodSeconds"] = o.TerminationGracePeriodSeconds
	}
	if !IsNil(o.Tolerations) {
		toSerialize["tolerations"] = o.Tolerations
	}
	if !IsNil(o.TopologySpreadConstraints) {
		toSerialize["topologySpreadConstraints"] = o.TopologySpreadConstraints
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubevirtIoApiCoreV1VirtualMachineInstanceSpec := _KubevirtIoApiCoreV1VirtualMachineInstanceSpec{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1VirtualMachineInstanceSpec)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1VirtualMachineInstanceSpec(varKubevirtIoApiCoreV1VirtualMachineInstanceSpec)

	return err
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceSpec
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceSpec {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec(val *KubevirtIoApiCoreV1VirtualMachineInstanceSpec) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


