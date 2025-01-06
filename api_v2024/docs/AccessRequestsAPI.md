# \AccessRequestsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveBulkAccessRequest**](AccessRequestsAPI.md#ApproveBulkAccessRequest) | **Post** /access-request-approvals/bulk-approve | Bulk Approve Access Request
[**CancelAccessRequest**](AccessRequestsAPI.md#CancelAccessRequest) | **Post** /access-requests/cancel | Cancel Access Request
[**CancelAccessRequestInBulk**](AccessRequestsAPI.md#CancelAccessRequestInBulk) | **Post** /access-requests/bulk-cancel | Bulk Cancel Access Request
[**CloseAccessRequest**](AccessRequestsAPI.md#CloseAccessRequest) | **Post** /access-requests/close | Close Access Request
[**CreateAccessRequest**](AccessRequestsAPI.md#CreateAccessRequest) | **Post** /access-requests | Submit Access Request
[**GetAccessRequestConfig**](AccessRequestsAPI.md#GetAccessRequestConfig) | **Get** /access-request-config | Get Access Request Configuration
[**ListAccessRequestStatus**](AccessRequestsAPI.md#ListAccessRequestStatus) | **Get** /access-request-status | Access Request Status
[**SetAccessRequestConfig**](AccessRequestsAPI.md#SetAccessRequestConfig) | **Put** /access-request-config | Update Access Request Configuration



## ApproveBulkAccessRequest

> map[string]interface{} ApproveBulkAccessRequest(ctx).BulkApproveAccessRequest(bulkApproveAccessRequest).Execute()

Bulk Approve Access Request



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
	bulkApproveAccessRequest := *openapiclient.NewBulkApproveAccessRequest([]string{"ApprovalIds_example"}, "I approve these request items") // BulkApproveAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.ApproveBulkAccessRequest(context.Background()).BulkApproveAccessRequest(bulkApproveAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.ApproveBulkAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApproveBulkAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.ApproveBulkAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApproveBulkAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkApproveAccessRequest** | [**BulkApproveAccessRequest**](BulkApproveAccessRequest.md) |  | 

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


## CancelAccessRequest

> map[string]interface{} CancelAccessRequest(ctx).CancelAccessRequest(cancelAccessRequest).Execute()

Cancel Access Request



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
	cancelAccessRequest := *openapiclient.NewCancelAccessRequest("2c9180835d2e5168015d32f890ca1581", "I requested this role by mistake.") // CancelAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.CancelAccessRequest(context.Background()).CancelAccessRequest(cancelAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.CancelAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.CancelAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelAccessRequest** | [**CancelAccessRequest**](CancelAccessRequest.md) |  | 

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


## CancelAccessRequestInBulk

> map[string]interface{} CancelAccessRequestInBulk(ctx).BulkCancelAccessRequest(bulkCancelAccessRequest).Execute()

Bulk Cancel Access Request



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
	bulkCancelAccessRequest := *openapiclient.NewBulkCancelAccessRequest([]string{"AccessRequestIds_example"}, "I requested this role by mistake.") // BulkCancelAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.CancelAccessRequestInBulk(context.Background()).BulkCancelAccessRequest(bulkCancelAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.CancelAccessRequestInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAccessRequestInBulk`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.CancelAccessRequestInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAccessRequestInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkCancelAccessRequest** | [**BulkCancelAccessRequest**](BulkCancelAccessRequest.md) |  | 

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


## CloseAccessRequest

> map[string]interface{} CloseAccessRequest(ctx).XSailPointExperimental(xSailPointExperimental).CloseAccessRequest(closeAccessRequest).Execute()

Close Access Request



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
	closeAccessRequest := *openapiclient.NewCloseAccessRequest([]string{"AccessRequestIds_example"}) // CloseAccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.CloseAccessRequest(context.Background()).XSailPointExperimental(xSailPointExperimental).CloseAccessRequest(closeAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.CloseAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloseAccessRequest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.CloseAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **closeAccessRequest** | [**CloseAccessRequest**](CloseAccessRequest.md) |  | 

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


## CreateAccessRequest

> AccessRequestResponse CreateAccessRequest(ctx).AccessRequest(accessRequest).Execute()

Submit Access Request



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
	accessRequest := *openapiclient.NewAccessRequest([]string{"2c918084660f45d6016617daa9210584"}, []openapiclient.AccessRequestItem{*openapiclient.NewAccessRequestItem("ACCESS_PROFILE", "2c9180835d2e5168015d32f890ca1581")}) // AccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.CreateAccessRequest(context.Background()).AccessRequest(accessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.CreateAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessRequest`: AccessRequestResponse
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.CreateAccessRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequest** | [**AccessRequest**](AccessRequest.md) |  | 

### Return type

[**AccessRequestResponse**](AccessRequestResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestConfig

> AccessRequestConfig GetAccessRequestConfig(ctx).Execute()

Get Access Request Configuration



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.GetAccessRequestConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.GetAccessRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestConfig`: AccessRequestConfig
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.GetAccessRequestConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestConfigRequest struct via the builder pattern


### Return type

[**AccessRequestConfig**](AccessRequestConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessRequestStatus

> []RequestedItemStatus ListAccessRequestStatus(ctx).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).AssignedTo(assignedTo).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).RequestState(requestState).Execute()

Access Request Status



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
	requestedFor := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity the requests were made for. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
	requestedBy := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity who made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
	regardingIdentity := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the specified identity who is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. (optional)
	assignedTo := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the specified identity who is the owner of the Identity Request Work Item. *me* indicates the current user. (optional)
	count := false // bool | If this is true, the *X-Total-Count* response header populates with the number of results that would be returned if limit and offset were ignored. (optional) (default to false)
	limit := int32(100) // int32 | Max number of results to return. (optional) (default to 250)
	offset := int32(10) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. (optional)
	filters := "accountActivityItemId eq "2c918086771c86df0177401efcdf54c0"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accessRequestId**: *in*  **accountActivityItemId**: *eq, in, ge, gt, le, lt, ne, isnull, sw* (optional)
	sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId, name** (optional)
	requestState := "request-state=EXECUTING" // string | Filter the results by the state of the request. The only valid value is *EXECUTING*. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.ListAccessRequestStatus(context.Background()).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).AssignedTo(assignedTo).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).RequestState(requestState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.ListAccessRequestStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessRequestStatus`: []RequestedItemStatus
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.ListAccessRequestStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRequestStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | Filter the results by the identity the requests were made for. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **requestedBy** | **string** | Filter the results by the identity who made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **regardingIdentity** | **string** | Filter the results by the specified identity who is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. | 
 **assignedTo** | **string** | Filter the results by the specified identity who is the owner of the Identity Request Work Item. *me* indicates the current user. | 
 **count** | **bool** | If this is true, the *X-Total-Count* response header populates with the number of results that would be returned if limit and offset were ignored. | [default to false]
 **limit** | **int32** | Max number of results to return. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accessRequestId**: *in*  **accountActivityItemId**: *eq, in, ge, gt, le, lt, ne, isnull, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId, name** | 
 **requestState** | **string** | Filter the results by the state of the request. The only valid value is *EXECUTING*. | 

### Return type

[**[]RequestedItemStatus**](RequestedItemStatus.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccessRequestConfig

> AccessRequestConfig SetAccessRequestConfig(ctx).AccessRequestConfig(accessRequestConfig).Execute()

Update Access Request Configuration



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
	accessRequestConfig := *openapiclient.NewAccessRequestConfig() // AccessRequestConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestsAPI.SetAccessRequestConfig(context.Background()).AccessRequestConfig(accessRequestConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.SetAccessRequestConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAccessRequestConfig`: AccessRequestConfig
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.SetAccessRequestConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestConfig** | [**AccessRequestConfig**](AccessRequestConfig.md) |  | 

### Return type

[**AccessRequestConfig**](AccessRequestConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

