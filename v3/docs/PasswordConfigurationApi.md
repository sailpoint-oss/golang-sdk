# \PasswordConfigurationApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordOrgConfig**](PasswordConfigurationApi.md#CreatePasswordOrgConfig) | **Post** /password-org-config | Create Password Org Config
[**GetPasswordOrgConfig**](PasswordConfigurationApi.md#GetPasswordOrgConfig) | **Get** /password-org-config | Get Password Org Config
[**UpdatePasswordOrgConfig**](PasswordConfigurationApi.md#UpdatePasswordOrgConfig) | **Put** /password-org-config | Update Password Org Config



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
    openapiclient "./openapi"
)

func main() {
    passwordOrgConfig := *openapiclient.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationApi.CreatePasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationApi.CreatePasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationApi.CreatePasswordOrgConfig`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationApi.GetPasswordOrgConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationApi.GetPasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationApi.GetPasswordOrgConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordOrgConfigRequest struct via the builder pattern


### Return type

[**PasswordOrgConfig**](PasswordOrgConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordOrgConfig

> PasswordOrgConfig UpdatePasswordOrgConfig(ctx).PasswordOrgConfig(passwordOrgConfig).Execute()

Update Password Org Config



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
    passwordOrgConfig := *openapiclient.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordConfigurationApi.UpdatePasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordConfigurationApi.UpdatePasswordOrgConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordOrgConfig`: PasswordOrgConfig
    fmt.Fprintf(os.Stdout, "Response from `PasswordConfigurationApi.UpdatePasswordOrgConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordOrgConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordOrgConfig** | [**PasswordOrgConfig**](PasswordOrgConfig.md) |  | 

### Return type

[**PasswordOrgConfig**](PasswordOrgConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

