# K8sIoV1PodDNSConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nameservers** | Pointer to **[]string** | A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed. | [optional] 
**Options** | Pointer to [**[]K8sIoV1PodDNSConfigOption**](K8sIoV1PodDNSConfigOption.md) | A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy. | [optional] 
**Searches** | Pointer to **[]string** | A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed. | [optional] 

## Methods

### NewK8sIoV1PodDNSConfig

`func NewK8sIoV1PodDNSConfig() *K8sIoV1PodDNSConfig`

NewK8sIoV1PodDNSConfig instantiates a new K8sIoV1PodDNSConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewK8sIoV1PodDNSConfigWithDefaults

`func NewK8sIoV1PodDNSConfigWithDefaults() *K8sIoV1PodDNSConfig`

NewK8sIoV1PodDNSConfigWithDefaults instantiates a new K8sIoV1PodDNSConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameservers

`func (o *K8sIoV1PodDNSConfig) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *K8sIoV1PodDNSConfig) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *K8sIoV1PodDNSConfig) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *K8sIoV1PodDNSConfig) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetOptions

`func (o *K8sIoV1PodDNSConfig) GetOptions() []K8sIoV1PodDNSConfigOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *K8sIoV1PodDNSConfig) GetOptionsOk() (*[]K8sIoV1PodDNSConfigOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *K8sIoV1PodDNSConfig) SetOptions(v []K8sIoV1PodDNSConfigOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *K8sIoV1PodDNSConfig) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSearches

`func (o *K8sIoV1PodDNSConfig) GetSearches() []string`

GetSearches returns the Searches field if non-nil, zero value otherwise.

### GetSearchesOk

`func (o *K8sIoV1PodDNSConfig) GetSearchesOk() (*[]string, bool)`

GetSearchesOk returns a tuple with the Searches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearches

`func (o *K8sIoV1PodDNSConfig) SetSearches(v []string)`

SetSearches sets Searches field to given value.

### HasSearches

`func (o *K8sIoV1PodDNSConfig) HasSearches() bool`

HasSearches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


