# KubevirtIoApiCoreV1CPU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cores** | Pointer to **int64** |  | [optional] 
**DedicatedCpuPlacement** | Pointer to **bool** |  | [optional] 
**Features** | Pointer to [**[]KubevirtIoApiCoreV1CPUFeature**](KubevirtIoApiCoreV1CPUFeature.md) |  | [optional] 
**IsolateEmulatorThread** | Pointer to **bool** |  | [optional] 
**MaxSockets** | Pointer to **int64** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Numa** | Pointer to [**KubevirtIoApiCoreV1NUMA**](KubevirtIoApiCoreV1NUMA.md) |  | [optional] 
**Realtime** | Pointer to [**KubevirtIoApiCoreV1Realtime**](KubevirtIoApiCoreV1Realtime.md) |  | [optional] 
**Sockets** | Pointer to **int64** |  | [optional] 
**Threads** | Pointer to **int64** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1CPU

`func NewKubevirtIoApiCoreV1CPU() *KubevirtIoApiCoreV1CPU`

NewKubevirtIoApiCoreV1CPU instantiates a new KubevirtIoApiCoreV1CPU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1CPUWithDefaults

`func NewKubevirtIoApiCoreV1CPUWithDefaults() *KubevirtIoApiCoreV1CPU`

NewKubevirtIoApiCoreV1CPUWithDefaults instantiates a new KubevirtIoApiCoreV1CPU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *KubevirtIoApiCoreV1CPU) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *KubevirtIoApiCoreV1CPU) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *KubevirtIoApiCoreV1CPU) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *KubevirtIoApiCoreV1CPU) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetDedicatedCpuPlacement

`func (o *KubevirtIoApiCoreV1CPU) GetDedicatedCpuPlacement() bool`

GetDedicatedCpuPlacement returns the DedicatedCpuPlacement field if non-nil, zero value otherwise.

### GetDedicatedCpuPlacementOk

`func (o *KubevirtIoApiCoreV1CPU) GetDedicatedCpuPlacementOk() (*bool, bool)`

GetDedicatedCpuPlacementOk returns a tuple with the DedicatedCpuPlacement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedCpuPlacement

`func (o *KubevirtIoApiCoreV1CPU) SetDedicatedCpuPlacement(v bool)`

SetDedicatedCpuPlacement sets DedicatedCpuPlacement field to given value.

### HasDedicatedCpuPlacement

`func (o *KubevirtIoApiCoreV1CPU) HasDedicatedCpuPlacement() bool`

HasDedicatedCpuPlacement returns a boolean if a field has been set.

### GetFeatures

`func (o *KubevirtIoApiCoreV1CPU) GetFeatures() []KubevirtIoApiCoreV1CPUFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *KubevirtIoApiCoreV1CPU) GetFeaturesOk() (*[]KubevirtIoApiCoreV1CPUFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *KubevirtIoApiCoreV1CPU) SetFeatures(v []KubevirtIoApiCoreV1CPUFeature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *KubevirtIoApiCoreV1CPU) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetIsolateEmulatorThread

`func (o *KubevirtIoApiCoreV1CPU) GetIsolateEmulatorThread() bool`

GetIsolateEmulatorThread returns the IsolateEmulatorThread field if non-nil, zero value otherwise.

### GetIsolateEmulatorThreadOk

`func (o *KubevirtIoApiCoreV1CPU) GetIsolateEmulatorThreadOk() (*bool, bool)`

GetIsolateEmulatorThreadOk returns a tuple with the IsolateEmulatorThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateEmulatorThread

`func (o *KubevirtIoApiCoreV1CPU) SetIsolateEmulatorThread(v bool)`

SetIsolateEmulatorThread sets IsolateEmulatorThread field to given value.

### HasIsolateEmulatorThread

`func (o *KubevirtIoApiCoreV1CPU) HasIsolateEmulatorThread() bool`

HasIsolateEmulatorThread returns a boolean if a field has been set.

### GetMaxSockets

`func (o *KubevirtIoApiCoreV1CPU) GetMaxSockets() int64`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *KubevirtIoApiCoreV1CPU) GetMaxSocketsOk() (*int64, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *KubevirtIoApiCoreV1CPU) SetMaxSockets(v int64)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *KubevirtIoApiCoreV1CPU) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### GetModel

`func (o *KubevirtIoApiCoreV1CPU) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *KubevirtIoApiCoreV1CPU) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *KubevirtIoApiCoreV1CPU) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *KubevirtIoApiCoreV1CPU) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNuma

`func (o *KubevirtIoApiCoreV1CPU) GetNuma() KubevirtIoApiCoreV1NUMA`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *KubevirtIoApiCoreV1CPU) GetNumaOk() (*KubevirtIoApiCoreV1NUMA, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *KubevirtIoApiCoreV1CPU) SetNuma(v KubevirtIoApiCoreV1NUMA)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *KubevirtIoApiCoreV1CPU) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetRealtime

`func (o *KubevirtIoApiCoreV1CPU) GetRealtime() KubevirtIoApiCoreV1Realtime`

GetRealtime returns the Realtime field if non-nil, zero value otherwise.

### GetRealtimeOk

`func (o *KubevirtIoApiCoreV1CPU) GetRealtimeOk() (*KubevirtIoApiCoreV1Realtime, bool)`

GetRealtimeOk returns a tuple with the Realtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealtime

`func (o *KubevirtIoApiCoreV1CPU) SetRealtime(v KubevirtIoApiCoreV1Realtime)`

SetRealtime sets Realtime field to given value.

### HasRealtime

`func (o *KubevirtIoApiCoreV1CPU) HasRealtime() bool`

HasRealtime returns a boolean if a field has been set.

### GetSockets

`func (o *KubevirtIoApiCoreV1CPU) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *KubevirtIoApiCoreV1CPU) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *KubevirtIoApiCoreV1CPU) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *KubevirtIoApiCoreV1CPU) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetThreads

`func (o *KubevirtIoApiCoreV1CPU) GetThreads() int64`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *KubevirtIoApiCoreV1CPU) GetThreadsOk() (*int64, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *KubevirtIoApiCoreV1CPU) SetThreads(v int64)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *KubevirtIoApiCoreV1CPU) HasThreads() bool`

HasThreads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


