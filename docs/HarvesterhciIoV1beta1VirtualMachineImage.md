# HarvesterhciIoV1beta1VirtualMachineImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] 
**Spec** | [**HarvesterhciIoV1beta1VirtualMachineImageSpec**](HarvesterhciIoV1beta1VirtualMachineImageSpec.md) |  | 
**Status** | Pointer to [**HarvesterhciIoV1beta1VirtualMachineImageStatus**](HarvesterhciIoV1beta1VirtualMachineImageStatus.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineImage

`func NewHarvesterhciIoV1beta1VirtualMachineImage(apiVersion string, kind string, spec HarvesterhciIoV1beta1VirtualMachineImageSpec, ) *HarvesterhciIoV1beta1VirtualMachineImage`

NewHarvesterhciIoV1beta1VirtualMachineImage instantiates a new HarvesterhciIoV1beta1VirtualMachineImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineImageWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineImageWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImage`

NewHarvesterhciIoV1beta1VirtualMachineImageWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetSpec() HarvesterhciIoV1beta1VirtualMachineImageSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineImageSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetSpec(v HarvesterhciIoV1beta1VirtualMachineImageSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetStatus() HarvesterhciIoV1beta1VirtualMachineImageStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineImageStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) SetStatus(v HarvesterhciIoV1beta1VirtualMachineImageStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


