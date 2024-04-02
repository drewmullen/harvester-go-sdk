# HarvesterhciIoV1beta1VirtualMachineRestoreSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionPolicy** | Pointer to **string** |  | [optional] 
**KeepMacAddress** | Pointer to **bool** |  | [optional] 
**NewVM** | Pointer to **bool** |  | [optional] 
**Target** | [**K8sIoV1TypedLocalObjectReference**](K8sIoV1TypedLocalObjectReference.md) |  | [default to {}]
**VirtualMachineBackupName** | **string** |  | [default to ""]
**VirtualMachineBackupNamespace** | **string** |  | [default to ""]

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineRestoreSpec

`func NewHarvesterhciIoV1beta1VirtualMachineRestoreSpec(target K8sIoV1TypedLocalObjectReference, virtualMachineBackupName string, virtualMachineBackupNamespace string, ) *HarvesterhciIoV1beta1VirtualMachineRestoreSpec`

NewHarvesterhciIoV1beta1VirtualMachineRestoreSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineRestoreSpecWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineRestoreSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineRestoreSpec`

NewHarvesterhciIoV1beta1VirtualMachineRestoreSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineRestoreSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionPolicy

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetDeletionPolicy() string`

GetDeletionPolicy returns the DeletionPolicy field if non-nil, zero value otherwise.

### GetDeletionPolicyOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetDeletionPolicyOk() (*string, bool)`

GetDeletionPolicyOk returns a tuple with the DeletionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionPolicy

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetDeletionPolicy(v string)`

SetDeletionPolicy sets DeletionPolicy field to given value.

### HasDeletionPolicy

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) HasDeletionPolicy() bool`

HasDeletionPolicy returns a boolean if a field has been set.

### GetKeepMacAddress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetKeepMacAddress() bool`

GetKeepMacAddress returns the KeepMacAddress field if non-nil, zero value otherwise.

### GetKeepMacAddressOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetKeepMacAddressOk() (*bool, bool)`

GetKeepMacAddressOk returns a tuple with the KeepMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepMacAddress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetKeepMacAddress(v bool)`

SetKeepMacAddress sets KeepMacAddress field to given value.

### HasKeepMacAddress

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) HasKeepMacAddress() bool`

HasKeepMacAddress returns a boolean if a field has been set.

### GetNewVM

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetNewVM() bool`

GetNewVM returns the NewVM field if non-nil, zero value otherwise.

### GetNewVMOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetNewVMOk() (*bool, bool)`

GetNewVMOk returns a tuple with the NewVM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVM

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetNewVM(v bool)`

SetNewVM sets NewVM field to given value.

### HasNewVM

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) HasNewVM() bool`

HasNewVM returns a boolean if a field has been set.

### GetTarget

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetTarget() K8sIoV1TypedLocalObjectReference`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetTargetOk() (*K8sIoV1TypedLocalObjectReference, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetTarget(v K8sIoV1TypedLocalObjectReference)`

SetTarget sets Target field to given value.


### GetVirtualMachineBackupName

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetVirtualMachineBackupName() string`

GetVirtualMachineBackupName returns the VirtualMachineBackupName field if non-nil, zero value otherwise.

### GetVirtualMachineBackupNameOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetVirtualMachineBackupNameOk() (*string, bool)`

GetVirtualMachineBackupNameOk returns a tuple with the VirtualMachineBackupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineBackupName

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetVirtualMachineBackupName(v string)`

SetVirtualMachineBackupName sets VirtualMachineBackupName field to given value.


### GetVirtualMachineBackupNamespace

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetVirtualMachineBackupNamespace() string`

GetVirtualMachineBackupNamespace returns the VirtualMachineBackupNamespace field if non-nil, zero value otherwise.

### GetVirtualMachineBackupNamespaceOk

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) GetVirtualMachineBackupNamespaceOk() (*string, bool)`

GetVirtualMachineBackupNamespaceOk returns a tuple with the VirtualMachineBackupNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachineBackupNamespace

`func (o *HarvesterhciIoV1beta1VirtualMachineRestoreSpec) SetVirtualMachineBackupNamespace(v string)`

SetVirtualMachineBackupNamespace sets VirtualMachineBackupNamespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


