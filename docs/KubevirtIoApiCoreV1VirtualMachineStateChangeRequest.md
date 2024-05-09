# KubevirtIoApiCoreV1VirtualMachineStateChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Indicates the type of action that is requested. e.g. Start or Stop | [default to ""]
**Data** | Pointer to **map[string]string** | Provides additional data in order to perform the Action | [optional] 
**Uid** | Pointer to **string** | Indicates the UUID of an existing Virtual Machine Instance that this change request applies to -- if applicable | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequest

`func NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequest(action string, ) *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest`

NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequest instantiates a new KubevirtIoApiCoreV1VirtualMachineStateChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequestWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequestWithDefaults() *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest`

NewKubevirtIoApiCoreV1VirtualMachineStateChangeRequestWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineStateChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetData

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUid

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *KubevirtIoApiCoreV1VirtualMachineStateChangeRequest) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


