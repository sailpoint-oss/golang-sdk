package sailpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

	t.Run("Test get transforms", func(t *testing.T) {

		resp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "v2024/transforms").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})

	t.Run("Test create workflow", func(t *testing.T) {

		jsonStr := `{
  "name": "Send Email2",
  "description": "Send an email to the identity who's attributes changed.",
  "definition": {
    "start": "Send Email Test",
    "steps": {
      "Send Email": {
        "actionId": "sp:send-email",
        "attributes": {
          "body": "This is a test",
          "from": "sailpoint@sailpoint.com",
          "recipientId.$": "$.identity.id",
          "subject": "test"
        },
        "nextStep": "success",
        "selectResult": null,
        "type": "action"
      },
      "success": {
        "type": "success"
      }
    }
  },
  "enabled": false,
  "trigger": {
    "type": "EVENT",
    "attributes": {
      "id": "idn:identity-attributes-changed",
      "filter": "$.changes[?(@.attribute == 'manager')]"
    }
  }
}`

		var result map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		resp, r, err := apiClient.Generic.DefaultAPI.GenericPost(context.TODO(), "v2024/workflows").RequestBody(result).Execute()
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}
		bodyString := string(bodyBytes)
		fmt.Println("Response Body:", bodyString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}
