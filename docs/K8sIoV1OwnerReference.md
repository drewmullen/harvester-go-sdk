# K8sIoV1OwnerReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | [default to ""]
**BlockOwnerDeletion** | Pointer to **bool** |  | [optional] 
**Controller** | Pointer to **bool** |  | [optional] 
**Kind** | **string** |  | [default to ""]
**Name** | **string** |  | [default to ""]
**Uid** | **string** |  | [default to ""]

## Methods

### NewK8sIoV1OwnerReference

`func NewK8sIoV1OwnerReference(apiVersion string, kind string, name string, uid string, ) *K8sIoV1OwnerReference`

NewK8sIoV1OwnerReference instantiates a new K8sIoV1OwnerReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1OwnerReferenceWithDefaults

`func NewK8sIoV1OwnerReferenceWithDefaults() *K8sIoV1OwnerReference`

NewK8sIoV1OwnerReferenceWithDefaults instantiates a new K8sIoV1OwnerReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1OwnerReference) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1OwnerReference) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1OwnerReference) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetBlockOwnerDeletion

`func (o *K8sIoV1OwnerReference) GetBlockOwnerDeletion() bool`

GetBlockOwnerDeletion returns the BlockOwnerDeletion field if non-nil, zero value otherwise.

### GetBlockOwnerDeletionOk

`func (o *K8sIoV1OwnerReference) GetBlockOwnerDeletionOk() (*bool, bool)`

GetBlockOwnerDeletionOk returns a tuple with the BlockOwnerDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockOwnerDeletion

`func (o *K8sIoV1OwnerReference) SetBlockOwnerDeletion(v bool)`

SetBlockOwnerDeletion sets BlockOwnerDeletion field to given value.

### HasBlockOwnerDeletion

`func (o *K8sIoV1OwnerReference) HasBlockOwnerDeletion() bool`

HasBlockOwnerDeletion returns a boolean if a field has been set.

### GetController

`func (o *K8sIoV1OwnerReference) GetController() bool`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *K8sIoV1OwnerReference) GetControllerOk() (*bool, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *K8sIoV1OwnerReference) SetController(v bool)`

SetController sets Controller field to given value.

### HasController

`func (o *K8sIoV1OwnerReference) HasController() bool`

HasController returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1OwnerReference) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1OwnerReference) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1OwnerReference) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *K8sIoV1OwnerReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *K8sIoV1OwnerReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *K8sIoV1OwnerReference) SetName(v string)`

SetName sets Name field to given value.


### GetUid

`func (o *K8sIoV1OwnerReference) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *K8sIoV1OwnerReference) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *K8sIoV1OwnerReference) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


