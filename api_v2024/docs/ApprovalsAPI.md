# \ApprovalsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApproval**](ApprovalsAPI.md#GetApproval) | **Get** /generic-approvals/{id} | Get an approval
[**GetApprovals**](ApprovalsAPI.md#GetApprovals) | **Get** /generic-approvals | Get Approvals
[**PatchApproval**](ApprovalsAPI.md#PatchApproval) | **Patch** /generic-approvals/{id} | Change an approval



## GetApproval

> Approval GetApproval(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Get an approval



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
    id := "38453251-6be2-5f8f-df93-5ce19e295837" // string | ID of the approval that is to be returned
    xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsAPI.GetApproval(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.GetApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApproval`: Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.GetApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the approval that is to be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Approval**](Approval.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovals

> []Approval GetApprovals(ctx).XSailPointExperimental(xSailPointExperimental).Mine(mine).RequesterId(requesterId).Filters(filters).Execute()

Get Approvals



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
    mine := true // bool | Returns the list of approvals for the current caller (optional)
    requesterId := "17e633e7d57e481569df76323169deb6a" // string | Returns the list of approvals for a given requester ID (optional)
    filters := "filters=status eq PENDING" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq*  **referenceType**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsAPI.GetApprovals(context.Background()).XSailPointExperimental(xSailPointExperimental).Mine(mine).RequesterId(requesterId).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.GetApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApprovals`: []Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.GetApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **mine** | **bool** | Returns the list of approvals for the current caller | 
 **requesterId** | **string** | Returns the list of approvals for a given requester ID | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq*  **referenceType**: *eq* | 

### Return type

[**[]Approval**](Approval.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApproval

> Approval PatchApproval(ctx).XSailPointExperimental(xSailPointExperimental).ApprovalDto(approvalDto).Execute()

Change an approval



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
    approvalDto := *openapiclient.NewApprovalDto() // ApprovalDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsAPI.PatchApproval(context.Background()).XSailPointExperimental(xSailPointExperimental).ApprovalDto(approvalDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsAPI.PatchApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchApproval`: Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsAPI.PatchApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **approvalDto** | [**ApprovalDto**](ApprovalDto.md) |  | 

### Return type

[**Approval**](Approval.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

