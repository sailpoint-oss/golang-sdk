# \PublicIdentitiesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublicIdentities**](PublicIdentitiesApi.md#GetPublicIdentities) | **Get** /public-identities | Get a list of public identities



## GetPublicIdentities

> []PublicIdentity GetPublicIdentities(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).AddCoreFilters(addCoreFilters).Sorters(sorters).Execute()

Get a list of public identities

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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "firstname eq "John"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **alias**: *eq, sw*  **email**: *eq, sw*  **firstname**: *eq, sw*  **lastname**: *eq, sw* (optional)
    addCoreFilters := false // bool | If *true*, only get identities which satisfy ALL the following criteria in addition to any criteria specified by *filters*:   - Should be either correlated or protected.   - Should not be \"spadmin\" or \"cloudadmin\".   - uid should not be null.   - lastname should not be null.   - email should not be null. (optional) (default to false)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PublicIdentitiesApi.GetPublicIdentities(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).AddCoreFilters(addCoreFilters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicIdentitiesApi.GetPublicIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicIdentities`: []PublicIdentity
    fmt.Fprintf(os.Stdout, "Response from `PublicIdentitiesApi.GetPublicIdentities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **alias**: *eq, sw*  **email**: *eq, sw*  **firstname**: *eq, sw*  **lastname**: *eq, sw* | 
 **addCoreFilters** | **bool** | If *true*, only get identities which satisfy ALL the following criteria in addition to any criteria specified by *filters*:   - Should be either correlated or protected.   - Should not be \&quot;spadmin\&quot; or \&quot;cloudadmin\&quot;.   - uid should not be null.   - lastname should not be null.   - email should not be null. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** | 

### Return type

[**[]PublicIdentity**](PublicIdentity.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

