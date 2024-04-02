# KubevirtIoApiCoreV1VirtualMachineInstanceMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Kind** | **string** |  | 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec.md) |  | [default to {}]
**Status** | Pointer to [**KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus**](KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus.md) |  | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigration

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigration(apiVersion string, kind string, spec KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec, ) *KubevirtIoApiCoreV1VirtualMachineInstanceMigration`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigration instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationWithDefaults() *KubevirtIoApiCoreV1VirtualMachineInstanceMigration`

NewKubevirtIoApiCoreV1VirtualMachineInstanceMigrationWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineInstanceMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetSpec() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetSpecOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetSpec(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationSpec)`

SetSpec sets Spec field to given value.


### GetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetStatus() KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) GetStatusOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) SetStatus(v KubevirtIoApiCoreV1VirtualMachineInstanceMigrationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KubevirtIoApiCoreV1VirtualMachineInstanceMigration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


