# K8sIoV1StatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Causes** | Pointer to [**[]K8sIoV1StatusCause**](K8sIoV1StatusCause.md) |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RetryAfterSeconds** | Pointer to **int32** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1StatusDetails

`func NewK8sIoV1StatusDetails() *K8sIoV1StatusDetails`

NewK8sIoV1StatusDetails instantiates a new K8sIoV1StatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1StatusDetailsWithDefaults

`func NewK8sIoV1StatusDetailsWithDefaults() *K8sIoV1StatusDetails`

NewK8sIoV1StatusDetailsWithDefaults instantiates a new K8sIoV1StatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCauses

`func (o *K8sIoV1StatusDetails) GetCauses() []K8sIoV1StatusCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *K8sIoV1StatusDetails) GetCausesOk() (*[]K8sIoV1StatusCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *K8sIoV1StatusDetails) SetCauses(v []K8sIoV1StatusCause)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *K8sIoV1StatusDetails) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetGroup

`func (o *K8sIoV1StatusDetails) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *K8sIoV1StatusDetails) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *K8sIoV1StatusDetails) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *K8sIoV1StatusDetails) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1StatusDetails) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1StatusDetails) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1StatusDetails) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *K8sIoV1StatusDetails) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *K8sIoV1StatusDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1StatusDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1StatusDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *K8sIoV1StatusDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRetryAfterSeconds

`func (o *K8sIoV1StatusDetails) GetRetryAfterSeconds() int32`

GetRetryAfterSeconds returns the RetryAfterSeconds field if non-nil, zero value otherwise.

### GetRetryAfterSecondsOk

`func (o *K8sIoV1StatusDetails) GetRetryAfterSecondsOk() (*int32, bool)`

GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSeconds

`func (o *K8sIoV1StatusDetails) SetRetryAfterSeconds(v int32)`

SetRetryAfterSeconds sets RetryAfterSeconds field to given value.

### HasRetryAfterSeconds

`func (o *K8sIoV1StatusDetails) HasRetryAfterSeconds() bool`

HasRetryAfterSeconds returns a boolean if a field has been set.

### GetUid

`func (o *K8sIoV1StatusDetails) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *K8sIoV1StatusDetails) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *K8sIoV1StatusDetails) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *K8sIoV1StatusDetails) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


