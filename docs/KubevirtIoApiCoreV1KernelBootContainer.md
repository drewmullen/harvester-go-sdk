# KubevirtIoApiCoreV1KernelBootContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | Image that contains initrd / kernel files. | [default to ""]
**ImagePullPolicy** | Pointer to **string** | Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images  Possible enum values:  - &#x60;\&quot;Always\&quot;&#x60; means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.  - &#x60;\&quot;IfNotPresent\&quot;&#x60; means that kubelet pulls if the image isn&#39;t present on disk. Container will fail if the image isn&#39;t present and the pull fails.  - &#x60;\&quot;Never\&quot;&#x60; means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn&#39;t present | [optional] 
**ImagePullSecret** | Pointer to **string** | ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist. | [optional] 
**InitrdPath** | Pointer to **string** | the fully-qualified path to the ramdisk image in the host OS | [optional] 
**KernelPath** | Pointer to **string** | The fully-qualified path to the kernel image in the host OS | [optional] 

## Methods

### NewKubevirtIoApiCoreV1KernelBootContainer

`func NewKubevirtIoApiCoreV1KernelBootContainer(image string, ) *KubevirtIoApiCoreV1KernelBootContainer`

NewKubevirtIoApiCoreV1KernelBootContainer instantiates a new KubevirtIoApiCoreV1KernelBootContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1KernelBootContainerWithDefaults

`func NewKubevirtIoApiCoreV1KernelBootContainerWithDefaults() *KubevirtIoApiCoreV1KernelBootContainer`

NewKubevirtIoApiCoreV1KernelBootContainerWithDefaults instantiates a new KubevirtIoApiCoreV1KernelBootContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *KubevirtIoApiCoreV1KernelBootContainer) SetImage(v string)`

SetImage sets Image field to given value.


### GetImagePullPolicy

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImagePullPolicy() string`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImagePullPolicyOk() (*string, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *KubevirtIoApiCoreV1KernelBootContainer) SetImagePullPolicy(v string)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *KubevirtIoApiCoreV1KernelBootContainer) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### GetImagePullSecret

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImagePullSecret() string`

GetImagePullSecret returns the ImagePullSecret field if non-nil, zero value otherwise.

### GetImagePullSecretOk

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetImagePullSecretOk() (*string, bool)`

GetImagePullSecretOk returns a tuple with the ImagePullSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullSecret

`func (o *KubevirtIoApiCoreV1KernelBootContainer) SetImagePullSecret(v string)`

SetImagePullSecret sets ImagePullSecret field to given value.

### HasImagePullSecret

`func (o *KubevirtIoApiCoreV1KernelBootContainer) HasImagePullSecret() bool`

HasImagePullSecret returns a boolean if a field has been set.

### GetInitrdPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetInitrdPath() string`

GetInitrdPath returns the InitrdPath field if non-nil, zero value otherwise.

### GetInitrdPathOk

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetInitrdPathOk() (*string, bool)`

GetInitrdPathOk returns a tuple with the InitrdPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitrdPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) SetInitrdPath(v string)`

SetInitrdPath sets InitrdPath field to given value.

### HasInitrdPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) HasInitrdPath() bool`

HasInitrdPath returns a boolean if a field has been set.

### GetKernelPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetKernelPath() string`

GetKernelPath returns the KernelPath field if non-nil, zero value otherwise.

### GetKernelPathOk

`func (o *KubevirtIoApiCoreV1KernelBootContainer) GetKernelPathOk() (*string, bool)`

GetKernelPathOk returns a tuple with the KernelPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) SetKernelPath(v string)`

SetKernelPath sets KernelPath field to given value.

### HasKernelPath

`func (o *KubevirtIoApiCoreV1KernelBootContainer) HasKernelPath() bool`

HasKernelPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


