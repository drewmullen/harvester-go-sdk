# KubevirtIoApiCoreV1BlockSize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Custom** | Pointer to [**KubevirtIoApiCoreV1CustomBlockSize**](KubevirtIoApiCoreV1CustomBlockSize.md) |  | [optional] 
**MatchVolume** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1BlockSize

`func NewKubevirtIoApiCoreV1BlockSize() *KubevirtIoApiCoreV1BlockSize`

NewKubevirtIoApiCoreV1BlockSize instantiates a new KubevirtIoApiCoreV1BlockSize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1BlockSizeWithDefaults

`func NewKubevirtIoApiCoreV1BlockSizeWithDefaults() *KubevirtIoApiCoreV1BlockSize`

NewKubevirtIoApiCoreV1BlockSizeWithDefaults instantiates a new KubevirtIoApiCoreV1BlockSize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustom

`func (o *KubevirtIoApiCoreV1BlockSize) GetCustom() KubevirtIoApiCoreV1CustomBlockSize`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *KubevirtIoApiCoreV1BlockSize) GetCustomOk() (*KubevirtIoApiCoreV1CustomBlockSize, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *KubevirtIoApiCoreV1BlockSize) SetCustom(v KubevirtIoApiCoreV1CustomBlockSize)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *KubevirtIoApiCoreV1BlockSize) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetMatchVolume

`func (o *KubevirtIoApiCoreV1BlockSize) GetMatchVolume() KubevirtIoApiCoreV1FeatureState`

GetMatchVolume returns the MatchVolume field if non-nil, zero value otherwise.

### GetMatchVolumeOk

`func (o *KubevirtIoApiCoreV1BlockSize) GetMatchVolumeOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetMatchVolumeOk returns a tuple with the MatchVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchVolume

`func (o *KubevirtIoApiCoreV1BlockSize) SetMatchVolume(v KubevirtIoApiCoreV1FeatureState)`

SetMatchVolume sets MatchVolume field to given value.

### HasMatchVolume

`func (o *KubevirtIoApiCoreV1BlockSize) HasMatchVolume() bool`

HasMatchVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


