# HarvesterhciIoV1beta1VirtualMachineRestore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**HarvesterhciIoV1beta1VirtualMachineRestoreSpec**](HarvesterhciIoV1beta1VirtualMachineRestoreSpec.md) |  | [default to {}]
**Status** | Pointer to [**HarvesterhciIoV1beta1VirtualMachineRestoreStatus**](HarvesterhciIoV1beta1VirtualMachineRestoreStatus.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineRestore

`func NewHarvesterhciIoV1beta1VirtualMachineRestore(spec HarvesterhciIoV1beta1VirtualMachineRestoreSpec, ) *HarvesterhciIoV1beta1VirtualMachineRestore`

NewHarvesterhciIoV1beta1VirtualMachineRestore instantiates a new HarvesterhciIoV1beta1VirtualMachineRestore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineRestoreWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineRestoreWithDefaults() *HarvesterhciIoV1beta1VirtualMachineRestore`

NewHarvesterhciIoV1beta1VirtualMachineRestoreWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineRestore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetSpec() HarvesterhciIoV1beta1VirtualMachineRestoreSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineRestoreSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetSpec(v HarvesterhciIoV1beta1VirtualMachineRestoreSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetStatus() HarvesterhciIoV1beta1VirtualMachineRestoreStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineRestoreStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) SetStatus(v HarvesterhciIoV1beta1VirtualMachineRestoreStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineRestore) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


