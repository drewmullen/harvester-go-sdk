/*
Harvester APIs

Testing RestoresAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_RestoresAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RestoresAPIService CreateNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.CreateNamespacedVirtualMachineRestore(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService DeleteNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.DeleteNamespacedVirtualMachineRestore(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService ListNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.ListNamespacedVirtualMachineRestore(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService ListVirtualMachineRestoreForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RestoresAPI.ListVirtualMachineRestoreForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService PatchNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.PatchNamespacedVirtualMachineRestore(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService ReadNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.ReadNamespacedVirtualMachineRestore(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RestoresAPIService ReplaceNamespacedVirtualMachineRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.RestoresAPI.ReplaceNamespacedVirtualMachineRestore(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
