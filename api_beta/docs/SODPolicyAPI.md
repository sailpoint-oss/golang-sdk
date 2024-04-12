# \SODPolicyAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomViolationReport**](SODPolicyAPI.md#GetCustomViolationReport) | **Get** /sod-violation-report/{reportResultId}/download/{fileName} | Download custom violation report
[**GetDefaultViolationReport**](SODPolicyAPI.md#GetDefaultViolationReport) | **Get** /sod-violation-report/{reportResultId}/download | Download violation report
[**GetSodAllReportRunStatus**](SODPolicyAPI.md#GetSodAllReportRunStatus) | **Get** /sod-violation-report | Get multi-report run task status
[**GetSodViolationReportRunStatus**](SODPolicyAPI.md#GetSodViolationReportRunStatus) | **Get** /sod-policies/sod-violation-report-status/{reportResultId} | Get violation report run status
[**StartSodAllPoliciesForOrg**](SODPolicyAPI.md#StartSodAllPoliciesForOrg) | **Post** /sod-violation-report/run | Runs all policies for org



## GetCustomViolationReport

> *os.File GetCustomViolationReport(ctx, reportResultId, fileName).Execute()

Download custom violation report



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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.
    fileName := "custom-name" // string | Custom Name for the  file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyAPI.GetCustomViolationReport(context.Background(), reportResultId, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetCustomViolationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomViolationReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetCustomViolationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportResultId** | **string** | The ID of the report reference to download. | 
**fileName** | **string** | Custom Name for the  file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomViolationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultViolationReport

> *os.File GetDefaultViolationReport(ctx, reportResultId).Execute()

Download violation report



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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyAPI.GetDefaultViolationReport(context.Background(), reportResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetDefaultViolationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultViolationReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetDefaultViolationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportResultId** | **string** | The ID of the report reference to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultViolationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodAllReportRunStatus

> ReportResultReference GetSodAllReportRunStatus(ctx).Execute()

Get multi-report run task status



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
    resp, r, err := apiClient.SODPolicyAPI.GetSodAllReportRunStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetSodAllReportRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodAllReportRunStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetSodAllReportRunStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodAllReportRunStatusRequest struct via the builder pattern


### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodViolationReportRunStatus

> ReportResultReference GetSodViolationReportRunStatus(ctx, reportResultId).Execute()

Get violation report run status



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
    reportResultId := "2e8d8180-24bc-4d21-91c6-7affdb473b0d" // string | The ID of the report reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyAPI.GetSodViolationReportRunStatus(context.Background(), reportResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetSodViolationReportRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodViolationReportRunStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetSodViolationReportRunStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportResultId** | **string** | The ID of the report reference to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodViolationReportRunStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartSodAllPoliciesForOrg

> ReportResultReference StartSodAllPoliciesForOrg(ctx).MultiPolicyRequest(multiPolicyRequest).Execute()

Runs all policies for org



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
    multiPolicyRequest := *openapiclient.NewMultiPolicyRequest() // MultiPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyAPI.StartSodAllPoliciesForOrg(context.Background()).MultiPolicyRequest(multiPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.StartSodAllPoliciesForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartSodAllPoliciesForOrg`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.StartSodAllPoliciesForOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartSodAllPoliciesForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiPolicyRequest** | [**MultiPolicyRequest**](MultiPolicyRequest.md) |  | 

### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

