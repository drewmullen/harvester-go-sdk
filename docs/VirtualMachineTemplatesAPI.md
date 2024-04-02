# \VirtualMachineTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#CreateNamespacedVirtualMachineTemplate) | **Post** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates | Create a Namespaced Virtual Machine Template
[**CreateNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#CreateNamespacedVirtualMachineTemplateVersion) | **Post** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions | Create a Namespaced Virtual Machine Template Version
[**DeleteNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#DeleteNamespacedVirtualMachineTemplate) | **Delete** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates/{name} | Delete a Namespaced Virtual Machine Template
[**DeleteNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#DeleteNamespacedVirtualMachineTemplateVersion) | **Delete** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions/{name} | Delete a Namespaced Virtual Machine Template Version
[**ListNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#ListNamespacedVirtualMachineTemplate) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates | List Namespaced Virtual Machine Templates
[**ListNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#ListNamespacedVirtualMachineTemplateVersion) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions | List Namespaced Virtual Machine Template Versions
[**ListVirtualMachineTemplateForAllNamespaces**](VirtualMachineTemplatesAPI.md#ListVirtualMachineTemplateForAllNamespaces) | **Get** /apis/harvesterhci.io/v1beta1/virtualmachinetemplates | List Virtual Machine Templates For All Namespaces
[**ListVirtualMachineTemplateVersionForAllNamespaces**](VirtualMachineTemplatesAPI.md#ListVirtualMachineTemplateVersionForAllNamespaces) | **Get** /apis/harvesterhci.io/v1beta1/virtualmachinetemplateversions | List Virtual Machine Template Versions For All Namespaces
[**PatchNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#PatchNamespacedVirtualMachineTemplate) | **Patch** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates/{name} | Patch a Namespaced Virtual Machine Template
[**PatchNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#PatchNamespacedVirtualMachineTemplateVersion) | **Patch** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions/{name} | Patch a Namespaced Virtual Machine Template Version
[**ReadNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#ReadNamespacedVirtualMachineTemplate) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates/{name} | Read a Namespaced Virtual Machine Template
[**ReadNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#ReadNamespacedVirtualMachineTemplateVersion) | **Get** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions/{name} | Read a Namespaced Virtual Machine Template Version
[**ReplaceNamespacedVirtualMachineTemplate**](VirtualMachineTemplatesAPI.md#ReplaceNamespacedVirtualMachineTemplate) | **Put** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplates/{name} | Replace a Namespaced Virtual Machine Template
[**ReplaceNamespacedVirtualMachineTemplateVersion**](VirtualMachineTemplatesAPI.md#ReplaceNamespacedVirtualMachineTemplateVersion) | **Put** /apis/harvesterhci.io/v1beta1/namespaces/{namespace}/virtualmachinetemplateversions/{name} | Replace a Namespaced Virtual Machine Template Version



## CreateNamespacedVirtualMachineTemplate

> HarvesterhciIoV1beta1VirtualMachineTemplate CreateNamespacedVirtualMachineTemplate(ctx, namespace).HarvesterhciIoV1beta1VirtualMachineTemplate(harvesterhciIoV1beta1VirtualMachineTemplate).Execute()

Create a Namespaced Virtual Machine Template



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
	harvesterhciIoV1beta1VirtualMachineTemplate := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplate("ApiVersion_example", "Kind_example") // HarvesterhciIoV1beta1VirtualMachineTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplate(context.Background(), namespace).HarvesterhciIoV1beta1VirtualMachineTemplate(harvesterhciIoV1beta1VirtualMachineTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedVirtualMachineTemplate`: HarvesterhciIoV1beta1VirtualMachineTemplate
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedVirtualMachineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harvesterhciIoV1beta1VirtualMachineTemplate** | [**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespacedVirtualMachineTemplateVersion

> HarvesterhciIoV1beta1VirtualMachineTemplateVersion CreateNamespacedVirtualMachineTemplateVersion(ctx, namespace).HarvesterhciIoV1beta1VirtualMachineTemplateVersion(harvesterhciIoV1beta1VirtualMachineTemplateVersion).Execute()

Create a Namespaced Virtual Machine Template Version



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
	harvesterhciIoV1beta1VirtualMachineTemplateVersion := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplateVersion("ApiVersion_example", "Kind_example", *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec("TemplateId_example")) // HarvesterhciIoV1beta1VirtualMachineTemplateVersion | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplateVersion(context.Background(), namespace).HarvesterhciIoV1beta1VirtualMachineTemplateVersion(harvesterhciIoV1beta1VirtualMachineTemplateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedVirtualMachineTemplateVersion`: HarvesterhciIoV1beta1VirtualMachineTemplateVersion
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.CreateNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **harvesterhciIoV1beta1VirtualMachineTemplateVersion** | [**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedVirtualMachineTemplate

> K8sIoV1Status DeleteNamespacedVirtualMachineTemplate(ctx, name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

Delete a Namespaced Virtual Machine Template



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplate(context.Background(), name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedVirtualMachineTemplate`: K8sIoV1Status
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedVirtualMachineTemplateRequest struct via the builder pattern


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


## DeleteNamespacedVirtualMachineTemplateVersion

> K8sIoV1Status DeleteNamespacedVirtualMachineTemplateVersion(ctx, name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

Delete a Namespaced Virtual Machine Template Version



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplateVersion(context.Background(), name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedVirtualMachineTemplateVersion`: K8sIoV1Status
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.DeleteNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


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


## ListNamespacedVirtualMachineTemplate

> HarvesterhciIoV1beta1VirtualMachineTemplateList ListNamespacedVirtualMachineTemplate(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Namespaced Virtual Machine Templates



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplate(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedVirtualMachineTemplate`: HarvesterhciIoV1beta1VirtualMachineTemplateList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedVirtualMachineTemplateRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineTemplateList**](HarvesterhciIoV1beta1VirtualMachineTemplateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedVirtualMachineTemplateVersion

> HarvesterhciIoV1beta1VirtualMachineTemplateVersionList ListNamespacedVirtualMachineTemplateVersion(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Namespaced Virtual Machine Template Versions



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplateVersion(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedVirtualMachineTemplateVersion`: HarvesterhciIoV1beta1VirtualMachineTemplateVersionList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ListNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersionList**](HarvesterhciIoV1beta1VirtualMachineTemplateVersionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineTemplateForAllNamespaces

> HarvesterhciIoV1beta1VirtualMachineTemplateList ListVirtualMachineTemplateForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machine Templates For All Namespaces



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ListVirtualMachineTemplateForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ListVirtualMachineTemplateForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineTemplateForAllNamespaces`: HarvesterhciIoV1beta1VirtualMachineTemplateList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ListVirtualMachineTemplateForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineTemplateForAllNamespacesRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineTemplateList**](HarvesterhciIoV1beta1VirtualMachineTemplateList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineTemplateVersionForAllNamespaces

> HarvesterhciIoV1beta1VirtualMachineTemplateVersionList ListVirtualMachineTemplateVersionForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machine Template Versions For All Namespaces



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ListVirtualMachineTemplateVersionForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ListVirtualMachineTemplateVersionForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineTemplateVersionForAllNamespaces`: HarvesterhciIoV1beta1VirtualMachineTemplateVersionList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ListVirtualMachineTemplateVersionForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineTemplateVersionForAllNamespacesRequest struct via the builder pattern


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

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersionList**](HarvesterhciIoV1beta1VirtualMachineTemplateVersionList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedVirtualMachineTemplate

> HarvesterhciIoV1beta1VirtualMachineTemplate PatchNamespacedVirtualMachineTemplate(ctx, name, namespace).Body(body).Execute()

Patch a Namespaced Virtual Machine Template



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplate(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespacedVirtualMachineTemplate`: HarvesterhciIoV1beta1VirtualMachineTemplate
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedVirtualMachineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedVirtualMachineTemplateVersion

> HarvesterhciIoV1beta1VirtualMachineTemplateVersion PatchNamespacedVirtualMachineTemplateVersion(ctx, name, namespace).Body(body).Execute()

Patch a Namespaced Virtual Machine Template Version



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplateVersion(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespacedVirtualMachineTemplateVersion`: HarvesterhciIoV1beta1VirtualMachineTemplateVersion
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.PatchNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedVirtualMachineTemplate

> HarvesterhciIoV1beta1VirtualMachineTemplate ReadNamespacedVirtualMachineTemplate(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Namespaced Virtual Machine Template



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplate(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedVirtualMachineTemplate`: HarvesterhciIoV1beta1VirtualMachineTemplate
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedVirtualMachineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedVirtualMachineTemplateVersion

> HarvesterhciIoV1beta1VirtualMachineTemplateVersion ReadNamespacedVirtualMachineTemplateVersion(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Namespaced Virtual Machine Template Version



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
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplateVersion(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedVirtualMachineTemplateVersion`: HarvesterhciIoV1beta1VirtualMachineTemplateVersion
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ReadNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedVirtualMachineTemplate

> HarvesterhciIoV1beta1VirtualMachineTemplate ReplaceNamespacedVirtualMachineTemplate(ctx, name, namespace).HarvesterhciIoV1beta1VirtualMachineTemplate(harvesterhciIoV1beta1VirtualMachineTemplate).Execute()

Replace a Namespaced Virtual Machine Template



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
	harvesterhciIoV1beta1VirtualMachineTemplate := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplate("ApiVersion_example", "Kind_example") // HarvesterhciIoV1beta1VirtualMachineTemplate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplate(context.Background(), name, namespace).HarvesterhciIoV1beta1VirtualMachineTemplate(harvesterhciIoV1beta1VirtualMachineTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNamespacedVirtualMachineTemplate`: HarvesterhciIoV1beta1VirtualMachineTemplate
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedVirtualMachineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harvesterhciIoV1beta1VirtualMachineTemplate** | [**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplate**](HarvesterhciIoV1beta1VirtualMachineTemplate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedVirtualMachineTemplateVersion

> HarvesterhciIoV1beta1VirtualMachineTemplateVersion ReplaceNamespacedVirtualMachineTemplateVersion(ctx, name, namespace).HarvesterhciIoV1beta1VirtualMachineTemplateVersion(harvesterhciIoV1beta1VirtualMachineTemplateVersion).Execute()

Replace a Namespaced Virtual Machine Template Version



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
	harvesterhciIoV1beta1VirtualMachineTemplateVersion := *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplateVersion("ApiVersion_example", "Kind_example", *openapiclient.NewHarvesterhciIoV1beta1VirtualMachineTemplateVersionSpec("TemplateId_example")) // HarvesterhciIoV1beta1VirtualMachineTemplateVersion | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplateVersion(context.Background(), name, namespace).HarvesterhciIoV1beta1VirtualMachineTemplateVersion(harvesterhciIoV1beta1VirtualMachineTemplateVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNamespacedVirtualMachineTemplateVersion`: HarvesterhciIoV1beta1VirtualMachineTemplateVersion
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachineTemplatesAPI.ReplaceNamespacedVirtualMachineTemplateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedVirtualMachineTemplateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **harvesterhciIoV1beta1VirtualMachineTemplateVersion** | [**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md) |  | 

### Return type

[**HarvesterhciIoV1beta1VirtualMachineTemplateVersion**](HarvesterhciIoV1beta1VirtualMachineTemplateVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

