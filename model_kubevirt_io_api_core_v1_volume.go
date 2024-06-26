/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoApiCoreV1Volume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1Volume{}

// KubevirtIoApiCoreV1Volume Volume represents a named volume in a vmi.
type KubevirtIoApiCoreV1Volume struct {
	// CloudInitConfigDrive represents a cloud-init Config Drive user-data source. The Config Drive data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest. More info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html
	CloudInitConfigDrive *KubevirtIoApiCoreV1CloudInitConfigDriveSource `json:"cloudInitConfigDrive,omitempty"`
	// CloudInitNoCloud represents a cloud-init NoCloud user-data source. The NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest. More info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html
	CloudInitNoCloud *KubevirtIoApiCoreV1CloudInitNoCloudSource `json:"cloudInitNoCloud,omitempty"`
	// ConfigMapSource represents a reference to a ConfigMap in the same namespace. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/
	ConfigMap *KubevirtIoApiCoreV1ConfigMapVolumeSource `json:"configMap,omitempty"`
	// ContainerDisk references a docker image, embedding a qcow or raw disk. More info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html
	ContainerDisk *KubevirtIoApiCoreV1ContainerDiskSource `json:"containerDisk,omitempty"`
	// DataVolume represents the dynamic creation a PVC for this volume as well as the process of populating that PVC with a disk image.
	DataVolume *KubevirtIoApiCoreV1DataVolumeSource `json:"dataVolume,omitempty"`
	// DownwardAPI represents downward API about the pod that should populate this volume
	DownwardAPI *KubevirtIoApiCoreV1DownwardAPIVolumeSource `json:"downwardAPI,omitempty"`
	// DownwardMetrics adds a very small disk to VMIs which contains a limited view of host and guest metrics. The disk content is compatible with vhostmd (https://github.com/vhostmd/vhostmd) and vm-dump-metrics.
	DownwardMetrics map[string]interface{} `json:"downwardMetrics,omitempty"`
	// EmptyDisk represents a temporary disk which shares the vmis lifecycle. More info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html
	EmptyDisk *KubevirtIoApiCoreV1EmptyDiskSource `json:"emptyDisk,omitempty"`
	// Ephemeral is a special volume source that \"wraps\" specified source and provides copy-on-write image on top of it.
	Ephemeral *KubevirtIoApiCoreV1EphemeralVolumeSource `json:"ephemeral,omitempty"`
	// HostDisk represents a disk created on the cluster level
	HostDisk *KubevirtIoApiCoreV1HostDisk `json:"hostDisk,omitempty"`
	// MemoryDump is attached to the virt launcher and is populated with a memory dump of the vmi
	MemoryDump *KubevirtIoApiCoreV1MemoryDumpVolumeSource `json:"memoryDump,omitempty"`
	// Volume's name. Must be a DNS_LABEL and unique within the vmi. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name"`
	// PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. Directly attached to the vmi via qemu. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
	PersistentVolumeClaim *KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource `json:"persistentVolumeClaim,omitempty"`
	// SecretVolumeSource represents a reference to a secret data in the same namespace. More info: https://kubernetes.io/docs/concepts/configuration/secret/
	Secret *KubevirtIoApiCoreV1SecretVolumeSource `json:"secret,omitempty"`
	// ServiceAccountVolumeSource represents a reference to a service account. There can only be one volume of this type! More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/
	ServiceAccount *KubevirtIoApiCoreV1ServiceAccountVolumeSource `json:"serviceAccount,omitempty"`
	// Represents a Sysprep volume source.
	Sysprep *KubevirtIoApiCoreV1SysprepSource `json:"sysprep,omitempty"`
}

type _KubevirtIoApiCoreV1Volume KubevirtIoApiCoreV1Volume

// NewKubevirtIoApiCoreV1Volume instantiates a new KubevirtIoApiCoreV1Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1Volume(name string) *KubevirtIoApiCoreV1Volume {
	this := KubevirtIoApiCoreV1Volume{}
	this.Name = name
	return &this
}

// NewKubevirtIoApiCoreV1VolumeWithDefaults instantiates a new KubevirtIoApiCoreV1Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1VolumeWithDefaults() *KubevirtIoApiCoreV1Volume {
	this := KubevirtIoApiCoreV1Volume{}
	var name string = ""
	this.Name = name
	return &this
}

// GetCloudInitConfigDrive returns the CloudInitConfigDrive field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetCloudInitConfigDrive() KubevirtIoApiCoreV1CloudInitConfigDriveSource {
	if o == nil || IsNil(o.CloudInitConfigDrive) {
		var ret KubevirtIoApiCoreV1CloudInitConfigDriveSource
		return ret
	}
	return *o.CloudInitConfigDrive
}

// GetCloudInitConfigDriveOk returns a tuple with the CloudInitConfigDrive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetCloudInitConfigDriveOk() (*KubevirtIoApiCoreV1CloudInitConfigDriveSource, bool) {
	if o == nil || IsNil(o.CloudInitConfigDrive) {
		return nil, false
	}
	return o.CloudInitConfigDrive, true
}

// HasCloudInitConfigDrive returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasCloudInitConfigDrive() bool {
	if o != nil && !IsNil(o.CloudInitConfigDrive) {
		return true
	}

	return false
}

// SetCloudInitConfigDrive gets a reference to the given KubevirtIoApiCoreV1CloudInitConfigDriveSource and assigns it to the CloudInitConfigDrive field.
func (o *KubevirtIoApiCoreV1Volume) SetCloudInitConfigDrive(v KubevirtIoApiCoreV1CloudInitConfigDriveSource) {
	o.CloudInitConfigDrive = &v
}

// GetCloudInitNoCloud returns the CloudInitNoCloud field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetCloudInitNoCloud() KubevirtIoApiCoreV1CloudInitNoCloudSource {
	if o == nil || IsNil(o.CloudInitNoCloud) {
		var ret KubevirtIoApiCoreV1CloudInitNoCloudSource
		return ret
	}
	return *o.CloudInitNoCloud
}

// GetCloudInitNoCloudOk returns a tuple with the CloudInitNoCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetCloudInitNoCloudOk() (*KubevirtIoApiCoreV1CloudInitNoCloudSource, bool) {
	if o == nil || IsNil(o.CloudInitNoCloud) {
		return nil, false
	}
	return o.CloudInitNoCloud, true
}

// HasCloudInitNoCloud returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasCloudInitNoCloud() bool {
	if o != nil && !IsNil(o.CloudInitNoCloud) {
		return true
	}

	return false
}

// SetCloudInitNoCloud gets a reference to the given KubevirtIoApiCoreV1CloudInitNoCloudSource and assigns it to the CloudInitNoCloud field.
func (o *KubevirtIoApiCoreV1Volume) SetCloudInitNoCloud(v KubevirtIoApiCoreV1CloudInitNoCloudSource) {
	o.CloudInitNoCloud = &v
}

// GetConfigMap returns the ConfigMap field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetConfigMap() KubevirtIoApiCoreV1ConfigMapVolumeSource {
	if o == nil || IsNil(o.ConfigMap) {
		var ret KubevirtIoApiCoreV1ConfigMapVolumeSource
		return ret
	}
	return *o.ConfigMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetConfigMapOk() (*KubevirtIoApiCoreV1ConfigMapVolumeSource, bool) {
	if o == nil || IsNil(o.ConfigMap) {
		return nil, false
	}
	return o.ConfigMap, true
}

// HasConfigMap returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasConfigMap() bool {
	if o != nil && !IsNil(o.ConfigMap) {
		return true
	}

	return false
}

// SetConfigMap gets a reference to the given KubevirtIoApiCoreV1ConfigMapVolumeSource and assigns it to the ConfigMap field.
func (o *KubevirtIoApiCoreV1Volume) SetConfigMap(v KubevirtIoApiCoreV1ConfigMapVolumeSource) {
	o.ConfigMap = &v
}

// GetContainerDisk returns the ContainerDisk field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetContainerDisk() KubevirtIoApiCoreV1ContainerDiskSource {
	if o == nil || IsNil(o.ContainerDisk) {
		var ret KubevirtIoApiCoreV1ContainerDiskSource
		return ret
	}
	return *o.ContainerDisk
}

// GetContainerDiskOk returns a tuple with the ContainerDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetContainerDiskOk() (*KubevirtIoApiCoreV1ContainerDiskSource, bool) {
	if o == nil || IsNil(o.ContainerDisk) {
		return nil, false
	}
	return o.ContainerDisk, true
}

// HasContainerDisk returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasContainerDisk() bool {
	if o != nil && !IsNil(o.ContainerDisk) {
		return true
	}

	return false
}

// SetContainerDisk gets a reference to the given KubevirtIoApiCoreV1ContainerDiskSource and assigns it to the ContainerDisk field.
func (o *KubevirtIoApiCoreV1Volume) SetContainerDisk(v KubevirtIoApiCoreV1ContainerDiskSource) {
	o.ContainerDisk = &v
}

// GetDataVolume returns the DataVolume field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetDataVolume() KubevirtIoApiCoreV1DataVolumeSource {
	if o == nil || IsNil(o.DataVolume) {
		var ret KubevirtIoApiCoreV1DataVolumeSource
		return ret
	}
	return *o.DataVolume
}

// GetDataVolumeOk returns a tuple with the DataVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetDataVolumeOk() (*KubevirtIoApiCoreV1DataVolumeSource, bool) {
	if o == nil || IsNil(o.DataVolume) {
		return nil, false
	}
	return o.DataVolume, true
}

// HasDataVolume returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasDataVolume() bool {
	if o != nil && !IsNil(o.DataVolume) {
		return true
	}

	return false
}

// SetDataVolume gets a reference to the given KubevirtIoApiCoreV1DataVolumeSource and assigns it to the DataVolume field.
func (o *KubevirtIoApiCoreV1Volume) SetDataVolume(v KubevirtIoApiCoreV1DataVolumeSource) {
	o.DataVolume = &v
}

// GetDownwardAPI returns the DownwardAPI field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetDownwardAPI() KubevirtIoApiCoreV1DownwardAPIVolumeSource {
	if o == nil || IsNil(o.DownwardAPI) {
		var ret KubevirtIoApiCoreV1DownwardAPIVolumeSource
		return ret
	}
	return *o.DownwardAPI
}

// GetDownwardAPIOk returns a tuple with the DownwardAPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetDownwardAPIOk() (*KubevirtIoApiCoreV1DownwardAPIVolumeSource, bool) {
	if o == nil || IsNil(o.DownwardAPI) {
		return nil, false
	}
	return o.DownwardAPI, true
}

// HasDownwardAPI returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasDownwardAPI() bool {
	if o != nil && !IsNil(o.DownwardAPI) {
		return true
	}

	return false
}

// SetDownwardAPI gets a reference to the given KubevirtIoApiCoreV1DownwardAPIVolumeSource and assigns it to the DownwardAPI field.
func (o *KubevirtIoApiCoreV1Volume) SetDownwardAPI(v KubevirtIoApiCoreV1DownwardAPIVolumeSource) {
	o.DownwardAPI = &v
}

// GetDownwardMetrics returns the DownwardMetrics field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetDownwardMetrics() map[string]interface{} {
	if o == nil || IsNil(o.DownwardMetrics) {
		var ret map[string]interface{}
		return ret
	}
	return o.DownwardMetrics
}

// GetDownwardMetricsOk returns a tuple with the DownwardMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetDownwardMetricsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DownwardMetrics) {
		return map[string]interface{}{}, false
	}
	return o.DownwardMetrics, true
}

// HasDownwardMetrics returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasDownwardMetrics() bool {
	if o != nil && !IsNil(o.DownwardMetrics) {
		return true
	}

	return false
}

// SetDownwardMetrics gets a reference to the given map[string]interface{} and assigns it to the DownwardMetrics field.
func (o *KubevirtIoApiCoreV1Volume) SetDownwardMetrics(v map[string]interface{}) {
	o.DownwardMetrics = v
}

// GetEmptyDisk returns the EmptyDisk field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetEmptyDisk() KubevirtIoApiCoreV1EmptyDiskSource {
	if o == nil || IsNil(o.EmptyDisk) {
		var ret KubevirtIoApiCoreV1EmptyDiskSource
		return ret
	}
	return *o.EmptyDisk
}

// GetEmptyDiskOk returns a tuple with the EmptyDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetEmptyDiskOk() (*KubevirtIoApiCoreV1EmptyDiskSource, bool) {
	if o == nil || IsNil(o.EmptyDisk) {
		return nil, false
	}
	return o.EmptyDisk, true
}

// HasEmptyDisk returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasEmptyDisk() bool {
	if o != nil && !IsNil(o.EmptyDisk) {
		return true
	}

	return false
}

// SetEmptyDisk gets a reference to the given KubevirtIoApiCoreV1EmptyDiskSource and assigns it to the EmptyDisk field.
func (o *KubevirtIoApiCoreV1Volume) SetEmptyDisk(v KubevirtIoApiCoreV1EmptyDiskSource) {
	o.EmptyDisk = &v
}

// GetEphemeral returns the Ephemeral field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetEphemeral() KubevirtIoApiCoreV1EphemeralVolumeSource {
	if o == nil || IsNil(o.Ephemeral) {
		var ret KubevirtIoApiCoreV1EphemeralVolumeSource
		return ret
	}
	return *o.Ephemeral
}

// GetEphemeralOk returns a tuple with the Ephemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetEphemeralOk() (*KubevirtIoApiCoreV1EphemeralVolumeSource, bool) {
	if o == nil || IsNil(o.Ephemeral) {
		return nil, false
	}
	return o.Ephemeral, true
}

// HasEphemeral returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasEphemeral() bool {
	if o != nil && !IsNil(o.Ephemeral) {
		return true
	}

	return false
}

// SetEphemeral gets a reference to the given KubevirtIoApiCoreV1EphemeralVolumeSource and assigns it to the Ephemeral field.
func (o *KubevirtIoApiCoreV1Volume) SetEphemeral(v KubevirtIoApiCoreV1EphemeralVolumeSource) {
	o.Ephemeral = &v
}

// GetHostDisk returns the HostDisk field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetHostDisk() KubevirtIoApiCoreV1HostDisk {
	if o == nil || IsNil(o.HostDisk) {
		var ret KubevirtIoApiCoreV1HostDisk
		return ret
	}
	return *o.HostDisk
}

// GetHostDiskOk returns a tuple with the HostDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetHostDiskOk() (*KubevirtIoApiCoreV1HostDisk, bool) {
	if o == nil || IsNil(o.HostDisk) {
		return nil, false
	}
	return o.HostDisk, true
}

// HasHostDisk returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasHostDisk() bool {
	if o != nil && !IsNil(o.HostDisk) {
		return true
	}

	return false
}

// SetHostDisk gets a reference to the given KubevirtIoApiCoreV1HostDisk and assigns it to the HostDisk field.
func (o *KubevirtIoApiCoreV1Volume) SetHostDisk(v KubevirtIoApiCoreV1HostDisk) {
	o.HostDisk = &v
}

// GetMemoryDump returns the MemoryDump field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetMemoryDump() KubevirtIoApiCoreV1MemoryDumpVolumeSource {
	if o == nil || IsNil(o.MemoryDump) {
		var ret KubevirtIoApiCoreV1MemoryDumpVolumeSource
		return ret
	}
	return *o.MemoryDump
}

// GetMemoryDumpOk returns a tuple with the MemoryDump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetMemoryDumpOk() (*KubevirtIoApiCoreV1MemoryDumpVolumeSource, bool) {
	if o == nil || IsNil(o.MemoryDump) {
		return nil, false
	}
	return o.MemoryDump, true
}

// HasMemoryDump returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasMemoryDump() bool {
	if o != nil && !IsNil(o.MemoryDump) {
		return true
	}

	return false
}

// SetMemoryDump gets a reference to the given KubevirtIoApiCoreV1MemoryDumpVolumeSource and assigns it to the MemoryDump field.
func (o *KubevirtIoApiCoreV1Volume) SetMemoryDump(v KubevirtIoApiCoreV1MemoryDumpVolumeSource) {
	o.MemoryDump = &v
}

// GetName returns the Name field value
func (o *KubevirtIoApiCoreV1Volume) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KubevirtIoApiCoreV1Volume) SetName(v string) {
	o.Name = v
}

// GetPersistentVolumeClaim returns the PersistentVolumeClaim field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetPersistentVolumeClaim() KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource {
	if o == nil || IsNil(o.PersistentVolumeClaim) {
		var ret KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource
		return ret
	}
	return *o.PersistentVolumeClaim
}

// GetPersistentVolumeClaimOk returns a tuple with the PersistentVolumeClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetPersistentVolumeClaimOk() (*KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource, bool) {
	if o == nil || IsNil(o.PersistentVolumeClaim) {
		return nil, false
	}
	return o.PersistentVolumeClaim, true
}

// HasPersistentVolumeClaim returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasPersistentVolumeClaim() bool {
	if o != nil && !IsNil(o.PersistentVolumeClaim) {
		return true
	}

	return false
}

// SetPersistentVolumeClaim gets a reference to the given KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource and assigns it to the PersistentVolumeClaim field.
func (o *KubevirtIoApiCoreV1Volume) SetPersistentVolumeClaim(v KubevirtIoApiCoreV1PersistentVolumeClaimVolumeSource) {
	o.PersistentVolumeClaim = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetSecret() KubevirtIoApiCoreV1SecretVolumeSource {
	if o == nil || IsNil(o.Secret) {
		var ret KubevirtIoApiCoreV1SecretVolumeSource
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetSecretOk() (*KubevirtIoApiCoreV1SecretVolumeSource, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given KubevirtIoApiCoreV1SecretVolumeSource and assigns it to the Secret field.
func (o *KubevirtIoApiCoreV1Volume) SetSecret(v KubevirtIoApiCoreV1SecretVolumeSource) {
	o.Secret = &v
}

// GetServiceAccount returns the ServiceAccount field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetServiceAccount() KubevirtIoApiCoreV1ServiceAccountVolumeSource {
	if o == nil || IsNil(o.ServiceAccount) {
		var ret KubevirtIoApiCoreV1ServiceAccountVolumeSource
		return ret
	}
	return *o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetServiceAccountOk() (*KubevirtIoApiCoreV1ServiceAccountVolumeSource, bool) {
	if o == nil || IsNil(o.ServiceAccount) {
		return nil, false
	}
	return o.ServiceAccount, true
}

// HasServiceAccount returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasServiceAccount() bool {
	if o != nil && !IsNil(o.ServiceAccount) {
		return true
	}

	return false
}

// SetServiceAccount gets a reference to the given KubevirtIoApiCoreV1ServiceAccountVolumeSource and assigns it to the ServiceAccount field.
func (o *KubevirtIoApiCoreV1Volume) SetServiceAccount(v KubevirtIoApiCoreV1ServiceAccountVolumeSource) {
	o.ServiceAccount = &v
}

// GetSysprep returns the Sysprep field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1Volume) GetSysprep() KubevirtIoApiCoreV1SysprepSource {
	if o == nil || IsNil(o.Sysprep) {
		var ret KubevirtIoApiCoreV1SysprepSource
		return ret
	}
	return *o.Sysprep
}

// GetSysprepOk returns a tuple with the Sysprep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1Volume) GetSysprepOk() (*KubevirtIoApiCoreV1SysprepSource, bool) {
	if o == nil || IsNil(o.Sysprep) {
		return nil, false
	}
	return o.Sysprep, true
}

// HasSysprep returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1Volume) HasSysprep() bool {
	if o != nil && !IsNil(o.Sysprep) {
		return true
	}

	return false
}

// SetSysprep gets a reference to the given KubevirtIoApiCoreV1SysprepSource and assigns it to the Sysprep field.
func (o *KubevirtIoApiCoreV1Volume) SetSysprep(v KubevirtIoApiCoreV1SysprepSource) {
	o.Sysprep = &v
}

func (o KubevirtIoApiCoreV1Volume) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1Volume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudInitConfigDrive) {
		toSerialize["cloudInitConfigDrive"] = o.CloudInitConfigDrive
	}
	if !IsNil(o.CloudInitNoCloud) {
		toSerialize["cloudInitNoCloud"] = o.CloudInitNoCloud
	}
	if !IsNil(o.ConfigMap) {
		toSerialize["configMap"] = o.ConfigMap
	}
	if !IsNil(o.ContainerDisk) {
		toSerialize["containerDisk"] = o.ContainerDisk
	}
	if !IsNil(o.DataVolume) {
		toSerialize["dataVolume"] = o.DataVolume
	}
	if !IsNil(o.DownwardAPI) {
		toSerialize["downwardAPI"] = o.DownwardAPI
	}
	if !IsNil(o.DownwardMetrics) {
		toSerialize["downwardMetrics"] = o.DownwardMetrics
	}
	if !IsNil(o.EmptyDisk) {
		toSerialize["emptyDisk"] = o.EmptyDisk
	}
	if !IsNil(o.Ephemeral) {
		toSerialize["ephemeral"] = o.Ephemeral
	}
	if !IsNil(o.HostDisk) {
		toSerialize["hostDisk"] = o.HostDisk
	}
	if !IsNil(o.MemoryDump) {
		toSerialize["memoryDump"] = o.MemoryDump
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PersistentVolumeClaim) {
		toSerialize["persistentVolumeClaim"] = o.PersistentVolumeClaim
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.ServiceAccount) {
		toSerialize["serviceAccount"] = o.ServiceAccount
	}
	if !IsNil(o.Sysprep) {
		toSerialize["sysprep"] = o.Sysprep
	}
	return toSerialize, nil
}

func (o *KubevirtIoApiCoreV1Volume) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKubevirtIoApiCoreV1Volume := _KubevirtIoApiCoreV1Volume{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoApiCoreV1Volume)

	if err != nil {
		return err
	}

	*o = KubevirtIoApiCoreV1Volume(varKubevirtIoApiCoreV1Volume)

	return err
}

type NullableKubevirtIoApiCoreV1Volume struct {
	value *KubevirtIoApiCoreV1Volume
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1Volume) Get() *KubevirtIoApiCoreV1Volume {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1Volume) Set(val *KubevirtIoApiCoreV1Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1Volume) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1Volume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1Volume(val *KubevirtIoApiCoreV1Volume) *NullableKubevirtIoApiCoreV1Volume {
	return &NullableKubevirtIoApiCoreV1Volume{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1Volume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1Volume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


