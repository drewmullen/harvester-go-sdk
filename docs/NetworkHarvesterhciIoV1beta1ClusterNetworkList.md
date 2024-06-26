# NetworkHarvesterhciIoV1beta1ClusterNetworkList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]NetworkHarvesterhciIoV1beta1ClusterNetwork**](NetworkHarvesterhciIoV1beta1ClusterNetwork.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [default to {}]

## Methods

### NewNetworkHarvesterhciIoV1beta1ClusterNetworkList

`func NewNetworkHarvesterhciIoV1beta1ClusterNetworkList(items []NetworkHarvesterhciIoV1beta1ClusterNetwork, metadata K8sIoV1ListMeta, ) *NetworkHarvesterhciIoV1beta1ClusterNetworkList`

NewNetworkHarvesterhciIoV1beta1ClusterNetworkList instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetworkList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1ClusterNetworkListWithDefaults

`func NewNetworkHarvesterhciIoV1beta1ClusterNetworkListWithDefaults() *NetworkHarvesterhciIoV1beta1ClusterNetworkList`

NewNetworkHarvesterhciIoV1beta1ClusterNetworkListWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetworkList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetItems() []NetworkHarvesterhciIoV1beta1ClusterNetwork`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetItemsOk() (*[]NetworkHarvesterhciIoV1beta1ClusterNetwork, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetItems(v []NetworkHarvesterhciIoV1beta1ClusterNetwork)`

SetItems sets Items field to given value.


### GetKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetworkList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


