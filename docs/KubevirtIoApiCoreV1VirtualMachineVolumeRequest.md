# KubevirtIoApiCoreV1VirtualMachineVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddVolumeOptions** | Pointer to [**KubevirtIoApiCoreV1AddVolumeOptions**](KubevirtIoApiCoreV1AddVolumeOptions.md) | AddVolumeOptions when set indicates a volume should be added. The details within this field specify how to add the volume | [optional] 
**RemoveVolumeOptions** | Pointer to [**KubevirtIoApiCoreV1RemoveVolumeOptions**](KubevirtIoApiCoreV1RemoveVolumeOptions.md) | RemoveVolumeOptions when set indicates a volume should be removed. The details within this field specify how to add the volume | [optional] 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineVolumeRequest

`func NewKubevirtIoApiCoreV1VirtualMachineVolumeRequest() *KubevirtIoApiCoreV1VirtualMachineVolumeRequest`

NewKubevirtIoApiCoreV1VirtualMachineVolumeRequest instantiates a new KubevirtIoApiCoreV1VirtualMachineVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineVolumeRequestWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineVolumeRequestWithDefaults() *KubevirtIoApiCoreV1VirtualMachineVolumeRequest`

NewKubevirtIoApiCoreV1VirtualMachineVolumeRequestWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) GetAddVolumeOptions() KubevirtIoApiCoreV1AddVolumeOptions`

GetAddVolumeOptions returns the AddVolumeOptions field if non-nil, zero value otherwise.

### GetAddVolumeOptionsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) GetAddVolumeOptionsOk() (*KubevirtIoApiCoreV1AddVolumeOptions, bool)`

GetAddVolumeOptionsOk returns a tuple with the AddVolumeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) SetAddVolumeOptions(v KubevirtIoApiCoreV1AddVolumeOptions)`

SetAddVolumeOptions sets AddVolumeOptions field to given value.

### HasAddVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) HasAddVolumeOptions() bool`

HasAddVolumeOptions returns a boolean if a field has been set.

### GetRemoveVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) GetRemoveVolumeOptions() KubevirtIoApiCoreV1RemoveVolumeOptions`

GetRemoveVolumeOptions returns the RemoveVolumeOptions field if non-nil, zero value otherwise.

### GetRemoveVolumeOptionsOk

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) GetRemoveVolumeOptionsOk() (*KubevirtIoApiCoreV1RemoveVolumeOptions, bool)`

GetRemoveVolumeOptionsOk returns a tuple with the RemoveVolumeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) SetRemoveVolumeOptions(v KubevirtIoApiCoreV1RemoveVolumeOptions)`

SetRemoveVolumeOptions sets RemoveVolumeOptions field to given value.

### HasRemoveVolumeOptions

`func (o *KubevirtIoApiCoreV1VirtualMachineVolumeRequest) HasRemoveVolumeOptions() bool`

HasRemoveVolumeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


