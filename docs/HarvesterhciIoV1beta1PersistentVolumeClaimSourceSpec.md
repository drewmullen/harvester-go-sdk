# HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**K8sIoV1PersistentVolumeClaimSpec**](K8sIoV1PersistentVolumeClaimSpec.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec

`func NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec() *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec`

NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec instantiates a new HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpecWithDefaults

`func NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpecWithDefaults() *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec`

NewHarvesterhciIoV1beta1PersistentVolumeClaimSourceSpecWithDefaults instantiates a new HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) GetSpec() K8sIoV1PersistentVolumeClaimSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) GetSpecOk() (*K8sIoV1PersistentVolumeClaimSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) SetSpec(v K8sIoV1PersistentVolumeClaimSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *HarvesterhciIoV1beta1PersistentVolumeClaimSourceSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


