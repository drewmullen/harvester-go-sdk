# K8sIoV1Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** |  | 
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to [**K8sIoV1StatusDetails**](K8sIoV1StatusDetails.md) |  | [optional] 
**Kind** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**K8sIoV1ListMeta**](K8sIoV1ListMeta.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewK8sIoV1Status

`func NewK8sIoV1Status(apiVersion string, kind string, ) *K8sIoV1Status`

NewK8sIoV1Status instantiates a new K8sIoV1Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1StatusWithDefaults

`func NewK8sIoV1StatusWithDefaults() *K8sIoV1Status`

NewK8sIoV1StatusWithDefaults instantiates a new K8sIoV1Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1Status) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1Status) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1Status) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetCode

`func (o *K8sIoV1Status) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *K8sIoV1Status) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *K8sIoV1Status) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *K8sIoV1Status) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *K8sIoV1Status) GetDetails() K8sIoV1StatusDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *K8sIoV1Status) GetDetailsOk() (*K8sIoV1StatusDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *K8sIoV1Status) SetDetails(v K8sIoV1StatusDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *K8sIoV1Status) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetKind

`func (o *K8sIoV1Status) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *K8sIoV1Status) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *K8sIoV1Status) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMessage

`func (o *K8sIoV1Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *K8sIoV1Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *K8sIoV1Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *K8sIoV1Status) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMetadata

`func (o *K8sIoV1Status) GetMetadata() K8sIoV1ListMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *K8sIoV1Status) GetMetadataOk() (*K8sIoV1ListMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *K8sIoV1Status) SetMetadata(v K8sIoV1ListMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *K8sIoV1Status) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReason

`func (o *K8sIoV1Status) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *K8sIoV1Status) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *K8sIoV1Status) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *K8sIoV1Status) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *K8sIoV1Status) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *K8sIoV1Status) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *K8sIoV1Status) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *K8sIoV1Status) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


