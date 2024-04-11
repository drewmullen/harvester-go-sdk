# NetworkHarvesterhciIoV1beta1ClusterNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Status** | Pointer to [**NetworkHarvesterhciIoV1beta1ClusterNetworkStatus**](NetworkHarvesterhciIoV1beta1ClusterNetworkStatus.md) |  | [optional] [default to {}]

## Methods

### NewNetworkHarvesterhciIoV1beta1ClusterNetwork

`func NewNetworkHarvesterhciIoV1beta1ClusterNetwork() *NetworkHarvesterhciIoV1beta1ClusterNetwork`

NewNetworkHarvesterhciIoV1beta1ClusterNetwork instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1ClusterNetworkWithDefaults

`func NewNetworkHarvesterhciIoV1beta1ClusterNetworkWithDefaults() *NetworkHarvesterhciIoV1beta1ClusterNetwork`

NewNetworkHarvesterhciIoV1beta1ClusterNetworkWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetStatus() NetworkHarvesterhciIoV1beta1ClusterNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetStatusOk() (*NetworkHarvesterhciIoV1beta1ClusterNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetStatus(v NetworkHarvesterhciIoV1beta1ClusterNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


