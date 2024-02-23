# KubevirtIoApiCoreV1Timer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hpet** | Pointer to [**KubevirtIoApiCoreV1HPETTimer**](KubevirtIoApiCoreV1HPETTimer.md) |  | [optional] 
**Hyperv** | Pointer to [**KubevirtIoApiCoreV1HypervTimer**](KubevirtIoApiCoreV1HypervTimer.md) |  | [optional] 
**Kvm** | Pointer to [**KubevirtIoApiCoreV1KVMTimer**](KubevirtIoApiCoreV1KVMTimer.md) |  | [optional] 
**Pit** | Pointer to [**KubevirtIoApiCoreV1PITTimer**](KubevirtIoApiCoreV1PITTimer.md) |  | [optional] 
**Rtc** | Pointer to [**KubevirtIoApiCoreV1RTCTimer**](KubevirtIoApiCoreV1RTCTimer.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Timer

`func NewKubevirtIoApiCoreV1Timer() *KubevirtIoApiCoreV1Timer`

NewKubevirtIoApiCoreV1Timer instantiates a new KubevirtIoApiCoreV1Timer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1TimerWithDefaults

`func NewKubevirtIoApiCoreV1TimerWithDefaults() *KubevirtIoApiCoreV1Timer`

NewKubevirtIoApiCoreV1TimerWithDefaults instantiates a new KubevirtIoApiCoreV1Timer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHpet

`func (o *KubevirtIoApiCoreV1Timer) GetHpet() KubevirtIoApiCoreV1HPETTimer`

GetHpet returns the Hpet field if non-nil, zero value otherwise.

### GetHpetOk

`func (o *KubevirtIoApiCoreV1Timer) GetHpetOk() (*KubevirtIoApiCoreV1HPETTimer, bool)`

GetHpetOk returns a tuple with the Hpet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpet

`func (o *KubevirtIoApiCoreV1Timer) SetHpet(v KubevirtIoApiCoreV1HPETTimer)`

SetHpet sets Hpet field to given value.

### HasHpet

`func (o *KubevirtIoApiCoreV1Timer) HasHpet() bool`

HasHpet returns a boolean if a field has been set.

### GetHyperv

`func (o *KubevirtIoApiCoreV1Timer) GetHyperv() KubevirtIoApiCoreV1HypervTimer`

GetHyperv returns the Hyperv field if non-nil, zero value otherwise.

### GetHypervOk

`func (o *KubevirtIoApiCoreV1Timer) GetHypervOk() (*KubevirtIoApiCoreV1HypervTimer, bool)`

GetHypervOk returns a tuple with the Hyperv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperv

`func (o *KubevirtIoApiCoreV1Timer) SetHyperv(v KubevirtIoApiCoreV1HypervTimer)`

SetHyperv sets Hyperv field to given value.

### HasHyperv

`func (o *KubevirtIoApiCoreV1Timer) HasHyperv() bool`

HasHyperv returns a boolean if a field has been set.

### GetKvm

`func (o *KubevirtIoApiCoreV1Timer) GetKvm() KubevirtIoApiCoreV1KVMTimer`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *KubevirtIoApiCoreV1Timer) GetKvmOk() (*KubevirtIoApiCoreV1KVMTimer, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *KubevirtIoApiCoreV1Timer) SetKvm(v KubevirtIoApiCoreV1KVMTimer)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *KubevirtIoApiCoreV1Timer) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetPit

`func (o *KubevirtIoApiCoreV1Timer) GetPit() KubevirtIoApiCoreV1PITTimer`

GetPit returns the Pit field if non-nil, zero value otherwise.

### GetPitOk

`func (o *KubevirtIoApiCoreV1Timer) GetPitOk() (*KubevirtIoApiCoreV1PITTimer, bool)`

GetPitOk returns a tuple with the Pit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPit

`func (o *KubevirtIoApiCoreV1Timer) SetPit(v KubevirtIoApiCoreV1PITTimer)`

SetPit sets Pit field to given value.

### HasPit

`func (o *KubevirtIoApiCoreV1Timer) HasPit() bool`

HasPit returns a boolean if a field has been set.

### GetRtc

`func (o *KubevirtIoApiCoreV1Timer) GetRtc() KubevirtIoApiCoreV1RTCTimer`

GetRtc returns the Rtc field if non-nil, zero value otherwise.

### GetRtcOk

`func (o *KubevirtIoApiCoreV1Timer) GetRtcOk() (*KubevirtIoApiCoreV1RTCTimer, bool)`

GetRtcOk returns a tuple with the Rtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtc

`func (o *KubevirtIoApiCoreV1Timer) SetRtc(v KubevirtIoApiCoreV1RTCTimer)`

SetRtc sets Rtc field to given value.

### HasRtc

`func (o *KubevirtIoApiCoreV1Timer) HasRtc() bool`

HasRtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


