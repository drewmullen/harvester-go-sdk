# HarvesterhciIoV1beta1VirtualMachineSourceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineSpec**](KubevirtIoApiCoreV1VirtualMachineSpec.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineSourceSpec

`func NewHarvesterhciIoV1beta1VirtualMachineSourceSpec() *HarvesterhciIoV1beta1VirtualMachineSourceSpec`

NewHarvesterhciIoV1beta1VirtualMachineSourceSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineSourceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineSourceSpecWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineSourceSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineSourceSpec`

NewHarvesterhciIoV1beta1VirtualMachineSourceSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineSourceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetSpec() KubevirtIoApiCoreV1VirtualMachineSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) SetSpec(v KubevirtIoApiCoreV1VirtualMachineSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineSourceSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


