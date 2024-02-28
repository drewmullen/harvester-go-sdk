/*
Harvester APIs

Testing SupportBundlesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/drewmullen/harvester-go-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_SupportBundlesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SupportBundlesAPIService CreateNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.CreateNamespacedSupportBundle(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService DeleteNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.DeleteNamespacedSupportBundle(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService ListNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.ListNamespacedSupportBundle(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService ListSupportBundleForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SupportBundlesAPI.ListSupportBundleForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService PatchNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.PatchNamespacedSupportBundle(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService ReadNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.ReadNamespacedSupportBundle(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SupportBundlesAPIService ReplaceNamespacedSupportBundle", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var namespace string

		resp, httpRes, err := apiClient.SupportBundlesAPI.ReplaceNamespacedSupportBundle(context.Background(), name, namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
