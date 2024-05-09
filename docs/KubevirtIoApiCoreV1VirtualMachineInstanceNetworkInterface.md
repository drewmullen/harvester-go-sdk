# KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InfoSource** | Pointer to **string** | Specifies the origin of the interface data collected. values: domain, guest-agent, multus-status. | [optional] 
**InterfaceName** | Pointer to **string** | The interface name inside the Virtual Machine | [optional] 
**IpAddress** | Pointer to **string** | IP address of a Virtual Machine interface. It is always the first item of IPs | [optional] 
**IpAddresses** | Pointer to **[]string** | List of all IP addresses of a Virtual Machine interface | [optional] 
**Mac** | Pointer to **string** | Hardware address of a Virtual Machine interface | [optional] 
**Name** | Pointer to **string** | Name of the interface, corresponds to name of the network assigned to the interface | [optional] 
**QueueCount** | Pointer to **int32** | Specifies how many queues are allocated by MultiQueue | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface() *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface`

NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterfaceWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterfaceWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface`

NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterfaceWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfoSource

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInfoSource() string`

GetInfoSource returns the InfoSource field if non-nil, zero value otherwise.

### GetInfoSourceOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInfoSourceOk() (*string, bool)`

GetInfoSourceOk returns a tuple with the InfoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoSource

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetInfoSource(v string)`

SetInfoSource sets InfoSource field to given value.

### HasInfoSource

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasInfoSource() bool`

HasInfoSource returns a boolean if a field has been set.

### GetInterfaceName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpAddresses

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetMac

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueueCount

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetQueueCount() int32`

GetQueueCount returns the QueueCount field if non-nil, zero value otherwise.

### GetQueueCountOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetQueueCountOk() (*int32, bool)`

GetQueueCountOk returns a tuple with the QueueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueCount

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetQueueCount(v int32)`

SetQueueCount sets QueueCount field to given value.

### HasQueueCount

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasQueueCount() bool`

HasQueueCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


