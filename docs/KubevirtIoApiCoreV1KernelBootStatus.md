# KubevirtIoApiCoreV1KernelBootStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitrdInfo** | Pointer to [**KubevirtIoApiCoreV1InitrdInfo**](KubevirtIoApiCoreV1InitrdInfo.md) | InitrdInfo show info about the initrd file | [optional] 
**KernelInfo** | Pointer to [**KubevirtIoApiCoreV1KernelInfo**](KubevirtIoApiCoreV1KernelInfo.md) | KernelInfo show info about the kernel image | [optional] 

## Methods

### NewKubevirtIoApiCoreV1KernelBootStatus

`func NewKubevirtIoApiCoreV1KernelBootStatus() *KubevirtIoApiCoreV1KernelBootStatus`

NewKubevirtIoApiCoreV1KernelBootStatus instantiates a new KubevirtIoApiCoreV1KernelBootStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1KernelBootStatusWithDefaults

`func NewKubevirtIoApiCoreV1KernelBootStatusWithDefaults() *KubevirtIoApiCoreV1KernelBootStatus`

NewKubevirtIoApiCoreV1KernelBootStatusWithDefaults instantiates a new KubevirtIoApiCoreV1KernelBootStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitrdInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) GetInitrdInfo() KubevirtIoApiCoreV1InitrdInfo`

GetInitrdInfo returns the InitrdInfo field if non-nil, zero value otherwise.

### GetInitrdInfoOk

`func (o *KubevirtIoApiCoreV1KernelBootStatus) GetInitrdInfoOk() (*KubevirtIoApiCoreV1InitrdInfo, bool)`

GetInitrdInfoOk returns a tuple with the InitrdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitrdInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) SetInitrdInfo(v KubevirtIoApiCoreV1InitrdInfo)`

SetInitrdInfo sets InitrdInfo field to given value.

### HasInitrdInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) HasInitrdInfo() bool`

HasInitrdInfo returns a boolean if a field has been set.

### GetKernelInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) GetKernelInfo() KubevirtIoApiCoreV1KernelInfo`

GetKernelInfo returns the KernelInfo field if non-nil, zero value otherwise.

### GetKernelInfoOk

`func (o *KubevirtIoApiCoreV1KernelBootStatus) GetKernelInfoOk() (*KubevirtIoApiCoreV1KernelInfo, bool)`

GetKernelInfoOk returns a tuple with the KernelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) SetKernelInfo(v KubevirtIoApiCoreV1KernelInfo)`

SetKernelInfo sets KernelInfo field to given value.

### HasKernelInfo

`func (o *KubevirtIoApiCoreV1KernelBootStatus) HasKernelInfo() bool`

HasKernelInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


