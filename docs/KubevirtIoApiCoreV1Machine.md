# KubevirtIoApiCoreV1Machine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | QEMU machine type is the actual chipset of the VirtualMachineInstance. | [optional] [default to ""]

## Methods

### NewKubevirtIoApiCoreV1Machine

`func NewKubevirtIoApiCoreV1Machine() *KubevirtIoApiCoreV1Machine`

NewKubevirtIoApiCoreV1Machine instantiates a new KubevirtIoApiCoreV1Machine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1MachineWithDefaults

`func NewKubevirtIoApiCoreV1MachineWithDefaults() *KubevirtIoApiCoreV1Machine`

NewKubevirtIoApiCoreV1MachineWithDefaults instantiates a new KubevirtIoApiCoreV1Machine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KubevirtIoApiCoreV1Machine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KubevirtIoApiCoreV1Machine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KubevirtIoApiCoreV1Machine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KubevirtIoApiCoreV1Machine) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


