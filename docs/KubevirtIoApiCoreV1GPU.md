# KubevirtIoApiCoreV1GPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | **string** |  | [default to ""]
**Name** | **string** |  | [default to ""]
**Tag** | Pointer to **string** |  | [optional] 
**VirtualGPUOptions** | Pointer to [**KubevirtIoApiCoreV1VGPUOptions**](KubevirtIoApiCoreV1VGPUOptions.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1GPU

`func NewKubevirtIoApiCoreV1GPU(deviceName string, name string, ) *KubevirtIoApiCoreV1GPU`

NewKubevirtIoApiCoreV1GPU instantiates a new KubevirtIoApiCoreV1GPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1GPUWithDefaults

`func NewKubevirtIoApiCoreV1GPUWithDefaults() *KubevirtIoApiCoreV1GPU`

NewKubevirtIoApiCoreV1GPUWithDefaults instantiates a new KubevirtIoApiCoreV1GPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *KubevirtIoApiCoreV1GPU) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *KubevirtIoApiCoreV1GPU) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *KubevirtIoApiCoreV1GPU) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetName

`func (o *KubevirtIoApiCoreV1GPU) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1GPU) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1GPU) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *KubevirtIoApiCoreV1GPU) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *KubevirtIoApiCoreV1GPU) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *KubevirtIoApiCoreV1GPU) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *KubevirtIoApiCoreV1GPU) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVirtualGPUOptions

`func (o *KubevirtIoApiCoreV1GPU) GetVirtualGPUOptions() KubevirtIoApiCoreV1VGPUOptions`

GetVirtualGPUOptions returns the VirtualGPUOptions field if non-nil, zero value otherwise.

### GetVirtualGPUOptionsOk

`func (o *KubevirtIoApiCoreV1GPU) GetVirtualGPUOptionsOk() (*KubevirtIoApiCoreV1VGPUOptions, bool)`

GetVirtualGPUOptionsOk returns a tuple with the VirtualGPUOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualGPUOptions

`func (o *KubevirtIoApiCoreV1GPU) SetVirtualGPUOptions(v KubevirtIoApiCoreV1VGPUOptions)`

SetVirtualGPUOptions sets VirtualGPUOptions field to given value.

### HasVirtualGPUOptions

`func (o *KubevirtIoApiCoreV1GPU) HasVirtualGPUOptions() bool`

HasVirtualGPUOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


