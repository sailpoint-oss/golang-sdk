# \WorkItemsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveApprovalItem**](WorkItemsAPI.md#ApproveApprovalItem) | **Post** /work-items/{id}/approve/{approvalItemId} | Approve an Approval Item
[**ApproveApprovalItemsInBulk**](WorkItemsAPI.md#ApproveApprovalItemsInBulk) | **Post** /work-items/bulk-approve/{id} | Bulk approve Approval Items
[**CompleteWorkItem**](WorkItemsAPI.md#CompleteWorkItem) | **Post** /work-items/{id} | Complete a Work Item
[**GetCompletedWorkItems**](WorkItemsAPI.md#GetCompletedWorkItems) | **Get** /work-items/completed | Completed Work Items
[**GetCountCompletedWorkItems**](WorkItemsAPI.md#GetCountCompletedWorkItems) | **Get** /work-items/completed/count | Count Completed Work Items
[**GetCountWorkItems**](WorkItemsAPI.md#GetCountWorkItems) | **Get** /work-items/count | Count Work Items
[**GetWorkItem**](WorkItemsAPI.md#GetWorkItem) | **Get** /work-items/{id} | Get a Work Item
[**GetWorkItemsSummary**](WorkItemsAPI.md#GetWorkItemsSummary) | **Get** /work-items/summary | Work Items Summary
[**ListWorkItems**](WorkItemsAPI.md#ListWorkItems) | **Get** /work-items | List Work Items
[**RejectApprovalItem**](WorkItemsAPI.md#RejectApprovalItem) | **Post** /work-items/{id}/reject/{approvalItemId} | Reject an Approval Item
[**RejectApprovalItemsInBulk**](WorkItemsAPI.md#RejectApprovalItemsInBulk) | **Post** /work-items/bulk-reject/{id} | Bulk reject Approval Items
[**SubmitAccountSelection**](WorkItemsAPI.md#SubmitAccountSelection) | **Post** /work-items/{id}/submit-account-selection | Submit Account Selections



## Approve an Approval Item

> WorkItems ApproveApprovalItem(ctx, id, approvalItemId).Execute()

This API approves an Approval Item. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.ApproveApprovalItem(context.Background(), id, approvalItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApproveApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApproveApprovalItem`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulk approve Approval Items

> WorkItems ApproveApprovalItemsInBulk(ctx, id).Execute()

This API bulk approves Approval Items. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.ApproveApprovalItemsInBulk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ApproveApprovalItemsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveApprovalItemsInBulk`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ApproveApprovalItemsInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveApprovalItemsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Complete a Work Item

> WorkItems CompleteWorkItem(ctx, id).Execute()

This API completes a work item. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.CompleteWorkItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.CompleteWorkItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteWorkItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.CompleteWorkItem`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Completed Work Items

> []WorkItems GetCompletedWorkItems(ctx).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()

This gets a collection of completed work items belonging to either the specified user(admin required), or the current user.

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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | The id of the owner of the work item list being requested.  Either an admin, or the owning/current user must make this request. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.GetCompletedWorkItems(context.Background()).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetCompletedWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompletedWorkItems`: []WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetCompletedWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompletedWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | The id of the owner of the work item list being requested.  Either an admin, or the owning/current user must make this request. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]WorkItems**](WorkItems.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Count Completed Work Items

> WorkItemsCount GetCountCompletedWorkItems(ctx).OwnerId(ownerId).Execute()

This gets a count of completed work items belonging to either the specified user(admin required), or the current user.

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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.GetCountCompletedWorkItems(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetCountCompletedWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountCompletedWorkItems`: WorkItemsCount
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetCountCompletedWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountCompletedWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsCount**](WorkItemsCount.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Count Work Items

> WorkItemsCount GetCountWorkItems(ctx).OwnerId(ownerId).Execute()

This gets a count of work items belonging to either the specified user(admin required), or the current user.

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
    ownerId := "ef38f94347e94562b5bb8424a56397d8" // string | ID of the work item owner. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.GetCountWorkItems(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetCountWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountWorkItems`: WorkItemsCount
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetCountWorkItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCountWorkItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsCount**](WorkItemsCount.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get a Work Item

> WorkItems GetWorkItem(ctx, id).Execute()

This gets the details of a Work Item belonging to either the specified user(admin required), or the current user.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the work item.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.GetWorkItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetWorkItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetWorkItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the work item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Work Items Summary

> WorkItemsSummary GetWorkItemsSummary(ctx).OwnerId(ownerId).Execute()

This gets a summary of work items belonging to either the specified user(admin required), or the current user.

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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.GetWorkItemsSummary(context.Background()).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.GetWorkItemsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkItemsSummary`: WorkItemsSummary
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.GetWorkItemsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkItemsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | ID of the work item owner. | 

### Return type

[**WorkItemsSummary**](WorkItemsSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Work Items

> []WorkItems ListWorkItems(ctx).Limit(limit).Offset(offset).Count(count).OwnerId(ownerId).Execute()

This gets a collection of work items belonging to either the specified user(admin required), or the current user.

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
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.ListWorkItems(context.Background()).Limit(limit).Offset(offset).Count(count).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.ListWorkItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkItems`: []WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.ListWorkItems`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reject an Approval Item

> WorkItems RejectApprovalItem(ctx, id, approvalItemId).Execute()

This API rejects an Approval Item. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.RejectApprovalItem(context.Background(), id, approvalItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.RejectApprovalItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectApprovalItem`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.RejectApprovalItem`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Bulk reject Approval Items

> WorkItems RejectApprovalItemsInBulk(ctx, id).Execute()

This API bulk rejects Approval Items. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.RejectApprovalItemsInBulk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.RejectApprovalItemsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectApprovalItemsInBulk`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.RejectApprovalItemsInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the work item | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectApprovalItemsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkItems**](WorkItems.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Submit Account Selections

> WorkItems SubmitAccountSelection(ctx, id).RequestBody(requestBody).Execute()

This API submits account selections. Either an admin, or the owning/current user must make this request.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Account Selection Data map, keyed on fieldName

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.WorkItemsAPI.SubmitAccountSelection(context.Background(), id).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkItemsAPI.SubmitAccountSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAccountSelection`: WorkItems
    fmt.Fprintf(os.Stdout, "Response from `WorkItemsAPI.SubmitAccountSelection`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

