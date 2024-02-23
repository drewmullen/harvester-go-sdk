# KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition.md) |  | [optional] 
**MigrationState** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**PhaseTransitionTimestamps** | Pointer to [**[]KubevirtIoApiCoreV1VirtualMachineInstanceMigrationPhaseTransitionTimestamp**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationPhaseTransitionTimestamp.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatusWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatusWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetConditions() []KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetConditionsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) SetConditions(v []KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetMigrationState() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState`

GetMigrationState returns the MigrationState field if non-nil, zero value otherwise.

### GetMigrationStateOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetMigrationStateOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState, bool)`

GetMigrationStateOk returns a tuple with the MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) SetMigrationState(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationState)`

SetMigrationState sets MigrationState field to given value.

### HasMigrationState

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) HasMigrationState() bool`

HasMigrationState returns a boolean if a field has been set.

### GetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetPhaseTransitionTimestamps() []KubevirtIoApiCoreV1VirtualMachineInstanceMigrationPhaseTransitionTimestamp`

GetPhaseTransitionTimestamps returns the PhaseTransitionTimestamps field if non-nil, zero value otherwise.

### GetPhaseTransitionTimestampsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) GetPhaseTransitionTimestampsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstanceMigrationPhaseTransitionTimestamp, bool)`

GetPhaseTransitionTimestampsOk returns a tuple with the PhaseTransitionTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) SetPhaseTransitionTimestamps(v []KubevirtIoApiCoreV1VirtualMachineInstanceMigrationPhaseTransitionTimestamp)`

SetPhaseTransitionTimestamps sets PhaseTransitionTimestamps field to given value.

### HasPhaseTransitionTimestamps

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus) HasPhaseTransitionTimestamps() bool`

HasPhaseTransitionTimestamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


