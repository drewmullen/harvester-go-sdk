# KubevirtIoApiCoreV1DHCPPrivateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Option** | **int32** | Option is an Integer value from 224-254 Required. | [default to 0]
**Value** | **string** | Value is a String value for the Option provided Required. | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1DHCPPrivateOptions

`func NewKubevirtIoApiCoreV1DHCPPrivateOptions(option int32, value string, ) *KubevirtIoApiCoreV1DHCPPrivateOptions`

NewKubevirtIoApiCoreV1DHCPPrivateOptions instantiates a new KubevirtIoApiCoreV1DHCPPrivateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DHCPPrivateOptionsWithDefaults

`func NewKubevirtIoApiCoreV1DHCPPrivateOptionsWithDefaults() *KubevirtIoApiCoreV1DHCPPrivateOptions`

NewKubevirtIoApiCoreV1DHCPPrivateOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1DHCPPrivateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOption

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) GetOption() int32`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) GetOptionOk() (*int32, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) SetOption(v int32)`

SetOption sets Option field to given value.


### GetValue

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KubevirtIoApiCoreV1DHCPPrivateOptions) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


