# \SODPolicyApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSodPolicy**](SODPolicyApi.md#CreateSodPolicy) | **Post** /sod-policies | Create SOD policy
[**DeleteSodPolicy**](SODPolicyApi.md#DeleteSodPolicy) | **Delete** /sod-policies/{id} | Delete SOD policy by ID
[**DeleteSodPolicySchedule**](SODPolicyApi.md#DeleteSodPolicySchedule) | **Delete** /sod-policies/{id}/schedule | Delete SOD policy schedule
[**GetSodPolicy**](SODPolicyApi.md#GetSodPolicy) | **Get** /sod-policies/{id} | Get SOD policy by ID
[**GetSodPolicySchedule**](SODPolicyApi.md#GetSodPolicySchedule) | **Get** /sod-policies/{id}/schedule | Get SOD policy schedule
[**GetSodViolationReportRunStatus**](SODPolicyApi.md#GetSodViolationReportRunStatus) | **Get** /sod-violation-report-status/{reportResultId} | Get violation report run status
[**GetSodViolationReportStatus**](SODPolicyApi.md#GetSodViolationReportStatus) | **Get** /sod-policies/{id}/violation-report | Get SOD violation report status
[**ListSodPolicies**](SODPolicyApi.md#ListSodPolicies) | **Get** /sod-policies | List SOD policies
[**PatchSodPolicy**](SODPolicyApi.md#PatchSodPolicy) | **Patch** /sod-policies/{id} | Patch SOD policy by ID
[**SetPolicySchedule**](SODPolicyApi.md#SetPolicySchedule) | **Put** /sod-policies/{id}/schedule | Update SOD Policy schedule
[**SetSodPolicy**](SODPolicyApi.md#SetSodPolicy) | **Put** /sod-policies/{id} | Update SOD policy by ID
[**StartSodPolicy**](SODPolicyApi.md#StartSodPolicy) | **Post** /sod-policies/{id}/violation-report/run | Runs SOD policy violation report



## CreateSodPolicy

> SodPolicy CreateSodPolicy(ctx).SodPolicy(sodPolicy).Execute()

Create SOD policy



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


## DeleteSodPolicy

> DeleteSodPolicy(ctx, id).Logical(logical).Execute()

Delete SOD policy by ID



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD Policy to delete.
    logical := true // bool | Indicates whether this is a soft delete (logical true) or a hard delete.  Soft delete marks the policy as deleted and just save it with this status. It could be fully deleted or recovered further.  Hard delete vise versa permanently delete SOD request during this call. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DeleteSodPolicy(context.Background(), id).Logical(logical).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DeleteSodPolicy``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSodPolicySchedule

> DeleteSodPolicySchedule(ctx, id).Execute()

Delete SOD policy schedule



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy the schedule must be deleted for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.DeleteSodPolicySchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.DeleteSodPolicySchedule``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodPolicy

> SodPolicy GetSodPolicy(ctx, id).Execute()

Get SOD policy by ID



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD Policy to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodPolicy`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSodPolicySchedule

> SodPolicySchedule GetSodPolicySchedule(ctx, id).Execute()

Get SOD policy schedule



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy schedule to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.GetSodPolicySchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.GetSodPolicySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSodPolicySchedule`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.GetSodPolicySchedule`: %v\n", resp)
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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the violation report to retrieve status for.

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
**id** | **string** | The ID of the violation report to retrieve status for. | 

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

List SOD policies



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
    filters := "id eq "bc693f07e7b645539626c25954c58554"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **name**: *eq*  **state**: *eq* (optional)

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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **name**: *eq*  **state**: *eq* | 

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

> SodPolicy PatchSodPolicy(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch SOD policy by ID



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
    id := "2c918083-5d19-1a86-015d-28455b4a2329" // string | The ID of the SOD policy being modified.
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.PatchSodPolicy(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
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
**id** | **string** | The ID of the SOD policy being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSodPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of SOD Policy update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * ownerRef * externalPolicyReference * compensatingControls * correctionAdvice * state * tags * violationOwnerAssignmentConfig * scheduled * conflictingAccessCriteria  | 

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


## SetPolicySchedule

> SodPolicySchedule SetPolicySchedule(ctx, id).SodPolicySchedule(sodPolicySchedule).Execute()

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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy to update its schedule.
    sodPolicySchedule := *openapiclient.NewSodPolicySchedule() // SodPolicySchedule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.SetPolicySchedule(context.Background(), id).SodPolicySchedule(sodPolicySchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.SetPolicySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPolicySchedule`: SodPolicySchedule
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.SetPolicySchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update its schedule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPolicyScheduleRequest struct via the builder pattern


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


## SetSodPolicy

> SodPolicy SetSodPolicy(ctx, id).SodPolicy(sodPolicy).Execute()

Update SOD policy by ID



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The ID of the SOD policy to update.
    sodPolicy := *openapiclient.NewSodPolicy() // SodPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.SetSodPolicy(context.Background(), id).SodPolicy(sodPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.SetSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSodPolicy`: SodPolicy
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.SetSodPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the SOD policy to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSodPolicyRequest struct via the builder pattern


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


## StartSodPolicy

> ReportResultReference StartSodPolicy(ctx, id).Execute()

Runs SOD policy violation report



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
    id := "ef38f943-47e9-4562-b5bb-8424a56397d8" // string | The SOD policy ID to run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODPolicyApi.StartSodPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODPolicyApi.StartSodPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartSodPolicy`: ReportResultReference
    fmt.Fprintf(os.Stdout, "Response from `SODPolicyApi.StartSodPolicy`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

