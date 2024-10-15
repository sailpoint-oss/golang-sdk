package sailpoint

import (
	"context"
	"fmt"
	"os"
	"testing"

	beta "github.com/sailpoint-oss/golang-sdk/v2/api_beta"
	v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_v3(t *testing.T) {

	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.V3.AccountsAPI.ListAccounts(context.TODO()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test paginate search API", func(t *testing.T) {

		search := v3.NewSearchWithDefaults()
		search.Indices = append(search.Indices, "identities")
		searchString := []byte(`
		{
		"indices": [
			"identities"
		],
		"query": {
			"query": "*"
		},
		"sort": [
			"-name"
		]
		}
		  `)
		search.UnmarshalJSON(searchString)
		resp, r, err := PaginateSearchApi(context.TODO(), apiClient, *search, 0, 10, 100)
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, len(resp), 100)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test List Transforms", func(t *testing.T) {

		resp, r, err := apiClient.V3.TransformsAPI.ListTransforms(context.TODO()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test Pagination", func(t *testing.T) {

		resp, r, err := Paginate[v3.Account](apiClient.V3.AccountsAPI.ListAccounts(context.TODO()), 0, 10, 100)

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, len(resp), 100)
		assert.Equal(t, 200, r.StatusCode)

	})

}

func Test_beta(t *testing.T) {

	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.Beta.AccountsAPI.ListAccounts(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test connector api", func(t *testing.T) {

		resp, r, err := apiClient.Beta.ConnectorsAPI.GetConnectorList(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test List Sources", func(t *testing.T) {

		resp, r, err := apiClient.Beta.SourcesAPI.ListSources(context.TODO()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test Pagination", func(t *testing.T) {

		resp, r, err := Paginate[beta.IdentityProfile](apiClient.Beta.IdentityProfilesAPI.ListIdentityProfiles(context.TODO()), 0, 1, 2)

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 2, len(resp))
		assert.Equal(t, 200, r.StatusCode)

	})

}

func Test_v2024(t *testing.T) {

	configuration := NewDefaultConfiguration()
	configuration.Experimental = true
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.V2024.AccountsAPI.ListAccounts(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test Experimental Identities API", func(t *testing.T) {

		resp, r, err := apiClient.V2024.IdentitiesAPI.ListIdentities(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})
}

func Test_generic(t *testing.T) {

	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "v2024/accounts").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}
