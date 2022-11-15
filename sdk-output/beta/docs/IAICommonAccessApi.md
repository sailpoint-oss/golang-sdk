# \IAICommonAccessApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommonAccessBulkUpdateStatus**](IAICommonAccessApi.md#CommonAccessBulkUpdateStatus) | **Post** /common-access/update-status | Bulk update common access status
[**CreateCommonAccess**](IAICommonAccessApi.md#CreateCommonAccess) | **Post** /common-access | Create common access items
[**GetCommonAccess**](IAICommonAccessApi.md#GetCommonAccess) | **Get** /common-access | Get a paginated list of common access



## CommonAccessBulkUpdateStatus

> map[string]interface{} CommonAccessBulkUpdateStatus(ctx).CommonAccessIDStatus(commonAccessIDStatus).Execute()

Bulk update common access status



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
    commonAccessIDStatus := []openapiclient.CommonAccessIDStatus{*openapiclient.NewCommonAccessIDStatus()} // []CommonAccessIDStatus | Confirm or deny in bulk the common access ids that are (or aren't) common access

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAICommonAccessApi.CommonAccessBulkUpdateStatus(context.Background()).CommonAccessIDStatus(commonAccessIDStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAICommonAccessApi.CommonAccessBulkUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommonAccessBulkUpdateStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IAICommonAccessApi.CommonAccessBulkUpdateStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonAccessBulkUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonAccessIDStatus** | [**[]CommonAccessIDStatus**](CommonAccessIDStatus.md) | Confirm or deny in bulk the common access ids that are (or aren&#39;t) common access | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCommonAccess

> CommonAccessItemResponse CreateCommonAccess(ctx).CommonAccessItemRequest(commonAccessItemRequest).Execute()

Create common access items



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
    commonAccessItemRequest := *openapiclient.NewCommonAccessItemRequest() // CommonAccessItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAICommonAccessApi.CreateCommonAccess(context.Background()).CommonAccessItemRequest(commonAccessItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAICommonAccessApi.CreateCommonAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCommonAccess`: CommonAccessItemResponse
    fmt.Fprintf(os.Stdout, "Response from `IAICommonAccessApi.CreateCommonAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommonAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonAccessItemRequest** | [**CommonAccessItemRequest**](CommonAccessItemRequest.md) |  | 

### Return type

[**CommonAccessItemResponse**](CommonAccessItemResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommonAccess

> []CommonAccessResponse GetCommonAccess(ctx).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()

Get a paginated list of common access



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "filters_example" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)  Filtering is supported for the following fields and operators:  **status**: *eq* \"CONFIRMED\" or \"DENIED\" **reviewedByUser** *eq* true or false **access.id**: *eq* \"id\" **access.type**: *eq* \"ROLE\" or \"ACCESS_PROFILE\" **access.name**: *sw* \"Administrator\" **access.description**: *sw* \"admin\" (optional)
    sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)  Sorting is supported for the following fields: **access.name,status**  By default the common access items are sorted by name, ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAICommonAccessApi.GetCommonAccess(context.Background()).Offset(offset).Limit(limit).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAICommonAccessApi.GetCommonAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommonAccess`: []CommonAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `IAICommonAccessApi.GetCommonAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommonAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)  Filtering is supported for the following fields and operators:  **status**: *eq* \&quot;CONFIRMED\&quot; or \&quot;DENIED\&quot; **reviewedByUser** *eq* true or false **access.id**: *eq* \&quot;id\&quot; **access.type**: *eq* \&quot;ROLE\&quot; or \&quot;ACCESS_PROFILE\&quot; **access.name**: *sw* \&quot;Administrator\&quot; **access.description**: *sw* \&quot;admin\&quot; | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)  Sorting is supported for the following fields: **access.name,status**  By default the common access items are sorted by name, ascending. | 

### Return type

[**[]CommonAccessResponse**](CommonAccessResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

