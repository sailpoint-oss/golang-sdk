package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
	v3 "github.com/sailpoint-oss/golang-sdk/sdk-output/v3"
)

func main() {

	ctx := context.TODO()

	configuration := sailpoint.NewDefaultConfiguration()

	apiClient := sailpoint.NewAPIClient(configuration)
	configuration.HTTPClient.RetryMax = 10

	getSearchResults(ctx, apiClient)

	//getAllPaginatedResults(ctx, apiClient)

}

func getResults(ctx context.Context, apiClient *sailpoint.APIClient) {
	resp, r, err := apiClient.V3.AccountsApi.ListAccounts(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0].Name)
}

func getSearchResults(ctx context.Context, apiClient *sailpoint.APIClient) {
	search := v3.NewSearch1WithDefaults()
	searchString := []byte(`
	{
	"indices": [
		"identities"
	],
	"query": {
		"query": "\"philip.ellis\"",
		"fields": [
		"name"
		]
	}
	}
	  `)
	search.UnmarshalJSON(searchString)
	resp, r, err := apiClient.V3.SearchApi.SearchPost(ctx).Search1(*search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0])
}

func getTransformResults(ctx context.Context, apiClient *sailpoint.APIClient) {
	resp, r, err := apiClient.V3.TransformsApi.GetTransformsList(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.GetTransformsList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	b, _ := json.Marshal(resp[0].Attributes)
	fmt.Fprintf(os.Stdout, "First response from `TransformsApi.GetTransformsList`: %v\n", string(b))
}

func getAllPaginatedResults(ctx context.Context, apiClient *sailpoint.APIClient) {
	resp, r, err := sailpoint.PaginateWithDefaults[v3.Account](apiClient.V3.AccountsApi.ListAccounts(ctx))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0].Name)
}
