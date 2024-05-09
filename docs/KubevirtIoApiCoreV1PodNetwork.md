# KubevirtIoApiCoreV1PodNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmIPv6NetworkCIDR** | Pointer to **string** | IPv6 CIDR for the vm network. Defaults to fd10:0:2::/120 if not specified. | [optional] 
**VmNetworkCIDR** | Pointer to **string** | CIDR for vm network. Default 10.0.2.0/24 if not specified. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1PodNetwork

`func NewKubevirtIoApiCoreV1PodNetwork() *KubevirtIoApiCoreV1PodNetwork`

NewKubevirtIoApiCoreV1PodNetwork instantiates a new KubevirtIoApiCoreV1PodNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1PodNetworkWithDefaults

`func NewKubevirtIoApiCoreV1PodNetworkWithDefaults() *KubevirtIoApiCoreV1PodNetwork`

NewKubevirtIoApiCoreV1PodNetworkWithDefaults instantiates a new KubevirtIoApiCoreV1PodNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmIPv6NetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) GetVmIPv6NetworkCIDR() string`

GetVmIPv6NetworkCIDR returns the VmIPv6NetworkCIDR field if non-nil, zero value otherwise.

### GetVmIPv6NetworkCIDROk

`func (o *KubevirtIoApiCoreV1PodNetwork) GetVmIPv6NetworkCIDROk() (*string, bool)`

GetVmIPv6NetworkCIDROk returns a tuple with the VmIPv6NetworkCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmIPv6NetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) SetVmIPv6NetworkCIDR(v string)`

SetVmIPv6NetworkCIDR sets VmIPv6NetworkCIDR field to given value.

### HasVmIPv6NetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) HasVmIPv6NetworkCIDR() bool`

HasVmIPv6NetworkCIDR returns a boolean if a field has been set.

### GetVmNetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) GetVmNetworkCIDR() string`

GetVmNetworkCIDR returns the VmNetworkCIDR field if non-nil, zero value otherwise.

### GetVmNetworkCIDROk

`func (o *KubevirtIoApiCoreV1PodNetwork) GetVmNetworkCIDROk() (*string, bool)`

GetVmNetworkCIDROk returns a tuple with the VmNetworkCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) SetVmNetworkCIDR(v string)`

SetVmNetworkCIDR sets VmNetworkCIDR field to given value.

### HasVmNetworkCIDR

`func (o *KubevirtIoApiCoreV1PodNetwork) HasVmNetworkCIDR() bool`

HasVmNetworkCIDR returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


