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

// checks if the HarvesterhciIoV1beta1UpgradeStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HarvesterhciIoV1beta1UpgradeStatus{}

// HarvesterhciIoV1beta1UpgradeStatus struct for HarvesterhciIoV1beta1UpgradeStatus
type HarvesterhciIoV1beta1UpgradeStatus struct {
	Conditions []HarvesterhciIoV1beta1Condition `json:"conditions,omitempty"`
	ImageID *string `json:"imageID,omitempty"`
	NodeStatuses *map[string]HarvesterhciIoV1beta1NodeUpgradeStatus `json:"nodeStatuses,omitempty"`
	PreviousVersion *string `json:"previousVersion,omitempty"`
	RepoInfo *string `json:"repoInfo,omitempty"`
	SingleNode *string `json:"singleNode,omitempty"`
	UpgradeLog *string `json:"upgradeLog,omitempty"`
}

// NewHarvesterhciIoV1beta1UpgradeStatus instantiates a new HarvesterhciIoV1beta1UpgradeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHarvesterhciIoV1beta1UpgradeStatus() *HarvesterhciIoV1beta1UpgradeStatus {
	this := HarvesterhciIoV1beta1UpgradeStatus{}
	return &this
}

// NewHarvesterhciIoV1beta1UpgradeStatusWithDefaults instantiates a new HarvesterhciIoV1beta1UpgradeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHarvesterhciIoV1beta1UpgradeStatusWithDefaults() *HarvesterhciIoV1beta1UpgradeStatus {
	this := HarvesterhciIoV1beta1UpgradeStatus{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetConditions() []HarvesterhciIoV1beta1Condition {
	if o == nil || IsNil(o.Conditions) {
		var ret []HarvesterhciIoV1beta1Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetConditionsOk() ([]HarvesterhciIoV1beta1Condition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []HarvesterhciIoV1beta1Condition and assigns it to the Conditions field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetConditions(v []HarvesterhciIoV1beta1Condition) {
	o.Conditions = v
}

// GetImageID returns the ImageID field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetImageID() string {
	if o == nil || IsNil(o.ImageID) {
		var ret string
		return ret
	}
	return *o.ImageID
}

// GetImageIDOk returns a tuple with the ImageID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetImageIDOk() (*string, bool) {
	if o == nil || IsNil(o.ImageID) {
		return nil, false
	}
	return o.ImageID, true
}

// HasImageID returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasImageID() bool {
	if o != nil && !IsNil(o.ImageID) {
		return true
	}

	return false
}

// SetImageID gets a reference to the given string and assigns it to the ImageID field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetImageID(v string) {
	o.ImageID = &v
}

// GetNodeStatuses returns the NodeStatuses field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetNodeStatuses() map[string]HarvesterhciIoV1beta1NodeUpgradeStatus {
	if o == nil || IsNil(o.NodeStatuses) {
		var ret map[string]HarvesterhciIoV1beta1NodeUpgradeStatus
		return ret
	}
	return *o.NodeStatuses
}

// GetNodeStatusesOk returns a tuple with the NodeStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetNodeStatusesOk() (*map[string]HarvesterhciIoV1beta1NodeUpgradeStatus, bool) {
	if o == nil || IsNil(o.NodeStatuses) {
		return nil, false
	}
	return o.NodeStatuses, true
}

// HasNodeStatuses returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasNodeStatuses() bool {
	if o != nil && !IsNil(o.NodeStatuses) {
		return true
	}

	return false
}

// SetNodeStatuses gets a reference to the given map[string]HarvesterhciIoV1beta1NodeUpgradeStatus and assigns it to the NodeStatuses field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetNodeStatuses(v map[string]HarvesterhciIoV1beta1NodeUpgradeStatus) {
	o.NodeStatuses = &v
}

// GetPreviousVersion returns the PreviousVersion field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetPreviousVersion() string {
	if o == nil || IsNil(o.PreviousVersion) {
		var ret string
		return ret
	}
	return *o.PreviousVersion
}

// GetPreviousVersionOk returns a tuple with the PreviousVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetPreviousVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousVersion) {
		return nil, false
	}
	return o.PreviousVersion, true
}

// HasPreviousVersion returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasPreviousVersion() bool {
	if o != nil && !IsNil(o.PreviousVersion) {
		return true
	}

	return false
}

// SetPreviousVersion gets a reference to the given string and assigns it to the PreviousVersion field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetPreviousVersion(v string) {
	o.PreviousVersion = &v
}

// GetRepoInfo returns the RepoInfo field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetRepoInfo() string {
	if o == nil || IsNil(o.RepoInfo) {
		var ret string
		return ret
	}
	return *o.RepoInfo
}

// GetRepoInfoOk returns a tuple with the RepoInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetRepoInfoOk() (*string, bool) {
	if o == nil || IsNil(o.RepoInfo) {
		return nil, false
	}
	return o.RepoInfo, true
}

// HasRepoInfo returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasRepoInfo() bool {
	if o != nil && !IsNil(o.RepoInfo) {
		return true
	}

	return false
}

// SetRepoInfo gets a reference to the given string and assigns it to the RepoInfo field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetRepoInfo(v string) {
	o.RepoInfo = &v
}

// GetSingleNode returns the SingleNode field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetSingleNode() string {
	if o == nil || IsNil(o.SingleNode) {
		var ret string
		return ret
	}
	return *o.SingleNode
}

// GetSingleNodeOk returns a tuple with the SingleNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetSingleNodeOk() (*string, bool) {
	if o == nil || IsNil(o.SingleNode) {
		return nil, false
	}
	return o.SingleNode, true
}

// HasSingleNode returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasSingleNode() bool {
	if o != nil && !IsNil(o.SingleNode) {
		return true
	}

	return false
}

// SetSingleNode gets a reference to the given string and assigns it to the SingleNode field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetSingleNode(v string) {
	o.SingleNode = &v
}

// GetUpgradeLog returns the UpgradeLog field value if set, zero value otherwise.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetUpgradeLog() string {
	if o == nil || IsNil(o.UpgradeLog) {
		var ret string
		return ret
	}
	return *o.UpgradeLog
}

// GetUpgradeLogOk returns a tuple with the UpgradeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) GetUpgradeLogOk() (*string, bool) {
	if o == nil || IsNil(o.UpgradeLog) {
		return nil, false
	}
	return o.UpgradeLog, true
}

// HasUpgradeLog returns a boolean if a field has been set.
func (o *HarvesterhciIoV1beta1UpgradeStatus) HasUpgradeLog() bool {
	if o != nil && !IsNil(o.UpgradeLog) {
		return true
	}

	return false
}

// SetUpgradeLog gets a reference to the given string and assigns it to the UpgradeLog field.
func (o *HarvesterhciIoV1beta1UpgradeStatus) SetUpgradeLog(v string) {
	o.UpgradeLog = &v
}

func (o HarvesterhciIoV1beta1UpgradeStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HarvesterhciIoV1beta1UpgradeStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ImageID) {
		toSerialize["imageID"] = o.ImageID
	}
	if !IsNil(o.NodeStatuses) {
		toSerialize["nodeStatuses"] = o.NodeStatuses
	}
	if !IsNil(o.PreviousVersion) {
		toSerialize["previousVersion"] = o.PreviousVersion
	}
	if !IsNil(o.RepoInfo) {
		toSerialize["repoInfo"] = o.RepoInfo
	}
	if !IsNil(o.SingleNode) {
		toSerialize["singleNode"] = o.SingleNode
	}
	if !IsNil(o.UpgradeLog) {
		toSerialize["upgradeLog"] = o.UpgradeLog
	}
	return toSerialize, nil
}

type NullableHarvesterhciIoV1beta1UpgradeStatus struct {
	value *HarvesterhciIoV1beta1UpgradeStatus
	isSet bool
}

func (v NullableHarvesterhciIoV1beta1UpgradeStatus) Get() *HarvesterhciIoV1beta1UpgradeStatus {
	return v.value
}

func (v *NullableHarvesterhciIoV1beta1UpgradeStatus) Set(val *HarvesterhciIoV1beta1UpgradeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHarvesterhciIoV1beta1UpgradeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHarvesterhciIoV1beta1UpgradeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHarvesterhciIoV1beta1UpgradeStatus(val *HarvesterhciIoV1beta1UpgradeStatus) *NullableHarvesterhciIoV1beta1UpgradeStatus {
	return &NullableHarvesterhciIoV1beta1UpgradeStatus{value: val, isSet: true}
}

func (v NullableHarvesterhciIoV1beta1UpgradeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHarvesterhciIoV1beta1UpgradeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


