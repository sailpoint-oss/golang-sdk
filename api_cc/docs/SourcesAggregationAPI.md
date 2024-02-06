# \SourcesAggregationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadAccounts**](SourcesAggregationAPI.md#LoadAccounts) | **Post** /cc/api/source/loadAccounts/{id} | Account Aggregation (File)
[**LoadEntitlements**](SourcesAggregationAPI.md#LoadEntitlements) | **Post** /cc/api/source/loadEntitlements/{id} | Account Aggregation (File)



## LoadAccounts

> map[string]interface{} LoadAccounts(ctx, id).ContentType(contentType).DisableOptimization(disableOptimization).File(file).Execute()

Account Aggregation (File)



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    disableOptimization := true // bool |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAggregationAPI.LoadAccounts(context.Background(), id).ContentType(contentType).DisableOptimization(disableOptimization).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAggregationAPI.LoadAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SourcesAggregationAPI.LoadAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **disableOptimization** | **bool** |  | 
 **file** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoadEntitlements

> map[string]interface{} LoadEntitlements(ctx, id).ContentType(contentType).File(file).Execute()

Account Aggregation (File)



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
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAggregationAPI.LoadEntitlements(context.Background(), id).ContentType(contentType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAggregationAPI.LoadEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LoadEntitlements`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SourcesAggregationAPI.LoadEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoadEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

