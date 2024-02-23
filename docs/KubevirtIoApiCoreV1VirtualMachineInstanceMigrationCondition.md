# KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | Pointer to **string** |  | [optional] [default to ""]
**LastTransitionTime** | Pointer to **string** |  | [optional] [default to ""]
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | [default to ""]
**Type** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition(status string, type_ string, ) *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationConditionWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationConditionWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationConditionWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastProbeTime() string`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastProbeTimeOk() (*string, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetLastProbeTime(v string)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


