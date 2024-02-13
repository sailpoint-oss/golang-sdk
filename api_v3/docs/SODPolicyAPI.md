# \SODPolicyAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSodPolicy**](SODPolicyAPI.md#CreateSodPolicy) | **Post** /sod-policies | Create SOD policy
[**DeleteSodPolicy**](SODPolicyAPI.md#DeleteSodPolicy) | **Delete** /sod-policies/{id} | Delete SOD policy by ID
[**DeleteSodPolicySchedule**](SODPolicyAPI.md#DeleteSodPolicySchedule) | **Delete** /sod-policies/{id}/schedule | Delete SOD policy schedule
[**GetCustomViolationReport**](SODPolicyAPI.md#GetCustomViolationReport) | **Get** /sod-violation-report/{reportResultId}/download/{fileName} | Download custom violation report
[**GetDefaultViolationReport**](SODPolicyAPI.md#GetDefaultViolationReport) | **Get** /sod-violation-report/{reportResultId}/download | Download violation report
[**GetSodAllReportRunStatus**](SODPolicyAPI.md#GetSodAllReportRunStatus) | **Get** /sod-violation-report | Get multi-report run task status
[**GetSodPolicy**](SODPolicyAPI.md#GetSodPolicy) | **Get** /sod-policies/{id} | Get SOD policy by ID
[**GetSodPolicySchedule**](SODPolicyAPI.md#GetSodPolicySchedule) | **Get** /sod-policies/{id}/schedule | Get SOD policy schedule
[**GetSodViolationReportRunStatus**](SODPolicyAPI.md#GetSodViolationReportRunStatus) | **Get** /sod-policies/sod-violation-report-status/{reportResultId} | Get violation report run status
[**GetSodViolationReportStatus**](SODPolicyAPI.md#GetSodViolationReportStatus) | **Get** /sod-policies/{id}/violation-report | Get SOD violation report status
[**ListSodPolicies**](SODPolicyAPI.md#ListSodPolicies) | **Get** /sod-policies | List SOD policies
[**PatchSodPolicy**](SODPolicyAPI.md#PatchSodPolicy) | **Patch** /sod-policies/{id} | Patch SOD policy by ID
[**PutPolicySchedule**](SODPolicyAPI.md#PutPolicySchedule) | **Put** /sod-policies/{id}/schedule | Update SOD Policy schedule
[**PutSodPolicy**](SODPolicyAPI.md#PutSodPolicy) | **Put** /sod-policies/{id} | Update SOD policy by ID
[**StartEvaluateSodPolicy**](SODPolicyAPI.md#StartEvaluateSodPolicy) | **Post** /sod-policies/{id}/evaluate | Evaluate one policy by ID
[**StartSodAllPoliciesForOrg**](SODPolicyAPI.md#StartSodAllPoliciesForOrg) | **Post** /sod-violation-report/run | Runs all policies for org
[**StartSodPolicy**](SODPolicyAPI.md#StartSodPolicy) | **Post** /sod-policies/{id}/violation-report/run | Runs SOD policy violation report



## Create SOD policy

> SodPolicy CreateSodPolicy(ctx).SodPolicy(sodPolicy).Execute()

This creates both General and Conflicting Access Based policy, with a limit of 50 entitlements for each (left & right) criteria for Conflicting Access Based SOD policy.
Requires role of ORG_ADMIN.

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
    sodPolicy := *sailpoint.NewSodPolicy() // SodPolicy | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.CreateSodPolicy(context.Background()).SodPolicy(sodPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.CreateSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.CreateSodPolicy`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete SOD policy by ID

> DeleteSodPolicy(ctx, id).Logical(logical).Execute()

This deletes a specified SOD policy.
Requires role of ORG_ADMIN.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD Policy to delete.
    logical := true // bool | Indicates whether this is a soft delete (logical true) or a hard delete.  Soft delete marks the policy as deleted and just save it with this status. It could be fully deleted or recovered further.  Hard delete vise versa permanently delete SOD request during this call. (optional) (default to true)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SODPolicyAPI.DeleteSodPolicy(context.Background(), id).Logical(logical).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.DeleteSodPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logical** | **bool** | Indicates whether this is a soft delete (logical true) or a hard delete.  Soft delete marks the policy as deleted and just save it with this status. It could be fully deleted or recovered further.  Hard delete vise versa permanently delete SOD request during this call. | [default to true]

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


## Delete SOD policy schedule

> DeleteSodPolicySchedule(ctx, id).Execute()

This deletes schedule for a specified SOD policy by ID.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy the schedule must be deleted for.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.SODPolicyAPI.DeleteSodPolicySchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.DeleteSodPolicySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy the schedule must be deleted for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSodPolicyScheduleRequest struct via the builder pattern


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


## Download custom violation report

> *os.File GetCustomViolationReport(ctx, reportResultId, fileName).Execute()

This allows to download a specified named violation report for a given report reference.

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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.
    fileName := "custom-name" // string | Custom Name for the  file.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetCustomViolationReport(context.Background(), reportResultId, fileName).Execute()
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


## Download violation report

> *os.File GetDefaultViolationReport(ctx, reportResultId).Execute()

This allows to download a violation report for a given report reference.

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
    reportResultId := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the report reference to download.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetDefaultViolationReport(context.Background(), reportResultId).Execute()
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


## Get multi-report run task status

> ReportResultReference GetSodAllReportRunStatus(ctx).Execute()

This endpoint gets the status for a violation report for all policy run.

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

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetSodAllReportRunStatus(context.Background()).Execute()
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


## Get SOD policy by ID

> SodPolicy GetSodPolicy(ctx, id).Execute()

This gets specified SOD policy.
Requires role of ORG_ADMIN.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD Policy to retrieve.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD Policy to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get SOD policy schedule

> SodPolicySchedule GetSodPolicySchedule(ctx, id).Execute()

This endpoint gets a specified SOD policy's schedule.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy schedule to retrieve.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetSodPolicySchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetSodPolicySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicySchedule`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetSodPolicySchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy schedule to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodPolicyScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SodPolicySchedule**](SodPolicySchedule.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get violation report run status

> ReportResultReference GetSodViolationReportRunStatus(ctx, reportResultId).Execute()

This gets the status for a violation report run task that has already been invoked.

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
    reportResultId := "2e8d8180-24bc-4d21-91c6-7affdb473b0d" // string | The ID of the report reference to retrieve.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetSodViolationReportRunStatus(context.Background(), reportResultId).Execute()
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


## Get SOD violation report status

> ReportResultReference GetSodViolationReportStatus(ctx, id).Execute()

This gets the status for a violation report run task that has already been invoked.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the violation report to retrieve status for.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.GetSodViolationReportStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.GetSodViolationReportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodViolationReportStatus`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.GetSodViolationReportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the violation report to retrieve status for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSodViolationReportStatusRequest struct via the builder pattern


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


## List SOD policies

> []SodPolicy ListSodPolicies(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

This gets list of all SOD policies.
Requires role of ORG_ADMIN

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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "bc693f07e7b645539626c25954c58554"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **name**: *eq*  **state**: *eq* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.ListSodPolicies(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.ListSodPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSodPolicies`: []SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.ListSodPolicies`: %v\n", resp)
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **name**: *eq*  **state**: *eq* | 

### Return type

[**[]SodPolicy**](SodPolicy.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch SOD policy by ID

> SodPolicy PatchSodPolicy(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Allows updating SOD Policy fields other than ["id","created","creatorId","policyQuery","type"] using the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.
Requires role of ORG_ADMIN.
This endpoint can only patch CONFLICTING_ACCESS_BASED type policies. Do not use this endpoint to patch general policies - doing so will build an API exception. 

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
    id := "2c918083-5d19-1a86-015d-28455b4a2329" // string | The ID of the SOD policy being modified.
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.PatchSodPolicy(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.PatchSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.PatchSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria  | 

### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update SOD Policy schedule

> SodPolicySchedule PutPolicySchedule(ctx, id).SodPolicySchedule(sodPolicySchedule).Execute()

This updates schedule for a specified SOD policy.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy to update its schedule.
    sodPolicySchedule := *sailpoint.NewSodPolicySchedule() // SodPolicySchedule | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.PutPolicySchedule(context.Background(), id).SodPolicySchedule(sodPolicySchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.PutPolicySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPolicySchedule`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.PutPolicySchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update its schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPolicyScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sodPolicySchedule** | [**SodPolicySchedule**](SodPolicySchedule.md) |  | 

### Return type

[**SodPolicySchedule**](SodPolicySchedule.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update SOD policy by ID

> SodPolicy PutSodPolicy(ctx, id).SodPolicy(sodPolicy).Execute()

This updates a specified SOD policy.
Requires role of ORG_ADMIN.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy to update.
    sodPolicy := *sailpoint.NewSodPolicy() // SodPolicy | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.PutSodPolicy(context.Background(), id).SodPolicy(sodPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.PutSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.PutSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sodPolicy** | [**SodPolicy**](SodPolicy.md) |  | 

### Return type

[**SodPolicy**](SodPolicy.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Evaluate one policy by ID

> ReportResultReference StartEvaluateSodPolicy(ctx, id).Execute()

Runs the scheduled report for the policy retrieved by passed policy ID.  The report schedule is fetched from the policy retrieved by ID.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The SOD policy ID to run.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.StartEvaluateSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.StartEvaluateSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartEvaluateSodPolicy`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.StartEvaluateSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The SOD policy ID to run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartEvaluateSodPolicyRequest struct via the builder pattern


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


## Runs all policies for org

> ReportResultReference StartSodAllPoliciesForOrg(ctx).MultiPolicyRequest(multiPolicyRequest).Execute()

Runs multi-policy report for the org. If a policy reports more than 5000 violations, the report mentions that the violation limit was exceeded for that policy. If the request is empty, the report runs for all policies. Otherwise, the report runs for only the filtered policy list provided.

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
    multiPolicyRequest := *sailpoint.NewMultiPolicyRequest() // MultiPolicyRequest |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.StartSodAllPoliciesForOrg(context.Background()).MultiPolicyRequest(multiPolicyRequest).Execute()
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


## Runs SOD policy violation report

> ReportResultReference StartSodPolicy(ctx, id).Execute()

This invokes processing of violation report for given SOD policy. If the policy reports more than 5000 violations, the report returns with violation limit exceeded message.

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The SOD policy ID to run.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODPolicyAPI.StartSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyAPI.StartSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartSodPolicy`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyAPI.StartSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The SOD policy ID to run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartSodPolicyRequest struct via the builder pattern


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

