# \EntitlementsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsBulkUpdate**](EntitlementsApi.md#EntitlementsBulkUpdate) | **Post** /entitlements/bulk-update | Bulk update an entitlement list
[**GetEntitlement**](EntitlementsApi.md#GetEntitlement) | **Get** /entitlements/{id} | Get an Entitlement
[**ListEntitlementParents**](EntitlementsApi.md#ListEntitlementParents) | **Get** /entitlements/{id}/parents | List of entitlements parents
[**ListEntitlementchildren**](EntitlementsApi.md#ListEntitlementchildren) | **Get** /entitlements/{id}/children | List of entitlements children
[**ListEntitlements**](EntitlementsApi.md#ListEntitlements) | **Get** /entitlements | Gets a list of entitlements.
[**PatchEntitlement**](EntitlementsApi.md#PatchEntitlement) | **Patch** /entitlements | Patch a specified Entitlement



## EntitlementsBulkUpdate

> EntitlementsBulkUpdate(ctx).EntitlementBulkUpdateRequest(entitlementBulkUpdateRequest).Execute()

Bulk update an entitlement list



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
    entitlementBulkUpdateRequest := *openapiclient.NewEntitlementBulkUpdateRequest([]string{"EntitlementIds_example"}, []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")}) // EntitlementBulkUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.EntitlementsBulkUpdate(context.Background()).EntitlementBulkUpdateRequest(entitlementBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.EntitlementsBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsBulkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementBulkUpdateRequest** | [**EntitlementBulkUpdateRequest**](EntitlementBulkUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlement

> Entitlement GetEntitlement(ctx, id).Execute()

Get an Entitlement



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
    id := "2c91808874ff91550175097daaec161c" // string | Entitlement Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.GetEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlement`: Entitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Entitlement**](Entitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementParents

> []Entitlement ListEntitlementParents(ctx, id).Limit(limit).Offset(offset).Count(count).Execute()

List of entitlements parents



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
    id := "2c91808c74ff913f0175097daa9d59cd" // string | Entitlement Id
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.ListEntitlementParents(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListEntitlementParents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementParents`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListEntitlementParents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementParentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

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


## ListEntitlementchildren

> []Entitlement ListEntitlementchildren(ctx, id).Limit(limit).Offset(offset).Count(count).Execute()

List of entitlements children



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
    id := "2c91808874ff91550175097daaec161c" // string | Entitlement Id
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.ListEntitlementchildren(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListEntitlementchildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlementchildren`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListEntitlementchildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entitlement Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementchildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

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


## ListEntitlements

> []Entitlement ListEntitlements(ctx).AccountId(accountId).SegmentedForIdentity(segmentedForIdentity).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Cursor(cursor).Execute()

Gets a list of entitlements.



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
    accountId := "ef38f94347e94562b5bb8424a56397d8" // string | The account ID. If specified, returns only entitlements associated with the given Account. Can not be specified with the **filters**, **segmented-for-identity**, **for-segment-ids**, or **include-unsegmented** param(s). (optional)
    segmentedForIdentity := "me" // string | If present and not empty, additionally filters Entitlements to those which are assigned to the Segment(s) which are visible to the Identity with the specified ID. By convention, the value **me** can stand in for the current user's Identity ID. Can not be specified with the **account-id** or **for-segment-ids** param(s). It is also illegal to specify a value that refers to a different user's Identity. (optional)
    forSegmentIds := "041727d4-7d95-4779-b891-93cf41e98249,a378c9fa-bae5-494c-804e-a1e30f69f649" // string | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs. Can not be specified with the **account-id** or **segmented-for-identity** param(s). (optional)
    includeUnsegmented := true // bool | Whether or not the response list should contain unsegmented Entitlements. If **for-segment-ids** and **segmented-for-identity** are both absent or empty, specifying **include-unsegmented=false** results in an error. (optional) (default to true)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name,-modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **id, name, created, modified, type, attribute, value, source.id** (optional)
    filters := "attribute eq "memberOf"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  **type**: *eq, in*  **attribute**: *eq, in*  **value**: *eq, in, sw*  **source.id**: *eq, in*  **requestable**: *eq* (optional)
    cursor := "9602877a-1c11-4c02-a399-7273ea2158da" // string | **WARNING: This parameter is deprecated and support will be removed in an upcoming release.**  This endpoint supports cursor-based pagination as an alternative to pagination using **limit** and **offset**. Use of cursors is optional but may deliver significantly better performance when paginating over large result sets, particularly when sorting is involved.  To obtain an initial cursor value, the initial request must be made with **offset=0** and **count=false**, or by omitting those parameters. Other query parameters may be specified on the initial request, to specify filtering and sorting for example. The response to this initial request will provide a cursor value in the **Slpt-Cursor** response header. Thereafter, passing in the last returned value of **Slpt-Cursor** as the value of the **cursor** parameter on the following request will return the next page of results, and so on. The end of the result set is signalled when a request returns without providing a value for **Slpt-Cursor**.  The **cursor** param is generally incompatible with other parameters; the second and subsequent requests should only pass that parameter and no others. A 400 Bad Request will be returned if any incompatible parameter values are passed along with the **cursor**.  Cursors expire after 15 minutes; a 400 Bad Request is returned if an expired value is provided for **cursor**. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.ListEntitlements(context.Background()).AccountId(accountId).SegmentedForIdentity(segmentedForIdentity).ForSegmentIds(forSegmentIds).IncludeUnsegmented(includeUnsegmented).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntitlements`: []Entitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account ID. If specified, returns only entitlements associated with the given Account. Can not be specified with the **filters**, **segmented-for-identity**, **for-segment-ids**, or **include-unsegmented** param(s). | 
 **segmentedForIdentity** | **string** | If present and not empty, additionally filters Entitlements to those which are assigned to the Segment(s) which are visible to the Identity with the specified ID. By convention, the value **me** can stand in for the current user&#39;s Identity ID. Can not be specified with the **account-id** or **for-segment-ids** param(s). It is also illegal to specify a value that refers to a different user&#39;s Identity. | 
 **forSegmentIds** | **string** | If present and not empty, additionally filters Access Profiles to those which are assigned to the Segment(s) with the specified IDs. Can not be specified with the **account-id** or **segmented-for-identity** param(s). | 
 **includeUnsegmented** | **bool** | Whether or not the response list should contain unsegmented Entitlements. If **for-segment-ids** and **segmented-for-identity** are both absent or empty, specifying **include-unsegmented&#x3D;false** results in an error. | [default to true]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **id, name, created, modified, type, attribute, value, source.id** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, in, sw*  **type**: *eq, in*  **attribute**: *eq, in*  **value**: *eq, in, sw*  **source.id**: *eq, in*  **requestable**: *eq* | 
 **cursor** | **string** | **WARNING: This parameter is deprecated and support will be removed in an upcoming release.**  This endpoint supports cursor-based pagination as an alternative to pagination using **limit** and **offset**. Use of cursors is optional but may deliver significantly better performance when paginating over large result sets, particularly when sorting is involved.  To obtain an initial cursor value, the initial request must be made with **offset&#x3D;0** and **count&#x3D;false**, or by omitting those parameters. Other query parameters may be specified on the initial request, to specify filtering and sorting for example. The response to this initial request will provide a cursor value in the **Slpt-Cursor** response header. Thereafter, passing in the last returned value of **Slpt-Cursor** as the value of the **cursor** parameter on the following request will return the next page of results, and so on. The end of the result set is signalled when a request returns without providing a value for **Slpt-Cursor**.  The **cursor** param is generally incompatible with other parameters; the second and subsequent requests should only pass that parameter and no others. A 400 Bad Request will be returned if any incompatible parameter values are passed along with the **cursor**.  Cursors expire after 15 minutes; a 400 Bad Request is returned if an expired value is provided for **cursor**. | 

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


## PatchEntitlement

> Entitlement PatchEntitlement(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a specified Entitlement



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
    id := "2c91808a7813090a017814121e121518" // string | ID of the Entitlement to patch
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntitlementsApi.PatchEntitlement(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.PatchEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchEntitlement`: Entitlement
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.PatchEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Entitlement to patch | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**Entitlement**](Entitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

