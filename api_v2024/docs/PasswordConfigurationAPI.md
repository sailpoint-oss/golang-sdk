# \PasswordConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordOrgConfig**](PasswordConfigurationAPI.md#CreatePasswordOrgConfig) | **Post** /password-org-config | Create Password Org Config
[**GetPasswordOrgConfig**](PasswordConfigurationAPI.md#GetPasswordOrgConfig) | **Get** /password-org-config | Get Password Org Config
[**PutPasswordOrgConfig**](PasswordConfigurationAPI.md#PutPasswordOrgConfig) | **Put** /password-org-config | Update Password Org Config



## CreatePasswordOrgConfig

> PasswordOrgConfig CreatePasswordOrgConfig(ctx).PasswordOrgConfig(passwordOrgConfig).Execute()

Create Password Org Config



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
    passwordOrgConfig := *openapiclient.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationAPI.CreatePasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationAPI.CreatePasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationAPI.CreatePasswordOrgConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordOrgConfig** | [**PasswordOrgConfig**](PasswordOrgConfig.md) |  | 

### Return type

[**PasswordOrgConfig**](PasswordOrgConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordOrgConfig

> PasswordOrgConfig GetPasswordOrgConfig(ctx).Execute()

Get Password Org Config



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationAPI.GetPasswordOrgConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationAPI.GetPasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationAPI.GetPasswordOrgConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordOrgConfigRequest struct via the builder pattern


### Return type

[**PasswordOrgConfig**](PasswordOrgConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPasswordOrgConfig

> PasswordOrgConfig PutPasswordOrgConfig(ctx).PasswordOrgConfig(passwordOrgConfig).Execute()

Update Password Org Config



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
    passwordOrgConfig := *openapiclient.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationAPI.PutPasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationAPI.PutPasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationAPI.PutPasswordOrgConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPasswordOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordOrgConfig** | [**PasswordOrgConfig**](PasswordOrgConfig.md) |  | 

### Return type

[**PasswordOrgConfig**](PasswordOrgConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

