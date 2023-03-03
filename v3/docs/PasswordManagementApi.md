# \PasswordManagementApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordChangeStatus**](PasswordManagementApi.md#GetPasswordChangeStatus) | **Get** /password-change-status/{id} | Get Password Change Request Status
[**QueryPasswordInfo**](PasswordManagementApi.md#QueryPasswordInfo) | **Post** /query-password-info | Query Password Info
[**SetPassword**](PasswordManagementApi.md#SetPassword) | **Post** /set-password | Set Identity&#39;s Password



## GetPasswordChangeStatus

> PasswordStatus GetPasswordChangeStatus(ctx, id).Execute()

Get Password Change Request Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "089899f13a8f4da7824996191587bab9" // string | Password change request ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementApi.GetPasswordChangeStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementApi.GetPasswordChangeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordChangeStatus`: PasswordStatus
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementApi.GetPasswordChangeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Password change request ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordChangeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordStatus**](PasswordStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    passwordInfoQueryDTO := *openapiclient.NewPasswordInfoQueryDTO() // PasswordInfoQueryDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementApi.QueryPasswordInfo(context.Background()).PasswordInfoQueryDTO(passwordInfoQueryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementApi.QueryPasswordInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryPasswordInfo`: PasswordInfo
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementApi.QueryPasswordInfo`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPassword

> PasswordChangeResponse SetPassword(ctx).PasswordChangeRequest(passwordChangeRequest).Execute()

Set Identity's Password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    passwordChangeRequest := *openapiclient.NewPasswordChangeRequest() // PasswordChangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordManagementApi.SetPassword(context.Background()).PasswordChangeRequest(passwordChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementApi.SetPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPassword`: PasswordChangeResponse
    fmt.Fprintf(os.Stdout, "Response from `PasswordManagementApi.SetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChangeRequest** | [**PasswordChangeRequest**](PasswordChangeRequest.md) |  | 

### Return type

[**PasswordChangeResponse**](PasswordChangeResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

