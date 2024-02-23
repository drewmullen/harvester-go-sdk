# KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | **string** |  | [default to ""]
**Hotpluggable** | Pointer to **bool** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource

`func NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource(claimName string, ) *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource`

NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSourceWithDefaults

`func NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSourceWithDefaults() *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource`

NewKubevirtIoApiCoreV1PersistentVolumeClaimVolumeSourceWithDefaults instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetHotpluggable

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetHotpluggable() bool`

GetHotpluggable returns the Hotpluggable field if non-nil, zero value otherwise.

### GetHotpluggableOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetHotpluggableOk() (*bool, bool)`

GetHotpluggableOk returns a tuple with the Hotpluggable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotpluggable

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) SetHotpluggable(v bool)`

SetHotpluggable sets Hotpluggable field to given value.

### HasHotpluggable

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) HasHotpluggable() bool`

HasHotpluggable returns a boolean if a field has been set.

### GetReadOnly

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


