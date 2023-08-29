# \GlobalTenantSecuritySettingsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsApi.md#CreateAuthOrgNetworkConfig) | **Post** /auth-org/network-config | Create security network configuration.
[**GetAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsApi.md#GetAuthOrgNetworkConfig) | **Get** /auth-org/network-config | Get security network configuration.
[**PatchAuthOrgNetworkConfig**](GlobalTenantSecuritySettingsApi.md#PatchAuthOrgNetworkConfig) | **Patch** /auth-org/network-config | Update security network configuration.



## CreateAuthOrgNetworkConfig

> NetworkConfiguration CreateAuthOrgNetworkConfig(ctx).NetworkConfiguration(networkConfiguration).Execute()

Create security network configuration.



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
    networkConfiguration := *openapiclient.NewNetworkConfiguration() // NetworkConfiguration | Network configuration creation request body.   The following constraints ensure the request body conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalTenantSecuritySettingsApi.CreateAuthOrgNetworkConfig(context.Background()).NetworkConfiguration(networkConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsApi.CreateAuthOrgNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthOrgNetworkConfig`: NetworkConfiguration
    fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsApi.CreateAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthOrgNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkConfiguration** | [**NetworkConfiguration**](NetworkConfiguration.md) | Network configuration creation request body.   The following constraints ensure the request body conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters. | 

### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthOrgNetworkConfig

> NetworkConfiguration GetAuthOrgNetworkConfig(ctx).Execute()

Get security network configuration.



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
    resp, r, err := apiClient.GlobalTenantSecuritySettingsApi.GetAuthOrgNetworkConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsApi.GetAuthOrgNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthOrgNetworkConfig`: NetworkConfiguration
    fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsApi.GetAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthOrgNetworkConfigRequest struct via the builder pattern


### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthOrgNetworkConfig

> NetworkConfiguration PatchAuthOrgNetworkConfig(ctx).JsonPatchOperation(jsonPatchOperation).Execute()

Update security network configuration.



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
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth org network configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Network Config conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GlobalTenantSecuritySettingsApi.PatchAuthOrgNetworkConfig(context.Background()).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalTenantSecuritySettingsApi.PatchAuthOrgNetworkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAuthOrgNetworkConfig`: NetworkConfiguration
    fmt.Fprintf(os.Stdout, "Response from `GlobalTenantSecuritySettingsApi.PatchAuthOrgNetworkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthOrgNetworkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth org network configuration update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Ensures that the patched Network Config conforms to certain logical guidelines, which are:   1. Each string element in the range array must be a valid ip address or ip subnet mask.   2. Each string element in the geolocation array must be 2 characters, and they can only be uppercase letters. | 

### Return type

[**NetworkConfiguration**](NetworkConfiguration.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

