# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **string** | Current is the identifier of the snapshot created for this checkpoint. | [default to ""]
**Previous** | **string** | Previous is the identifier of the snapshot from the previous checkpoint. | [default to ""]

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint(current string, previous string, ) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpointWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpointWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpointWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) GetCurrent() string`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) GetCurrentOk() (*string, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) SetCurrent(v string)`

SetCurrent sets Current field to given value.


### GetPrevious

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeCheckpoint) SetPrevious(v string)`

SetPrevious sets Previous field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


