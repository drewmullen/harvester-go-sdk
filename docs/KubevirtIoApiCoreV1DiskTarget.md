# KubevirtIoApiCoreV1DiskTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | Pointer to **string** | Bus indicates the type of disk device to emulate. supported values: virtio, sata, scsi, usb. | [optional] 
**PciAddress** | Pointer to **string** | If specified, the virtual disk will be placed on the guests pci address with the specified PCI address. For example: 0000:81:01.10 | [optional] 
**Readonly** | Pointer to **bool** | ReadOnly. Defaults to false. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1DiskTarget

`func NewKubevirtIoApiCoreV1DiskTarget() *KubevirtIoApiCoreV1DiskTarget`

NewKubevirtIoApiCoreV1DiskTarget instantiates a new KubevirtIoApiCoreV1DiskTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DiskTargetWithDefaults

`func NewKubevirtIoApiCoreV1DiskTargetWithDefaults() *KubevirtIoApiCoreV1DiskTarget`

NewKubevirtIoApiCoreV1DiskTargetWithDefaults instantiates a new KubevirtIoApiCoreV1DiskTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *KubevirtIoApiCoreV1DiskTarget) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *KubevirtIoApiCoreV1DiskTarget) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *KubevirtIoApiCoreV1DiskTarget) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *KubevirtIoApiCoreV1DiskTarget) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetPciAddress

`func (o *KubevirtIoApiCoreV1DiskTarget) GetPciAddress() string`

GetPciAddress returns the PciAddress field if non-nil, zero value otherwise.

### GetPciAddressOk

`func (o *KubevirtIoApiCoreV1DiskTarget) GetPciAddressOk() (*string, bool)`

GetPciAddressOk returns a tuple with the PciAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciAddress

`func (o *KubevirtIoApiCoreV1DiskTarget) SetPciAddress(v string)`

SetPciAddress sets PciAddress field to given value.

### HasPciAddress

`func (o *KubevirtIoApiCoreV1DiskTarget) HasPciAddress() bool`

HasPciAddress returns a boolean if a field has been set.

### GetReadonly

`func (o *KubevirtIoApiCoreV1DiskTarget) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *KubevirtIoApiCoreV1DiskTarget) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *KubevirtIoApiCoreV1DiskTarget) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.

### HasReadonly

`func (o *KubevirtIoApiCoreV1DiskTarget) HasReadonly() bool`

HasReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


