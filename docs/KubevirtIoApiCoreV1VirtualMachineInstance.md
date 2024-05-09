# KubevirtIoApiCoreV1VirtualMachineInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**KubevirtIoApiCoreV1VirtualMachineInstanceSpec**](KubevirtIoApiCoreV1VirtualMachineInstanceSpec.md) | VirtualMachineInstance Spec contains the VirtualMachineInstance specification. | [default to {}]
**Status** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceStatus**](KubevirtIoApiCoreV1VirtualMachineInstanceStatus.md) | Status is the high level overview of how the VirtualMachineInstance is doing. It contains information available to controllers and users. | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstance

`func NewKubevirtIoApiCoreV1VirtualMachineInstance(spec KubevirtIoApiCoreV1VirtualMachineInstanceSpec, ) *KubevirtIoApiCoreV1VirtualMachineInstance`

NewKubevirtIoApiCoreV1VirtualMachineInstance instantiates a new KubevirtIoApiCoreV1VirtualMachineInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstance`

NewKubevirtIoApiCoreV1VirtualMachineInstanceWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetSpec() KubevirtIoApiCoreV1VirtualMachineInstanceSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetSpec(v KubevirtIoApiCoreV1VirtualMachineInstanceSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetStatus() KubevirtIoApiCoreV1VirtualMachineInstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) GetStatusOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) SetStatus(v KubevirtIoApiCoreV1VirtualMachineInstanceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


