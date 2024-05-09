# KubevirtIoApiCoreV1Volume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudInitConfigDrive** | Pointer to [**KubevirtIoApiCoreV1CloudInitConfigDriveSource**](KubevirtIoApiCoreV1CloudInitConfigDriveSource.md) | CloudInitConfigDrive represents a cloud-init Config Drive user-data source. The Config Drive data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest. More info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html | [optional] 
**CloudInitNoCloud** | Pointer to [**KubevirtIoApiCoreV1CloudInitNoCloudSource**](KubevirtIoApiCoreV1CloudInitNoCloudSource.md) | CloudInitNoCloud represents a cloud-init NoCloud user-data source. The NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest. More info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html | [optional] 
**ConfigMap** | Pointer to [**KubevirtIoApiCoreV1ConfigMapVolumeSource**](KubevirtIoApiCoreV1ConfigMapVolumeSource.md) | ConfigMapSource represents a reference to a ConfigMap in the same namespace. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/ | [optional] 
**ContainerDisk** | Pointer to [**KubevirtIoApiCoreV1ContainerDiskSource**](KubevirtIoApiCoreV1ContainerDiskSource.md) | ContainerDisk references a docker image, embedding a qcow or raw disk. More info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html | [optional] 
**DataVolume** | Pointer to [**KubevirtIoApiCoreV1DataVolumeSource**](KubevirtIoApiCoreV1DataVolumeSource.md) | DataVolume represents the dynamic creation a PVC for this volume as well as the process of populating that PVC with a disk image. | [optional] 
**DownwardAPI** | Pointer to [**KubevirtIoApiCoreV1DownwardAPIVolumeSource**](KubevirtIoApiCoreV1DownwardAPIVolumeSource.md) | DownwardAPI represents downward API about the pod that should populate this volume | [optional] 
**DownwardMetrics** | Pointer to **map[string]interface{}** | DownwardMetrics adds a very small disk to VMIs which contains a limited view of host and guest metrics. The disk content is compatible with vhostmd (https://github.com/vhostmd/vhostmd) and vm-dump-metrics. | [optional] 
**EmptyDisk** | Pointer to [**KubevirtIoApiCoreV1EmptyDiskSource**](KubevirtIoApiCoreV1EmptyDiskSource.md) | EmptyDisk represents a temporary disk which shares the vmis lifecycle. More info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html | [optional] 
**Ephemeral** | Pointer to [**KubevirtIoApiCoreV1EphemeralVolumeSource**](KubevirtIoApiCoreV1EphemeralVolumeSource.md) | Ephemeral is a special volume source that \&quot;wraps\&quot; specified source and provides copy-on-write image on top of it. | [optional] 
**HostDisk** | Pointer to [**KubevirtIoApiCoreV1HostDisk**](KubevirtIoApiCoreV1HostDisk.md) | HostDisk represents a disk created on the cluster level | [optional] 
**MemoryDump** | Pointer to [**KubevirtIoApiCoreV1MemoryDumpVolumeSource**](KubevirtIoApiCoreV1MemoryDumpVolumeSource.md) | MemoryDump is attached to the virt launcher and is populated with a memory dump of the vmi | [optional] 
**Name** | **string** | Volume&#39;s name. Must be a DNS_LABEL and unique within the vmi. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [default to ""]
**PersistentVolumeClaim** | Pointer to [**KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource**](KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource.md) | PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. Directly attached to the vmi via qemu. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims | [optional] 
**Secret** | Pointer to [**KubevirtIoApiCoreV1SecretVolumeSource**](KubevirtIoApiCoreV1SecretVolumeSource.md) | SecretVolumeSource represents a reference to a secret data in the same namespace. More info: https://kubernetes.io/docs/concepts/configuration/secret/ | [optional] 
**ServiceAccount** | Pointer to [**KubevirtIoApiCoreV1ServiceAccountVolumeSource**](KubevirtIoApiCoreV1ServiceAccountVolumeSource.md) | ServiceAccountVolumeSource represents a reference to a service account. There can only be one volume of this type! More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/ | [optional] 
**Sysprep** | Pointer to [**KubevirtIoApiCoreV1SysprepSource**](KubevirtIoApiCoreV1SysprepSource.md) | Represents a Sysprep volume source. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Volume

`func NewKubevirtIoApiCoreV1Volume(name string, ) *KubevirtIoApiCoreV1Volume`

NewKubevirtIoApiCoreV1Volume instantiates a new KubevirtIoApiCoreV1Volume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VolumeWithDefaults

`func NewKubevirtIoApiCoreV1VolumeWithDefaults() *KubevirtIoApiCoreV1Volume`

NewKubevirtIoApiCoreV1VolumeWithDefaults instantiates a new KubevirtIoApiCoreV1Volume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudInitConfigDrive

`func (o *KubevirtIoApiCoreV1Volume) GetCloudInitConfigDrive() KubevirtIoApiCoreV1CloudInitConfigDriveSource`

GetCloudInitConfigDrive returns the CloudInitConfigDrive field if non-nil, zero value otherwise.

### GetCloudInitConfigDriveOk

`func (o *KubevirtIoApiCoreV1Volume) GetCloudInitConfigDriveOk() (*KubevirtIoApiCoreV1CloudInitConfigDriveSource, bool)`

GetCloudInitConfigDriveOk returns a tuple with the CloudInitConfigDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitConfigDrive

`func (o *KubevirtIoApiCoreV1Volume) SetCloudInitConfigDrive(v KubevirtIoApiCoreV1CloudInitConfigDriveSource)`

SetCloudInitConfigDrive sets CloudInitConfigDrive field to given value.

### HasCloudInitConfigDrive

`func (o *KubevirtIoApiCoreV1Volume) HasCloudInitConfigDrive() bool`

HasCloudInitConfigDrive returns a boolean if a field has been set.

### GetCloudInitNoCloud

`func (o *KubevirtIoApiCoreV1Volume) GetCloudInitNoCloud() KubevirtIoApiCoreV1CloudInitNoCloudSource`

GetCloudInitNoCloud returns the CloudInitNoCloud field if non-nil, zero value otherwise.

### GetCloudInitNoCloudOk

`func (o *KubevirtIoApiCoreV1Volume) GetCloudInitNoCloudOk() (*KubevirtIoApiCoreV1CloudInitNoCloudSource, bool)`

GetCloudInitNoCloudOk returns a tuple with the CloudInitNoCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitNoCloud

`func (o *KubevirtIoApiCoreV1Volume) SetCloudInitNoCloud(v KubevirtIoApiCoreV1CloudInitNoCloudSource)`

SetCloudInitNoCloud sets CloudInitNoCloud field to given value.

### HasCloudInitNoCloud

`func (o *KubevirtIoApiCoreV1Volume) HasCloudInitNoCloud() bool`

HasCloudInitNoCloud returns a boolean if a field has been set.

### GetConfigMap

`func (o *KubevirtIoApiCoreV1Volume) GetConfigMap() KubevirtIoApiCoreV1ConfigMapVolumeSource`

GetConfigMap returns the ConfigMap field if non-nil, zero value otherwise.

### GetConfigMapOk

`func (o *KubevirtIoApiCoreV1Volume) GetConfigMapOk() (*KubevirtIoApiCoreV1ConfigMapVolumeSource, bool)`

GetConfigMapOk returns a tuple with the ConfigMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMap

`func (o *KubevirtIoApiCoreV1Volume) SetConfigMap(v KubevirtIoApiCoreV1ConfigMapVolumeSource)`

SetConfigMap sets ConfigMap field to given value.

### HasConfigMap

`func (o *KubevirtIoApiCoreV1Volume) HasConfigMap() bool`

HasConfigMap returns a boolean if a field has been set.

### GetContainerDisk

`func (o *KubevirtIoApiCoreV1Volume) GetContainerDisk() KubevirtIoApiCoreV1ContainerDiskSource`

GetContainerDisk returns the ContainerDisk field if non-nil, zero value otherwise.

### GetContainerDiskOk

`func (o *KubevirtIoApiCoreV1Volume) GetContainerDiskOk() (*KubevirtIoApiCoreV1ContainerDiskSource, bool)`

GetContainerDiskOk returns a tuple with the ContainerDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDisk

`func (o *KubevirtIoApiCoreV1Volume) SetContainerDisk(v KubevirtIoApiCoreV1ContainerDiskSource)`

SetContainerDisk sets ContainerDisk field to given value.

### HasContainerDisk

`func (o *KubevirtIoApiCoreV1Volume) HasContainerDisk() bool`

HasContainerDisk returns a boolean if a field has been set.

### GetDataVolume

`func (o *KubevirtIoApiCoreV1Volume) GetDataVolume() KubevirtIoApiCoreV1DataVolumeSource`

GetDataVolume returns the DataVolume field if non-nil, zero value otherwise.

### GetDataVolumeOk

`func (o *KubevirtIoApiCoreV1Volume) GetDataVolumeOk() (*KubevirtIoApiCoreV1DataVolumeSource, bool)`

GetDataVolumeOk returns a tuple with the DataVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVolume

`func (o *KubevirtIoApiCoreV1Volume) SetDataVolume(v KubevirtIoApiCoreV1DataVolumeSource)`

SetDataVolume sets DataVolume field to given value.

### HasDataVolume

`func (o *KubevirtIoApiCoreV1Volume) HasDataVolume() bool`

HasDataVolume returns a boolean if a field has been set.

### GetDownwardAPI

`func (o *KubevirtIoApiCoreV1Volume) GetDownwardAPI() KubevirtIoApiCoreV1DownwardAPIVolumeSource`

GetDownwardAPI returns the DownwardAPI field if non-nil, zero value otherwise.

### GetDownwardAPIOk

`func (o *KubevirtIoApiCoreV1Volume) GetDownwardAPIOk() (*KubevirtIoApiCoreV1DownwardAPIVolumeSource, bool)`

GetDownwardAPIOk returns a tuple with the DownwardAPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardAPI

`func (o *KubevirtIoApiCoreV1Volume) SetDownwardAPI(v KubevirtIoApiCoreV1DownwardAPIVolumeSource)`

SetDownwardAPI sets DownwardAPI field to given value.

### HasDownwardAPI

`func (o *KubevirtIoApiCoreV1Volume) HasDownwardAPI() bool`

HasDownwardAPI returns a boolean if a field has been set.

### GetDownwardMetrics

`func (o *KubevirtIoApiCoreV1Volume) GetDownwardMetrics() map[string]interface{}`

GetDownwardMetrics returns the DownwardMetrics field if non-nil, zero value otherwise.

### GetDownwardMetricsOk

`func (o *KubevirtIoApiCoreV1Volume) GetDownwardMetricsOk() (*map[string]interface{}, bool)`

GetDownwardMetricsOk returns a tuple with the DownwardMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownwardMetrics

`func (o *KubevirtIoApiCoreV1Volume) SetDownwardMetrics(v map[string]interface{})`

SetDownwardMetrics sets DownwardMetrics field to given value.

### HasDownwardMetrics

`func (o *KubevirtIoApiCoreV1Volume) HasDownwardMetrics() bool`

HasDownwardMetrics returns a boolean if a field has been set.

### GetEmptyDisk

`func (o *KubevirtIoApiCoreV1Volume) GetEmptyDisk() KubevirtIoApiCoreV1EmptyDiskSource`

GetEmptyDisk returns the EmptyDisk field if non-nil, zero value otherwise.

### GetEmptyDiskOk

`func (o *KubevirtIoApiCoreV1Volume) GetEmptyDiskOk() (*KubevirtIoApiCoreV1EmptyDiskSource, bool)`

GetEmptyDiskOk returns a tuple with the EmptyDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyDisk

`func (o *KubevirtIoApiCoreV1Volume) SetEmptyDisk(v KubevirtIoApiCoreV1EmptyDiskSource)`

SetEmptyDisk sets EmptyDisk field to given value.

### HasEmptyDisk

`func (o *KubevirtIoApiCoreV1Volume) HasEmptyDisk() bool`

HasEmptyDisk returns a boolean if a field has been set.

### GetEphemeral

`func (o *KubevirtIoApiCoreV1Volume) GetEphemeral() KubevirtIoApiCoreV1EphemeralVolumeSource`

GetEphemeral returns the Ephemeral field if non-nil, zero value otherwise.

### GetEphemeralOk

`func (o *KubevirtIoApiCoreV1Volume) GetEphemeralOk() (*KubevirtIoApiCoreV1EphemeralVolumeSource, bool)`

GetEphemeralOk returns a tuple with the Ephemeral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral

`func (o *KubevirtIoApiCoreV1Volume) SetEphemeral(v KubevirtIoApiCoreV1EphemeralVolumeSource)`

SetEphemeral sets Ephemeral field to given value.

### HasEphemeral

`func (o *KubevirtIoApiCoreV1Volume) HasEphemeral() bool`

HasEphemeral returns a boolean if a field has been set.

### GetHostDisk

`func (o *KubevirtIoApiCoreV1Volume) GetHostDisk() KubevirtIoApiCoreV1HostDisk`

GetHostDisk returns the HostDisk field if non-nil, zero value otherwise.

### GetHostDiskOk

`func (o *KubevirtIoApiCoreV1Volume) GetHostDiskOk() (*KubevirtIoApiCoreV1HostDisk, bool)`

GetHostDiskOk returns a tuple with the HostDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDisk

`func (o *KubevirtIoApiCoreV1Volume) SetHostDisk(v KubevirtIoApiCoreV1HostDisk)`

SetHostDisk sets HostDisk field to given value.

### HasHostDisk

`func (o *KubevirtIoApiCoreV1Volume) HasHostDisk() bool`

HasHostDisk returns a boolean if a field has been set.

### GetMemoryDump

`func (o *KubevirtIoApiCoreV1Volume) GetMemoryDump() KubevirtIoApiCoreV1MemoryDumpVolumeSource`

GetMemoryDump returns the MemoryDump field if non-nil, zero value otherwise.

### GetMemoryDumpOk

`func (o *KubevirtIoApiCoreV1Volume) GetMemoryDumpOk() (*KubevirtIoApiCoreV1MemoryDumpVolumeSource, bool)`

GetMemoryDumpOk returns a tuple with the MemoryDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDump

`func (o *KubevirtIoApiCoreV1Volume) SetMemoryDump(v KubevirtIoApiCoreV1MemoryDumpVolumeSource)`

SetMemoryDump sets MemoryDump field to given value.

### HasMemoryDump

`func (o *KubevirtIoApiCoreV1Volume) HasMemoryDump() bool`

HasMemoryDump returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Volume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Volume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Volume) SetName(v string)`

SetName sets Name field to given value.


### GetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1Volume) GetPersistentVolumeClaim() KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource`

GetPersistentVolumeClaim returns the PersistentVolumeClaim field if non-nil, zero value otherwise.

### GetPersistentVolumeClaimOk

`func (o *KubevirtIoApiCoreV1Volume) GetPersistentVolumeClaimOk() (*KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource, bool)`

GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1Volume) SetPersistentVolumeClaim(v KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource)`

SetPersistentVolumeClaim sets PersistentVolumeClaim field to given value.

### HasPersistentVolumeClaim

`func (o *KubevirtIoApiCoreV1Volume) HasPersistentVolumeClaim() bool`

HasPersistentVolumeClaim returns a boolean if a field has been set.

### GetSecret

`func (o *KubevirtIoApiCoreV1Volume) GetSecret() KubevirtIoApiCoreV1SecretVolumeSource`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *KubevirtIoApiCoreV1Volume) GetSecretOk() (*KubevirtIoApiCoreV1SecretVolumeSource, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *KubevirtIoApiCoreV1Volume) SetSecret(v KubevirtIoApiCoreV1SecretVolumeSource)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *KubevirtIoApiCoreV1Volume) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetServiceAccount

`func (o *KubevirtIoApiCoreV1Volume) GetServiceAccount() KubevirtIoApiCoreV1ServiceAccountVolumeSource`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *KubevirtIoApiCoreV1Volume) GetServiceAccountOk() (*KubevirtIoApiCoreV1ServiceAccountVolumeSource, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *KubevirtIoApiCoreV1Volume) SetServiceAccount(v KubevirtIoApiCoreV1ServiceAccountVolumeSource)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *KubevirtIoApiCoreV1Volume) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetSysprep

`func (o *KubevirtIoApiCoreV1Volume) GetSysprep() KubevirtIoApiCoreV1SysprepSource`

GetSysprep returns the Sysprep field if non-nil, zero value otherwise.

### GetSysprepOk

`func (o *KubevirtIoApiCoreV1Volume) GetSysprepOk() (*KubevirtIoApiCoreV1SysprepSource, bool)`

GetSysprepOk returns a tuple with the Sysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysprep

`func (o *KubevirtIoApiCoreV1Volume) SetSysprep(v KubevirtIoApiCoreV1SysprepSource)`

SetSysprep sets Sysprep field to given value.

### HasSysprep

`func (o *KubevirtIoApiCoreV1Volume) HasSysprep() bool`

HasSysprep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


