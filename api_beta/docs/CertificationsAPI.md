# \CertificationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentityCertificationItemPermissions**](CertificationsAPI.md#GetIdentityCertificationItemPermissions) | **Get** /certifications/{certificationId}/access-review-items/{itemId}/permissions | Permissions for Entitlement Certification Item
[**GetIdentityCertificationPendingTasks**](CertificationsAPI.md#GetIdentityCertificationPendingTasks) | **Get** /certifications/{id}/tasks-pending | Pending Certification Tasks
[**GetIdentityCertificationTaskStatus**](CertificationsAPI.md#GetIdentityCertificationTaskStatus) | **Get** /certifications/{id}/tasks/{taskId} | Certification Task Status
[**ListCertificationReviewers**](CertificationsAPI.md#ListCertificationReviewers) | **Get** /certifications/{id}/reviewers | List of Reviewers for certification
[**ListCertifications**](CertificationsAPI.md#ListCertifications) | **Get** /certifications | Certifications by IDs
[**SubmitReassignCertsAsync**](CertificationsAPI.md#SubmitReassignCertsAsync) | **Post** /certifications/{id}/reassign-async | Reassign Certifications Asynchronously



## GetIdentityCertificationItemPermissions

> []PermissionDto GetIdentityCertificationItemPermissions(ctx, certificationId, itemId).Filters(filters).Limit(limit).Offset(offset).Count(count).Execute()

Permissions for Entitlement Certification Item



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
	certificationId := "ef38f94347e94562b5bb8424a56397d8" // string | The certification ID
	itemId := "2c91808671bcbab40171bd945d961227" // string | The certification item ID
	filters := "target eq "SYS.OBJAUTH2"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **target**: *eq, sw*  **rights**: *ca*  All field values (second filter operands) are case-insensitive for this API.  Only a single *and* or *or* composite filter operator may be used. It must also be used between a target filter and a rights filter, not between 2 filters for the same field.  For example, the following is valid: `?filters=rights+ca+(%22CREATE%22)+and+target+eq+%22SYS.OBJAUTH2%22`  The following is invalid: `?filters=rights+ca+(%22CREATE%22)+and+rights+ca+(%SELECT%22)` (optional)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.GetIdentityCertificationItemPermissions(context.Background(), certificationId, itemId).Filters(filters).Limit(limit).Offset(offset).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.GetIdentityCertificationItemPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityCertificationItemPermissions`: []PermissionDto
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.GetIdentityCertificationItemPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificationId** | **string** | The certification ID | 
**itemId** | **string** | The certification item ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityCertificationItemPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **target**: *eq, sw*  **rights**: *ca*  All field values (second filter operands) are case-insensitive for this API.  Only a single *and* or *or* composite filter operator may be used. It must also be used between a target filter and a rights filter, not between 2 filters for the same field.  For example, the following is valid: &#x60;?filters&#x3D;rights+ca+(%22CREATE%22)+and+target+eq+%22SYS.OBJAUTH2%22&#x60;  The following is invalid: &#x60;?filters&#x3D;rights+ca+(%22CREATE%22)+and+rights+ca+(%SELECT%22)&#x60; | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]PermissionDto**](PermissionDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityCertificationPendingTasks

> []IdentityCertificationTask GetIdentityCertificationPendingTasks(ctx, id).Execute()

Pending Certification Tasks



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
	id := "id_example" // string | The identity campaign certification ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.GetIdentityCertificationPendingTasks(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.GetIdentityCertificationPendingTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityCertificationPendingTasks`: []IdentityCertificationTask
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.GetIdentityCertificationPendingTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityCertificationPendingTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IdentityCertificationTask**](IdentityCertificationTask.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityCertificationTaskStatus

> IdentityCertificationTask GetIdentityCertificationTaskStatus(ctx, id, taskId).Execute()

Certification Task Status



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
	id := "id_example" // string | The identity campaign certification ID
	taskId := "taskId_example" // string | The certification task ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.GetIdentityCertificationTaskStatus(context.Background(), id, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.GetIdentityCertificationTaskStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityCertificationTaskStatus`: IdentityCertificationTask
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.GetIdentityCertificationTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 
**taskId** | **string** | The certification task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityCertificationTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityCertificationTask**](IdentityCertificationTask.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertificationReviewers

> []IdentityReferenceWithNameAndEmail ListCertificationReviewers(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List of Reviewers for certification



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | The certification ID
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "name eq "Bob"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **email**: *eq, sw* (optional)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, email** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.ListCertificationReviewers(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.ListCertificationReviewers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertificationReviewers`: []IdentityReferenceWithNameAndEmail
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.ListCertificationReviewers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCertificationReviewersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **email**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, email** | 

### Return type

[**[]IdentityReferenceWithNameAndEmail**](IdentityReferenceWithNameAndEmail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCertifications

> []CertificationDto ListCertifications(ctx).ReviewerIdentitiy(reviewerIdentitiy).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Certifications by IDs



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
	reviewerIdentitiy := "reviewerIdentitiy_example" // string | The ID of reviewer identity. *me* indicates the current user. (optional)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "filters_example" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **phase**: *eq*  **completed**: *eq, ne*  **campaignRef.campaignType**: *eq, in*  **campaignRef.id**: *eq, in* (optional)
	sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, due, signed** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.ListCertifications(context.Background()).ReviewerIdentitiy(reviewerIdentitiy).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.ListCertifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCertifications`: []CertificationDto
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.ListCertifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewerIdentitiy** | **string** | The ID of reviewer identity. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **phase**: *eq*  **completed**: *eq, ne*  **campaignRef.campaignType**: *eq, in*  **campaignRef.id**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, due, signed** | 

### Return type

[**[]CertificationDto**](CertificationDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitReassignCertsAsync

> IdentityCertificationTask SubmitReassignCertsAsync(ctx, id).ReviewReassign(reviewReassign).Execute()

Reassign Certifications Asynchronously



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID
	reviewReassign := *openapiclient.NewReviewReassign([]openapiclient.ReassignReference{*openapiclient.NewReassignReference("ef38f94347e94562b5bb8424a56397d8", "ITEM")}, "ef38f94347e94562b5bb8424a56397d8", "reassigned for some reason") // ReviewReassign | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CertificationsAPI.SubmitReassignCertsAsync(context.Background(), id).ReviewReassign(reviewReassign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificationsAPI.SubmitReassignCertsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitReassignCertsAsync`: IdentityCertificationTask
	fmt.Fprintf(os.Stdout, "Response from `CertificationsAPI.SubmitReassignCertsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitReassignCertsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewReassign** | [**ReviewReassign**](ReviewReassign.md) |  | 

### Return type

[**IdentityCertificationTask**](IdentityCertificationTask.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

