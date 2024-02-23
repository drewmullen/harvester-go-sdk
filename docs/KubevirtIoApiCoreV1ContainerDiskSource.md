# KubevirtIoApiCoreV1ContainerDiskSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** |  | [default to ""]
**ImagePullPolicy** | Pointer to **string** |  | [optional] 
**ImagePullSecret** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1ContainerDiskSource

`func NewKubevirtIoApiCoreV1ContainerDiskSource(image string, ) *KubevirtIoApiCoreV1ContainerDiskSource`

NewKubevirtIoApiCoreV1ContainerDiskSource instantiates a new KubevirtIoApiCoreV1ContainerDiskSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ContainerDiskSourceWithDefaults

`func NewKubevirtIoApiCoreV1ContainerDiskSourceWithDefaults() *KubevirtIoApiCoreV1ContainerDiskSource`

NewKubevirtIoApiCoreV1ContainerDiskSourceWithDefaults instantiates a new KubevirtIoApiCoreV1ContainerDiskSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) SetImage(v string)`

SetImage sets Image field to given value.


### GetImagePullPolicy

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetImagePullSecret

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImagePullSecret() string`

GetImagePullSecret returns the ImagePullSecret field if non-nil, zero value otherwise.

### GetImagePullSecretOk

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetImagePullSecretOk() (*string, bool)`

GetImagePullSecretOk returns a tuple with the ImagePullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecret

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) SetImagePullSecret(v string)`

SetImagePullSecret sets ImagePullSecret field to given value.

### HasImagePullSecret

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) HasImagePullSecret() bool`

HasImagePullSecret returns a boolean if a field has been set.

### GetPath

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KubevirtIoApiCoreV1ContainerDiskSource) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


