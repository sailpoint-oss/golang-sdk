# \ReportsDataExtractionAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](ReportsDataExtractionAPI.md#CancelReport) | **Post** /reports/{id}/cancel | Cancel Report
[**GetReport**](ReportsDataExtractionAPI.md#GetReport) | **Get** /reports/{taskResultId} | Get Report File
[**GetReportResult**](ReportsDataExtractionAPI.md#GetReportResult) | **Get** /reports/{taskResultId}/result | Get Report Result
[**StartReport**](ReportsDataExtractionAPI.md#StartReport) | **Post** /reports/run | Run Report



## Cancel Report

> CancelReport(ctx, id).Execute()

Cancels a running report.

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
    id := "a1ed223247144cc29d23c632624b4767" // string | ID of the running Report to cancel

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.ReportsDataExtractionAPI.CancelReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.CancelReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the running Report to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Report File

> *os.File GetReport(ctx, taskResultId).FileFormat(fileFormat).Name(name).Auditable(auditable).Execute()

Gets a report in file format.

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
    taskResultId := "ef38f94347e94562b5bb8424a56397d8" // string | Unique identifier of the task result which handled report
    fileFormat := "csv" // string | Output format of the requested report file
    name := "Identities Details Report" // string | preferred Report file name, by default will be used report name from task result. (optional)
    auditable := true // bool | Enables auditing for current report download. Will create an audit event and sent it to the REPORT cloud-audit kafka topic.  Event will be created if there is any result present by requested taskResultId. (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ReportsDataExtractionAPI.GetReport(context.Background(), taskResultId).FileFormat(fileFormat).Name(name).Auditable(auditable).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReportsDataExtractionAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskResultId** | **string** | Unique identifier of the task result which handled report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileFormat** | **string** | Output format of the requested report file | 
 **name** | **string** | preferred Report file name, by default will be used report name from task result. | 
 **auditable** | **bool** | Enables auditing for current report download. Will create an audit event and sent it to the REPORT cloud-audit kafka topic.  Event will be created if there is any result present by requested taskResultId. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Report Result

> ReportResults GetReportResult(ctx, taskResultId).Completed(completed).Execute()

Get the report results for a report that was run or is running. Returns empty report result in case there are no active task definitions with used in payload task definition name.

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
    taskResultId := "ef38f94347e94562b5bb8424a56397d8" // string | Unique identifier of the task result which handled report
    completed := true // bool | state of task result to apply ordering when results are fetching from the DB (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ReportsDataExtractionAPI.GetReportResult(context.Background(), taskResultId).Completed(completed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.GetReportResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportResult`: ReportResults
    fmt.Fprintf(os.Stdout, "Response from `ReportsDataExtractionAPI.GetReportResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskResultId** | **string** | Unique identifier of the task result which handled report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **completed** | **bool** | state of task result to apply ordering when results are fetching from the DB | [default to false]

### Return type

[**ReportResults**](ReportResults.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Run Report

> TaskResultDetails StartReport(ctx).ReportDetails(reportDetails).Execute()

Runs a report according to input report details. If non-concurrent task is already running then it returns, otherwise new task creates and returns.

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
    reportDetails := *sailpoint.NewReportDetails() // ReportDetails | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ReportsDataExtractionAPI.StartReport(context.Background()).ReportDetails(reportDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.StartReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartReport`: TaskResultDetails
    fmt.Fprintf(os.Stdout, "Response from `ReportsDataExtractionAPI.StartReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportDetails** | [**ReportDetails**](ReportDetails.md) |  | 

### Return type

[**TaskResultDetails**](TaskResultDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

