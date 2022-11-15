# \WorkItemsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveApprovalItem**](WorkItemsApi.md#ApproveApprovalItem) | **Post** /work-items/{id}/approve/{approvalItemId} | Approve an Approval Item
[**BulkApproveApprovalItem**](WorkItemsApi.md#BulkApproveApprovalItem) | **Post** /work-items/bulk-approve/{id} | Bulk approve Approval Items
[**BulkRejectApprovalItem**](WorkItemsApi.md#BulkRejectApprovalItem) | **Post** /work-items/bulk-reject/{id} | Bulk reject Approval Items
[**CompleteWorkItem**](WorkItemsApi.md#CompleteWorkItem) | **Post** /work-items/{id} | Complete a Work Item
[**CompletedWorkItems**](WorkItemsApi.md#CompletedWorkItems) | **Get** /work-items/completed | Completed Work Items
[**CountCompletedWorkItems**](WorkItemsApi.md#CountCompletedWorkItems) | **Get** /work-items/completed/count | Count Completed Work Items
[**CountWorkItems**](WorkItemsApi.md#CountWorkItems) | **Get** /work-items/count | Count Work Items
[**GetWorkItems**](WorkItemsApi.md#GetWorkItems) | **Get** /work-items/{id} | Get a Work Item
[**ListWorkItems**](WorkItemsApi.md#ListWorkItems) | **Get** /work-items | List Work Items
[**RejectApprovalItem**](WorkItemsApi.md#RejectApprovalItem) | **Post** /work-items/{id}/reject/{approvalItemId} | Reject an Approval Item
[**SubmitAccountSelection**](WorkItemsApi.md#SubmitAccountSelection) | **Post** /work-items/{id}/submit-account-selection | Submit Account Selections
[**SummaryWorkItems**](WorkItemsApi.md#SummaryWorkItems) | **Get** /work-items/summary | Work Items Summary



## ApproveApprovalItem

> WorkItems ApproveApprovalItem(ctx, id, approvalItemId).Execute()

Approve an Approval Item



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.ApproveApprovalItem(context.Background(), id, approvalItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.ApproveApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.ApproveApprovalItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 
**approvalItemId** | **string** | The ID of the approval item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveApprovalItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkApproveApprovalItem

> WorkItems BulkApproveApprovalItem(ctx, id).Execute()

Bulk approve Approval Items



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.BulkApproveApprovalItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.BulkApproveApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkApproveApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.BulkApproveApprovalItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkApproveApprovalItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRejectApprovalItem

> WorkItems BulkRejectApprovalItem(ctx, id).Execute()

Bulk reject Approval Items



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.BulkRejectApprovalItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.BulkRejectApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkRejectApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.BulkRejectApprovalItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkRejectApprovalItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteWorkItem

> WorkItems CompleteWorkItem(ctx, id).Execute()

Complete a Work Item



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.CompleteWorkItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.CompleteWorkItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteWorkItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.CompleteWorkItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompletedWorkItems

> []WorkItems CompletedWorkItems(ctx).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()

Completed Work Items



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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | The id of the owner of the work item list being requested.  Either an admin, or the owning/current user must make this request. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.CompletedWorkItems(context.Background()).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.CompletedWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompletedWorkItems`: []WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.CompletedWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompletedWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | The id of the owner of the work item list being requested.  Either an admin, or the owning/current user must make this request. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountCompletedWorkItems

> WorkItemsCount CountCompletedWorkItems(ctx).OwnerId(ownerId).Execute()

Count Completed Work Items



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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.CountCompletedWorkItems(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.CountCompletedWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountCompletedWorkItems`: WorkItemsCount
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.CountCompletedWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountCompletedWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsCount**](WorkItemsCount.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountWorkItems

> WorkItemsCount CountWorkItems(ctx).OwnerId(ownerId).Execute()

Count Work Items



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
    ownerId := "ef38f94347e94562b5bb8424a56397d8" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.CountWorkItems(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.CountWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountWorkItems`: WorkItemsCount
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.CountWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsCount**](WorkItemsCount.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkItems

> WorkItems GetWorkItems(ctx, id).Execute()

Get a Work Item



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
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the work item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.GetWorkItems(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.GetWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItems`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.GetWorkItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the work item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkItems

> []WorkItems ListWorkItems(ctx).Limit(limit).Offset(offset).Count(count).OwnerId(ownerId).Execute()

List Work Items



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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.ListWorkItems(context.Background()).Limit(limit).Offset(offset).Count(count).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.ListWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkItems`: []WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.ListWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**[]WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectApprovalItem

> WorkItems RejectApprovalItem(ctx, id, approvalItemId).Execute()

Reject an Approval Item



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.RejectApprovalItem(context.Background(), id, approvalItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.RejectApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.RejectApprovalItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 
**approvalItemId** | **string** | The ID of the approval item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectApprovalItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitAccountSelection

> WorkItems SubmitAccountSelection(ctx, id).RequestBody(requestBody).Execute()

Submit Account Selections



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Account Selection Data map, keyed on fieldName

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.SubmitAccountSelection(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.SubmitAccountSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAccountSelection`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.SubmitAccountSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAccountSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** | Account Selection Data map, keyed on fieldName | 

### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummaryWorkItems

> WorkItemsSummary SummaryWorkItems(ctx).OwnerId(ownerId).Execute()

Work Items Summary



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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsApi.SummaryWorkItems(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsApi.SummaryWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SummaryWorkItems`: WorkItemsSummary
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsApi.SummaryWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSummaryWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsSummary**](WorkItemsSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

