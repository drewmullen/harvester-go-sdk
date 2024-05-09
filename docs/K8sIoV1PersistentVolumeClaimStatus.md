# K8sIoV1PersistentVolumeClaimStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | Pointer to **[]string** | accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 | [optional] 
**AllocatedResources** | Pointer to **map[string]string** | allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature. | [optional] 
**Capacity** | Pointer to **map[string]string** | capacity represents the actual resources of the underlying volume. | [optional] 
**Conditions** | Pointer to [**[]K8sIoV1PersistentVolumeClaimCondition**](K8sIoV1PersistentVolumeClaimCondition.md) | conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to &#39;ResizeStarted&#39;. | [optional] 
**Phase** | Pointer to **string** | phase represents the current phase of PersistentVolumeClaim.  Possible enum values:  - &#x60;\&quot;Bound\&quot;&#x60; used for PersistentVolumeClaims that are bound  - &#x60;\&quot;Lost\&quot;&#x60; used for PersistentVolumeClaims that lost their underlying PersistentVolume. The claim was bound to a PersistentVolume and this volume does not exist any longer and all data on it was lost.  - &#x60;\&quot;Pending\&quot;&#x60; used for PersistentVolumeClaims that are not yet bound | [optional] 
**ResizeStatus** | Pointer to **string** | resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature. | [optional] 

## Methods

### NewK8sIoV1PersistentVolumeClaimStatus

`func NewK8sIoV1PersistentVolumeClaimStatus() *K8sIoV1PersistentVolumeClaimStatus`

NewK8sIoV1PersistentVolumeClaimStatus instantiates a new K8sIoV1PersistentVolumeClaimStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PersistentVolumeClaimStatusWithDefaults

`func NewK8sIoV1PersistentVolumeClaimStatusWithDefaults() *K8sIoV1PersistentVolumeClaimStatus`

NewK8sIoV1PersistentVolumeClaimStatusWithDefaults instantiates a new K8sIoV1PersistentVolumeClaimStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.

### HasAccessModes

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasAccessModes() bool`

HasAccessModes returns a boolean if a field has been set.

### GetAllocatedResources

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetAllocatedResources() map[string]string`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetAllocatedResourcesOk() (*map[string]string, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetAllocatedResources(v map[string]string)`

SetAllocatedResources sets AllocatedResources field to given value.

### HasAllocatedResources

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasAllocatedResources() bool`

HasAllocatedResources returns a boolean if a field has been set.

### GetCapacity

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetCapacity() map[string]string`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetCapacityOk() (*map[string]string, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetCapacity(v map[string]string)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetConditions

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetConditions() []K8sIoV1PersistentVolumeClaimCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetConditionsOk() (*[]K8sIoV1PersistentVolumeClaimCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetConditions(v []K8sIoV1PersistentVolumeClaimCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetResizeStatus

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetResizeStatus() string`

GetResizeStatus returns the ResizeStatus field if non-nil, zero value otherwise.

### GetResizeStatusOk

`func (o *K8sIoV1PersistentVolumeClaimStatus) GetResizeStatusOk() (*string, bool)`

GetResizeStatusOk returns a tuple with the ResizeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizeStatus

`func (o *K8sIoV1PersistentVolumeClaimStatus) SetResizeStatus(v string)`

SetResizeStatus sets ResizeStatus field to given value.

### HasResizeStatus

`func (o *K8sIoV1PersistentVolumeClaimStatus) HasResizeStatus() bool`

HasResizeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


