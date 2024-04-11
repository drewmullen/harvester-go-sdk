# NetworkHarvesterhciIoV1beta1NodeNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | Pointer to [**NetworkHarvesterhciIoV1beta1NodeNetworkSpec**](NetworkHarvesterhciIoV1beta1NodeNetworkSpec.md) |  | [optional] [default to {}]
**Status** | Pointer to [**NetworkHarvesterhciIoV1beta1NodeNetworkStatus**](NetworkHarvesterhciIoV1beta1NodeNetworkStatus.md) |  | [optional] [default to {}]

## Methods

### NewNetworkHarvesterhciIoV1beta1NodeNetwork

`func NewNetworkHarvesterhciIoV1beta1NodeNetwork() *NetworkHarvesterhciIoV1beta1NodeNetwork`

NewNetworkHarvesterhciIoV1beta1NodeNetwork instantiates a new NetworkHarvesterhciIoV1beta1NodeNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1NodeNetworkWithDefaults

`func NewNetworkHarvesterhciIoV1beta1NodeNetworkWithDefaults() *NetworkHarvesterhciIoV1beta1NodeNetwork`

NewNetworkHarvesterhciIoV1beta1NodeNetworkWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1NodeNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetSpec() NetworkHarvesterhciIoV1beta1NodeNetworkSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetSpecOk() (*NetworkHarvesterhciIoV1beta1NodeNetworkSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) SetSpec(v NetworkHarvesterhciIoV1beta1NodeNetworkSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetStatus() NetworkHarvesterhciIoV1beta1NodeNetworkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) GetStatusOk() (*NetworkHarvesterhciIoV1beta1NodeNetworkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) SetStatus(v NetworkHarvesterhciIoV1beta1NodeNetworkStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetwork) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


