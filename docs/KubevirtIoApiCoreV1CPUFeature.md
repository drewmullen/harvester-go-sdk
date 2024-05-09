# KubevirtIoApiCoreV1CPUFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the CPU feature | [default to ""]
**Policy** | Pointer to **string** | Policy is the CPU feature attribute which can have the following attributes: force    - The virtual CPU will claim the feature is supported regardless of it being supported by host CPU. require  - Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it. optional - The feature will be supported by virtual CPU if and only if it is supported by host CPU. disable  - The feature will not be supported by virtual CPU. forbid   - Guest creation will fail if the feature is supported by host CPU. Defaults to require | [optional] 

## Methods

### NewKubevirtIoApiCoreV1CPUFeature

`func NewKubevirtIoApiCoreV1CPUFeature(name string, ) *KubevirtIoApiCoreV1CPUFeature`

NewKubevirtIoApiCoreV1CPUFeature instantiates a new KubevirtIoApiCoreV1CPUFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1CPUFeatureWithDefaults

`func NewKubevirtIoApiCoreV1CPUFeatureWithDefaults() *KubevirtIoApiCoreV1CPUFeature`

NewKubevirtIoApiCoreV1CPUFeatureWithDefaults instantiates a new KubevirtIoApiCoreV1CPUFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KubevirtIoApiCoreV1CPUFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1CPUFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1CPUFeature) SetName(v string)`

SetName sets Name field to given value.


### GetPolicy

`func (o *KubevirtIoApiCoreV1CPUFeature) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KubevirtIoApiCoreV1CPUFeature) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KubevirtIoApiCoreV1CPUFeature) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KubevirtIoApiCoreV1CPUFeature) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


