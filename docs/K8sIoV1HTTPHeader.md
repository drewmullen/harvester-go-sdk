# K8sIoV1HTTPHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header. | [default to ""]
**Value** | **string** | The header field value | [default to ""]

## Methods

### NewK8sIoV1HTTPHeader

`func NewK8sIoV1HTTPHeader(name string, value string, ) *K8sIoV1HTTPHeader`

NewK8sIoV1HTTPHeader instantiates a new K8sIoV1HTTPHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1HTTPHeaderWithDefaults

`func NewK8sIoV1HTTPHeaderWithDefaults() *K8sIoV1HTTPHeader`

NewK8sIoV1HTTPHeaderWithDefaults instantiates a new K8sIoV1HTTPHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *K8sIoV1HTTPHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1HTTPHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1HTTPHeader) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *K8sIoV1HTTPHeader) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *K8sIoV1HTTPHeader) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *K8sIoV1HTTPHeader) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


