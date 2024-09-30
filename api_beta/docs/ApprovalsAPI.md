# \ApprovalsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApproval**](ApprovalsAPI.md#GetApproval) | **Get** /generic-approvals/{id} | Get Approval
[**GetApprovals**](ApprovalsAPI.md#GetApprovals) | **Get** /generic-approvals | Get Approvals



## GetApproval

> Approval GetApproval(ctx, id).Execute()

Get Approval



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
	id := "38453251-6be2-5f8f-df93-5ce19e295837" // string | ID of the approval that to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.GetApproval(context.Background(), id).Execute()
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
**id** | **string** | ID of the approval that to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Approval**](Approval.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApprovals

> []Approval GetApprovals(ctx).Mine(mine).RequesterId(requesterId).Filters(filters).Execute()

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
	mine := true // bool | Returns the list of approvals for the current caller. (optional)
	requesterId := "17e633e7d57e481569df76323169deb6a" // string | Returns the list of approvals for a given requester ID. (optional)
	filters := "filters=status eq PENDING" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq*  **referenceType**: *eq* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApprovalsAPI.GetApprovals(context.Background()).Mine(mine).RequesterId(requesterId).Filters(filters).Execute()
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
 **mine** | **bool** | Returns the list of approvals for the current caller. | 
 **requesterId** | **string** | Returns the list of approvals for a given requester ID. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq*  **referenceType**: *eq* | 

### Return type

[**[]Approval**](Approval.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

