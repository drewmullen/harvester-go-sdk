/*
Harvester APIs

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3{}

// KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 struct for KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3
type KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 struct {
	CertConfigMap *string `json:"certConfigMap,omitempty"`
	SecretRef *string `json:"secretRef,omitempty"`
	Url string `json:"url"`
}

type _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3(url string) *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3{}
	this.Url = url
	return &this
}

// NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3WithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3WithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 {
	this := KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3{}
	var url string = ""
	this.Url = url
	return &this
}

// GetCertConfigMap returns the CertConfigMap field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetCertConfigMap() string {
	if o == nil || IsNil(o.CertConfigMap) {
		var ret string
		return ret
	}
	return *o.CertConfigMap
}

// GetCertConfigMapOk returns a tuple with the CertConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetCertConfigMapOk() (*string, bool) {
	if o == nil || IsNil(o.CertConfigMap) {
		return nil, false
	}
	return o.CertConfigMap, true
}

// HasCertConfigMap returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) HasCertConfigMap() bool {
	if o != nil && !IsNil(o.CertConfigMap) {
		return true
	}

	return false
}

// SetCertConfigMap gets a reference to the given string and assigns it to the CertConfigMap field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) SetCertConfigMap(v string) {
	o.CertConfigMap = &v
}

// GetSecretRef returns the SecretRef field value if set, zero value otherwise.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetSecretRef() string {
	if o == nil || IsNil(o.SecretRef) {
		var ret string
		return ret
	}
	return *o.SecretRef
}

// GetSecretRefOk returns a tuple with the SecretRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetSecretRefOk() (*string, bool) {
	if o == nil || IsNil(o.SecretRef) {
		return nil, false
	}
	return o.SecretRef, true
}

// HasSecretRef returns a boolean if a field has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) HasSecretRef() bool {
	if o != nil && !IsNil(o.SecretRef) {
		return true
	}

	return false
}

// SetSecretRef gets a reference to the given string and assigns it to the SecretRef field.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) SetSecretRef(v string) {
	o.SecretRef = &v
}

// GetUrl returns the Url field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) SetUrl(v string) {
	o.Url = v
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CertConfigMap) {
		toSerialize["certConfigMap"] = o.CertConfigMap
	}
	if !IsNil(o.SecretRef) {
		toSerialize["secretRef"] = o.SecretRef
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 := _KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3)

	if err != nil {
		return err
	}

	*o = KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3(varKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3)

	return err
}

type NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 struct {
	value *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3
	isSet bool
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) Get() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 {
	return v.value
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) Set(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) {
	v.value = val
	v.isSet = true
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) IsSet() bool {
	return v.isSet
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3(val *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3 {
	return &NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3{value: val, isSet: true}
}

func (v NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


