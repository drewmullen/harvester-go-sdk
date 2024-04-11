# KubevirtIoApiCoreV1VirtualMachineInstanceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | [**[]KubevirtIoApiCoreV1VirtualMachineInstance**](KubevirtIoApiCoreV1VirtualMachineInstance.md) |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceList

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceList(items []KubevirtIoApiCoreV1VirtualMachineInstance, ) *KubevirtIoApiCoreV1VirtualMachineInstanceList`

NewKubevirtIoApiCoreV1VirtualMachineInstanceList instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceListWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceListWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceList`

NewKubevirtIoApiCoreV1VirtualMachineInstanceListWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetItems() []KubevirtIoApiCoreV1VirtualMachineInstance`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetItemsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstance, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetItems(v []KubevirtIoApiCoreV1VirtualMachineInstance)`

SetItems sets Items field to given value.


### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


