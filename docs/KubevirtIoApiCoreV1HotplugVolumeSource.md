# KubevirtIoApiCoreV1HotplugVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVolume** | Pointer to [**KubevirtIoApiCoreV1DataVolumeSource**](KubevirtIoApiCoreV1DataVolumeSource.md) |  | [optional] 
**PersistentVolumeClaim** | Pointer to [**KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource**](KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1HotplugVolumeSource

`func NewKubevirtIoApiCoreV1HotplugVolumeSource() *KubevirtIoApiCoreV1HotplugVolumeSource`

NewKubevirtIoApiCoreV1HotplugVolumeSource instantiates a new KubevirtIoApiCoreV1HotplugVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1HotplugVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1HotplugVolumeSourceWithDefaults() *KubevirtIoApiCoreV1HotplugVolumeSource`

NewKubevirtIoApiCoreV1HotplugVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1HotplugVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVolume

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) GetDataVolume() KubevirtIoApiCoreV1DataVolumeSource`

GetDataVolume returns the DataVolume field if non-nil, zero value otherwise.

### GetDataVolumeOk

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) GetDataVolumeOk() (*KubevirtIoApiCoreV1DataVolumeSource, bool)`

GetDataVolumeOk returns a tuple with the DataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVolume

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) SetDataVolume(v KubevirtIoApiCoreV1DataVolumeSource)`

SetDataVolume sets DataVolume field to given value.

### HasDataVolume

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) HasDataVolume() bool`

HasDataVolume returns a boolean if a field has been set.

### GetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) GetPersistentVolumeClaim() KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) GetPersistentVolumeClaimOk() (*KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) SetPersistentVolumeClaim(v KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1HotplugVolumeSource) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


