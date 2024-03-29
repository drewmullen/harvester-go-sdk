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

// checks if the HarvesterhciIoV1beta1VirtualMachineImageStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1VirtualMachineImageStatus{}

// HarvesterhciIoV1beta1VirtualMachineImageStatus struct for HarvesterhciIoV1beta1VirtualMachineImageStatus
type HarvesterhciIoV1beta1VirtualMachineImageStatus struct {
	AppliedUrl *string `json:"appliedUrl,omitempty"`
	Conditions []HarvesterhciIoV1beta1Condition `json:"conditions,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
	LastFailedTime *string `json:"lastFailedTime,omitempty"`
	Progress *int32 `json:"progress,omitempty"`
	Size *int64 `json:"size,omitempty"`
	StorageClassName *string `json:"storageClassName,omitempty"`
}

// NewHarvesterhciIoV1beta1VirtualMachineImageStatus instantiates a new HarvesterhciIoV1beta1VirtualMachineImageStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1VirtualMachineImageStatus() *HarvesterhciIoV1beta1VirtualMachineImageStatus {
	this := HarvesterhciIoV1beta1VirtualMachineImageStatus{}
	var failed int32 = 0
	this.Failed = &failed
	return &this
}

// NewHarvesterhciIoV1beta1VirtualMachineImageStatusWithDefaults instantiates a new HarvesterhciIoV1beta1VirtualMachineImageStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1VirtualMachineImageStatusWithDefaults() *HarvesterhciIoV1beta1VirtualMachineImageStatus {
	this := HarvesterhciIoV1beta1VirtualMachineImageStatus{}
	var failed int32 = 0
	this.Failed = &failed
	return &this
}

// GetAppliedUrl returns the AppliedUrl field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetAppliedUrl() string {
	if o == nil || IsNil(o.AppliedUrl) {
		var ret string
		return ret
	}
	return *o.AppliedUrl
}

// GetAppliedUrlOk returns a tuple with the AppliedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetAppliedUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AppliedUrl) {
		return nil, false
	}
	return o.AppliedUrl, true
}

// HasAppliedUrl returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasAppliedUrl() bool {
	if o != nil && !IsNil(o.AppliedUrl) {
		return true
	}

	return false
}

// SetAppliedUrl gets a reference to the given string and assigns it to the AppliedUrl field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetAppliedUrl(v string) {
	o.AppliedUrl = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetConditions() []HarvesterhciIoV1beta1Condition {
	if o == nil || IsNil(o.Conditions) {
		var ret []HarvesterhciIoV1beta1Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetConditionsOk() ([]HarvesterhciIoV1beta1Condition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []HarvesterhciIoV1beta1Condition and assigns it to the Conditions field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetConditions(v []HarvesterhciIoV1beta1Condition) {
	o.Conditions = v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetFailed(v int32) {
	o.Failed = &v
}

// GetLastFailedTime returns the LastFailedTime field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetLastFailedTime() string {
	if o == nil || IsNil(o.LastFailedTime) {
		var ret string
		return ret
	}
	return *o.LastFailedTime
}

// GetLastFailedTimeOk returns a tuple with the LastFailedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetLastFailedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastFailedTime) {
		return nil, false
	}
	return o.LastFailedTime, true
}

// HasLastFailedTime returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasLastFailedTime() bool {
	if o != nil && !IsNil(o.LastFailedTime) {
		return true
	}

	return false
}

// SetLastFailedTime gets a reference to the given string and assigns it to the LastFailedTime field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetLastFailedTime(v string) {
	o.LastFailedTime = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetProgress(v int32) {
	o.Progress = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetSize(v int64) {
	o.Size = &v
}

// GetStorageClassName returns the StorageClassName field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetStorageClassName() string {
	if o == nil || IsNil(o.StorageClassName) {
		var ret string
		return ret
	}
	return *o.StorageClassName
}

// GetStorageClassNameOk returns a tuple with the StorageClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) GetStorageClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.StorageClassName) {
		return nil, false
	}
	return o.StorageClassName, true
}

// HasStorageClassName returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) HasStorageClassName() bool {
	if o != nil && !IsNil(o.StorageClassName) {
		return true
	}

	return false
}

// SetStorageClassName gets a reference to the given string and assigns it to the StorageClassName field.
func (o *HarvesterhciIoV1beta1VirtualMachineImageStatus) SetStorageClassName(v string) {
	o.StorageClassName = &v
}

func (o HarvesterhciIoV1beta1VirtualMachineImageStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1VirtualMachineImageStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppliedUrl) {
		toSerialize["appliedUrl"] = o.AppliedUrl
	}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.LastFailedTime) {
		toSerialize["lastFailedTime"] = o.LastFailedTime
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.StorageClassName) {
		toSerialize["storageClassName"] = o.StorageClassName
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1VirtualMachineImageStatus struct {
	value *HarvesterhciIoV1beta1VirtualMachineImageStatus
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) Get() *HarvesterhciIoV1beta1VirtualMachineImageStatus {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) Set(val *HarvesterhciIoV1beta1VirtualMachineImageStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1VirtualMachineImageStatus(val *HarvesterhciIoV1beta1VirtualMachineImageStatus) *NullableHarvesterhciIoV1beta1VirtualMachineImageStatus {
	return &NullableHarvesterhciIoV1beta1VirtualMachineImageStatus{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1VirtualMachineImageStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


