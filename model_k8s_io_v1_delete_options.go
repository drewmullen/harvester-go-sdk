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

// checks if the K8sIoV1DeleteOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &K8sIoV1DeleteOptions{}

// K8sIoV1DeleteOptions DeleteOptions may be provided when deleting an API object.
type K8sIoV1DeleteOptions struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// When present, indicates that modifications should not be persisted. An invalid or unrecognized dryRun directive will result in an error response and no further processing of the request. Valid values are: - All: all dry run stages will be processed
	DryRun []string `json:"dryRun,omitempty"`
	// The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
	GracePeriodSeconds *int64 `json:"gracePeriodSeconds,omitempty"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both.
	OrphanDependents *bool `json:"orphanDependents,omitempty"`
	// Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.
	Preconditions *K8sIoV1Preconditions `json:"preconditions,omitempty"`
	// Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground.
	PropagationPolicy *string `json:"propagationPolicy,omitempty"`
}

// NewK8sIoV1DeleteOptions instantiates a new K8sIoV1DeleteOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewK8sIoV1DeleteOptions() *K8sIoV1DeleteOptions {
	this := K8sIoV1DeleteOptions{}
	return &this
}

// NewK8sIoV1DeleteOptionsWithDefaults instantiates a new K8sIoV1DeleteOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewK8sIoV1DeleteOptionsWithDefaults() *K8sIoV1DeleteOptions {
	this := K8sIoV1DeleteOptions{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *K8sIoV1DeleteOptions) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetDryRun() []string {
	if o == nil || IsNil(o.DryRun) {
		var ret []string
		return ret
	}
	return o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetDryRunOk() ([]string, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}
	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given []string and assigns it to the DryRun field.
func (o *K8sIoV1DeleteOptions) SetDryRun(v []string) {
	o.DryRun = v
}

// GetGracePeriodSeconds returns the GracePeriodSeconds field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetGracePeriodSeconds() int64 {
	if o == nil || IsNil(o.GracePeriodSeconds) {
		var ret int64
		return ret
	}
	return *o.GracePeriodSeconds
}

// GetGracePeriodSecondsOk returns a tuple with the GracePeriodSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetGracePeriodSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.GracePeriodSeconds) {
		return nil, false
	}
	return o.GracePeriodSeconds, true
}

// HasGracePeriodSeconds returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasGracePeriodSeconds() bool {
	if o != nil && !IsNil(o.GracePeriodSeconds) {
		return true
	}

	return false
}

// SetGracePeriodSeconds gets a reference to the given int64 and assigns it to the GracePeriodSeconds field.
func (o *K8sIoV1DeleteOptions) SetGracePeriodSeconds(v int64) {
	o.GracePeriodSeconds = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *K8sIoV1DeleteOptions) SetKind(v string) {
	o.Kind = &v
}

// GetOrphanDependents returns the OrphanDependents field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetOrphanDependents() bool {
	if o == nil || IsNil(o.OrphanDependents) {
		var ret bool
		return ret
	}
	return *o.OrphanDependents
}

// GetOrphanDependentsOk returns a tuple with the OrphanDependents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetOrphanDependentsOk() (*bool, bool) {
	if o == nil || IsNil(o.OrphanDependents) {
		return nil, false
	}
	return o.OrphanDependents, true
}

// HasOrphanDependents returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasOrphanDependents() bool {
	if o != nil && !IsNil(o.OrphanDependents) {
		return true
	}

	return false
}

// SetOrphanDependents gets a reference to the given bool and assigns it to the OrphanDependents field.
func (o *K8sIoV1DeleteOptions) SetOrphanDependents(v bool) {
	o.OrphanDependents = &v
}

// GetPreconditions returns the Preconditions field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetPreconditions() K8sIoV1Preconditions {
	if o == nil || IsNil(o.Preconditions) {
		var ret K8sIoV1Preconditions
		return ret
	}
	return *o.Preconditions
}

// GetPreconditionsOk returns a tuple with the Preconditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetPreconditionsOk() (*K8sIoV1Preconditions, bool) {
	if o == nil || IsNil(o.Preconditions) {
		return nil, false
	}
	return o.Preconditions, true
}

// HasPreconditions returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasPreconditions() bool {
	if o != nil && !IsNil(o.Preconditions) {
		return true
	}

	return false
}

// SetPreconditions gets a reference to the given K8sIoV1Preconditions and assigns it to the Preconditions field.
func (o *K8sIoV1DeleteOptions) SetPreconditions(v K8sIoV1Preconditions) {
	o.Preconditions = &v
}

// GetPropagationPolicy returns the PropagationPolicy field value if set, zero value otherwise.
func (o *K8sIoV1DeleteOptions) GetPropagationPolicy() string {
	if o == nil || IsNil(o.PropagationPolicy) {
		var ret string
		return ret
	}
	return *o.PropagationPolicy
}

// GetPropagationPolicyOk returns a tuple with the PropagationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *K8sIoV1DeleteOptions) GetPropagationPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PropagationPolicy) {
		return nil, false
	}
	return o.PropagationPolicy, true
}

// HasPropagationPolicy returns a boolean if a field has been set.
func (o *K8sIoV1DeleteOptions) HasPropagationPolicy() bool {
	if o != nil && !IsNil(o.PropagationPolicy) {
		return true
	}

	return false
}

// SetPropagationPolicy gets a reference to the given string and assigns it to the PropagationPolicy field.
func (o *K8sIoV1DeleteOptions) SetPropagationPolicy(v string) {
	o.PropagationPolicy = &v
}

func (o K8sIoV1DeleteOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o K8sIoV1DeleteOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.DryRun) {
		toSerialize["dryRun"] = o.DryRun
	}
	if !IsNil(o.GracePeriodSeconds) {
		toSerialize["gracePeriodSeconds"] = o.GracePeriodSeconds
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.OrphanDependents) {
		toSerialize["orphanDependents"] = o.OrphanDependents
	}
	if !IsNil(o.Preconditions) {
		toSerialize["preconditions"] = o.Preconditions
	}
	if !IsNil(o.PropagationPolicy) {
		toSerialize["propagationPolicy"] = o.PropagationPolicy
	}
	return toSerialize, nil
}

type NullableK8sIoV1DeleteOptions struct {
	value *K8sIoV1DeleteOptions
	isSet bool
}

func (v NullableK8sIoV1DeleteOptions) Get() *K8sIoV1DeleteOptions {
	return v.value
}

func (v *NullableK8sIoV1DeleteOptions) Set(val *K8sIoV1DeleteOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableK8sIoV1DeleteOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableK8sIoV1DeleteOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableK8sIoV1DeleteOptions(val *K8sIoV1DeleteOptions) *NullableK8sIoV1DeleteOptions {
	return &NullableK8sIoV1DeleteOptions{value: val, isSet: true}
}

func (v NullableK8sIoV1DeleteOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableK8sIoV1DeleteOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


