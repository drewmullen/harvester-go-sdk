# KubevirtIoApiCoreV1SEV

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attestation** | Pointer to **map[string]interface{}** | If specified, run the attestation process for a vmi. | [optional] 
**DhCert** | Pointer to **string** | Base64 encoded guest owner&#39;s Diffie-Hellman key. | [optional] 
**Policy** | Pointer to [**KubevirtIoApiCoreV1SEVPolicy**](KubevirtIoApiCoreV1SEVPolicy.md) | Guest policy flags as defined in AMD SEV API specification. Note: due to security reasons it is not allowed to enable guest debugging. Therefore NoDebug flag is not exposed to users and is always true. | [optional] 
**Session** | Pointer to **string** | Base64 encoded session blob. | [optional] 

## Methods

### NewKubevirtIoApiCoreV1SEV

`func NewKubevirtIoApiCoreV1SEV() *KubevirtIoApiCoreV1SEV`

NewKubevirtIoApiCoreV1SEV instantiates a new KubevirtIoApiCoreV1SEV object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1SEVWithDefaults

`func NewKubevirtIoApiCoreV1SEVWithDefaults() *KubevirtIoApiCoreV1SEV`

NewKubevirtIoApiCoreV1SEVWithDefaults instantiates a new KubevirtIoApiCoreV1SEV object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttestation

`func (o *KubevirtIoApiCoreV1SEV) GetAttestation() map[string]interface{}`

GetAttestation returns the Attestation field if non-nil, zero value otherwise.

### GetAttestationOk

`func (o *KubevirtIoApiCoreV1SEV) GetAttestationOk() (*map[string]interface{}, bool)`

GetAttestationOk returns a tuple with the Attestation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestation

`func (o *KubevirtIoApiCoreV1SEV) SetAttestation(v map[string]interface{})`

SetAttestation sets Attestation field to given value.

### HasAttestation

`func (o *KubevirtIoApiCoreV1SEV) HasAttestation() bool`

HasAttestation returns a boolean if a field has been set.

### GetDhCert

`func (o *KubevirtIoApiCoreV1SEV) GetDhCert() string`

GetDhCert returns the DhCert field if non-nil, zero value otherwise.

### GetDhCertOk

`func (o *KubevirtIoApiCoreV1SEV) GetDhCertOk() (*string, bool)`

GetDhCertOk returns a tuple with the DhCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhCert

`func (o *KubevirtIoApiCoreV1SEV) SetDhCert(v string)`

SetDhCert sets DhCert field to given value.

### HasDhCert

`func (o *KubevirtIoApiCoreV1SEV) HasDhCert() bool`

HasDhCert returns a boolean if a field has been set.

### GetPolicy

`func (o *KubevirtIoApiCoreV1SEV) GetPolicy() KubevirtIoApiCoreV1SEVPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *KubevirtIoApiCoreV1SEV) GetPolicyOk() (*KubevirtIoApiCoreV1SEVPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *KubevirtIoApiCoreV1SEV) SetPolicy(v KubevirtIoApiCoreV1SEVPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *KubevirtIoApiCoreV1SEV) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetSession

`func (o *KubevirtIoApiCoreV1SEV) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *KubevirtIoApiCoreV1SEV) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *KubevirtIoApiCoreV1SEV) SetSession(v string)`

SetSession sets Session field to given value.

### HasSession

`func (o *KubevirtIoApiCoreV1SEV) HasSession() bool`

HasSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


