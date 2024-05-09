# K8sIoV1ExecAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **[]string** | Command is the command line to execute inside the container, the working directory for the command  is root (&#39;/&#39;) in the container&#39;s filesystem. The command is simply exec&#39;d, it is not run inside a shell, so traditional shell instructions (&#39;|&#39;, etc) won&#39;t work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy. | [optional] 

## Methods

### NewK8sIoV1ExecAction

`func NewK8sIoV1ExecAction() *K8sIoV1ExecAction`

NewK8sIoV1ExecAction instantiates a new K8sIoV1ExecAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ExecActionWithDefaults

`func NewK8sIoV1ExecActionWithDefaults() *K8sIoV1ExecAction`

NewK8sIoV1ExecActionWithDefaults instantiates a new K8sIoV1ExecAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *K8sIoV1ExecAction) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *K8sIoV1ExecAction) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *K8sIoV1ExecAction) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *K8sIoV1ExecAction) HasCommand() bool`

HasCommand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


