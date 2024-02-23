# K8sIoV1ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claims** | Pointer to [**[]K8sIoV1ResourceClaim**](K8sIoV1ResourceClaim.md) |  | [optional] 
**Limits** | Pointer to **map[string]string** |  | [optional] 
**Requests** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewK8sIoV1ResourceRequirements

`func NewK8sIoV1ResourceRequirements() *K8sIoV1ResourceRequirements`

NewK8sIoV1ResourceRequirements instantiates a new K8sIoV1ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ResourceRequirementsWithDefaults

`func NewK8sIoV1ResourceRequirementsWithDefaults() *K8sIoV1ResourceRequirements`

NewK8sIoV1ResourceRequirementsWithDefaults instantiates a new K8sIoV1ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaims

`func (o *K8sIoV1ResourceRequirements) GetClaims() []K8sIoV1ResourceClaim`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *K8sIoV1ResourceRequirements) GetClaimsOk() (*[]K8sIoV1ResourceClaim, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *K8sIoV1ResourceRequirements) SetClaims(v []K8sIoV1ResourceClaim)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *K8sIoV1ResourceRequirements) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### GetLimits

`func (o *K8sIoV1ResourceRequirements) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *K8sIoV1ResourceRequirements) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *K8sIoV1ResourceRequirements) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *K8sIoV1ResourceRequirements) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequests

`func (o *K8sIoV1ResourceRequirements) GetRequests() map[string]string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *K8sIoV1ResourceRequirements) GetRequestsOk() (*map[string]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *K8sIoV1ResourceRequirements) SetRequests(v map[string]string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *K8sIoV1ResourceRequirements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


