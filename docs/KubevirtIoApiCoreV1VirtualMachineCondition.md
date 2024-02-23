# KubevirtIoApiCoreV1VirtualMachineCondition

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

### NewKubevirtIoApiCoreV1VirtualMachineCondition

`func NewKubevirtIoApiCoreV1VirtualMachineCondition(status string, type_ string, ) *KubevirtIoApiCoreV1VirtualMachineCondition`

NewKubevirtIoApiCoreV1VirtualMachineCondition instantiates a new KubevirtIoApiCoreV1VirtualMachineCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineConditionWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineConditionWithDefaults() *KubevirtIoApiCoreV1VirtualMachineCondition`

NewKubevirtIoApiCoreV1VirtualMachineConditionWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetLastProbeTime() string`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetLastProbeTimeOk() (*string, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetLastProbeTime(v string)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubevirtIoApiCoreV1VirtualMachineCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


