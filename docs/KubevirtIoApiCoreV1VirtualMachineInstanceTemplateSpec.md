# KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceSpec**](KubevirtIoApiCoreV1VirtualMachineInstanceSpec.md) | VirtualMachineInstance Spec contains the VirtualMachineInstance specification. | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec() *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec`

NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpecWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpecWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec`

NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpecWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) GetSpec() KubevirtIoApiCoreV1VirtualMachineInstanceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) SetSpec(v KubevirtIoApiCoreV1VirtualMachineInstanceSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


