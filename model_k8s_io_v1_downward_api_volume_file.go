/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package harvester

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the K8sIoV1DownwardAPIVolumeFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1DownwardAPIVolumeFile{}

// K8sIoV1DownwardAPIVolumeFile DownwardAPIVolumeFile represents information to create the file containing the pod field
type K8sIoV1DownwardAPIVolumeFile struct {
	// Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.
	FieldRef *K8sIoV1ObjectFieldSelector `json:"fieldRef,omitempty"`
	// Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	Mode *int32 `json:"mode,omitempty"`
	// Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'
	Path string `json:"path"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.
	ResourceFieldRef *K8sIoV1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
}

type _K8sIoV1DownwardAPIVolumeFile K8sIoV1DownwardAPIVolumeFile

// NewK8sIoV1DownwardAPIVolumeFile instantiates a new K8sIoV1DownwardAPIVolumeFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1DownwardAPIVolumeFile(path string) *K8sIoV1DownwardAPIVolumeFile {
	this := K8sIoV1DownwardAPIVolumeFile{}
	this.Path = path
	return &this
}

// NewK8sIoV1DownwardAPIVolumeFileWithDefaults instantiates a new K8sIoV1DownwardAPIVolumeFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1DownwardAPIVolumeFileWithDefaults() *K8sIoV1DownwardAPIVolumeFile {
	this := K8sIoV1DownwardAPIVolumeFile{}
	var path string = ""
	this.Path = path
	return &this
}

// GetFieldRef returns the FieldRef field value if set, zero value otherwise.
func (o *K8sIoV1DownwardAPIVolumeFile) GetFieldRef() K8sIoV1ObjectFieldSelector {
	if o == nil || IsNil(o.FieldRef) {
		var ret K8sIoV1ObjectFieldSelector
		return ret
	}
	return *o.FieldRef
}

// GetFieldRefOk returns a tuple with the FieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) GetFieldRefOk() (*K8sIoV1ObjectFieldSelector, bool) {
	if o == nil || IsNil(o.FieldRef) {
		return nil, false
	}
	return o.FieldRef, true
}

// HasFieldRef returns a boolean if a field has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) HasFieldRef() bool {
	if o != nil && !IsNil(o.FieldRef) {
		return true
	}

	return false
}

// SetFieldRef gets a reference to the given K8sIoV1ObjectFieldSelector and assigns it to the FieldRef field.
func (o *K8sIoV1DownwardAPIVolumeFile) SetFieldRef(v K8sIoV1ObjectFieldSelector) {
	o.FieldRef = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *K8sIoV1DownwardAPIVolumeFile) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *K8sIoV1DownwardAPIVolumeFile) SetMode(v int32) {
	o.Mode = &v
}

// GetPath returns the Path field value
func (o *K8sIoV1DownwardAPIVolumeFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *K8sIoV1DownwardAPIVolumeFile) SetPath(v string) {
	o.Path = v
}

// GetResourceFieldRef returns the ResourceFieldRef field value if set, zero value otherwise.
func (o *K8sIoV1DownwardAPIVolumeFile) GetResourceFieldRef() K8sIoV1ResourceFieldSelector {
	if o == nil || IsNil(o.ResourceFieldRef) {
		var ret K8sIoV1ResourceFieldSelector
		return ret
	}
	return *o.ResourceFieldRef
}

// GetResourceFieldRefOk returns a tuple with the ResourceFieldRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) GetResourceFieldRefOk() (*K8sIoV1ResourceFieldSelector, bool) {
	if o == nil || IsNil(o.ResourceFieldRef) {
		return nil, false
	}
	return o.ResourceFieldRef, true
}

// HasResourceFieldRef returns a boolean if a field has been set.
func (o *K8sIoV1DownwardAPIVolumeFile) HasResourceFieldRef() bool {
	if o != nil && !IsNil(o.ResourceFieldRef) {
		return true
	}

	return false
}

// SetResourceFieldRef gets a reference to the given K8sIoV1ResourceFieldSelector and assigns it to the ResourceFieldRef field.
func (o *K8sIoV1DownwardAPIVolumeFile) SetResourceFieldRef(v K8sIoV1ResourceFieldSelector) {
	o.ResourceFieldRef = &v
}

func (o K8sIoV1DownwardAPIVolumeFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1DownwardAPIVolumeFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldRef) {
		toSerialize["fieldRef"] = o.FieldRef
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.ResourceFieldRef) {
		toSerialize["resourceFieldRef"] = o.ResourceFieldRef
	}
	return toSerialize, nil
}

func (o *K8sIoV1DownwardAPIVolumeFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"path",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varK8sIoV1DownwardAPIVolumeFile := _K8sIoV1DownwardAPIVolumeFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varK8sIoV1DownwardAPIVolumeFile)

	if err != nil {
		return err
	}

	*o = K8sIoV1DownwardAPIVolumeFile(varK8sIoV1DownwardAPIVolumeFile)

	return err
}

type NullableK8sIoV1DownwardAPIVolumeFile struct {
	value *K8sIoV1DownwardAPIVolumeFile
	isSet bool
}

func (v NullableK8sIoV1DownwardAPIVolumeFile) Get() *K8sIoV1DownwardAPIVolumeFile {
	return v.value
}

func (v *NullableK8sIoV1DownwardAPIVolumeFile) Set(val *K8sIoV1DownwardAPIVolumeFile) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1DownwardAPIVolumeFile) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1DownwardAPIVolumeFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1DownwardAPIVolumeFile(val *K8sIoV1DownwardAPIVolumeFile) *NullableK8sIoV1DownwardAPIVolumeFile {
	return &NullableK8sIoV1DownwardAPIVolumeFile{value: val, isSet: true}
}

func (v NullableK8sIoV1DownwardAPIVolumeFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1DownwardAPIVolumeFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


