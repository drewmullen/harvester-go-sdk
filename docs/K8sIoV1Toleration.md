# K8sIoV1Toleration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**TolerationSeconds** | Pointer to **int64** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1Toleration

`func NewK8sIoV1Toleration() *K8sIoV1Toleration`

NewK8sIoV1Toleration instantiates a new K8sIoV1Toleration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1TolerationWithDefaults

`func NewK8sIoV1TolerationWithDefaults() *K8sIoV1Toleration`

NewK8sIoV1TolerationWithDefaults instantiates a new K8sIoV1Toleration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *K8sIoV1Toleration) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *K8sIoV1Toleration) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *K8sIoV1Toleration) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *K8sIoV1Toleration) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *K8sIoV1Toleration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *K8sIoV1Toleration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *K8sIoV1Toleration) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *K8sIoV1Toleration) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOperator

`func (o *K8sIoV1Toleration) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *K8sIoV1Toleration) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *K8sIoV1Toleration) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *K8sIoV1Toleration) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetTolerationSeconds

`func (o *K8sIoV1Toleration) GetTolerationSeconds() int64`

GetTolerationSeconds returns the TolerationSeconds field if non-nil, zero value otherwise.

### GetTolerationSecondsOk

`func (o *K8sIoV1Toleration) GetTolerationSecondsOk() (*int64, bool)`

GetTolerationSecondsOk returns a tuple with the TolerationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTolerationSeconds

`func (o *K8sIoV1Toleration) SetTolerationSeconds(v int64)`

SetTolerationSeconds sets TolerationSeconds field to given value.

### HasTolerationSeconds

`func (o *K8sIoV1Toleration) HasTolerationSeconds() bool`

HasTolerationSeconds returns a boolean if a field has been set.

### GetValue

`func (o *K8sIoV1Toleration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *K8sIoV1Toleration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *K8sIoV1Toleration) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *K8sIoV1Toleration) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


