# \SuggestedEntitlementDescriptionAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSedBatchStats**](SuggestedEntitlementDescriptionAPI.md#GetSedBatchStats) | **Get** /suggested-entitlement-description-batches/{batchId}/stats | Submit Sed Batch Stats Request
[**GetSedBatches**](SuggestedEntitlementDescriptionAPI.md#GetSedBatches) | **Get** /suggested-entitlement-description-batches | List Sed Batch Request
[**ListSeds**](SuggestedEntitlementDescriptionAPI.md#ListSeds) | **Get** /suggested-entitlement-descriptions | List Suggested Entitlement Description
[**PatchSed**](SuggestedEntitlementDescriptionAPI.md#PatchSed) | **Patch** /suggested-entitlement-descriptions | Patch Suggested Entitlement Description
[**SubmitSedApproval**](SuggestedEntitlementDescriptionAPI.md#SubmitSedApproval) | **Post** /suggested-entitlement-description-approvals | Submit Bulk Approval Request
[**SubmitSedAssignment**](SuggestedEntitlementDescriptionAPI.md#SubmitSedAssignment) | **Post** /suggested-entitlement-description-assignments | Submit Sed Assignment Request
[**SubmitSedBatchRequest**](SuggestedEntitlementDescriptionAPI.md#SubmitSedBatchRequest) | **Post** /suggested-entitlement-description-batches | Submit Sed Batch Request



## GetSedBatchStats

> SedBatchStats GetSedBatchStats(ctx, batchId).XSailPointExperimental(xSailPointExperimental).Execute()

Submit Sed Batch Stats Request



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
    batchId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | Batch Id
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.GetSedBatchStats(context.Background(), batchId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.GetSedBatchStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSedBatchStats`: SedBatchStats
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.GetSedBatchStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | Batch Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSedBatchStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SedBatchStats**](SedBatchStats.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSedBatches

> SedBatchStatus GetSedBatches(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

List Sed Batch Request



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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.GetSedBatches(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.GetSedBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSedBatches`: SedBatchStatus
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.GetSedBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSedBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SedBatchStatus**](SedBatchStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSeds

> []Sed ListSeds(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Filters(filters).Count(count).CountOnly(countOnly).RequestedByAnyone(requestedByAnyone).ShowPendingStatusOnly(showPendingStatusOnly).Execute()

List Suggested Entitlement Description



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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    limit := int64(limit=0) // int64 | Integer specifying the maximum number of records to return in a single API call.  The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional)
    filters := "displayName co "Read and Write"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **batchId**: *eq*  **status**: *eq, ne, in*  **displayName**: *eq, co* (optional)
    count := false // bool | If `true` it will populate the `X-Total-Count` response header with the number of results that would be returned if `limit` and `offset` were ignored.  The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). Since requesting a total count can have a performance impact, it is recommended not to send `count=true` if that value will not be used. (optional)
    countOnly := false // bool | If `true` it will populate the `X-Total-Count` response header with the number of results that would be returned if `limit` and `offset` were ignored. This parameter differs from the Coun parameter in that this one skip executing the actual query and always return an empty array. (optional)
    requestedByAnyone := false // bool | By default, the ListSeds API will only return items that you have requested to be generated.   This option will allow you to see all items that have been requested (optional)
    showPendingStatusOnly := false // bool | Will limit records to items that are in \"suggested\" or \"approved\" status (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.ListSeds(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Filters(filters).Count(count).CountOnly(countOnly).RequestedByAnyone(requestedByAnyone).ShowPendingStatusOnly(showPendingStatusOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.ListSeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSeds`: []Sed
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.ListSeds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSedsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int64** | Integer specifying the maximum number of records to return in a single API call.  The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **batchId**: *eq*  **status**: *eq, ne, in*  **displayName**: *eq, co* | 
 **count** | **bool** | If &#x60;true&#x60; it will populate the &#x60;X-Total-Count&#x60; response header with the number of results that would be returned if &#x60;limit&#x60; and &#x60;offset&#x60; were ignored.  The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). Since requesting a total count can have a performance impact, it is recommended not to send &#x60;count&#x3D;true&#x60; if that value will not be used. | 
 **countOnly** | **bool** | If &#x60;true&#x60; it will populate the &#x60;X-Total-Count&#x60; response header with the number of results that would be returned if &#x60;limit&#x60; and &#x60;offset&#x60; were ignored. This parameter differs from the Coun parameter in that this one skip executing the actual query and always return an empty array. | 
 **requestedByAnyone** | **bool** | By default, the ListSeds API will only return items that you have requested to be generated.   This option will allow you to see all items that have been requested | 
 **showPendingStatusOnly** | **bool** | Will limit records to items that are in \&quot;suggested\&quot; or \&quot;approved\&quot; status | 

### Return type

[**[]Sed**](Sed.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSed

> Sed PatchSed(ctx, id).XSailPointExperimental(xSailPointExperimental).SedPatch(sedPatch).Execute()

Patch Suggested Entitlement Description



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
    id := "ebab396f-0af1-4050-89b7-dafc63ec70e7" // string | id is sed id
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    sedPatch := []openapiclient.SedPatch{*openapiclient.NewSedPatch()} // []SedPatch | Sed Patch Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.PatchSed(context.Background(), id).XSailPointExperimental(xSailPointExperimental).SedPatch(sedPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.PatchSed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSed`: Sed
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.PatchSed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id is sed id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sedPatch** | [**[]SedPatch**](SedPatch.md) | Sed Patch Request | 

### Return type

[**Sed**](Sed.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitSedApproval

> []SedApprovalStatus SubmitSedApproval(ctx).XSailPointExperimental(xSailPointExperimental).SedApproval(sedApproval).Execute()

Submit Bulk Approval Request



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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    sedApproval := []openapiclient.SedApproval{*openapiclient.NewSedApproval()} // []SedApproval | Sed Approval

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.SubmitSedApproval(context.Background()).XSailPointExperimental(xSailPointExperimental).SedApproval(sedApproval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.SubmitSedApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitSedApproval`: []SedApprovalStatus
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.SubmitSedApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitSedApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sedApproval** | [**[]SedApproval**](SedApproval.md) | Sed Approval | 

### Return type

[**[]SedApprovalStatus**](SedApprovalStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitSedAssignment

> SedAssignmentResponse SubmitSedAssignment(ctx).XSailPointExperimental(xSailPointExperimental).SedAssignment(sedAssignment).Execute()

Submit Sed Assignment Request



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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    sedAssignment := *openapiclient.NewSedAssignment() // SedAssignment | Sed Assignment Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.SubmitSedAssignment(context.Background()).XSailPointExperimental(xSailPointExperimental).SedAssignment(sedAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.SubmitSedAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitSedAssignment`: SedAssignmentResponse
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.SubmitSedAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitSedAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sedAssignment** | [**SedAssignment**](SedAssignment.md) | Sed Assignment Request | 

### Return type

[**SedAssignmentResponse**](SedAssignmentResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitSedBatchRequest

> SedBatchResponse SubmitSedBatchRequest(ctx).XSailPointExperimental(xSailPointExperimental).SedBatchRequest(sedBatchRequest).Execute()

Submit Sed Batch Request



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
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
    sedBatchRequest := *openapiclient.NewSedBatchRequest() // SedBatchRequest | Sed Batch Request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SuggestedEntitlementDescriptionAPI.SubmitSedBatchRequest(context.Background()).XSailPointExperimental(xSailPointExperimental).SedBatchRequest(sedBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SuggestedEntitlementDescriptionAPI.SubmitSedBatchRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitSedBatchRequest`: SedBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `SuggestedEntitlementDescriptionAPI.SubmitSedBatchRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitSedBatchRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sedBatchRequest** | [**SedBatchRequest**](SedBatchRequest.md) | Sed Batch Request | 

### Return type

[**SedBatchResponse**](SedBatchResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

