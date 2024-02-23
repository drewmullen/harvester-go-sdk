# K8sIoV1PersistentVolumeClaimVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimName** | **string** |  | [default to ""]
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewK8sIoV1PersistentVolumeClaimVolumeSource

`func NewK8sIoV1PersistentVolumeClaimVolumeSource(claimName string, ) *K8sIoV1PersistentVolumeClaimVolumeSource`

NewK8sIoV1PersistentVolumeClaimVolumeSource instantiates a new K8sIoV1PersistentVolumeClaimVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimVolumeSourceWithDefaults

`func NewK8sIoV1PersistentVolumeClaimVolumeSourceWithDefaults() *K8sIoV1PersistentVolumeClaimVolumeSource`

NewK8sIoV1PersistentVolumeClaimVolumeSourceWithDefaults instantiates a new K8sIoV1PersistentVolumeClaimVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimName

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) GetClaimName() string`

GetClaimName returns the ClaimName field if non-nil, zero value otherwise.

### GetClaimNameOk

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) GetClaimNameOk() (*string, bool)`

GetClaimNameOk returns a tuple with the ClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimName

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) SetClaimName(v string)`

SetClaimName sets ClaimName field to given value.


### GetReadOnly

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *K8sIoV1PersistentVolumeClaimVolumeSource) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


