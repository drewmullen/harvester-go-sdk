# \ImagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedVirtualMachineImage**](ImagesAPI.md#CreateNamespacedVirtualMachineImage) | **Post** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages | Create a Virtual Machine Image
[**DeleteNamespacedVirtualMachineImage**](ImagesAPI.md#DeleteNamespacedVirtualMachineImage) | **Delete** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages/{name} | Delete a Virtual Machine Image
[**ListNamespacedVirtualMachineImage**](ImagesAPI.md#ListNamespacedVirtualMachineImage) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages | List Virtual Machine Images
[**ListVirtualMachineImageForAllNamespaces**](ImagesAPI.md#ListVirtualMachineImageForAllNamespaces) | **Get** /apis/harvesterhci.io/v1beta1/virtualmachineimages | List Virtual Machine Images For All Namespaces
[**PatchNamespacedVirtualMachineImage**](ImagesAPI.md#PatchNamespacedVirtualMachineImage) | **Patch** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages/{name} | Patch a Virtual Machine Image
[**ReadNamespacedVirtualMachineImage**](ImagesAPI.md#ReadNamespacedVirtualMachineImage) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages/{name} | Read a Virtual Machine Image
[**ReplaceNamespacedVirtualMachineImage**](ImagesAPI.md#ReplaceNamespacedVirtualMachineImage) | **Put** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachineimages/{name} | Replace a Virtual Machine Image



## CreateNamespacedVirtualMachineImage

> HarvesterhciIoV1beta1VirtualMachineImage CreateNamespacedVirtualMachineImage(ctx, namespace).Body(body).Execute()

Create a Virtual Machine Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	body := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineImage("ApiVersion_example", "Kind_example", *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineImageSpec("DisplayName_example", "SourceType_example")) // HarvesterhciIoV1beta1VirtualMachineImage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.CreateNamespacedVirtualMachineImage(context.Background(), namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.CreateNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedVirtualMachineImage`: HarvesterhciIoV1beta1VirtualMachineImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.CreateNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedVirtualMachineImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedVirtualMachineImage

> K8sIoV1Status DeleteNamespacedVirtualMachineImage(ctx, name, namespace).Body(body).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

Delete a Virtual Machine Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	body := *openapiclient.NewK8sIoV1DeleteOptions("ApiVersion_example", "Kind_example") // K8sIoV1DeleteOptions | 
	gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
	orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
	propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.DeleteNamespacedVirtualMachineImage(context.Background(), name, namespace).Body(body).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.DeleteNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedVirtualMachineImage`: K8sIoV1Status
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.DeleteNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedVirtualMachineImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**K8sIoV1DeleteOptions**](K8sIoV1DeleteOptions.md) |  | 
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


## ListNamespacedVirtualMachineImage

> HarvesterhciIoV1beta1VirtualMachineImageList ListNamespacedVirtualMachineImage(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machine Images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	resp, r, err := apiClient.ImagesAPI.ListNamespacedVirtualMachineImage(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedVirtualMachineImage`: HarvesterhciIoV1beta1VirtualMachineImageList
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedVirtualMachineImageRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineImageList**](HarvesterhciIoV1beta1VirtualMachineImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineImageForAllNamespaces

> HarvesterhciIoV1beta1VirtualMachineImageList ListVirtualMachineImageForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machine Images For All Namespaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
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
	resp, r, err := apiClient.ImagesAPI.ListVirtualMachineImageForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListVirtualMachineImageForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineImageForAllNamespaces`: HarvesterhciIoV1beta1VirtualMachineImageList
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListVirtualMachineImageForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineImageForAllNamespacesRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineImageList**](HarvesterhciIoV1beta1VirtualMachineImageList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedVirtualMachineImage

> HarvesterhciIoV1beta1VirtualMachineImage PatchNamespacedVirtualMachineImage(ctx, name, namespace).Body(body).Execute()

Patch a Virtual Machine Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.PatchNamespacedVirtualMachineImage(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.PatchNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespacedVirtualMachineImage`: HarvesterhciIoV1beta1VirtualMachineImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.PatchNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedVirtualMachineImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedVirtualMachineImage

> HarvesterhciIoV1beta1VirtualMachineImage ReadNamespacedVirtualMachineImage(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Virtual Machine Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	exact := true // bool | Should the export be exact. Exact export maintains cluster-specific fields like 'Namespace'. (optional)
	export := true // bool | Should this value be exported. Export strips fields that a user can not specify. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ReadNamespacedVirtualMachineImage(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ReadNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedVirtualMachineImage`: HarvesterhciIoV1beta1VirtualMachineImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ReadNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedVirtualMachineImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/json;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedVirtualMachineImage

> HarvesterhciIoV1beta1VirtualMachineImage ReplaceNamespacedVirtualMachineImage(ctx, name, namespace).Body(body).Execute()

Replace a Virtual Machine Image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | Name of the resource
	namespace := "namespace_example" // string | Object name and auth scope, such as for teams and projects
	body := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineImage("ApiVersion_example", "Kind_example", *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineImageSpec("DisplayName_example", "SourceType_example")) // HarvesterhciIoV1beta1VirtualMachineImage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ReplaceNamespacedVirtualMachineImage(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ReplaceNamespacedVirtualMachineImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNamespacedVirtualMachineImage`: HarvesterhciIoV1beta1VirtualMachineImage
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ReplaceNamespacedVirtualMachineImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedVirtualMachineImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineImage**](HarvesterhciIoV1beta1VirtualMachineImage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

