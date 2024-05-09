# K8sIoV1DownwardAPIVolumeFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldRef** | Pointer to [**K8sIoV1ObjectFieldSelector**](K8sIoV1ObjectFieldSelector.md) | Required: Selects a field of the pod: only annotations, labels, name and namespace are supported. | [optional] 
**Mode** | Pointer to **int32** | Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. | [optional] 
**Path** | **string** | Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the &#39;..&#39; path. Must be utf-8 encoded. The first item of the relative path must not start with &#39;..&#39; | [default to ""]
**ResourceFieldRef** | Pointer to [**K8sIoV1ResourceFieldSelector**](K8sIoV1ResourceFieldSelector.md) | Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported. | [optional] 

## Methods

### NewK8sIoV1DownwardAPIVolumeFile

`func NewK8sIoV1DownwardAPIVolumeFile(path string, ) *K8sIoV1DownwardAPIVolumeFile`

NewK8sIoV1DownwardAPIVolumeFile instantiates a new K8sIoV1DownwardAPIVolumeFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1DownwardAPIVolumeFileWithDefaults

`func NewK8sIoV1DownwardAPIVolumeFileWithDefaults() *K8sIoV1DownwardAPIVolumeFile`

NewK8sIoV1DownwardAPIVolumeFileWithDefaults instantiates a new K8sIoV1DownwardAPIVolumeFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) GetFieldRef() K8sIoV1ObjectFieldSelector`

GetFieldRef returns the FieldRef field if non-nil, zero value otherwise.

### GetFieldRefOk

`func (o *K8sIoV1DownwardAPIVolumeFile) GetFieldRefOk() (*K8sIoV1ObjectFieldSelector, bool)`

GetFieldRefOk returns a tuple with the FieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) SetFieldRef(v K8sIoV1ObjectFieldSelector)`

SetFieldRef sets FieldRef field to given value.

### HasFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) HasFieldRef() bool`

HasFieldRef returns a boolean if a field has been set.

### GetMode

`func (o *K8sIoV1DownwardAPIVolumeFile) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *K8sIoV1DownwardAPIVolumeFile) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *K8sIoV1DownwardAPIVolumeFile) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *K8sIoV1DownwardAPIVolumeFile) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *K8sIoV1DownwardAPIVolumeFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *K8sIoV1DownwardAPIVolumeFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *K8sIoV1DownwardAPIVolumeFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetResourceFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) GetResourceFieldRef() K8sIoV1ResourceFieldSelector`

GetResourceFieldRef returns the ResourceFieldRef field if non-nil, zero value otherwise.

### GetResourceFieldRefOk

`func (o *K8sIoV1DownwardAPIVolumeFile) GetResourceFieldRefOk() (*K8sIoV1ResourceFieldSelector, bool)`

GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) SetResourceFieldRef(v K8sIoV1ResourceFieldSelector)`

SetResourceFieldRef sets ResourceFieldRef field to given value.

### HasResourceFieldRef

`func (o *K8sIoV1DownwardAPIVolumeFile) HasResourceFieldRef() bool`

HasResourceFieldRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


