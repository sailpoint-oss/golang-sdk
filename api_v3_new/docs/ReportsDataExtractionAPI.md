# \ReportsDataExtractionAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](ReportsDataExtractionAPI.md#cancel-report) | **Post** /reports/{id}/cancel | Cancel Report
[**GetReport**](ReportsDataExtractionAPI.md#get-report) | **Get** /reports/{taskResultId} | Get Report File
[**GetReportResult**](ReportsDataExtractionAPI.md#get-report-result) | **Get** /reports/{taskResultId}/result | Get Report Result
[**StartReport**](ReportsDataExtractionAPI.md#start-report) | **Post** /reports/run | Run Report



## cancel-report


Cancels a running report.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | id | **string** | True  | ID of the running Report to cancel


### Return type

 (empty response body)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
204 | No content - indicates the request was successful but there is no content to be returned in the response. | 
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //CancelReport

    id := "a1ed223247144cc29d23c632624b4767"



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.ReportsDataExtractionAPI.CancelReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.CancelReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```




## get-report


Gets a report in file format.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | taskResultId | **string** | True  | Unique identifier of the task result which handled report
  Query | fileFormat | **string** | True  | Output format of the requested report file
  Query | name | **string** |   (optional) | preferred Report file name, by default will be used report name from task result.
  Query | auditable | **bool** |   (optional) (default to false) | Enables auditing for current report download. Will create an audit event and sent it to the REPORT cloud-audit kafka topic.  Event will be created if there is any result present by requested taskResultId.


### Return type

[***os.File**](*os.File.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | Report file in selected format. CSV by default. | *os.File
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //GetReport

    taskResultId := "ef38f94347e94562b5bb8424a56397d8"
    fileFormat := "csv"
    //name := "Identities Details Report"
    //auditable := true



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ReportsDataExtractionAPI.GetReport(context.Background(), taskResultId).FileFormat(fileFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ReportsDataExtractionAPI.GetReport`: %v\n", resp)
}
```




## get-report-result


Get the report results for a report that was run or is running. Returns empty report result in case there are no active task definitions with used in payload task definition name.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | taskResultId | **string** | True  | Unique identifier of the task result which handled report
  Query | completed | **bool** |   (optional) (default to false) | state of task result to apply ordering when results are fetching from the DB


### Return type

[**ReportResults**](ReportResults.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | Details about report that was run or is running. | ReportResults
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //GetReportResult

    taskResultId := "ef38f94347e94562b5bb8424a56397d8"
    //completed := true



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ReportsDataExtractionAPI.GetReportResult(context.Background(), taskResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsDataExtractionAPI.GetReportResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportResult`: ReportResults
    fmt.Fprintf(os.Stdout, "Response from `ReportsDataExtractionAPI.GetReportResult`: %v\n", resp)
}
```




## start-report


Runs a report according to input report details. If non-concurrent task is already running then it returns, otherwise new task creates and returns.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
 Body  | reportDetails | [**ReportDetails**](ReportDetails.md) | True  | 


### Return type

[**TaskResultDetails**](TaskResultDetails.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | Details about running report task. | TaskResultDetails
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //StartReport

    reportDetails := *sailpoint.NewReportDetails()



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



