package sailpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/sailpoint-oss/golang-sdk/v3/api_accounts"
	"github.com/sailpoint-oss/golang-sdk/v3/api_transforms"
	"github.com/sailpoint-oss/golang-sdk/v3/api_workflows"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test_generic validates the generic pass-through API client which is used
// when a typed per-partition client is not available.
func Test_generic(t *testing.T) {

	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {

		resp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "accounts/v1").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})

	t.Run("Test get transforms", func(t *testing.T) {

		resp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "transforms/v1").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})

	t.Run("Test create workflow", func(t *testing.T) {
		randomName := fmt.Sprintf("Test Workflow %s", randString(8))

		jsonStr := fmt.Sprintf(`
		{
			"name": "%s",
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
		}`, randomName)

		var result map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		resp, r, err := apiClient.Generic.DefaultAPI.GenericPost(context.TODO(), "workflows/v1").RequestBody(result).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

		getResp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "workflows/v1").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
			t.FailNow()
		}
		require.Nil(t, err)
		require.NotNil(t, getResp)
		assert.Equal(t, 200, r.StatusCode)

		jsonBytes, err := getResp.MarshalJSON()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error marshaling response: %v\n", err)
			t.FailNow()
		}

		var response []map[string]interface{}
		err = json.Unmarshal(jsonBytes, &response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshaling response: %v\n", err)
			t.FailNow()
		}

		for _, wf := range response {
			if wf["name"] == randomName {
				id, _ := wf["id"].(string)
				_, r, err := apiClient.Generic.DefaultAPI.GenericDelete(context.TODO(), "workflows/v1/"+id).Execute()
				if err != nil {
					fmt.Fprintf(os.Stderr, "during test: %v\n", err)
					t.FailNow()
				}
				require.Nil(t, err)
				assert.Equal(t, 204, r.StatusCode)
			}
		}
	})
}

// Test_accounts validates the typed per-partition accounts client
// via the top-level APIClient (c.AccountsAPI).
func Test_accounts(t *testing.T) {
	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Accounts", func(t *testing.T) {
		resp, r, err := apiClient.AccountsAPI.ListAccountsV1(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}

// Test_transforms validates the typed per-partition transforms client
// via the top-level APIClient (c.TransformsAPI).
func Test_transforms(t *testing.T) {
	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Transforms", func(t *testing.T) {

		resp, r, err := apiClient.TransformsAPI.ListTransformsV1(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}

// Test_workflows validates the typed per-partition workflows client
// via the top-level APIClient (c.WorkflowsAPI). Creates, lists, then deletes a workflow.
func Test_workflows(t *testing.T) {
	configuration := NewDefaultConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test create workflow", func(t *testing.T) {
		randomName := fmt.Sprintf("Test Workflow %s", randString(8))

		trigger := api_workflows.NewWorkflowtrigger("EVENT", map[string]interface{}{
			"id":     "idn:identity-attributes-changed",
			"filter": "$.changes[?(@.attribute == 'manager')]",
		})
		description := "Send an email to the identity who's attributes changed."
		enabled := false
		req := api_workflows.NewCreateWorkflowV1Request(randomName)
		req.SetDescription(description)
		req.SetEnabled(enabled)
		req.SetTrigger(*trigger)

		resp, r, err := apiClient.WorkflowsAPI.CreateWorkflowV1(context.TODO()).
			CreateWorkflowV1Request(*req).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

		workflows, r, err := apiClient.WorkflowsAPI.ListWorkflowsV1(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
			t.FailNow()
		}
		require.Nil(t, err)
		assert.Equal(t, 200, r.StatusCode)

		for _, wf := range workflows {
			if wf.GetName() == randomName {
				r, err := apiClient.WorkflowsAPI.DeleteWorkflowV1(context.TODO(), wf.GetId()).Execute()
				if err != nil {
					fmt.Fprintf(os.Stderr, "during test: %v\n", err)
					t.FailNow()
				}
				require.Nil(t, err)
				assert.Equal(t, 204, r.StatusCode)
			}
		}
	})
}

// Test_typed_direct demonstrates using per-partition packages directly
// without going through the top-level APIClient.
func Test_typed_direct(t *testing.T) {
	configuration := NewDefaultConfiguration()

	t.Run("Test List Accounts direct", func(t *testing.T) {
		cfg := api_accounts.NewConfiguration(NewPartitionConfiguration(configuration))
		client := api_accounts.NewAPIClient(cfg)

		resp, r, err := client.AccountsAPI.ListAccountsV1(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})

	t.Run("Test List Transforms direct", func(t *testing.T) {
		cfg := api_transforms.NewConfiguration(NewPartitionConfiguration(configuration))
		client := api_transforms.NewAPIClient(cfg)

		resp, r, err := client.TransformsAPI.ListTransformsV1(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}

// Test_nerm validates the NERM and NERM v2025 API clients.
// The test is skipped when NermBaseURL is not configured.
func Test_nerm(t *testing.T) {
	configuration := NewDefaultConfiguration()
	if configuration.ClientConfiguration.NermBaseURL == "" {
		t.Skip("NERM not configured: set nermbaseurl in config or SAIL_NERM_BASE_URL env var")
	}
	apiClient := NewAPIClient(configuration)

	t.Run("NERM list users", func(t *testing.T) {
		resp, r, err := apiClient.NERM.UsersAPI.GetUsers(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})

	t.Run("NERM v2025 list delegations", func(t *testing.T) {
		resp, r, err := apiClient.NERMV2025.DelegationsAPI.DelegationsGet(context.TODO()).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)
	})
}

// Helper function to generate a random string
func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
