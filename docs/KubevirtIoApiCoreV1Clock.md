# KubevirtIoApiCoreV1Clock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timer** | Pointer to [**KubevirtIoApiCoreV1Timer**](KubevirtIoApiCoreV1Timer.md) | Timer specifies whih timers are attached to the vmi. | [optional] 
**Timezone** | Pointer to **string** | Timezone sets the guest clock to the specified timezone. Zone name follows the TZ environment variable format (e.g. &#39;America/New_York&#39;). | [optional] 
**Utc** | Pointer to [**KubevirtIoApiCoreV1ClockOffsetUTC**](KubevirtIoApiCoreV1ClockOffsetUTC.md) | UTC sets the guest clock to UTC on each boot. If an offset is specified, guest changes to the clock will be kept during reboots and are not reset. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Clock

`func NewKubevirtIoApiCoreV1Clock() *KubevirtIoApiCoreV1Clock`

NewKubevirtIoApiCoreV1Clock instantiates a new KubevirtIoApiCoreV1Clock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ClockWithDefaults

`func NewKubevirtIoApiCoreV1ClockWithDefaults() *KubevirtIoApiCoreV1Clock`

NewKubevirtIoApiCoreV1ClockWithDefaults instantiates a new KubevirtIoApiCoreV1Clock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimer

`func (o *KubevirtIoApiCoreV1Clock) GetTimer() KubevirtIoApiCoreV1Timer`

GetTimer returns the Timer field if non-nil, zero value otherwise.

### GetTimerOk

`func (o *KubevirtIoApiCoreV1Clock) GetTimerOk() (*KubevirtIoApiCoreV1Timer, bool)`

GetTimerOk returns a tuple with the Timer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimer

`func (o *KubevirtIoApiCoreV1Clock) SetTimer(v KubevirtIoApiCoreV1Timer)`

SetTimer sets Timer field to given value.

### HasTimer

`func (o *KubevirtIoApiCoreV1Clock) HasTimer() bool`

HasTimer returns a boolean if a field has been set.

### GetTimezone

`func (o *KubevirtIoApiCoreV1Clock) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *KubevirtIoApiCoreV1Clock) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *KubevirtIoApiCoreV1Clock) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *KubevirtIoApiCoreV1Clock) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUtc

`func (o *KubevirtIoApiCoreV1Clock) GetUtc() KubevirtIoApiCoreV1ClockOffsetUTC`

GetUtc returns the Utc field if non-nil, zero value otherwise.

### GetUtcOk

`func (o *KubevirtIoApiCoreV1Clock) GetUtcOk() (*KubevirtIoApiCoreV1ClockOffsetUTC, bool)`

GetUtcOk returns a tuple with the Utc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtc

`func (o *KubevirtIoApiCoreV1Clock) SetUtc(v KubevirtIoApiCoreV1ClockOffsetUTC)`

SetUtc sets Utc field to given value.

### HasUtc

`func (o *KubevirtIoApiCoreV1Clock) HasUtc() bool`

HasUtc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


