# HarvesterhciIoV1beta1VirtualMachineBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**HarvesterhciIoV1beta1VirtualMachineBackupSpec**](HarvesterhciIoV1beta1VirtualMachineBackupSpec.md) |  | [default to {}]
**Status** | Pointer to [**HarvesterhciIoV1beta1VirtualMachineBackupStatus**](HarvesterhciIoV1beta1VirtualMachineBackupStatus.md) |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineBackup

`func NewHarvesterhciIoV1beta1VirtualMachineBackup(spec HarvesterhciIoV1beta1VirtualMachineBackupSpec, ) *HarvesterhciIoV1beta1VirtualMachineBackup`

NewHarvesterhciIoV1beta1VirtualMachineBackup instantiates a new HarvesterhciIoV1beta1VirtualMachineBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineBackupWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineBackupWithDefaults() *HarvesterhciIoV1beta1VirtualMachineBackup`

NewHarvesterhciIoV1beta1VirtualMachineBackupWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetSpec() HarvesterhciIoV1beta1VirtualMachineBackupSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetSpecOk() (*HarvesterhciIoV1beta1VirtualMachineBackupSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetSpec(v HarvesterhciIoV1beta1VirtualMachineBackupSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetStatus() HarvesterhciIoV1beta1VirtualMachineBackupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) GetStatusOk() (*HarvesterhciIoV1beta1VirtualMachineBackupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) SetStatus(v HarvesterhciIoV1beta1VirtualMachineBackupStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HarvesterhciIoV1beta1VirtualMachineBackup) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


