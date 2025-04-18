# \AccessRequestApprovalsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveAccessRequest**](AccessRequestApprovalsAPI.md#ApproveAccessRequest) | **Post** /access-request-approvals/{approvalId}/approve | Approve Access Request Approval
[**ForwardAccessRequest**](AccessRequestApprovalsAPI.md#ForwardAccessRequest) | **Post** /access-request-approvals/{approvalId}/forward | Forward Access Request Approval
[**GetAccessRequestApprovalSummary**](AccessRequestApprovalsAPI.md#GetAccessRequestApprovalSummary) | **Get** /access-request-approvals/approval-summary | Get Access Requests Approvals Number
[**ListAccessRequestApprovers**](AccessRequestApprovalsAPI.md#ListAccessRequestApprovers) | **Get** /access-request-approvals/{accessRequestId}/approvers | Access Request Approvers
[**ListCompletedApprovals**](AccessRequestApprovalsAPI.md#ListCompletedApprovals) | **Get** /access-request-approvals/completed | Completed Access Request Approvals List
[**ListPendingApprovals**](AccessRequestApprovalsAPI.md#ListPendingApprovals) | **Get** /access-request-approvals/pending | Pending Access Request Approvals List
[**RejectAccessRequest**](AccessRequestApprovalsAPI.md#RejectAccessRequest) | **Post** /access-request-approvals/{approvalId}/reject | Reject Access Request Approval



## ApproveAccessRequest

> map[string]interface{} ApproveAccessRequest(ctx, approvalId).CommentDto(commentDto).Execute()

Approve Access Request Approval



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
	approvalId := "2c91808b7294bea301729568c68c002e" // string | Approval ID.
	commentDto := *openapiclient.NewCommentDto() // CommentDto | Reviewer's comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.ApproveAccessRequest(context.Background(), approvalId).CommentDto(commentDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.ApproveAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.ApproveAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | Approval ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentDto** | [**CommentDto**](CommentDto.md) | Reviewer&#39;s comment. | 

### Return type

**map[string]interface{}**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForwardAccessRequest

> map[string]interface{} ForwardAccessRequest(ctx, approvalId).ForwardApprovalDto(forwardApprovalDto).Execute()

Forward Access Request Approval



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
	approvalId := "2c91808b7294bea301729568c68c002e" // string | Approval ID.
	forwardApprovalDto := *openapiclient.NewForwardApprovalDto("2c91808568c529c60168cca6f90c1314", "2c91808568c529c60168cca6f90c1313") // ForwardApprovalDto | Information about the forwarded approval.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.ForwardAccessRequest(context.Background(), approvalId).ForwardApprovalDto(forwardApprovalDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.ForwardAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForwardAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.ForwardAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | Approval ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiForwardAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forwardApprovalDto** | [**ForwardApprovalDto**](ForwardApprovalDto.md) | Information about the forwarded approval. | 

### Return type

**map[string]interface{}**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestApprovalSummary

> ApprovalSummary GetAccessRequestApprovalSummary(ctx).OwnerId(ownerId).FromDate(fromDate).Execute()

Get Access Requests Approvals Number



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
	ownerId := "2c91808568c529c60168cca6f90c1313" // string | The ID of the owner or approver identity of the approvals. If present, the value returns approval summary for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN user can also fetch all the approvals in the org, when owner-id is not used.    * Non ORG_ADMIN users can only specify *me* or pass their own identity ID value. (optional)
	fromDate := "from-date=2020-03-19T19:59:11Z" // string | This is the date and time the results will be shown from. It must be in a valid ISO-8601 format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.GetAccessRequestApprovalSummary(context.Background()).OwnerId(ownerId).FromDate(fromDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.GetAccessRequestApprovalSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestApprovalSummary`: ApprovalSummary
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.GetAccessRequestApprovalSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestApprovalSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | The ID of the owner or approver identity of the approvals. If present, the value returns approval summary for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN user can also fetch all the approvals in the org, when owner-id is not used.    * Non ORG_ADMIN users can only specify *me* or pass their own identity ID value. | 
 **fromDate** | **string** | This is the date and time the results will be shown from. It must be in a valid ISO-8601 format. | 

### Return type

[**ApprovalSummary**](ApprovalSummary.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessRequestApprovers

> []AccessRequestApproversListResponse ListAccessRequestApprovers(ctx, accessRequestId).Limit(limit).Offset(offset).Count(count).Execute()

Access Request Approvers



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
	accessRequestId := "2c91808568c529c60168cca6f90c1313" // string | Access Request ID.
	limit := int32(100) // int32 | Max number of results to return. (optional) (default to 250)
	offset := int32(10) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. (optional)
	count := false // bool | If this is true, the *X-Total-Count* response header populates with the number of results that would be returned if limit and offset were ignored. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.ListAccessRequestApprovers(context.Background(), accessRequestId).Limit(limit).Offset(offset).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.ListAccessRequestApprovers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessRequestApprovers`: []AccessRequestApproversListResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.ListAccessRequestApprovers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessRequestId** | **string** | Access Request ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRequestApproversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. | 
 **count** | **bool** | If this is true, the *X-Total-Count* response header populates with the number of results that would be returned if limit and offset were ignored. | [default to false]

### Return type

[**[]AccessRequestApproversListResponse**](AccessRequestApproversListResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompletedApprovals

> []CompletedApproval ListCompletedApprovals(ctx).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Completed Access Request Approvals List



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
	ownerId := "2c91808568c529c60168cca6f90c1313" // string | If present, the value returns only completed approvals for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN users can also fetch all the approvals in the org, when owner-id is not used.    * Non-ORG_ADMIN users can only specify *me* or pass their own identity ID value. (optional)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "id eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **requestedFor.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **modified**: *gt, lt, ge, le, eq, in, ne, sw* (optional)
	sorters := "modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.ListCompletedApprovals(context.Background()).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.ListCompletedApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompletedApprovals`: []CompletedApproval
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.ListCompletedApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompletedApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | If present, the value returns only completed approvals for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN users can also fetch all the approvals in the org, when owner-id is not used.    * Non-ORG_ADMIN users can only specify *me* or pass their own identity ID value. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **requestedFor.id**: *eq, in, ge, gt, le, lt, ne, isnull, sw*  **modified**: *gt, lt, ge, le, eq, in, ne, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]CompletedApproval**](CompletedApproval.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingApprovals

> []PendingApproval ListPendingApprovals(ctx).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Pending Access Request Approvals List



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
	ownerId := "2c91808568c529c60168cca6f90c1313" // string | If present, the value returns only pending approvals for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN users can also fetch all the approvals in the org, when owner-id is not used.    * Non-ORG_ADMIN users can only specify *me* or pass their own identity ID value. (optional)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "id eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **requestedFor.id**: *eq, in*  **modified**: *gt, lt, ge, le, eq, in*  **accessRequestId**: *eq, in* (optional)
	sorters := "modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.ListPendingApprovals(context.Background()).OwnerId(ownerId).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.ListPendingApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPendingApprovals`: []PendingApproval
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.ListPendingApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPendingApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | If present, the value returns only pending approvals for the specified identity.    * ORG_ADMIN users can call this with any identity ID value.    * ORG_ADMIN users can also fetch all the approvals in the org, when owner-id is not used.    * Non-ORG_ADMIN users can only specify *me* or pass their own identity ID value. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **requestedFor.id**: *eq, in*  **modified**: *gt, lt, ge, le, eq, in*  **accessRequestId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]PendingApproval**](PendingApproval.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectAccessRequest

> map[string]interface{} RejectAccessRequest(ctx, approvalId).CommentDto(commentDto).Execute()

Reject Access Request Approval



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
	approvalId := "2c91808b7294bea301729568c68c002e" // string | Approval ID.
	commentDto := *openapiclient.NewCommentDto() // CommentDto | Reviewer's comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestApprovalsAPI.RejectAccessRequest(context.Background(), approvalId).CommentDto(commentDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestApprovalsAPI.RejectAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RejectAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestApprovalsAPI.RejectAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**approvalId** | **string** | Approval ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentDto** | [**CommentDto**](CommentDto.md) | Reviewer&#39;s comment. | 

### Return type

**map[string]interface{}**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

