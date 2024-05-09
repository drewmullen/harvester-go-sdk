# KubevirtIoApiCoreV1HostDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | **string** | DeviceName is the resource name of the host device exposed by a device plugin | [default to ""]
**Name** | **string** |  | [default to ""]
**Tag** | Pointer to **string** | If specified, the virtual network interface address and its tag will be provided to the guest via config drive | [optional] 

## Methods

### NewKubevirtIoApiCoreV1HostDevice

`func NewKubevirtIoApiCoreV1HostDevice(deviceName string, name string, ) *KubevirtIoApiCoreV1HostDevice`

NewKubevirtIoApiCoreV1HostDevice instantiates a new KubevirtIoApiCoreV1HostDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1HostDeviceWithDefaults

`func NewKubevirtIoApiCoreV1HostDeviceWithDefaults() *KubevirtIoApiCoreV1HostDevice`

NewKubevirtIoApiCoreV1HostDeviceWithDefaults instantiates a new KubevirtIoApiCoreV1HostDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *KubevirtIoApiCoreV1HostDevice) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *KubevirtIoApiCoreV1HostDevice) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *KubevirtIoApiCoreV1HostDevice) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetName

`func (o *KubevirtIoApiCoreV1HostDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1HostDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1HostDevice) SetName(v string)`

SetName sets Name field to given value.


### GetTag

`func (o *KubevirtIoApiCoreV1HostDevice) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *KubevirtIoApiCoreV1HostDevice) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *KubevirtIoApiCoreV1HostDevice) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *KubevirtIoApiCoreV1HostDevice) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


