# \SystemApi

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshIdentities**](SystemApi.md#RefreshIdentities) | **Post** /cc/api/system/refreshIdentities | Refresh Identities



## RefreshIdentities

> map[string]interface{} RefreshIdentities(ctx).ContentType(contentType).RefreshIdentitiesRequest(refreshIdentitiesRequest).Execute()

Refresh Identities



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
    contentType := "application/json" // string |  (optional)
    refreshIdentitiesRequest := *openapiclient.NewRefreshIdentitiesRequest() // RefreshIdentitiesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.RefreshIdentities(context.Background()).ContentType(contentType).RefreshIdentitiesRequest(refreshIdentitiesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.RefreshIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshIdentities`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SystemApi.RefreshIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 
 **refreshIdentitiesRequest** | [**RefreshIdentitiesRequest**](RefreshIdentitiesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

