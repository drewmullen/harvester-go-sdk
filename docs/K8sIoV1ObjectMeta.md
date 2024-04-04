# K8sIoV1ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | Pointer to **map[string]string** |  | [optional] 
**CreationTimestamp** | Pointer to **string** |  | [optional] [default to "{}"]
**DeletionGracePeriodSeconds** | Pointer to **int64** |  | [optional] 
**DeletionTimestamp** | Pointer to **string** |  | [optional] [default to ""]
**Finalizers** | Pointer to **[]string** |  | [optional] 
**GenerateName** | Pointer to **string** |  | [optional] 
**Generation** | Pointer to **int64** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**ManagedFields** | Pointer to [**[]K8sIoV1ManagedFieldsEntry**](K8sIoV1ManagedFieldsEntry.md) |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**OwnerReferences** | Pointer to [**[]K8sIoV1OwnerReference**](K8sIoV1OwnerReference.md) |  | [optional] 
**ResourceVersion** | Pointer to **string** |  | [optional] 
**SelfLink** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1ObjectMeta

`func NewK8sIoV1ObjectMeta(name string, ) *K8sIoV1ObjectMeta`

NewK8sIoV1ObjectMeta instantiates a new K8sIoV1ObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ObjectMetaWithDefaults

`func NewK8sIoV1ObjectMetaWithDefaults() *K8sIoV1ObjectMeta`

NewK8sIoV1ObjectMetaWithDefaults instantiates a new K8sIoV1ObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotations

`func (o *K8sIoV1ObjectMeta) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *K8sIoV1ObjectMeta) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *K8sIoV1ObjectMeta) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *K8sIoV1ObjectMeta) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *K8sIoV1ObjectMeta) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *K8sIoV1ObjectMeta) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *K8sIoV1ObjectMeta) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *K8sIoV1ObjectMeta) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetDeletionGracePeriodSeconds

`func (o *K8sIoV1ObjectMeta) GetDeletionGracePeriodSeconds() int64`

GetDeletionGracePeriodSeconds returns the DeletionGracePeriodSeconds field if non-nil, zero value otherwise.

### GetDeletionGracePeriodSecondsOk

`func (o *K8sIoV1ObjectMeta) GetDeletionGracePeriodSecondsOk() (*int64, bool)`

GetDeletionGracePeriodSecondsOk returns a tuple with the DeletionGracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionGracePeriodSeconds

`func (o *K8sIoV1ObjectMeta) SetDeletionGracePeriodSeconds(v int64)`

SetDeletionGracePeriodSeconds sets DeletionGracePeriodSeconds field to given value.

### HasDeletionGracePeriodSeconds

`func (o *K8sIoV1ObjectMeta) HasDeletionGracePeriodSeconds() bool`

HasDeletionGracePeriodSeconds returns a boolean if a field has been set.

### GetDeletionTimestamp

`func (o *K8sIoV1ObjectMeta) GetDeletionTimestamp() string`

GetDeletionTimestamp returns the DeletionTimestamp field if non-nil, zero value otherwise.

### GetDeletionTimestampOk

`func (o *K8sIoV1ObjectMeta) GetDeletionTimestampOk() (*string, bool)`

GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTimestamp

`func (o *K8sIoV1ObjectMeta) SetDeletionTimestamp(v string)`

SetDeletionTimestamp sets DeletionTimestamp field to given value.

### HasDeletionTimestamp

`func (o *K8sIoV1ObjectMeta) HasDeletionTimestamp() bool`

HasDeletionTimestamp returns a boolean if a field has been set.

### GetFinalizers

`func (o *K8sIoV1ObjectMeta) GetFinalizers() []string`

GetFinalizers returns the Finalizers field if non-nil, zero value otherwise.

### GetFinalizersOk

`func (o *K8sIoV1ObjectMeta) GetFinalizersOk() (*[]string, bool)`

GetFinalizersOk returns a tuple with the Finalizers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizers

`func (o *K8sIoV1ObjectMeta) SetFinalizers(v []string)`

SetFinalizers sets Finalizers field to given value.

### HasFinalizers

`func (o *K8sIoV1ObjectMeta) HasFinalizers() bool`

HasFinalizers returns a boolean if a field has been set.

### GetGenerateName

`func (o *K8sIoV1ObjectMeta) GetGenerateName() string`

GetGenerateName returns the GenerateName field if non-nil, zero value otherwise.

### GetGenerateNameOk

`func (o *K8sIoV1ObjectMeta) GetGenerateNameOk() (*string, bool)`

GetGenerateNameOk returns a tuple with the GenerateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateName

`func (o *K8sIoV1ObjectMeta) SetGenerateName(v string)`

SetGenerateName sets GenerateName field to given value.

### HasGenerateName

`func (o *K8sIoV1ObjectMeta) HasGenerateName() bool`

HasGenerateName returns a boolean if a field has been set.

### GetGeneration

`func (o *K8sIoV1ObjectMeta) GetGeneration() int64`

GetGeneration returns the Generation field if non-nil, zero value otherwise.

### GetGenerationOk

`func (o *K8sIoV1ObjectMeta) GetGenerationOk() (*int64, bool)`

GetGenerationOk returns a tuple with the Generation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneration

`func (o *K8sIoV1ObjectMeta) SetGeneration(v int64)`

SetGeneration sets Generation field to given value.

### HasGeneration

`func (o *K8sIoV1ObjectMeta) HasGeneration() bool`

HasGeneration returns a boolean if a field has been set.

### GetLabels

`func (o *K8sIoV1ObjectMeta) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *K8sIoV1ObjectMeta) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *K8sIoV1ObjectMeta) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *K8sIoV1ObjectMeta) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetManagedFields

`func (o *K8sIoV1ObjectMeta) GetManagedFields() []K8sIoV1ManagedFieldsEntry`

GetManagedFields returns the ManagedFields field if non-nil, zero value otherwise.

### GetManagedFieldsOk

`func (o *K8sIoV1ObjectMeta) GetManagedFieldsOk() (*[]K8sIoV1ManagedFieldsEntry, bool)`

GetManagedFieldsOk returns a tuple with the ManagedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedFields

`func (o *K8sIoV1ObjectMeta) SetManagedFields(v []K8sIoV1ManagedFieldsEntry)`

SetManagedFields sets ManagedFields field to given value.

### HasManagedFields

`func (o *K8sIoV1ObjectMeta) HasManagedFields() bool`

HasManagedFields returns a boolean if a field has been set.

### GetName

`func (o *K8sIoV1ObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1ObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1ObjectMeta) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *K8sIoV1ObjectMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *K8sIoV1ObjectMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *K8sIoV1ObjectMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *K8sIoV1ObjectMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwnerReferences

`func (o *K8sIoV1ObjectMeta) GetOwnerReferences() []K8sIoV1OwnerReference`

GetOwnerReferences returns the OwnerReferences field if non-nil, zero value otherwise.

### GetOwnerReferencesOk

`func (o *K8sIoV1ObjectMeta) GetOwnerReferencesOk() (*[]K8sIoV1OwnerReference, bool)`

GetOwnerReferencesOk returns a tuple with the OwnerReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerReferences

`func (o *K8sIoV1ObjectMeta) SetOwnerReferences(v []K8sIoV1OwnerReference)`

SetOwnerReferences sets OwnerReferences field to given value.

### HasOwnerReferences

`func (o *K8sIoV1ObjectMeta) HasOwnerReferences() bool`

HasOwnerReferences returns a boolean if a field has been set.

### GetResourceVersion

`func (o *K8sIoV1ObjectMeta) GetResourceVersion() string`

GetResourceVersion returns the ResourceVersion field if non-nil, zero value otherwise.

### GetResourceVersionOk

`func (o *K8sIoV1ObjectMeta) GetResourceVersionOk() (*string, bool)`

GetResourceVersionOk returns a tuple with the ResourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersion

`func (o *K8sIoV1ObjectMeta) SetResourceVersion(v string)`

SetResourceVersion sets ResourceVersion field to given value.

### HasResourceVersion

`func (o *K8sIoV1ObjectMeta) HasResourceVersion() bool`

HasResourceVersion returns a boolean if a field has been set.

### GetSelfLink

`func (o *K8sIoV1ObjectMeta) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *K8sIoV1ObjectMeta) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *K8sIoV1ObjectMeta) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.

### HasSelfLink

`func (o *K8sIoV1ObjectMeta) HasSelfLink() bool`

HasSelfLink returns a boolean if a field has been set.

### GetUid

`func (o *K8sIoV1ObjectMeta) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *K8sIoV1ObjectMeta) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *K8sIoV1ObjectMeta) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *K8sIoV1ObjectMeta) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


