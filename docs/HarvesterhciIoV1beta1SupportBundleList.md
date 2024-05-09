# HarvesterhciIoV1beta1SupportBundleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]HarvesterhciIoV1beta1SupportBundle**](HarvesterhciIoV1beta1SupportBundle.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [default to {}]

## Methods

### NewHarvesterhciIoV1beta1SupportBundleList

`func NewHarvesterhciIoV1beta1SupportBundleList(items []HarvesterhciIoV1beta1SupportBundle, metadata K8sIoV1ListMeta, ) *HarvesterhciIoV1beta1SupportBundleList`

NewHarvesterhciIoV1beta1SupportBundleList instantiates a new HarvesterhciIoV1beta1SupportBundleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1SupportBundleListWithDefaults

`func NewHarvesterhciIoV1beta1SupportBundleListWithDefaults() *HarvesterhciIoV1beta1SupportBundleList`

NewHarvesterhciIoV1beta1SupportBundleListWithDefaults instantiates a new HarvesterhciIoV1beta1SupportBundleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetItems() []HarvesterhciIoV1beta1SupportBundle`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetItemsOk() (*[]HarvesterhciIoV1beta1SupportBundle, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HarvesterhciIoV1beta1SupportBundleList) SetItems(v []HarvesterhciIoV1beta1SupportBundle)`

SetItems sets Items field to given value.


### GetKind

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1SupportBundleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HarvesterhciIoV1beta1SupportBundleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1SupportBundleList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1SupportBundleList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


