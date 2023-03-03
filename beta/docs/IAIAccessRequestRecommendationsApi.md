# \IAIAccessRequestRecommendationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessRequestRecommendationsIgnoredItem**](IAIAccessRequestRecommendationsApi.md#AddAccessRequestRecommendationsIgnoredItem) | **Post** /ai-access-request-recommendations/ignored-items | Notification of Ignored Access Request Recommendations
[**AddAccessRequestRecommendationsRequestedItem**](IAIAccessRequestRecommendationsApi.md#AddAccessRequestRecommendationsRequestedItem) | **Post** /ai-access-request-recommendations/requested-items | Notification of Requested Access Request Recommendations
[**AddAccessRequestRecommendationsViewedItem**](IAIAccessRequestRecommendationsApi.md#AddAccessRequestRecommendationsViewedItem) | **Post** /ai-access-request-recommendations/viewed-items | Notification of Viewed Access Request Recommendations
[**AddAccessRequestRecommendationsViewedItems**](IAIAccessRequestRecommendationsApi.md#AddAccessRequestRecommendationsViewedItems) | **Post** /ai-access-request-recommendations/viewed-items/bulk-create | Notification of Viewed Access Request Recommendations in Bulk
[**GetAccessRequestRecommendations**](IAIAccessRequestRecommendationsApi.md#GetAccessRequestRecommendations) | **Get** /ai-access-request-recommendations | Identity Access Request Recommendations
[**GetAccessRequestRecommendationsIgnoredItems**](IAIAccessRequestRecommendationsApi.md#GetAccessRequestRecommendationsIgnoredItems) | **Get** /ai-access-request-recommendations/ignored-items | List of Ignored Access Request Recommendations
[**GetAccessRequestRecommendationsRequestedItems**](IAIAccessRequestRecommendationsApi.md#GetAccessRequestRecommendationsRequestedItems) | **Get** /ai-access-request-recommendations/requested-items | List of Requested Access Request Recommendations
[**GetAccessRequestRecommendationsViewedItems**](IAIAccessRequestRecommendationsApi.md#GetAccessRequestRecommendationsViewedItems) | **Get** /ai-access-request-recommendations/viewed-items | List of Viewed Access Request Recommendations
[**GetMessageCatalogs**](IAIAccessRequestRecommendationsApi.md#GetMessageCatalogs) | **Get** /translation-catalogs/{catalog-id} | Get Message catalogs



## AddAccessRequestRecommendationsIgnoredItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsIgnoredItem(ctx).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Notification of Ignored Access Request Recommendations



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
    accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access item to ignore for an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsIgnoredItem(context.Background()).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsIgnoredItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessRequestRecommendationsIgnoredItem`: AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsIgnoredItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsIgnoredItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access item to ignore for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsRequestedItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsRequestedItem(ctx).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Notification of Requested Access Request Recommendations



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
    accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access item that was requested for an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsRequestedItem(context.Background()).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsRequestedItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessRequestRecommendationsRequestedItem`: AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsRequestedItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsRequestedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access item that was requested for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsViewedItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsViewedItem(ctx).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Notification of Viewed Access Request Recommendations



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
    accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access that was viewed for an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItem(context.Background()).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessRequestRecommendationsViewedItem`: AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsViewedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access that was viewed for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsViewedItems

> []AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsViewedItems(ctx).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Notification of Viewed Access Request Recommendations in Bulk



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
    accessRequestRecommendationActionItemDto := []openapiclient.AccessRequestRecommendationActionItemDto{*openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem())} // []AccessRequestRecommendationActionItemDto | The recommended access items that were viewed for an identity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItems(context.Background()).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAccessRequestRecommendationsViewedItems`: []AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.AddAccessRequestRecommendationsViewedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsViewedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestRecommendationActionItemDto** | [**[]AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access items that were viewed for an identity. | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendations

> []AccessRequestRecommendationItemDetail GetAccessRequestRecommendations(ctx).IdentityId(identityId).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Filters(filters).Sorters(sorters).Execute()

Identity Access Request Recommendations



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
    identityId := "2c91808570313110017040b06f344ec9" // string | Get access request recommendations for an identityId. *me* indicates the current user. (optional) (default to "me")
    limit := int32(56) // int32 | Max number of results to return. (optional) (default to 15)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    includeTranslationMessages := false // bool | If *true* it will populate a list of translation messages in the response. (optional) (default to false)
    filters := "filters_example" // string | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.name**: *co*  **access.type**: *eq, in*  **access.description**: *co* (optional)
    sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name, access.type**  By default the recommendations are sorted by highest confidence first. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendations(context.Background()).IdentityId(identityId).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestRecommendations`: []AccessRequestRecommendationItemDetail
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityId** | **string** | Get access request recommendations for an identityId. *me* indicates the current user. | [default to &quot;me&quot;]
 **limit** | **int32** | Max number of results to return. | [default to 15]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **includeTranslationMessages** | **bool** | If *true* it will populate a list of translation messages in the response. | [default to false]
 **filters** | **string** | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.name**: *co*  **access.type**: *eq, in*  **access.description**: *co* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name, access.type**  By default the recommendations are sorted by highest confidence first. | 

### Return type

[**[]AccessRequestRecommendationItemDetail**](AccessRequestRecommendationItemDetail.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsIgnoredItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsIgnoredItems(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List of Ignored Access Request Recommendations



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
    filters := "filters_example" // string | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
    sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsIgnoredItems(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsIgnoredItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestRecommendationsIgnoredItems`: []AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsIgnoredItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsIgnoredItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsRequestedItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsRequestedItems(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List of Requested Access Request Recommendations



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
    filters := "filters_example" // string | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
    sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsRequestedItems(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsRequestedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestRecommendationsRequestedItems`: []AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsRequestedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsRequestedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsViewedItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsViewedItems(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List of Viewed Access Request Recommendations



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
    filters := "filters_example" // string | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
    sorters := "sorters_example" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsViewedItems(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsViewedItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestRecommendationsViewedItems`: []AccessRequestRecommendationActionItemResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.GetAccessRequestRecommendationsViewedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsViewedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter recommendations using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageCatalogs

> []MessageCatalogDto GetMessageCatalogs(ctx, catalogId).Execute()

Get Message catalogs



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
    catalogId := "catalogId_example" // string | The ID of the message catalog.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIAccessRequestRecommendationsApi.GetMessageCatalogs(context.Background(), catalogId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsApi.GetMessageCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageCatalogs`: []MessageCatalogDto
    fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsApi.GetMessageCatalogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **string** | The ID of the message catalog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageCatalogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MessageCatalogDto**](MessageCatalogDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

