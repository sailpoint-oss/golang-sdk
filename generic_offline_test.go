package sailpoint

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sailpoint-oss/golang-sdk/v3/generic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// capturedRequest is what the test server saw for one generic request.
type capturedRequest struct {
	method      string
	rawPath     string
	query       map[string][]string
	contentType string
	auth        string
	body        string
}

// newGenericTestClient returns a generic client pointed at a local server, so
// these tests need no tenant.
func newGenericTestClient(t *testing.T, response string) (*generic.APIClient, *capturedRequest) {
	t.Helper()
	captured := &capturedRequest{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*captured = capturedRequest{
			method:      r.Method,
			rawPath:     r.URL.EscapedPath(),
			query:       r.URL.Query(),
			contentType: r.Header.Get("Content-Type"),
			auth:        r.Header.Get("Authorization"),
			body:        string(body),
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, response)
	}))
	t.Cleanup(server.Close)

	cfg := generic.NewConfiguration("", "", server.URL, "", "token-1", "test", false)
	return generic.NewAPIClient(cfg), captured
}

func Test_generic_offline(t *testing.T) {
	t.Run("keeps the slashes of a multi-segment path and adds no version", func(t *testing.T) {
		client, captured := newGenericTestClient(t, `[{"id":"1"}]`)

		resp, r, err := client.DefaultAPI.GenericGet(context.TODO(), "/entitlements/v1/identities/2c91808a/entitlements").
			Limit(10).Filters(`name eq "x"`).Execute()

		require.NoError(t, err)
		assert.Equal(t, 200, r.StatusCode)
		assert.Equal(t, "/entitlements/v1/identities/2c91808a/entitlements", captured.rawPath)
		assert.Equal(t, []string{"10"}, captured.query["limit"])
		assert.Equal(t, []string{`name eq "x"`}, captured.query["filters"])
		assert.Equal(t, "Bearer token-1", captured.auth)
		assert.Equal(t, []interface{}{map[string]interface{}{"id": "1"}}, resp)
	})

	t.Run("encodes each segment and keeps a query string from the path", func(t *testing.T) {
		client, captured := newGenericTestClient(t, `{}`)

		resp, _, err := client.DefaultAPI.GenericGet(context.TODO(), "accounts/v1/a b?detailLevel=SLIM").Limit(5).Execute()

		require.NoError(t, err)
		assert.Equal(t, "/accounts/v1/a%20b", captured.rawPath)
		assert.Equal(t, []string{"SLIM"}, captured.query["detailLevel"])
		assert.Equal(t, []string{"5"}, captured.query["limit"])
		assert.Equal(t, map[string]interface{}{}, resp)
	})

	t.Run("sends a JSON Patch list with the json-patch content type", func(t *testing.T) {
		client, captured := newGenericTestClient(t, `{}`)
		patch := []map[string]interface{}{{"op": "replace", "path": "/description", "value": "New"}}

		_, _, err := client.DefaultAPI.GenericPatch(context.TODO(), "transforms/v1/1234").RequestBody(patch).Execute()

		require.NoError(t, err)
		assert.Equal(t, http.MethodPatch, captured.method)
		assert.Equal(t, "application/json-patch+json", captured.contentType)
		var sent []map[string]interface{}
		require.NoError(t, json.Unmarshal([]byte(captured.body), &sent))
		assert.Equal(t, patch, sent)
	})

	t.Run("sends a POST body as JSON to a v2 path", func(t *testing.T) {
		client, captured := newGenericTestClient(t, `{"id":"1"}`)

		_, _, err := client.DefaultAPI.GenericPost(context.TODO(), "sources/v2/2c91808a/provisioning-policies").
			RequestBody(map[string]interface{}{"name": "Test", "type": "lower"}).Execute()

		require.NoError(t, err)
		assert.Equal(t, http.MethodPost, captured.method)
		assert.Equal(t, "/sources/v2/2c91808a/provisioning-policies", captured.rawPath)
		assert.JSONEq(t, `{"name":"Test","type":"lower"}`, captured.body)
	})
}
