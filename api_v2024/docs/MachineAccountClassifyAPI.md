# \MachineAccountClassifyAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendClassifyMachineAccount**](MachineAccountClassifyAPI.md#SendClassifyMachineAccount) | **Post** /accounts/{id}/classify | Classify a Single Machine Account



## SendClassifyMachineAccount

> SendClassifyMachineAccount200Response SendClassifyMachineAccount(ctx, id).ClassificationMode(classificationMode).Execute()

Classify a Single Machine Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Account ID.
	classificationMode := "forceMachine" // string | Specifies how the accounts should be classified.        default - uses criteria to classify account as machine or human, excludes accounts that were manually classified.       ignoreManual - like default, but includes accounts that were manually classified.       forceMachine - forces account to be classified as machine.       forceHuman - forces account to be classified as human. (optional) (default to "default")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAccountClassifyAPI.SendClassifyMachineAccount(context.Background(), id).ClassificationMode(classificationMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountClassifyAPI.SendClassifyMachineAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendClassifyMachineAccount`: SendClassifyMachineAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `MachineAccountClassifyAPI.SendClassifyMachineAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendClassifyMachineAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classificationMode** | **string** | Specifies how the accounts should be classified.        default - uses criteria to classify account as machine or human, excludes accounts that were manually classified.       ignoreManual - like default, but includes accounts that were manually classified.       forceMachine - forces account to be classified as machine.       forceHuman - forces account to be classified as human. | [default to &quot;default&quot;]

### Return type

[**SendClassifyMachineAccount200Response**](SendClassifyMachineAccount200Response.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

