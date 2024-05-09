# KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortRequested** | Pointer to **bool** | Indicates that the migration has been requested to abort | [optional] 
**AbortStatus** | Pointer to **string** | Indicates the final status of the live migration abortion | [optional] 
**Completed** | Pointer to **bool** | Indicates the migration completed | [optional] 
**EndTimestamp** | Pointer to **string** | The time the migration action ended | [optional] 
**Failed** | Pointer to **bool** | Indicates that the migration failed | [optional] 
**MigrationConfiguration** | Pointer to [**KubevirtIoApiCoreV1MigrationConfiguration**](KubevirtIoApiCoreV1MigrationConfiguration.md) | Migration configurations to apply | [optional] 
**MigrationPolicyName** | Pointer to **string** | Name of the migration policy. If string is empty, no policy is matched | [optional] 
**MigrationUid** | Pointer to **string** | The VirtualMachineInstanceMigration object associated with this migration | [optional] 
**Mode** | Pointer to **string** | Lets us know if the vmi is currently running pre or post copy migration | [optional] 
**SourceNode** | Pointer to **string** | The source node that the VMI originated on | [optional] 
**StartTimestamp** | Pointer to **string** | The time the migration action began | [optional] 
**TargetAttachmentPodUID** | Pointer to **string** | The UID of the target attachment pod for hotplug volumes | [optional] 
**TargetCPUSet** | Pointer to **[]int32** | If the VMI requires dedicated CPUs, this field will hold the dedicated CPU set on the target node | [optional] 
**TargetDirectMigrationNodePorts** | Pointer to **map[string]int32** | The list of ports opened for live migration on the destination node | [optional] 
**TargetNode** | Pointer to **string** | The target node that the VMI is moving to | [optional] 
**TargetNodeAddress** | Pointer to **string** | The address of the target node to use for the migration | [optional] 
**TargetNodeDomainDetected** | Pointer to **bool** | The Target Node has seen the Domain Start Event | [optional] 
**TargetNodeDomainReadyTimestamp** | Pointer to **string** | The timestamp at which the target node detects the domain is active | [optional] 
**TargetNodeTopology** | Pointer to **string** | If the VMI requires dedicated CPUs, this field will hold the numa topology on the target node | [optional] 
**TargetPod** | Pointer to **string** | The target pod that the VMI is moving to | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationState instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStateWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStateWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStateWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortRequested

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortRequested() bool`

GetAbortRequested returns the AbortRequested field if non-nil, zero value otherwise.

### GetAbortRequestedOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortRequestedOk() (*bool, bool)`

GetAbortRequestedOk returns a tuple with the AbortRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortRequested

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetAbortRequested(v bool)`

SetAbortRequested sets AbortRequested field to given value.

### HasAbortRequested

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasAbortRequested() bool`

HasAbortRequested returns a boolean if a field has been set.

### GetAbortStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortStatus() string`

GetAbortStatus returns the AbortStatus field if non-nil, zero value otherwise.

### GetAbortStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetAbortStatusOk() (*string, bool)`

GetAbortStatusOk returns a tuple with the AbortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetAbortStatus(v string)`

SetAbortStatus sets AbortStatus field to given value.

### HasAbortStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasAbortStatus() bool`

HasAbortStatus returns a boolean if a field has been set.

### GetCompleted

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetCompleted() bool`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetCompletedOk() (*bool, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetCompleted(v bool)`

SetCompleted sets Completed field to given value.

### HasCompleted

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasCompleted() bool`

HasCompleted returns a boolean if a field has been set.

### GetEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetFailed

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetMigrationConfiguration

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationConfiguration() KubevirtIoApiCoreV1MigrationConfiguration`

GetMigrationConfiguration returns the MigrationConfiguration field if non-nil, zero value otherwise.

### GetMigrationConfigurationOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationConfigurationOk() (*KubevirtIoApiCoreV1MigrationConfiguration, bool)`

GetMigrationConfigurationOk returns a tuple with the MigrationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationConfiguration

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationConfiguration(v KubevirtIoApiCoreV1MigrationConfiguration)`

SetMigrationConfiguration sets MigrationConfiguration field to given value.

### HasMigrationConfiguration

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationConfiguration() bool`

HasMigrationConfiguration returns a boolean if a field has been set.

### GetMigrationPolicyName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationPolicyName() string`

GetMigrationPolicyName returns the MigrationPolicyName field if non-nil, zero value otherwise.

### GetMigrationPolicyNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationPolicyNameOk() (*string, bool)`

GetMigrationPolicyNameOk returns a tuple with the MigrationPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationPolicyName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationPolicyName(v string)`

SetMigrationPolicyName sets MigrationPolicyName field to given value.

### HasMigrationPolicyName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationPolicyName() bool`

HasMigrationPolicyName returns a boolean if a field has been set.

### GetMigrationUid

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationUid() string`

GetMigrationUid returns the MigrationUid field if non-nil, zero value otherwise.

### GetMigrationUidOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMigrationUidOk() (*string, bool)`

GetMigrationUidOk returns a tuple with the MigrationUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationUid

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMigrationUid(v string)`

SetMigrationUid sets MigrationUid field to given value.

### HasMigrationUid

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMigrationUid() bool`

HasMigrationUid returns a boolean if a field has been set.

### GetMode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSourceNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetSourceNode() string`

GetSourceNode returns the SourceNode field if non-nil, zero value otherwise.

### GetSourceNodeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetSourceNodeOk() (*string, bool)`

GetSourceNodeOk returns a tuple with the SourceNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetSourceNode(v string)`

SetSourceNode sets SourceNode field to given value.

### HasSourceNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasSourceNode() bool`

HasSourceNode returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.

### GetTargetAttachmentPodUID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetAttachmentPodUID() string`

GetTargetAttachmentPodUID returns the TargetAttachmentPodUID field if non-nil, zero value otherwise.

### GetTargetAttachmentPodUIDOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetAttachmentPodUIDOk() (*string, bool)`

GetTargetAttachmentPodUIDOk returns a tuple with the TargetAttachmentPodUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttachmentPodUID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetAttachmentPodUID(v string)`

SetTargetAttachmentPodUID sets TargetAttachmentPodUID field to given value.

### HasTargetAttachmentPodUID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetAttachmentPodUID() bool`

HasTargetAttachmentPodUID returns a boolean if a field has been set.

### GetTargetCPUSet

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetCPUSet() []int32`

GetTargetCPUSet returns the TargetCPUSet field if non-nil, zero value otherwise.

### GetTargetCPUSetOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetCPUSetOk() (*[]int32, bool)`

GetTargetCPUSetOk returns a tuple with the TargetCPUSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCPUSet

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetCPUSet(v []int32)`

SetTargetCPUSet sets TargetCPUSet field to given value.

### HasTargetCPUSet

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetCPUSet() bool`

HasTargetCPUSet returns a boolean if a field has been set.

### GetTargetDirectMigrationNodePorts

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetDirectMigrationNodePorts() map[string]int32`

GetTargetDirectMigrationNodePorts returns the TargetDirectMigrationNodePorts field if non-nil, zero value otherwise.

### GetTargetDirectMigrationNodePortsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetDirectMigrationNodePortsOk() (*map[string]int32, bool)`

GetTargetDirectMigrationNodePortsOk returns a tuple with the TargetDirectMigrationNodePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectMigrationNodePorts

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetDirectMigrationNodePorts(v map[string]int32)`

SetTargetDirectMigrationNodePorts sets TargetDirectMigrationNodePorts field to given value.

### HasTargetDirectMigrationNodePorts

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetDirectMigrationNodePorts() bool`

HasTargetDirectMigrationNodePorts returns a boolean if a field has been set.

### GetTargetNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNode() string`

GetTargetNode returns the TargetNode field if non-nil, zero value otherwise.

### GetTargetNodeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeOk() (*string, bool)`

GetTargetNodeOk returns a tuple with the TargetNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNode(v string)`

SetTargetNode sets TargetNode field to given value.

### HasTargetNode

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNode() bool`

HasTargetNode returns a boolean if a field has been set.

### GetTargetNodeAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeAddress() string`

GetTargetNodeAddress returns the TargetNodeAddress field if non-nil, zero value otherwise.

### GetTargetNodeAddressOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeAddressOk() (*string, bool)`

GetTargetNodeAddressOk returns a tuple with the TargetNodeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodeAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeAddress(v string)`

SetTargetNodeAddress sets TargetNodeAddress field to given value.

### HasTargetNodeAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeAddress() bool`

HasTargetNodeAddress returns a boolean if a field has been set.

### GetTargetNodeDomainDetected

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainDetected() bool`

GetTargetNodeDomainDetected returns the TargetNodeDomainDetected field if non-nil, zero value otherwise.

### GetTargetNodeDomainDetectedOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainDetectedOk() (*bool, bool)`

GetTargetNodeDomainDetectedOk returns a tuple with the TargetNodeDomainDetected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodeDomainDetected

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeDomainDetected(v bool)`

SetTargetNodeDomainDetected sets TargetNodeDomainDetected field to given value.

### HasTargetNodeDomainDetected

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeDomainDetected() bool`

HasTargetNodeDomainDetected returns a boolean if a field has been set.

### GetTargetNodeDomainReadyTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainReadyTimestamp() string`

GetTargetNodeDomainReadyTimestamp returns the TargetNodeDomainReadyTimestamp field if non-nil, zero value otherwise.

### GetTargetNodeDomainReadyTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeDomainReadyTimestampOk() (*string, bool)`

GetTargetNodeDomainReadyTimestampOk returns a tuple with the TargetNodeDomainReadyTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodeDomainReadyTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeDomainReadyTimestamp(v string)`

SetTargetNodeDomainReadyTimestamp sets TargetNodeDomainReadyTimestamp field to given value.

### HasTargetNodeDomainReadyTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeDomainReadyTimestamp() bool`

HasTargetNodeDomainReadyTimestamp returns a boolean if a field has been set.

### GetTargetNodeTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeTopology() string`

GetTargetNodeTopology returns the TargetNodeTopology field if non-nil, zero value otherwise.

### GetTargetNodeTopologyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetNodeTopologyOk() (*string, bool)`

GetTargetNodeTopologyOk returns a tuple with the TargetNodeTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNodeTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetNodeTopology(v string)`

SetTargetNodeTopology sets TargetNodeTopology field to given value.

### HasTargetNodeTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetNodeTopology() bool`

HasTargetNodeTopology returns a boolean if a field has been set.

### GetTargetPod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetPod() string`

GetTargetPod returns the TargetPod field if non-nil, zero value otherwise.

### GetTargetPodOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) GetTargetPodOk() (*string, bool)`

GetTargetPodOk returns a tuple with the TargetPod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) SetTargetPod(v string)`

SetTargetPod sets TargetPod field to given value.

### HasTargetPod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState) HasTargetPod() bool`

HasTargetPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


