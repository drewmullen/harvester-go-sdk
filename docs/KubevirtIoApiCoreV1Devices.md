# KubevirtIoApiCoreV1Devices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoattachGraphicsDevice** | Pointer to **bool** | Whether to attach the default graphics device or not. VNC will not be available if set to false. Defaults to true. | [optional] 
**AutoattachInputDevice** | Pointer to **bool** | Whether to attach an Input Device. Defaults to false. | [optional] 
**AutoattachMemBalloon** | Pointer to **bool** | Whether to attach the Memory balloon device with default period. Period can be adjusted in virt-config. Defaults to true. | [optional] 
**AutoattachPodInterface** | Pointer to **bool** | Whether to attach a pod network interface. Defaults to true. | [optional] 
**AutoattachSerialConsole** | Pointer to **bool** | Whether to attach the default virtio-serial console or not. Serial console access will not be available if set to false. Defaults to true. | [optional] 
**AutoattachVSOCK** | Pointer to **bool** | Whether to attach the VSOCK CID to the VM or not. VSOCK access will be available if set to true. Defaults to false. | [optional] 
**BlockMultiQueue** | Pointer to **bool** | Whether or not to enable virtio multi-queue for block devices. Defaults to false. | [optional] 
**ClientPassthrough** | Pointer to **map[string]interface{}** | To configure and access client devices such as redirecting USB | [optional] 
**DisableHotplug** | Pointer to **bool** | DisableHotplug disabled the ability to hotplug disks. | [optional] 
**Disks** | Pointer to [**[]KubevirtIoApiCoreV1Disk**](KubevirtIoApiCoreV1Disk.md) | Disks describes disks, cdroms and luns which are connected to the vmi. | [optional] 
**DownwardMetrics** | Pointer to **map[string]interface{}** | DownwardMetrics creates a virtio serials for exposing the downward metrics to the vmi. | [optional] 
**Filesystems** | Pointer to [**[]KubevirtIoApiCoreV1Filesystem**](KubevirtIoApiCoreV1Filesystem.md) | Filesystems describes filesystem which is connected to the vmi. | [optional] 
**Gpus** | Pointer to [**[]KubevirtIoApiCoreV1GPU**](KubevirtIoApiCoreV1GPU.md) | Whether to attach a GPU device to the vmi. | [optional] 
**HostDevices** | Pointer to [**[]KubevirtIoApiCoreV1HostDevice**](KubevirtIoApiCoreV1HostDevice.md) | Whether to attach a host device to the vmi. | [optional] 
**Inputs** | Pointer to [**[]KubevirtIoApiCoreV1Input**](KubevirtIoApiCoreV1Input.md) | Inputs describe input devices | [optional] 
**Interfaces** | Pointer to [**[]KubevirtIoApiCoreV1Interface**](KubevirtIoApiCoreV1Interface.md) | Interfaces describe network interfaces which are added to the vmi. | [optional] 
**LogSerialConsole** | Pointer to **bool** | Whether to log the auto-attached default serial console or not. Serial console logs will be collect to a file and then streamed from a named &#x60;guest-console-log&#x60;. Not relevant if autoattachSerialConsole is disabled. Defaults to cluster wide setting on VirtualMachineOptions. | [optional] 
**NetworkInterfaceMultiqueue** | Pointer to **bool** | If specified, virtual network interfaces configured with a virtio bus will also enable the vhost multiqueue feature for network devices. The number of queues created depends on additional factors of the VirtualMachineInstance, like the number of guest CPUs. | [optional] 
**Rng** | Pointer to **map[string]interface{}** | Whether to have random number generator from host | [optional] 
**Sound** | Pointer to [**KubevirtIoApiCoreV1SoundDevice**](KubevirtIoApiCoreV1SoundDevice.md) | Whether to emulate a sound device. | [optional] 
**Tpm** | Pointer to [**KubevirtIoApiCoreV1TPMDevice**](KubevirtIoApiCoreV1TPMDevice.md) | Whether to emulate a TPM device. | [optional] 
**UseVirtioTransitional** | Pointer to **bool** | Fall back to legacy virtio 0.9 support if virtio bus is selected on devices. This is helpful for old machines like CentOS6 or RHEL6 which do not understand virtio_non_transitional (virtio 1.0). | [optional] 
**Watchdog** | Pointer to [**KubevirtIoApiCoreV1Watchdog**](KubevirtIoApiCoreV1Watchdog.md) | Watchdog describes a watchdog device which can be added to the vmi. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Devices

`func NewKubevirtIoApiCoreV1Devices() *KubevirtIoApiCoreV1Devices`

NewKubevirtIoApiCoreV1Devices instantiates a new KubevirtIoApiCoreV1Devices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1DevicesWithDefaults

`func NewKubevirtIoApiCoreV1DevicesWithDefaults() *KubevirtIoApiCoreV1Devices`

NewKubevirtIoApiCoreV1DevicesWithDefaults instantiates a new KubevirtIoApiCoreV1Devices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoattachGraphicsDevice

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachGraphicsDevice() bool`

GetAutoattachGraphicsDevice returns the AutoattachGraphicsDevice field if non-nil, zero value otherwise.

### GetAutoattachGraphicsDeviceOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachGraphicsDeviceOk() (*bool, bool)`

GetAutoattachGraphicsDeviceOk returns a tuple with the AutoattachGraphicsDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachGraphicsDevice

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachGraphicsDevice(v bool)`

SetAutoattachGraphicsDevice sets AutoattachGraphicsDevice field to given value.

### HasAutoattachGraphicsDevice

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachGraphicsDevice() bool`

HasAutoattachGraphicsDevice returns a boolean if a field has been set.

### GetAutoattachInputDevice

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachInputDevice() bool`

GetAutoattachInputDevice returns the AutoattachInputDevice field if non-nil, zero value otherwise.

### GetAutoattachInputDeviceOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachInputDeviceOk() (*bool, bool)`

GetAutoattachInputDeviceOk returns a tuple with the AutoattachInputDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachInputDevice

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachInputDevice(v bool)`

SetAutoattachInputDevice sets AutoattachInputDevice field to given value.

### HasAutoattachInputDevice

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachInputDevice() bool`

HasAutoattachInputDevice returns a boolean if a field has been set.

### GetAutoattachMemBalloon

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachMemBalloon() bool`

GetAutoattachMemBalloon returns the AutoattachMemBalloon field if non-nil, zero value otherwise.

### GetAutoattachMemBalloonOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachMemBalloonOk() (*bool, bool)`

GetAutoattachMemBalloonOk returns a tuple with the AutoattachMemBalloon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachMemBalloon

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachMemBalloon(v bool)`

SetAutoattachMemBalloon sets AutoattachMemBalloon field to given value.

### HasAutoattachMemBalloon

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachMemBalloon() bool`

HasAutoattachMemBalloon returns a boolean if a field has been set.

### GetAutoattachPodInterface

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachPodInterface() bool`

GetAutoattachPodInterface returns the AutoattachPodInterface field if non-nil, zero value otherwise.

### GetAutoattachPodInterfaceOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachPodInterfaceOk() (*bool, bool)`

GetAutoattachPodInterfaceOk returns a tuple with the AutoattachPodInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachPodInterface

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachPodInterface(v bool)`

SetAutoattachPodInterface sets AutoattachPodInterface field to given value.

### HasAutoattachPodInterface

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachPodInterface() bool`

HasAutoattachPodInterface returns a boolean if a field has been set.

### GetAutoattachSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachSerialConsole() bool`

GetAutoattachSerialConsole returns the AutoattachSerialConsole field if non-nil, zero value otherwise.

### GetAutoattachSerialConsoleOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachSerialConsoleOk() (*bool, bool)`

GetAutoattachSerialConsoleOk returns a tuple with the AutoattachSerialConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachSerialConsole(v bool)`

SetAutoattachSerialConsole sets AutoattachSerialConsole field to given value.

### HasAutoattachSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachSerialConsole() bool`

HasAutoattachSerialConsole returns a boolean if a field has been set.

### GetAutoattachVSOCK

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachVSOCK() bool`

GetAutoattachVSOCK returns the AutoattachVSOCK field if non-nil, zero value otherwise.

### GetAutoattachVSOCKOk

`func (o *KubevirtIoApiCoreV1Devices) GetAutoattachVSOCKOk() (*bool, bool)`

GetAutoattachVSOCKOk returns a tuple with the AutoattachVSOCK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoattachVSOCK

`func (o *KubevirtIoApiCoreV1Devices) SetAutoattachVSOCK(v bool)`

SetAutoattachVSOCK sets AutoattachVSOCK field to given value.

### HasAutoattachVSOCK

`func (o *KubevirtIoApiCoreV1Devices) HasAutoattachVSOCK() bool`

HasAutoattachVSOCK returns a boolean if a field has been set.

### GetBlockMultiQueue

`func (o *KubevirtIoApiCoreV1Devices) GetBlockMultiQueue() bool`

GetBlockMultiQueue returns the BlockMultiQueue field if non-nil, zero value otherwise.

### GetBlockMultiQueueOk

`func (o *KubevirtIoApiCoreV1Devices) GetBlockMultiQueueOk() (*bool, bool)`

GetBlockMultiQueueOk returns a tuple with the BlockMultiQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockMultiQueue

`func (o *KubevirtIoApiCoreV1Devices) SetBlockMultiQueue(v bool)`

SetBlockMultiQueue sets BlockMultiQueue field to given value.

### HasBlockMultiQueue

`func (o *KubevirtIoApiCoreV1Devices) HasBlockMultiQueue() bool`

HasBlockMultiQueue returns a boolean if a field has been set.

### GetClientPassthrough

`func (o *KubevirtIoApiCoreV1Devices) GetClientPassthrough() map[string]interface{}`

GetClientPassthrough returns the ClientPassthrough field if non-nil, zero value otherwise.

### GetClientPassthroughOk

`func (o *KubevirtIoApiCoreV1Devices) GetClientPassthroughOk() (*map[string]interface{}, bool)`

GetClientPassthroughOk returns a tuple with the ClientPassthrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPassthrough

`func (o *KubevirtIoApiCoreV1Devices) SetClientPassthrough(v map[string]interface{})`

SetClientPassthrough sets ClientPassthrough field to given value.

### HasClientPassthrough

`func (o *KubevirtIoApiCoreV1Devices) HasClientPassthrough() bool`

HasClientPassthrough returns a boolean if a field has been set.

### GetDisableHotplug

`func (o *KubevirtIoApiCoreV1Devices) GetDisableHotplug() bool`

GetDisableHotplug returns the DisableHotplug field if non-nil, zero value otherwise.

### GetDisableHotplugOk

`func (o *KubevirtIoApiCoreV1Devices) GetDisableHotplugOk() (*bool, bool)`

GetDisableHotplugOk returns a tuple with the DisableHotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHotplug

`func (o *KubevirtIoApiCoreV1Devices) SetDisableHotplug(v bool)`

SetDisableHotplug sets DisableHotplug field to given value.

### HasDisableHotplug

`func (o *KubevirtIoApiCoreV1Devices) HasDisableHotplug() bool`

HasDisableHotplug returns a boolean if a field has been set.

### GetDisks

`func (o *KubevirtIoApiCoreV1Devices) GetDisks() []KubevirtIoApiCoreV1Disk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *KubevirtIoApiCoreV1Devices) GetDisksOk() (*[]KubevirtIoApiCoreV1Disk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *KubevirtIoApiCoreV1Devices) SetDisks(v []KubevirtIoApiCoreV1Disk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *KubevirtIoApiCoreV1Devices) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetDownwardMetrics

`func (o *KubevirtIoApiCoreV1Devices) GetDownwardMetrics() map[string]interface{}`

GetDownwardMetrics returns the DownwardMetrics field if non-nil, zero value otherwise.

### GetDownwardMetricsOk

`func (o *KubevirtIoApiCoreV1Devices) GetDownwardMetricsOk() (*map[string]interface{}, bool)`

GetDownwardMetricsOk returns a tuple with the DownwardMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardMetrics

`func (o *KubevirtIoApiCoreV1Devices) SetDownwardMetrics(v map[string]interface{})`

SetDownwardMetrics sets DownwardMetrics field to given value.

### HasDownwardMetrics

`func (o *KubevirtIoApiCoreV1Devices) HasDownwardMetrics() bool`

HasDownwardMetrics returns a boolean if a field has been set.

### GetFilesystems

`func (o *KubevirtIoApiCoreV1Devices) GetFilesystems() []KubevirtIoApiCoreV1Filesystem`

GetFilesystems returns the Filesystems field if non-nil, zero value otherwise.

### GetFilesystemsOk

`func (o *KubevirtIoApiCoreV1Devices) GetFilesystemsOk() (*[]KubevirtIoApiCoreV1Filesystem, bool)`

GetFilesystemsOk returns a tuple with the Filesystems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystems

`func (o *KubevirtIoApiCoreV1Devices) SetFilesystems(v []KubevirtIoApiCoreV1Filesystem)`

SetFilesystems sets Filesystems field to given value.

### HasFilesystems

`func (o *KubevirtIoApiCoreV1Devices) HasFilesystems() bool`

HasFilesystems returns a boolean if a field has been set.

### GetGpus

`func (o *KubevirtIoApiCoreV1Devices) GetGpus() []KubevirtIoApiCoreV1GPU`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *KubevirtIoApiCoreV1Devices) GetGpusOk() (*[]KubevirtIoApiCoreV1GPU, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *KubevirtIoApiCoreV1Devices) SetGpus(v []KubevirtIoApiCoreV1GPU)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *KubevirtIoApiCoreV1Devices) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetHostDevices

`func (o *KubevirtIoApiCoreV1Devices) GetHostDevices() []KubevirtIoApiCoreV1HostDevice`

GetHostDevices returns the HostDevices field if non-nil, zero value otherwise.

### GetHostDevicesOk

`func (o *KubevirtIoApiCoreV1Devices) GetHostDevicesOk() (*[]KubevirtIoApiCoreV1HostDevice, bool)`

GetHostDevicesOk returns a tuple with the HostDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevices

`func (o *KubevirtIoApiCoreV1Devices) SetHostDevices(v []KubevirtIoApiCoreV1HostDevice)`

SetHostDevices sets HostDevices field to given value.

### HasHostDevices

`func (o *KubevirtIoApiCoreV1Devices) HasHostDevices() bool`

HasHostDevices returns a boolean if a field has been set.

### GetInputs

`func (o *KubevirtIoApiCoreV1Devices) GetInputs() []KubevirtIoApiCoreV1Input`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *KubevirtIoApiCoreV1Devices) GetInputsOk() (*[]KubevirtIoApiCoreV1Input, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *KubevirtIoApiCoreV1Devices) SetInputs(v []KubevirtIoApiCoreV1Input)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *KubevirtIoApiCoreV1Devices) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetInterfaces

`func (o *KubevirtIoApiCoreV1Devices) GetInterfaces() []KubevirtIoApiCoreV1Interface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *KubevirtIoApiCoreV1Devices) GetInterfacesOk() (*[]KubevirtIoApiCoreV1Interface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *KubevirtIoApiCoreV1Devices) SetInterfaces(v []KubevirtIoApiCoreV1Interface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *KubevirtIoApiCoreV1Devices) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLogSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) GetLogSerialConsole() bool`

GetLogSerialConsole returns the LogSerialConsole field if non-nil, zero value otherwise.

### GetLogSerialConsoleOk

`func (o *KubevirtIoApiCoreV1Devices) GetLogSerialConsoleOk() (*bool, bool)`

GetLogSerialConsoleOk returns a tuple with the LogSerialConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) SetLogSerialConsole(v bool)`

SetLogSerialConsole sets LogSerialConsole field to given value.

### HasLogSerialConsole

`func (o *KubevirtIoApiCoreV1Devices) HasLogSerialConsole() bool`

HasLogSerialConsole returns a boolean if a field has been set.

### GetNetworkInterfaceMultiqueue

`func (o *KubevirtIoApiCoreV1Devices) GetNetworkInterfaceMultiqueue() bool`

GetNetworkInterfaceMultiqueue returns the NetworkInterfaceMultiqueue field if non-nil, zero value otherwise.

### GetNetworkInterfaceMultiqueueOk

`func (o *KubevirtIoApiCoreV1Devices) GetNetworkInterfaceMultiqueueOk() (*bool, bool)`

GetNetworkInterfaceMultiqueueOk returns a tuple with the NetworkInterfaceMultiqueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceMultiqueue

`func (o *KubevirtIoApiCoreV1Devices) SetNetworkInterfaceMultiqueue(v bool)`

SetNetworkInterfaceMultiqueue sets NetworkInterfaceMultiqueue field to given value.

### HasNetworkInterfaceMultiqueue

`func (o *KubevirtIoApiCoreV1Devices) HasNetworkInterfaceMultiqueue() bool`

HasNetworkInterfaceMultiqueue returns a boolean if a field has been set.

### GetRng

`func (o *KubevirtIoApiCoreV1Devices) GetRng() map[string]interface{}`

GetRng returns the Rng field if non-nil, zero value otherwise.

### GetRngOk

`func (o *KubevirtIoApiCoreV1Devices) GetRngOk() (*map[string]interface{}, bool)`

GetRngOk returns a tuple with the Rng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng

`func (o *KubevirtIoApiCoreV1Devices) SetRng(v map[string]interface{})`

SetRng sets Rng field to given value.

### HasRng

`func (o *KubevirtIoApiCoreV1Devices) HasRng() bool`

HasRng returns a boolean if a field has been set.

### GetSound

`func (o *KubevirtIoApiCoreV1Devices) GetSound() KubevirtIoApiCoreV1SoundDevice`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *KubevirtIoApiCoreV1Devices) GetSoundOk() (*KubevirtIoApiCoreV1SoundDevice, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *KubevirtIoApiCoreV1Devices) SetSound(v KubevirtIoApiCoreV1SoundDevice)`

SetSound sets Sound field to given value.

### HasSound

`func (o *KubevirtIoApiCoreV1Devices) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetTpm

`func (o *KubevirtIoApiCoreV1Devices) GetTpm() KubevirtIoApiCoreV1TPMDevice`

GetTpm returns the Tpm field if non-nil, zero value otherwise.

### GetTpmOk

`func (o *KubevirtIoApiCoreV1Devices) GetTpmOk() (*KubevirtIoApiCoreV1TPMDevice, bool)`

GetTpmOk returns a tuple with the Tpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpm

`func (o *KubevirtIoApiCoreV1Devices) SetTpm(v KubevirtIoApiCoreV1TPMDevice)`

SetTpm sets Tpm field to given value.

### HasTpm

`func (o *KubevirtIoApiCoreV1Devices) HasTpm() bool`

HasTpm returns a boolean if a field has been set.

### GetUseVirtioTransitional

`func (o *KubevirtIoApiCoreV1Devices) GetUseVirtioTransitional() bool`

GetUseVirtioTransitional returns the UseVirtioTransitional field if non-nil, zero value otherwise.

### GetUseVirtioTransitionalOk

`func (o *KubevirtIoApiCoreV1Devices) GetUseVirtioTransitionalOk() (*bool, bool)`

GetUseVirtioTransitionalOk returns a tuple with the UseVirtioTransitional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVirtioTransitional

`func (o *KubevirtIoApiCoreV1Devices) SetUseVirtioTransitional(v bool)`

SetUseVirtioTransitional sets UseVirtioTransitional field to given value.

### HasUseVirtioTransitional

`func (o *KubevirtIoApiCoreV1Devices) HasUseVirtioTransitional() bool`

HasUseVirtioTransitional returns a boolean if a field has been set.

### GetWatchdog

`func (o *KubevirtIoApiCoreV1Devices) GetWatchdog() KubevirtIoApiCoreV1Watchdog`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *KubevirtIoApiCoreV1Devices) GetWatchdogOk() (*KubevirtIoApiCoreV1Watchdog, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *KubevirtIoApiCoreV1Devices) SetWatchdog(v KubevirtIoApiCoreV1Watchdog)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *KubevirtIoApiCoreV1Devices) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


