# NetworkHarvesterhciIoV1beta1NodeNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]NetworkHarvesterhciIoV1beta1NodeNetwork**](NetworkHarvesterhciIoV1beta1NodeNetwork.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [default to {}]

## Methods

### NewNetworkHarvesterhciIoV1beta1NodeNetworkList

`func NewNetworkHarvesterhciIoV1beta1NodeNetworkList(items []NetworkHarvesterhciIoV1beta1NodeNetwork, metadata K8sIoV1ListMeta, ) *NetworkHarvesterhciIoV1beta1NodeNetworkList`

NewNetworkHarvesterhciIoV1beta1NodeNetworkList instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1NodeNetworkListWithDefaults

`func NewNetworkHarvesterhciIoV1beta1NodeNetworkListWithDefaults() *NetworkHarvesterhciIoV1beta1NodeNetworkList`

NewNetworkHarvesterhciIoV1beta1NodeNetworkListWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetItems() []NetworkHarvesterhciIoV1beta1NodeNetwork`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetItemsOk() (*[]NetworkHarvesterhciIoV1beta1NodeNetwork, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetItems(v []NetworkHarvesterhciIoV1beta1NodeNetwork)`

SetItems sets Items field to given value.


### GetKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


