# KubevirtIoApiCoreV1VirtualMachineInstanceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VSOCKCID** | Pointer to **int64** | VSOCKCID is used to track the allocated VSOCK CID in the VM. | [optional] 
**ActivePods** | Pointer to **map[string]string** | ActivePods is a mapping of pod UID to node name. It is possible for multiple pods to be running for a single VMI during migration. | [optional] 
**Conditions** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineInstanceCondition**](KubevirtIoApiCoreV1VirtualMachineInstanceCondition.md) | Conditions are specific points in VirtualMachineInstance&#39;s pod runtime. | [optional] 
**CurrentCPUTopology** | Pointer to [**KubevirtIoApiCoreV1CPUTopology**](KubevirtIoApiCoreV1CPUTopology.md) | CurrentCPUTopology specifies the current CPU topology used by the VM workload. Current topology may differ from the desired topology in the spec while CPU hotplug takes place. | [optional] 
**EvacuationNodeName** | Pointer to **string** | EvacuationNodeName is used to track the eviction process of a VMI. It stores the name of the node that we want to evacuate. It is meant to be used by KubeVirt core components only and can&#39;t be set or modified by users. | [optional] 
**FsFreezeStatus** | Pointer to **string** | FSFreezeStatus is the state of the fs of the guest it can be either frozen or thawed | [optional] 
**GuestOSInfo** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceGuestOSInfo**](KubevirtIoApiCoreV1VirtualMachineInstanceGuestOSInfo.md) | Guest OS Information | [optional] [default to {}]
**Interfaces** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface**](KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface.md) | Interfaces represent the details of available network interfaces. | [optional] 
**KernelBootStatus** | Pointer to [**KubevirtIoApiCoreV1KernelBootStatus**](KubevirtIoApiCoreV1KernelBootStatus.md) | KernelBootStatus contains info about the kernelBootContainer | [optional] 
**LauncherContainerImageVersion** | Pointer to **string** | LauncherContainerImageVersion indicates what container image is currently active for the vmi. | [optional] 
**Machine** | Pointer to [**KubevirtIoApiCoreV1Machine**](KubevirtIoApiCoreV1Machine.md) | Machine shows the final resulting qemu machine type. This can be different than the machine type selected in the spec, due to qemus machine type alias mechanism. | [optional] 
**Memory** | Pointer to [**KubevirtIoApiCoreV1MemoryStatus**](KubevirtIoApiCoreV1MemoryStatus.md) | Memory shows various informations about the VirtualMachine memory. | [optional] 
**MigrationMethod** | Pointer to **string** | Represents the method using which the vmi can be migrated: live migration or block migration | [optional] 
**MigrationState** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState.md) | Represents the status of a live migration | [optional] 
**MigrationTransport** | Pointer to **string** | This represents the migration transport | [optional] 
**NodeName** | Pointer to **string** | NodeName is the name where the VirtualMachineInstance is currently running. | [optional] 
**Phase** | Pointer to **string** | Phase is the status of the VirtualMachineInstance in kubernetes world. It is not the VirtualMachineInstance status, but partially correlates to it. | [optional] 
**PhaseTransitionTimestamps** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp**](KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp.md) | PhaseTransitionTimestamp is the timestamp of when the last phase change occurred | [optional] 
**QosClass** | Pointer to **string** | The Quality of Service (QOS) classification assigned to the virtual machine instance based on resource requirements See PodQOSClass type for available QOS classes More info: https://git.k8s.io/community/contributors/design-proposals/node/resource-qos.md | [optional] 
**Reason** | Pointer to **string** | A brief CamelCase message indicating details about why the VMI is in this state. e.g. &#39;NodeUnresponsive&#39; | [optional] 
**RuntimeUser** | Pointer to **int64** | RuntimeUser is used to determine what user will be used in launcher | [optional] [default to 0]
**SelinuxContext** | Pointer to **string** | SELinuxContext is the actual SELinux context of the virt-launcher pod | [optional] 
**TopologyHints** | Pointer to [**KubevirtIoApiCoreV1TopologyHints**](KubevirtIoApiCoreV1TopologyHints.md) |  | [optional] 
**VirtualMachineRevisionName** | Pointer to **string** | VirtualMachineRevisionName is used to get the vm revision of the vmi when doing an online vm snapshot | [optional] 
**VolumeStatus** | Pointer to [**[]KubevirtIoApiCoreV1VolumeStatus**](KubevirtIoApiCoreV1VolumeStatus.md) | VolumeStatus contains the statuses of all the volumes | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceStatus

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceStatus() *KubevirtIoApiCoreV1VirtualMachineInstanceStatus`

NewKubevirtIoApiCoreV1VirtualMachineInstanceStatus instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceStatusWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceStatusWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceStatus`

NewKubevirtIoApiCoreV1VirtualMachineInstanceStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVSOCKCID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVSOCKCID() int64`

GetVSOCKCID returns the VSOCKCID field if non-nil, zero value otherwise.

### GetVSOCKCIDOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVSOCKCIDOk() (*int64, bool)`

GetVSOCKCIDOk returns a tuple with the VSOCKCID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSOCKCID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetVSOCKCID(v int64)`

SetVSOCKCID sets VSOCKCID field to given value.

### HasVSOCKCID

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasVSOCKCID() bool`

HasVSOCKCID returns a boolean if a field has been set.

### GetActivePods

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetActivePods() map[string]string`

GetActivePods returns the ActivePods field if non-nil, zero value otherwise.

### GetActivePodsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetActivePodsOk() (*map[string]string, bool)`

GetActivePodsOk returns a tuple with the ActivePods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePods

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetActivePods(v map[string]string)`

SetActivePods sets ActivePods field to given value.

### HasActivePods

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasActivePods() bool`

HasActivePods returns a boolean if a field has been set.

### GetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetConditions() []KubevirtIoApiCoreV1VirtualMachineInstanceCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetConditionsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstanceCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetConditions(v []KubevirtIoApiCoreV1VirtualMachineInstanceCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCurrentCPUTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetCurrentCPUTopology() KubevirtIoApiCoreV1CPUTopology`

GetCurrentCPUTopology returns the CurrentCPUTopology field if non-nil, zero value otherwise.

### GetCurrentCPUTopologyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetCurrentCPUTopologyOk() (*KubevirtIoApiCoreV1CPUTopology, bool)`

GetCurrentCPUTopologyOk returns a tuple with the CurrentCPUTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCPUTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetCurrentCPUTopology(v KubevirtIoApiCoreV1CPUTopology)`

SetCurrentCPUTopology sets CurrentCPUTopology field to given value.

### HasCurrentCPUTopology

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasCurrentCPUTopology() bool`

HasCurrentCPUTopology returns a boolean if a field has been set.

### GetEvacuationNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetEvacuationNodeName() string`

GetEvacuationNodeName returns the EvacuationNodeName field if non-nil, zero value otherwise.

### GetEvacuationNodeNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetEvacuationNodeNameOk() (*string, bool)`

GetEvacuationNodeNameOk returns a tuple with the EvacuationNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvacuationNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetEvacuationNodeName(v string)`

SetEvacuationNodeName sets EvacuationNodeName field to given value.

### HasEvacuationNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasEvacuationNodeName() bool`

HasEvacuationNodeName returns a boolean if a field has been set.

### GetFsFreezeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetFsFreezeStatus() string`

GetFsFreezeStatus returns the FsFreezeStatus field if non-nil, zero value otherwise.

### GetFsFreezeStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetFsFreezeStatusOk() (*string, bool)`

GetFsFreezeStatusOk returns a tuple with the FsFreezeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsFreezeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetFsFreezeStatus(v string)`

SetFsFreezeStatus sets FsFreezeStatus field to given value.

### HasFsFreezeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasFsFreezeStatus() bool`

HasFsFreezeStatus returns a boolean if a field has been set.

### GetGuestOSInfo

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetGuestOSInfo() KubevirtIoApiCoreV1VirtualMachineInstanceGuestOSInfo`

GetGuestOSInfo returns the GuestOSInfo field if non-nil, zero value otherwise.

### GetGuestOSInfoOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetGuestOSInfoOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceGuestOSInfo, bool)`

GetGuestOSInfoOk returns a tuple with the GuestOSInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOSInfo

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetGuestOSInfo(v KubevirtIoApiCoreV1VirtualMachineInstanceGuestOSInfo)`

SetGuestOSInfo sets GuestOSInfo field to given value.

### HasGuestOSInfo

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasGuestOSInfo() bool`

HasGuestOSInfo returns a boolean if a field has been set.

### GetInterfaces

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetInterfaces() []KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetInterfacesOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetInterfaces(v []KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetKernelBootStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetKernelBootStatus() KubevirtIoApiCoreV1KernelBootStatus`

GetKernelBootStatus returns the KernelBootStatus field if non-nil, zero value otherwise.

### GetKernelBootStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetKernelBootStatusOk() (*KubevirtIoApiCoreV1KernelBootStatus, bool)`

GetKernelBootStatusOk returns a tuple with the KernelBootStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelBootStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetKernelBootStatus(v KubevirtIoApiCoreV1KernelBootStatus)`

SetKernelBootStatus sets KernelBootStatus field to given value.

### HasKernelBootStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasKernelBootStatus() bool`

HasKernelBootStatus returns a boolean if a field has been set.

### GetLauncherContainerImageVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetLauncherContainerImageVersion() string`

GetLauncherContainerImageVersion returns the LauncherContainerImageVersion field if non-nil, zero value otherwise.

### GetLauncherContainerImageVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetLauncherContainerImageVersionOk() (*string, bool)`

GetLauncherContainerImageVersionOk returns a tuple with the LauncherContainerImageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLauncherContainerImageVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetLauncherContainerImageVersion(v string)`

SetLauncherContainerImageVersion sets LauncherContainerImageVersion field to given value.

### HasLauncherContainerImageVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasLauncherContainerImageVersion() bool`

HasLauncherContainerImageVersion returns a boolean if a field has been set.

### GetMachine

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMachine() KubevirtIoApiCoreV1Machine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMachineOk() (*KubevirtIoApiCoreV1Machine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetMachine(v KubevirtIoApiCoreV1Machine)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMemory() KubevirtIoApiCoreV1MemoryStatus`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMemoryOk() (*KubevirtIoApiCoreV1MemoryStatus, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetMemory(v KubevirtIoApiCoreV1MemoryStatus)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrationMethod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationMethod() string`

GetMigrationMethod returns the MigrationMethod field if non-nil, zero value otherwise.

### GetMigrationMethodOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationMethodOk() (*string, bool)`

GetMigrationMethodOk returns a tuple with the MigrationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationMethod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetMigrationMethod(v string)`

SetMigrationMethod sets MigrationMethod field to given value.

### HasMigrationMethod

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasMigrationMethod() bool`

HasMigrationMethod returns a boolean if a field has been set.

### GetMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationState() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState`

GetMigrationState returns the MigrationState field if non-nil, zero value otherwise.

### GetMigrationStateOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationStateOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState, bool)`

GetMigrationStateOk returns a tuple with the MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetMigrationState(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState)`

SetMigrationState sets MigrationState field to given value.

### HasMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasMigrationState() bool`

HasMigrationState returns a boolean if a field has been set.

### GetMigrationTransport

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationTransport() string`

GetMigrationTransport returns the MigrationTransport field if non-nil, zero value otherwise.

### GetMigrationTransportOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetMigrationTransportOk() (*string, bool)`

GetMigrationTransportOk returns a tuple with the MigrationTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationTransport

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetMigrationTransport(v string)`

SetMigrationTransport sets MigrationTransport field to given value.

### HasMigrationTransport

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasMigrationTransport() bool`

HasMigrationTransport returns a boolean if a field has been set.

### GetNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetPhaseTransitionTimestamps() []KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp`

GetPhaseTransitionTimestamps returns the PhaseTransitionTimestamps field if non-nil, zero value otherwise.

### GetPhaseTransitionTimestampsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetPhaseTransitionTimestampsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp, bool)`

GetPhaseTransitionTimestampsOk returns a tuple with the PhaseTransitionTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetPhaseTransitionTimestamps(v []KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp)`

SetPhaseTransitionTimestamps sets PhaseTransitionTimestamps field to given value.

### HasPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasPhaseTransitionTimestamps() bool`

HasPhaseTransitionTimestamps returns a boolean if a field has been set.

### GetQosClass

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.

### HasQosClass

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasQosClass() bool`

HasQosClass returns a boolean if a field has been set.

### GetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRuntimeUser

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetRuntimeUser() int64`

GetRuntimeUser returns the RuntimeUser field if non-nil, zero value otherwise.

### GetRuntimeUserOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetRuntimeUserOk() (*int64, bool)`

GetRuntimeUserOk returns a tuple with the RuntimeUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeUser

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetRuntimeUser(v int64)`

SetRuntimeUser sets RuntimeUser field to given value.

### HasRuntimeUser

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasRuntimeUser() bool`

HasRuntimeUser returns a boolean if a field has been set.

### GetSelinuxContext

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetSelinuxContext() string`

GetSelinuxContext returns the SelinuxContext field if non-nil, zero value otherwise.

### GetSelinuxContextOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetSelinuxContextOk() (*string, bool)`

GetSelinuxContextOk returns a tuple with the SelinuxContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelinuxContext

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetSelinuxContext(v string)`

SetSelinuxContext sets SelinuxContext field to given value.

### HasSelinuxContext

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasSelinuxContext() bool`

HasSelinuxContext returns a boolean if a field has been set.

### GetTopologyHints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetTopologyHints() KubevirtIoApiCoreV1TopologyHints`

GetTopologyHints returns the TopologyHints field if non-nil, zero value otherwise.

### GetTopologyHintsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetTopologyHintsOk() (*KubevirtIoApiCoreV1TopologyHints, bool)`

GetTopologyHintsOk returns a tuple with the TopologyHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyHints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetTopologyHints(v KubevirtIoApiCoreV1TopologyHints)`

SetTopologyHints sets TopologyHints field to given value.

### HasTopologyHints

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasTopologyHints() bool`

HasTopologyHints returns a boolean if a field has been set.

### GetVirtualMachineRevisionName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVirtualMachineRevisionName() string`

GetVirtualMachineRevisionName returns the VirtualMachineRevisionName field if non-nil, zero value otherwise.

### GetVirtualMachineRevisionNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVirtualMachineRevisionNameOk() (*string, bool)`

GetVirtualMachineRevisionNameOk returns a tuple with the VirtualMachineRevisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineRevisionName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetVirtualMachineRevisionName(v string)`

SetVirtualMachineRevisionName sets VirtualMachineRevisionName field to given value.

### HasVirtualMachineRevisionName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasVirtualMachineRevisionName() bool`

HasVirtualMachineRevisionName returns a boolean if a field has been set.

### GetVolumeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVolumeStatus() []KubevirtIoApiCoreV1VolumeStatus`

GetVolumeStatus returns the VolumeStatus field if non-nil, zero value otherwise.

### GetVolumeStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) GetVolumeStatusOk() (*[]KubevirtIoApiCoreV1VolumeStatus, bool)`

GetVolumeStatusOk returns a tuple with the VolumeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) SetVolumeStatus(v []KubevirtIoApiCoreV1VolumeStatus)`

SetVolumeStatus sets VolumeStatus field to given value.

### HasVolumeStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceStatus) HasVolumeStatus() bool`

HasVolumeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


