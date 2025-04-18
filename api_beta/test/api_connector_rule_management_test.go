/*
Identity Security Cloud Beta API

Testing ConnectorRuleManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_beta_ConnectorRuleManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectorRuleManagementAPIService CreateConnectorRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectorRuleManagementAPI.CreateConnectorRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorRuleManagementAPIService DeleteConnectorRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ConnectorRuleManagementAPI.DeleteConnectorRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorRuleManagementAPIService GetConnectorRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ConnectorRuleManagementAPI.GetConnectorRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorRuleManagementAPIService GetConnectorRuleList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectorRuleManagementAPI.GetConnectorRuleList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorRuleManagementAPIService UpdateConnectorRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ConnectorRuleManagementAPI.UpdateConnectorRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectorRuleManagementAPIService ValidateConnectorRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectorRuleManagementAPI.ValidateConnectorRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
