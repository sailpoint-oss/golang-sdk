# \SODPolicyApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSodPolicy**](SODPolicyApi.md#CreateSodPolicy) | **Post** /sod-policies | Create SOD Policy
[**DeleteSodPolicyById**](SODPolicyApi.md#DeleteSodPolicyById) | **Delete** /sod-policies/{id} | Delete SOD Policy by ID
[**DeleteSodPolicyScheduleById**](SODPolicyApi.md#DeleteSodPolicyScheduleById) | **Delete** /sod-policies/{id}/schedule | Delete SOD Policy Schedule
[**DownloadCustomViolationReport**](SODPolicyApi.md#DownloadCustomViolationReport) | **Get** /sod-violation-report/{reportResultId}/download/{fileName} | Download custom violation report
[**DownloadDefaultViolationReport**](SODPolicyApi.md#DownloadDefaultViolationReport) | **Get** /sod-violation-report/{reportResultId}/download | Download violation report
[**GetSodAllReportRunStatus**](SODPolicyApi.md#GetSodAllReportRunStatus) | **Get** /sod-violation-report | Get multi-report run task status
[**GetSodPolicyById**](SODPolicyApi.md#GetSodPolicyById) | **Get** /sod-policies/{id} | Get SOD Policy By ID
[**GetSodPolicyScheduleById**](SODPolicyApi.md#GetSodPolicyScheduleById) | **Get** /sod-policies/{id}/schedule | Get SOD Policy Schedule
[**GetSodViolationReportRunStatus**](SODPolicyApi.md#GetSodViolationReportRunStatus) | **Get** /sod-violation-report-status/{reportResultId} | Get violation report run status
[**GetSodViolationReportStatus**](SODPolicyApi.md#GetSodViolationReportStatus) | **Get** /sod-policies/{id}/violation-report | Get SOD violation report status
[**ListSodPolicies**](SODPolicyApi.md#ListSodPolicies) | **Get** /sod-policies | List SOD Policies
[**PatchSodPolicy**](SODPolicyApi.md#PatchSodPolicy) | **Patch** /sod-policies/{id} | Update a SOD Policy
[**RunAllPoliciesForOrg**](SODPolicyApi.md#RunAllPoliciesForOrg) | **Post** /sod-violation-report/run | Runs all policies for Org.
[**RunSodPolicy**](SODPolicyApi.md#RunSodPolicy) | **Post** /sod-policies/{id}/violation-report/run | Runs SOD Policy Violation Report
[**UpdatePolicyById**](SODPolicyApi.md#UpdatePolicyById) | **Put** /sod-policies/{id} | Update SOD Policy By ID
[**UpdatePolicyScheduleById**](SODPolicyApi.md#UpdatePolicyScheduleById) | **Put** /sod-policies/{id}/schedule | Update SOD Policy schedule



## CreateSodPolicy

> SodPolicy CreateSodPolicy(ctx).SodPolicy(sodPolicy).Execute()

Create SOD Policy



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
    sodPolicy := *openapiclient.NewSodPolicy() // SodPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.CreateSodPolicy(context.Background()).SodPolicy(sodPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.CreateSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.CreateSodPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sodPolicy** | [**SodPolicy**](SodPolicy.md) |  | 

### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSodPolicyById

> DeleteSodPolicyById(ctx, id).Logical(logical).Execute()

Delete SOD Policy by ID



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the SOD Policy to delete.
    logical := true // bool | whether this is soft delete i.e. logical true or hard delete (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DeleteSodPolicyById(context.Background(), id).Logical(logical).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DeleteSodPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD Policy to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSodPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logical** | **bool** | whether this is soft delete i.e. logical true or hard delete | [default to true]

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSodPolicyScheduleById

> DeleteSodPolicyScheduleById(ctx, id).Execute()

Delete SOD Policy Schedule



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the SOD Policy for which the schedule needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DeleteSodPolicyScheduleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DeleteSodPolicyScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD Policy for which the schedule needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSodPolicyScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCustomViolationReport

> *os.File DownloadCustomViolationReport(ctx, reportResultId, fileName).Execute()

Download custom violation report



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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.
    fileName := "custom-name" // string | Custom Name for the  file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DownloadCustomViolationReport(context.Background(), reportResultId, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DownloadCustomViolationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadCustomViolationReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.DownloadCustomViolationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportResultId** | **string** | The ID of the report reference to download. | 
**fileName** | **string** | Custom Name for the  file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCustomViolationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDefaultViolationReport

> *os.File DownloadDefaultViolationReport(ctx, reportResultId).Execute()

Download violation report



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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DownloadDefaultViolationReport(context.Background(), reportResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DownloadDefaultViolationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadDefaultViolationReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.DownloadDefaultViolationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportResultId** | **string** | The ID of the report reference to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDefaultViolationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodAllReportRunStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodAllReportRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodAllReportRunStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodAllReportRunStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodAllReportRunStatusRequest struct via the builder pattern


### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodPolicyById

> SodPolicy GetSodPolicyById(ctx, id).Execute()

Get SOD Policy By ID



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodPolicyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicyById`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object reference to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodPolicyScheduleById

> SodPolicySchedule GetSodPolicyScheduleById(ctx, id).Execute()

Get SOD Policy Schedule



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodPolicyScheduleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodPolicyScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicyScheduleById`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodPolicyScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object reference to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodPolicyScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SodPolicySchedule**](SodPolicySchedule.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    reportResultId := "2e8d8180-24bc-4d21-91c6-7affdb473b0d" // string | The ID of the report reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodViolationReportRunStatus(context.Background(), reportResultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodViolationReportRunStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodViolationReportRunStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodViolationReportRunStatus`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodViolationReportStatus

> ReportResultReference GetSodViolationReportStatus(ctx, id).Execute()

Get SOD violation report status



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodViolationReportStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodViolationReportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodViolationReportStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodViolationReportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the object reference to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodViolationReportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSodPolicies

> []SodPolicy ListSodPolicies(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

List SOD Policies



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "bc693f07e7b645539626c25954c58554"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* <br/> **name**: *eq* <br/> **state**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.ListSodPolicies(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.ListSodPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSodPolicies`: []SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.ListSodPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSodPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* &lt;br/&gt; **name**: *eq* &lt;br/&gt; **state**: *eq* | 

### Return type

[**[]SodPolicy**](SodPolicy.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSodPolicy

> SodPolicy PatchSodPolicy(ctx, id).RequestBody(requestBody).Execute()

Update a SOD Policy



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
    id := "2c9180835d191a86015d28455b4a2329" // string | The ID of the SOD Policy being modified.
    requestBody := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.PatchSodPolicy(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.PatchSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.PatchSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD Policy being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria  | 

### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAllPoliciesForOrg

> ReportResultReference RunAllPoliciesForOrg(ctx).MultiPolicyRequest(multiPolicyRequest).Execute()

Runs all policies for Org.



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
    multiPolicyRequest := *openapiclient.NewMultiPolicyRequest() // MultiPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.RunAllPoliciesForOrg(context.Background()).MultiPolicyRequest(multiPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.RunAllPoliciesForOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunAllPoliciesForOrg`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.RunAllPoliciesForOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunAllPoliciesForOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiPolicyRequest** | [**MultiPolicyRequest**](MultiPolicyRequest.md) |  | 

### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunSodPolicy

> ReportResultReference RunSodPolicy(ctx, id).Execute()

Runs SOD Policy Violation Report



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the SOD Policy to run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.RunSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.RunSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunSodPolicy`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.RunSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD Policy to run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReportResultReference**](ReportResultReference.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyById

> SodPolicy UpdatePolicyById(ctx, id).SodPolicy(sodPolicy).Execute()

Update SOD Policy By ID



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the SOD policy to update.
    sodPolicy := *openapiclient.NewSodPolicy() // SodPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.UpdatePolicyById(context.Background(), id).SodPolicy(sodPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.UpdatePolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyById`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.UpdatePolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sodPolicy** | [**SodPolicy**](SodPolicy.md) |  | 

### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyScheduleById

> SodPolicySchedule UpdatePolicyScheduleById(ctx, id).SodPolicySchedule(sodPolicySchedule).Execute()

Update SOD Policy schedule



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the SOD policy to update its schedule.
    sodPolicySchedule := *openapiclient.NewSodPolicySchedule() // SodPolicySchedule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.UpdatePolicyScheduleById(context.Background(), id).SodPolicySchedule(sodPolicySchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.UpdatePolicyScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyScheduleById`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.UpdatePolicyScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update its schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sodPolicySchedule** | [**SodPolicySchedule**](SodPolicySchedule.md) |  | 

### Return type

[**SodPolicySchedule**](SodPolicySchedule.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

