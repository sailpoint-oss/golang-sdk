# \GovernanceGroupsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteWorkGroups**](GovernanceGroupsApi.md#BulkDeleteWorkGroups) | **Post** /workgroups/bulk-delete | Bulk delete work groups
[**CreateWorkgroup**](GovernanceGroupsApi.md#CreateWorkgroup) | **Post** /workgroups | Create Work Group
[**DeleteWorkgroup**](GovernanceGroupsApi.md#DeleteWorkgroup) | **Delete** /v2/workgroups/{workgroupId} | Delete Work Group By Id
[**GetWorkgroup**](GovernanceGroupsApi.md#GetWorkgroup) | **Get** /v2/workgroups/{workgroupId} | Get Work Group By Id
[**ListWorkgroupConnections**](GovernanceGroupsApi.md#ListWorkgroupConnections) | **Get** /workgroups/{workgroupId}/connections | List Work Group Connections
[**ListWorkgroupMembers**](GovernanceGroupsApi.md#ListWorkgroupMembers) | **Get** /workgroups/{workgroupId}/members | List Work Group Members
[**ListWorkgroups**](GovernanceGroupsApi.md#ListWorkgroups) | **Get** /workgroups | List Work Groups
[**ModifyWorkgroupMembers**](GovernanceGroupsApi.md#ModifyWorkgroupMembers) | **Post** /workgroups/{workgroupId}/members | Modify Work Group Members
[**UpdateWorkgroup**](GovernanceGroupsApi.md#UpdateWorkgroup) | **Patch** /v2/workgroups/{workgroupId} | Update Work Group By Id



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
    openapiclient "./openapi"
)

func main() {
    bulkDeleteWorkGroupsRequest := *openapiclient.NewBulkDeleteWorkGroupsRequest() // BulkDeleteWorkGroupsRequest | Work group ids to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.BulkDeleteWorkGroups(context.Background()).BulkDeleteWorkGroupsRequest(bulkDeleteWorkGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.BulkDeleteWorkGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteWorkGroups`: BulkDeleteWorkGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.BulkDeleteWorkGroups`: %v\n", resp)
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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkgroup

> []Workgroup CreateWorkgroup(ctx).CreateWorkgroupRequest(createWorkgroupRequest).Execute()

Create Work Group



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
    createWorkgroupRequest := *openapiclient.NewCreateWorkgroupRequest() // CreateWorkgroupRequest | Work group to create.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.CreateWorkgroup(context.Background()).CreateWorkgroupRequest(createWorkgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.CreateWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkgroup`: []Workgroup
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.CreateWorkgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkgroupRequest** | [**CreateWorkgroupRequest**](CreateWorkgroupRequest.md) | Work group to create. | 

### Return type

[**[]Workgroup**](Workgroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroup

> DeleteWorkgroup(ctx).Execute()

Delete Work Group By Id



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
    resp, r, err := apiClient.GovernanceGroupsApi.DeleteWorkgroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.DeleteWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkgroup

> Workgroup GetWorkgroup(ctx).Execute()

Get Work Group By Id



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
    resp, r, err := apiClient.GovernanceGroupsApi.GetWorkgroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.GetWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkgroup`: Workgroup
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.GetWorkgroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkgroupRequest struct via the builder pattern


### Return type

[**Workgroup**](Workgroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupConnections

> []WorkgroupConnection ListWorkgroupConnections(ctx).Execute()

List Work Group Connections



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
    resp, r, err := apiClient.GovernanceGroupsApi.ListWorkgroupConnections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListWorkgroupConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroupConnections`: []WorkgroupConnection
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListWorkgroupConnections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupConnectionsRequest struct via the builder pattern


### Return type

[**[]WorkgroupConnection**](WorkgroupConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupMembers

> []WorkgroupMember ListWorkgroupMembers(ctx).Execute()

List Work Group Members



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
    resp, r, err := apiClient.GovernanceGroupsApi.ListWorkgroupMembers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroupMembers`: []WorkgroupMember
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListWorkgroupMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupMembersRequest struct via the builder pattern


### Return type

[**[]WorkgroupMember**](WorkgroupMember.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroups

> []Workgroup ListWorkgroups(ctx).Execute()

List Work Groups



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
    resp, r, err := apiClient.GovernanceGroupsApi.ListWorkgroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListWorkgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroups`: []Workgroup
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListWorkgroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupsRequest struct via the builder pattern


### Return type

[**[]Workgroup**](Workgroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWorkgroupMembers

> ModifyWorkgroupMembers(ctx).ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest).Execute()

Modify Work Group Members



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
    modifyWorkgroupMembersRequest := *openapiclient.NewModifyWorkgroupMembersRequest() // ModifyWorkgroupMembersRequest | Add/Remove workgroup member ids.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.ModifyWorkgroupMembers(context.Background()).ModifyWorkgroupMembersRequest(modifyWorkgroupMembersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ModifyWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modifyWorkgroupMembersRequest** | [**ModifyWorkgroupMembersRequest**](ModifyWorkgroupMembersRequest.md) | Add/Remove workgroup member ids. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkgroup

> Workgroup UpdateWorkgroup(ctx).CreateWorkgroupRequest(createWorkgroupRequest).Execute()

Update Work Group By Id



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
    createWorkgroupRequest := *openapiclient.NewCreateWorkgroupRequest() // CreateWorkgroupRequest | Work group to modify.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.UpdateWorkgroup(context.Background()).CreateWorkgroupRequest(createWorkgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.UpdateWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkgroup`: Workgroup
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.UpdateWorkgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkgroupRequest** | [**CreateWorkgroupRequest**](CreateWorkgroupRequest.md) | Work group to modify. | 

### Return type

[**Workgroup**](Workgroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

