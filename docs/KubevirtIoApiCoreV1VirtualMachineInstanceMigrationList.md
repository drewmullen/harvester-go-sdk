# KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Items** | [**[]KubevirtIoApiCoreV1VirtualMachineInstanceMigration**](KubevirtIoApiCoreV1VirtualMachineInstanceMigration.md) |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationList

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationList(apiVersion string, items []KubevirtIoApiCoreV1VirtualMachineInstanceMigration, kind string, ) *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationList instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationListWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationListWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationListWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetItems() []KubevirtIoApiCoreV1VirtualMachineInstanceMigration`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetItemsOk() (*[]KubevirtIoApiCoreV1VirtualMachineInstanceMigration, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) SetItems(v []KubevirtIoApiCoreV1VirtualMachineInstanceMigration)`

SetItems sets Items field to given value.


### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigrationList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


