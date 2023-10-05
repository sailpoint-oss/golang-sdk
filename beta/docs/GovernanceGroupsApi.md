# \GovernanceGroupsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkgroup**](GovernanceGroupsApi.md#CreateWorkgroup) | **Post** /workgroups | Create a new Governance Group.
[**DeleteWorkgroup**](GovernanceGroupsApi.md#DeleteWorkgroup) | **Delete** /workgroups/{id} | Delete a Governance Group
[**DeleteWorkgroupMembers**](GovernanceGroupsApi.md#DeleteWorkgroupMembers) | **Post** /workgroups/{workgrouId}/members/bulk-delete | Remove members from Governance Group
[**DeleteWorkgroupsInBulk**](GovernanceGroupsApi.md#DeleteWorkgroupsInBulk) | **Post** /workgroups/bulk-delete | Delete Governance Group(s)
[**GetWorkgroup**](GovernanceGroupsApi.md#GetWorkgroup) | **Get** /workgroups/{id} | Get Governance Group by Id
[**ListConnections**](GovernanceGroupsApi.md#ListConnections) | **Get** /workgroups/{workgrouId}/connections | List connections for Governance Group
[**ListWorkgroupMembers**](GovernanceGroupsApi.md#ListWorkgroupMembers) | **Get** /workgroups/{workgrouId}/members | List Governance Group Members
[**ListWorkgroups**](GovernanceGroupsApi.md#ListWorkgroups) | **Get** /workgroups | List Governance Groups
[**PatchWorkgroup**](GovernanceGroupsApi.md#PatchWorkgroup) | **Patch** /workgroups/{id} | Patch a Governance Group
[**UpdateWorkgroupMembers**](GovernanceGroupsApi.md#UpdateWorkgroupMembers) | **Post** /workgroups/{workgrouId}/members/bulk-add | Add members to Governance Group



## CreateWorkgroup

> WorkgroupDto CreateWorkgroup(ctx).WorkgroupDto(workgroupDto).Execute()

Create a new Governance Group.



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
    workgroupDto := *openapiclient.NewWorkgroupDto() // WorkgroupDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.CreateWorkgroup(context.Background()).WorkgroupDto(workgroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.CreateWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkgroup`: WorkgroupDto
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.CreateWorkgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workgroupDto** | [**WorkgroupDto**](WorkgroupDto.md) |  | 

### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroup

> DeleteWorkgroup(ctx, id).Execute()

Delete a Governance Group



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
    id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.DeleteWorkgroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.DeleteWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Governance Group | 

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


## DeleteWorkgroupMembers

> []WorkgroupMemberDeleteItem DeleteWorkgroupMembers(ctx, workgroupId).BaseReferenceDto1(baseReferenceDto1).Execute()

Remove members from Governance Group



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
    workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
    baseReferenceDto1 := []openapiclient.BaseReferenceDto1{*openapiclient.NewBaseReferenceDto1()} // []BaseReferenceDto1 | List of identities to be removed from  a Governance Group members list.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.DeleteWorkgroupMembers(context.Background(), workgroupId).BaseReferenceDto1(baseReferenceDto1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.DeleteWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkgroupMembers`: []WorkgroupMemberDeleteItem
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.DeleteWorkgroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | ID of the Governance Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseReferenceDto1** | [**[]BaseReferenceDto1**](BaseReferenceDto1.md) | List of identities to be removed from  a Governance Group members list. | 

### Return type

[**[]WorkgroupMemberDeleteItem**](WorkgroupMemberDeleteItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroupsInBulk

> []WorkgroupDeleteItem DeleteWorkgroupsInBulk(ctx).WorkgroupBulkDeleteRequest(workgroupBulkDeleteRequest).Execute()

Delete Governance Group(s)



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
    workgroupBulkDeleteRequest := *openapiclient.NewWorkgroupBulkDeleteRequest() // WorkgroupBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.DeleteWorkgroupsInBulk(context.Background()).WorkgroupBulkDeleteRequest(workgroupBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.DeleteWorkgroupsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkgroupsInBulk`: []WorkgroupDeleteItem
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.DeleteWorkgroupsInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workgroupBulkDeleteRequest** | [**WorkgroupBulkDeleteRequest**](WorkgroupBulkDeleteRequest.md) |  | 

### Return type

[**[]WorkgroupDeleteItem**](WorkgroupDeleteItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkgroup

> WorkgroupDto GetWorkgroup(ctx, id).Execute()

Get Governance Group by Id



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
    id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.GetWorkgroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.GetWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkgroup`: WorkgroupDto
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.GetWorkgroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Governance Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> []WorkgroupConnectionDto ListConnections(ctx, workgroupId).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()

List connections for Governance Group



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
    workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.ListConnections(context.Background(), workgroupId).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: []WorkgroupConnectionDto
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | ID of the Governance Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** | 

### Return type

[**[]WorkgroupConnectionDto**](WorkgroupConnectionDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupMembers

> []BaseReferenceDto1 ListWorkgroupMembers(ctx, workgroupId).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()

List Governance Group Members



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
    workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.ListWorkgroupMembers(context.Background(), workgroupId).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroupMembers`: []BaseReferenceDto1
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListWorkgroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | ID of the Governance Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** | 

### Return type

[**[]BaseReferenceDto1**](BaseReferenceDto1.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroups

> []WorkgroupDto ListWorkgroups(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()

List Governance Groups



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name sw "Test"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Governance Group fields and operators: **id**: *eq, in*  **name**: *eq, sw* (optional)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.ListWorkgroups(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.ListWorkgroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkgroups`: []WorkgroupDto
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.ListWorkgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Governance Group fields and operators: **id**: *eq, in*  **name**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created, modified** | 

### Return type

[**[]WorkgroupDto**](WorkgroupDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkgroup

> WorkgroupDto PatchWorkgroup(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a Governance Group



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
    id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.PatchWorkgroup(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.PatchWorkgroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchWorkgroup`: WorkgroupDto
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.PatchWorkgroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Governance Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkgroupMembers

> []WorkgroupMemberAddItem UpdateWorkgroupMembers(ctx, workgroupId).BaseReferenceDto1(baseReferenceDto1).Execute()

Add members to Governance Group



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
    workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
    baseReferenceDto1 := []openapiclient.BaseReferenceDto1{*openapiclient.NewBaseReferenceDto1()} // []BaseReferenceDto1 | List of identities to be added to a Governance Group members list.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GovernanceGroupsApi.UpdateWorkgroupMembers(context.Background(), workgroupId).BaseReferenceDto1(baseReferenceDto1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsApi.UpdateWorkgroupMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkgroupMembers`: []WorkgroupMemberAddItem
    fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsApi.UpdateWorkgroupMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workgroupId** | **string** | ID of the Governance Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseReferenceDto1** | [**[]BaseReferenceDto1**](BaseReferenceDto1.md) | List of identities to be added to a Governance Group members list. | 

### Return type

[**[]WorkgroupMemberAddItem**](WorkgroupMemberAddItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

