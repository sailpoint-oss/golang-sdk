# \TenantAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenant**](TenantAPI.md#GetTenant) | **Get** /tenant | Get Tenant Information.



## GetTenant

> Tenant GetTenant(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

Get Tenant Information.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TenantAPI.GetTenant(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenant`: Tenant
    fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Tenant**](Tenant.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

