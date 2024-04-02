# K8sIoV1PersistentVolumeClaimList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Items** | [**[]K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md) |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [optional] [default to {}]

## Methods

### NewK8sIoV1PersistentVolumeClaimList

`func NewK8sIoV1PersistentVolumeClaimList(apiVersion string, items []K8sIoV1PersistentVolumeClaim, kind string, ) *K8sIoV1PersistentVolumeClaimList`

NewK8sIoV1PersistentVolumeClaimList instantiates a new K8sIoV1PersistentVolumeClaimList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimListWithDefaults

`func NewK8sIoV1PersistentVolumeClaimListWithDefaults() *K8sIoV1PersistentVolumeClaimList`

NewK8sIoV1PersistentVolumeClaimListWithDefaults instantiates a new K8sIoV1PersistentVolumeClaimList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1PersistentVolumeClaimList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1PersistentVolumeClaimList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1PersistentVolumeClaimList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetItems

`func (o *K8sIoV1PersistentVolumeClaimList) GetItems() []K8sIoV1PersistentVolumeClaim`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *K8sIoV1PersistentVolumeClaimList) GetItemsOk() (*[]K8sIoV1PersistentVolumeClaim, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *K8sIoV1PersistentVolumeClaimList) SetItems(v []K8sIoV1PersistentVolumeClaim)`

SetItems sets Items field to given value.


### GetKind

`func (o *K8sIoV1PersistentVolumeClaimList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1PersistentVolumeClaimList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1PersistentVolumeClaimList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *K8sIoV1PersistentVolumeClaimList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *K8sIoV1PersistentVolumeClaimList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *K8sIoV1PersistentVolumeClaimList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *K8sIoV1PersistentVolumeClaimList) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


