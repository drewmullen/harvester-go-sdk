/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoApiCoreV1Interface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1Interface{}

// KubevirtIoApiCoreV1Interface struct for KubevirtIoApiCoreV1Interface
type KubevirtIoApiCoreV1Interface struct {
	// If specified, the ACPI index is used to provide network interface device naming, that is stable across changes in PCI addresses assigned to the device. This value is required to be unique across all devices and be between 1 and (16*1024-1).
	AcpiIndex *int32 `json:"acpiIndex,omitempty"`
	// Binding specifies the binding plugin that will be used to connect the interface to the guest. It provides an alternative to InterfaceBindingMethod. version: 1alphav1
	Binding *KubevirtIoApiCoreV1PluginBinding `json:"binding,omitempty"`
	// BootOrder is an integer value > 0, used to determine ordering of boot devices. Lower values take precedence. Each interface or disk that has a boot order must have a unique value. Interfaces without a boot order are not tried.
	BootOrder *int32 `json:"bootOrder,omitempty"`
	// InterfaceBridge connects to a given network via a linux bridge.
	Bridge map[string]interface{} `json:"bridge,omitempty"`
	// If specified the network interface will pass additional DHCP options to the VMI
	DhcpOptions *KubevirtIoApiCoreV1DHCPOptions `json:"dhcpOptions,omitempty"`
	// Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.
	MacAddress *string `json:"macAddress,omitempty"`
	// InterfaceMacvtap connects to a given network by extending the Kubernetes node's L2 networks via a macvtap interface.
	Macvtap map[string]interface{} `json:"macvtap,omitempty"`
	// InterfaceMasquerade connects to a given network using netfilter rules to nat the traffic.
	Masquerade map[string]interface{} `json:"masquerade,omitempty"`
	// Interface model. One of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio. Defaults to virtio.
	Model *string `json:"model,omitempty"`
	// Logical name of the interface as well as a reference to the associated networks. Must match the Name of a Network.
	Name string `json:"name"`
	// InterfacePasst connects to a given network.
	Passt map[string]interface{} `json:"passt,omitempty"`
	// If specified, the virtual network interface will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10
	PciAddress *string `json:"pciAddress,omitempty"`
	// List of ports to be forwarded to the virtual machine.
	Ports []KubevirtIoApiCoreV1Port `json:"ports,omitempty"`
	// InterfaceSlirp connects to a given network using QEMU user networking mode.
	Slirp map[string]interface{} `json:"slirp,omitempty"`
	// InterfaceSRIOV connects to a given network by passing-through an SR-IOV PCI device via vfio.
	Sriov map[string]interface{} `json:"sriov,omitempty"`
	// State represents the requested operational state of the interface. The (only) value supported is `absent`, expressing a request to remove the interface.
	State *string `json:"state,omitempty"`
	// If specified, the virtual network interface address and its tag will be provided to the guest via config drive
	Tag *string `json:"tag,omitempty"`
}

type _KubevirtIoApiCoreV1Interface KubevirtIoApiCoreV1Interface

// NewKubevirtIoApiCoreV1Interface instantiates a new KubevirtIoApiCoreV1Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1Interface(name string) *KubevirtIoApiCoreV1Interface {
	this := KubevirtIoApiCoreV1Interface{}
	this.Name = name
	return &this
}

// NewKubevirtIoApiCoreV1InterfaceWithDefaults instantiates a new KubevirtIoApiCoreV1Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1InterfaceWithDefaults() *KubevirtIoApiCoreV1Interface {
	this := KubevirtIoApiCoreV1Interface{}
	var name string = ""
	this.Name = name
	return &this
}

// GetAcpiIndex returns the AcpiIndex field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetAcpiIndex() int32 {
	if o == nil || IsNil(o.AcpiIndex) {
		var ret int32
		return ret
	}
	return *o.AcpiIndex
}

// GetAcpiIndexOk returns a tuple with the AcpiIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetAcpiIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.AcpiIndex) {
		return nil, false
	}
	return o.AcpiIndex, true
}

// HasAcpiIndex returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasAcpiIndex() bool {
	if o != nil && !IsNil(o.AcpiIndex) {
		return true
	}

	return false
}

// SetAcpiIndex gets a reference to the given int32 and assigns it to the AcpiIndex field.
func (o *KubevirtIoApiCoreV1Interface) SetAcpiIndex(v int32) {
	o.AcpiIndex = &v
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetBinding() KubevirtIoApiCoreV1PluginBinding {
	if o == nil || IsNil(o.Binding) {
		var ret KubevirtIoApiCoreV1PluginBinding
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetBindingOk() (*KubevirtIoApiCoreV1PluginBinding, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given KubevirtIoApiCoreV1PluginBinding and assigns it to the Binding field.
func (o *KubevirtIoApiCoreV1Interface) SetBinding(v KubevirtIoApiCoreV1PluginBinding) {
	o.Binding = &v
}

// GetBootOrder returns the BootOrder field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetBootOrder() int32 {
	if o == nil || IsNil(o.BootOrder) {
		var ret int32
		return ret
	}
	return *o.BootOrder
}

// GetBootOrderOk returns a tuple with the BootOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetBootOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.BootOrder) {
		return nil, false
	}
	return o.BootOrder, true
}

// HasBootOrder returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasBootOrder() bool {
	if o != nil && !IsNil(o.BootOrder) {
		return true
	}

	return false
}

// SetBootOrder gets a reference to the given int32 and assigns it to the BootOrder field.
func (o *KubevirtIoApiCoreV1Interface) SetBootOrder(v int32) {
	o.BootOrder = &v
}

// GetBridge returns the Bridge field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetBridge() map[string]interface{} {
	if o == nil || IsNil(o.Bridge) {
		var ret map[string]interface{}
		return ret
	}
	return o.Bridge
}

// GetBridgeOk returns a tuple with the Bridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetBridgeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Bridge) {
		return map[string]interface{}{}, false
	}
	return o.Bridge, true
}

// HasBridge returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasBridge() bool {
	if o != nil && !IsNil(o.Bridge) {
		return true
	}

	return false
}

// SetBridge gets a reference to the given map[string]interface{} and assigns it to the Bridge field.
func (o *KubevirtIoApiCoreV1Interface) SetBridge(v map[string]interface{}) {
	o.Bridge = v
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetDhcpOptions() KubevirtIoApiCoreV1DHCPOptions {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret KubevirtIoApiCoreV1DHCPOptions
		return ret
	}
	return *o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetDhcpOptionsOk() (*KubevirtIoApiCoreV1DHCPOptions, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given KubevirtIoApiCoreV1DHCPOptions and assigns it to the DhcpOptions field.
func (o *KubevirtIoApiCoreV1Interface) SetDhcpOptions(v KubevirtIoApiCoreV1DHCPOptions) {
	o.DhcpOptions = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetMacAddress() string {
	if o == nil || IsNil(o.MacAddress) {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetMacAddressOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddress) {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasMacAddress() bool {
	if o != nil && !IsNil(o.MacAddress) {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *KubevirtIoApiCoreV1Interface) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetMacvtap returns the Macvtap field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetMacvtap() map[string]interface{} {
	if o == nil || IsNil(o.Macvtap) {
		var ret map[string]interface{}
		return ret
	}
	return o.Macvtap
}

// GetMacvtapOk returns a tuple with the Macvtap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetMacvtapOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Macvtap) {
		return map[string]interface{}{}, false
	}
	return o.Macvtap, true
}

// HasMacvtap returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasMacvtap() bool {
	if o != nil && !IsNil(o.Macvtap) {
		return true
	}

	return false
}

// SetMacvtap gets a reference to the given map[string]interface{} and assigns it to the Macvtap field.
func (o *KubevirtIoApiCoreV1Interface) SetMacvtap(v map[string]interface{}) {
	o.Macvtap = v
}

// GetMasquerade returns the Masquerade field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetMasquerade() map[string]interface{} {
	if o == nil || IsNil(o.Masquerade) {
		var ret map[string]interface{}
		return ret
	}
	return o.Masquerade
}

// GetMasqueradeOk returns a tuple with the Masquerade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetMasqueradeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Masquerade) {
		return map[string]interface{}{}, false
	}
	return o.Masquerade, true
}

// HasMasquerade returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasMasquerade() bool {
	if o != nil && !IsNil(o.Masquerade) {
		return true
	}

	return false
}

// SetMasquerade gets a reference to the given map[string]interface{} and assigns it to the Masquerade field.
func (o *KubevirtIoApiCoreV1Interface) SetMasquerade(v map[string]interface{}) {
	o.Masquerade = v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *KubevirtIoApiCoreV1Interface) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value
func (o *KubevirtIoApiCoreV1Interface) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoApiCoreV1Interface) SetName(v string) {
	o.Name = v
}

// GetPasst returns the Passt field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetPasst() map[string]interface{} {
	if o == nil || IsNil(o.Passt) {
		var ret map[string]interface{}
		return ret
	}
	return o.Passt
}

// GetPasstOk returns a tuple with the Passt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetPasstOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Passt) {
		return map[string]interface{}{}, false
	}
	return o.Passt, true
}

// HasPasst returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasPasst() bool {
	if o != nil && !IsNil(o.Passt) {
		return true
	}

	return false
}

// SetPasst gets a reference to the given map[string]interface{} and assigns it to the Passt field.
func (o *KubevirtIoApiCoreV1Interface) SetPasst(v map[string]interface{}) {
	o.Passt = v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetPciAddress() string {
	if o == nil || IsNil(o.PciAddress) {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetPciAddressOk() (*string, bool) {
	if o == nil || IsNil(o.PciAddress) {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasPciAddress() bool {
	if o != nil && !IsNil(o.PciAddress) {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *KubevirtIoApiCoreV1Interface) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetPorts() []KubevirtIoApiCoreV1Port {
	if o == nil || IsNil(o.Ports) {
		var ret []KubevirtIoApiCoreV1Port
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetPortsOk() ([]KubevirtIoApiCoreV1Port, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []KubevirtIoApiCoreV1Port and assigns it to the Ports field.
func (o *KubevirtIoApiCoreV1Interface) SetPorts(v []KubevirtIoApiCoreV1Port) {
	o.Ports = v
}

// GetSlirp returns the Slirp field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetSlirp() map[string]interface{} {
	if o == nil || IsNil(o.Slirp) {
		var ret map[string]interface{}
		return ret
	}
	return o.Slirp
}

// GetSlirpOk returns a tuple with the Slirp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetSlirpOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Slirp) {
		return map[string]interface{}{}, false
	}
	return o.Slirp, true
}

// HasSlirp returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasSlirp() bool {
	if o != nil && !IsNil(o.Slirp) {
		return true
	}

	return false
}

// SetSlirp gets a reference to the given map[string]interface{} and assigns it to the Slirp field.
func (o *KubevirtIoApiCoreV1Interface) SetSlirp(v map[string]interface{}) {
	o.Slirp = v
}

// GetSriov returns the Sriov field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetSriov() map[string]interface{} {
	if o == nil || IsNil(o.Sriov) {
		var ret map[string]interface{}
		return ret
	}
	return o.Sriov
}

// GetSriovOk returns a tuple with the Sriov field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetSriovOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Sriov) {
		return map[string]interface{}{}, false
	}
	return o.Sriov, true
}

// HasSriov returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasSriov() bool {
	if o != nil && !IsNil(o.Sriov) {
		return true
	}

	return false
}

// SetSriov gets a reference to the given map[string]interface{} and assigns it to the Sriov field.
func (o *KubevirtIoApiCoreV1Interface) SetSriov(v map[string]interface{}) {
	o.Sriov = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *KubevirtIoApiCoreV1Interface) SetState(v string) {
	o.State = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Interface) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Interface) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Interface) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *KubevirtIoApiCoreV1Interface) SetTag(v string) {
	o.Tag = &v
}

func (o KubevirtIoApiCoreV1Interface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1Interface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcpiIndex) {
		toSerialize["acpiIndex"] = o.AcpiIndex
	}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	if !IsNil(o.BootOrder) {
		toSerialize["bootOrder"] = o.BootOrder
	}
	if !IsNil(o.Bridge) {
		toSerialize["bridge"] = o.Bridge
	}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcpOptions"] = o.DhcpOptions
	}
	if !IsNil(o.MacAddress) {
		toSerialize["macAddress"] = o.MacAddress
	}
	if !IsNil(o.Macvtap) {
		toSerialize["macvtap"] = o.Macvtap
	}
	if !IsNil(o.Masquerade) {
		toSerialize["masquerade"] = o.Masquerade
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Passt) {
		toSerialize["passt"] = o.Passt
	}
	if !IsNil(o.PciAddress) {
		toSerialize["pciAddress"] = o.PciAddress
	}
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.Slirp) {
		toSerialize["slirp"] = o.Slirp
	}
	if !IsNil(o.Sriov) {
		toSerialize["sriov"] = o.Sriov
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1Interface) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubevirtIoApiCoreV1Interface := _KubevirtIoApiCoreV1Interface{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1Interface)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1Interface(varKubevirtIoApiCoreV1Interface)

	return err
}

type NullableKubevirtIoApiCoreV1Interface struct {
	value *KubevirtIoApiCoreV1Interface
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1Interface) Get() *KubevirtIoApiCoreV1Interface {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1Interface) Set(val *KubevirtIoApiCoreV1Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1Interface) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1Interface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1Interface(val *KubevirtIoApiCoreV1Interface) *NullableKubevirtIoApiCoreV1Interface {
	return &NullableKubevirtIoApiCoreV1Interface{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1Interface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1Interface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


