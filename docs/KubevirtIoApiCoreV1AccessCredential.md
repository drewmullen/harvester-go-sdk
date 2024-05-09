# KubevirtIoApiCoreV1AccessCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPublicKey** | Pointer to [**KubevirtIoApiCoreV1SSHPublicKeyAccessCredential**](KubevirtIoApiCoreV1SSHPublicKeyAccessCredential.md) | SSHPublicKey represents the source and method of applying a ssh public key into a guest virtual machine. | [optional] 
**UserPassword** | Pointer to [**KubevirtIoApiCoreV1UserPasswordAccessCredential**](KubevirtIoApiCoreV1UserPasswordAccessCredential.md) | UserPassword represents the source and method for applying a guest user&#39;s password | [optional] 

## Methods

### NewKubevirtIoApiCoreV1AccessCredential

`func NewKubevirtIoApiCoreV1AccessCredential() *KubevirtIoApiCoreV1AccessCredential`

NewKubevirtIoApiCoreV1AccessCredential instantiates a new KubevirtIoApiCoreV1AccessCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1AccessCredentialWithDefaults

`func NewKubevirtIoApiCoreV1AccessCredentialWithDefaults() *KubevirtIoApiCoreV1AccessCredential`

NewKubevirtIoApiCoreV1AccessCredentialWithDefaults instantiates a new KubevirtIoApiCoreV1AccessCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshPublicKey

`func (o *KubevirtIoApiCoreV1AccessCredential) GetSshPublicKey() KubevirtIoApiCoreV1SSHPublicKeyAccessCredential`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *KubevirtIoApiCoreV1AccessCredential) GetSshPublicKeyOk() (*KubevirtIoApiCoreV1SSHPublicKeyAccessCredential, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *KubevirtIoApiCoreV1AccessCredential) SetSshPublicKey(v KubevirtIoApiCoreV1SSHPublicKeyAccessCredential)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *KubevirtIoApiCoreV1AccessCredential) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### GetUserPassword

`func (o *KubevirtIoApiCoreV1AccessCredential) GetUserPassword() KubevirtIoApiCoreV1UserPasswordAccessCredential`

GetUserPassword returns the UserPassword field if non-nil, zero value otherwise.

### GetUserPasswordOk

`func (o *KubevirtIoApiCoreV1AccessCredential) GetUserPasswordOk() (*KubevirtIoApiCoreV1UserPasswordAccessCredential, bool)`

GetUserPasswordOk returns a tuple with the UserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPassword

`func (o *KubevirtIoApiCoreV1AccessCredential) SetUserPassword(v KubevirtIoApiCoreV1UserPasswordAccessCredential)`

SetUserPassword sets UserPassword field to given value.

### HasUserPassword

`func (o *KubevirtIoApiCoreV1AccessCredential) HasUserPassword() bool`

HasUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


