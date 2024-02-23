/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the KubevirtIoApiCoreV1PersistentVolumeClaimInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoApiCoreV1PersistentVolumeClaimInfo{}

// KubevirtIoApiCoreV1PersistentVolumeClaimInfo struct for KubevirtIoApiCoreV1PersistentVolumeClaimInfo
type KubevirtIoApiCoreV1PersistentVolumeClaimInfo struct {
	AccessModes []string `json:"accessModes,omitempty"`
	Capacity *map[string]string `json:"capacity,omitempty"`
	FilesystemOverhead *string `json:"filesystemOverhead,omitempty"`
	Preallocated *bool `json:"preallocated,omitempty"`
	Requests *map[string]string `json:"requests,omitempty"`
	VolumeMode *string `json:"volumeMode,omitempty"`
}

// NewKubevirtIoApiCoreV1PersistentVolumeClaimInfo instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoApiCoreV1PersistentVolumeClaimInfo() *KubevirtIoApiCoreV1PersistentVolumeClaimInfo {
	this := KubevirtIoApiCoreV1PersistentVolumeClaimInfo{}
	return &this
}

// NewKubevirtIoApiCoreV1PersistentVolumeClaimInfoWithDefaults instantiates a new KubevirtIoApiCoreV1PersistentVolumeClaimInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoApiCoreV1PersistentVolumeClaimInfoWithDefaults() *KubevirtIoApiCoreV1PersistentVolumeClaimInfo {
	this := KubevirtIoApiCoreV1PersistentVolumeClaimInfo{}
	return &this
}

// GetAccessModes returns the AccessModes field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetAccessModes() []string {
	if o == nil || IsNil(o.AccessModes) {
		var ret []string
		return ret
	}
	return o.AccessModes
}

// GetAccessModesOk returns a tuple with the AccessModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetAccessModesOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessModes) {
		return nil, false
	}
	return o.AccessModes, true
}

// HasAccessModes returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasAccessModes() bool {
	if o != nil && !IsNil(o.AccessModes) {
		return true
	}

	return false
}

// SetAccessModes gets a reference to the given []string and assigns it to the AccessModes field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetAccessModes(v []string) {
	o.AccessModes = v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetCapacity() map[string]string {
	if o == nil || IsNil(o.Capacity) {
		var ret map[string]string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetCapacityOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Capacity) {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasCapacity() bool {
	if o != nil && !IsNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given map[string]string and assigns it to the Capacity field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetCapacity(v map[string]string) {
	o.Capacity = &v
}

// GetFilesystemOverhead returns the FilesystemOverhead field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetFilesystemOverhead() string {
	if o == nil || IsNil(o.FilesystemOverhead) {
		var ret string
		return ret
	}
	return *o.FilesystemOverhead
}

// GetFilesystemOverheadOk returns a tuple with the FilesystemOverhead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetFilesystemOverheadOk() (*string, bool) {
	if o == nil || IsNil(o.FilesystemOverhead) {
		return nil, false
	}
	return o.FilesystemOverhead, true
}

// HasFilesystemOverhead returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasFilesystemOverhead() bool {
	if o != nil && !IsNil(o.FilesystemOverhead) {
		return true
	}

	return false
}

// SetFilesystemOverhead gets a reference to the given string and assigns it to the FilesystemOverhead field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetFilesystemOverhead(v string) {
	o.FilesystemOverhead = &v
}

// GetPreallocated returns the Preallocated field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetPreallocated() bool {
	if o == nil || IsNil(o.Preallocated) {
		var ret bool
		return ret
	}
	return *o.Preallocated
}

// GetPreallocatedOk returns a tuple with the Preallocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetPreallocatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Preallocated) {
		return nil, false
	}
	return o.Preallocated, true
}

// HasPreallocated returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasPreallocated() bool {
	if o != nil && !IsNil(o.Preallocated) {
		return true
	}

	return false
}

// SetPreallocated gets a reference to the given bool and assigns it to the Preallocated field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetPreallocated(v bool) {
	o.Preallocated = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetRequests() map[string]string {
	if o == nil || IsNil(o.Requests) {
		var ret map[string]string
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetRequestsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given map[string]string and assigns it to the Requests field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetRequests(v map[string]string) {
	o.Requests = &v
}

// GetVolumeMode returns the VolumeMode field value if set, zero value otherwise.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetVolumeMode() string {
	if o == nil || IsNil(o.VolumeMode) {
		var ret string
		return ret
	}
	return *o.VolumeMode
}

// GetVolumeModeOk returns a tuple with the VolumeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) GetVolumeModeOk() (*string, bool) {
	if o == nil || IsNil(o.VolumeMode) {
		return nil, false
	}
	return o.VolumeMode, true
}

// HasVolumeMode returns a boolean if a field has been set.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) HasVolumeMode() bool {
	if o != nil && !IsNil(o.VolumeMode) {
		return true
	}

	return false
}

// SetVolumeMode gets a reference to the given string and assigns it to the VolumeMode field.
func (o *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) SetVolumeMode(v string) {
	o.VolumeMode = &v
}

func (o KubevirtIoApiCoreV1PersistentVolumeClaimInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoApiCoreV1PersistentVolumeClaimInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessModes) {
		toSerialize["accessModes"] = o.AccessModes
	}
	if !IsNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !IsNil(o.FilesystemOverhead) {
		toSerialize["filesystemOverhead"] = o.FilesystemOverhead
	}
	if !IsNil(o.Preallocated) {
		toSerialize["preallocated"] = o.Preallocated
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.VolumeMode) {
		toSerialize["volumeMode"] = o.VolumeMode
	}
	return toSerialize, nil
}

type NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo struct {
	value *KubevirtIoApiCoreV1PersistentVolumeClaimInfo
	isSet bool
}

func (v NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) Get() *KubevirtIoApiCoreV1PersistentVolumeClaimInfo {
	return v.value
}

func (v *NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) Set(val *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo(val *KubevirtIoApiCoreV1PersistentVolumeClaimInfo) *NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo {
	return &NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo{value: val, isSet: true}
}

func (v NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoApiCoreV1PersistentVolumeClaimInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


