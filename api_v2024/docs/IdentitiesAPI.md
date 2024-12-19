# \IdentitiesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIdentity**](IdentitiesAPI.md#DeleteIdentity) | **Delete** /identities/{id} | Delete identity
[**GetIdentity**](IdentitiesAPI.md#GetIdentity) | **Get** /identities/{id} | Identity Details
[**GetIdentityOwnershipDetails**](IdentitiesAPI.md#GetIdentityOwnershipDetails) | **Get** /identities/{identityId}/ownership | Get ownership details
[**GetRoleAssignment**](IdentitiesAPI.md#GetRoleAssignment) | **Get** /identities/{identityId}/role-assignments/{assignmentId} | Role assignment details
[**GetRoleAssignments**](IdentitiesAPI.md#GetRoleAssignments) | **Get** /identities/{identityId}/role-assignments | List role assignments
[**ListIdentities**](IdentitiesAPI.md#ListIdentities) | **Get** /identities | List Identities
[**ResetIdentity**](IdentitiesAPI.md#ResetIdentity) | **Post** /identities/{id}/reset | Reset an identity
[**SendIdentityVerificationAccountToken**](IdentitiesAPI.md#SendIdentityVerificationAccountToken) | **Post** /identities/{id}/verification/account/send | Send password reset email
[**StartIdentitiesInvite**](IdentitiesAPI.md#StartIdentitiesInvite) | **Post** /identities/invite | Invite identities to register
[**StartIdentityProcessing**](IdentitiesAPI.md#StartIdentityProcessing) | **Post** /identities/process | Process a list of identityIds
[**SynchronizeAttributesForIdentity**](IdentitiesAPI.md#SynchronizeAttributesForIdentity) | **Post** /identities/{identityId}/synchronize-attributes | Attribute synchronization for single identity.



## DeleteIdentity

> DeleteIdentity(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Delete identity



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitiesAPI.DeleteIdentity(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.DeleteIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentity

> Identity GetIdentity(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Identity Details



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentity(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentity`: Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**Identity**](Identity.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityOwnershipDetails

> IdentityOwnershipAssociationDetails GetIdentityOwnershipDetails(ctx, identityId).XSailPointExperimental(xSailPointExperimental).Execute()

Get ownership details



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
	identityId := "ff8081814d2a8036014d701f3fbf53fa" // string | Identity ID.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetIdentityOwnershipDetails(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetIdentityOwnershipDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdentityOwnershipDetails`: IdentityOwnershipAssociationDetails
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetIdentityOwnershipDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | Identity ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityOwnershipDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**IdentityOwnershipAssociationDetails**](IdentityOwnershipAssociationDetails.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAssignment

> RoleAssignmentDto GetRoleAssignment(ctx, identityId, assignmentId).XSailPointExperimental(xSailPointExperimental).Execute()

Role assignment details

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
	identityId := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Id
	assignmentId := "1cbb0705b38c4226b1334eadd8874086" // string | Assignment Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetRoleAssignment(context.Background(), identityId, assignmentId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetRoleAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleAssignment`: RoleAssignmentDto
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | Identity Id | 
**assignmentId** | **string** | Assignment Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**RoleAssignmentDto**](RoleAssignmentDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleAssignments

> []GetRoleAssignments200ResponseInner GetRoleAssignments(ctx, identityId).XSailPointExperimental(xSailPointExperimental).RoleId(roleId).RoleName(roleName).Execute()

List role assignments



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
	identityId := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Id to get the role assignments for
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	roleId := "e7697a1e96d04db1ac7b0f4544915d2c" // string | Role Id to filter the role assignments with (optional)
	roleName := "Engineer" // string | Role name to filter the role assignments with (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.GetRoleAssignments(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).RoleId(roleId).RoleName(roleName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.GetRoleAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleAssignments`: []GetRoleAssignments200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.GetRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | Identity Id to get the role assignments for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **roleId** | **string** | Role Id to filter the role assignments with | 
 **roleName** | **string** | Role name to filter the role assignments with | 

### Return type

[**[]GetRoleAssignments200ResponseInner**](GetRoleAssignments200ResponseInner.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentities

> []Identity ListIdentities(ctx).XSailPointExperimental(xSailPointExperimental).Filters(filters).Sorters(sorters).DefaultFilter(defaultFilter).Count(count).Limit(limit).Offset(offset).Execute()

List Identities



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
	filters := "id eq "6c9079b270a266a60170a2779fcb0006" or correlated eq false" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **alias**: *eq, sw*  **firstname**: *eq, sw*  **lastname**: *eq, sw*  **email**: *eq, sw*  **cloudStatus**: *eq*  **processingState**: *eq*  **correlated**: *eq*  **protected**: *eq* (optional)
	sorters := "name,-cloudStatus" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, alias, cloudStatus** (optional)
	defaultFilter := "NONE" // string | Adds additional filter to filters query parameter.  CORRELATED_ONLY adds correlated=true and returns only identities that are correlated.  NONE does not add any and returns all identities that satisfy filters query parameter. (optional) (default to "CORRELATED_ONLY")
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.ListIdentities(context.Background()).XSailPointExperimental(xSailPointExperimental).Filters(filters).Sorters(sorters).DefaultFilter(defaultFilter).Count(count).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.ListIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdentities`: []Identity
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.ListIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **alias**: *eq, sw*  **firstname**: *eq, sw*  **lastname**: *eq, sw*  **email**: *eq, sw*  **cloudStatus**: *eq*  **processingState**: *eq*  **correlated**: *eq*  **protected**: *eq* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, alias, cloudStatus** | 
 **defaultFilter** | **string** | Adds additional filter to filters query parameter.  CORRELATED_ONLY adds correlated&#x3D;true and returns only identities that are correlated.  NONE does not add any and returns all identities that satisfy filters query parameter. | [default to &quot;CORRELATED_ONLY&quot;]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]Identity**](Identity.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetIdentity

> ResetIdentity(ctx, identityId).XSailPointExperimental(xSailPointExperimental).Execute()

Reset an identity



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
	identityId := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitiesAPI.ResetIdentity(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.ResetIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | Identity Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendIdentityVerificationAccountToken

> SendIdentityVerificationAccountToken(ctx, id).XSailPointExperimental(xSailPointExperimental).SendAccountVerificationRequest(sendAccountVerificationRequest).Execute()

Send password reset email



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity ID
	sendAccountVerificationRequest := *openapiclient.NewSendAccountVerificationRequest("EMAIL_WORK") // SendAccountVerificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IdentitiesAPI.SendIdentityVerificationAccountToken(context.Background(), id).XSailPointExperimental(xSailPointExperimental).SendAccountVerificationRequest(sendAccountVerificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.SendIdentityVerificationAccountToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendIdentityVerificationAccountTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

 **sendAccountVerificationRequest** | [**SendAccountVerificationRequest**](SendAccountVerificationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartIdentitiesInvite

> TaskStatus StartIdentitiesInvite(ctx).XSailPointExperimental(xSailPointExperimental).InviteIdentitiesRequest(inviteIdentitiesRequest).Execute()

Invite identities to register



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
	inviteIdentitiesRequest := *openapiclient.NewInviteIdentitiesRequest() // InviteIdentitiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.StartIdentitiesInvite(context.Background()).XSailPointExperimental(xSailPointExperimental).InviteIdentitiesRequest(inviteIdentitiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.StartIdentitiesInvite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartIdentitiesInvite`: TaskStatus
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.StartIdentitiesInvite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartIdentitiesInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **inviteIdentitiesRequest** | [**InviteIdentitiesRequest**](InviteIdentitiesRequest.md) |  | 

### Return type

[**TaskStatus**](TaskStatus.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartIdentityProcessing

> TaskResultResponse StartIdentityProcessing(ctx).XSailPointExperimental(xSailPointExperimental).ProcessIdentitiesRequest(processIdentitiesRequest).Execute()

Process a list of identityIds



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
	processIdentitiesRequest := *openapiclient.NewProcessIdentitiesRequest() // ProcessIdentitiesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.StartIdentityProcessing(context.Background()).XSailPointExperimental(xSailPointExperimental).ProcessIdentitiesRequest(processIdentitiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.StartIdentityProcessing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartIdentityProcessing`: TaskResultResponse
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.StartIdentityProcessing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartIdentityProcessingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **processIdentitiesRequest** | [**ProcessIdentitiesRequest**](ProcessIdentitiesRequest.md) |  | 

### Return type

[**TaskResultResponse**](TaskResultResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SynchronizeAttributesForIdentity

> IdentitySyncJob SynchronizeAttributesForIdentity(ctx, identityId).XSailPointExperimental(xSailPointExperimental).Execute()

Attribute synchronization for single identity.



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
	identityId := "identityId_example" // string | The Identity id
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IdentitiesAPI.SynchronizeAttributesForIdentity(context.Background(), identityId).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesAPI.SynchronizeAttributesForIdentity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SynchronizeAttributesForIdentity`: IdentitySyncJob
	fmt.Fprintf(os.Stdout, "Response from `IdentitiesAPI.SynchronizeAttributesForIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The Identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSynchronizeAttributesForIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**IdentitySyncJob**](IdentitySyncJob.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth), [applicationAuth](../README.md#applicationAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

