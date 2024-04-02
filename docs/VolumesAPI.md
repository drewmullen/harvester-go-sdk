# \VolumesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedPersistentVolumeClaim**](VolumesAPI.md#CreateNamespacedPersistentVolumeClaim) | **Post** /api/v1/namespaces/{namespace}/persistentvolumeclaims | Create a Namespaced Persistent Volume Claim
[**DeleteNamespacedPersistentVolumeClaim**](VolumesAPI.md#DeleteNamespacedPersistentVolumeClaim) | **Delete** /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} | Delete a Namespaced Persistent Volume Claim
[**ListNamespacedPersistentVolumeClaim**](VolumesAPI.md#ListNamespacedPersistentVolumeClaim) | **Get** /api/v1/namespaces/{namespace}/persistentvolumeclaims | List Namespaced Persistent Volume Claims
[**ListPersistentVolumeClaimForAllNamespaces**](VolumesAPI.md#ListPersistentVolumeClaimForAllNamespaces) | **Get** /api/v1/persistentvolumeclaims | List Persistent Volume Claims For All Namespaces
[**PatchNamespacedPersistentVolumeClaim**](VolumesAPI.md#PatchNamespacedPersistentVolumeClaim) | **Patch** /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} | Patch a Namespaced Persistent Volume Claim
[**ReadNamespacedPersistentVolumeClaim**](VolumesAPI.md#ReadNamespacedPersistentVolumeClaim) | **Get** /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} | Read a Namespaced Persistent Volume Claim
[**ReplaceNamespacedPersistentVolumeClaim**](VolumesAPI.md#ReplaceNamespacedPersistentVolumeClaim) | **Put** /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name} | Replace a Namespaced Persistent Volume Claim



## CreateNamespacedPersistentVolumeClaim

> K8sIoV1PersistentVolumeClaim CreateNamespacedPersistentVolumeClaim(ctx, namespace).K8sIoV1PersistentVolumeClaim(k8sIoV1PersistentVolumeClaim).Execute()

Create a Namespaced Persistent Volume Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	k8sIoV1PersistentVolumeClaim := *openapiclient.NewK8sIoV1PersistentVolumeClaim("ApiVersion_example", "Kind_example") // K8sIoV1PersistentVolumeClaim | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.CreateNamespacedPersistentVolumeClaim(context.Background(), namespace).K8sIoV1PersistentVolumeClaim(k8sIoV1PersistentVolumeClaim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.CreateNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedPersistentVolumeClaim`: K8sIoV1PersistentVolumeClaim
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.CreateNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **k8sIoV1PersistentVolumeClaim** | [**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md) |  | 

### Return type

[**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedPersistentVolumeClaim

> K8sIoV1Status DeleteNamespacedPersistentVolumeClaim(ctx, name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

Delete a Namespaced Persistent Volume Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	k8sIoV1DeleteOptions := *openapiclient.NewK8sIoV1DeleteOptions("ApiVersion_example", "Kind_example") // K8sIoV1DeleteOptions | 
	gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
	orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
	propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.DeleteNamespacedPersistentVolumeClaim(context.Background(), name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.DeleteNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedPersistentVolumeClaim`: K8sIoV1Status
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.DeleteNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **k8sIoV1DeleteOptions** | [**K8sIoV1DeleteOptions**](K8sIoV1DeleteOptions.md) |  | 
 **gracePeriodSeconds** | **int32** | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **bool** | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **string** | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. | 

### Return type

[**K8sIoV1Status**](K8sIoV1Status.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedPersistentVolumeClaim

> K8sIoV1PersistentVolumeClaimList ListNamespacedPersistentVolumeClaim(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Namespaced Persistent Volume Claims



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	continue_ := "continue__example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)
	fieldSelector := "fieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. (optional)
	includeUninitialized := true // bool | If true, partially initialized resources are included in the response. (optional)
	labelSelector := "labelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything (optional)
	limit := int32(56) // int32 | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
	resourceVersion := "resourceVersion_example" // string | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. (optional)
	timeoutSeconds := int32(56) // int32 | TimeoutSeconds for the list/watch call. (optional)
	watch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.ListNamespacedPersistentVolumeClaim(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.ListNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedPersistentVolumeClaim`: K8sIoV1PersistentVolumeClaimList
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.ListNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **continue_** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **fieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **includeUninitialized** | **bool** | If true, partially initialized resources are included in the response. | 
 **labelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything | 
 **limit** | **int32** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **resourceVersion** | **string** | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. | 
 **timeoutSeconds** | **int32** | TimeoutSeconds for the list/watch call. | 
 **watch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 

### Return type

[**K8sIoV1PersistentVolumeClaimList**](K8sIoV1PersistentVolumeClaimList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersistentVolumeClaimForAllNamespaces

> K8sIoV1PersistentVolumeClaimList ListPersistentVolumeClaimForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Persistent Volume Claims For All Namespaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	continue_ := "continue__example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)
	fieldSelector := "fieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. (optional)
	includeUninitialized := true // bool | If true, partially initialized resources are included in the response. (optional)
	labelSelector := "labelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything (optional)
	limit := int32(56) // int32 | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
	resourceVersion := "resourceVersion_example" // string | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. (optional)
	timeoutSeconds := int32(56) // int32 | TimeoutSeconds for the list/watch call. (optional)
	watch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.ListPersistentVolumeClaimForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.ListPersistentVolumeClaimForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPersistentVolumeClaimForAllNamespaces`: K8sIoV1PersistentVolumeClaimList
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.ListPersistentVolumeClaimForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersistentVolumeClaimForAllNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continue_** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 
 **fieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **includeUninitialized** | **bool** | If true, partially initialized resources are included in the response. | 
 **labelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything | 
 **limit** | **int32** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **resourceVersion** | **string** | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. | 
 **timeoutSeconds** | **int32** | TimeoutSeconds for the list/watch call. | 
 **watch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 

### Return type

[**K8sIoV1PersistentVolumeClaimList**](K8sIoV1PersistentVolumeClaimList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedPersistentVolumeClaim

> K8sIoV1PersistentVolumeClaim PatchNamespacedPersistentVolumeClaim(ctx, name, namespace).Body(body).Execute()

Patch a Namespaced Persistent Volume Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PatchNamespacedPersistentVolumeClaim(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PatchNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespacedPersistentVolumeClaim`: K8sIoV1PersistentVolumeClaim
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PatchNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedPersistentVolumeClaim

> K8sIoV1PersistentVolumeClaim ReadNamespacedPersistentVolumeClaim(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Namespaced Persistent Volume Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	exact := true // bool | Should the export be exact. Exact export maintains cluster-specific fields like 'Namespace'. (optional)
	export := true // bool | Should this value be exported. Export strips fields that a user can not specify. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.ReadNamespacedPersistentVolumeClaim(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.ReadNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedPersistentVolumeClaim`: K8sIoV1PersistentVolumeClaim
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.ReadNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedPersistentVolumeClaim

> K8sIoV1PersistentVolumeClaim ReplaceNamespacedPersistentVolumeClaim(ctx, name, namespace).K8sIoV1PersistentVolumeClaim(k8sIoV1PersistentVolumeClaim).Execute()

Replace a Namespaced Persistent Volume Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	k8sIoV1PersistentVolumeClaim := *openapiclient.NewK8sIoV1PersistentVolumeClaim("ApiVersion_example", "Kind_example") // K8sIoV1PersistentVolumeClaim | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.ReplaceNamespacedPersistentVolumeClaim(context.Background(), name, namespace).K8sIoV1PersistentVolumeClaim(k8sIoV1PersistentVolumeClaim).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.ReplaceNamespacedPersistentVolumeClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNamespacedPersistentVolumeClaim`: K8sIoV1PersistentVolumeClaim
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.ReplaceNamespacedPersistentVolumeClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedPersistentVolumeClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **k8sIoV1PersistentVolumeClaim** | [**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md) |  | 

### Return type

[**K8sIoV1PersistentVolumeClaim**](K8sIoV1PersistentVolumeClaim.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

