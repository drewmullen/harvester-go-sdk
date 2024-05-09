# KubevirtIoApiCoreV1ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to **map[string]string** | Limits describes the maximum amount of compute resources allowed. Valid resource keys are \&quot;memory\&quot; and \&quot;cpu\&quot;. | [optional] 
**OvercommitGuestOverhead** | Pointer to **bool** | Don&#39;t ask the scheduler to take the guest-management overhead into account. Instead put the overhead only into the container&#39;s memory limit. This can lead to crashes if all memory is in use on a node. Defaults to false. | [optional] 
**Requests** | Pointer to **map[string]string** | Requests is a description of the initial vmi resources. Valid resource keys are \&quot;memory\&quot; and \&quot;cpu\&quot;. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1ResourceRequirements

`func NewKubevirtIoApiCoreV1ResourceRequirements() *KubevirtIoApiCoreV1ResourceRequirements`

NewKubevirtIoApiCoreV1ResourceRequirements instantiates a new KubevirtIoApiCoreV1ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ResourceRequirementsWithDefaults

`func NewKubevirtIoApiCoreV1ResourceRequirementsWithDefaults() *KubevirtIoApiCoreV1ResourceRequirements`

NewKubevirtIoApiCoreV1ResourceRequirementsWithDefaults instantiates a new KubevirtIoApiCoreV1ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetLimits() map[string]string`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetLimitsOk() (*map[string]string, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KubevirtIoApiCoreV1ResourceRequirements) SetLimits(v map[string]string)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KubevirtIoApiCoreV1ResourceRequirements) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetOvercommitGuestOverhead

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetOvercommitGuestOverhead() bool`

GetOvercommitGuestOverhead returns the OvercommitGuestOverhead field if non-nil, zero value otherwise.

### GetOvercommitGuestOverheadOk

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetOvercommitGuestOverheadOk() (*bool, bool)`

GetOvercommitGuestOverheadOk returns a tuple with the OvercommitGuestOverhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvercommitGuestOverhead

`func (o *KubevirtIoApiCoreV1ResourceRequirements) SetOvercommitGuestOverhead(v bool)`

SetOvercommitGuestOverhead sets OvercommitGuestOverhead field to given value.

### HasOvercommitGuestOverhead

`func (o *KubevirtIoApiCoreV1ResourceRequirements) HasOvercommitGuestOverhead() bool`

HasOvercommitGuestOverhead returns a boolean if a field has been set.

### GetRequests

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetRequests() map[string]string`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *KubevirtIoApiCoreV1ResourceRequirements) GetRequestsOk() (*map[string]string, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *KubevirtIoApiCoreV1ResourceRequirements) SetRequests(v map[string]string)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *KubevirtIoApiCoreV1ResourceRequirements) HasRequests() bool`

HasRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


