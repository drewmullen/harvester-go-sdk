# KubevirtIoApiCoreV1Probe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exec** | Pointer to [**K8sIoV1ExecAction**](K8sIoV1ExecAction.md) |  | [optional] 
**FailureThreshold** | Pointer to **int32** |  | [optional] 
**GuestAgentPing** | Pointer to **map[string]interface{}** |  | [optional] 
**HttpGet** | Pointer to [**K8sIoV1HTTPGetAction**](K8sIoV1HTTPGetAction.md) |  | [optional] 
**InitialDelaySeconds** | Pointer to **int32** |  | [optional] 
**PeriodSeconds** | Pointer to **int32** |  | [optional] 
**SuccessThreshold** | Pointer to **int32** |  | [optional] 
**TcpSocket** | Pointer to [**K8sIoV1TCPSocketAction**](K8sIoV1TCPSocketAction.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Probe

`func NewKubevirtIoApiCoreV1Probe() *KubevirtIoApiCoreV1Probe`

NewKubevirtIoApiCoreV1Probe instantiates a new KubevirtIoApiCoreV1Probe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1ProbeWithDefaults

`func NewKubevirtIoApiCoreV1ProbeWithDefaults() *KubevirtIoApiCoreV1Probe`

NewKubevirtIoApiCoreV1ProbeWithDefaults instantiates a new KubevirtIoApiCoreV1Probe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExec

`func (o *KubevirtIoApiCoreV1Probe) GetExec() K8sIoV1ExecAction`

GetExec returns the Exec field if non-nil, zero value otherwise.

### GetExecOk

`func (o *KubevirtIoApiCoreV1Probe) GetExecOk() (*K8sIoV1ExecAction, bool)`

GetExecOk returns a tuple with the Exec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExec

`func (o *KubevirtIoApiCoreV1Probe) SetExec(v K8sIoV1ExecAction)`

SetExec sets Exec field to given value.

### HasExec

`func (o *KubevirtIoApiCoreV1Probe) HasExec() bool`

HasExec returns a boolean if a field has been set.

### GetFailureThreshold

`func (o *KubevirtIoApiCoreV1Probe) GetFailureThreshold() int32`

GetFailureThreshold returns the FailureThreshold field if non-nil, zero value otherwise.

### GetFailureThresholdOk

`func (o *KubevirtIoApiCoreV1Probe) GetFailureThresholdOk() (*int32, bool)`

GetFailureThresholdOk returns a tuple with the FailureThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureThreshold

`func (o *KubevirtIoApiCoreV1Probe) SetFailureThreshold(v int32)`

SetFailureThreshold sets FailureThreshold field to given value.

### HasFailureThreshold

`func (o *KubevirtIoApiCoreV1Probe) HasFailureThreshold() bool`

HasFailureThreshold returns a boolean if a field has been set.

### GetGuestAgentPing

`func (o *KubevirtIoApiCoreV1Probe) GetGuestAgentPing() map[string]interface{}`

GetGuestAgentPing returns the GuestAgentPing field if non-nil, zero value otherwise.

### GetGuestAgentPingOk

`func (o *KubevirtIoApiCoreV1Probe) GetGuestAgentPingOk() (*map[string]interface{}, bool)`

GetGuestAgentPingOk returns a tuple with the GuestAgentPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestAgentPing

`func (o *KubevirtIoApiCoreV1Probe) SetGuestAgentPing(v map[string]interface{})`

SetGuestAgentPing sets GuestAgentPing field to given value.

### HasGuestAgentPing

`func (o *KubevirtIoApiCoreV1Probe) HasGuestAgentPing() bool`

HasGuestAgentPing returns a boolean if a field has been set.

### GetHttpGet

`func (o *KubevirtIoApiCoreV1Probe) GetHttpGet() K8sIoV1HTTPGetAction`

GetHttpGet returns the HttpGet field if non-nil, zero value otherwise.

### GetHttpGetOk

`func (o *KubevirtIoApiCoreV1Probe) GetHttpGetOk() (*K8sIoV1HTTPGetAction, bool)`

GetHttpGetOk returns a tuple with the HttpGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpGet

`func (o *KubevirtIoApiCoreV1Probe) SetHttpGet(v K8sIoV1HTTPGetAction)`

SetHttpGet sets HttpGet field to given value.

### HasHttpGet

`func (o *KubevirtIoApiCoreV1Probe) HasHttpGet() bool`

HasHttpGet returns a boolean if a field has been set.

### GetInitialDelaySeconds

`func (o *KubevirtIoApiCoreV1Probe) GetInitialDelaySeconds() int32`

GetInitialDelaySeconds returns the InitialDelaySeconds field if non-nil, zero value otherwise.

### GetInitialDelaySecondsOk

`func (o *KubevirtIoApiCoreV1Probe) GetInitialDelaySecondsOk() (*int32, bool)`

GetInitialDelaySecondsOk returns a tuple with the InitialDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDelaySeconds

`func (o *KubevirtIoApiCoreV1Probe) SetInitialDelaySeconds(v int32)`

SetInitialDelaySeconds sets InitialDelaySeconds field to given value.

### HasInitialDelaySeconds

`func (o *KubevirtIoApiCoreV1Probe) HasInitialDelaySeconds() bool`

HasInitialDelaySeconds returns a boolean if a field has been set.

### GetPeriodSeconds

`func (o *KubevirtIoApiCoreV1Probe) GetPeriodSeconds() int32`

GetPeriodSeconds returns the PeriodSeconds field if non-nil, zero value otherwise.

### GetPeriodSecondsOk

`func (o *KubevirtIoApiCoreV1Probe) GetPeriodSecondsOk() (*int32, bool)`

GetPeriodSecondsOk returns a tuple with the PeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodSeconds

`func (o *KubevirtIoApiCoreV1Probe) SetPeriodSeconds(v int32)`

SetPeriodSeconds sets PeriodSeconds field to given value.

### HasPeriodSeconds

`func (o *KubevirtIoApiCoreV1Probe) HasPeriodSeconds() bool`

HasPeriodSeconds returns a boolean if a field has been set.

### GetSuccessThreshold

`func (o *KubevirtIoApiCoreV1Probe) GetSuccessThreshold() int32`

GetSuccessThreshold returns the SuccessThreshold field if non-nil, zero value otherwise.

### GetSuccessThresholdOk

`func (o *KubevirtIoApiCoreV1Probe) GetSuccessThresholdOk() (*int32, bool)`

GetSuccessThresholdOk returns a tuple with the SuccessThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessThreshold

`func (o *KubevirtIoApiCoreV1Probe) SetSuccessThreshold(v int32)`

SetSuccessThreshold sets SuccessThreshold field to given value.

### HasSuccessThreshold

`func (o *KubevirtIoApiCoreV1Probe) HasSuccessThreshold() bool`

HasSuccessThreshold returns a boolean if a field has been set.

### GetTcpSocket

`func (o *KubevirtIoApiCoreV1Probe) GetTcpSocket() K8sIoV1TCPSocketAction`

GetTcpSocket returns the TcpSocket field if non-nil, zero value otherwise.

### GetTcpSocketOk

`func (o *KubevirtIoApiCoreV1Probe) GetTcpSocketOk() (*K8sIoV1TCPSocketAction, bool)`

GetTcpSocketOk returns a tuple with the TcpSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSocket

`func (o *KubevirtIoApiCoreV1Probe) SetTcpSocket(v K8sIoV1TCPSocketAction)`

SetTcpSocket sets TcpSocket field to given value.

### HasTcpSocket

`func (o *KubevirtIoApiCoreV1Probe) HasTcpSocket() bool`

HasTcpSocket returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *KubevirtIoApiCoreV1Probe) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *KubevirtIoApiCoreV1Probe) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *KubevirtIoApiCoreV1Probe) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *KubevirtIoApiCoreV1Probe) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


