# HarvesterhciIoV1beta1KeyPairList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | [**[]HarvesterhciIoV1beta1KeyPair**](HarvesterhciIoV1beta1KeyPair.md) |  | 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [default to {}]

## Methods

### NewHarvesterhciIoV1beta1KeyPairList

`func NewHarvesterhciIoV1beta1KeyPairList(items []HarvesterhciIoV1beta1KeyPair, metadata K8sIoV1ListMeta, ) *HarvesterhciIoV1beta1KeyPairList`

NewHarvesterhciIoV1beta1KeyPairList instantiates a new HarvesterhciIoV1beta1KeyPairList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1KeyPairListWithDefaults

`func NewHarvesterhciIoV1beta1KeyPairListWithDefaults() *HarvesterhciIoV1beta1KeyPairList`

NewHarvesterhciIoV1beta1KeyPairListWithDefaults instantiates a new HarvesterhciIoV1beta1KeyPairList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1KeyPairList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1KeyPairList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1KeyPairList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HarvesterhciIoV1beta1KeyPairList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *HarvesterhciIoV1beta1KeyPairList) GetItems() []HarvesterhciIoV1beta1KeyPair`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HarvesterhciIoV1beta1KeyPairList) GetItemsOk() (*[]HarvesterhciIoV1beta1KeyPair, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HarvesterhciIoV1beta1KeyPairList) SetItems(v []HarvesterhciIoV1beta1KeyPair)`

SetItems sets Items field to given value.


### GetKind

`func (o *HarvesterhciIoV1beta1KeyPairList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1KeyPairList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1KeyPairList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HarvesterhciIoV1beta1KeyPairList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *HarvesterhciIoV1beta1KeyPairList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1KeyPairList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1KeyPairList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


