# KubevirtIoApiCoreV1DomainSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chassis** | Pointer to [**KubevirtIoApiCoreV1Chassis**](KubevirtIoApiCoreV1Chassis.md) | Chassis specifies the chassis info passed to the domain. | [optional] 
**Clock** | Pointer to [**KubevirtIoApiCoreV1Clock**](KubevirtIoApiCoreV1Clock.md) | Clock sets the clock and timers of the vmi. | [optional] 
**Cpu** | Pointer to [**KubevirtIoApiCoreV1CPU**](KubevirtIoApiCoreV1CPU.md) | CPU allow specified the detailed CPU topology inside the vmi. | [optional] 
**Devices** | [**KubevirtIoApiCoreV1Devices**](KubevirtIoApiCoreV1Devices.md) | Devices allows adding disks, network interfaces, and others | [default to {}]
**Features** | Pointer to [**KubevirtIoApiCoreV1Features**](KubevirtIoApiCoreV1Features.md) | Features like acpi, apic, hyperv, smm. | [optional] 
**Firmware** | Pointer to [**KubevirtIoApiCoreV1Firmware**](KubevirtIoApiCoreV1Firmware.md) | Firmware. | [optional] 
**IoThreadsPolicy** | Pointer to **string** | Controls whether or not disks will share IOThreads. Omitting IOThreadsPolicy disables use of IOThreads. One of: shared, auto | [optional] 
**LaunchSecurity** | Pointer to [**KubevirtIoApiCoreV1LaunchSecurity**](KubevirtIoApiCoreV1LaunchSecurity.md) | Launch Security setting of the vmi. | [optional] 
**Machine** | Pointer to [**KubevirtIoApiCoreV1Machine**](KubevirtIoApiCoreV1Machine.md) | Machine type. | [optional] 
**Memory** | Pointer to [**KubevirtIoApiCoreV1Memory**](KubevirtIoApiCoreV1Memory.md) | Memory allow specifying the VMI memory features. | [optional] 
**Resources** | Pointer to [**KubevirtIoApiCoreV1ResourceRequirements**](KubevirtIoApiCoreV1ResourceRequirements.md) | Resources describes the Compute Resources required by this vmi. | [optional] [default to {}]

## Methods

### NewKubevirtIoApiCoreV1DomainSpec

`func NewKubevirtIoApiCoreV1DomainSpec(devices KubevirtIoApiCoreV1Devices, ) *KubevirtIoApiCoreV1DomainSpec`

NewKubevirtIoApiCoreV1DomainSpec instantiates a new KubevirtIoApiCoreV1DomainSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DomainSpecWithDefaults

`func NewKubevirtIoApiCoreV1DomainSpecWithDefaults() *KubevirtIoApiCoreV1DomainSpec`

NewKubevirtIoApiCoreV1DomainSpecWithDefaults instantiates a new KubevirtIoApiCoreV1DomainSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChassis

`func (o *KubevirtIoApiCoreV1DomainSpec) GetChassis() KubevirtIoApiCoreV1Chassis`

GetChassis returns the Chassis field if non-nil, zero value otherwise.

### GetChassisOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetChassisOk() (*KubevirtIoApiCoreV1Chassis, bool)`

GetChassisOk returns a tuple with the Chassis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassis

`func (o *KubevirtIoApiCoreV1DomainSpec) SetChassis(v KubevirtIoApiCoreV1Chassis)`

SetChassis sets Chassis field to given value.

### HasChassis

`func (o *KubevirtIoApiCoreV1DomainSpec) HasChassis() bool`

HasChassis returns a boolean if a field has been set.

### GetClock

`func (o *KubevirtIoApiCoreV1DomainSpec) GetClock() KubevirtIoApiCoreV1Clock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetClockOk() (*KubevirtIoApiCoreV1Clock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *KubevirtIoApiCoreV1DomainSpec) SetClock(v KubevirtIoApiCoreV1Clock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *KubevirtIoApiCoreV1DomainSpec) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetCpu

`func (o *KubevirtIoApiCoreV1DomainSpec) GetCpu() KubevirtIoApiCoreV1CPU`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetCpuOk() (*KubevirtIoApiCoreV1CPU, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *KubevirtIoApiCoreV1DomainSpec) SetCpu(v KubevirtIoApiCoreV1CPU)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *KubevirtIoApiCoreV1DomainSpec) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDevices

`func (o *KubevirtIoApiCoreV1DomainSpec) GetDevices() KubevirtIoApiCoreV1Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetDevicesOk() (*KubevirtIoApiCoreV1Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *KubevirtIoApiCoreV1DomainSpec) SetDevices(v KubevirtIoApiCoreV1Devices)`

SetDevices sets Devices field to given value.


### GetFeatures

`func (o *KubevirtIoApiCoreV1DomainSpec) GetFeatures() KubevirtIoApiCoreV1Features`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetFeaturesOk() (*KubevirtIoApiCoreV1Features, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *KubevirtIoApiCoreV1DomainSpec) SetFeatures(v KubevirtIoApiCoreV1Features)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *KubevirtIoApiCoreV1DomainSpec) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFirmware

`func (o *KubevirtIoApiCoreV1DomainSpec) GetFirmware() KubevirtIoApiCoreV1Firmware`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetFirmwareOk() (*KubevirtIoApiCoreV1Firmware, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *KubevirtIoApiCoreV1DomainSpec) SetFirmware(v KubevirtIoApiCoreV1Firmware)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *KubevirtIoApiCoreV1DomainSpec) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetIoThreadsPolicy

`func (o *KubevirtIoApiCoreV1DomainSpec) GetIoThreadsPolicy() string`

GetIoThreadsPolicy returns the IoThreadsPolicy field if non-nil, zero value otherwise.

### GetIoThreadsPolicyOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetIoThreadsPolicyOk() (*string, bool)`

GetIoThreadsPolicyOk returns a tuple with the IoThreadsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIoThreadsPolicy

`func (o *KubevirtIoApiCoreV1DomainSpec) SetIoThreadsPolicy(v string)`

SetIoThreadsPolicy sets IoThreadsPolicy field to given value.

### HasIoThreadsPolicy

`func (o *KubevirtIoApiCoreV1DomainSpec) HasIoThreadsPolicy() bool`

HasIoThreadsPolicy returns a boolean if a field has been set.

### GetLaunchSecurity

`func (o *KubevirtIoApiCoreV1DomainSpec) GetLaunchSecurity() KubevirtIoApiCoreV1LaunchSecurity`

GetLaunchSecurity returns the LaunchSecurity field if non-nil, zero value otherwise.

### GetLaunchSecurityOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetLaunchSecurityOk() (*KubevirtIoApiCoreV1LaunchSecurity, bool)`

GetLaunchSecurityOk returns a tuple with the LaunchSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchSecurity

`func (o *KubevirtIoApiCoreV1DomainSpec) SetLaunchSecurity(v KubevirtIoApiCoreV1LaunchSecurity)`

SetLaunchSecurity sets LaunchSecurity field to given value.

### HasLaunchSecurity

`func (o *KubevirtIoApiCoreV1DomainSpec) HasLaunchSecurity() bool`

HasLaunchSecurity returns a boolean if a field has been set.

### GetMachine

`func (o *KubevirtIoApiCoreV1DomainSpec) GetMachine() KubevirtIoApiCoreV1Machine`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetMachineOk() (*KubevirtIoApiCoreV1Machine, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *KubevirtIoApiCoreV1DomainSpec) SetMachine(v KubevirtIoApiCoreV1Machine)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *KubevirtIoApiCoreV1DomainSpec) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *KubevirtIoApiCoreV1DomainSpec) GetMemory() KubevirtIoApiCoreV1Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetMemoryOk() (*KubevirtIoApiCoreV1Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *KubevirtIoApiCoreV1DomainSpec) SetMemory(v KubevirtIoApiCoreV1Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *KubevirtIoApiCoreV1DomainSpec) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetResources

`func (o *KubevirtIoApiCoreV1DomainSpec) GetResources() KubevirtIoApiCoreV1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *KubevirtIoApiCoreV1DomainSpec) GetResourcesOk() (*KubevirtIoApiCoreV1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *KubevirtIoApiCoreV1DomainSpec) SetResources(v KubevirtIoApiCoreV1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *KubevirtIoApiCoreV1DomainSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


