package sailpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	beta "github.com/sailpoint-oss/golang-sdk/v2/api_beta"
	v2024 "github.com/sailpoint-oss/golang-sdk/v2/api_v2024"
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

func Test_v2026(t *testing.T) {
	configuration := NewDefaultConfiguration()
	configuration.Experimental = true
	apiClient := NewAPIClient(configuration)

	t.Run("Test List Tasks", func(t *testing.T) {
		resp, r, err := apiClient.V2026.TaskManagementAPI.GetTaskStatusList(context.TODO()).Execute()
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
		// Generate a random string for the workflow name
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

		resp, r, err := apiClient.Generic.DefaultAPI.GenericPost(context.TODO(), "v2024/workflows").RequestBody(result).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test`: %v\n", err)
		}
		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, r.StatusCode)

		getResp, r, err := apiClient.Generic.DefaultAPI.GenericGet(context.TODO(), "v2024/workflows").Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "during test: %v\n", err)
			t.FailNow()
		}
		require.Nil(t, err)
		require.NotNil(t, getResp)
		assert.Equal(t, 200, r.StatusCode)

		// Convert getResp to []v2024.Workflow
		jsonBytes, err := getResp.MarshalJSON()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error marshaling response: %v\n", err)
			t.FailNow()
		}

		var response []v2024.Workflow
		err = json.Unmarshal(jsonBytes, &response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error unmarshaling response: %v\n", err)
			t.FailNow()
		}

		for _, wf := range response {
			if *wf.Name == randomName {
				_, r, err := apiClient.Generic.DefaultAPI.GenericDelete(context.TODO(), "v2024/workflows/"+*wf.Id).Execute()
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

// Helper function to generate a random string
func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// ============================================================================
// DEVREL-2463: Enum Forward Compatibility Tests
// ============================================================================
// Tests that enum types accept unknown values for forward compatibility
// while maintaining backward compatibility with existing code.

// Test that known enum values work as before (backward compatibility)
func TestEnumUnmarshal_KnownValue(t *testing.T) {
	jsonData := []byte(`"ACCOUNT"`)

	var dtoType v3.DtoType
	err := json.Unmarshal(jsonData, &dtoType)

	require.NoError(t, err, "Unmarshaling known enum value should succeed")
	assert.Equal(t, v3.DTOTYPE_ACCOUNT, dtoType, "Enum should match known constant")
	assert.True(t, dtoType.IsValid(), "Known value should be valid")
}

// Test that unknown enum values are accepted (forward compatibility)
func TestEnumUnmarshal_UnknownValue(t *testing.T) {
	jsonData := []byte(`"FUTURE_TYPE"`)

	var dtoType v3.DtoType
	err := json.Unmarshal(jsonData, &dtoType)

	require.NoError(t, err, "Unmarshaling unknown enum value should succeed")
	assert.Equal(t, v3.DtoType("FUTURE_TYPE"), dtoType, "Enum should store raw value")
	assert.False(t, dtoType.IsValid(), "Unknown value should not be valid")
}

// Test that IsValid() works correctly for known values
func TestEnumIsValid_KnownValue(t *testing.T) {
	dtoType := v3.DTOTYPE_IDENTITY

	assert.True(t, dtoType.IsValid(), "Known enum constant should be valid")
}

// Test that IsValid() returns false for unknown values
func TestEnumIsValid_UnknownValue(t *testing.T) {
	dtoType := v3.DtoType("UNKNOWN_TYPE")

	assert.False(t, dtoType.IsValid(), "Unknown enum value should not be valid")
}

// Test round-trip marshaling preserves unknown values
func TestEnumMarshal_UnknownValue(t *testing.T) {
	original := []byte(`"NEW_ENUM_VALUE"`)

	var dtoType v3.DtoType
	err := json.Unmarshal(original, &dtoType)
	require.NoError(t, err, "Unmarshal should succeed")

	marshaled, err := json.Marshal(dtoType)
	require.NoError(t, err, "Marshal should succeed")

	assert.JSONEq(t, string(original), string(marshaled),
		"Round-trip should preserve unknown enum value")
}

// Test that NewXXXFromValue accepts known values
func TestEnumConstructor_KnownValue(t *testing.T) {
	dtoType, err := v3.NewDtoTypeFromValue("ACCOUNT")

	require.NoError(t, err, "Constructor should accept known value")
	require.NotNil(t, dtoType, "Constructor should return non-nil pointer")
	assert.Equal(t, v3.DTOTYPE_ACCOUNT, *dtoType, "Value should match constant")
	assert.True(t, dtoType.IsValid(), "Known value should be valid")
}

// Test that NewXXXFromValue accepts unknown values (NEW behavior)
func TestEnumConstructor_UnknownValue(t *testing.T) {
	dtoType, err := v3.NewDtoTypeFromValue("MODIFY_ACCESS")

	require.NoError(t, err, "Constructor should accept unknown value")
	require.NotNil(t, dtoType, "Constructor should return non-nil pointer")
	assert.Equal(t, v3.DtoType("MODIFY_ACCESS"), *dtoType, "Value should be stored")
	assert.False(t, dtoType.IsValid(), "Unknown value should not be valid")
}

// Test the real-world scenario that was failing (customer example)
func TestAccessRequestType_ModifyAccess(t *testing.T) {
	// Simulate API response with new MODIFY_ACCESS value
	jsonData := []byte(`"MODIFY_ACCESS"`)

	var dtoType v3.DtoType // Using DtoType as proxy since this tests enum behavior
	err := json.Unmarshal(jsonData, &dtoType)

	require.NoError(t, err,
		"Should handle MODIFY_ACCESS enum value without error (customer scenario)")
	assert.Equal(t, v3.DtoType("MODIFY_ACCESS"), dtoType)
}

// Test struct with enum field handles unknown value
func TestStructWithEnum_UnknownValue(t *testing.T) {
	type TestStruct struct {
		Type v3.DtoType `json:"type"`
		Name string     `json:"name"`
	}

	jsonData := []byte(`{"type": "FUTURE_TYPE", "name": "test"}`)

	var obj TestStruct
	err := json.Unmarshal(jsonData, &obj)

	require.NoError(t, err, "Struct unmarshal should succeed with unknown enum")
	assert.Equal(t, v3.DtoType("FUTURE_TYPE"), obj.Type)
	assert.Equal(t, "test", obj.Name)
	assert.False(t, obj.Type.IsValid(), "Unknown enum in struct should not be valid")
}

// Test Ptr() method still works with unknown values
func TestEnumPtr_UnknownValue(t *testing.T) {
	dtoType := v3.DtoType("UNKNOWN")
	ptr := dtoType.Ptr()

	require.NotNil(t, ptr, "Ptr() should return non-nil pointer")
	assert.Equal(t, dtoType, *ptr, "Dereferenced pointer should match original value")
}

// Test NullableXXX with unknown value
func TestNullableEnum_UnknownValue(t *testing.T) {
	jsonData := []byte(`"FUTURE_TYPE"`)

	var nullable v3.NullableDtoType
	err := json.Unmarshal(jsonData, &nullable)

	require.NoError(t, err, "Nullable enum unmarshal should succeed")
	assert.True(t, nullable.IsSet(), "Nullable should be set")

	value := nullable.Get()
	require.NotNil(t, value, "Nullable value should not be nil")
	assert.Equal(t, v3.DtoType("FUTURE_TYPE"), *value)
	assert.False(t, value.IsValid(), "Unknown value should not be valid")
}
