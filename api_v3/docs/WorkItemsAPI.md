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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.ApproveApprovalItem(context.Background(), id, approvalItemId).Execute()
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


## ApproveApprovalItemsInBulk

> WorkItems ApproveApprovalItemsInBulk(ctx, id).Execute()

Bulk approve Approval Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.ApproveApprovalItemsInBulk(context.Background(), id).Execute()
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.CompleteWorkItem(context.Background(), id).Execute()
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


## GetCompletedWorkItems

> []WorkItems GetCompletedWorkItems(ctx).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()

Completed Work Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | The id of the owner of the work item list being requested.  Either an admin, or the owning/current user must make this request. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.GetCompletedWorkItems(context.Background()).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Execute()
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


## GetCountCompletedWorkItems

> WorkItemsCount GetCountCompletedWorkItems(ctx).OwnerId(ownerId).Execute()

Count Completed Work Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.GetCountCompletedWorkItems(context.Background()).OwnerId(ownerId).Execute()
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


## GetCountWorkItems

> WorkItemsCount GetCountWorkItems(ctx).OwnerId(ownerId).Execute()

Count Work Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    ownerId := "ef38f94347e94562b5bb8424a56397d8" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.GetCountWorkItems(context.Background()).OwnerId(ownerId).Execute()
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


## GetWorkItem

> WorkItems GetWorkItem(ctx, id).Execute()

Get a Work Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the work item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.GetWorkItem(context.Background(), id).Execute()
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


## GetWorkItemsSummary

> WorkItemsSummary GetWorkItemsSummary(ctx).OwnerId(ownerId).Execute()

Work Items Summary



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.GetWorkItemsSummary(context.Background()).OwnerId(ownerId).Execute()
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    ownerId := "1211bcaa32112bcef6122adb21cef1ac" // string | ID of the work item owner. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.ListWorkItems(context.Background()).Limit(limit).Offset(offset).Count(count).OwnerId(ownerId).Execute()
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    approvalItemId := "1211bcaa32112bcef6122adb21cef1ac" // string | The ID of the approval item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.RejectApprovalItem(context.Background(), id, approvalItemId).Execute()
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


## RejectApprovalItemsInBulk

> WorkItems RejectApprovalItemsInBulk(ctx, id).Execute()

Bulk reject Approval Items



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.RejectApprovalItemsInBulk(context.Background(), id).Execute()
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the work item
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Account Selection Data map, keyed on fieldName

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkItemsAPI.SubmitAccountSelection(context.Background(), id).RequestBody(requestBody).Execute()
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

