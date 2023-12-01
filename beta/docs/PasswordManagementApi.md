# \PasswordManagementAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateDigitToken**](PasswordManagementAPI.md#GenerateDigitToken) | **Post** /generate-password-reset-token/digit | Generate a digit token
[**GetIdentityPasswordChangeStatus**](PasswordManagementAPI.md#GetIdentityPasswordChangeStatus) | **Get** /password-change-status/{id} | Get Password Change Request Status
[**QueryPasswordInfo**](PasswordManagementAPI.md#QueryPasswordInfo) | **Post** /query-password-info | Query Password Info
[**SetIdentityPassword**](PasswordManagementAPI.md#SetIdentityPassword) | **Post** /set-password | Set Identity&#39;s Password



## GenerateDigitToken

> PasswordDigitToken GenerateDigitToken(ctx).PasswordDigitTokenReset(passwordDigitTokenReset).Execute()

Generate a digit token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    passwordDigitTokenReset := *openapiclient.NewPasswordDigitTokenReset("Abby.Smith") // PasswordDigitTokenReset | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementAPI.GenerateDigitToken(context.Background()).PasswordDigitTokenReset(passwordDigitTokenReset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.GenerateDigitToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDigitToken`: PasswordDigitToken
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.GenerateDigitToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDigitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordDigitTokenReset** | [**PasswordDigitTokenReset**](PasswordDigitTokenReset.md) |  | 

### Return type

[**PasswordDigitToken**](PasswordDigitToken.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityPasswordChangeStatus

> PasswordStatus GetIdentityPasswordChangeStatus(ctx, id).Execute()

Get Password Change Request Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementAPI.GetIdentityPasswordChangeStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.GetIdentityPasswordChangeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityPasswordChangeStatus`: PasswordStatus
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.GetIdentityPasswordChangeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityPasswordChangeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordStatus**](PasswordStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryPasswordInfo

> PasswordInfo QueryPasswordInfo(ctx).PasswordInfoQueryDTO(passwordInfoQueryDTO).Execute()

Query Password Info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    passwordInfoQueryDTO := *openapiclient.NewPasswordInfoQueryDTO() // PasswordInfoQueryDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementAPI.QueryPasswordInfo(context.Background()).PasswordInfoQueryDTO(passwordInfoQueryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.QueryPasswordInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryPasswordInfo`: PasswordInfo
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.QueryPasswordInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryPasswordInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordInfoQueryDTO** | [**PasswordInfoQueryDTO**](PasswordInfoQueryDTO.md) |  | 

### Return type

[**PasswordInfo**](PasswordInfo.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIdentityPassword

> PasswordChangeResponse SetIdentityPassword(ctx).PasswordChangeRequest(passwordChangeRequest).Execute()

Set Identity's Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    passwordChangeRequest := *openapiclient.NewPasswordChangeRequest() // PasswordChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementAPI.SetIdentityPassword(context.Background()).PasswordChangeRequest(passwordChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.SetIdentityPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetIdentityPassword`: PasswordChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.SetIdentityPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetIdentityPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChangeRequest** | [**PasswordChangeRequest**](PasswordChangeRequest.md) |  | 

### Return type

[**PasswordChangeResponse**](PasswordChangeResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

