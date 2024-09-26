# \CustomPasswordInstructionsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#CreateCustomPasswordInstructions) | **Post** /custom-password-instructions | Create Custom Password Instructions
[**DeleteCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#DeleteCustomPasswordInstructions) | **Delete** /custom-password-instructions/{pageId} | Delete Custom Password Instructions by page ID
[**GetCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#GetCustomPasswordInstructions) | **Get** /custom-password-instructions/{pageId} | Get Custom Password Instructions by Page ID



## CreateCustomPasswordInstructions

> CustomPasswordInstruction CreateCustomPasswordInstructions(ctx).CustomPasswordInstruction(customPasswordInstruction).Execute()

Create Custom Password Instructions



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
	customPasswordInstruction := *openapiclient.NewCustomPasswordInstruction() // CustomPasswordInstruction | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPasswordInstructionsAPI.CreateCustomPasswordInstructions(context.Background()).CustomPasswordInstruction(customPasswordInstruction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPasswordInstructionsAPI.CreateCustomPasswordInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomPasswordInstructions`: CustomPasswordInstruction
	fmt.Fprintf(os.Stdout, "Response from `CustomPasswordInstructionsAPI.CreateCustomPasswordInstructions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomPasswordInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customPasswordInstruction** | [**CustomPasswordInstruction**](CustomPasswordInstruction.md) |  | 

### Return type

[**CustomPasswordInstruction**](CustomPasswordInstruction.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomPasswordInstructions

> DeleteCustomPasswordInstructions(ctx, pageId).Locale(locale).Execute()

Delete Custom Password Instructions by page ID



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
	pageId := "mfa:select" // string | The page ID of custom password instructions to delete.
	locale := "locale_example" // string | The locale for the custom instructions, a BCP47 language tag. The default value is \\\"default\\\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomPasswordInstructionsAPI.DeleteCustomPasswordInstructions(context.Background(), pageId).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPasswordInstructionsAPI.DeleteCustomPasswordInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The page ID of custom password instructions to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomPasswordInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;. | 

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPasswordInstructions

> CustomPasswordInstruction GetCustomPasswordInstructions(ctx, pageId).Locale(locale).Execute()

Get Custom Password Instructions by Page ID



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
	pageId := "mfa:select" // string | The page ID of custom password instructions to query.
	locale := "locale_example" // string | The locale for the custom instructions, a BCP47 language tag. The default value is \\\"default\\\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomPasswordInstructionsAPI.GetCustomPasswordInstructions(context.Background(), pageId).Locale(locale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomPasswordInstructionsAPI.GetCustomPasswordInstructions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomPasswordInstructions`: CustomPasswordInstruction
	fmt.Fprintf(os.Stdout, "Response from `CustomPasswordInstructionsAPI.GetCustomPasswordInstructions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pageId** | **string** | The page ID of custom password instructions to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomPasswordInstructionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locale** | **string** | The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;. | 

### Return type

[**CustomPasswordInstruction**](CustomPasswordInstruction.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

