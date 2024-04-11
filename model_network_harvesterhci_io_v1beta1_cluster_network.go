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

// checks if the NetworkHarvesterhciIoV1beta1ClusterNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkHarvesterhciIoV1beta1ClusterNetwork{}

// NetworkHarvesterhciIoV1beta1ClusterNetwork struct for NetworkHarvesterhciIoV1beta1ClusterNetwork
type NetworkHarvesterhciIoV1beta1ClusterNetwork struct {
	ApiVersion *string `json:"apiVersion,omitempty"`
	Kind *string `json:"kind,omitempty"`
	Metadata *K8sIoV1ObjectMeta `json:"metadata,omitempty"`
	Status *NetworkHarvesterhciIoV1beta1ClusterNetworkStatus `json:"status,omitempty"`
}

// NewNetworkHarvesterhciIoV1beta1ClusterNetwork instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkHarvesterhciIoV1beta1ClusterNetwork() *NetworkHarvesterhciIoV1beta1ClusterNetwork {
	this := NetworkHarvesterhciIoV1beta1ClusterNetwork{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var status NetworkHarvesterhciIoV1beta1ClusterNetworkStatus
	this.Status = &status
	return &this
}

// NewNetworkHarvesterhciIoV1beta1ClusterNetworkWithDefaults instantiates a new NetworkHarvesterhciIoV1beta1ClusterNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkHarvesterhciIoV1beta1ClusterNetworkWithDefaults() *NetworkHarvesterhciIoV1beta1ClusterNetwork {
	this := NetworkHarvesterhciIoV1beta1ClusterNetwork{}
	var metadata K8sIoV1ObjectMeta
	this.Metadata = &metadata
	var status NetworkHarvesterhciIoV1beta1ClusterNetworkStatus
	this.Status = &status
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetMetadata() K8sIoV1ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret K8sIoV1ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetMetadataOk() (*K8sIoV1ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given K8sIoV1ObjectMeta and assigns it to the Metadata field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetMetadata(v K8sIoV1ObjectMeta) {
	o.Metadata = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetStatus() NetworkHarvesterhciIoV1beta1ClusterNetworkStatus {
	if o == nil || IsNil(o.Status) {
		var ret NetworkHarvesterhciIoV1beta1ClusterNetworkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) GetStatusOk() (*NetworkHarvesterhciIoV1beta1ClusterNetworkStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NetworkHarvesterhciIoV1beta1ClusterNetworkStatus and assigns it to the Status field.
func (o *NetworkHarvesterhciIoV1beta1ClusterNetwork) SetStatus(v NetworkHarvesterhciIoV1beta1ClusterNetworkStatus) {
	o.Status = &v
}

func (o NetworkHarvesterhciIoV1beta1ClusterNetwork) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkHarvesterhciIoV1beta1ClusterNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableNetworkHarvesterhciIoV1beta1ClusterNetwork struct {
	value *NetworkHarvesterhciIoV1beta1ClusterNetwork
	isSet bool
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) Get() *NetworkHarvesterhciIoV1beta1ClusterNetwork {
	return v.value
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) Set(val *NetworkHarvesterhciIoV1beta1ClusterNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkHarvesterhciIoV1beta1ClusterNetwork(val *NetworkHarvesterhciIoV1beta1ClusterNetwork) *NullableNetworkHarvesterhciIoV1beta1ClusterNetwork {
	return &NullableNetworkHarvesterhciIoV1beta1ClusterNetwork{value: val, isSet: true}
}

func (v NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkHarvesterhciIoV1beta1ClusterNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


