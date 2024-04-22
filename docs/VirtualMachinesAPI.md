# \VirtualMachinesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedVirtualMachine**](VirtualMachinesAPI.md#CreateNamespacedVirtualMachine) | **Post** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines | Create a Namespaced Virtual Machine
[**DeleteNamespacedVirtualMachine**](VirtualMachinesAPI.md#DeleteNamespacedVirtualMachine) | **Delete** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines/{name} | Delete a Namespaced Virtual Machine
[**ListNamespacedVirtualMachine**](VirtualMachinesAPI.md#ListNamespacedVirtualMachine) | **Get** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines | List Namespaced Virtual Machines
[**ListNamespacedVirtualMachineInstance**](VirtualMachinesAPI.md#ListNamespacedVirtualMachineInstance) | **Get** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachineinstances | List Namespaced Virtual Machine Instances
[**ListVirtualMachineForAllNamespaces**](VirtualMachinesAPI.md#ListVirtualMachineForAllNamespaces) | **Get** /apis/kubevirt.io/v1/virtualmachines | List Virtual Machines For All Namespaces
[**ListVirtualMachineInstanceForAllNamespaces**](VirtualMachinesAPI.md#ListVirtualMachineInstanceForAllNamespaces) | **Get** /apis/kubevirt.io/v1/virtualmachineinstances | List Virtual Machine Instances For All Namespaces
[**PatchNamespacedVirtualMachine**](VirtualMachinesAPI.md#PatchNamespacedVirtualMachine) | **Patch** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines/{name} | Patch a Namespaced Virtual Machine
[**ReadNamespacedVirtualMachine**](VirtualMachinesAPI.md#ReadNamespacedVirtualMachine) | **Get** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines/{name} | Read a Namespaced Virtual Machine
[**ReadNamespacedVirtualMachineInstance**](VirtualMachinesAPI.md#ReadNamespacedVirtualMachineInstance) | **Get** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachineinstances/{name} | Read a Namespaced Virtual Machine Instance
[**ReplaceNamespacedVirtualMachine**](VirtualMachinesAPI.md#ReplaceNamespacedVirtualMachine) | **Put** /apis/kubevirt.io/v1/namespaces/{namespace}/virtualmachines/{name} | Replace a Namespaced Virtual Machine



## CreateNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachine CreateNamespacedVirtualMachine(ctx, namespace).KubevirtIoApiCoreV1VirtualMachine(kubevirtIoApiCoreV1VirtualMachine).Execute()

Create a Namespaced Virtual Machine



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
	kubevirtIoApiCoreV1VirtualMachine := *openapiclient.NewKubevirtIoApiCoreV1VirtualMachine(*openapiclient.NewKubevirtIoApiCoreV1VirtualMachineSpec(*openapiclient.NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec())) // KubevirtIoApiCoreV1VirtualMachine | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.CreateNamespacedVirtualMachine(context.Background(), namespace).KubevirtIoApiCoreV1VirtualMachine(kubevirtIoApiCoreV1VirtualMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.CreateNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.CreateNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubevirtIoApiCoreV1VirtualMachine** | [**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md) |  | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachine DeleteNamespacedVirtualMachine(ctx, name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).RemovedDisks(removedDisks).Execute()

Delete a Namespaced Virtual Machine



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
	k8sIoV1DeleteOptions := *openapiclient.NewK8sIoV1DeleteOptions() // K8sIoV1DeleteOptions | 
	gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
	orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
	propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)
	removedDisks := "removedDisks_example" // string | The disks that should be removed when deleting a virtual machine instance. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.DeleteNamespacedVirtualMachine(context.Background(), name, namespace).K8sIoV1DeleteOptions(k8sIoV1DeleteOptions).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).RemovedDisks(removedDisks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.DeleteNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.DeleteNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **k8sIoV1DeleteOptions** | [**K8sIoV1DeleteOptions**](K8sIoV1DeleteOptions.md) |  | 
 **gracePeriodSeconds** | **int32** | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **bool** | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **string** | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. | 
 **removedDisks** | **string** | The disks that should be removed when deleting a virtual machine instance. | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachineList ListNamespacedVirtualMachine(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Namespaced Virtual Machines



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
	resp, r, err := apiClient.VirtualMachinesAPI.ListNamespacedVirtualMachine(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ListNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachineList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ListNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedVirtualMachineRequest struct via the builder pattern


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

[**KubevirtIoApiCoreV1VirtualMachineList**](KubevirtIoApiCoreV1VirtualMachineList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedVirtualMachineInstance

> KubevirtIoApiCoreV1VirtualMachineInstanceList ListNamespacedVirtualMachineInstance(ctx, namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Namespaced Virtual Machine Instances



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
	resp, r, err := apiClient.VirtualMachinesAPI.ListNamespacedVirtualMachineInstance(context.Background(), namespace).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ListNamespacedVirtualMachineInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNamespacedVirtualMachineInstance`: KubevirtIoApiCoreV1VirtualMachineInstanceList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ListNamespacedVirtualMachineInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedVirtualMachineInstanceRequest struct via the builder pattern


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

[**KubevirtIoApiCoreV1VirtualMachineInstanceList**](KubevirtIoApiCoreV1VirtualMachineInstanceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineForAllNamespaces

> KubevirtIoApiCoreV1VirtualMachineList ListVirtualMachineForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machines For All Namespaces



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
	resp, r, err := apiClient.VirtualMachinesAPI.ListVirtualMachineForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ListVirtualMachineForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineForAllNamespaces`: KubevirtIoApiCoreV1VirtualMachineList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ListVirtualMachineForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineForAllNamespacesRequest struct via the builder pattern


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

[**KubevirtIoApiCoreV1VirtualMachineList**](KubevirtIoApiCoreV1VirtualMachineList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualMachineInstanceForAllNamespaces

> KubevirtIoApiCoreV1VirtualMachineInstanceList ListVirtualMachineInstanceForAllNamespaces(ctx).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()

List Virtual Machine Instances For All Namespaces



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
	resp, r, err := apiClient.VirtualMachinesAPI.ListVirtualMachineInstanceForAllNamespaces(context.Background()).Continue_(continue_).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).LabelSelector(labelSelector).Limit(limit).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Watch(watch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ListVirtualMachineInstanceForAllNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVirtualMachineInstanceForAllNamespaces`: KubevirtIoApiCoreV1VirtualMachineInstanceList
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ListVirtualMachineInstanceForAllNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualMachineInstanceForAllNamespacesRequest struct via the builder pattern


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

[**KubevirtIoApiCoreV1VirtualMachineInstanceList**](KubevirtIoApiCoreV1VirtualMachineInstanceList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachine PatchNamespacedVirtualMachine(ctx, name, namespace).Body(body).Execute()

Patch a Namespaced Virtual Machine



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
	resp, r, err := apiClient.VirtualMachinesAPI.PatchNamespacedVirtualMachine(context.Background(), name, namespace).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.PatchNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.PatchNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachine ReadNamespacedVirtualMachine(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Namespaced Virtual Machine



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
	resp, r, err := apiClient.VirtualMachinesAPI.ReadNamespacedVirtualMachine(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ReadNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ReadNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedVirtualMachineInstance

> KubevirtIoApiCoreV1VirtualMachineInstance ReadNamespacedVirtualMachineInstance(ctx, name, namespace).Exact(exact).Export(export).Execute()

Read a Namespaced Virtual Machine Instance



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
	resp, r, err := apiClient.VirtualMachinesAPI.ReadNamespacedVirtualMachineInstance(context.Background(), name, namespace).Exact(exact).Export(export).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ReadNamespacedVirtualMachineInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadNamespacedVirtualMachineInstance`: KubevirtIoApiCoreV1VirtualMachineInstance
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ReadNamespacedVirtualMachineInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedVirtualMachineInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exact** | **bool** | Should the export be exact. Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 
 **export** | **bool** | Should this value be exported. Export strips fields that a user can not specify. | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachineInstance**](KubevirtIoApiCoreV1VirtualMachineInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json;stream=watch, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedVirtualMachine

> KubevirtIoApiCoreV1VirtualMachine ReplaceNamespacedVirtualMachine(ctx, name, namespace).KubevirtIoApiCoreV1VirtualMachine(kubevirtIoApiCoreV1VirtualMachine).Execute()

Replace a Namespaced Virtual Machine



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
	kubevirtIoApiCoreV1VirtualMachine := *openapiclient.NewKubevirtIoApiCoreV1VirtualMachine(*openapiclient.NewKubevirtIoApiCoreV1VirtualMachineSpec(*openapiclient.NewKubevirtIoApiCoreV1VirtualMachineInstanceTemplateSpec())) // KubevirtIoApiCoreV1VirtualMachine | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VirtualMachinesAPI.ReplaceNamespacedVirtualMachine(context.Background(), name, namespace).KubevirtIoApiCoreV1VirtualMachine(kubevirtIoApiCoreV1VirtualMachine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachinesAPI.ReplaceNamespacedVirtualMachine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceNamespacedVirtualMachine`: KubevirtIoApiCoreV1VirtualMachine
	fmt.Fprintf(os.Stdout, "Response from `VirtualMachinesAPI.ReplaceNamespacedVirtualMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the resource | 
**namespace** | **string** | Object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedVirtualMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kubevirtIoApiCoreV1VirtualMachine** | [**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md) |  | 

### Return type

[**KubevirtIoApiCoreV1VirtualMachine**](KubevirtIoApiCoreV1VirtualMachine.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/yaml
- **Accept**: application/json, application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

