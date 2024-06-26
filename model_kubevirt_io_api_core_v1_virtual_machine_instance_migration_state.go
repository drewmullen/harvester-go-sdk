/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState{}

// KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState struct for KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState
type KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState struct {
	// Indicates that the migration has been requested to abort
	AbortRequested *bool `json:"abortRequested,omitempty"`
	// Indicates the final status of the live migration abortion
	AbortStatus *string `json:"abortStatus,omitempty"`
	// Indicates the migration completed
	Completed *bool `json:"completed,omitempty"`
	// The time the migration action ended
	EndTimestamp *string `json:"endTimestamp,omitempty"`
	// Indicates that the migration failed
	Failed *bool `json:"failed,omitempty"`
	// Migration configurations to apply
	MigrationConfiguration *KubevirtIoApiCoreV1MigrationConfiguration `json:"migrationConfiguration,omitempty"`
	// Name of the migration policy. If string is empty, no policy is matched
	MigrationPolicyName *string `json:"migrationPolicyName,omitempty"`
	// The VirtualMachineInstanceMigration object associated with this migration
	MigrationUid *string `json:"migrationUid,omitempty"`
	// Lets us know if the vmi is currently running pre or post copy migration
	Mode *string `json:"mode,omitempty"`
	// The source node that the VMI originated on
	SourceNode *string `json:"sourceNode,omitempty"`
	// The time the migration action began
	StartTimestamp *string `json:"startTimestamp,omitempty"`
	// The UID of the target attachment pod for hotplug volumes
	TargetAttachmentPodUID *string `json:"targetAttachmentPodUID,omitempty"`
	// If the VMI requires dedicated CPUs, this field will hold the dedicated CPU set on the target node
	TargetCPUSet []int32 `json:"targetCPUSet,omitempty"`
	// The list of ports opened for live migration on the destination node
	TargetDirectMigrationNodePorts *map[string]int32 `json:"targetDirectMigrationNodePorts,omitempty"`
	// The target node that the VMI is moving to
	TargetNode *string `json:"targetNode,omitempty"`
	// The address of the target node to use for the migration
	TargetNodeAddress *string `json:"targetNodeAddress,omitempty"`
	// The Target Node has seen the Domain Start Event
	TargetNodeDomainDetected *bool `json:"targetNodeDomainDetected,omitempty"`
	// The timestamp at which the target node detects the domain is active
	TargetNodeDomainReadyTimestamp *string `json:"targetNodeDomainReadyTimestamp,omitempty"`
	// If the VMI requires dedicated CPUs, this field will hold the numa topology on the target node
	TargetNodeTopology *string `json:"targetNodeTopology,omitempty"`
	// The target pod that the VMI is moving to
	TargetPod *string `json:"targetPod,omitempty"`
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState{}
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStateWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStateWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState{}
	return &this
}

// GetAbortRequested returns the AbortRequested field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortRequested() bool {
	if o == nil || IsNil(o.AbortRequested) {
		var ret bool
		return ret
	}
	return *o.AbortRequested
}

// GetAbortRequestedOk returns a tuple with the AbortRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.AbortRequested) {
		return nil, false
	}
	return o.AbortRequested, true
}

// HasAbortRequested returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasAbortRequested() bool {
	if o != nil && !IsNil(o.AbortRequested) {
		return true
	}

	return false
}

// SetAbortRequested gets a reference to the given bool and assigns it to the AbortRequested field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetAbortRequested(v bool) {
	o.AbortRequested = &v
}

// GetAbortStatus returns the AbortStatus field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortStatus() string {
	if o == nil || IsNil(o.AbortStatus) {
		var ret string
		return ret
	}
	return *o.AbortStatus
}

// GetAbortStatusOk returns a tuple with the AbortStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortStatusOk() (*string, bool) {
	if o == nil || IsNil(o.AbortStatus) {
		return nil, false
	}
	return o.AbortStatus, true
}

// HasAbortStatus returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasAbortStatus() bool {
	if o != nil && !IsNil(o.AbortStatus) {
		return true
	}

	return false
}

// SetAbortStatus gets a reference to the given string and assigns it to the AbortStatus field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetAbortStatus(v string) {
	o.AbortStatus = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetCompleted() bool {
	if o == nil || IsNil(o.Completed) {
		var ret bool
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasCompleted() bool {
	if o != nil && !IsNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given bool and assigns it to the Completed field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetCompleted(v bool) {
	o.Completed = &v
}

// GetEndTimestamp returns the EndTimestamp field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetEndTimestamp() string {
	if o == nil || IsNil(o.EndTimestamp) {
		var ret string
		return ret
	}
	return *o.EndTimestamp
}

// GetEndTimestampOk returns a tuple with the EndTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetEndTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.EndTimestamp) {
		return nil, false
	}
	return o.EndTimestamp, true
}

// HasEndTimestamp returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasEndTimestamp() bool {
	if o != nil && !IsNil(o.EndTimestamp) {
		return true
	}

	return false
}

// SetEndTimestamp gets a reference to the given string and assigns it to the EndTimestamp field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetEndTimestamp(v string) {
	o.EndTimestamp = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetFailed(v bool) {
	o.Failed = &v
}

// GetMigrationConfiguration returns the MigrationConfiguration field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationConfiguration() KubevirtIoApiCoreV1MigrationConfiguration {
	if o == nil || IsNil(o.MigrationConfiguration) {
		var ret KubevirtIoApiCoreV1MigrationConfiguration
		return ret
	}
	return *o.MigrationConfiguration
}

// GetMigrationConfigurationOk returns a tuple with the MigrationConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationConfigurationOk() (*KubevirtIoApiCoreV1MigrationConfiguration, bool) {
	if o == nil || IsNil(o.MigrationConfiguration) {
		return nil, false
	}
	return o.MigrationConfiguration, true
}

// HasMigrationConfiguration returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationConfiguration() bool {
	if o != nil && !IsNil(o.MigrationConfiguration) {
		return true
	}

	return false
}

// SetMigrationConfiguration gets a reference to the given KubevirtIoApiCoreV1MigrationConfiguration and assigns it to the MigrationConfiguration field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationConfiguration(v KubevirtIoApiCoreV1MigrationConfiguration) {
	o.MigrationConfiguration = &v
}

// GetMigrationPolicyName returns the MigrationPolicyName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationPolicyName() string {
	if o == nil || IsNil(o.MigrationPolicyName) {
		var ret string
		return ret
	}
	return *o.MigrationPolicyName
}

// GetMigrationPolicyNameOk returns a tuple with the MigrationPolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.MigrationPolicyName) {
		return nil, false
	}
	return o.MigrationPolicyName, true
}

// HasMigrationPolicyName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationPolicyName() bool {
	if o != nil && !IsNil(o.MigrationPolicyName) {
		return true
	}

	return false
}

// SetMigrationPolicyName gets a reference to the given string and assigns it to the MigrationPolicyName field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationPolicyName(v string) {
	o.MigrationPolicyName = &v
}

// GetMigrationUid returns the MigrationUid field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationUid() string {
	if o == nil || IsNil(o.MigrationUid) {
		var ret string
		return ret
	}
	return *o.MigrationUid
}

// GetMigrationUidOk returns a tuple with the MigrationUid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationUidOk() (*string, bool) {
	if o == nil || IsNil(o.MigrationUid) {
		return nil, false
	}
	return o.MigrationUid, true
}

// HasMigrationUid returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationUid() bool {
	if o != nil && !IsNil(o.MigrationUid) {
		return true
	}

	return false
}

// SetMigrationUid gets a reference to the given string and assigns it to the MigrationUid field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationUid(v string) {
	o.MigrationUid = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMode(v string) {
	o.Mode = &v
}

// GetSourceNode returns the SourceNode field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetSourceNode() string {
	if o == nil || IsNil(o.SourceNode) {
		var ret string
		return ret
	}
	return *o.SourceNode
}

// GetSourceNodeOk returns a tuple with the SourceNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetSourceNodeOk() (*string, bool) {
	if o == nil || IsNil(o.SourceNode) {
		return nil, false
	}
	return o.SourceNode, true
}

// HasSourceNode returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasSourceNode() bool {
	if o != nil && !IsNil(o.SourceNode) {
		return true
	}

	return false
}

// SetSourceNode gets a reference to the given string and assigns it to the SourceNode field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetSourceNode(v string) {
	o.SourceNode = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetStartTimestamp() string {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret string
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetStartTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given string and assigns it to the StartTimestamp field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetStartTimestamp(v string) {
	o.StartTimestamp = &v
}

// GetTargetAttachmentPodUID returns the TargetAttachmentPodUID field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetAttachmentPodUID() string {
	if o == nil || IsNil(o.TargetAttachmentPodUID) {
		var ret string
		return ret
	}
	return *o.TargetAttachmentPodUID
}

// GetTargetAttachmentPodUIDOk returns a tuple with the TargetAttachmentPodUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetAttachmentPodUIDOk() (*string, bool) {
	if o == nil || IsNil(o.TargetAttachmentPodUID) {
		return nil, false
	}
	return o.TargetAttachmentPodUID, true
}

// HasTargetAttachmentPodUID returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetAttachmentPodUID() bool {
	if o != nil && !IsNil(o.TargetAttachmentPodUID) {
		return true
	}

	return false
}

// SetTargetAttachmentPodUID gets a reference to the given string and assigns it to the TargetAttachmentPodUID field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetAttachmentPodUID(v string) {
	o.TargetAttachmentPodUID = &v
}

// GetTargetCPUSet returns the TargetCPUSet field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetCPUSet() []int32 {
	if o == nil || IsNil(o.TargetCPUSet) {
		var ret []int32
		return ret
	}
	return o.TargetCPUSet
}

// GetTargetCPUSetOk returns a tuple with the TargetCPUSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetCPUSetOk() ([]int32, bool) {
	if o == nil || IsNil(o.TargetCPUSet) {
		return nil, false
	}
	return o.TargetCPUSet, true
}

// HasTargetCPUSet returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetCPUSet() bool {
	if o != nil && !IsNil(o.TargetCPUSet) {
		return true
	}

	return false
}

// SetTargetCPUSet gets a reference to the given []int32 and assigns it to the TargetCPUSet field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetCPUSet(v []int32) {
	o.TargetCPUSet = v
}

// GetTargetDirectMigrationNodePorts returns the TargetDirectMigrationNodePorts field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetDirectMigrationNodePorts() map[string]int32 {
	if o == nil || IsNil(o.TargetDirectMigrationNodePorts) {
		var ret map[string]int32
		return ret
	}
	return *o.TargetDirectMigrationNodePorts
}

// GetTargetDirectMigrationNodePortsOk returns a tuple with the TargetDirectMigrationNodePorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetDirectMigrationNodePortsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.TargetDirectMigrationNodePorts) {
		return nil, false
	}
	return o.TargetDirectMigrationNodePorts, true
}

// HasTargetDirectMigrationNodePorts returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetDirectMigrationNodePorts() bool {
	if o != nil && !IsNil(o.TargetDirectMigrationNodePorts) {
		return true
	}

	return false
}

// SetTargetDirectMigrationNodePorts gets a reference to the given map[string]int32 and assigns it to the TargetDirectMigrationNodePorts field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetDirectMigrationNodePorts(v map[string]int32) {
	o.TargetDirectMigrationNodePorts = &v
}

// GetTargetNode returns the TargetNode field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNode() string {
	if o == nil || IsNil(o.TargetNode) {
		var ret string
		return ret
	}
	return *o.TargetNode
}

// GetTargetNodeOk returns a tuple with the TargetNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNode) {
		return nil, false
	}
	return o.TargetNode, true
}

// HasTargetNode returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNode() bool {
	if o != nil && !IsNil(o.TargetNode) {
		return true
	}

	return false
}

// SetTargetNode gets a reference to the given string and assigns it to the TargetNode field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNode(v string) {
	o.TargetNode = &v
}

// GetTargetNodeAddress returns the TargetNodeAddress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeAddress() string {
	if o == nil || IsNil(o.TargetNodeAddress) {
		var ret string
		return ret
	}
	return *o.TargetNodeAddress
}

// GetTargetNodeAddressOk returns a tuple with the TargetNodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeAddressOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNodeAddress) {
		return nil, false
	}
	return o.TargetNodeAddress, true
}

// HasTargetNodeAddress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeAddress() bool {
	if o != nil && !IsNil(o.TargetNodeAddress) {
		return true
	}

	return false
}

// SetTargetNodeAddress gets a reference to the given string and assigns it to the TargetNodeAddress field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeAddress(v string) {
	o.TargetNodeAddress = &v
}

// GetTargetNodeDomainDetected returns the TargetNodeDomainDetected field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainDetected() bool {
	if o == nil || IsNil(o.TargetNodeDomainDetected) {
		var ret bool
		return ret
	}
	return *o.TargetNodeDomainDetected
}

// GetTargetNodeDomainDetectedOk returns a tuple with the TargetNodeDomainDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainDetectedOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetNodeDomainDetected) {
		return nil, false
	}
	return o.TargetNodeDomainDetected, true
}

// HasTargetNodeDomainDetected returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeDomainDetected() bool {
	if o != nil && !IsNil(o.TargetNodeDomainDetected) {
		return true
	}

	return false
}

// SetTargetNodeDomainDetected gets a reference to the given bool and assigns it to the TargetNodeDomainDetected field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeDomainDetected(v bool) {
	o.TargetNodeDomainDetected = &v
}

// GetTargetNodeDomainReadyTimestamp returns the TargetNodeDomainReadyTimestamp field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainReadyTimestamp() string {
	if o == nil || IsNil(o.TargetNodeDomainReadyTimestamp) {
		var ret string
		return ret
	}
	return *o.TargetNodeDomainReadyTimestamp
}

// GetTargetNodeDomainReadyTimestampOk returns a tuple with the TargetNodeDomainReadyTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainReadyTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNodeDomainReadyTimestamp) {
		return nil, false
	}
	return o.TargetNodeDomainReadyTimestamp, true
}

// HasTargetNodeDomainReadyTimestamp returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeDomainReadyTimestamp() bool {
	if o != nil && !IsNil(o.TargetNodeDomainReadyTimestamp) {
		return true
	}

	return false
}

// SetTargetNodeDomainReadyTimestamp gets a reference to the given string and assigns it to the TargetNodeDomainReadyTimestamp field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeDomainReadyTimestamp(v string) {
	o.TargetNodeDomainReadyTimestamp = &v
}

// GetTargetNodeTopology returns the TargetNodeTopology field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeTopology() string {
	if o == nil || IsNil(o.TargetNodeTopology) {
		var ret string
		return ret
	}
	return *o.TargetNodeTopology
}

// GetTargetNodeTopologyOk returns a tuple with the TargetNodeTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeTopologyOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNodeTopology) {
		return nil, false
	}
	return o.TargetNodeTopology, true
}

// HasTargetNodeTopology returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeTopology() bool {
	if o != nil && !IsNil(o.TargetNodeTopology) {
		return true
	}

	return false
}

// SetTargetNodeTopology gets a reference to the given string and assigns it to the TargetNodeTopology field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeTopology(v string) {
	o.TargetNodeTopology = &v
}

// GetTargetPod returns the TargetPod field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetPod() string {
	if o == nil || IsNil(o.TargetPod) {
		var ret string
		return ret
	}
	return *o.TargetPod
}

// GetTargetPodOk returns a tuple with the TargetPod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetPodOk() (*string, bool) {
	if o == nil || IsNil(o.TargetPod) {
		return nil, false
	}
	return o.TargetPod, true
}

// HasTargetPod returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetPod() bool {
	if o != nil && !IsNil(o.TargetPod) {
		return true
	}

	return false
}

// SetTargetPod gets a reference to the given string and assigns it to the TargetPod field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetPod(v string) {
	o.TargetPod = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AbortRequested) {
		toSerialize["abortRequested"] = o.AbortRequested
	}
	if !IsNil(o.AbortStatus) {
		toSerialize["abortStatus"] = o.AbortStatus
	}
	if !IsNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if !IsNil(o.EndTimestamp) {
		toSerialize["endTimestamp"] = o.EndTimestamp
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.MigrationConfiguration) {
		toSerialize["migrationConfiguration"] = o.MigrationConfiguration
	}
	if !IsNil(o.MigrationPolicyName) {
		toSerialize["migrationPolicyName"] = o.MigrationPolicyName
	}
	if !IsNil(o.MigrationUid) {
		toSerialize["migrationUid"] = o.MigrationUid
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.SourceNode) {
		toSerialize["sourceNode"] = o.SourceNode
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	if !IsNil(o.TargetAttachmentPodUID) {
		toSerialize["targetAttachmentPodUID"] = o.TargetAttachmentPodUID
	}
	if !IsNil(o.TargetCPUSet) {
		toSerialize["targetCPUSet"] = o.TargetCPUSet
	}
	if !IsNil(o.TargetDirectMigrationNodePorts) {
		toSerialize["targetDirectMigrationNodePorts"] = o.TargetDirectMigrationNodePorts
	}
	if !IsNil(o.TargetNode) {
		toSerialize["targetNode"] = o.TargetNode
	}
	if !IsNil(o.TargetNodeAddress) {
		toSerialize["targetNodeAddress"] = o.TargetNodeAddress
	}
	if !IsNil(o.TargetNodeDomainDetected) {
		toSerialize["targetNodeDomainDetected"] = o.TargetNodeDomainDetected
	}
	if !IsNil(o.TargetNodeDomainReadyTimestamp) {
		toSerialize["targetNodeDomainReadyTimestamp"] = o.TargetNodeDomainReadyTimestamp
	}
	if !IsNil(o.TargetNodeTopology) {
		toSerialize["targetNodeTopology"] = o.TargetNodeTopology
	}
	if !IsNil(o.TargetPod) {
		toSerialize["targetPod"] = o.TargetPod
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState(val *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


