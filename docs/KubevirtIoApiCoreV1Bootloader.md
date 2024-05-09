# KubevirtIoApiCoreV1Bootloader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bios** | Pointer to [**KubevirtIoApiCoreV1BIOS**](KubevirtIoApiCoreV1BIOS.md) | If set (default), BIOS will be used. | [optional] 
**Efi** | Pointer to [**KubevirtIoApiCoreV1EFI**](KubevirtIoApiCoreV1EFI.md) | If set, EFI will be used instead of BIOS. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Bootloader

`func NewKubevirtIoApiCoreV1Bootloader() *KubevirtIoApiCoreV1Bootloader`

NewKubevirtIoApiCoreV1Bootloader instantiates a new KubevirtIoApiCoreV1Bootloader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1BootloaderWithDefaults

`func NewKubevirtIoApiCoreV1BootloaderWithDefaults() *KubevirtIoApiCoreV1Bootloader`

NewKubevirtIoApiCoreV1BootloaderWithDefaults instantiates a new KubevirtIoApiCoreV1Bootloader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBios

`func (o *KubevirtIoApiCoreV1Bootloader) GetBios() KubevirtIoApiCoreV1BIOS`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *KubevirtIoApiCoreV1Bootloader) GetBiosOk() (*KubevirtIoApiCoreV1BIOS, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *KubevirtIoApiCoreV1Bootloader) SetBios(v KubevirtIoApiCoreV1BIOS)`

SetBios sets Bios field to given value.

### HasBios

`func (o *KubevirtIoApiCoreV1Bootloader) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetEfi

`func (o *KubevirtIoApiCoreV1Bootloader) GetEfi() KubevirtIoApiCoreV1EFI`

GetEfi returns the Efi field if non-nil, zero value otherwise.

### GetEfiOk

`func (o *KubevirtIoApiCoreV1Bootloader) GetEfiOk() (*KubevirtIoApiCoreV1EFI, bool)`

GetEfiOk returns a tuple with the Efi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfi

`func (o *KubevirtIoApiCoreV1Bootloader) SetEfi(v KubevirtIoApiCoreV1EFI)`

SetEfi sets Efi field to given value.

### HasEfi

`func (o *KubevirtIoApiCoreV1Bootloader) HasEfi() bool`

HasEfi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


