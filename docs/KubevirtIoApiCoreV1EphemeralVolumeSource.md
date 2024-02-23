# KubevirtIoApiCoreV1EphemeralVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersistentVolumeClaim** | Pointer to [**K8sIoV1PersistentVolumeClaimVolumeSource**](K8sIoV1PersistentVolumeClaimVolumeSource.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1EphemeralVolumeSource

`func NewKubevirtIoApiCoreV1EphemeralVolumeSource() *KubevirtIoApiCoreV1EphemeralVolumeSource`

NewKubevirtIoApiCoreV1EphemeralVolumeSource instantiates a new KubevirtIoApiCoreV1EphemeralVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1EphemeralVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1EphemeralVolumeSourceWithDefaults() *KubevirtIoApiCoreV1EphemeralVolumeSource`

NewKubevirtIoApiCoreV1EphemeralVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1EphemeralVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1EphemeralVolumeSource) GetPersistentVolumeClaim() K8sIoV1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *KubevirtIoApiCoreV1EphemeralVolumeSource) GetPersistentVolumeClaimOk() (*K8sIoV1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1EphemeralVolumeSource) SetPersistentVolumeClaim(v K8sIoV1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1EphemeralVolumeSource) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


