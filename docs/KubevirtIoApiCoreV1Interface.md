# KubevirtIoApiCoreV1Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcpiIndex** | Pointer to **int32** | If specified, the ACPI index is used to provide network interface device naming, that is stable across changes in PCI addresses assigned to the device. This value is required to be unique across all devices and be between 1 and (16*1024-1). | [optional] 
**Binding** | Pointer to [**KubevirtIoApiCoreV1PluginBinding**](KubevirtIoApiCoreV1PluginBinding.md) | Binding specifies the binding plugin that will be used to connect the interface to the guest. It provides an alternative to InterfaceBindingMethod. version: 1alphav1 | [optional] 
**BootOrder** | Pointer to **int32** | BootOrder is an integer value &gt; 0, used to determine ordering of boot devices. Lower values take precedence. Each interface or disk that has a boot order must have a unique value. Interfaces without a boot order are not tried. | [optional] 
**Bridge** | Pointer to **map[string]interface{}** | InterfaceBridge connects to a given network via a linux bridge. | [optional] 
**DhcpOptions** | Pointer to [**KubevirtIoApiCoreV1DHCPOptions**](KubevirtIoApiCoreV1DHCPOptions.md) | If specified the network interface will pass additional DHCP options to the VMI | [optional] 
**MacAddress** | Pointer to **string** | Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF. | [optional] 
**Macvtap** | Pointer to **map[string]interface{}** | InterfaceMacvtap connects to a given network by extending the Kubernetes node&#39;s L2 networks via a macvtap interface. | [optional] 
**Masquerade** | Pointer to **map[string]interface{}** | InterfaceMasquerade connects to a given network using netfilter rules to nat the traffic. | [optional] 
**Model** | Pointer to **string** | Interface model. One of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio. Defaults to virtio. | [optional] 
**Name** | **string** | Logical name of the interface as well as a reference to the associated networks. Must match the Name of a Network. | [default to ""]
**Passt** | Pointer to **map[string]interface{}** | InterfacePasst connects to a given network. | [optional] 
**PciAddress** | Pointer to **string** | If specified, the virtual network interface will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10 | [optional] 
**Ports** | Pointer to [**[]KubevirtIoApiCoreV1Port**](KubevirtIoApiCoreV1Port.md) | List of ports to be forwarded to the virtual machine. | [optional] 
**Slirp** | Pointer to **map[string]interface{}** | InterfaceSlirp connects to a given network using QEMU user networking mode. | [optional] 
**Sriov** | Pointer to **map[string]interface{}** | InterfaceSRIOV connects to a given network by passing-through an SR-IOV PCI device via vfio. | [optional] 
**State** | Pointer to **string** | State represents the requested operational state of the interface. The (only) value supported is &#x60;absent&#x60;, expressing a request to remove the interface. | [optional] 
**Tag** | Pointer to **string** | If specified, the virtual network interface address and its tag will be provided to the guest via config drive | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Interface

`func NewKubevirtIoApiCoreV1Interface(name string, ) *KubevirtIoApiCoreV1Interface`

NewKubevirtIoApiCoreV1Interface instantiates a new KubevirtIoApiCoreV1Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1InterfaceWithDefaults

`func NewKubevirtIoApiCoreV1InterfaceWithDefaults() *KubevirtIoApiCoreV1Interface`

NewKubevirtIoApiCoreV1InterfaceWithDefaults instantiates a new KubevirtIoApiCoreV1Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpiIndex

`func (o *KubevirtIoApiCoreV1Interface) GetAcpiIndex() int32`

GetAcpiIndex returns the AcpiIndex field if non-nil, zero value otherwise.

### GetAcpiIndexOk

`func (o *KubevirtIoApiCoreV1Interface) GetAcpiIndexOk() (*int32, bool)`

GetAcpiIndexOk returns a tuple with the AcpiIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpiIndex

`func (o *KubevirtIoApiCoreV1Interface) SetAcpiIndex(v int32)`

SetAcpiIndex sets AcpiIndex field to given value.

### HasAcpiIndex

`func (o *KubevirtIoApiCoreV1Interface) HasAcpiIndex() bool`

HasAcpiIndex returns a boolean if a field has been set.

### GetBinding

`func (o *KubevirtIoApiCoreV1Interface) GetBinding() KubevirtIoApiCoreV1PluginBinding`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *KubevirtIoApiCoreV1Interface) GetBindingOk() (*KubevirtIoApiCoreV1PluginBinding, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *KubevirtIoApiCoreV1Interface) SetBinding(v KubevirtIoApiCoreV1PluginBinding)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *KubevirtIoApiCoreV1Interface) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetBootOrder

`func (o *KubevirtIoApiCoreV1Interface) GetBootOrder() int32`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *KubevirtIoApiCoreV1Interface) GetBootOrderOk() (*int32, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *KubevirtIoApiCoreV1Interface) SetBootOrder(v int32)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *KubevirtIoApiCoreV1Interface) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.

### GetBridge

`func (o *KubevirtIoApiCoreV1Interface) GetBridge() map[string]interface{}`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *KubevirtIoApiCoreV1Interface) GetBridgeOk() (*map[string]interface{}, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *KubevirtIoApiCoreV1Interface) SetBridge(v map[string]interface{})`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *KubevirtIoApiCoreV1Interface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *KubevirtIoApiCoreV1Interface) GetDhcpOptions() KubevirtIoApiCoreV1DHCPOptions`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *KubevirtIoApiCoreV1Interface) GetDhcpOptionsOk() (*KubevirtIoApiCoreV1DHCPOptions, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *KubevirtIoApiCoreV1Interface) SetDhcpOptions(v KubevirtIoApiCoreV1DHCPOptions)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *KubevirtIoApiCoreV1Interface) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetMacAddress

`func (o *KubevirtIoApiCoreV1Interface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *KubevirtIoApiCoreV1Interface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *KubevirtIoApiCoreV1Interface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *KubevirtIoApiCoreV1Interface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMacvtap

`func (o *KubevirtIoApiCoreV1Interface) GetMacvtap() map[string]interface{}`

GetMacvtap returns the Macvtap field if non-nil, zero value otherwise.

### GetMacvtapOk

`func (o *KubevirtIoApiCoreV1Interface) GetMacvtapOk() (*map[string]interface{}, bool)`

GetMacvtapOk returns a tuple with the Macvtap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacvtap

`func (o *KubevirtIoApiCoreV1Interface) SetMacvtap(v map[string]interface{})`

SetMacvtap sets Macvtap field to given value.

### HasMacvtap

`func (o *KubevirtIoApiCoreV1Interface) HasMacvtap() bool`

HasMacvtap returns a boolean if a field has been set.

### GetMasquerade

`func (o *KubevirtIoApiCoreV1Interface) GetMasquerade() map[string]interface{}`

GetMasquerade returns the Masquerade field if non-nil, zero value otherwise.

### GetMasqueradeOk

`func (o *KubevirtIoApiCoreV1Interface) GetMasqueradeOk() (*map[string]interface{}, bool)`

GetMasqueradeOk returns a tuple with the Masquerade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasquerade

`func (o *KubevirtIoApiCoreV1Interface) SetMasquerade(v map[string]interface{})`

SetMasquerade sets Masquerade field to given value.

### HasMasquerade

`func (o *KubevirtIoApiCoreV1Interface) HasMasquerade() bool`

HasMasquerade returns a boolean if a field has been set.

### GetModel

`func (o *KubevirtIoApiCoreV1Interface) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *KubevirtIoApiCoreV1Interface) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *KubevirtIoApiCoreV1Interface) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *KubevirtIoApiCoreV1Interface) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Interface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Interface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Interface) SetName(v string)`

SetName sets Name field to given value.


### GetPasst

`func (o *KubevirtIoApiCoreV1Interface) GetPasst() map[string]interface{}`

GetPasst returns the Passt field if non-nil, zero value otherwise.

### GetPasstOk

`func (o *KubevirtIoApiCoreV1Interface) GetPasstOk() (*map[string]interface{}, bool)`

GetPasstOk returns a tuple with the Passt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasst

`func (o *KubevirtIoApiCoreV1Interface) SetPasst(v map[string]interface{})`

SetPasst sets Passt field to given value.

### HasPasst

`func (o *KubevirtIoApiCoreV1Interface) HasPasst() bool`

HasPasst returns a boolean if a field has been set.

### GetPciAddress

`func (o *KubevirtIoApiCoreV1Interface) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *KubevirtIoApiCoreV1Interface) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *KubevirtIoApiCoreV1Interface) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *KubevirtIoApiCoreV1Interface) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetPorts

`func (o *KubevirtIoApiCoreV1Interface) GetPorts() []KubevirtIoApiCoreV1Port`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KubevirtIoApiCoreV1Interface) GetPortsOk() (*[]KubevirtIoApiCoreV1Port, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KubevirtIoApiCoreV1Interface) SetPorts(v []KubevirtIoApiCoreV1Port)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *KubevirtIoApiCoreV1Interface) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSlirp

`func (o *KubevirtIoApiCoreV1Interface) GetSlirp() map[string]interface{}`

GetSlirp returns the Slirp field if non-nil, zero value otherwise.

### GetSlirpOk

`func (o *KubevirtIoApiCoreV1Interface) GetSlirpOk() (*map[string]interface{}, bool)`

GetSlirpOk returns a tuple with the Slirp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlirp

`func (o *KubevirtIoApiCoreV1Interface) SetSlirp(v map[string]interface{})`

SetSlirp sets Slirp field to given value.

### HasSlirp

`func (o *KubevirtIoApiCoreV1Interface) HasSlirp() bool`

HasSlirp returns a boolean if a field has been set.

### GetSriov

`func (o *KubevirtIoApiCoreV1Interface) GetSriov() map[string]interface{}`

GetSriov returns the Sriov field if non-nil, zero value otherwise.

### GetSriovOk

`func (o *KubevirtIoApiCoreV1Interface) GetSriovOk() (*map[string]interface{}, bool)`

GetSriovOk returns a tuple with the Sriov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSriov

`func (o *KubevirtIoApiCoreV1Interface) SetSriov(v map[string]interface{})`

SetSriov sets Sriov field to given value.

### HasSriov

`func (o *KubevirtIoApiCoreV1Interface) HasSriov() bool`

HasSriov returns a boolean if a field has been set.

### GetState

`func (o *KubevirtIoApiCoreV1Interface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KubevirtIoApiCoreV1Interface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KubevirtIoApiCoreV1Interface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *KubevirtIoApiCoreV1Interface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTag

`func (o *KubevirtIoApiCoreV1Interface) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *KubevirtIoApiCoreV1Interface) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *KubevirtIoApiCoreV1Interface) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *KubevirtIoApiCoreV1Interface) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


