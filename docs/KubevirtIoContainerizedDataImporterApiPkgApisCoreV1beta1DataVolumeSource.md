# KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blank** | Pointer to **map[string]interface{}** | DataVolumeBlankImage provides the parameters to create a new raw blank image for the PVC | [optional] 
**Gcs** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS.md) |  | [optional] 
**Http** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP.md) |  | [optional] 
**Imageio** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceImageIO**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceImageIO.md) |  | [optional] 
**Pvc** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourcePVC**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourcePVC.md) |  | [optional] 
**Registry** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry.md) |  | [optional] 
**S3** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3.md) |  | [optional] 
**Snapshot** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot.md) |  | [optional] 
**Upload** | Pointer to **map[string]interface{}** | DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source | [optional] 
**Vddk** | Pointer to [**KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceVDDK**](KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceVDDK.md) |  | [optional] 

## Methods

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceWithDefaults

`func NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceWithDefaults() *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource`

NewKubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceWithDefaults instantiates a new KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlank

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetBlank() map[string]interface{}`

GetBlank returns the Blank field if non-nil, zero value otherwise.

### GetBlankOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetBlankOk() (*map[string]interface{}, bool)`

GetBlankOk returns a tuple with the Blank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlank

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetBlank(v map[string]interface{})`

SetBlank sets Blank field to given value.

### HasBlank

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasBlank() bool`

HasBlank returns a boolean if a field has been set.

### GetGcs

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetGcs() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetGcsOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetGcs(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceGCS)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetHttp

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetHttp() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetHttpOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetHttp(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceHTTP)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetImageio

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetImageio() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceImageIO`

GetImageio returns the Imageio field if non-nil, zero value otherwise.

### GetImageioOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetImageioOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceImageIO, bool)`

GetImageioOk returns a tuple with the Imageio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageio

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetImageio(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceImageIO)`

SetImageio sets Imageio field to given value.

### HasImageio

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasImageio() bool`

HasImageio returns a boolean if a field has been set.

### GetPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetPvc() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourcePVC`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetPvcOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourcePVC, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetPvc(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourcePVC)`

SetPvc sets Pvc field to given value.

### HasPvc

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasPvc() bool`

HasPvc returns a boolean if a field has been set.

### GetRegistry

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetRegistry() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetRegistryOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetRegistry(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceRegistry)`

SetRegistry sets Registry field to given value.

### HasRegistry

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasRegistry() bool`

HasRegistry returns a boolean if a field has been set.

### GetS3

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetS3() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetS3Ok() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetS3(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceS3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetSnapshot

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetSnapshot() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetSnapshotOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetSnapshot(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetUpload

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetUpload() map[string]interface{}`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetUploadOk() (*map[string]interface{}, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetUpload(v map[string]interface{})`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetVddk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetVddk() KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceVDDK`

GetVddk returns the Vddk field if non-nil, zero value otherwise.

### GetVddkOk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) GetVddkOk() (*KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceVDDK, bool)`

GetVddkOk returns a tuple with the Vddk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVddk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) SetVddk(v KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSourceVDDK)`

SetVddk sets Vddk field to given value.

### HasVddk

`func (o *KubevirtIoContainerizedDataImporterApiPkgApisCoreV1beta1DataVolumeSource) HasVddk() bool`

HasVddk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


