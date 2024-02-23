# K8sIoV1TypedObjectReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroup** | Pointer to **string** |  | [optional] 
**Kind** | **string** |  | [default to ""]
**Name** | **string** |  | [default to ""]
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1TypedObjectReference

`func NewK8sIoV1TypedObjectReference(kind string, name string, ) *K8sIoV1TypedObjectReference`

NewK8sIoV1TypedObjectReference instantiates a new K8sIoV1TypedObjectReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1TypedObjectReferenceWithDefaults

`func NewK8sIoV1TypedObjectReferenceWithDefaults() *K8sIoV1TypedObjectReference`

NewK8sIoV1TypedObjectReferenceWithDefaults instantiates a new K8sIoV1TypedObjectReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroup

`func (o *K8sIoV1TypedObjectReference) GetApiGroup() string`

GetApiGroup returns the ApiGroup field if non-nil, zero value otherwise.

### GetApiGroupOk

`func (o *K8sIoV1TypedObjectReference) GetApiGroupOk() (*string, bool)`

GetApiGroupOk returns a tuple with the ApiGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroup

`func (o *K8sIoV1TypedObjectReference) SetApiGroup(v string)`

SetApiGroup sets ApiGroup field to given value.

### HasApiGroup

`func (o *K8sIoV1TypedObjectReference) HasApiGroup() bool`

HasApiGroup returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1TypedObjectReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1TypedObjectReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1TypedObjectReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *K8sIoV1TypedObjectReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1TypedObjectReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1TypedObjectReference) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *K8sIoV1TypedObjectReference) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *K8sIoV1TypedObjectReference) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *K8sIoV1TypedObjectReference) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *K8sIoV1TypedObjectReference) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


