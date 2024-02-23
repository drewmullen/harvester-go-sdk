# KubevirtIoApiCoreV1VolumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HotplugVolume** | Pointer to [**KubevirtIoApiCoreV1HotplugVolumeStatus**](KubevirtIoApiCoreV1HotplugVolumeStatus.md) |  | [optional] 
**MemoryDumpVolume** | Pointer to [**KubevirtIoApiCoreV1DomainMemoryDumpInfo**](KubevirtIoApiCoreV1DomainMemoryDumpInfo.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | [default to ""]
**PersistentVolumeClaimInfo** | Pointer to [**KubevirtIoApiCoreV1PersistentVolumeClaimInfo**](KubevirtIoApiCoreV1PersistentVolumeClaimInfo.md) |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**Target** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1VolumeStatus

`func NewKubevirtIoApiCoreV1VolumeStatus(name string, target string, ) *KubevirtIoApiCoreV1VolumeStatus`

NewKubevirtIoApiCoreV1VolumeStatus instantiates a new KubevirtIoApiCoreV1VolumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VolumeStatusWithDefaults

`func NewKubevirtIoApiCoreV1VolumeStatusWithDefaults() *KubevirtIoApiCoreV1VolumeStatus`

NewKubevirtIoApiCoreV1VolumeStatusWithDefaults instantiates a new KubevirtIoApiCoreV1VolumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHotplugVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetHotplugVolume() KubevirtIoApiCoreV1HotplugVolumeStatus`

GetHotplugVolume returns the HotplugVolume field if non-nil, zero value otherwise.

### GetHotplugVolumeOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetHotplugVolumeOk() (*KubevirtIoApiCoreV1HotplugVolumeStatus, bool)`

GetHotplugVolumeOk returns a tuple with the HotplugVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplugVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetHotplugVolume(v KubevirtIoApiCoreV1HotplugVolumeStatus)`

SetHotplugVolume sets HotplugVolume field to given value.

### HasHotplugVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasHotplugVolume() bool`

HasHotplugVolume returns a boolean if a field has been set.

### GetMemoryDumpVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetMemoryDumpVolume() KubevirtIoApiCoreV1DomainMemoryDumpInfo`

GetMemoryDumpVolume returns the MemoryDumpVolume field if non-nil, zero value otherwise.

### GetMemoryDumpVolumeOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetMemoryDumpVolumeOk() (*KubevirtIoApiCoreV1DomainMemoryDumpInfo, bool)`

GetMemoryDumpVolumeOk returns a tuple with the MemoryDumpVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDumpVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetMemoryDumpVolume(v KubevirtIoApiCoreV1DomainMemoryDumpInfo)`

SetMemoryDumpVolume sets MemoryDumpVolume field to given value.

### HasMemoryDumpVolume

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasMemoryDumpVolume() bool`

HasMemoryDumpVolume returns a boolean if a field has been set.

### GetMessage

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetName(v string)`

SetName sets Name field to given value.


### GetPersistentVolumeClaimInfo

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetPersistentVolumeClaimInfo() KubevirtIoApiCoreV1PersistentVolumeClaimInfo`

GetPersistentVolumeClaimInfo returns the PersistentVolumeClaimInfo field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimInfoOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetPersistentVolumeClaimInfoOk() (*KubevirtIoApiCoreV1PersistentVolumeClaimInfo, bool)`

GetPersistentVolumeClaimInfoOk returns a tuple with the PersistentVolumeClaimInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaimInfo

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetPersistentVolumeClaimInfo(v KubevirtIoApiCoreV1PersistentVolumeClaimInfo)`

SetPersistentVolumeClaimInfo sets PersistentVolumeClaimInfo field to given value.

### HasPersistentVolumeClaimInfo

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasPersistentVolumeClaimInfo() bool`

HasPersistentVolumeClaimInfo returns a boolean if a field has been set.

### GetPhase

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetReason

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSize

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *KubevirtIoApiCoreV1VolumeStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTarget

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *KubevirtIoApiCoreV1VolumeStatus) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *KubevirtIoApiCoreV1VolumeStatus) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


