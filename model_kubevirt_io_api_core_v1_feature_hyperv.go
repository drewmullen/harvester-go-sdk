/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1FeatureHyperv type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1FeatureHyperv{}

// KubevirtIoApiCoreV1FeatureHyperv struct for KubevirtIoApiCoreV1FeatureHyperv
type KubevirtIoApiCoreV1FeatureHyperv struct {
	Evmcs *KubevirtIoApiCoreV1FeatureState `json:"evmcs,omitempty"`
	Frequencies *KubevirtIoApiCoreV1FeatureState `json:"frequencies,omitempty"`
	Ipi *KubevirtIoApiCoreV1FeatureState `json:"ipi,omitempty"`
	Reenlightenment *KubevirtIoApiCoreV1FeatureState `json:"reenlightenment,omitempty"`
	Relaxed *KubevirtIoApiCoreV1FeatureState `json:"relaxed,omitempty"`
	Reset *KubevirtIoApiCoreV1FeatureState `json:"reset,omitempty"`
	Runtime *KubevirtIoApiCoreV1FeatureState `json:"runtime,omitempty"`
	Spinlocks *KubevirtIoApiCoreV1FeatureSpinlocks `json:"spinlocks,omitempty"`
	Synic *KubevirtIoApiCoreV1FeatureState `json:"synic,omitempty"`
	Synictimer *KubevirtIoApiCoreV1SyNICTimer `json:"synictimer,omitempty"`
	Tlbflush *KubevirtIoApiCoreV1FeatureState `json:"tlbflush,omitempty"`
	Vapic *KubevirtIoApiCoreV1FeatureState `json:"vapic,omitempty"`
	Vendorid *KubevirtIoApiCoreV1FeatureVendorID `json:"vendorid,omitempty"`
	Vpindex *KubevirtIoApiCoreV1FeatureState `json:"vpindex,omitempty"`
}

// NewKubevirtIoApiCoreV1FeatureHyperv instantiates a new KubevirtIoApiCoreV1FeatureHyperv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1FeatureHyperv() *KubevirtIoApiCoreV1FeatureHyperv {
	this := KubevirtIoApiCoreV1FeatureHyperv{}
	return &this
}

// NewKubevirtIoApiCoreV1FeatureHypervWithDefaults instantiates a new KubevirtIoApiCoreV1FeatureHyperv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1FeatureHypervWithDefaults() *KubevirtIoApiCoreV1FeatureHyperv {
	this := KubevirtIoApiCoreV1FeatureHyperv{}
	return &this
}

// GetEvmcs returns the Evmcs field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetEvmcs() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Evmcs) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Evmcs
}

// GetEvmcsOk returns a tuple with the Evmcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetEvmcsOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Evmcs) {
		return nil, false
	}
	return o.Evmcs, true
}

// HasEvmcs returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasEvmcs() bool {
	if o != nil && !IsNil(o.Evmcs) {
		return true
	}

	return false
}

// SetEvmcs gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Evmcs field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetEvmcs(v KubevirtIoApiCoreV1FeatureState) {
	o.Evmcs = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetFrequencies() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Frequencies) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetFrequenciesOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}
	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasFrequencies() bool {
	if o != nil && !IsNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Frequencies field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetFrequencies(v KubevirtIoApiCoreV1FeatureState) {
	o.Frequencies = &v
}

// GetIpi returns the Ipi field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetIpi() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Ipi) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Ipi
}

// GetIpiOk returns a tuple with the Ipi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetIpiOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Ipi) {
		return nil, false
	}
	return o.Ipi, true
}

// HasIpi returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasIpi() bool {
	if o != nil && !IsNil(o.Ipi) {
		return true
	}

	return false
}

// SetIpi gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Ipi field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetIpi(v KubevirtIoApiCoreV1FeatureState) {
	o.Ipi = &v
}

// GetReenlightenment returns the Reenlightenment field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReenlightenment() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Reenlightenment) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Reenlightenment
}

// GetReenlightenmentOk returns a tuple with the Reenlightenment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReenlightenmentOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Reenlightenment) {
		return nil, false
	}
	return o.Reenlightenment, true
}

// HasReenlightenment returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasReenlightenment() bool {
	if o != nil && !IsNil(o.Reenlightenment) {
		return true
	}

	return false
}

// SetReenlightenment gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Reenlightenment field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetReenlightenment(v KubevirtIoApiCoreV1FeatureState) {
	o.Reenlightenment = &v
}

// GetRelaxed returns the Relaxed field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRelaxed() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Relaxed) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Relaxed
}

// GetRelaxedOk returns a tuple with the Relaxed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRelaxedOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Relaxed) {
		return nil, false
	}
	return o.Relaxed, true
}

// HasRelaxed returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasRelaxed() bool {
	if o != nil && !IsNil(o.Relaxed) {
		return true
	}

	return false
}

// SetRelaxed gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Relaxed field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetRelaxed(v KubevirtIoApiCoreV1FeatureState) {
	o.Relaxed = &v
}

// GetReset returns the Reset field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetReset() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Reset) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Reset
}

// GetResetOk returns a tuple with the Reset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetResetOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Reset) {
		return nil, false
	}
	return o.Reset, true
}

// HasReset returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasReset() bool {
	if o != nil && !IsNil(o.Reset) {
		return true
	}

	return false
}

// SetReset gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Reset field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetReset(v KubevirtIoApiCoreV1FeatureState) {
	o.Reset = &v
}

// GetRuntime returns the Runtime field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRuntime() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Runtime) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Runtime
}

// GetRuntimeOk returns a tuple with the Runtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetRuntimeOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Runtime) {
		return nil, false
	}
	return o.Runtime, true
}

// HasRuntime returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasRuntime() bool {
	if o != nil && !IsNil(o.Runtime) {
		return true
	}

	return false
}

// SetRuntime gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Runtime field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetRuntime(v KubevirtIoApiCoreV1FeatureState) {
	o.Runtime = &v
}

// GetSpinlocks returns the Spinlocks field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSpinlocks() KubevirtIoApiCoreV1FeatureSpinlocks {
	if o == nil || IsNil(o.Spinlocks) {
		var ret KubevirtIoApiCoreV1FeatureSpinlocks
		return ret
	}
	return *o.Spinlocks
}

// GetSpinlocksOk returns a tuple with the Spinlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSpinlocksOk() (*KubevirtIoApiCoreV1FeatureSpinlocks, bool) {
	if o == nil || IsNil(o.Spinlocks) {
		return nil, false
	}
	return o.Spinlocks, true
}

// HasSpinlocks returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSpinlocks() bool {
	if o != nil && !IsNil(o.Spinlocks) {
		return true
	}

	return false
}

// SetSpinlocks gets a reference to the given KubevirtIoApiCoreV1FeatureSpinlocks and assigns it to the Spinlocks field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSpinlocks(v KubevirtIoApiCoreV1FeatureSpinlocks) {
	o.Spinlocks = &v
}

// GetSynic returns the Synic field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynic() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Synic) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Synic
}

// GetSynicOk returns a tuple with the Synic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynicOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Synic) {
		return nil, false
	}
	return o.Synic, true
}

// HasSynic returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSynic() bool {
	if o != nil && !IsNil(o.Synic) {
		return true
	}

	return false
}

// SetSynic gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Synic field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSynic(v KubevirtIoApiCoreV1FeatureState) {
	o.Synic = &v
}

// GetSynictimer returns the Synictimer field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynictimer() KubevirtIoApiCoreV1SyNICTimer {
	if o == nil || IsNil(o.Synictimer) {
		var ret KubevirtIoApiCoreV1SyNICTimer
		return ret
	}
	return *o.Synictimer
}

// GetSynictimerOk returns a tuple with the Synictimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetSynictimerOk() (*KubevirtIoApiCoreV1SyNICTimer, bool) {
	if o == nil || IsNil(o.Synictimer) {
		return nil, false
	}
	return o.Synictimer, true
}

// HasSynictimer returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasSynictimer() bool {
	if o != nil && !IsNil(o.Synictimer) {
		return true
	}

	return false
}

// SetSynictimer gets a reference to the given KubevirtIoApiCoreV1SyNICTimer and assigns it to the Synictimer field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetSynictimer(v KubevirtIoApiCoreV1SyNICTimer) {
	o.Synictimer = &v
}

// GetTlbflush returns the Tlbflush field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetTlbflush() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Tlbflush) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Tlbflush
}

// GetTlbflushOk returns a tuple with the Tlbflush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetTlbflushOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Tlbflush) {
		return nil, false
	}
	return o.Tlbflush, true
}

// HasTlbflush returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasTlbflush() bool {
	if o != nil && !IsNil(o.Tlbflush) {
		return true
	}

	return false
}

// SetTlbflush gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Tlbflush field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetTlbflush(v KubevirtIoApiCoreV1FeatureState) {
	o.Tlbflush = &v
}

// GetVapic returns the Vapic field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVapic() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Vapic) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Vapic
}

// GetVapicOk returns a tuple with the Vapic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVapicOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Vapic) {
		return nil, false
	}
	return o.Vapic, true
}

// HasVapic returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVapic() bool {
	if o != nil && !IsNil(o.Vapic) {
		return true
	}

	return false
}

// SetVapic gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Vapic field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVapic(v KubevirtIoApiCoreV1FeatureState) {
	o.Vapic = &v
}

// GetVendorid returns the Vendorid field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVendorid() KubevirtIoApiCoreV1FeatureVendorID {
	if o == nil || IsNil(o.Vendorid) {
		var ret KubevirtIoApiCoreV1FeatureVendorID
		return ret
	}
	return *o.Vendorid
}

// GetVendoridOk returns a tuple with the Vendorid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVendoridOk() (*KubevirtIoApiCoreV1FeatureVendorID, bool) {
	if o == nil || IsNil(o.Vendorid) {
		return nil, false
	}
	return o.Vendorid, true
}

// HasVendorid returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVendorid() bool {
	if o != nil && !IsNil(o.Vendorid) {
		return true
	}

	return false
}

// SetVendorid gets a reference to the given KubevirtIoApiCoreV1FeatureVendorID and assigns it to the Vendorid field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVendorid(v KubevirtIoApiCoreV1FeatureVendorID) {
	o.Vendorid = &v
}

// GetVpindex returns the Vpindex field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVpindex() KubevirtIoApiCoreV1FeatureState {
	if o == nil || IsNil(o.Vpindex) {
		var ret KubevirtIoApiCoreV1FeatureState
		return ret
	}
	return *o.Vpindex
}

// GetVpindexOk returns a tuple with the Vpindex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) GetVpindexOk() (*KubevirtIoApiCoreV1FeatureState, bool) {
	if o == nil || IsNil(o.Vpindex) {
		return nil, false
	}
	return o.Vpindex, true
}

// HasVpindex returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1FeatureHyperv) HasVpindex() bool {
	if o != nil && !IsNil(o.Vpindex) {
		return true
	}

	return false
}

// SetVpindex gets a reference to the given KubevirtIoApiCoreV1FeatureState and assigns it to the Vpindex field.
func (o *KubevirtIoApiCoreV1FeatureHyperv) SetVpindex(v KubevirtIoApiCoreV1FeatureState) {
	o.Vpindex = &v
}

func (o KubevirtIoApiCoreV1FeatureHyperv) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1FeatureHyperv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Evmcs) {
		toSerialize["evmcs"] = o.Evmcs
	}
	if !IsNil(o.Frequencies) {
		toSerialize["frequencies"] = o.Frequencies
	}
	if !IsNil(o.Ipi) {
		toSerialize["ipi"] = o.Ipi
	}
	if !IsNil(o.Reenlightenment) {
		toSerialize["reenlightenment"] = o.Reenlightenment
	}
	if !IsNil(o.Relaxed) {
		toSerialize["relaxed"] = o.Relaxed
	}
	if !IsNil(o.Reset) {
		toSerialize["reset"] = o.Reset
	}
	if !IsNil(o.Runtime) {
		toSerialize["runtime"] = o.Runtime
	}
	if !IsNil(o.Spinlocks) {
		toSerialize["spinlocks"] = o.Spinlocks
	}
	if !IsNil(o.Synic) {
		toSerialize["synic"] = o.Synic
	}
	if !IsNil(o.Synictimer) {
		toSerialize["synictimer"] = o.Synictimer
	}
	if !IsNil(o.Tlbflush) {
		toSerialize["tlbflush"] = o.Tlbflush
	}
	if !IsNil(o.Vapic) {
		toSerialize["vapic"] = o.Vapic
	}
	if !IsNil(o.Vendorid) {
		toSerialize["vendorid"] = o.Vendorid
	}
	if !IsNil(o.Vpindex) {
		toSerialize["vpindex"] = o.Vpindex
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1FeatureHyperv struct {
	value *KubevirtIoApiCoreV1FeatureHyperv
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1FeatureHyperv) Get() *KubevirtIoApiCoreV1FeatureHyperv {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1FeatureHyperv) Set(val *KubevirtIoApiCoreV1FeatureHyperv) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1FeatureHyperv) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1FeatureHyperv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1FeatureHyperv(val *KubevirtIoApiCoreV1FeatureHyperv) *NullableKubevirtIoApiCoreV1FeatureHyperv {
	return &NullableKubevirtIoApiCoreV1FeatureHyperv{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1FeatureHyperv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1FeatureHyperv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


