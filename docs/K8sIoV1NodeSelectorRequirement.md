# K8sIoV1NodeSelectorRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | [default to ""]
**Operator** | **string** |  | [default to ""]
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewK8sIoV1NodeSelectorRequirement

`func NewK8sIoV1NodeSelectorRequirement(key string, operator string, ) *K8sIoV1NodeSelectorRequirement`

NewK8sIoV1NodeSelectorRequirement instantiates a new K8sIoV1NodeSelectorRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1NodeSelectorRequirementWithDefaults

`func NewK8sIoV1NodeSelectorRequirementWithDefaults() *K8sIoV1NodeSelectorRequirement`

NewK8sIoV1NodeSelectorRequirementWithDefaults instantiates a new K8sIoV1NodeSelectorRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *K8sIoV1NodeSelectorRequirement) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *K8sIoV1NodeSelectorRequirement) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *K8sIoV1NodeSelectorRequirement) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *K8sIoV1NodeSelectorRequirement) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *K8sIoV1NodeSelectorRequirement) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *K8sIoV1NodeSelectorRequirement) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *K8sIoV1NodeSelectorRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *K8sIoV1NodeSelectorRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *K8sIoV1NodeSelectorRequirement) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *K8sIoV1NodeSelectorRequirement) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


