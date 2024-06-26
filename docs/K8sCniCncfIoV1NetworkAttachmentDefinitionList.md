# K8sCniCncfIoV1NetworkAttachmentDefinitionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Items** | [**[]K8sCniCncfIoV1NetworkAttachmentDefinition**](K8sCniCncfIoV1NetworkAttachmentDefinition.md) |  | 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [default to {}]

## Methods

### NewK8sCniCncfIoV1NetworkAttachmentDefinitionList

`func NewK8sCniCncfIoV1NetworkAttachmentDefinitionList(items []K8sCniCncfIoV1NetworkAttachmentDefinition, metadata K8sIoV1ListMeta, ) *K8sCniCncfIoV1NetworkAttachmentDefinitionList`

NewK8sCniCncfIoV1NetworkAttachmentDefinitionList instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinitionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sCniCncfIoV1NetworkAttachmentDefinitionListWithDefaults

`func NewK8sCniCncfIoV1NetworkAttachmentDefinitionListWithDefaults() *K8sCniCncfIoV1NetworkAttachmentDefinitionList`

NewK8sCniCncfIoV1NetworkAttachmentDefinitionListWithDefaults instantiates a new K8sCniCncfIoV1NetworkAttachmentDefinitionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetItems() []K8sCniCncfIoV1NetworkAttachmentDefinition`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetItemsOk() (*[]K8sCniCncfIoV1NetworkAttachmentDefinition, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetItems(v []K8sCniCncfIoV1NetworkAttachmentDefinition)`

SetItems sets Items field to given value.


### GetKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *K8sCniCncfIoV1NetworkAttachmentDefinitionList) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


