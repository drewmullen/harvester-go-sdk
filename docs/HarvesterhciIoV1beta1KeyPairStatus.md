# HarvesterhciIoV1beta1KeyPairStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**FingerPrint** | Pointer to **string** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1KeyPairStatus

`func NewHarvesterhciIoV1beta1KeyPairStatus() *HarvesterhciIoV1beta1KeyPairStatus`

NewHarvesterhciIoV1beta1KeyPairStatus instantiates a new HarvesterhciIoV1beta1KeyPairStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1KeyPairStatusWithDefaults

`func NewHarvesterhciIoV1beta1KeyPairStatusWithDefaults() *HarvesterhciIoV1beta1KeyPairStatus`

NewHarvesterhciIoV1beta1KeyPairStatusWithDefaults instantiates a new HarvesterhciIoV1beta1KeyPairStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *HarvesterhciIoV1beta1KeyPairStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1KeyPairStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1KeyPairStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1KeyPairStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetFingerPrint

`func (o *HarvesterhciIoV1beta1KeyPairStatus) GetFingerPrint() string`

GetFingerPrint returns the FingerPrint field if non-nil, zero value otherwise.

### GetFingerPrintOk

`func (o *HarvesterhciIoV1beta1KeyPairStatus) GetFingerPrintOk() (*string, bool)`

GetFingerPrintOk returns a tuple with the FingerPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerPrint

`func (o *HarvesterhciIoV1beta1KeyPairStatus) SetFingerPrint(v string)`

SetFingerPrint sets FingerPrint field to given value.

### HasFingerPrint

`func (o *HarvesterhciIoV1beta1KeyPairStatus) HasFingerPrint() bool`

HasFingerPrint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


