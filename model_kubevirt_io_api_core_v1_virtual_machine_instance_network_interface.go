/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface{}

// KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface struct for KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface
type KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface struct {
	InfoSource *string `json:"infoSource,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	Mac *string `json:"mac,omitempty"`
	Name *string `json:"name,omitempty"`
	QueueCount *int32 `json:"queueCount,omitempty"`
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface() *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface{}
	return &this
}

// NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterfaceWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterfaceWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface {
	this := KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface{}
	return &this
}

// GetInfoSource returns the InfoSource field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInfoSource() string {
	if o == nil || IsNil(o.InfoSource) {
		var ret string
		return ret
	}
	return *o.InfoSource
}

// GetInfoSourceOk returns a tuple with the InfoSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInfoSourceOk() (*string, bool) {
	if o == nil || IsNil(o.InfoSource) {
		return nil, false
	}
	return o.InfoSource, true
}

// HasInfoSource returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasInfoSource() bool {
	if o != nil && !IsNil(o.InfoSource) {
		return true
	}

	return false
}

// SetInfoSource gets a reference to the given string and assigns it to the InfoSource field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetInfoSource(v string) {
	o.InfoSource = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInterfaceName() string {
	if o == nil || IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasInterfaceName() bool {
	if o != nil && !IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddresses() []string {
	if o == nil || IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasIpAddresses() bool {
	if o != nil && !IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetName(v string) {
	o.Name = &v
}

// GetQueueCount returns the QueueCount field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetQueueCount() int32 {
	if o == nil || IsNil(o.QueueCount) {
		var ret int32
		return ret
	}
	return *o.QueueCount
}

// GetQueueCountOk returns a tuple with the QueueCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) GetQueueCountOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueCount) {
		return nil, false
	}
	return o.QueueCount, true
}

// HasQueueCount returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) HasQueueCount() bool {
	if o != nil && !IsNil(o.QueueCount) {
		return true
	}

	return false
}

// SetQueueCount gets a reference to the given int32 and assigns it to the QueueCount field.
func (o *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) SetQueueCount(v int32) {
	o.QueueCount = &v
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InfoSource) {
		toSerialize["infoSource"] = o.InfoSource
	}
	if !IsNil(o.InterfaceName) {
		toSerialize["interfaceName"] = o.InterfaceName
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.QueueCount) {
		toSerialize["queueCount"] = o.QueueCount
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface struct {
	value *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) Get() *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) Set(val *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface(val *KubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) *NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface {
	return &NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1VirtualMachineInstanceNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


