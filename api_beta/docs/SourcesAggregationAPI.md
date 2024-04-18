# \SourcesAggregationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportAccounts**](SourcesAggregationAPI.md#ImportAccounts) | **Post** /sources/{id}/load-accounts | Account Aggregation



## ImportAccounts

> LoadAccountsTask ImportAccounts(ctx, id).File(file).DisableOptimization(disableOptimization).Execute()

Account Aggregation



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Source Id
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)
    disableOptimization := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAggregationAPI.ImportAccounts(context.Background(), id).File(file).DisableOptimization(disableOptimization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAggregationAPI.ImportAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportAccounts`: LoadAccountsTask
    fmt.Fprintf(os.Stdout, "Response from `SourcesAggregationAPI.ImportAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 
 **disableOptimization** | **bool** |  | 

### Return type

[**LoadAccountsTask**](LoadAccountsTask.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

