# KubevirtIoApiCoreV1RemoveVolumeOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **[]string** | When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed | [optional] 
**Name** | **string** | Name represents the name that maps to both the disk and volume that should be removed | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1RemoveVolumeOptions

`func NewKubevirtIoApiCoreV1RemoveVolumeOptions(name string, ) *KubevirtIoApiCoreV1RemoveVolumeOptions`

NewKubevirtIoApiCoreV1RemoveVolumeOptions instantiates a new KubevirtIoApiCoreV1RemoveVolumeOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1RemoveVolumeOptionsWithDefaults

`func NewKubevirtIoApiCoreV1RemoveVolumeOptionsWithDefaults() *KubevirtIoApiCoreV1RemoveVolumeOptions`

NewKubevirtIoApiCoreV1RemoveVolumeOptionsWithDefaults instantiates a new KubevirtIoApiCoreV1RemoveVolumeOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetDryRun() []string`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetDryRunOk() (*[]string, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) SetDryRun(v []string)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1RemoveVolumeOptions) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


