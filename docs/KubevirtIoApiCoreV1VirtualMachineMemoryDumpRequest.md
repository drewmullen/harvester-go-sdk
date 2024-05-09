# KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | **string** | ClaimName is the name of the pvc that will contain the memory dump | [default to ""]
**EndTimestamp** | Pointer to **string** | EndTimestamp represents the time the memory dump was completed | [optional] 
**FileName** | Pointer to **string** | FileName represents the name of the output file | [optional] 
**Message** | Pointer to **string** | Message is a detailed message about failure of the memory dump | [optional] 
**Phase** | **string** | Phase represents the memory dump phase | [default to ""]
**Remove** | Pointer to **bool** | Remove represents request of dissociating the memory dump pvc | [optional] 
**StartTimestamp** | Pointer to **string** | StartTimestamp represents the time the memory dump started | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest

`func NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest(claimName string, phase string, ) *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest`

NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest instantiates a new KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequestWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequestWithDefaults() *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest`

NewKubevirtIoApiCoreV1VirtualMachineMemoryDumpRequestWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetEndTimestamp() string`

GetEndTimestamp returns the EndTimestamp field if non-nil, zero value otherwise.

### GetEndTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetEndTimestampOk() (*string, bool)`

GetEndTimestampOk returns a tuple with the EndTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetEndTimestamp(v string)`

SetEndTimestamp sets EndTimestamp field to given value.

### HasEndTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) HasEndTimestamp() bool`

HasEndTimestamp returns a boolean if a field has been set.

### GetFileName

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetRemove

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetRemove() bool`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetRemoveOk() (*bool, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetRemove(v bool)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) HasRemove() bool`

HasRemove returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *KubevirtIoApiCoreV1VirtualMachineMemoryDumpRequest) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


