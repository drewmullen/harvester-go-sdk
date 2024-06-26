# K8sIoV1PersistentVolumeClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**K8sIoV1ObjectMeta**](K8sIoV1ObjectMeta.md) | Standard object&#39;s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata | [optional] [default to {}]
**Spec** | Pointer to [**K8sIoV1PersistentVolumeClaimSpec**](K8sIoV1PersistentVolumeClaimSpec.md) | spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | [optional] [default to {}]
**Status** | Pointer to [**K8sIoV1PersistentVolumeClaimStatus**](K8sIoV1PersistentVolumeClaimStatus.md) | status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | [optional] [default to {}]

## Methods

### NewK8sIoV1PersistentVolumeClaim

`func NewK8sIoV1PersistentVolumeClaim() *K8sIoV1PersistentVolumeClaim`

NewK8sIoV1PersistentVolumeClaim instantiates a new K8sIoV1PersistentVolumeClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimWithDefaults

`func NewK8sIoV1PersistentVolumeClaimWithDefaults() *K8sIoV1PersistentVolumeClaim`

NewK8sIoV1PersistentVolumeClaimWithDefaults instantiates a new K8sIoV1PersistentVolumeClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1PersistentVolumeClaim) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1PersistentVolumeClaim) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1PersistentVolumeClaim) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *K8sIoV1PersistentVolumeClaim) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1PersistentVolumeClaim) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1PersistentVolumeClaim) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1PersistentVolumeClaim) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *K8sIoV1PersistentVolumeClaim) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *K8sIoV1PersistentVolumeClaim) GetMetadata() K8sIoV1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *K8sIoV1PersistentVolumeClaim) GetMetadataOk() (*K8sIoV1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *K8sIoV1PersistentVolumeClaim) SetMetadata(v K8sIoV1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *K8sIoV1PersistentVolumeClaim) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *K8sIoV1PersistentVolumeClaim) GetSpec() K8sIoV1PersistentVolumeClaimSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *K8sIoV1PersistentVolumeClaim) GetSpecOk() (*K8sIoV1PersistentVolumeClaimSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *K8sIoV1PersistentVolumeClaim) SetSpec(v K8sIoV1PersistentVolumeClaimSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *K8sIoV1PersistentVolumeClaim) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetStatus

`func (o *K8sIoV1PersistentVolumeClaim) GetStatus() K8sIoV1PersistentVolumeClaimStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *K8sIoV1PersistentVolumeClaim) GetStatusOk() (*K8sIoV1PersistentVolumeClaimStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *K8sIoV1PersistentVolumeClaim) SetStatus(v K8sIoV1PersistentVolumeClaimStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *K8sIoV1PersistentVolumeClaim) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


