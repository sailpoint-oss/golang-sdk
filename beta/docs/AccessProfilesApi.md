# \AccessProfilesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessProfile**](AccessProfilesApi.md#CreateAccessProfile) | **Post** /access-profiles | Create an Access Profile
[**DeleteAccessProfile**](AccessProfilesApi.md#DeleteAccessProfile) | **Delete** /access-profiles/{id} | Delete the specified Access Profile
[**DeleteAccessProfilesInBulk**](AccessProfilesApi.md#DeleteAccessProfilesInBulk) | **Post** /access-profiles/bulk-delete | Delete Access Profile(s)
[**GetAccessProfile**](AccessProfilesApi.md#GetAccessProfile) | **Get** /access-profiles/{id} | Get an Access Profile
[**GetAccessProfileEntitlements**](AccessProfilesApi.md#GetAccessProfileEntitlements) | **Get** /access-profiles/{id}/entitlements | List Access Profile&#39;s Entitlements
[**ListAccessProfiles**](AccessProfilesApi.md#ListAccessProfiles) | **Get** /access-profiles | List Access Profiles
[**PatchAccessProfile**](AccessProfilesApi.md#PatchAccessProfile) | **Patch** /access-profiles/{id} | Patch a specified Access Profile



## CreateAccessProfile

> AccessProfile CreateAccessProfile(ctx).AccessProfile(accessProfile).Execute()

Create an Access Profile



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
    accessProfile := *openapiclient.NewAccessProfile("Employee-database-read-write", *openapiclient.NewOwnerReference(), *openapiclient.NewAccessProfileSourceRef()) // AccessProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.CreateAccessProfile(context.Background()).AccessProfile(accessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.CreateAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.CreateAccessProfile`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessProfile

> DeleteAccessProfile(ctx, id).Execute()

Delete the specified Access Profile



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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the Access Profile to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.DeleteAccessProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.DeleteAccessProfile``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessProfilesInBulk

> AccessProfileBulkDeleteResponse DeleteAccessProfilesInBulk(ctx).AccessProfileBulkDeleteRequest(accessProfileBulkDeleteRequest).Execute()

Delete Access Profile(s)



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
    accessProfileBulkDeleteRequest := *openapiclient.NewAccessProfileBulkDeleteRequest() // AccessProfileBulkDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.DeleteAccessProfilesInBulk(context.Background()).AccessProfileBulkDeleteRequest(accessProfileBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.DeleteAccessProfilesInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessProfilesInBulk`: AccessProfileBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.DeleteAccessProfilesInBulk`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessProfile

> AccessProfile GetAccessProfile(ctx, id).Execute()

Get an Access Profile



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
    id := "2c9180837ca6693d017ca8d097500149" // string | ID of the Access Profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.GetAccessProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.GetAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.GetAccessProfile`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessProfileEntitlements

> []Entitlement GetAccessProfileEntitlements(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Access Profile's Entitlements



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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the containing Access Profile
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "attribute eq "memberOf"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Entitlement fields and operators: **id**: *eq, in*  **name**: *eq, sw*  **attribute**: *eq, sw*  **value**: *eq, sw*  **created, modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **source.id**: *eq, in* (optional)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, attribute, value, created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.GetAccessProfileEntitlements(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.GetAccessProfileEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessProfileEntitlements`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.GetAccessProfileEntitlements`: %v\n", resp)
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following Entitlement fields and operators: **id**: *eq, in*  **name**: *eq, sw*  **attribute**: *eq, sw*  **value**: *eq, sw*  **created, modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **source.id**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, attribute, value, created, modified** | 

### Return type

[**[]Entitlement**](Entitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessProfiles

> []AccessProfile ListAccessProfiles(ctx).ForSubadmin(forSubadmin).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Execute()

List Access Profiles



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
    forSubadmin := "8c190e6787aa4ed9a90bd9d5344523fb" // string | If provided, filters the returned list according to what is visible to the indicated ROLE_SUBADMIN or SOURCE_SUBADMIN Identity. The value of the parameter is either an Identity ID, or the special value **me**, which is shorthand for the calling Identity's ID.  A 400 Bad Request error is returned if the **for-subadmin** parameter is specified for an Identity that is not a subadmin. (optional)
    limit := int32(50) // int32 | Note that for this API the maximum value for limit is 50. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 50)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name eq "SailPoint Support"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **created, modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **requestable**: *eq*  **source.id**: *eq, in* (optional)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)
    forSegmentIds := "0b5c9f25-83c6-4762-9073-e38f7bb2ae26,2e8d8180-24bc-4d21-91c6-7affdb473b0d" // string | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs.  If segmentation is currently unavailable, specifying this parameter results in an error. (optional)
    includeUnsegmented := false // bool | Whether or not the response list should contain unsegmented Access Profiles. If *for-segment-ids* is absent or empty, specifying *include-unsegmented* as false results in an error. (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.ListAccessProfiles(context.Background()).ForSubadmin(forSubadmin).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.ListAccessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessProfiles`: []AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.ListAccessProfiles`: %v\n", resp)
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **created, modified**: *gt, lt, ge, le*  **owner.id**: *eq, in*  **requestable**: *eq*  **source.id**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** | 
 **forSegmentIds** | **string** | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs.  If segmentation is currently unavailable, specifying this parameter results in an error. | 
 **includeUnsegmented** | **bool** | Whether or not the response list should contain unsegmented Access Profiles. If *for-segment-ids* is absent or empty, specifying *include-unsegmented* as false results in an error. | [default to true]

### Return type

[**[]AccessProfile**](AccessProfile.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccessProfile

> AccessProfile PatchAccessProfile(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a specified Access Profile



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
    id := "2c91808a7813090a017814121919ecca" // string | ID of the Access Profile to patch
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessProfilesApi.PatchAccessProfile(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessProfilesApi.PatchAccessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAccessProfile`: AccessProfile
    fmt.Fprintf(os.Stdout, "Response from `AccessProfilesApi.PatchAccessProfile`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

