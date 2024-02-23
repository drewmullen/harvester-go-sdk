# KubevirtIoApiCoreV1VirtualMachineList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Items** | [**[]KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md) |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineList

`func NewKubevirtIoApiCoreV1VirtualMachineList(apiVersion string, items []KubevirtIoApiCoreV1VirtualMachine, kind string, ) *KubevirtIoApiCoreV1VirtualMachineList`

NewKubevirtIoApiCoreV1VirtualMachineList instantiates a new KubevirtIoApiCoreV1VirtualMachineList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineListWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineListWithDefaults() *KubevirtIoApiCoreV1VirtualMachineList`

NewKubevirtIoApiCoreV1VirtualMachineListWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetItems() []KubevirtIoApiCoreV1VirtualMachine`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetItemsOk() (*[]KubevirtIoApiCoreV1VirtualMachine, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineList) SetItems(v []KubevirtIoApiCoreV1VirtualMachine)`

SetItems sets Items field to given value.


### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


