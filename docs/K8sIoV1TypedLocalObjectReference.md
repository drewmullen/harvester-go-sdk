# K8sIoV1TypedLocalObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | [default to ""]
**Name** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1TypedLocalObjectReference

`func NewK8sIoV1TypedLocalObjectReference(kind string, name string, ) *K8sIoV1TypedLocalObjectReference`

NewK8sIoV1TypedLocalObjectReference instantiates a new K8sIoV1TypedLocalObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1TypedLocalObjectReferenceWithDefaults

`func NewK8sIoV1TypedLocalObjectReferenceWithDefaults() *K8sIoV1TypedLocalObjectReference`

NewK8sIoV1TypedLocalObjectReferenceWithDefaults instantiates a new K8sIoV1TypedLocalObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *K8sIoV1TypedLocalObjectReference) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *K8sIoV1TypedLocalObjectReference) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *K8sIoV1TypedLocalObjectReference) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *K8sIoV1TypedLocalObjectReference) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1TypedLocalObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1TypedLocalObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1TypedLocalObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *K8sIoV1TypedLocalObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1TypedLocalObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1TypedLocalObjectReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


