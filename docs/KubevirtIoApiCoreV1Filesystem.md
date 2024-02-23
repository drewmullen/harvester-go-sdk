# KubevirtIoApiCoreV1Filesystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to ""]
**Virtiofs** | **map[string]interface{}** |  | 

## Methods

### NewKubevirtIoApiCoreV1Filesystem

`func NewKubevirtIoApiCoreV1Filesystem(name string, virtiofs map[string]interface{}, ) *KubevirtIoApiCoreV1Filesystem`

NewKubevirtIoApiCoreV1Filesystem instantiates a new KubevirtIoApiCoreV1Filesystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1FilesystemWithDefaults

`func NewKubevirtIoApiCoreV1FilesystemWithDefaults() *KubevirtIoApiCoreV1Filesystem`

NewKubevirtIoApiCoreV1FilesystemWithDefaults instantiates a new KubevirtIoApiCoreV1Filesystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubevirtIoApiCoreV1Filesystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Filesystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Filesystem) SetName(v string)`

SetName sets Name field to given value.


### GetVirtiofs

`func (o *KubevirtIoApiCoreV1Filesystem) GetVirtiofs() map[string]interface{}`

GetVirtiofs returns the Virtiofs field if non-nil, zero value otherwise.

### GetVirtiofsOk

`func (o *KubevirtIoApiCoreV1Filesystem) GetVirtiofsOk() (*map[string]interface{}, bool)`

GetVirtiofsOk returns a tuple with the Virtiofs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtiofs

`func (o *KubevirtIoApiCoreV1Filesystem) SetVirtiofs(v map[string]interface{})`

SetVirtiofs sets Virtiofs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


