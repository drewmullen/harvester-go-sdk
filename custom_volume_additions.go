package harvester

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ExportVolumeInput struct {
	DisplayName      string `json:"displayName"`
	Namespace        string `json:"namespace"`
	StorageClassName string `json:"storageClassName"`
}

type CloneVolumeInput struct {
	Name string `json:"name"`
}

type SnapshotVolumeInput struct {
	Name string `json:"name"`
}

var PersistentVolumeClaimActionExport = "export"
var PersistentVolumeClaimActionClone = "clone"
var PersistentVolumeClaimActionSnapshot = "snapshot"

func getActions() []string {
	return []string{
		PersistentVolumeClaimActionExport,
		PersistentVolumeClaimActionClone,
		PersistentVolumeClaimActionSnapshot,
	}
}

type ApiPostNamespacedPersistentVolumeClaimRequest struct {
	ctx        context.Context
	ApiService *VolumesAPIService
	name       string
	namespace  string
	body       *map[string]interface{}
	action     string
}

func (r ApiPostNamespacedPersistentVolumeClaimRequest) Body(body map[string]interface{}) ApiPostNamespacedPersistentVolumeClaimRequest {
	r.body = &body
	return r
}

func (r ApiPostNamespacedPersistentVolumeClaimRequest) Action(action string) ApiPostNamespacedPersistentVolumeClaimRequest {
	r.action = action
	return r
}

func (r ApiPostNamespacedPersistentVolumeClaimRequest) Execute() error {
	return r.ApiService.PostNamespacedPersistentVolumeClaimExecute(r)
}

/*
PostNamespacedPersistentVolumeClaim Post a Namespaced Persistent Volume Claim

Post a PersistentVolumeClaim object.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Name of the resource
	@param namespace Object name and auth scope, such as for teams and projects
	@return ApiPostNamespacedPersistentVolumeClaimRequest
*/
func (a *VolumesAPIService) PostNamespacedPersistentVolumeClaim(ctx context.Context, name string, namespace string) ApiPostNamespacedPersistentVolumeClaimRequest {
	return ApiPostNamespacedPersistentVolumeClaimRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
		namespace:  namespace,
	}
}

// Execute executes the request
//
//	@return K8sIoV1PersistentVolumeClaim
func (a *VolumesAPIService) PostNamespacedPersistentVolumeClaimExecute(r ApiPostNamespacedPersistentVolumeClaimRequest) error {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	// accept body relative to the action provided
	// set query parameter for ?action
	// returns nothing
	//

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumesAPIService.PostNamespacedPersistentVolumeClaim")
	if err != nil {
		return &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/harvester/persistentvolumeclaims/{namespace}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return reportError("body is required and must be specified")
	}

	localVarQueryParams.Add("action", r.action)

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-post+json", "application/merge-post+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return err
	}

	if localVarHTTPResponse.StatusCode == 204 {
		return nil
	}
	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return newErr
	}
	return fmt.Errorf("Unreachable")
}
