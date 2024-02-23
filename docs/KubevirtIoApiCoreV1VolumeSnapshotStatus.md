# KubevirtIoApiCoreV1VolumeSnapshotStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**Name** | **string** |  | [default to ""]
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VolumeSnapshotStatus

`func NewKubevirtIoApiCoreV1VolumeSnapshotStatus(enabled bool, name string, ) *KubevirtIoApiCoreV1VolumeSnapshotStatus`

NewKubevirtIoApiCoreV1VolumeSnapshotStatus instantiates a new KubevirtIoApiCoreV1VolumeSnapshotStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VolumeSnapshotStatusWithDefaults

`func NewKubevirtIoApiCoreV1VolumeSnapshotStatusWithDefaults() *KubevirtIoApiCoreV1VolumeSnapshotStatus`

NewKubevirtIoApiCoreV1VolumeSnapshotStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VolumeSnapshotStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubevirtIoApiCoreV1VolumeSnapshotStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


