# KubevirtIoApiCoreV1Features

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acpi** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) | ACPI enables/disables ACPI inside the guest. Defaults to enabled. | [optional] [default to {}]
**Apic** | Pointer to [**KubevirtIoApiCoreV1FeatureAPIC**](KubevirtIoApiCoreV1FeatureAPIC.md) | Defaults to the machine type setting. | [optional] 
**Hyperv** | Pointer to [**KubevirtIoApiCoreV1FeatureHyperv**](KubevirtIoApiCoreV1FeatureHyperv.md) | Defaults to the machine type setting. | [optional] 
**Kvm** | Pointer to [**KubevirtIoApiCoreV1FeatureKVM**](KubevirtIoApiCoreV1FeatureKVM.md) | Configure how KVM presence is exposed to the guest. | [optional] 
**Pvspinlock** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) | Notify the guest that the host supports paravirtual spinlocks. For older kernels this feature should be explicitly disabled. | [optional] 
**Smm** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) | SMM enables/disables System Management Mode. TSEG not yet implemented. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Features

`func NewKubevirtIoApiCoreV1Features() *KubevirtIoApiCoreV1Features`

NewKubevirtIoApiCoreV1Features instantiates a new KubevirtIoApiCoreV1Features object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1FeaturesWithDefaults

`func NewKubevirtIoApiCoreV1FeaturesWithDefaults() *KubevirtIoApiCoreV1Features`

NewKubevirtIoApiCoreV1FeaturesWithDefaults instantiates a new KubevirtIoApiCoreV1Features object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *KubevirtIoApiCoreV1Features) GetAcpi() KubevirtIoApiCoreV1FeatureState`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *KubevirtIoApiCoreV1Features) GetAcpiOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *KubevirtIoApiCoreV1Features) SetAcpi(v KubevirtIoApiCoreV1FeatureState)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *KubevirtIoApiCoreV1Features) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetApic

`func (o *KubevirtIoApiCoreV1Features) GetApic() KubevirtIoApiCoreV1FeatureAPIC`

GetApic returns the Apic field if non-nil, zero value otherwise.

### GetApicOk

`func (o *KubevirtIoApiCoreV1Features) GetApicOk() (*KubevirtIoApiCoreV1FeatureAPIC, bool)`

GetApicOk returns a tuple with the Apic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApic

`func (o *KubevirtIoApiCoreV1Features) SetApic(v KubevirtIoApiCoreV1FeatureAPIC)`

SetApic sets Apic field to given value.

### HasApic

`func (o *KubevirtIoApiCoreV1Features) HasApic() bool`

HasApic returns a boolean if a field has been set.

### GetHyperv

`func (o *KubevirtIoApiCoreV1Features) GetHyperv() KubevirtIoApiCoreV1FeatureHyperv`

GetHyperv returns the Hyperv field if non-nil, zero value otherwise.

### GetHypervOk

`func (o *KubevirtIoApiCoreV1Features) GetHypervOk() (*KubevirtIoApiCoreV1FeatureHyperv, bool)`

GetHypervOk returns a tuple with the Hyperv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperv

`func (o *KubevirtIoApiCoreV1Features) SetHyperv(v KubevirtIoApiCoreV1FeatureHyperv)`

SetHyperv sets Hyperv field to given value.

### HasHyperv

`func (o *KubevirtIoApiCoreV1Features) HasHyperv() bool`

HasHyperv returns a boolean if a field has been set.

### GetKvm

`func (o *KubevirtIoApiCoreV1Features) GetKvm() KubevirtIoApiCoreV1FeatureKVM`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *KubevirtIoApiCoreV1Features) GetKvmOk() (*KubevirtIoApiCoreV1FeatureKVM, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *KubevirtIoApiCoreV1Features) SetKvm(v KubevirtIoApiCoreV1FeatureKVM)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *KubevirtIoApiCoreV1Features) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetPvspinlock

`func (o *KubevirtIoApiCoreV1Features) GetPvspinlock() KubevirtIoApiCoreV1FeatureState`

GetPvspinlock returns the Pvspinlock field if non-nil, zero value otherwise.

### GetPvspinlockOk

`func (o *KubevirtIoApiCoreV1Features) GetPvspinlockOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetPvspinlockOk returns a tuple with the Pvspinlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvspinlock

`func (o *KubevirtIoApiCoreV1Features) SetPvspinlock(v KubevirtIoApiCoreV1FeatureState)`

SetPvspinlock sets Pvspinlock field to given value.

### HasPvspinlock

`func (o *KubevirtIoApiCoreV1Features) HasPvspinlock() bool`

HasPvspinlock returns a boolean if a field has been set.

### GetSmm

`func (o *KubevirtIoApiCoreV1Features) GetSmm() KubevirtIoApiCoreV1FeatureState`

GetSmm returns the Smm field if non-nil, zero value otherwise.

### GetSmmOk

`func (o *KubevirtIoApiCoreV1Features) GetSmmOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetSmmOk returns a tuple with the Smm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmm

`func (o *KubevirtIoApiCoreV1Features) SetSmm(v KubevirtIoApiCoreV1FeatureState)`

SetSmm sets Smm field to given value.

### HasSmm

`func (o *KubevirtIoApiCoreV1Features) HasSmm() bool`

HasSmm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


