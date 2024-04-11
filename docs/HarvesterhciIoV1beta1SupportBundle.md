# HarvesterhciIoV1beta1SupportBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**HarvesterhciIoV1beta1SupportBundleSpec**](HarvesterhciIoV1beta1SupportBundleSpec.md) |  | [default to {}]
**Status** | Pointer to [**HarvesterhciIoV1beta1SupportBundleStatus**](HarvesterhciIoV1beta1SupportBundleStatus.md) |  | [optional] [default to {}]

## Methods

### NewHarvesterhciIoV1beta1SupportBundle

`func NewHarvesterhciIoV1beta1SupportBundle(spec HarvesterhciIoV1beta1SupportBundleSpec, ) *HarvesterhciIoV1beta1SupportBundle`

NewHarvesterhciIoV1beta1SupportBundle instantiates a new HarvesterhciIoV1beta1SupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1SupportBundleWithDefaults

`func NewHarvesterhciIoV1beta1SupportBundleWithDefaults() *HarvesterhciIoV1beta1SupportBundle`

NewHarvesterhciIoV1beta1SupportBundleWithDefaults instantiates a new HarvesterhciIoV1beta1SupportBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundle) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1SupportBundle) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundle) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HarvesterhciIoV1beta1SupportBundle) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *HarvesterhciIoV1beta1SupportBundle) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1SupportBundle) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1SupportBundle) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HarvesterhciIoV1beta1SupportBundle) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *HarvesterhciIoV1beta1SupportBundle) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1SupportBundle) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1SupportBundle) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1SupportBundle) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1SupportBundle) GetSpec() HarvesterhciIoV1beta1SupportBundleSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1SupportBundle) GetSpecOk() (*HarvesterhciIoV1beta1SupportBundleSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1SupportBundle) SetSpec(v HarvesterhciIoV1beta1SupportBundleSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *HarvesterhciIoV1beta1SupportBundle) GetStatus() HarvesterhciIoV1beta1SupportBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HarvesterhciIoV1beta1SupportBundle) GetStatusOk() (*HarvesterhciIoV1beta1SupportBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HarvesterhciIoV1beta1SupportBundle) SetStatus(v HarvesterhciIoV1beta1SupportBundleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HarvesterhciIoV1beta1SupportBundle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


