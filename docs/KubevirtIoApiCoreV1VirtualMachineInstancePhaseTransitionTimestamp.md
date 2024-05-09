# KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | Pointer to **string** | Phase is the status of the VirtualMachineInstance in kubernetes world. It is not the VirtualMachineInstance status, but partially correlates to it. | [optional] 
**PhaseTransitionTimestamp** | Pointer to **string** | PhaseTransitionTimestamp is the timestamp of when the phase change occurred | [optional] [default to "{}"]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp

`func NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp() *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp`

NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp instantiates a new KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestampWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestampWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp`

NewKubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestampWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPhaseTransitionTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseTransitionTimestamp() string`

GetPhaseTransitionTimestamp returns the PhaseTransitionTimestamp field if non-nil, zero value otherwise.

### GetPhaseTransitionTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) GetPhaseTransitionTimestampOk() (*string, bool)`

GetPhaseTransitionTimestampOk returns a tuple with the PhaseTransitionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseTransitionTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) SetPhaseTransitionTimestamp(v string)`

SetPhaseTransitionTimestamp sets PhaseTransitionTimestamp field to given value.

### HasPhaseTransitionTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineInstancePhaseTransitionTimestamp) HasPhaseTransitionTimestamp() bool`

HasPhaseTransitionTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


