# \AccessRequestsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelAccessRequest**](AccessRequestsAPI.md#CancelAccessRequest) | **Post** /access-requests/cancel | Cancel Access Request
[**CreateAccessRequest**](AccessRequestsAPI.md#CreateAccessRequest) | **Post** /access-requests | Submit an Access Request
[**GetAccessRequestConfig**](AccessRequestsAPI.md#GetAccessRequestConfig) | **Get** /access-request-config | Get Access Request Configuration
[**ListAccessRequestStatus**](AccessRequestsAPI.md#ListAccessRequestStatus) | **Get** /access-request-status | Access Request Status
[**SetAccessRequestConfig**](AccessRequestsAPI.md#SetAccessRequestConfig) | **Put** /access-request-config | Update Access Request Configuration



## Cancel Access Request

> map[string]interface{} CancelAccessRequest(ctx).CancelAccessRequest(cancelAccessRequest).Execute()

This API endpoint cancels a pending access request. An access request can be cancelled only if it has not passed the approval step.
Any token with ORG_ADMIN authority or token of the user who originally requested the access request is required to cancel it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    cancelAccessRequest := *sailpoint.NewCancelAccessRequest("2c9180835d2e5168015d32f890ca1581", "I requested this role by mistake.") // CancelAccessRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessRequestsAPI.CancelAccessRequest(context.Background()).CancelAccessRequest(cancelAccessRequest).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Submit an Access Request

> map[string]interface{} CreateAccessRequest(ctx).AccessRequest(accessRequest).Execute()

This submits the access request into IdentityNow, where it will follow any IdentityNow approval processes.

Access requests are processed asynchronously by IdentityNow.  A success response from this endpoint means the request
has been submitted to IDN and is queued for processing.  Because this endpoint is asynchronous, it will not return an error
if you submit duplicate access requests in quick succession, or you submit an access request for access that is already in progress, approved, or rejected.
It is best practice to check for any existing access requests that reference the same access items before submitting a new access request.  This can
be accomplished by using the [access request status](https://developer.sailpoint.com/idn/api/v3/list-access-request-status) or the [pending access request approvals](https://developer.sailpoint.com/idn/api/v3/list-pending-approvals) endpoints.  You can also
use the [search API](https://developer.sailpoint.com/idn/api/v3/search) to check the existing access items that an identity has before submitting
an access request to ensure you are not requesting access that is already granted.

There are two types of access request:

__GRANT_ACCESS__
* Can be requested for multiple identities in a single request.
* Supports self request and request on behalf of other users. Refer to the [Get Access Request Configuration](https://developer.sailpoint.com/idn/api/v3/get-access-request-config) endpoint for request configuration options.  
* Allows any authenticated token (except API) to call this endpoint to request to grant access to themselves. Depending on the configuration, a user can request access for others.
* Roles, access profiles and entitlements can be requested.
* While requesting entitlements, maximum of 25 entitlements and 10 recipients are allowed in a request.
 
__REVOKE_ACCESS__
* Can only be requested for a single identity at a time.
* You cannot use an access request to revoke access from an identity if that access has been granted by role membership or by birthright provisioning. 
* Does not support self request. Only manager can request to revoke access for their directly managed employees.
* If a `removeDate` is specified, then the access will be removed on that date and time only for roles and access profiles. Entitlements are currently unsupported for `removeDate`.
* Roles, access profiles, and entitlements can be requested for revocation.
* Revoke requests for entitlements are limited to 1 entitlement per access request currently.
* [Roles, Access Profiles] You can specify a `removeDate` if the access doesn't already have a sunset date. The `removeDate` must be a future date, in the UTC timezone. 
* Allows a manager to request to revoke access for direct employees. A token with ORG_ADMIN authority can also request to revoke access from anyone.

>**Note:** There is no indication to the approver in the IdentityNow UI that the approval request is for a revoke action. Take this into consideration when calling this API.

A token with API authority cannot be used to call this endpoint. 


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    accessRequest := *sailpoint.NewAccessRequest([]string{"2c918084660f45d6016617daa9210584"}, []sailpoint.AccessRequestItem{*sailpoint.NewAccessRequestItem("ACCESS_PROFILE", "2c9180835d2e5168015d32f890ca1581")}) // AccessRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessRequestsAPI.CreateAccessRequest(context.Background()).AccessRequest(accessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.CreateAccessRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessRequest`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Access Request Configuration

> AccessRequestConfig GetAccessRequestConfig(ctx).Execute()

This endpoint returns the current access-request configuration.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessRequestsAPI.GetAccessRequestConfig(context.Background()).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Access Request Status

> []RequestedItemStatus ListAccessRequestStatus(ctx).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).AssignedTo(assignedTo).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).Execute()

The Access Request Status API returns a list of access request statuses based on the specified query parameters.
Any token with any authority can request their own status. A token with ORG_ADMIN authority is required to call this API to get a list of statuses for other users.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity for which the requests were made. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    requestedBy := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the identity that made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. (optional)
    regardingIdentity := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the specified identity which is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. (optional)
    assignedTo := "2c9180877b2b6ea4017b2c545f971429" // string | Filter the results by the specified identity which is the owner of the Identity Request Work Item. *me* indicates the current user. (optional)
    count := false // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. (optional) (default to false)
    limit := int32(100) // int32 | Max number of results to return. (optional) (default to 250)
    offset := int32(10) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. (optional)
    filters := "accountActivityItemId eq "2c918086771c86df0177401efcdf54c0"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accountActivityItemId**: *eq, in, ge, gt, le, lt, ne, isnull, sw* (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId, name** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessRequestsAPI.ListAccessRequestStatus(context.Background()).RequestedFor(requestedFor).RequestedBy(requestedBy).RegardingIdentity(regardingIdentity).AssignedTo(assignedTo).Count(count).Limit(limit).Offset(offset).Filters(filters).Sorters(sorters).Execute()
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
 **requestedFor** | **string** | Filter the results by the identity for which the requests were made. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **requestedBy** | **string** | Filter the results by the identity that made the requests. *me* indicates the current user. Mutually exclusive with *regarding-identity*. | 
 **regardingIdentity** | **string** | Filter the results by the specified identity which is either the requester or target of the requests. *me* indicates the current user. Mutually exclusive with *requested-for* and *requested-by*. | 
 **assignedTo** | **string** | Filter the results by the specified identity which is the owner of the Identity Request Work Item. *me* indicates the current user. | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored. | [default to false]
 **limit** | **int32** | Max number of results to return. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. Defaults to 0 if not specified. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **accountActivityItemId**: *eq, in, ge, gt, le, lt, ne, isnull, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified, accountActivityItemId, name** | 

### Return type

[**[]RequestedItemStatus**](RequestedItemStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Access Request Configuration

> AccessRequestConfig SetAccessRequestConfig(ctx).AccessRequestConfig(accessRequestConfig).Execute()

This endpoint replaces the current access-request configuration.
A token with ORG_ADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    accessRequestConfig := *sailpoint.NewAccessRequestConfig() // AccessRequestConfig | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessRequestsAPI.SetAccessRequestConfig(context.Background()).AccessRequestConfig(accessRequestConfig).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

