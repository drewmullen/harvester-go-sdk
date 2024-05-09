# KubevirtIoApiCoreV1VirtualMachineStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineCondition**](KubevirtIoApiCoreV1VirtualMachineCondition.md) | Hold the state information of the VirtualMachine and its VirtualMachineInstance | [optional] 
**Created** | Pointer to **bool** | Created indicates if the virtual machine is created in the cluster | [optional] 
**DesiredGeneration** | Pointer to **int64** | DesiredGeneration is the generation which is desired for the VMI. This will be used in comparisons with ObservedGeneration to understand when the VMI is out of sync. This will be changed at the same time as ObservedGeneration to remove errors which could occur if Generation is updated through an Update() before ObservedGeneration in Status. | [optional] 
**MemoryDumpRequest** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest**](KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest.md) | MemoryDumpRequest tracks memory dump request phase and info of getting a memory dump to the given pvc | [optional] 
**ObservedGeneration** | Pointer to **int64** | ObservedGeneration is the generation observed by the vmi when started. | [optional] 
**PrintableStatus** | Pointer to **string** | PrintableStatus is a human readable, high-level representation of the status of the virtual machine | [optional] 
**Ready** | Pointer to **bool** | Ready indicates if the virtual machine is running and ready | [optional] 
**RestoreInProgress** | Pointer to **string** | RestoreInProgress is the name of the VirtualMachineRestore currently executing | [optional] 
**SnapshotInProgress** | Pointer to **string** | SnapshotInProgress is the name of the VirtualMachineSnapshot currently executing | [optional] 
**StartFailure** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineStartFailure**](KubevirtIoApiCoreV1VirtualMachineStartFailure.md) | StartFailure tracks consecutive VMI startup failures for the purposes of crash loop backoffs | [optional] 
**StateChangeRequests** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineStateChangeRequest**](KubevirtIoApiCoreV1VirtualMachineStateChangeRequest.md) | StateChangeRequests indicates a list of actions that should be taken on a VMI e.g. stop a specific VMI then start a new one. | [optional] 
**VolumeRequests** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineVolumeRequest**](KubevirtIoApiCoreV1VirtualMachineVolumeRequest.md) | VolumeRequests indicates a list of volumes add or remove from the VMI template and hotplug on an active running VMI. | [optional] 
**VolumeSnapshotStatuses** | Pointer to [**[]KubevirtIoApiCoreV1VolumeSnapshotStatus**](KubevirtIoApiCoreV1VolumeSnapshotStatus.md) | VolumeSnapshotStatuses indicates a list of statuses whether snapshotting is supported by each volume. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineStatus

`func NewKubevirtIoApiCoreV1VirtualMachineStatus() *KubevirtIoApiCoreV1VirtualMachineStatus`

NewKubevirtIoApiCoreV1VirtualMachineStatus instantiates a new KubevirtIoApiCoreV1VirtualMachineStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineStatusWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineStatusWithDefaults() *KubevirtIoApiCoreV1VirtualMachineStatus`

NewKubevirtIoApiCoreV1VirtualMachineStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetConditions() []KubevirtIoApiCoreV1VirtualMachineCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetConditionsOk() (*[]KubevirtIoApiCoreV1VirtualMachineCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetConditions(v []KubevirtIoApiCoreV1VirtualMachineCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreated

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDesiredGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetDesiredGeneration() int64`

GetDesiredGeneration returns the DesiredGeneration field if non-nil, zero value otherwise.

### GetDesiredGenerationOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetDesiredGenerationOk() (*int64, bool)`

GetDesiredGenerationOk returns a tuple with the DesiredGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetDesiredGeneration(v int64)`

SetDesiredGeneration sets DesiredGeneration field to given value.

### HasDesiredGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasDesiredGeneration() bool`

HasDesiredGeneration returns a boolean if a field has been set.

### GetMemoryDumpRequest

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetMemoryDumpRequest() KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest`

GetMemoryDumpRequest returns the MemoryDumpRequest field if non-nil, zero value otherwise.

### GetMemoryDumpRequestOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetMemoryDumpRequestOk() (*KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest, bool)`

GetMemoryDumpRequestOk returns a tuple with the MemoryDumpRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDumpRequest

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetMemoryDumpRequest(v KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest)`

SetMemoryDumpRequest sets MemoryDumpRequest field to given value.

### HasMemoryDumpRequest

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasMemoryDumpRequest() bool`

HasMemoryDumpRequest returns a boolean if a field has been set.

### GetObservedGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetObservedGeneration() int64`

GetObservedGeneration returns the ObservedGeneration field if non-nil, zero value otherwise.

### GetObservedGenerationOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetObservedGenerationOk() (*int64, bool)`

GetObservedGenerationOk returns a tuple with the ObservedGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetObservedGeneration(v int64)`

SetObservedGeneration sets ObservedGeneration field to given value.

### HasObservedGeneration

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasObservedGeneration() bool`

HasObservedGeneration returns a boolean if a field has been set.

### GetPrintableStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetPrintableStatus() string`

GetPrintableStatus returns the PrintableStatus field if non-nil, zero value otherwise.

### GetPrintableStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetPrintableStatusOk() (*string, bool)`

GetPrintableStatusOk returns a tuple with the PrintableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintableStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetPrintableStatus(v string)`

SetPrintableStatus sets PrintableStatus field to given value.

### HasPrintableStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasPrintableStatus() bool`

HasPrintableStatus returns a boolean if a field has been set.

### GetReady

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetReady(v bool)`

SetReady sets Ready field to given value.

### HasReady

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasReady() bool`

HasReady returns a boolean if a field has been set.

### GetRestoreInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetRestoreInProgress() string`

GetRestoreInProgress returns the RestoreInProgress field if non-nil, zero value otherwise.

### GetRestoreInProgressOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetRestoreInProgressOk() (*string, bool)`

GetRestoreInProgressOk returns a tuple with the RestoreInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetRestoreInProgress(v string)`

SetRestoreInProgress sets RestoreInProgress field to given value.

### HasRestoreInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasRestoreInProgress() bool`

HasRestoreInProgress returns a boolean if a field has been set.

### GetSnapshotInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetSnapshotInProgress() string`

GetSnapshotInProgress returns the SnapshotInProgress field if non-nil, zero value otherwise.

### GetSnapshotInProgressOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetSnapshotInProgressOk() (*string, bool)`

GetSnapshotInProgressOk returns a tuple with the SnapshotInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetSnapshotInProgress(v string)`

SetSnapshotInProgress sets SnapshotInProgress field to given value.

### HasSnapshotInProgress

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasSnapshotInProgress() bool`

HasSnapshotInProgress returns a boolean if a field has been set.

### GetStartFailure

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStartFailure() KubevirtIoApiCoreV1VirtualMachineStartFailure`

GetStartFailure returns the StartFailure field if non-nil, zero value otherwise.

### GetStartFailureOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStartFailureOk() (*KubevirtIoApiCoreV1VirtualMachineStartFailure, bool)`

GetStartFailureOk returns a tuple with the StartFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFailure

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetStartFailure(v KubevirtIoApiCoreV1VirtualMachineStartFailure)`

SetStartFailure sets StartFailure field to given value.

### HasStartFailure

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasStartFailure() bool`

HasStartFailure returns a boolean if a field has been set.

### GetStateChangeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStateChangeRequests() []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest`

GetStateChangeRequests returns the StateChangeRequests field if non-nil, zero value otherwise.

### GetStateChangeRequestsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetStateChangeRequestsOk() (*[]KubevirtIoApiCoreV1VirtualMachineStateChangeRequest, bool)`

GetStateChangeRequestsOk returns a tuple with the StateChangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetStateChangeRequests(v []KubevirtIoApiCoreV1VirtualMachineStateChangeRequest)`

SetStateChangeRequests sets StateChangeRequests field to given value.

### HasStateChangeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasStateChangeRequests() bool`

HasStateChangeRequests returns a boolean if a field has been set.

### GetVolumeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeRequests() []KubevirtIoApiCoreV1VirtualMachineVolumeRequest`

GetVolumeRequests returns the VolumeRequests field if non-nil, zero value otherwise.

### GetVolumeRequestsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeRequestsOk() (*[]KubevirtIoApiCoreV1VirtualMachineVolumeRequest, bool)`

GetVolumeRequestsOk returns a tuple with the VolumeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetVolumeRequests(v []KubevirtIoApiCoreV1VirtualMachineVolumeRequest)`

SetVolumeRequests sets VolumeRequests field to given value.

### HasVolumeRequests

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasVolumeRequests() bool`

HasVolumeRequests returns a boolean if a field has been set.

### GetVolumeSnapshotStatuses

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeSnapshotStatuses() []KubevirtIoApiCoreV1VolumeSnapshotStatus`

GetVolumeSnapshotStatuses returns the VolumeSnapshotStatuses field if non-nil, zero value otherwise.

### GetVolumeSnapshotStatusesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) GetVolumeSnapshotStatusesOk() (*[]KubevirtIoApiCoreV1VolumeSnapshotStatus, bool)`

GetVolumeSnapshotStatusesOk returns a tuple with the VolumeSnapshotStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSnapshotStatuses

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) SetVolumeSnapshotStatuses(v []KubevirtIoApiCoreV1VolumeSnapshotStatus)`

SetVolumeSnapshotStatuses sets VolumeSnapshotStatuses field to given value.

### HasVolumeSnapshotStatuses

`func (o *KubevirtIoApiCoreV1VirtualMachineStatus) HasVolumeSnapshotStatuses() bool`

HasVolumeSnapshotStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


