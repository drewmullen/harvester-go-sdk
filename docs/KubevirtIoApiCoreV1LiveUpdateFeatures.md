# KubevirtIoApiCoreV1LiveUpdateFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinity** | Pointer to **map[string]interface{}** |  | [optional] 
**Cpu** | Pointer to [**KubevirtIoApiCoreV1LiveUpdateCPU**](KubevirtIoApiCoreV1LiveUpdateCPU.md) |  | [optional] 
**Memory** | Pointer to [**KubevirtIoApiCoreV1LiveUpdateMemory**](KubevirtIoApiCoreV1LiveUpdateMemory.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1LiveUpdateFeatures

`func NewKubevirtIoApiCoreV1LiveUpdateFeatures() *KubevirtIoApiCoreV1LiveUpdateFeatures`

NewKubevirtIoApiCoreV1LiveUpdateFeatures instantiates a new KubevirtIoApiCoreV1LiveUpdateFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1LiveUpdateFeaturesWithDefaults

`func NewKubevirtIoApiCoreV1LiveUpdateFeaturesWithDefaults() *KubevirtIoApiCoreV1LiveUpdateFeatures`

NewKubevirtIoApiCoreV1LiveUpdateFeaturesWithDefaults instantiates a new KubevirtIoApiCoreV1LiveUpdateFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinity

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetAffinity() map[string]interface{}`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetAffinityOk() (*map[string]interface{}, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetAffinity(v map[string]interface{})`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetCpu

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetCpu() KubevirtIoApiCoreV1LiveUpdateCPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetCpuOk() (*KubevirtIoApiCoreV1LiveUpdateCPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetCpu(v KubevirtIoApiCoreV1LiveUpdateCPU)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetMemory() KubevirtIoApiCoreV1LiveUpdateMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) GetMemoryOk() (*KubevirtIoApiCoreV1LiveUpdateMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) SetMemory(v KubevirtIoApiCoreV1LiveUpdateMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KubevirtIoApiCoreV1LiveUpdateFeatures) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


