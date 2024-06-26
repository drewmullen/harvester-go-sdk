/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1DHCPOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1DHCPOptions{}

// KubevirtIoApiCoreV1DHCPOptions Extra DHCP options to use in the interface.
type KubevirtIoApiCoreV1DHCPOptions struct {
	// If specified will pass option 67 to interface's DHCP server
	BootFileName *string `json:"bootFileName,omitempty"`
	// If specified will pass the configured NTP server to the VM via DHCP option 042.
	NtpServers []string `json:"ntpServers,omitempty"`
	// If specified will pass extra DHCP options for private use, range: 224-254
	PrivateOptions []KubevirtIoApiCoreV1DHCPPrivateOptions `json:"privateOptions,omitempty"`
	// If specified will pass option 66 to interface's DHCP server
	TftpServerName *string `json:"tftpServerName,omitempty"`
}

// NewKubevirtIoApiCoreV1DHCPOptions instantiates a new KubevirtIoApiCoreV1DHCPOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1DHCPOptions() *KubevirtIoApiCoreV1DHCPOptions {
	this := KubevirtIoApiCoreV1DHCPOptions{}
	return &this
}

// NewKubevirtIoApiCoreV1DHCPOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1DHCPOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1DHCPOptionsWithDefaults() *KubevirtIoApiCoreV1DHCPOptions {
	this := KubevirtIoApiCoreV1DHCPOptions{}
	return &this
}

// GetBootFileName returns the BootFileName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetBootFileName() string {
	if o == nil || IsNil(o.BootFileName) {
		var ret string
		return ret
	}
	return *o.BootFileName
}

// GetBootFileNameOk returns a tuple with the BootFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetBootFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.BootFileName) {
		return nil, false
	}
	return o.BootFileName, true
}

// HasBootFileName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) HasBootFileName() bool {
	if o != nil && !IsNil(o.BootFileName) {
		return true
	}

	return false
}

// SetBootFileName gets a reference to the given string and assigns it to the BootFileName field.
func (o *KubevirtIoApiCoreV1DHCPOptions) SetBootFileName(v string) {
	o.BootFileName = &v
}

// GetNtpServers returns the NtpServers field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetNtpServers() []string {
	if o == nil || IsNil(o.NtpServers) {
		var ret []string
		return ret
	}
	return o.NtpServers
}

// GetNtpServersOk returns a tuple with the NtpServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetNtpServersOk() ([]string, bool) {
	if o == nil || IsNil(o.NtpServers) {
		return nil, false
	}
	return o.NtpServers, true
}

// HasNtpServers returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) HasNtpServers() bool {
	if o != nil && !IsNil(o.NtpServers) {
		return true
	}

	return false
}

// SetNtpServers gets a reference to the given []string and assigns it to the NtpServers field.
func (o *KubevirtIoApiCoreV1DHCPOptions) SetNtpServers(v []string) {
	o.NtpServers = v
}

// GetPrivateOptions returns the PrivateOptions field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetPrivateOptions() []KubevirtIoApiCoreV1DHCPPrivateOptions {
	if o == nil || IsNil(o.PrivateOptions) {
		var ret []KubevirtIoApiCoreV1DHCPPrivateOptions
		return ret
	}
	return o.PrivateOptions
}

// GetPrivateOptionsOk returns a tuple with the PrivateOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetPrivateOptionsOk() ([]KubevirtIoApiCoreV1DHCPPrivateOptions, bool) {
	if o == nil || IsNil(o.PrivateOptions) {
		return nil, false
	}
	return o.PrivateOptions, true
}

// HasPrivateOptions returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) HasPrivateOptions() bool {
	if o != nil && !IsNil(o.PrivateOptions) {
		return true
	}

	return false
}

// SetPrivateOptions gets a reference to the given []KubevirtIoApiCoreV1DHCPPrivateOptions and assigns it to the PrivateOptions field.
func (o *KubevirtIoApiCoreV1DHCPOptions) SetPrivateOptions(v []KubevirtIoApiCoreV1DHCPPrivateOptions) {
	o.PrivateOptions = v
}

// GetTftpServerName returns the TftpServerName field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetTftpServerName() string {
	if o == nil || IsNil(o.TftpServerName) {
		var ret string
		return ret
	}
	return *o.TftpServerName
}

// GetTftpServerNameOk returns a tuple with the TftpServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) GetTftpServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.TftpServerName) {
		return nil, false
	}
	return o.TftpServerName, true
}

// HasTftpServerName returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1DHCPOptions) HasTftpServerName() bool {
	if o != nil && !IsNil(o.TftpServerName) {
		return true
	}

	return false
}

// SetTftpServerName gets a reference to the given string and assigns it to the TftpServerName field.
func (o *KubevirtIoApiCoreV1DHCPOptions) SetTftpServerName(v string) {
	o.TftpServerName = &v
}

func (o KubevirtIoApiCoreV1DHCPOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1DHCPOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BootFileName) {
		toSerialize["bootFileName"] = o.BootFileName
	}
	if !IsNil(o.NtpServers) {
		toSerialize["ntpServers"] = o.NtpServers
	}
	if !IsNil(o.PrivateOptions) {
		toSerialize["privateOptions"] = o.PrivateOptions
	}
	if !IsNil(o.TftpServerName) {
		toSerialize["tftpServerName"] = o.TftpServerName
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1DHCPOptions struct {
	value *KubevirtIoApiCoreV1DHCPOptions
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1DHCPOptions) Get() *KubevirtIoApiCoreV1DHCPOptions {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1DHCPOptions) Set(val *KubevirtIoApiCoreV1DHCPOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1DHCPOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1DHCPOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1DHCPOptions(val *KubevirtIoApiCoreV1DHCPOptions) *NullableKubevirtIoApiCoreV1DHCPOptions {
	return &NullableKubevirtIoApiCoreV1DHCPOptions{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1DHCPOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1DHCPOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


