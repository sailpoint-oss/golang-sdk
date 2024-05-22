# \SourcesUncorrelatedAccountsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImportUncorrelatedAccounts**](SourcesUncorrelatedAccountsAPI.md#ImportUncorrelatedAccounts) | **Post** /sources/{id}/load-uncorrelated-accounts | Process Uncorrelated Accounts



## ImportUncorrelatedAccounts

> LoadUncorrelatedAccountsTask ImportUncorrelatedAccounts(ctx, id).File(file).Execute()

Process Uncorrelated Accounts



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
    id := "75dbec1ebe154d5785da27b95e1dd5d7" // string | Source Id
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesUncorrelatedAccountsAPI.ImportUncorrelatedAccounts(context.Background(), id).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesUncorrelatedAccountsAPI.ImportUncorrelatedAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportUncorrelatedAccounts`: LoadUncorrelatedAccountsTask
    fmt.Fprintf(os.Stdout, "Response from `SourcesUncorrelatedAccountsAPI.ImportUncorrelatedAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportUncorrelatedAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**LoadUncorrelatedAccountsTask**](LoadUncorrelatedAccountsTask.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

