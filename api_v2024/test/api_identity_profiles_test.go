/*
Identity Security Cloud V2024 API

Testing IdentityProfilesAPIService

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

func Test_api_v2024_IdentityProfilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentityProfilesAPIService CreateIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.CreateIdentityProfile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService DeleteIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.IdentityProfilesAPI.DeleteIdentityProfile(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService DeleteIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.DeleteIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService ExportIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.ExportIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService GetDefaultIdentityAttributeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.IdentityProfilesAPI.GetDefaultIdentityAttributeConfig(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService GetIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.IdentityProfilesAPI.GetIdentityProfile(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService ImportIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.ImportIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService ListIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.ListIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService ShowIdentityPreview", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentityProfilesAPI.ShowIdentityPreview(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService SyncIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.IdentityProfilesAPI.SyncIdentityProfile(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesAPIService UpdateIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.IdentityProfilesAPI.UpdateIdentityProfile(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}