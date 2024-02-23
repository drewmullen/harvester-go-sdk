# HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus

`func NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus() *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus`

NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatusWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus`

NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionStatus) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


