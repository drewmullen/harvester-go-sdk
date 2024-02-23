# KubevirtIoApiCoreV1MemoryDumpVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | **string** |  | [default to ""]
**Hotpluggable** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1MemoryDumpVolumeSource

`func NewKubevirtIoApiCoreV1MemoryDumpVolumeSource(claimName string, ) *KubevirtIoApiCoreV1MemoryDumpVolumeSource`

NewKubevirtIoApiCoreV1MemoryDumpVolumeSource instantiates a new KubevirtIoApiCoreV1MemoryDumpVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1MemoryDumpVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1MemoryDumpVolumeSourceWithDefaults() *KubevirtIoApiCoreV1MemoryDumpVolumeSource`

NewKubevirtIoApiCoreV1MemoryDumpVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1MemoryDumpVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetHotpluggable

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetHotpluggable() bool`

GetHotpluggable returns the Hotpluggable field if non-nil, zero value otherwise.

### GetHotpluggableOk

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetHotpluggableOk() (*bool, bool)`

GetHotpluggableOk returns a tuple with the Hotpluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotpluggable

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) SetHotpluggable(v bool)`

SetHotpluggable sets Hotpluggable field to given value.

### HasHotpluggable

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) HasHotpluggable() bool`

HasHotpluggable returns a boolean if a field has been set.

### GetReadOnly

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *KubevirtIoApiCoreV1MemoryDumpVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


