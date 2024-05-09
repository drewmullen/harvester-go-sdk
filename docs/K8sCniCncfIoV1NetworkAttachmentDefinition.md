# K8sCniCncfIoV1NetworkAttachmentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) |  | [optional] [default to {}]
**Spec** | [**K8sCniCncfIoV1NetworkAttachmentDefinitionSpec**](K8sCniCncfIoV1NetworkAttachmentDefinitionSpec.md) |  | [default to {}]

## Methods

### NewK8sCniCncfIoV1NetworkAttachmentDefinition

`func NewK8sCniCncfIoV1NetworkAttachmentDefinition(spec K8sCniCncfIoV1NetworkAttachmentDefinitionSpec, ) *K8sCniCncfIoV1NetworkAttachmentDefinition`

NewK8sCniCncfIoV1NetworkAttachmentDefinition instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sCniCncfIoV1NetworkAttachmentDefinitionWithDefaults

`func NewK8sCniCncfIoV1NetworkAttachmentDefinitionWithDefaults() *K8sCniCncfIoV1NetworkAttachmentDefinition`

NewK8sCniCncfIoV1NetworkAttachmentDefinitionWithDefaults instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetSpec() K8sCniCncfIoV1NetworkAttachmentDefinitionSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) GetSpecOk() (*K8sCniCncfIoV1NetworkAttachmentDefinitionSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinition) SetSpec(v K8sCniCncfIoV1NetworkAttachmentDefinitionSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


