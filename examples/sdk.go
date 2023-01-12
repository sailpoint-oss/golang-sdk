package main

import (
	"context"
	"fmt"
	"os"

	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
	sailpointsdk "github.com/sailpoint-oss/golang-sdk/sdk-output/v3"
)

func main() {

	ctx := context.TODO()

	configuration := sailpoint.NewConfiguration(sailpoint.ClientConfiguration{ClientId: "", ClientSecret: "", BaseURL: "", TokenURL: ""})

	apiClient := sailpoint.NewAPIClient(configuration)

	getResults(ctx, apiClient)

	getAllPaginatedResults(ctx, apiClient)

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

func getAllPaginatedResults(ctx context.Context, apiClient *sailpoint.APIClient) {
	resp, r, err := sailpoint.Paginate[sailpointsdk.Account](apiClient.V3.AccountsApi.ListAccounts(ctx), 0, 1000)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ListAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccounts`: []Account
	fmt.Fprintf(os.Stdout, "First response from `AccountsApi.ListAccount`: %v\n", resp[0].Name)
}
