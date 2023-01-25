# \SourcesAggregationApi

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoadAccounts**](SourcesAggregationApi.md#LoadAccounts) | **Post** /cc/api/source/loadAccounts/{id} | Account Aggregation (File)



## LoadAccounts

> LoadAccounts(ctx, id).ContentType(contentType).DisableOptimization(disableOptimization).File(file).Execute()

Account Aggregation (File)



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
    id := "id_example" // string | 
    contentType := "application/x-www-form-urlencoded" // string |  (optional)
    disableOptimization := true // bool |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAggregationApi.LoadAccounts(context.Background(), id).ContentType(contentType).DisableOptimization(disableOptimization).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAggregationApi.LoadAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

