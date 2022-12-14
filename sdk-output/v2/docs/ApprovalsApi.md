# \ApprovalsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveAccessRequest**](ApprovalsApi.md#ApproveAccessRequest) | **Post** /approvals/{approvalId}/approve-request | Approves an access request.
[**ForwardAccessRequestApproval**](ApprovalsApi.md#ForwardAccessRequestApproval) | **Post** /approvals/{approvalId}/forward | Forwards an access request approval.
[**ListAccessRequestApproval**](ApprovalsApi.md#ListAccessRequestApproval) | **Get** /approvals | Lists the approvals.
[**RejectAccessRequest**](ApprovalsApi.md#RejectAccessRequest) | **Post** /approvals/{approvalId}/reject-request | Rejects an access request.



## ApproveAccessRequest

> Approval ApproveAccessRequest(ctx, approvalId).CommentEto(commentEto).Execute()

Approves an access request.



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
    approvalId := "approvalId_example" // string | ID of an access request approval.
    commentEto := *openapiclient.NewCommentEto() // CommentEto | Reviewer's comment. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.ApproveAccessRequest(context.Background(), approvalId).CommentEto(commentEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.ApproveAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveAccessRequest`: Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.ApproveAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | ID of an access request approval. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentEto** | [**CommentEto**](CommentEto.md) | Reviewer&#39;s comment. | 

### Return type

[**Approval**](Approval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForwardAccessRequestApproval

> Approval ForwardAccessRequestApproval(ctx, approvalId).ForwardApprovalEto(forwardApprovalEto).Execute()

Forwards an access request approval.



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
    approvalId := "approvalId_example" // string | ID of an access request approval.
    forwardApprovalEto := *openapiclient.NewForwardApprovalEto("NewOwnerId_example") // ForwardApprovalEto | Information about the forwarded approval.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.ForwardAccessRequestApproval(context.Background(), approvalId).ForwardApprovalEto(forwardApprovalEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.ForwardAccessRequestApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ForwardAccessRequestApproval`: Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.ForwardAccessRequestApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | ID of an access request approval. | 

### Other Parameters

Other parameters are passed through a pointer to a apiForwardAccessRequestApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forwardApprovalEto** | [**ForwardApprovalEto**](ForwardApprovalEto.md) | Information about the forwarded approval. | 

### Return type

[**Approval**](Approval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessRequestApproval

> []Approval ListAccessRequestApproval(ctx).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists the approvals.



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
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.ListAccessRequestApproval(context.Background()).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.ListAccessRequestApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessRequestApproval`: []Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.ListAccessRequestApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRequestApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]Approval**](Approval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectAccessRequest

> Approval RejectAccessRequest(ctx, approvalId).CommentEto(commentEto).Execute()

Rejects an access request.



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
    approvalId := "approvalId_example" // string | ID of an access request approval.
    commentEto := *openapiclient.NewCommentEto() // CommentEto | Reason about the approval rejection. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApprovalsApi.RejectAccessRequest(context.Background(), approvalId).CommentEto(commentEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApprovalsApi.RejectAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectAccessRequest`: Approval
    fmt.Fprintf(os.Stdout, "Response from `ApprovalsApi.RejectAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | ID of an access request approval. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentEto** | [**CommentEto**](CommentEto.md) | Reason about the approval rejection. | 

### Return type

[**Approval**](Approval.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

