# KubevirtIoApiCoreV1PITTimer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Present** | Pointer to **bool** | Enabled set to false makes sure that the machine type or a preset can&#39;t add the timer. Defaults to true. | [optional] 
**TickPolicy** | Pointer to **string** | TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest. One of \&quot;delay\&quot;, \&quot;catchup\&quot;, \&quot;discard\&quot;. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1PITTimer

`func NewKubevirtIoApiCoreV1PITTimer() *KubevirtIoApiCoreV1PITTimer`

NewKubevirtIoApiCoreV1PITTimer instantiates a new KubevirtIoApiCoreV1PITTimer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1PITTimerWithDefaults

`func NewKubevirtIoApiCoreV1PITTimerWithDefaults() *KubevirtIoApiCoreV1PITTimer`

NewKubevirtIoApiCoreV1PITTimerWithDefaults instantiates a new KubevirtIoApiCoreV1PITTimer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPresent

`func (o *KubevirtIoApiCoreV1PITTimer) GetPresent() bool`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *KubevirtIoApiCoreV1PITTimer) GetPresentOk() (*bool, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *KubevirtIoApiCoreV1PITTimer) SetPresent(v bool)`

SetPresent sets Present field to given value.

### HasPresent

`func (o *KubevirtIoApiCoreV1PITTimer) HasPresent() bool`

HasPresent returns a boolean if a field has been set.

### GetTickPolicy

`func (o *KubevirtIoApiCoreV1PITTimer) GetTickPolicy() string`

GetTickPolicy returns the TickPolicy field if non-nil, zero value otherwise.

### GetTickPolicyOk

`func (o *KubevirtIoApiCoreV1PITTimer) GetTickPolicyOk() (*string, bool)`

GetTickPolicyOk returns a tuple with the TickPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickPolicy

`func (o *KubevirtIoApiCoreV1PITTimer) SetTickPolicy(v string)`

SetTickPolicy sets TickPolicy field to given value.

### HasTickPolicy

`func (o *KubevirtIoApiCoreV1PITTimer) HasTickPolicy() bool`

HasTickPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


