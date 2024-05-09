# NetworkHarvesterhciIoV1beta1Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastTransitionTime** | Pointer to **string** | Last time the condition transitioned from one status to another. | [optional] 
**LastUpdateTime** | Pointer to **string** | The last time this condition was updated. | [optional] 
**Message** | Pointer to **string** | Human-readable message indicating details about last transition | [optional] 
**Reason** | Pointer to **string** | The reason for the condition&#39;s last transition. | [optional] 
**Status** | **string** | Status of the condition, one of True, False, Unknown. | [default to ""]
**Type** | **string** | Type of the condition. | [default to ""]

## Methods

### NewNetworkHarvesterhciIoV1beta1Condition

`func NewNetworkHarvesterhciIoV1beta1Condition(status string, type_ string, ) *NetworkHarvesterhciIoV1beta1Condition`

NewNetworkHarvesterhciIoV1beta1Condition instantiates a new NetworkHarvesterhciIoV1beta1Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1ConditionWithDefaults

`func NewNetworkHarvesterhciIoV1beta1ConditionWithDefaults() *NetworkHarvesterhciIoV1beta1Condition`

NewNetworkHarvesterhciIoV1beta1ConditionWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastTransitionTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetLastUpdateTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetLastUpdateTime() string`

GetLastUpdateTime returns the LastUpdateTime field if non-nil, zero value otherwise.

### GetLastUpdateTimeOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetLastUpdateTimeOk() (*string, bool)`

GetLastUpdateTimeOk returns a tuple with the LastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetLastUpdateTime(v string)`

SetLastUpdateTime sets LastUpdateTime field to given value.

### HasLastUpdateTime

`func (o *NetworkHarvesterhciIoV1beta1Condition) HasLastUpdateTime() bool`

HasLastUpdateTime returns a boolean if a field has been set.

### GetMessage

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NetworkHarvesterhciIoV1beta1Condition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NetworkHarvesterhciIoV1beta1Condition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkHarvesterhciIoV1beta1Condition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkHarvesterhciIoV1beta1Condition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


