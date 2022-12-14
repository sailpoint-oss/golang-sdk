/*
SailPoint SaaS API

Testing IdentitiesApiService

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

func Test_openapi_IdentitiesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentitiesApiService CreateIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentitiesApi.CreateIdentity(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService CreateLauncher", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string
		var accessProfileId string

		resp, httpRes, err := apiClient.IdentitiesApi.CreateLauncher(context.Background(), identityIdOrAlias, accessProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService DeleteIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.DeleteIdentity(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService DeleteLauncher", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string
		var launcherId string

		resp, httpRes, err := apiClient.IdentitiesApi.DeleteLauncher(context.Background(), identityIdOrAlias, launcherId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService GetIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.GetIdentity(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService ListApprovals", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.ListApprovals(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService ListApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.ListApps(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService ListIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentitiesApi.ListIdentities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService ListLaunchers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.ListLaunchers(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService LockIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentitiesApi.LockIdentities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesApiService UpdateIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityIdOrAlias string

		resp, httpRes, err := apiClient.IdentitiesApi.UpdateIdentity(context.Background(), identityIdOrAlias).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}