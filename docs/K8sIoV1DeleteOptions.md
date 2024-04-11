# K8sIoV1DeleteOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**DryRun** | Pointer to **[]string** |  | [optional] 
**GracePeriodSeconds** | Pointer to **int64** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**OrphanDependents** | Pointer to **bool** |  | [optional] 
**Preconditions** | Pointer to [**K8sIoV1Preconditions**](K8sIoV1Preconditions.md) |  | [optional] 
**PropagationPolicy** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1DeleteOptions

`func NewK8sIoV1DeleteOptions() *K8sIoV1DeleteOptions`

NewK8sIoV1DeleteOptions instantiates a new K8sIoV1DeleteOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1DeleteOptionsWithDefaults

`func NewK8sIoV1DeleteOptionsWithDefaults() *K8sIoV1DeleteOptions`

NewK8sIoV1DeleteOptionsWithDefaults instantiates a new K8sIoV1DeleteOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1DeleteOptions) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1DeleteOptions) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1DeleteOptions) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *K8sIoV1DeleteOptions) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetDryRun

`func (o *K8sIoV1DeleteOptions) GetDryRun() []string`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *K8sIoV1DeleteOptions) GetDryRunOk() (*[]string, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *K8sIoV1DeleteOptions) SetDryRun(v []string)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *K8sIoV1DeleteOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetGracePeriodSeconds

`func (o *K8sIoV1DeleteOptions) GetGracePeriodSeconds() int64`

GetGracePeriodSeconds returns the GracePeriodSeconds field if non-nil, zero value otherwise.

### GetGracePeriodSecondsOk

`func (o *K8sIoV1DeleteOptions) GetGracePeriodSecondsOk() (*int64, bool)`

GetGracePeriodSecondsOk returns a tuple with the GracePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodSeconds

`func (o *K8sIoV1DeleteOptions) SetGracePeriodSeconds(v int64)`

SetGracePeriodSeconds sets GracePeriodSeconds field to given value.

### HasGracePeriodSeconds

`func (o *K8sIoV1DeleteOptions) HasGracePeriodSeconds() bool`

HasGracePeriodSeconds returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1DeleteOptions) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1DeleteOptions) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1DeleteOptions) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *K8sIoV1DeleteOptions) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetOrphanDependents

`func (o *K8sIoV1DeleteOptions) GetOrphanDependents() bool`

GetOrphanDependents returns the OrphanDependents field if non-nil, zero value otherwise.

### GetOrphanDependentsOk

`func (o *K8sIoV1DeleteOptions) GetOrphanDependentsOk() (*bool, bool)`

GetOrphanDependentsOk returns a tuple with the OrphanDependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrphanDependents

`func (o *K8sIoV1DeleteOptions) SetOrphanDependents(v bool)`

SetOrphanDependents sets OrphanDependents field to given value.

### HasOrphanDependents

`func (o *K8sIoV1DeleteOptions) HasOrphanDependents() bool`

HasOrphanDependents returns a boolean if a field has been set.

### GetPreconditions

`func (o *K8sIoV1DeleteOptions) GetPreconditions() K8sIoV1Preconditions`

GetPreconditions returns the Preconditions field if non-nil, zero value otherwise.

### GetPreconditionsOk

`func (o *K8sIoV1DeleteOptions) GetPreconditionsOk() (*K8sIoV1Preconditions, bool)`

GetPreconditionsOk returns a tuple with the Preconditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditions

`func (o *K8sIoV1DeleteOptions) SetPreconditions(v K8sIoV1Preconditions)`

SetPreconditions sets Preconditions field to given value.

### HasPreconditions

`func (o *K8sIoV1DeleteOptions) HasPreconditions() bool`

HasPreconditions returns a boolean if a field has been set.

### GetPropagationPolicy

`func (o *K8sIoV1DeleteOptions) GetPropagationPolicy() string`

GetPropagationPolicy returns the PropagationPolicy field if non-nil, zero value otherwise.

### GetPropagationPolicyOk

`func (o *K8sIoV1DeleteOptions) GetPropagationPolicyOk() (*string, bool)`

GetPropagationPolicyOk returns a tuple with the PropagationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationPolicy

`func (o *K8sIoV1DeleteOptions) SetPropagationPolicy(v string)`

SetPropagationPolicy sets PropagationPolicy field to given value.

### HasPropagationPolicy

`func (o *K8sIoV1DeleteOptions) HasPropagationPolicy() bool`

HasPropagationPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


