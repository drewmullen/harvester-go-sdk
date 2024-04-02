# K8sIoV1TCPSocketAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Port** | **string** |  | [default to "{}"]

## Methods

### NewK8sIoV1TCPSocketAction

`func NewK8sIoV1TCPSocketAction(port string, ) *K8sIoV1TCPSocketAction`

NewK8sIoV1TCPSocketAction instantiates a new K8sIoV1TCPSocketAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1TCPSocketActionWithDefaults

`func NewK8sIoV1TCPSocketActionWithDefaults() *K8sIoV1TCPSocketAction`

NewK8sIoV1TCPSocketActionWithDefaults instantiates a new K8sIoV1TCPSocketAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *K8sIoV1TCPSocketAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *K8sIoV1TCPSocketAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *K8sIoV1TCPSocketAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *K8sIoV1TCPSocketAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *K8sIoV1TCPSocketAction) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *K8sIoV1TCPSocketAction) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *K8sIoV1TCPSocketAction) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


