# KubevirtIoApiCoreV1Firmware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bootloader** | Pointer to [**KubevirtIoApiCoreV1Bootloader**](KubevirtIoApiCoreV1Bootloader.md) | Settings to control the bootloader that is used. | [optional] 
**KernelBoot** | Pointer to [**KubevirtIoApiCoreV1KernelBoot**](KubevirtIoApiCoreV1KernelBoot.md) | Settings to set the kernel for booting. | [optional] 
**Serial** | Pointer to **string** | The system-serial-number in SMBIOS | [optional] 
**Uuid** | Pointer to **string** | UUID reported by the vmi bios. Defaults to a random generated uid. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Firmware

`func NewKubevirtIoApiCoreV1Firmware() *KubevirtIoApiCoreV1Firmware`

NewKubevirtIoApiCoreV1Firmware instantiates a new KubevirtIoApiCoreV1Firmware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1FirmwareWithDefaults

`func NewKubevirtIoApiCoreV1FirmwareWithDefaults() *KubevirtIoApiCoreV1Firmware`

NewKubevirtIoApiCoreV1FirmwareWithDefaults instantiates a new KubevirtIoApiCoreV1Firmware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootloader

`func (o *KubevirtIoApiCoreV1Firmware) GetBootloader() KubevirtIoApiCoreV1Bootloader`

GetBootloader returns the Bootloader field if non-nil, zero value otherwise.

### GetBootloaderOk

`func (o *KubevirtIoApiCoreV1Firmware) GetBootloaderOk() (*KubevirtIoApiCoreV1Bootloader, bool)`

GetBootloaderOk returns a tuple with the Bootloader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootloader

`func (o *KubevirtIoApiCoreV1Firmware) SetBootloader(v KubevirtIoApiCoreV1Bootloader)`

SetBootloader sets Bootloader field to given value.

### HasBootloader

`func (o *KubevirtIoApiCoreV1Firmware) HasBootloader() bool`

HasBootloader returns a boolean if a field has been set.

### GetKernelBoot

`func (o *KubevirtIoApiCoreV1Firmware) GetKernelBoot() KubevirtIoApiCoreV1KernelBoot`

GetKernelBoot returns the KernelBoot field if non-nil, zero value otherwise.

### GetKernelBootOk

`func (o *KubevirtIoApiCoreV1Firmware) GetKernelBootOk() (*KubevirtIoApiCoreV1KernelBoot, bool)`

GetKernelBootOk returns a tuple with the KernelBoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelBoot

`func (o *KubevirtIoApiCoreV1Firmware) SetKernelBoot(v KubevirtIoApiCoreV1KernelBoot)`

SetKernelBoot sets KernelBoot field to given value.

### HasKernelBoot

`func (o *KubevirtIoApiCoreV1Firmware) HasKernelBoot() bool`

HasKernelBoot returns a boolean if a field has been set.

### GetSerial

`func (o *KubevirtIoApiCoreV1Firmware) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *KubevirtIoApiCoreV1Firmware) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *KubevirtIoApiCoreV1Firmware) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *KubevirtIoApiCoreV1Firmware) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUuid

`func (o *KubevirtIoApiCoreV1Firmware) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *KubevirtIoApiCoreV1Firmware) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *KubevirtIoApiCoreV1Firmware) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *KubevirtIoApiCoreV1Firmware) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


