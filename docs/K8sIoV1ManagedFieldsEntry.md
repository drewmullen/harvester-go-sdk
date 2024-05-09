# K8sIoV1ManagedFieldsEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the version of this resource that this field set applies to. The format is \&quot;group/version\&quot; just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted. | [optional] 
**FieldsType** | Pointer to **string** | FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: \&quot;FieldsV1\&quot; | [optional] 
**FieldsV1** | Pointer to **map[string]interface{}** | FieldsV1 holds the first JSON version format as described in the \&quot;FieldsV1\&quot; type. | [optional] 
**Manager** | Pointer to **string** | Manager is an identifier of the workflow managing these fields. | [optional] 
**Operation** | Pointer to **string** | Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are &#39;Apply&#39; and &#39;Update&#39;. | [optional] 
**Subresource** | Pointer to **string** | Subresource is the name of the subresource used to update that object, or empty string if the object was updated through the main resource. The value of this field is used to distinguish between managers, even if they share the same name. For example, a status update will be distinct from a regular update using the same manager name. Note that the APIVersion field is not related to the Subresource field and it always corresponds to the version of the main resource. | [optional] 
**Time** | Pointer to **string** | Time is the timestamp of when the ManagedFields entry was added. The timestamp will also be updated if a field is added, the manager changes any of the owned fields value or removes a field. The timestamp does not update when a field is removed from the entry because another manager took it over. | [optional] 

## Methods

### NewK8sIoV1ManagedFieldsEntry

`func NewK8sIoV1ManagedFieldsEntry() *K8sIoV1ManagedFieldsEntry`

NewK8sIoV1ManagedFieldsEntry instantiates a new K8sIoV1ManagedFieldsEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1ManagedFieldsEntryWithDefaults

`func NewK8sIoV1ManagedFieldsEntryWithDefaults() *K8sIoV1ManagedFieldsEntry`

NewK8sIoV1ManagedFieldsEntryWithDefaults instantiates a new K8sIoV1ManagedFieldsEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *K8sIoV1ManagedFieldsEntry) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *K8sIoV1ManagedFieldsEntry) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *K8sIoV1ManagedFieldsEntry) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *K8sIoV1ManagedFieldsEntry) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetFieldsType

`func (o *K8sIoV1ManagedFieldsEntry) GetFieldsType() string`

GetFieldsType returns the FieldsType field if non-nil, zero value otherwise.

### GetFieldsTypeOk

`func (o *K8sIoV1ManagedFieldsEntry) GetFieldsTypeOk() (*string, bool)`

GetFieldsTypeOk returns a tuple with the FieldsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsType

`func (o *K8sIoV1ManagedFieldsEntry) SetFieldsType(v string)`

SetFieldsType sets FieldsType field to given value.

### HasFieldsType

`func (o *K8sIoV1ManagedFieldsEntry) HasFieldsType() bool`

HasFieldsType returns a boolean if a field has been set.

### GetFieldsV1

`func (o *K8sIoV1ManagedFieldsEntry) GetFieldsV1() map[string]interface{}`

GetFieldsV1 returns the FieldsV1 field if non-nil, zero value otherwise.

### GetFieldsV1Ok

`func (o *K8sIoV1ManagedFieldsEntry) GetFieldsV1Ok() (*map[string]interface{}, bool)`

GetFieldsV1Ok returns a tuple with the FieldsV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldsV1

`func (o *K8sIoV1ManagedFieldsEntry) SetFieldsV1(v map[string]interface{})`

SetFieldsV1 sets FieldsV1 field to given value.

### HasFieldsV1

`func (o *K8sIoV1ManagedFieldsEntry) HasFieldsV1() bool`

HasFieldsV1 returns a boolean if a field has been set.

### GetManager

`func (o *K8sIoV1ManagedFieldsEntry) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *K8sIoV1ManagedFieldsEntry) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *K8sIoV1ManagedFieldsEntry) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *K8sIoV1ManagedFieldsEntry) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetOperation

`func (o *K8sIoV1ManagedFieldsEntry) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *K8sIoV1ManagedFieldsEntry) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *K8sIoV1ManagedFieldsEntry) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *K8sIoV1ManagedFieldsEntry) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetSubresource

`func (o *K8sIoV1ManagedFieldsEntry) GetSubresource() string`

GetSubresource returns the Subresource field if non-nil, zero value otherwise.

### GetSubresourceOk

`func (o *K8sIoV1ManagedFieldsEntry) GetSubresourceOk() (*string, bool)`

GetSubresourceOk returns a tuple with the Subresource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresource

`func (o *K8sIoV1ManagedFieldsEntry) SetSubresource(v string)`

SetSubresource sets Subresource field to given value.

### HasSubresource

`func (o *K8sIoV1ManagedFieldsEntry) HasSubresource() bool`

HasSubresource returns a boolean if a field has been set.

### GetTime

`func (o *K8sIoV1ManagedFieldsEntry) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *K8sIoV1ManagedFieldsEntry) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *K8sIoV1ManagedFieldsEntry) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *K8sIoV1ManagedFieldsEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


