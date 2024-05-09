# KubevirtIoApiCoreV1Input

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | Pointer to **string** | Bus indicates the bus of input device to emulate. Supported values: virtio, usb. | [optional] 
**Name** | **string** | Name is the device name | [default to ""]
**Type** | **string** | Type indicated the type of input device. Supported values: tablet. | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1Input

`func NewKubevirtIoApiCoreV1Input(name string, type_ string, ) *KubevirtIoApiCoreV1Input`

NewKubevirtIoApiCoreV1Input instantiates a new KubevirtIoApiCoreV1Input object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1InputWithDefaults

`func NewKubevirtIoApiCoreV1InputWithDefaults() *KubevirtIoApiCoreV1Input`

NewKubevirtIoApiCoreV1InputWithDefaults instantiates a new KubevirtIoApiCoreV1Input object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *KubevirtIoApiCoreV1Input) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *KubevirtIoApiCoreV1Input) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *KubevirtIoApiCoreV1Input) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *KubevirtIoApiCoreV1Input) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Input) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Input) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Input) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *KubevirtIoApiCoreV1Input) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubevirtIoApiCoreV1Input) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubevirtIoApiCoreV1Input) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


