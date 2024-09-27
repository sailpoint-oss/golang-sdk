# \PasswordManagementAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitToken**](PasswordManagementAPI.md#CreateDigitToken) | **Post** /generate-password-reset-token/digit | Generate a digit token
[**GetPasswordChangeStatus**](PasswordManagementAPI.md#GetPasswordChangeStatus) | **Get** /password-change-status/{id} | Get Password Change Request Status
[**QueryPasswordInfo**](PasswordManagementAPI.md#QueryPasswordInfo) | **Post** /query-password-info | Query Password Info
[**SetPassword**](PasswordManagementAPI.md#SetPassword) | **Post** /set-password | Set Identity&#39;s Password



## CreateDigitToken

> PasswordDigitToken CreateDigitToken(ctx).XSailPointExperimental(xSailPointExperimental).PasswordDigitTokenReset(passwordDigitTokenReset).Execute()

Generate a digit token



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
	passwordDigitTokenReset := *openapiclient.NewPasswordDigitTokenReset("Abby.Smith") // PasswordDigitTokenReset | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordManagementAPI.CreateDigitToken(context.Background()).XSailPointExperimental(xSailPointExperimental).PasswordDigitTokenReset(passwordDigitTokenReset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.CreateDigitToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDigitToken`: PasswordDigitToken
	fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.CreateDigitToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **passwordDigitTokenReset** | [**PasswordDigitTokenReset**](PasswordDigitTokenReset.md) |  | 

### Return type

[**PasswordDigitToken**](PasswordDigitToken.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	id := "089899f13a8f4da7824996191587bab9" // string | Password change request ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordManagementAPI.GetPasswordChangeStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.GetPasswordChangeStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordChangeStatus`: PasswordStatus
	fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.GetPasswordChangeStatus`: %v\n", resp)
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
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

[applicationAuth](../README.md#applicationAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	passwordChangeRequest := *openapiclient.NewPasswordChangeRequest() // PasswordChangeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PasswordManagementAPI.SetPassword(context.Background()).PasswordChangeRequest(passwordChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PasswordManagementAPI.SetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPassword`: PasswordChangeResponse
	fmt.Fprintf(os.Stdout, "Response from `PasswordManagementAPI.SetPassword`: %v\n", resp)
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

[applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

