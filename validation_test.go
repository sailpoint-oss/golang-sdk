package sailpoint

import (
	"context"
	"fmt"
	"os"
	"testing"

	beta "github.com/sailpoint-oss/golang-sdk/beta"
	v3 "github.com/sailpoint-oss/golang-sdk/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_v3(t *testing.T) {

	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.V3.AccountsApi.ListAccounts(context.TODO()).Execute()

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

		resp, r, err := apiClient.V3.TransformsApi.ListTransforms(context.TODO()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test Pagination", func(t *testing.T) {

		resp, r, err := Paginate[v3.Account](apiClient.V3.AccountsApi.ListAccounts(context.TODO()), 0, 10, 100)

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

		resp, r, err := apiClient.Beta.AccountsApi.ListAccounts(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test connector api", func(t *testing.T) {

		resp, r, err := apiClient.Beta.ConnectorsApi.GetConnectorList(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test List Sources", func(t *testing.T) {

		resp, r, err := apiClient.Beta.SourcesApi.ListSources(context.TODO()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

	})

	t.Run("Test Pagination", func(t *testing.T) {

		resp, r, err := Paginate[beta.IdentityProfile](apiClient.Beta.IdentityProfilesApi.ListIdentityProfiles(context.TODO()), 0, 1, 2)

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 2, len(resp))
		assert.Equal(t, 200, r.StatusCode)

	})

}
func Test_v2(t *testing.T) {
	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	resp, r, err := apiClient.V2.GovernanceGroupsApi.ListWorkgroups(context.TODO()).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, r.StatusCode)
}

func Test_cc(t *testing.T) {
	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	resp, r, err := apiClient.CC.AccountsApi.ListAccounts(context.TODO()).Execute()

	require.Nil(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, r.StatusCode)
}
