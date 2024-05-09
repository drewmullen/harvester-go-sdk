# KubevirtIoApiCoreV1Realtime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mask** | Pointer to **string** | Mask defines the vcpu mask expression that defines which vcpus are used for realtime. Format matches libvirt&#39;s expressions. Example: \&quot;0-3,^1\&quot;,\&quot;0,2,3\&quot;,\&quot;2-3\&quot; | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Realtime

`func NewKubevirtIoApiCoreV1Realtime() *KubevirtIoApiCoreV1Realtime`

NewKubevirtIoApiCoreV1Realtime instantiates a new KubevirtIoApiCoreV1Realtime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1RealtimeWithDefaults

`func NewKubevirtIoApiCoreV1RealtimeWithDefaults() *KubevirtIoApiCoreV1Realtime`

NewKubevirtIoApiCoreV1RealtimeWithDefaults instantiates a new KubevirtIoApiCoreV1Realtime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMask

`func (o *KubevirtIoApiCoreV1Realtime) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *KubevirtIoApiCoreV1Realtime) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *KubevirtIoApiCoreV1Realtime) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *KubevirtIoApiCoreV1Realtime) HasMask() bool`

HasMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


