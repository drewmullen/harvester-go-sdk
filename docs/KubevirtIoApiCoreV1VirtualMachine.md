# KubevirtIoApiCoreV1VirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**KubevirtIoApiCoreV1VirtualMachineSpec**](KubevirtIoApiCoreV1VirtualMachineSpec.md) | Spec contains the specification of VirtualMachineInstance created | [default to {}]
**Status** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineStatus**](KubevirtIoApiCoreV1VirtualMachineStatus.md) | Status holds the current state of the controller and brief information about its associated VirtualMachineInstance | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachine

`func NewKubevirtIoApiCoreV1VirtualMachine(spec KubevirtIoApiCoreV1VirtualMachineSpec, ) *KubevirtIoApiCoreV1VirtualMachine`

NewKubevirtIoApiCoreV1VirtualMachine instantiates a new KubevirtIoApiCoreV1VirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineWithDefaults() *KubevirtIoApiCoreV1VirtualMachine`

NewKubevirtIoApiCoreV1VirtualMachineWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachine) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachine) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachine) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubevirtIoApiCoreV1VirtualMachine) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachine) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachine) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetSpec() KubevirtIoApiCoreV1VirtualMachineSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachine) SetSpec(v KubevirtIoApiCoreV1VirtualMachineSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetStatus() KubevirtIoApiCoreV1VirtualMachineStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachine) GetStatusOk() (*KubevirtIoApiCoreV1VirtualMachineStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachine) SetStatus(v KubevirtIoApiCoreV1VirtualMachineStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubevirtIoApiCoreV1VirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


