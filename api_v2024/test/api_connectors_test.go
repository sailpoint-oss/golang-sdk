/*
Identity Security Cloud V2024 API

Testing ConnectorsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_v2024_ConnectorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectorsAPIService CreateCustomConnector", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectorsAPI.CreateCustomConnector(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService DeleteCustomConnector", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		httpRes, err := apiClient.ConnectorsAPI.DeleteCustomConnector(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnector", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnector(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnectorCorrelationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnectorCorrelationConfig(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnectorList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnectorList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnectorSourceConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnectorSourceConfig(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnectorSourceTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnectorSourceTemplate(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService GetConnectorTranslations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string
		var locale string

		resp, httpRes, err := apiClient.ConnectorsAPI.GetConnectorTranslations(context.Background(), scriptName, locale).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService PutCorrelationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.PutCorrelationConfig(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService PutSourceConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.PutSourceConfig(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService PutSourceTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.PutSourceTemplate(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService PutTranslations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string
		var locale string

		resp, httpRes, err := apiClient.ConnectorsAPI.PutTranslations(context.Background(), scriptName, locale).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorsAPIService UpdateConnector", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.ConnectorsAPI.UpdateConnector(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}