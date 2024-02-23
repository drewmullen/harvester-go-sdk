# K8sIoV1ResourceFieldSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | Pointer to **string** |  | [optional] 
**Divisor** | Pointer to **string** |  | [optional] 
**Resource** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1ResourceFieldSelector

`func NewK8sIoV1ResourceFieldSelector(resource string, ) *K8sIoV1ResourceFieldSelector`

NewK8sIoV1ResourceFieldSelector instantiates a new K8sIoV1ResourceFieldSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ResourceFieldSelectorWithDefaults

`func NewK8sIoV1ResourceFieldSelectorWithDefaults() *K8sIoV1ResourceFieldSelector`

NewK8sIoV1ResourceFieldSelectorWithDefaults instantiates a new K8sIoV1ResourceFieldSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerName

`func (o *K8sIoV1ResourceFieldSelector) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *K8sIoV1ResourceFieldSelector) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *K8sIoV1ResourceFieldSelector) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *K8sIoV1ResourceFieldSelector) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDivisor

`func (o *K8sIoV1ResourceFieldSelector) GetDivisor() string`

GetDivisor returns the Divisor field if non-nil, zero value otherwise.

### GetDivisorOk

`func (o *K8sIoV1ResourceFieldSelector) GetDivisorOk() (*string, bool)`

GetDivisorOk returns a tuple with the Divisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisor

`func (o *K8sIoV1ResourceFieldSelector) SetDivisor(v string)`

SetDivisor sets Divisor field to given value.

### HasDivisor

`func (o *K8sIoV1ResourceFieldSelector) HasDivisor() bool`

HasDivisor returns a boolean if a field has been set.

### GetResource

`func (o *K8sIoV1ResourceFieldSelector) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *K8sIoV1ResourceFieldSelector) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *K8sIoV1ResourceFieldSelector) SetResource(v string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


