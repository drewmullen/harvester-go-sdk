# KubevirtIoApiCoreV1FeatureHyperv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evmcs** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Frequencies** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Ipi** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Reenlightenment** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Relaxed** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Reset** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Runtime** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Spinlocks** | Pointer to [**KubevirtIoApiCoreV1FeatureSpinlocks**](KubevirtIoApiCoreV1FeatureSpinlocks.md) |  | [optional] 
**Synic** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Synictimer** | Pointer to [**KubevirtIoApiCoreV1SyNICTimer**](KubevirtIoApiCoreV1SyNICTimer.md) |  | [optional] 
**Tlbflush** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Vapic** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 
**Vendorid** | Pointer to [**KubevirtIoApiCoreV1FeatureVendorID**](KubevirtIoApiCoreV1FeatureVendorID.md) |  | [optional] 
**Vpindex** | Pointer to [**KubevirtIoApiCoreV1FeatureState**](KubevirtIoApiCoreV1FeatureState.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1FeatureHyperv

`func NewKubevirtIoApiCoreV1FeatureHyperv() *KubevirtIoApiCoreV1FeatureHyperv`

NewKubevirtIoApiCoreV1FeatureHyperv instantiates a new KubevirtIoApiCoreV1FeatureHyperv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1FeatureHypervWithDefaults

`func NewKubevirtIoApiCoreV1FeatureHypervWithDefaults() *KubevirtIoApiCoreV1FeatureHyperv`

NewKubevirtIoApiCoreV1FeatureHypervWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureHyperv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvmcs

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetEvmcs() KubevirtIoApiCoreV1FeatureState`

GetEvmcs returns the Evmcs field if non-nil, zero value otherwise.

### GetEvmcsOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetEvmcsOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetEvmcsOk returns a tuple with the Evmcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmcs

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetEvmcs(v KubevirtIoApiCoreV1FeatureState)`

SetEvmcs sets Evmcs field to given value.

### HasEvmcs

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasEvmcs() bool`

HasEvmcs returns a boolean if a field has been set.

### GetFrequencies

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetFrequencies() KubevirtIoApiCoreV1FeatureState`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetFrequenciesOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetFrequencies(v KubevirtIoApiCoreV1FeatureState)`

SetFrequencies sets Frequencies field to given value.

### HasFrequencies

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasFrequencies() bool`

HasFrequencies returns a boolean if a field has been set.

### GetIpi

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetIpi() KubevirtIoApiCoreV1FeatureState`

GetIpi returns the Ipi field if non-nil, zero value otherwise.

### GetIpiOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetIpiOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetIpiOk returns a tuple with the Ipi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpi

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetIpi(v KubevirtIoApiCoreV1FeatureState)`

SetIpi sets Ipi field to given value.

### HasIpi

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasIpi() bool`

HasIpi returns a boolean if a field has been set.

### GetReenlightenment

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReenlightenment() KubevirtIoApiCoreV1FeatureState`

GetReenlightenment returns the Reenlightenment field if non-nil, zero value otherwise.

### GetReenlightenmentOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReenlightenmentOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetReenlightenmentOk returns a tuple with the Reenlightenment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenlightenment

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetReenlightenment(v KubevirtIoApiCoreV1FeatureState)`

SetReenlightenment sets Reenlightenment field to given value.

### HasReenlightenment

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasReenlightenment() bool`

HasReenlightenment returns a boolean if a field has been set.

### GetRelaxed

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRelaxed() KubevirtIoApiCoreV1FeatureState`

GetRelaxed returns the Relaxed field if non-nil, zero value otherwise.

### GetRelaxedOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRelaxedOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetRelaxedOk returns a tuple with the Relaxed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelaxed

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetRelaxed(v KubevirtIoApiCoreV1FeatureState)`

SetRelaxed sets Relaxed field to given value.

### HasRelaxed

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasRelaxed() bool`

HasRelaxed returns a boolean if a field has been set.

### GetReset

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReset() KubevirtIoApiCoreV1FeatureState`

GetReset returns the Reset field if non-nil, zero value otherwise.

### GetResetOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetResetOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetResetOk returns a tuple with the Reset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReset

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetReset(v KubevirtIoApiCoreV1FeatureState)`

SetReset sets Reset field to given value.

### HasReset

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasReset() bool`

HasReset returns a boolean if a field has been set.

### GetRuntime

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRuntime() KubevirtIoApiCoreV1FeatureState`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRuntimeOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetRuntime(v KubevirtIoApiCoreV1FeatureState)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSpinlocks() KubevirtIoApiCoreV1FeatureSpinlocks`

GetSpinlocks returns the Spinlocks field if non-nil, zero value otherwise.

### GetSpinlocksOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSpinlocksOk() (*KubevirtIoApiCoreV1FeatureSpinlocks, bool)`

GetSpinlocksOk returns a tuple with the Spinlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSpinlocks(v KubevirtIoApiCoreV1FeatureSpinlocks)`

SetSpinlocks sets Spinlocks field to given value.

### HasSpinlocks

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSpinlocks() bool`

HasSpinlocks returns a boolean if a field has been set.

### GetSynic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynic() KubevirtIoApiCoreV1FeatureState`

GetSynic returns the Synic field if non-nil, zero value otherwise.

### GetSynicOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynicOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetSynicOk returns a tuple with the Synic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSynic(v KubevirtIoApiCoreV1FeatureState)`

SetSynic sets Synic field to given value.

### HasSynic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSynic() bool`

HasSynic returns a boolean if a field has been set.

### GetSynictimer

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynictimer() KubevirtIoApiCoreV1SyNICTimer`

GetSynictimer returns the Synictimer field if non-nil, zero value otherwise.

### GetSynictimerOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynictimerOk() (*KubevirtIoApiCoreV1SyNICTimer, bool)`

GetSynictimerOk returns a tuple with the Synictimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynictimer

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSynictimer(v KubevirtIoApiCoreV1SyNICTimer)`

SetSynictimer sets Synictimer field to given value.

### HasSynictimer

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSynictimer() bool`

HasSynictimer returns a boolean if a field has been set.

### GetTlbflush

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetTlbflush() KubevirtIoApiCoreV1FeatureState`

GetTlbflush returns the Tlbflush field if non-nil, zero value otherwise.

### GetTlbflushOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetTlbflushOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetTlbflushOk returns a tuple with the Tlbflush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlbflush

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetTlbflush(v KubevirtIoApiCoreV1FeatureState)`

SetTlbflush sets Tlbflush field to given value.

### HasTlbflush

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasTlbflush() bool`

HasTlbflush returns a boolean if a field has been set.

### GetVapic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVapic() KubevirtIoApiCoreV1FeatureState`

GetVapic returns the Vapic field if non-nil, zero value otherwise.

### GetVapicOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVapicOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetVapicOk returns a tuple with the Vapic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVapic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVapic(v KubevirtIoApiCoreV1FeatureState)`

SetVapic sets Vapic field to given value.

### HasVapic

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVapic() bool`

HasVapic returns a boolean if a field has been set.

### GetVendorid

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVendorid() KubevirtIoApiCoreV1FeatureVendorID`

GetVendorid returns the Vendorid field if non-nil, zero value otherwise.

### GetVendoridOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVendoridOk() (*KubevirtIoApiCoreV1FeatureVendorID, bool)`

GetVendoridOk returns a tuple with the Vendorid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorid

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVendorid(v KubevirtIoApiCoreV1FeatureVendorID)`

SetVendorid sets Vendorid field to given value.

### HasVendorid

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVendorid() bool`

HasVendorid returns a boolean if a field has been set.

### GetVpindex

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVpindex() KubevirtIoApiCoreV1FeatureState`

GetVpindex returns the Vpindex field if non-nil, zero value otherwise.

### GetVpindexOk

`func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVpindexOk() (*KubevirtIoApiCoreV1FeatureState, bool)`

GetVpindexOk returns a tuple with the Vpindex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpindex

`func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVpindex(v KubevirtIoApiCoreV1FeatureState)`

SetVpindex sets Vpindex field to given value.

### HasVpindex

`func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVpindex() bool`

HasVpindex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


