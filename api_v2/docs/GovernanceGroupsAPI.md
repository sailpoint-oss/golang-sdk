# \GovernanceGroupsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteWorkGroups**](GovernanceGroupsAPI.md#BulkDeleteWorkGroups) | **Post** /workgroups/bulk-delete | Bulk delete work groups
[**CreateWorkgroup**](GovernanceGroupsAPI.md#CreateWorkgroup) | **Post** /workgroups | Create Work Group
[**DeleteWorkgroup**](GovernanceGroupsAPI.md#DeleteWorkgroup) | **Delete** /workgroups/{workgroupId} | Delete Work Group By Id
[**GetWorkgroup**](GovernanceGroupsAPI.md#GetWorkgroup) | **Get** /workgroups/{workgroupId} | Get Work Group By Id
[**ListWorkgroupConnections**](GovernanceGroupsAPI.md#ListWorkgroupConnections) | **Get** /workgroups/{workgroupId}/connections | List Work Group Connections
[**ListWorkgroupMembers**](GovernanceGroupsAPI.md#ListWorkgroupMembers) | **Get** /workgroups/{workgroupId}/members | List Work Group Members
[**ListWorkgroups**](GovernanceGroupsAPI.md#ListWorkgroups) | **Get** /workgroups | List Work Groups
[**ModifyWorkgroupMembers**](GovernanceGroupsAPI.md#ModifyWorkgroupMembers) | **Post** /workgroups/{workgroupId}/members | Modify Work Group Members
[**UpdateWorkgroup**](GovernanceGroupsAPI.md#UpdateWorkgroup) | **Patch** /workgroups/{workgroupId} | Update Work Group By Id



## BulkDeleteWorkGroups

> BulkDeleteWorkGroups200Response BulkDeleteWorkGroups(ctx).BulkDeleteWorkGroupsRequest(bulkDeleteWorkGroupsRequest).Execute()

Bulk delete work groups



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
    bulkDeleteWorkGroupsRequest := *openapiclient.NewBulkDeleteWorkGroupsRequest() // BulkDeleteWorkGroupsRequest | Work group ids to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.BulkDeleteWorkGroups(context.Background()).BulkDeleteWorkGroupsRequest(bulkDeleteWorkGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.BulkDeleteWorkGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteWorkGroups`: BulkDeleteWorkGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.BulkDeleteWorkGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteWorkGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkDeleteWorkGroupsRequest** | [**BulkDeleteWorkGroupsRequest**](BulkDeleteWorkGroupsRequest.md) | Work group ids to delete | 

### Return type

[**BulkDeleteWorkGroups200Response**](BulkDeleteWorkGroups200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkgroup

> []ListWorkgroups200ResponseInner CreateWorkgroup(ctx).CreateWorkgroupRequest(createWorkgroupRequest).Execute()

Create Work Group



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
    createWorkgroupRequest := *openapiclient.NewCreateWorkgroupRequest() // CreateWorkgroupRequest | Work group to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.CreateWorkgroup(context.Background()).CreateWorkgroupRequest(createWorkgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.CreateWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkgroup`: []ListWorkgroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.CreateWorkgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkgroupRequest** | [**CreateWorkgroupRequest**](CreateWorkgroupRequest.md) | Work group to create. | 

### Return type

[**[]ListWorkgroups200ResponseInner**](ListWorkgroups200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroup

> DeleteWorkgroup(ctx, workgroupId).Execute()

Delete Work Group By Id



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GovernanceGroupsAPI.DeleteWorkgroup(context.Background(), workgroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.DeleteWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupRequest struct via the builder pattern


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


## GetWorkgroup

> ListWorkgroups200ResponseInner GetWorkgroup(ctx, workgroupId).Execute()

Get Work Group By Id



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.GetWorkgroup(context.Background(), workgroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.GetWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkgroup`: ListWorkgroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.GetWorkgroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListWorkgroups200ResponseInner**](ListWorkgroups200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupConnections

> []ListWorkgroupConnections200ResponseInner ListWorkgroupConnections(ctx, workgroupId).Execute()

List Work Group Connections



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.ListWorkgroupConnections(context.Background(), workgroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ListWorkgroupConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroupConnections`: []ListWorkgroupConnections200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.ListWorkgroupConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListWorkgroupConnections200ResponseInner**](ListWorkgroupConnections200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupMembers

> []ListWorkgroupMembers200ResponseInner ListWorkgroupMembers(ctx, workgroupId).Limit(limit).Offset(offset).Filters(filters).Execute()

List Work Group Members



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID
    limit := int32(250) // int32 | Max number of results to return (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. (optional) (default to 0)
    filters := "filters_example" // string | Filter results using the following syntax. [{property:name, value: \"Tyler\", operation: EQ}] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.ListWorkgroupMembers(context.Background(), workgroupId).Limit(limit).Offset(offset).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ListWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroupMembers`: []ListWorkgroupMembers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.ListWorkgroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. | [default to 0]
 **filters** | **string** | Filter results using the following syntax. [{property:name, value: \&quot;Tyler\&quot;, operation: EQ}] | 

### Return type

[**[]ListWorkgroupMembers200ResponseInner**](ListWorkgroupMembers200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroups

> []ListWorkgroups200ResponseInner ListWorkgroups(ctx).Limit(limit).Offset(offset).Filters(filters).Execute()

List Work Groups



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
    limit := int32(250) // int32 | Max number of results to return (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. (optional) (default to 0)
    filters := "filters_example" // string | Filter results using the following syntax. [{property:name, value: \"Tyler\", operation: EQ}] (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.ListWorkgroups(context.Background()).Limit(limit).Offset(offset).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ListWorkgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroups`: []ListWorkgroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.ListWorkgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. | [default to 0]
 **filters** | **string** | Filter results using the following syntax. [{property:name, value: \&quot;Tyler\&quot;, operation: EQ}] | 

### Return type

[**[]ListWorkgroups200ResponseInner**](ListWorkgroups200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWorkgroupMembers

> ModifyWorkgroupMembers(ctx, workgroupId).ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest).Execute()

Modify Work Group Members



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID
    modifyWorkgroupMembersRequest := *openapiclient.NewModifyWorkgroupMembersRequest() // ModifyWorkgroupMembersRequest | Add/Remove workgroup member ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GovernanceGroupsAPI.ModifyWorkgroupMembers(context.Background(), workgroupId).ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ModifyWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyWorkgroupMembersRequest** | [**ModifyWorkgroupMembersRequest**](ModifyWorkgroupMembersRequest.md) | Add/Remove workgroup member ids. | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkgroup

> ListWorkgroups200ResponseInner UpdateWorkgroup(ctx, workgroupId).CreateWorkgroupRequest(createWorkgroupRequest).Execute()

Update Work Group By Id



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
    workgroupId := "ef38f94347e94562b5bb8424a56397d8" // string | The workgroup ID
    createWorkgroupRequest := *openapiclient.NewCreateWorkgroupRequest() // CreateWorkgroupRequest | Work group to modify.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsAPI.UpdateWorkgroup(context.Background(), workgroupId).CreateWorkgroupRequest(createWorkgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.UpdateWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkgroup`: ListWorkgroups200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.UpdateWorkgroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | The workgroup ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkgroupRequest** | [**CreateWorkgroupRequest**](CreateWorkgroupRequest.md) | Work group to modify. | 

### Return type

[**ListWorkgroups200ResponseInner**](ListWorkgroups200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

