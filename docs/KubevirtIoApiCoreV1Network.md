# KubevirtIoApiCoreV1Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Multus** | Pointer to [**KubevirtIoApiCoreV1MultusNetwork**](KubevirtIoApiCoreV1MultusNetwork.md) |  | [optional] 
**Name** | **string** | Network name. Must be a DNS_LABEL and unique within the vm. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names | [default to ""]
**Pod** | Pointer to [**KubevirtIoApiCoreV1PodNetwork**](KubevirtIoApiCoreV1PodNetwork.md) |  | [optional] 

## Methods

### NewKubevirtIoApiCoreV1Network

`func NewKubevirtIoApiCoreV1Network(name string, ) *KubevirtIoApiCoreV1Network`

NewKubevirtIoApiCoreV1Network instantiates a new KubevirtIoApiCoreV1Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoApiCoreV1NetworkWithDefaults

`func NewKubevirtIoApiCoreV1NetworkWithDefaults() *KubevirtIoApiCoreV1Network`

NewKubevirtIoApiCoreV1NetworkWithDefaults instantiates a new KubevirtIoApiCoreV1Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultus

`func (o *KubevirtIoApiCoreV1Network) GetMultus() KubevirtIoApiCoreV1MultusNetwork`

GetMultus returns the Multus field if non-nil, zero value otherwise.

### GetMultusOk

`func (o *KubevirtIoApiCoreV1Network) GetMultusOk() (*KubevirtIoApiCoreV1MultusNetwork, bool)`

GetMultusOk returns a tuple with the Multus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultus

`func (o *KubevirtIoApiCoreV1Network) SetMultus(v KubevirtIoApiCoreV1MultusNetwork)`

SetMultus sets Multus field to given value.

### HasMultus

`func (o *KubevirtIoApiCoreV1Network) HasMultus() bool`

HasMultus returns a boolean if a field has been set.

### GetName

`func (o *KubevirtIoApiCoreV1Network) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubevirtIoApiCoreV1Network) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubevirtIoApiCoreV1Network) SetName(v string)`

SetName sets Name field to given value.


### GetPod

`func (o *KubevirtIoApiCoreV1Network) GetPod() KubevirtIoApiCoreV1PodNetwork`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *KubevirtIoApiCoreV1Network) GetPodOk() (*KubevirtIoApiCoreV1PodNetwork, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *KubevirtIoApiCoreV1Network) SetPod(v KubevirtIoApiCoreV1PodNetwork)`

SetPod sets Pod field to given value.

### HasPod

`func (o *KubevirtIoApiCoreV1Network) HasPod() bool`

HasPod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


