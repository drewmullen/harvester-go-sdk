# HarvesterhciIoV1beta1UpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]HarvesterhciIoV1beta1Condition**](HarvesterhciIoV1beta1Condition.md) |  | [optional] 
**ImageID** | Pointer to **string** |  | [optional] 
**NodeStatuses** | Pointer to [**map[string]HarvesterhciIoV1beta1NodeUpgradeStatus**](HarvesterhciIoV1beta1NodeUpgradeStatus.md) |  | [optional] 
**PreviousVersion** | Pointer to **string** |  | [optional] 
**RepoInfo** | Pointer to **string** |  | [optional] 
**SingleNode** | Pointer to **string** |  | [optional] 
**UpgradeLog** | Pointer to **string** |  | [optional] 

## Methods

### NewHarvesterhciIoV1beta1UpgradeStatus

`func NewHarvesterhciIoV1beta1UpgradeStatus() *HarvesterhciIoV1beta1UpgradeStatus`

NewHarvesterhciIoV1beta1UpgradeStatus instantiates a new HarvesterhciIoV1beta1UpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1UpgradeStatusWithDefaults

`func NewHarvesterhciIoV1beta1UpgradeStatusWithDefaults() *HarvesterhciIoV1beta1UpgradeStatus`

NewHarvesterhciIoV1beta1UpgradeStatusWithDefaults instantiates a new HarvesterhciIoV1beta1UpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetConditions() []HarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetConditionsOk() (*[]HarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetConditions(v []HarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetImageID

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetImageID() string`

GetImageID returns the ImageID field if non-nil, zero value otherwise.

### GetImageIDOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetImageIDOk() (*string, bool)`

GetImageIDOk returns a tuple with the ImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageID

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetImageID(v string)`

SetImageID sets ImageID field to given value.

### HasImageID

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasImageID() bool`

HasImageID returns a boolean if a field has been set.

### GetNodeStatuses

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetNodeStatuses() map[string]HarvesterhciIoV1beta1NodeUpgradeStatus`

GetNodeStatuses returns the NodeStatuses field if non-nil, zero value otherwise.

### GetNodeStatusesOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetNodeStatusesOk() (*map[string]HarvesterhciIoV1beta1NodeUpgradeStatus, bool)`

GetNodeStatusesOk returns a tuple with the NodeStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatuses

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetNodeStatuses(v map[string]HarvesterhciIoV1beta1NodeUpgradeStatus)`

SetNodeStatuses sets NodeStatuses field to given value.

### HasNodeStatuses

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasNodeStatuses() bool`

HasNodeStatuses returns a boolean if a field has been set.

### GetPreviousVersion

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### GetRepoInfo

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetRepoInfo() string`

GetRepoInfo returns the RepoInfo field if non-nil, zero value otherwise.

### GetRepoInfoOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetRepoInfoOk() (*string, bool)`

GetRepoInfoOk returns a tuple with the RepoInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoInfo

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetRepoInfo(v string)`

SetRepoInfo sets RepoInfo field to given value.

### HasRepoInfo

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasRepoInfo() bool`

HasRepoInfo returns a boolean if a field has been set.

### GetSingleNode

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetSingleNode() string`

GetSingleNode returns the SingleNode field if non-nil, zero value otherwise.

### GetSingleNodeOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetSingleNodeOk() (*string, bool)`

GetSingleNodeOk returns a tuple with the SingleNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleNode

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetSingleNode(v string)`

SetSingleNode sets SingleNode field to given value.

### HasSingleNode

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasSingleNode() bool`

HasSingleNode returns a boolean if a field has been set.

### GetUpgradeLog

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetUpgradeLog() string`

GetUpgradeLog returns the UpgradeLog field if non-nil, zero value otherwise.

### GetUpgradeLogOk

`func (o *HarvesterhciIoV1beta1UpgradeStatus) GetUpgradeLogOk() (*string, bool)`

GetUpgradeLogOk returns a tuple with the UpgradeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeLog

`func (o *HarvesterhciIoV1beta1UpgradeStatus) SetUpgradeLog(v string)`

SetUpgradeLog sets UpgradeLog field to given value.

### HasUpgradeLog

`func (o *HarvesterhciIoV1beta1UpgradeStatus) HasUpgradeLog() bool`

HasUpgradeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


