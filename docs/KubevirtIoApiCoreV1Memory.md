# KubevirtIoApiCoreV1Memory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Guest** | Pointer to **string** | Guest allows to specifying the amount of memory which is visible inside the Guest OS. The Guest must lie between Requests and Limits from the resources section. Defaults to the requested memory in the resources section if not specified. | [optional] 
**Hugepages** | Pointer to [**KubevirtIoApiCoreV1Hugepages**](KubevirtIoApiCoreV1Hugepages.md) | Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory. | [optional] 
**MaxGuest** | Pointer to **string** | MaxGuest allows to specify the maximum amount of memory which is visible inside the Guest OS. The delta between MaxGuest and Guest is the amount of memory that can be hot(un)plugged. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Memory

`func NewKubevirtIoApiCoreV1Memory() *KubevirtIoApiCoreV1Memory`

NewKubevirtIoApiCoreV1Memory instantiates a new KubevirtIoApiCoreV1Memory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1MemoryWithDefaults

`func NewKubevirtIoApiCoreV1MemoryWithDefaults() *KubevirtIoApiCoreV1Memory`

NewKubevirtIoApiCoreV1MemoryWithDefaults instantiates a new KubevirtIoApiCoreV1Memory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuest

`func (o *KubevirtIoApiCoreV1Memory) GetGuest() string`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *KubevirtIoApiCoreV1Memory) GetGuestOk() (*string, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *KubevirtIoApiCoreV1Memory) SetGuest(v string)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *KubevirtIoApiCoreV1Memory) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHugepages

`func (o *KubevirtIoApiCoreV1Memory) GetHugepages() KubevirtIoApiCoreV1Hugepages`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *KubevirtIoApiCoreV1Memory) GetHugepagesOk() (*KubevirtIoApiCoreV1Hugepages, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *KubevirtIoApiCoreV1Memory) SetHugepages(v KubevirtIoApiCoreV1Hugepages)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *KubevirtIoApiCoreV1Memory) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetMaxGuest

`func (o *KubevirtIoApiCoreV1Memory) GetMaxGuest() string`

GetMaxGuest returns the MaxGuest field if non-nil, zero value otherwise.

### GetMaxGuestOk

`func (o *KubevirtIoApiCoreV1Memory) GetMaxGuestOk() (*string, bool)`

GetMaxGuestOk returns a tuple with the MaxGuest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGuest

`func (o *KubevirtIoApiCoreV1Memory) SetMaxGuest(v string)`

SetMaxGuest sets MaxGuest field to given value.

### HasMaxGuest

`func (o *KubevirtIoApiCoreV1Memory) HasMaxGuest() bool`

HasMaxGuest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


