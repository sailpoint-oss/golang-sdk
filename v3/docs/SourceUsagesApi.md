# \SourceUsagesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusBySourceId**](SourceUsagesAPI.md#GetStatusBySourceId) | **Get** /source-usages/{sourceId}/status | Finds status of source usage
[**GetUsagesBySourceId**](SourceUsagesAPI.md#GetUsagesBySourceId) | **Get** /source-usages/{sourceId}/summaries | Returns source usage insights



## GetStatusBySourceId

> SourceUsageStatus GetStatusBySourceId(ctx, sourceId).Execute()

Finds status of source usage



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | ID of IDN source

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceUsagesAPI.GetStatusBySourceId(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceUsagesAPI.GetStatusBySourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatusBySourceId`: SourceUsageStatus
    fmt.Fprintf(os.Stdout, "Response from `SourceUsagesAPI.GetStatusBySourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | ID of IDN source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusBySourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceUsageStatus**](SourceUsageStatus.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsagesBySourceId

> []SourceUsage GetUsagesBySourceId(ctx, sourceId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()

Returns source usage insights



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    sourceId := "2c9180835d191a86015d28455b4a2329" // string | ID of IDN source
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "-date" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **date** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceUsagesAPI.GetUsagesBySourceId(context.Background(), sourceId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceUsagesAPI.GetUsagesBySourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsagesBySourceId`: []SourceUsage
    fmt.Fprintf(os.Stdout, "Response from `SourceUsagesAPI.GetUsagesBySourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | ID of IDN source | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsagesBySourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **date** | 

### Return type

[**[]SourceUsage**](SourceUsage.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

