# HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**KeyPairIds** | Pointer to **[]string** |  | [optional] 
**TemplateId** | **string** |  | [default to ""]
**Vm** | Pointer to [**HarvesterhciIoV1beta1VirtualMachineSourceSpec**](HarvesterhciIoV1beta1VirtualMachineSourceSpec.md) |  | [optional] [default to {}]

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec

`func NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec(templateId string, ) *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec`

NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpecWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec`

NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageId

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetKeyPairIds

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetKeyPairIds() []string`

GetKeyPairIds returns the KeyPairIds field if non-nil, zero value otherwise.

### GetKeyPairIdsOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetKeyPairIdsOk() (*[]string, bool)`

GetKeyPairIdsOk returns a tuple with the KeyPairIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPairIds

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) SetKeyPairIds(v []string)`

SetKeyPairIds sets KeyPairIds field to given value.

### HasKeyPairIds

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) HasKeyPairIds() bool`

HasKeyPairIds returns a boolean if a field has been set.

### GetTemplateId

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetVm

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetVm() HarvesterhciIoV1beta1VirtualMachineSourceSpec`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) GetVmOk() (*HarvesterhciIoV1beta1VirtualMachineSourceSpec, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) SetVm(v HarvesterhciIoV1beta1VirtualMachineSourceSpec)`

SetVm sets Vm field to given value.

### HasVm

`func (o *HarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


