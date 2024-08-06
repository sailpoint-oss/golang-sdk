# \CustomPasswordInstructionsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#CreateCustomPasswordInstructions) | **Post** /custom-password-instructions | Create Custom Password Instructions
[**DeleteCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#DeleteCustomPasswordInstructions) | **Delete** /custom-password-instructions/{pageId} | Delete Custom Password Instructions by page ID
[**GetCustomPasswordInstructions**](CustomPasswordInstructionsAPI.md#GetCustomPasswordInstructions) | **Get** /custom-password-instructions/{pageId} | Get Custom Password Instructions by Page ID



## CreateCustomPasswordInstructions

> CustomPasswordInstruction CreateCustomPasswordInstructions(ctx).XSailPointExperimental(xSailPointExperimental).CustomPasswordInstruction(customPasswordInstruction).Execute()

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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    customPasswordInstruction := *openapiclient.NewCustomPasswordInstruction() // CustomPasswordInstruction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPasswordInstructionsAPI.CreateCustomPasswordInstructions(context.Background()).XSailPointExperimental(xSailPointExperimental).CustomPasswordInstruction(customPasswordInstruction).Execute()
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
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **customPasswordInstruction** | [**CustomPasswordInstruction**](CustomPasswordInstruction.md) |  | 

### Return type

[**CustomPasswordInstruction**](CustomPasswordInstruction.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomPasswordInstructions

> DeleteCustomPasswordInstructions(ctx, pageId).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()

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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    locale := "locale_example" // string | The locale for the custom instructions, a BCP47 language tag. The default value is \\\"default\\\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomPasswordInstructionsAPI.DeleteCustomPasswordInstructions(context.Background(), pageId).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **locale** | **string** | The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;. | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomPasswordInstructions

> CustomPasswordInstruction GetCustomPasswordInstructions(ctx, pageId).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()

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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    locale := "locale_example" // string | The locale for the custom instructions, a BCP47 language tag. The default value is \\\"default\\\". (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPasswordInstructionsAPI.GetCustomPasswordInstructions(context.Background(), pageId).XSailPointExperimental(xSailPointExperimental).Locale(locale).Execute()
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **locale** | **string** | The locale for the custom instructions, a BCP47 language tag. The default value is \\\&quot;default\\\&quot;. | 

### Return type

[**CustomPasswordInstruction**](CustomPasswordInstruction.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

