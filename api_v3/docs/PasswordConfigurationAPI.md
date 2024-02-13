# \PasswordConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordOrgConfig**](PasswordConfigurationAPI.md#CreatePasswordOrgConfig) | **Post** /password-org-config | Create Password Org Config
[**GetPasswordOrgConfig**](PasswordConfigurationAPI.md#GetPasswordOrgConfig) | **Get** /password-org-config | Get Password Org Config
[**PutPasswordOrgConfig**](PasswordConfigurationAPI.md#PutPasswordOrgConfig) | **Put** /password-org-config | Update Password Org Config



## Create Password Org Config

> PasswordOrgConfig CreatePasswordOrgConfig(ctx).PasswordOrgConfig(passwordOrgConfig).Execute()

This API creates the password org config. Unspecified fields will use default value.
To be able to use the custom password instructions, you must set the `customInstructionsEnabled` field to "true".
Requires ORG_ADMIN, API role or authorization scope of 'idn:password-org-config:write'

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
    passwordOrgConfig := *sailpoint.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordConfigurationAPI.CreatePasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
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


## Get Password Org Config

> PasswordOrgConfig GetPasswordOrgConfig(ctx).Execute()

This API returns the password org config . Requires ORG_ADMIN, API role or authorization scope of 'idn:password-org-config:read'

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

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordConfigurationAPI.GetPasswordOrgConfig(context.Background()).Execute()
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


## Update Password Org Config

> PasswordOrgConfig PutPasswordOrgConfig(ctx).PasswordOrgConfig(passwordOrgConfig).Execute()

This API updates the password org config for specified fields. Other fields will keep original value.
You must set the `customInstructionsEnabled` field to "true" to be able to use custom password instructions. 
Requires ORG_ADMIN, API role or authorization scope of 'idn:password-org-config:write'

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
    passwordOrgConfig := *sailpoint.NewPasswordOrgConfig() // PasswordOrgConfig | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordConfigurationAPI.PutPasswordOrgConfig(context.Background()).PasswordOrgConfig(passwordOrgConfig).Execute()
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

