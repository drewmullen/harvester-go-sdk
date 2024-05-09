# KubevirtIoApiCoreV1ClockOffsetUTC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetSeconds** | Pointer to **int32** | OffsetSeconds specifies an offset in seconds, relative to UTC. If set, guest changes to the clock will be kept during reboots and not reset. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1ClockOffsetUTC

`func NewKubevirtIoApiCoreV1ClockOffsetUTC() *KubevirtIoApiCoreV1ClockOffsetUTC`

NewKubevirtIoApiCoreV1ClockOffsetUTC instantiates a new KubevirtIoApiCoreV1ClockOffsetUTC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ClockOffsetUTCWithDefaults

`func NewKubevirtIoApiCoreV1ClockOffsetUTCWithDefaults() *KubevirtIoApiCoreV1ClockOffsetUTC`

NewKubevirtIoApiCoreV1ClockOffsetUTCWithDefaults instantiates a new KubevirtIoApiCoreV1ClockOffsetUTC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffsetSeconds

`func (o *KubevirtIoApiCoreV1ClockOffsetUTC) GetOffsetSeconds() int32`

GetOffsetSeconds returns the OffsetSeconds field if non-nil, zero value otherwise.

### GetOffsetSecondsOk

`func (o *KubevirtIoApiCoreV1ClockOffsetUTC) GetOffsetSecondsOk() (*int32, bool)`

GetOffsetSecondsOk returns a tuple with the OffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetSeconds

`func (o *KubevirtIoApiCoreV1ClockOffsetUTC) SetOffsetSeconds(v int32)`

SetOffsetSeconds sets OffsetSeconds field to given value.

### HasOffsetSeconds

`func (o *KubevirtIoApiCoreV1ClockOffsetUTC) HasOffsetSeconds() bool`

HasOffsetSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


