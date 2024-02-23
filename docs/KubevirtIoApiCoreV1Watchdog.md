# KubevirtIoApiCoreV1Watchdog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**I6300esb** | Pointer to [**KubevirtIoApiCoreV1I6300ESBWatchdog**](KubevirtIoApiCoreV1I6300ESBWatchdog.md) |  | [optional] 
**Name** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1Watchdog

`func NewKubevirtIoApiCoreV1Watchdog(name string, ) *KubevirtIoApiCoreV1Watchdog`

NewKubevirtIoApiCoreV1Watchdog instantiates a new KubevirtIoApiCoreV1Watchdog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1WatchdogWithDefaults

`func NewKubevirtIoApiCoreV1WatchdogWithDefaults() *KubevirtIoApiCoreV1Watchdog`

NewKubevirtIoApiCoreV1WatchdogWithDefaults instantiates a new KubevirtIoApiCoreV1Watchdog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetI6300esb

`func (o *KubevirtIoApiCoreV1Watchdog) GetI6300esb() KubevirtIoApiCoreV1I6300ESBWatchdog`

GetI6300esb returns the I6300esb field if non-nil, zero value otherwise.

### GetI6300esbOk

`func (o *KubevirtIoApiCoreV1Watchdog) GetI6300esbOk() (*KubevirtIoApiCoreV1I6300ESBWatchdog, bool)`

GetI6300esbOk returns a tuple with the I6300esb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI6300esb

`func (o *KubevirtIoApiCoreV1Watchdog) SetI6300esb(v KubevirtIoApiCoreV1I6300ESBWatchdog)`

SetI6300esb sets I6300esb field to given value.

### HasI6300esb

`func (o *KubevirtIoApiCoreV1Watchdog) HasI6300esb() bool`

HasI6300esb returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Watchdog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Watchdog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Watchdog) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


