# NetworkHarvesterhciIoV1beta1NetworkInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Index of the NIC | [default to 0]
**MasterIndex** | Pointer to **int32** | Index of the NIC&#39;s master | [optional] 
**Name** | **string** | Name of the NIC | [default to ""]
**State** | **string** | State of the NIC, up/down/unknown | [default to ""]
**Type** | **string** | Interface type of the NIC | [default to ""]
**UsedByManagementNetwork** | Pointer to **bool** | Specify whether used by management network or not | [optional] 
**UsedByVlanNetwork** | Pointer to **bool** | Specify whether used by VLAN network or not | [optional] 

## Methods

### NewNetworkHarvesterhciIoV1beta1NetworkInterface

`func NewNetworkHarvesterhciIoV1beta1NetworkInterface(index int32, name string, state string, type_ string, ) *NetworkHarvesterhciIoV1beta1NetworkInterface`

NewNetworkHarvesterhciIoV1beta1NetworkInterface instantiates a new NetworkHarvesterhciIoV1beta1NetworkInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkHarvesterhciIoV1beta1NetworkInterfaceWithDefaults

`func NewNetworkHarvesterhciIoV1beta1NetworkInterfaceWithDefaults() *NetworkHarvesterhciIoV1beta1NetworkInterface`

NewNetworkHarvesterhciIoV1beta1NetworkInterfaceWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1NetworkInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMasterIndex

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetMasterIndex() int32`

GetMasterIndex returns the MasterIndex field if non-nil, zero value otherwise.

### GetMasterIndexOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetMasterIndexOk() (*int32, bool)`

GetMasterIndexOk returns a tuple with the MasterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterIndex

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetMasterIndex(v int32)`

SetMasterIndex sets MasterIndex field to given value.

### HasMasterIndex

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) HasMasterIndex() bool`

HasMasterIndex returns a boolean if a field has been set.

### GetName

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetType(v string)`

SetType sets Type field to given value.


### GetUsedByManagementNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetUsedByManagementNetwork() bool`

GetUsedByManagementNetwork returns the UsedByManagementNetwork field if non-nil, zero value otherwise.

### GetUsedByManagementNetworkOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetUsedByManagementNetworkOk() (*bool, bool)`

GetUsedByManagementNetworkOk returns a tuple with the UsedByManagementNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByManagementNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetUsedByManagementNetwork(v bool)`

SetUsedByManagementNetwork sets UsedByManagementNetwork field to given value.

### HasUsedByManagementNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) HasUsedByManagementNetwork() bool`

HasUsedByManagementNetwork returns a boolean if a field has been set.

### GetUsedByVlanNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetUsedByVlanNetwork() bool`

GetUsedByVlanNetwork returns the UsedByVlanNetwork field if non-nil, zero value otherwise.

### GetUsedByVlanNetworkOk

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) GetUsedByVlanNetworkOk() (*bool, bool)`

GetUsedByVlanNetworkOk returns a tuple with the UsedByVlanNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByVlanNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) SetUsedByVlanNetwork(v bool)`

SetUsedByVlanNetwork sets UsedByVlanNetwork field to given value.

### HasUsedByVlanNetwork

`func (o *NetworkHarvesterhciIoV1beta1NetworkInterface) HasUsedByVlanNetwork() bool`

HasUsedByVlanNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


