# KubevirtIoApiCoreV1DHCPOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootFileName** | Pointer to **string** | If specified will pass option 67 to interface&#39;s DHCP server | [optional] 
**NtpServers** | Pointer to **[]string** | If specified will pass the configured NTP server to the VM via DHCP option 042. | [optional] 
**PrivateOptions** | Pointer to [**[]KubevirtIoApiCoreV1DHCPPrivateOptions**](KubevirtIoApiCoreV1DHCPPrivateOptions.md) | If specified will pass extra DHCP options for private use, range: 224-254 | [optional] 
**TftpServerName** | Pointer to **string** | If specified will pass option 66 to interface&#39;s DHCP server | [optional] 

## Methods

### NewKubevirtIoApiCoreV1DHCPOptions

`func NewKubevirtIoApiCoreV1DHCPOptions() *KubevirtIoApiCoreV1DHCPOptions`

NewKubevirtIoApiCoreV1DHCPOptions instantiates a new KubevirtIoApiCoreV1DHCPOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DHCPOptionsWithDefaults

`func NewKubevirtIoApiCoreV1DHCPOptionsWithDefaults() *KubevirtIoApiCoreV1DHCPOptions`

NewKubevirtIoApiCoreV1DHCPOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1DHCPOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootFileName

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetBootFileName() string`

GetBootFileName returns the BootFileName field if non-nil, zero value otherwise.

### GetBootFileNameOk

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetBootFileNameOk() (*string, bool)`

GetBootFileNameOk returns a tuple with the BootFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootFileName

`func (o *KubevirtIoApiCoreV1DHCPOptions) SetBootFileName(v string)`

SetBootFileName sets BootFileName field to given value.

### HasBootFileName

`func (o *KubevirtIoApiCoreV1DHCPOptions) HasBootFileName() bool`

HasBootFileName returns a boolean if a field has been set.

### GetNtpServers

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *KubevirtIoApiCoreV1DHCPOptions) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *KubevirtIoApiCoreV1DHCPOptions) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetPrivateOptions

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetPrivateOptions() []KubevirtIoApiCoreV1DHCPPrivateOptions`

GetPrivateOptions returns the PrivateOptions field if non-nil, zero value otherwise.

### GetPrivateOptionsOk

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetPrivateOptionsOk() (*[]KubevirtIoApiCoreV1DHCPPrivateOptions, bool)`

GetPrivateOptionsOk returns a tuple with the PrivateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateOptions

`func (o *KubevirtIoApiCoreV1DHCPOptions) SetPrivateOptions(v []KubevirtIoApiCoreV1DHCPPrivateOptions)`

SetPrivateOptions sets PrivateOptions field to given value.

### HasPrivateOptions

`func (o *KubevirtIoApiCoreV1DHCPOptions) HasPrivateOptions() bool`

HasPrivateOptions returns a boolean if a field has been set.

### GetTftpServerName

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetTftpServerName() string`

GetTftpServerName returns the TftpServerName field if non-nil, zero value otherwise.

### GetTftpServerNameOk

`func (o *KubevirtIoApiCoreV1DHCPOptions) GetTftpServerNameOk() (*string, bool)`

GetTftpServerNameOk returns a tuple with the TftpServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpServerName

`func (o *KubevirtIoApiCoreV1DHCPOptions) SetTftpServerName(v string)`

SetTftpServerName sets TftpServerName field to given value.

### HasTftpServerName

`func (o *KubevirtIoApiCoreV1DHCPOptions) HasTftpServerName() bool`

HasTftpServerName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


