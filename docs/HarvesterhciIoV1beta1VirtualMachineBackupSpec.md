# HarvesterhciIoV1beta1VirtualMachineBackupSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**K8sIoV1TypedLocalObjectReference**](K8sIoV1TypedLocalObjectReference.md) |  | [default to {}]
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineBackupSpec

`func NewHarvesterhciIoV1beta1VirtualMachineBackupSpec(source K8sIoV1TypedLocalObjectReference, ) *HarvesterhciIoV1beta1VirtualMachineBackupSpec`

NewHarvesterhciIoV1beta1VirtualMachineBackupSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineBackupSpecWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineBackupSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineBackupSpec`

NewHarvesterhciIoV1beta1VirtualMachineBackupSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineBackupSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) GetSource() K8sIoV1TypedLocalObjectReference`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) GetSourceOk() (*K8sIoV1TypedLocalObjectReference, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) SetSource(v K8sIoV1TypedLocalObjectReference)`

SetSource sets Source field to given value.


### GetType

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HarvesterhciIoV1beta1VirtualMachineBackupSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


