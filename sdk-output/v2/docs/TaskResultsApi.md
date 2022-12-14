# \TaskResultsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackgroundTaskResults**](TaskResultsApi.md#GetBackgroundTaskResults) | **Get** /task-results/{taskResultIdOrName} | Retrieve Result of Background Task



## GetBackgroundTaskResults

> TaskResult GetBackgroundTaskResults(ctx, taskResultIdOrName).Execute()

Retrieve Result of Background Task



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
    taskResultIdOrName := "taskResultIdOrName_example" // string | ID or name of a task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskResultsApi.GetBackgroundTaskResults(context.Background(), taskResultIdOrName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskResultsApi.GetBackgroundTaskResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackgroundTaskResults`: TaskResult
    fmt.Fprintf(os.Stdout, "Response from `TaskResultsApi.GetBackgroundTaskResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskResultIdOrName** | **string** | ID or name of a task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackgroundTaskResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskResult**](TaskResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

