# KubevirtIoApiCoreV1InstancetypeMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InferFromVolume** | Pointer to **string** | InferFromVolume lists the name of a volume that should be used to infer or discover the instancetype to be used through known annotations on the underlying resource. Once applied to the InstancetypeMatcher this field is removed. | [optional] 
**InferFromVolumeFailurePolicy** | Pointer to **string** | InferFromVolumeFailurePolicy controls what should happen on failure when inferring the instancetype. Allowed values are: \&quot;RejectInferFromVolumeFailure\&quot; and \&quot;IgnoreInferFromVolumeFailure\&quot;. If not specified, \&quot;RejectInferFromVolumeFailure\&quot; is used by default. | [optional] 
**Kind** | Pointer to **string** | Kind specifies which instancetype resource is referenced. Allowed values are: \&quot;VirtualMachineInstancetype\&quot; and \&quot;VirtualMachineClusterInstancetype\&quot;. If not specified, \&quot;VirtualMachineClusterInstancetype\&quot; is used by default. | [optional] 
**Name** | Pointer to **string** | Name is the name of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype | [optional] 
**RevisionName** | Pointer to **string** | RevisionName specifies a ControllerRevision containing a specific copy of the VirtualMachineInstancetype or VirtualMachineClusterInstancetype to be used. This is initially captured the first time the instancetype is applied to the VirtualMachineInstance. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1InstancetypeMatcher

`func NewKubevirtIoApiCoreV1InstancetypeMatcher() *KubevirtIoApiCoreV1InstancetypeMatcher`

NewKubevirtIoApiCoreV1InstancetypeMatcher instantiates a new KubevirtIoApiCoreV1InstancetypeMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1InstancetypeMatcherWithDefaults

`func NewKubevirtIoApiCoreV1InstancetypeMatcherWithDefaults() *KubevirtIoApiCoreV1InstancetypeMatcher`

NewKubevirtIoApiCoreV1InstancetypeMatcherWithDefaults instantiates a new KubevirtIoApiCoreV1InstancetypeMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInferFromVolume

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetInferFromVolume() string`

GetInferFromVolume returns the InferFromVolume field if non-nil, zero value otherwise.

### GetInferFromVolumeOk

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetInferFromVolumeOk() (*string, bool)`

GetInferFromVolumeOk returns a tuple with the InferFromVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferFromVolume

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) SetInferFromVolume(v string)`

SetInferFromVolume sets InferFromVolume field to given value.

### HasInferFromVolume

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) HasInferFromVolume() bool`

HasInferFromVolume returns a boolean if a field has been set.

### GetInferFromVolumeFailurePolicy

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetInferFromVolumeFailurePolicy() string`

GetInferFromVolumeFailurePolicy returns the InferFromVolumeFailurePolicy field if non-nil, zero value otherwise.

### GetInferFromVolumeFailurePolicyOk

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetInferFromVolumeFailurePolicyOk() (*string, bool)`

GetInferFromVolumeFailurePolicyOk returns a tuple with the InferFromVolumeFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferFromVolumeFailurePolicy

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) SetInferFromVolumeFailurePolicy(v string)`

SetInferFromVolumeFailurePolicy sets InferFromVolumeFailurePolicy field to given value.

### HasInferFromVolumeFailurePolicy

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) HasInferFromVolumeFailurePolicy() bool`

HasInferFromVolumeFailurePolicy returns a boolean if a field has been set.

### GetKind

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRevisionName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetRevisionName() string`

GetRevisionName returns the RevisionName field if non-nil, zero value otherwise.

### GetRevisionNameOk

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) GetRevisionNameOk() (*string, bool)`

GetRevisionNameOk returns a tuple with the RevisionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) SetRevisionName(v string)`

SetRevisionName sets RevisionName field to given value.

### HasRevisionName

`func (o *KubevirtIoApiCoreV1InstancetypeMatcher) HasRevisionName() bool`

HasRevisionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


