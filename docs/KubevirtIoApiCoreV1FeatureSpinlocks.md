# KubevirtIoApiCoreV1FeatureSpinlocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled determines if the feature should be enabled or disabled on the guest. Defaults to true. | [optional] 
**Spinlocks** | Pointer to **int64** | Retries indicates the number of retries. Must be a value greater or equal 4096. Defaults to 4096. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1FeatureSpinlocks

`func NewKubevirtIoApiCoreV1FeatureSpinlocks() *KubevirtIoApiCoreV1FeatureSpinlocks`

NewKubevirtIoApiCoreV1FeatureSpinlocks instantiates a new KubevirtIoApiCoreV1FeatureSpinlocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1FeatureSpinlocksWithDefaults

`func NewKubevirtIoApiCoreV1FeatureSpinlocksWithDefaults() *KubevirtIoApiCoreV1FeatureSpinlocks`

NewKubevirtIoApiCoreV1FeatureSpinlocksWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureSpinlocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetSpinlocks() int64`

GetSpinlocks returns the Spinlocks field if non-nil, zero value otherwise.

### GetSpinlocksOk

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) GetSpinlocksOk() (*int64, bool)`

GetSpinlocksOk returns a tuple with the Spinlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) SetSpinlocks(v int64)`

SetSpinlocks sets Spinlocks field to given value.

### HasSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureSpinlocks) HasSpinlocks() bool`

HasSpinlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


