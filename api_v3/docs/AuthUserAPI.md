# \AuthUserAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthUser**](AuthUserAPI.md#GetAuthUser) | **Get** /auth-users/{id} | Auth User Details
[**PatchAuthUser**](AuthUserAPI.md#PatchAuthUser) | **Patch** /auth-users/{id} | Auth User Update



## Auth User Details

> AuthUser GetAuthUser(ctx, id).Execute()

Return the specified user's authentication system details.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity ID

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AuthUserAPI.GetAuthUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthUserAPI.GetAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthUserAPI.GetAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthUser**](AuthUser.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Auth User Update

> AuthUser PatchAuthUser(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Use a PATCH request to update an existing user in the authentication system.
Use this endpoint to modify these fields: 
  * `capabilities`

A '400.1.1 Illegal update attempt' detail code indicates that you attempted to PATCH a field that is not allowed.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity ID
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth user update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AuthUserAPI.PatchAuthUser(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthUserAPI.PatchAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAuthUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthUserAPI.PatchAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth user update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**AuthUser**](AuthUser.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

