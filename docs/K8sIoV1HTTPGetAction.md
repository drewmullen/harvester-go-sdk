# K8sIoV1HTTPGetAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**HttpHeaders** | Pointer to [**[]K8sIoV1HTTPHeader**](K8sIoV1HTTPHeader.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Port** | **string** |  | [default to "{}"]
**Scheme** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1HTTPGetAction

`func NewK8sIoV1HTTPGetAction(port string, ) *K8sIoV1HTTPGetAction`

NewK8sIoV1HTTPGetAction instantiates a new K8sIoV1HTTPGetAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1HTTPGetActionWithDefaults

`func NewK8sIoV1HTTPGetActionWithDefaults() *K8sIoV1HTTPGetAction`

NewK8sIoV1HTTPGetActionWithDefaults instantiates a new K8sIoV1HTTPGetAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *K8sIoV1HTTPGetAction) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *K8sIoV1HTTPGetAction) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *K8sIoV1HTTPGetAction) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *K8sIoV1HTTPGetAction) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHttpHeaders

`func (o *K8sIoV1HTTPGetAction) GetHttpHeaders() []K8sIoV1HTTPHeader`

GetHttpHeaders returns the HttpHeaders field if non-nil, zero value otherwise.

### GetHttpHeadersOk

`func (o *K8sIoV1HTTPGetAction) GetHttpHeadersOk() (*[]K8sIoV1HTTPHeader, bool)`

GetHttpHeadersOk returns a tuple with the HttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeaders

`func (o *K8sIoV1HTTPGetAction) SetHttpHeaders(v []K8sIoV1HTTPHeader)`

SetHttpHeaders sets HttpHeaders field to given value.

### HasHttpHeaders

`func (o *K8sIoV1HTTPGetAction) HasHttpHeaders() bool`

HasHttpHeaders returns a boolean if a field has been set.

### GetPath

`func (o *K8sIoV1HTTPGetAction) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *K8sIoV1HTTPGetAction) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *K8sIoV1HTTPGetAction) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *K8sIoV1HTTPGetAction) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPort

`func (o *K8sIoV1HTTPGetAction) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *K8sIoV1HTTPGetAction) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *K8sIoV1HTTPGetAction) SetPort(v string)`

SetPort sets Port field to given value.


### GetScheme

`func (o *K8sIoV1HTTPGetAction) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *K8sIoV1HTTPGetAction) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *K8sIoV1HTTPGetAction) SetScheme(v string)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *K8sIoV1HTTPGetAction) HasScheme() bool`

HasScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


