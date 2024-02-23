# KubevirtIoApiCoreV1Disk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockSize** | Pointer to [**KubevirtIoApiCoreV1BlockSize**](KubevirtIoApiCoreV1BlockSize.md) |  | [optional] 
**BootOrder** | Pointer to **int32** |  | [optional] 
**Cache** | Pointer to **string** |  | [optional] 
**Cdrom** | Pointer to [**KubevirtIoApiCoreV1CDRomTarget**](KubevirtIoApiCoreV1CDRomTarget.md) |  | [optional] 
**DedicatedIOThread** | Pointer to **bool** |  | [optional] 
**Disk** | Pointer to [**KubevirtIoApiCoreV1DiskTarget**](KubevirtIoApiCoreV1DiskTarget.md) |  | [optional] 
**ErrorPolicy** | Pointer to **string** |  | [optional] 
**Io** | Pointer to **string** |  | [optional] 
**Lun** | Pointer to [**KubevirtIoApiCoreV1LunTarget**](KubevirtIoApiCoreV1LunTarget.md) |  | [optional] 
**Name** | **string** |  | [default to ""]
**Serial** | Pointer to **string** |  | [optional] 
**Shareable** | Pointer to **bool** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Disk

`func NewKubevirtIoApiCoreV1Disk(name string, ) *KubevirtIoApiCoreV1Disk`

NewKubevirtIoApiCoreV1Disk instantiates a new KubevirtIoApiCoreV1Disk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DiskWithDefaults

`func NewKubevirtIoApiCoreV1DiskWithDefaults() *KubevirtIoApiCoreV1Disk`

NewKubevirtIoApiCoreV1DiskWithDefaults instantiates a new KubevirtIoApiCoreV1Disk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockSize

`func (o *KubevirtIoApiCoreV1Disk) GetBlockSize() KubevirtIoApiCoreV1BlockSize`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *KubevirtIoApiCoreV1Disk) GetBlockSizeOk() (*KubevirtIoApiCoreV1BlockSize, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *KubevirtIoApiCoreV1Disk) SetBlockSize(v KubevirtIoApiCoreV1BlockSize)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *KubevirtIoApiCoreV1Disk) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBootOrder

`func (o *KubevirtIoApiCoreV1Disk) GetBootOrder() int32`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *KubevirtIoApiCoreV1Disk) GetBootOrderOk() (*int32, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *KubevirtIoApiCoreV1Disk) SetBootOrder(v int32)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *KubevirtIoApiCoreV1Disk) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.

### GetCache

`func (o *KubevirtIoApiCoreV1Disk) GetCache() string`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *KubevirtIoApiCoreV1Disk) GetCacheOk() (*string, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *KubevirtIoApiCoreV1Disk) SetCache(v string)`

SetCache sets Cache field to given value.

### HasCache

`func (o *KubevirtIoApiCoreV1Disk) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetCdrom

`func (o *KubevirtIoApiCoreV1Disk) GetCdrom() KubevirtIoApiCoreV1CDRomTarget`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *KubevirtIoApiCoreV1Disk) GetCdromOk() (*KubevirtIoApiCoreV1CDRomTarget, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *KubevirtIoApiCoreV1Disk) SetCdrom(v KubevirtIoApiCoreV1CDRomTarget)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *KubevirtIoApiCoreV1Disk) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetDedicatedIOThread

`func (o *KubevirtIoApiCoreV1Disk) GetDedicatedIOThread() bool`

GetDedicatedIOThread returns the DedicatedIOThread field if non-nil, zero value otherwise.

### GetDedicatedIOThreadOk

`func (o *KubevirtIoApiCoreV1Disk) GetDedicatedIOThreadOk() (*bool, bool)`

GetDedicatedIOThreadOk returns a tuple with the DedicatedIOThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedIOThread

`func (o *KubevirtIoApiCoreV1Disk) SetDedicatedIOThread(v bool)`

SetDedicatedIOThread sets DedicatedIOThread field to given value.

### HasDedicatedIOThread

`func (o *KubevirtIoApiCoreV1Disk) HasDedicatedIOThread() bool`

HasDedicatedIOThread returns a boolean if a field has been set.

### GetDisk

`func (o *KubevirtIoApiCoreV1Disk) GetDisk() KubevirtIoApiCoreV1DiskTarget`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *KubevirtIoApiCoreV1Disk) GetDiskOk() (*KubevirtIoApiCoreV1DiskTarget, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *KubevirtIoApiCoreV1Disk) SetDisk(v KubevirtIoApiCoreV1DiskTarget)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *KubevirtIoApiCoreV1Disk) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetErrorPolicy

`func (o *KubevirtIoApiCoreV1Disk) GetErrorPolicy() string`

GetErrorPolicy returns the ErrorPolicy field if non-nil, zero value otherwise.

### GetErrorPolicyOk

`func (o *KubevirtIoApiCoreV1Disk) GetErrorPolicyOk() (*string, bool)`

GetErrorPolicyOk returns a tuple with the ErrorPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPolicy

`func (o *KubevirtIoApiCoreV1Disk) SetErrorPolicy(v string)`

SetErrorPolicy sets ErrorPolicy field to given value.

### HasErrorPolicy

`func (o *KubevirtIoApiCoreV1Disk) HasErrorPolicy() bool`

HasErrorPolicy returns a boolean if a field has been set.

### GetIo

`func (o *KubevirtIoApiCoreV1Disk) GetIo() string`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *KubevirtIoApiCoreV1Disk) GetIoOk() (*string, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *KubevirtIoApiCoreV1Disk) SetIo(v string)`

SetIo sets Io field to given value.

### HasIo

`func (o *KubevirtIoApiCoreV1Disk) HasIo() bool`

HasIo returns a boolean if a field has been set.

### GetLun

`func (o *KubevirtIoApiCoreV1Disk) GetLun() KubevirtIoApiCoreV1LunTarget`

GetLun returns the Lun field if non-nil, zero value otherwise.

### GetLunOk

`func (o *KubevirtIoApiCoreV1Disk) GetLunOk() (*KubevirtIoApiCoreV1LunTarget, bool)`

GetLunOk returns a tuple with the Lun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLun

`func (o *KubevirtIoApiCoreV1Disk) SetLun(v KubevirtIoApiCoreV1LunTarget)`

SetLun sets Lun field to given value.

### HasLun

`func (o *KubevirtIoApiCoreV1Disk) HasLun() bool`

HasLun returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Disk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Disk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Disk) SetName(v string)`

SetName sets Name field to given value.


### GetSerial

`func (o *KubevirtIoApiCoreV1Disk) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *KubevirtIoApiCoreV1Disk) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *KubevirtIoApiCoreV1Disk) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *KubevirtIoApiCoreV1Disk) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetShareable

`func (o *KubevirtIoApiCoreV1Disk) GetShareable() bool`

GetShareable returns the Shareable field if non-nil, zero value otherwise.

### GetShareableOk

`func (o *KubevirtIoApiCoreV1Disk) GetShareableOk() (*bool, bool)`

GetShareableOk returns a tuple with the Shareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareable

`func (o *KubevirtIoApiCoreV1Disk) SetShareable(v bool)`

SetShareable sets Shareable field to given value.

### HasShareable

`func (o *KubevirtIoApiCoreV1Disk) HasShareable() bool`

HasShareable returns a boolean if a field has been set.

### GetTag

`func (o *KubevirtIoApiCoreV1Disk) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *KubevirtIoApiCoreV1Disk) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *KubevirtIoApiCoreV1Disk) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *KubevirtIoApiCoreV1Disk) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


