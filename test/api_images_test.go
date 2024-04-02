/*
Harvester APIs

Testing ImagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package harvester

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/drewmullen/harvester-go-sdk"
)

func Test_harvester_ImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesAPIService CreateNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.CreateNamespacedVirtualMachineImage(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService DeleteNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.DeleteNamespacedVirtualMachineImage(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ListNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.ListNamespacedVirtualMachineImage(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ListVirtualMachineImageForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ImagesAPI.ListVirtualMachineImageForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService PatchNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.PatchNamespacedVirtualMachineImage(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ReadNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.ReadNamespacedVirtualMachineImage(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ReplaceNamespacedVirtualMachineImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.ImagesAPI.ReplaceNamespacedVirtualMachineImage(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
