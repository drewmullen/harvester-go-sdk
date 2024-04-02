# K8sIoV1PersistentVolumeClaimCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProbeTime** | Pointer to **string** |  | [optional] [default to "{}"]
**LastTransitionTime** | Pointer to **string** |  | [optional] [default to "{}"]
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | [default to ""]
**Type** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1PersistentVolumeClaimCondition

`func NewK8sIoV1PersistentVolumeClaimCondition(status string, type_ string, ) *K8sIoV1PersistentVolumeClaimCondition`

NewK8sIoV1PersistentVolumeClaimCondition instantiates a new K8sIoV1PersistentVolumeClaimCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimConditionWithDefaults

`func NewK8sIoV1PersistentVolumeClaimConditionWithDefaults() *K8sIoV1PersistentVolumeClaimCondition`

NewK8sIoV1PersistentVolumeClaimConditionWithDefaults instantiates a new K8sIoV1PersistentVolumeClaimCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProbeTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetLastProbeTime() string`

GetLastProbeTime returns the LastProbeTime field if non-nil, zero value otherwise.

### GetLastProbeTimeOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetLastProbeTimeOk() (*string, bool)`

GetLastProbeTimeOk returns a tuple with the LastProbeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProbeTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetLastProbeTime(v string)`

SetLastProbeTime sets LastProbeTime field to given value.

### HasLastProbeTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) HasLastProbeTime() bool`

HasLastProbeTime returns a boolean if a field has been set.

### GetLastTransitionTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.

### HasLastTransitionTime

`func (o *K8sIoV1PersistentVolumeClaimCondition) HasLastTransitionTime() bool`

HasLastTransitionTime returns a boolean if a field has been set.

### GetMessage

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *K8sIoV1PersistentVolumeClaimCondition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *K8sIoV1PersistentVolumeClaimCondition) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *K8sIoV1PersistentVolumeClaimCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *K8sIoV1PersistentVolumeClaimCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


