# KubevirtIoApiCoreV1VirtualMachineSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataVolumeTemplates** | Pointer to [**[]KubevirtIoApiCoreV1DataVolumeTemplateSpec**](KubevirtIoApiCoreV1DataVolumeTemplateSpec.md) | dataVolumeTemplates is a list of dataVolumes that the VirtualMachineInstance template can reference. DataVolumes in this list are dynamically created for the VirtualMachine and are tied to the VirtualMachine&#39;s life-cycle. | [optional] 
**Instancetype** | Pointer to [**KubevirtIoApiCoreV1InstancetypeMatcher**](KubevirtIoApiCoreV1InstancetypeMatcher.md) | InstancetypeMatcher references a instancetype that is used to fill fields in Template | [optional] 
**LiveUpdateFeatures** | Pointer to [**KubevirtIoApiCoreV1LiveUpdateFeatures**](KubevirtIoApiCoreV1LiveUpdateFeatures.md) | LiveUpdateFeatures references a configuration of hotpluggable resources | [optional] 
**Preference** | Pointer to [**KubevirtIoApiCoreV1PreferenceMatcher**](KubevirtIoApiCoreV1PreferenceMatcher.md) | PreferenceMatcher references a set of preference that is used to fill fields in Template | [optional] 
**RunStrategy** | Pointer to **string** | Running state indicates the requested running state of the VirtualMachineInstance mutually exclusive with Running | [optional] 
**Running** | Pointer to **bool** | Running controls whether the associatied VirtualMachineInstance is created or not Mutually exclusive with RunStrategy | [optional] 
**Template** | [**KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec**](KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec.md) | Template is the direct specification of VirtualMachineInstance | 

## Methods

### NewKubevirtIoApiCoreV1VirtualMachineSpec

`func NewKubevirtIoApiCoreV1VirtualMachineSpec(template KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec, ) *KubevirtIoApiCoreV1VirtualMachineSpec`

NewKubevirtIoApiCoreV1VirtualMachineSpec instantiates a new KubevirtIoApiCoreV1VirtualMachineSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1VirtualMachineSpecWithDefaults

`func NewKubevirtIoApiCoreV1VirtualMachineSpecWithDefaults() *KubevirtIoApiCoreV1VirtualMachineSpec`

NewKubevirtIoApiCoreV1VirtualMachineSpecWithDefaults instantiates a new KubevirtIoApiCoreV1VirtualMachineSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataVolumeTemplates

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetDataVolumeTemplates() []KubevirtIoApiCoreV1DataVolumeTemplateSpec`

GetDataVolumeTemplates returns the DataVolumeTemplates field if non-nil, zero value otherwise.

### GetDataVolumeTemplatesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetDataVolumeTemplatesOk() (*[]KubevirtIoApiCoreV1DataVolumeTemplateSpec, bool)`

GetDataVolumeTemplatesOk returns a tuple with the DataVolumeTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataVolumeTemplates

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetDataVolumeTemplates(v []KubevirtIoApiCoreV1DataVolumeTemplateSpec)`

SetDataVolumeTemplates sets DataVolumeTemplates field to given value.

### HasDataVolumeTemplates

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasDataVolumeTemplates() bool`

HasDataVolumeTemplates returns a boolean if a field has been set.

### GetInstancetype

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetInstancetype() KubevirtIoApiCoreV1InstancetypeMatcher`

GetInstancetype returns the Instancetype field if non-nil, zero value otherwise.

### GetInstancetypeOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetInstancetypeOk() (*KubevirtIoApiCoreV1InstancetypeMatcher, bool)`

GetInstancetypeOk returns a tuple with the Instancetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancetype

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetInstancetype(v KubevirtIoApiCoreV1InstancetypeMatcher)`

SetInstancetype sets Instancetype field to given value.

### HasInstancetype

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasInstancetype() bool`

HasInstancetype returns a boolean if a field has been set.

### GetLiveUpdateFeatures

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetLiveUpdateFeatures() KubevirtIoApiCoreV1LiveUpdateFeatures`

GetLiveUpdateFeatures returns the LiveUpdateFeatures field if non-nil, zero value otherwise.

### GetLiveUpdateFeaturesOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetLiveUpdateFeaturesOk() (*KubevirtIoApiCoreV1LiveUpdateFeatures, bool)`

GetLiveUpdateFeaturesOk returns a tuple with the LiveUpdateFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveUpdateFeatures

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetLiveUpdateFeatures(v KubevirtIoApiCoreV1LiveUpdateFeatures)`

SetLiveUpdateFeatures sets LiveUpdateFeatures field to given value.

### HasLiveUpdateFeatures

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasLiveUpdateFeatures() bool`

HasLiveUpdateFeatures returns a boolean if a field has been set.

### GetPreference

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetPreference() KubevirtIoApiCoreV1PreferenceMatcher`

GetPreference returns the Preference field if non-nil, zero value otherwise.

### GetPreferenceOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetPreferenceOk() (*KubevirtIoApiCoreV1PreferenceMatcher, bool)`

GetPreferenceOk returns a tuple with the Preference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreference

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetPreference(v KubevirtIoApiCoreV1PreferenceMatcher)`

SetPreference sets Preference field to given value.

### HasPreference

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasPreference() bool`

HasPreference returns a boolean if a field has been set.

### GetRunStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetRunStrategy() string`

GetRunStrategy returns the RunStrategy field if non-nil, zero value otherwise.

### GetRunStrategyOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetRunStrategyOk() (*string, bool)`

GetRunStrategyOk returns a tuple with the RunStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetRunStrategy(v string)`

SetRunStrategy sets RunStrategy field to given value.

### HasRunStrategy

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasRunStrategy() bool`

HasRunStrategy returns a boolean if a field has been set.

### GetRunning

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTemplate

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetTemplate() KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) GetTemplateOk() (*KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *KubevirtIoApiCoreV1VirtualMachineSpec) SetTemplate(v KubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


