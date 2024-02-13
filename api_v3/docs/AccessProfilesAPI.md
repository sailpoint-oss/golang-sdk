# \AccessProfilesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessProfile**](AccessProfilesAPI.md#CreateAccessProfile) | **Post** /access-profiles | Create an Access Profile
[**DeleteAccessProfile**](AccessProfilesAPI.md#DeleteAccessProfile) | **Delete** /access-profiles/{id} | Delete the specified Access Profile
[**DeleteAccessProfilesInBulk**](AccessProfilesAPI.md#DeleteAccessProfilesInBulk) | **Post** /access-profiles/bulk-delete | Delete Access Profile(s)
[**GetAccessProfile**](AccessProfilesAPI.md#GetAccessProfile) | **Get** /access-profiles/{id} | Get an Access Profile
[**GetAccessProfileEntitlements**](AccessProfilesAPI.md#GetAccessProfileEntitlements) | **Get** /access-profiles/{id}/entitlements | List Access Profile&#39;s Entitlements
[**ListAccessProfiles**](AccessProfilesAPI.md#ListAccessProfiles) | **Get** /access-profiles | List Access Profiles
[**PatchAccessProfile**](AccessProfilesAPI.md#PatchAccessProfile) | **Patch** /access-profiles/{id} | Patch a specified Access Profile



## Create an Access Profile

> AccessProfile CreateAccessProfile(ctx).AccessProfile(accessProfile).Execute()

This API creates an Access Profile.
A token with API, ORG_ADMIN, ROLE_ADMIN, ROLE_SUBADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API. In addition, a token with only ROLE_SUBADMIN or SOURCE_SUBADMIN authority must be associated with the Access Profile's Source.
The maximum supported length for the description field is 2000 characters. Longer descriptions will be preserved for existing access profiles, however, any new access profiles as well as any updates to existing descriptions will be limited to 2000 characters.

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
    accessProfile := *sailpoint.NewAccessProfile("Employee-database-read-write", *sailpoint.NewOwnerReference(), *sailpoint.NewAccessProfileSourceRef()) // AccessProfile | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.CreateAccessProfile(context.Background()).AccessProfile(accessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.CreateAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.CreateAccessProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfile** | [**AccessProfile**](AccessProfile.md) |  | 

### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete the specified Access Profile

> DeleteAccessProfile(ctx, id).Execute()

This API deletes an existing Access Profile.

The Access Profile must not be in use, for example, Access Profile can not be deleted if they belong to an Application, Life Cycle State or a Role. If it is, a 400 error is returned.

A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to invoke this API. In addition, a SOURCE_SUBADMIN token must be able to administer the Source associated with the Access Profile.

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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the Access Profile to delete

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.AccessProfilesAPI.DeleteAccessProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.DeleteAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Profile to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessProfileRequest struct via the builder pattern


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


## Delete Access Profile(s)

> AccessProfileBulkDeleteResponse DeleteAccessProfilesInBulk(ctx).AccessProfileBulkDeleteRequest(accessProfileBulkDeleteRequest).Execute()

This API initiates a bulk deletion of one or more Access Profiles.

By default, if any of the indicated Access Profiles are in use, no deletions will be performed and the **inUse** field of the response indicates the usages that must be removed first. If the request field **bestEffortOnly** is **true**, however, usages are reported in the **inUse** response field but all other indicated Access Profiles will be deleted.

A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API. In addition, a SOURCE_SUBADMIN may only use this API to delete Access Profiles which are associated with Sources they are able to administer.

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
    accessProfileBulkDeleteRequest := *sailpoint.NewAccessProfileBulkDeleteRequest() // AccessProfileBulkDeleteRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.DeleteAccessProfilesInBulk(context.Background()).AccessProfileBulkDeleteRequest(accessProfileBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.DeleteAccessProfilesInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessProfilesInBulk`: AccessProfileBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.DeleteAccessProfilesInBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessProfilesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessProfileBulkDeleteRequest** | [**AccessProfileBulkDeleteRequest**](AccessProfileBulkDeleteRequest.md) |  | 

### Return type

[**AccessProfileBulkDeleteResponse**](AccessProfileBulkDeleteResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get an Access Profile

> AccessProfile GetAccessProfile(ctx, id).Execute()

This API returns an Access Profile by its ID.

A token with API, ORG_ADMIN, ROLE_ADMIN, ROLE_SUBADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Access Profile

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.GetAccessProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.GetAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.GetAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Access Profile's Entitlements

> []Entitlement GetAccessProfileEntitlements(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

This API lists the Entitlements associated with a given Access Profile

A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to invoke this API. In addition, a token with SOURCE_SUBADMIN authority must have access to the Source associated with the given Access Profile

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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the containing Access Profile
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "attribute eq "memberOf"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **attribute**: *eq, sw*  **value**: *eq, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **source.id**: *eq, in* (optional)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, attribute, value, created, modified** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.GetAccessProfileEntitlements(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.GetAccessProfileEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessProfileEntitlements`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.GetAccessProfileEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the containing Access Profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessProfileEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **attribute**: *eq, sw*  **value**: *eq, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **source.id**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, attribute, value, created, modified** | 

### Return type

[**[]Entitlement**](Entitlement.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Access Profiles

> []AccessProfile ListAccessProfiles(ctx).ForSubadmin(forSubadmin).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Execute()

This API returns a list of Access Profiles.

A token with API, ORG_ADMIN, ROLE_ADMIN, ROLE_SUBADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

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
    forSubadmin := "8c190e6787aa4ed9a90bd9d5344523fb" // string | If provided, filters the returned list according to what is visible to the indicated ROLE_SUBADMIN or SOURCE_SUBADMIN Identity. The value of the parameter is either an Identity ID, or the special value **me**, which is shorthand for the calling Identity's ID.  A 400 Bad Request error is returned if the **for-subadmin** parameter is specified for an Identity that is not a subadmin. (optional)
    limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name eq "SailPoint Support"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **requestable**: *eq*  **source.id**: *eq, in*  Composite operators supported: *and, or* (optional)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)
    forSegmentIds := "0b5c9f25-83c6-4762-9073-e38f7bb2ae26,2e8d8180-24bc-4d21-91c6-7affdb473b0d" // string | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs.  If segmentation is currently unavailable, specifying this parameter results in an error. (optional)
    includeUnsegmented := false // bool | Whether or not the response list should contain unsegmented Access Profiles. If *for-segment-ids* is absent or empty, specifying *include-unsegmented* as false results in an error. (optional) (default to true)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.ListAccessProfiles(context.Background()).ForSubadmin(forSubadmin).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.ListAccessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessProfiles`: []AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.ListAccessProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forSubadmin** | **string** | If provided, filters the returned list according to what is visible to the indicated ROLE_SUBADMIN or SOURCE_SUBADMIN Identity. The value of the parameter is either an Identity ID, or the special value **me**, which is shorthand for the calling Identity&#39;s ID.  A 400 Bad Request error is returned if the **for-subadmin** parameter is specified for an Identity that is not a subadmin. | 
 **limit** | **int32** | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 50]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **created**: *gt, lt, ge, le*  **modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **requestable**: *eq*  **source.id**: *eq, in*  Composite operators supported: *and, or* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** | 
 **forSegmentIds** | **string** | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs.  If segmentation is currently unavailable, specifying this parameter results in an error. | 
 **includeUnsegmented** | **bool** | Whether or not the response list should contain unsegmented Access Profiles. If *for-segment-ids* is absent or empty, specifying *include-unsegmented* as false results in an error. | [default to true]

### Return type

[**[]AccessProfile**](AccessProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch a specified Access Profile

> AccessProfile PatchAccessProfile(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

This API updates an existing Access Profile. The following fields are patchable:

**name**

**description**

**enabled**

**owner**

**requestable**

**accessRequestConfig**

**revokeRequestConfig**

**segments**

**entitlements**

**provisioningCriteria**

**source** (must be updated with entitlements belonging to new source in the same API call)

If you need to change the `source` of the access profile, you can do so only if you update the `entitlements` in the same API call.  The new entitlements can only come from the target source that you want to change to.  Look for the example "Replace Source" in the examples dropdown.

A token with API, ORG_ADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API. In addition, a SOURCE_SUBADMIN may only use this API to patch Access Profiles which are associated with Sources they are able to administer.
>  The maximum supported length for the description field is 2000 characters. Longer descriptions will be preserved for existing access profiles, however, any new access profiles as well as any updates to existing descriptions will be limited to 2000 characters.

> You can only add or replace **entitlements** that exist on the source that the access profile is attached to. You can use the **list entitlements** endpoint with the **filters** query parameter to get a list of available entitlements on the access profile's source.

>  Patching the value of the **requestable** field is only supported for customers enabled with the new Request Center. Otherwise, attempting to modify this field results in a 400 error.

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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the Access Profile to patch
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.AccessProfilesAPI.PatchAccessProfile(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesAPI.PatchAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesAPI.PatchAccessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Access Profile to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**AccessProfile**](AccessProfile.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

