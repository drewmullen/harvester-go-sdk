# KubevirtIoApiCoreV1Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Port** | **int32** |  | [default to 0]
**Protocol** | Pointer to **string** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Port

`func NewKubevirtIoApiCoreV1Port(port int32, ) *KubevirtIoApiCoreV1Port`

NewKubevirtIoApiCoreV1Port instantiates a new KubevirtIoApiCoreV1Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1PortWithDefaults

`func NewKubevirtIoApiCoreV1PortWithDefaults() *KubevirtIoApiCoreV1Port`

NewKubevirtIoApiCoreV1PortWithDefaults instantiates a new KubevirtIoApiCoreV1Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubevirtIoApiCoreV1Port) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Port) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Port) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubevirtIoApiCoreV1Port) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *KubevirtIoApiCoreV1Port) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KubevirtIoApiCoreV1Port) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KubevirtIoApiCoreV1Port) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *KubevirtIoApiCoreV1Port) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KubevirtIoApiCoreV1Port) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KubevirtIoApiCoreV1Port) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KubevirtIoApiCoreV1Port) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


