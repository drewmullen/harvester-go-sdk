# KubevirtIoApiCoreV1DataVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hotpluggable** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | [default to ""]

## Methods

### NewKubevirtIoApiCoreV1DataVolumeSource

`func NewKubevirtIoApiCoreV1DataVolumeSource(name string, ) *KubevirtIoApiCoreV1DataVolumeSource`

NewKubevirtIoApiCoreV1DataVolumeSource instantiates a new KubevirtIoApiCoreV1DataVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DataVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1DataVolumeSourceWithDefaults() *KubevirtIoApiCoreV1DataVolumeSource`

NewKubevirtIoApiCoreV1DataVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1DataVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHotpluggable

`func (o *KubevirtIoApiCoreV1DataVolumeSource) GetHotpluggable() bool`

GetHotpluggable returns the Hotpluggable field if non-nil, zero value otherwise.

### GetHotpluggableOk

`func (o *KubevirtIoApiCoreV1DataVolumeSource) GetHotpluggableOk() (*bool, bool)`

GetHotpluggableOk returns a tuple with the Hotpluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotpluggable

`func (o *KubevirtIoApiCoreV1DataVolumeSource) SetHotpluggable(v bool)`

SetHotpluggable sets Hotpluggable field to given value.

### HasHotpluggable

`func (o *KubevirtIoApiCoreV1DataVolumeSource) HasHotpluggable() bool`

HasHotpluggable returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1DataVolumeSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1DataVolumeSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1DataVolumeSource) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


