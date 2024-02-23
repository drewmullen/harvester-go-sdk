# K8sIoV1ObjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1ObjectMeta

`func NewK8sIoV1ObjectMeta(name string, ) *K8sIoV1ObjectMeta`

NewK8sIoV1ObjectMeta instantiates a new K8sIoV1ObjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ObjectMetaWithDefaults

`func NewK8sIoV1ObjectMetaWithDefaults() *K8sIoV1ObjectMeta`

NewK8sIoV1ObjectMetaWithDefaults instantiates a new K8sIoV1ObjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *K8sIoV1ObjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1ObjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1ObjectMeta) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *K8sIoV1ObjectMeta) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *K8sIoV1ObjectMeta) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *K8sIoV1ObjectMeta) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *K8sIoV1ObjectMeta) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


