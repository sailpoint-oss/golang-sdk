# \PublicIdentitiesConfigAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicIdentityConfig**](PublicIdentitiesConfigAPI.md#GetPublicIdentityConfig) | **Get** /public-identities-config | Get the Public Identities Configuration
[**UpdatePublicIdentityConfig**](PublicIdentitiesConfigAPI.md#UpdatePublicIdentityConfig) | **Put** /public-identities-config | Update the Public Identities Configuration



## GetPublicIdentityConfig

> PublicIdentityConfig GetPublicIdentityConfig(ctx).Execute()

Get the Public Identities Configuration



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
    resp, r, err := apiClient.PublicIdentitiesConfigAPI.GetPublicIdentityConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIdentitiesConfigAPI.GetPublicIdentityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicIdentityConfig`: PublicIdentityConfig
    fmt.Fprintf(os.Stdout, "Response from `PublicIdentitiesConfigAPI.GetPublicIdentityConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicIdentityConfigRequest struct via the builder pattern


### Return type

[**PublicIdentityConfig**](PublicIdentityConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePublicIdentityConfig

> PublicIdentityConfig UpdatePublicIdentityConfig(ctx).PublicIdentityConfig(publicIdentityConfig).Execute()

Update the Public Identities Configuration



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
    publicIdentityConfig := *openapiclient.NewPublicIdentityConfig() // PublicIdentityConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicIdentitiesConfigAPI.UpdatePublicIdentityConfig(context.Background()).PublicIdentityConfig(publicIdentityConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIdentitiesConfigAPI.UpdatePublicIdentityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePublicIdentityConfig`: PublicIdentityConfig
    fmt.Fprintf(os.Stdout, "Response from `PublicIdentitiesConfigAPI.UpdatePublicIdentityConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePublicIdentityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicIdentityConfig** | [**PublicIdentityConfig**](PublicIdentityConfig.md) |  | 

### Return type

[**PublicIdentityConfig**](PublicIdentityConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

