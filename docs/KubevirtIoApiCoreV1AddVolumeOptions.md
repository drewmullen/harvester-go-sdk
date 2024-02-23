# KubevirtIoApiCoreV1AddVolumeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | [**KubevirtIoApiCoreV1Disk**](KubevirtIoApiCoreV1Disk.md) |  | 
**DryRun** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | [default to ""]
**VolumeSource** | [**KubevirtIoApiCoreV1HotplugVolumeSource**](KubevirtIoApiCoreV1HotplugVolumeSource.md) |  | 

## Methods

### NewKubevirtIoApiCoreV1AddVolumeOptions

`func NewKubevirtIoApiCoreV1AddVolumeOptions(disk KubevirtIoApiCoreV1Disk, name string, volumeSource KubevirtIoApiCoreV1HotplugVolumeSource, ) *KubevirtIoApiCoreV1AddVolumeOptions`

NewKubevirtIoApiCoreV1AddVolumeOptions instantiates a new KubevirtIoApiCoreV1AddVolumeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1AddVolumeOptionsWithDefaults

`func NewKubevirtIoApiCoreV1AddVolumeOptionsWithDefaults() *KubevirtIoApiCoreV1AddVolumeOptions`

NewKubevirtIoApiCoreV1AddVolumeOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1AddVolumeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetDisk() KubevirtIoApiCoreV1Disk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetDiskOk() (*KubevirtIoApiCoreV1Disk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) SetDisk(v KubevirtIoApiCoreV1Disk)`

SetDisk sets Disk field to given value.


### GetDryRun

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetDryRun() []string`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetDryRunOk() (*[]string, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) SetDryRun(v []string)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) SetName(v string)`

SetName sets Name field to given value.


### GetVolumeSource

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetVolumeSource() KubevirtIoApiCoreV1HotplugVolumeSource`

GetVolumeSource returns the VolumeSource field if non-nil, zero value otherwise.

### GetVolumeSourceOk

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) GetVolumeSourceOk() (*KubevirtIoApiCoreV1HotplugVolumeSource, bool)`

GetVolumeSourceOk returns a tuple with the VolumeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSource

`func (o *KubevirtIoApiCoreV1AddVolumeOptions) SetVolumeSource(v KubevirtIoApiCoreV1HotplugVolumeSource)`

SetVolumeSource sets VolumeSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


