# HarvesterhciIoV1beta1VirtualMachineImageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **string** |  | [optional] [default to ""]
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | **string** |  | [default to ""]
**PvcName** | Pointer to **string** |  | [optional] [default to ""]
**PvcNamespace** | Pointer to **string** |  | [optional] [default to ""]
**Retry** | Pointer to **int32** |  | [optional] [default to 0]
**SourceType** | **string** |  | [default to ""]
**StorageClassParameters** | Pointer to **map[string]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewHarvesterhciIoV1beta1VirtualMachineImageSpec

`func NewHarvesterhciIoV1beta1VirtualMachineImageSpec(displayName string, sourceType string, ) *HarvesterhciIoV1beta1VirtualMachineImageSpec`

NewHarvesterhciIoV1beta1VirtualMachineImageSpec instantiates a new HarvesterhciIoV1beta1VirtualMachineImageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHarvesterhciIoV1beta1VirtualMachineImageSpecWithDefaults

`func NewHarvesterhciIoV1beta1VirtualMachineImageSpecWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImageSpec`

NewHarvesterhciIoV1beta1VirtualMachineImageSpecWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPvcName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetPvcName() string`

GetPvcName returns the PvcName field if non-nil, zero value otherwise.

### GetPvcNameOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetPvcNameOk() (*string, bool)`

GetPvcNameOk returns a tuple with the PvcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetPvcName(v string)`

SetPvcName sets PvcName field to given value.

### HasPvcName

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasPvcName() bool`

HasPvcName returns a boolean if a field has been set.

### GetPvcNamespace

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetPvcNamespace() string`

GetPvcNamespace returns the PvcNamespace field if non-nil, zero value otherwise.

### GetPvcNamespaceOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetPvcNamespaceOk() (*string, bool)`

GetPvcNamespaceOk returns a tuple with the PvcNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcNamespace

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetPvcNamespace(v string)`

SetPvcNamespace sets PvcNamespace field to given value.

### HasPvcNamespace

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasPvcNamespace() bool`

HasPvcNamespace returns a boolean if a field has been set.

### GetRetry

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### GetSourceType

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetStorageClassParameters

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetStorageClassParameters() map[string]string`

GetStorageClassParameters returns the StorageClassParameters field if non-nil, zero value otherwise.

### GetStorageClassParametersOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetStorageClassParametersOk() (*map[string]string, bool)`

GetStorageClassParametersOk returns a tuple with the StorageClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassParameters

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetStorageClassParameters(v map[string]string)`

SetStorageClassParameters sets StorageClassParameters field to given value.

### HasStorageClassParameters

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasStorageClassParameters() bool`

HasStorageClassParameters returns a boolean if a field has been set.

### GetUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HarvesterhciIoV1beta1VirtualMachineImageSpec) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


