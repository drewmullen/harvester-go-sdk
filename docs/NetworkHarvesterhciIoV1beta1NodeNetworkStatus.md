# NetworkHarvesterhciIoV1beta1NodeNetworkStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]NetworkHarvesterhciIoV1beta1Condition**](NetworkHarvesterhciIoV1beta1Condition.md) |  | [optional] 
**NetworkIDs** | Pointer to **[]int32** |  | [optional] 
**NetworkLinkStatus** | Pointer to [**map[string]NetworkHarvesterhciIoV1beta1LinkStatus**](NetworkHarvesterhciIoV1beta1LinkStatus.md) |  | [optional] 
**Nics** | Pointer to [**[]NetworkHarvesterhciIoV1beta1NetworkInterface**](NetworkHarvesterhciIoV1beta1NetworkInterface.md) |  | [optional] 

## Methods

### NewNetworkHarvesterhciIoV1beta1NodeNetworkStatus

`func NewNetworkHarvesterhciIoV1beta1NodeNetworkStatus() *NetworkHarvesterhciIoV1beta1NodeNetworkStatus`

NewNetworkHarvesterhciIoV1beta1NodeNetworkStatus instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1NodeNetworkStatusWithDefaults

`func NewNetworkHarvesterhciIoV1beta1NodeNetworkStatusWithDefaults() *NetworkHarvesterhciIoV1beta1NodeNetworkStatus`

NewNetworkHarvesterhciIoV1beta1NodeNetworkStatusWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1NodeNetworkStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetConditions() []NetworkHarvesterhciIoV1beta1Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetConditionsOk() (*[]NetworkHarvesterhciIoV1beta1Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) SetConditions(v []NetworkHarvesterhciIoV1beta1Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetNetworkIDs

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNetworkIDs() []int32`

GetNetworkIDs returns the NetworkIDs field if non-nil, zero value otherwise.

### GetNetworkIDsOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNetworkIDsOk() (*[]int32, bool)`

GetNetworkIDsOk returns a tuple with the NetworkIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIDs

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) SetNetworkIDs(v []int32)`

SetNetworkIDs sets NetworkIDs field to given value.

### HasNetworkIDs

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) HasNetworkIDs() bool`

HasNetworkIDs returns a boolean if a field has been set.

### GetNetworkLinkStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNetworkLinkStatus() map[string]NetworkHarvesterhciIoV1beta1LinkStatus`

GetNetworkLinkStatus returns the NetworkLinkStatus field if non-nil, zero value otherwise.

### GetNetworkLinkStatusOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNetworkLinkStatusOk() (*map[string]NetworkHarvesterhciIoV1beta1LinkStatus, bool)`

GetNetworkLinkStatusOk returns a tuple with the NetworkLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLinkStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) SetNetworkLinkStatus(v map[string]NetworkHarvesterhciIoV1beta1LinkStatus)`

SetNetworkLinkStatus sets NetworkLinkStatus field to given value.

### HasNetworkLinkStatus

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) HasNetworkLinkStatus() bool`

HasNetworkLinkStatus returns a boolean if a field has been set.

### GetNics

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNics() []NetworkHarvesterhciIoV1beta1NetworkInterface`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) GetNicsOk() (*[]NetworkHarvesterhciIoV1beta1NetworkInterface, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) SetNics(v []NetworkHarvesterhciIoV1beta1NetworkInterface)`

SetNics sets Nics field to given value.

### HasNics

`func (o *NetworkHarvesterhciIoV1beta1NodeNetworkStatus) HasNics() bool`

HasNics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


