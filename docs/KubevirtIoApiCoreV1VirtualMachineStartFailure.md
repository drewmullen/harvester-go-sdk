# KubevirtIoApiCoreV1VirtualMachineStartFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsecutiveFailCount** | Pointer to **int32** |  | [optional] 
**LastFailedVMIUID** | Pointer to **string** |  | [optional] 
**RetryAfterTimestamp** | Pointer to **string** | Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers. | [optional] [default to ""]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineStartFailure

`func NewKubevirtIoApiCoreV1VirtualMachineStartFailure() *KubevirtIoApiCoreV1VirtualMachineStartFailure`

NewKubevirtIoApiCoreV1VirtualMachineStartFailure instantiates a new KubevirtIoApiCoreV1VirtualMachineStartFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineStartFailureWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineStartFailureWithDefaults() *KubevirtIoApiCoreV1VirtualMachineStartFailure`

NewKubevirtIoApiCoreV1VirtualMachineStartFailureWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineStartFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsecutiveFailCount

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetConsecutiveFailCount() int32`

GetConsecutiveFailCount returns the ConsecutiveFailCount field if non-nil, zero value otherwise.

### GetConsecutiveFailCountOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetConsecutiveFailCountOk() (*int32, bool)`

GetConsecutiveFailCountOk returns a tuple with the ConsecutiveFailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailCount

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) SetConsecutiveFailCount(v int32)`

SetConsecutiveFailCount sets ConsecutiveFailCount field to given value.

### HasConsecutiveFailCount

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) HasConsecutiveFailCount() bool`

HasConsecutiveFailCount returns a boolean if a field has been set.

### GetLastFailedVMIUID

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetLastFailedVMIUID() string`

GetLastFailedVMIUID returns the LastFailedVMIUID field if non-nil, zero value otherwise.

### GetLastFailedVMIUIDOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetLastFailedVMIUIDOk() (*string, bool)`

GetLastFailedVMIUIDOk returns a tuple with the LastFailedVMIUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedVMIUID

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) SetLastFailedVMIUID(v string)`

SetLastFailedVMIUID sets LastFailedVMIUID field to given value.

### HasLastFailedVMIUID

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) HasLastFailedVMIUID() bool`

HasLastFailedVMIUID returns a boolean if a field has been set.

### GetRetryAfterTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetRetryAfterTimestamp() string`

GetRetryAfterTimestamp returns the RetryAfterTimestamp field if non-nil, zero value otherwise.

### GetRetryAfterTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) GetRetryAfterTimestampOk() (*string, bool)`

GetRetryAfterTimestampOk returns a tuple with the RetryAfterTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) SetRetryAfterTimestamp(v string)`

SetRetryAfterTimestamp sets RetryAfterTimestamp field to given value.

### HasRetryAfterTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineStartFailure) HasRetryAfterTimestamp() bool`

HasRetryAfterTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


