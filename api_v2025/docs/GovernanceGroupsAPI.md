# \GovernanceGroupsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkgroup**](GovernanceGroupsAPI.md#CreateWorkgroup) | **Post** /workgroups | Create a new Governance Group.
[**DeleteWorkgroup**](GovernanceGroupsAPI.md#DeleteWorkgroup) | **Delete** /workgroups/{id} | Delete a Governance Group
[**DeleteWorkgroupMembers**](GovernanceGroupsAPI.md#DeleteWorkgroupMembers) | **Post** /workgroups/{workgroupId}/members/bulk-delete | Remove members from Governance Group
[**DeleteWorkgroupsInBulk**](GovernanceGroupsAPI.md#DeleteWorkgroupsInBulk) | **Post** /workgroups/bulk-delete | Delete Governance Group(s)
[**GetWorkgroup**](GovernanceGroupsAPI.md#GetWorkgroup) | **Get** /workgroups/{id} | Get Governance Group by Id
[**ListConnections**](GovernanceGroupsAPI.md#ListConnections) | **Get** /workgroups/{workgroupId}/connections | List connections for Governance Group
[**ListWorkgroupMembers**](GovernanceGroupsAPI.md#ListWorkgroupMembers) | **Get** /workgroups/{workgroupId}/members | List Governance Group Members
[**ListWorkgroups**](GovernanceGroupsAPI.md#ListWorkgroups) | **Get** /workgroups | List Governance Groups
[**PatchWorkgroup**](GovernanceGroupsAPI.md#PatchWorkgroup) | **Patch** /workgroups/{id} | Patch a Governance Group
[**UpdateWorkgroupMembers**](GovernanceGroupsAPI.md#UpdateWorkgroupMembers) | **Post** /workgroups/{workgroupId}/members/bulk-add | Add members to Governance Group



## CreateWorkgroup

> WorkgroupDto CreateWorkgroup(ctx).XSailPointExperimental(xSailPointExperimental).WorkgroupDto(workgroupDto).Execute()

Create a new Governance Group.



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
	workgroupDto := *openapiclient.NewWorkgroupDto() // WorkgroupDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.CreateWorkgroup(context.Background()).XSailPointExperimental(xSailPointExperimental).WorkgroupDto(workgroupDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.CreateWorkgroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkgroup`: WorkgroupDto
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.CreateWorkgroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **workgroupDto** | [**WorkgroupDto**](WorkgroupDto.md) |  | 

### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroup

> DeleteWorkgroup(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Delete a Governance Group



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
	id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GovernanceGroupsAPI.DeleteWorkgroup(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
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
**id** | **string** | ID of the Governance Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroupMembers

> []WorkgroupMemberDeleteItem DeleteWorkgroupMembers(ctx, workgroupId).XSailPointExperimental(xSailPointExperimental).IdentityPreviewResponseIdentity(identityPreviewResponseIdentity).Execute()

Remove members from Governance Group



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
	workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	identityPreviewResponseIdentity := []openapiclient.IdentityPreviewResponseIdentity{*openapiclient.NewIdentityPreviewResponseIdentity()} // []IdentityPreviewResponseIdentity | List of identities to be removed from  a Governance Group members list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.DeleteWorkgroupMembers(context.Background(), workgroupId).XSailPointExperimental(xSailPointExperimental).IdentityPreviewResponseIdentity(identityPreviewResponseIdentity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.DeleteWorkgroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkgroupMembers`: []WorkgroupMemberDeleteItem
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.DeleteWorkgroupMembers`: %v\n", resp)
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **identityPreviewResponseIdentity** | [**[]IdentityPreviewResponseIdentity**](IdentityPreviewResponseIdentity.md) | List of identities to be removed from  a Governance Group members list. | 

### Return type

[**[]WorkgroupMemberDeleteItem**](WorkgroupMemberDeleteItem.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkgroupsInBulk

> []WorkgroupDeleteItem DeleteWorkgroupsInBulk(ctx).XSailPointExperimental(xSailPointExperimental).WorkgroupBulkDeleteRequest(workgroupBulkDeleteRequest).Execute()

Delete Governance Group(s)



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
	workgroupBulkDeleteRequest := *openapiclient.NewWorkgroupBulkDeleteRequest() // WorkgroupBulkDeleteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.DeleteWorkgroupsInBulk(context.Background()).XSailPointExperimental(xSailPointExperimental).WorkgroupBulkDeleteRequest(workgroupBulkDeleteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.DeleteWorkgroupsInBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkgroupsInBulk`: []WorkgroupDeleteItem
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.DeleteWorkgroupsInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkgroupsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **workgroupBulkDeleteRequest** | [**WorkgroupBulkDeleteRequest**](WorkgroupBulkDeleteRequest.md) |  | 

### Return type

[**[]WorkgroupDeleteItem**](WorkgroupDeleteItem.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkgroup

> WorkgroupDto GetWorkgroup(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Get Governance Group by Id



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
	id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.GetWorkgroup(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.GetWorkgroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkgroup`: WorkgroupDto
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.GetWorkgroup`: %v\n", resp)
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> []WorkgroupConnectionDto ListConnections(ctx, workgroupId).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()

List connections for Governance Group



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
	workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.ListConnections(context.Background(), workgroupId).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ListConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnections`: []WorkgroupConnectionDto
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.ListConnections`: %v\n", resp)
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** | 

### Return type

[**[]WorkgroupConnectionDto**](WorkgroupConnectionDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroupMembers

> []ListWorkgroupMembers200ResponseInner ListWorkgroupMembers(ctx, workgroupId).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()

List Governance Group Members



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
	workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.ListWorkgroupMembers(context.Background(), workgroupId).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Execute()
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
**workgroupId** | **string** | ID of the Governance Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** | 

### Return type

[**[]ListWorkgroupMembers200ResponseInner**](ListWorkgroupMembers200ResponseInner.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkgroups

> []WorkgroupDto ListWorkgroups(ctx).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()

List Governance Groups



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
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "name sw "Test"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **name**: *eq, sw, in*  **memberships.identityId**: *eq, in* (optional)
	sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified, id, description** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.ListWorkgroups(context.Background()).XSailPointExperimental(xSailPointExperimental).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.ListWorkgroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkgroups`: []WorkgroupDto
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.ListWorkgroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkgroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in, sw*  **name**: *eq, sw, in*  **memberships.identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified, id, description** | 

### Return type

[**[]WorkgroupDto**](WorkgroupDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkgroup

> WorkgroupDto PatchWorkgroup(ctx, id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a Governance Group



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
	id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Governance Group
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.PatchWorkgroup(context.Background(), id).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.PatchWorkgroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchWorkgroup`: WorkgroupDto
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.PatchWorkgroup`: %v\n", resp)
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**WorkgroupDto**](WorkgroupDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkgroupMembers

> []WorkgroupMemberAddItem UpdateWorkgroupMembers(ctx, workgroupId).XSailPointExperimental(xSailPointExperimental).IdentityPreviewResponseIdentity(identityPreviewResponseIdentity).Execute()

Add members to Governance Group



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
	workgroupId := "2c91808a7813090a017814121919ecca" // string | ID of the Governance Group.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	identityPreviewResponseIdentity := []openapiclient.IdentityPreviewResponseIdentity{*openapiclient.NewIdentityPreviewResponseIdentity()} // []IdentityPreviewResponseIdentity | List of identities to be added to a Governance Group members list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GovernanceGroupsAPI.UpdateWorkgroupMembers(context.Background(), workgroupId).XSailPointExperimental(xSailPointExperimental).IdentityPreviewResponseIdentity(identityPreviewResponseIdentity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GovernanceGroupsAPI.UpdateWorkgroupMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkgroupMembers`: []WorkgroupMemberAddItem
	fmt.Fprintf(os.Stdout, "Response from `GovernanceGroupsAPI.UpdateWorkgroupMembers`: %v\n", resp)
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

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **identityPreviewResponseIdentity** | [**[]IdentityPreviewResponseIdentity**](IdentityPreviewResponseIdentity.md) | List of identities to be added to a Governance Group members list. | 

### Return type

[**[]WorkgroupMemberAddItem**](WorkgroupMemberAddItem.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

