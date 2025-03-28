# \IAIAccessRequestRecommendationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAccessRequestRecommendationsIgnoredItem**](IAIAccessRequestRecommendationsAPI.md#AddAccessRequestRecommendationsIgnoredItem) | **Post** /ai-access-request-recommendations/ignored-items | Ignore Access Request Recommendation
[**AddAccessRequestRecommendationsRequestedItem**](IAIAccessRequestRecommendationsAPI.md#AddAccessRequestRecommendationsRequestedItem) | **Post** /ai-access-request-recommendations/requested-items | Accept Access Request Recommendation
[**AddAccessRequestRecommendationsViewedItem**](IAIAccessRequestRecommendationsAPI.md#AddAccessRequestRecommendationsViewedItem) | **Post** /ai-access-request-recommendations/viewed-items | Mark Viewed Access Request Recommendations
[**AddAccessRequestRecommendationsViewedItems**](IAIAccessRequestRecommendationsAPI.md#AddAccessRequestRecommendationsViewedItems) | **Post** /ai-access-request-recommendations/viewed-items/bulk-create | Bulk Mark Viewed Access Request Recommendations
[**GetAccessRequestRecommendations**](IAIAccessRequestRecommendationsAPI.md#GetAccessRequestRecommendations) | **Get** /ai-access-request-recommendations | Identity Access Request Recommendations
[**GetAccessRequestRecommendationsConfig**](IAIAccessRequestRecommendationsAPI.md#GetAccessRequestRecommendationsConfig) | **Get** /ai-access-request-recommendations/config | Get Access Request Recommendations config
[**GetAccessRequestRecommendationsIgnoredItems**](IAIAccessRequestRecommendationsAPI.md#GetAccessRequestRecommendationsIgnoredItems) | **Get** /ai-access-request-recommendations/ignored-items | List Ignored Access Request Recommendations
[**GetAccessRequestRecommendationsRequestedItems**](IAIAccessRequestRecommendationsAPI.md#GetAccessRequestRecommendationsRequestedItems) | **Get** /ai-access-request-recommendations/requested-items | List Accepted Access Request Recommendations
[**GetAccessRequestRecommendationsViewedItems**](IAIAccessRequestRecommendationsAPI.md#GetAccessRequestRecommendationsViewedItems) | **Get** /ai-access-request-recommendations/viewed-items | List Viewed Access Request Recommendations
[**SetAccessRequestRecommendationsConfig**](IAIAccessRequestRecommendationsAPI.md#SetAccessRequestRecommendationsConfig) | **Put** /ai-access-request-recommendations/config | Update Access Request Recommendations config



## AddAccessRequestRecommendationsIgnoredItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsIgnoredItem(ctx).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Ignore Access Request Recommendation



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
	accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access item to ignore for an identity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsIgnoredItem(context.Background()).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsIgnoredItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAccessRequestRecommendationsIgnoredItem`: AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsIgnoredItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsIgnoredItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access item to ignore for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsRequestedItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsRequestedItem(ctx).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Accept Access Request Recommendation



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
	accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access item that was requested for an identity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsRequestedItem(context.Background()).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsRequestedItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAccessRequestRecommendationsRequestedItem`: AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsRequestedItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsRequestedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access item that was requested for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsViewedItem

> AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsViewedItem(ctx).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Mark Viewed Access Request Recommendations



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
	accessRequestRecommendationActionItemDto := *openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem()) // AccessRequestRecommendationActionItemDto | The recommended access that was viewed for an identity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItem(context.Background()).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAccessRequestRecommendationsViewedItem`: AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsViewedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **accessRequestRecommendationActionItemDto** | [**AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access that was viewed for an identity. | 

### Return type

[**AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAccessRequestRecommendationsViewedItems

> []AccessRequestRecommendationActionItemResponseDto AddAccessRequestRecommendationsViewedItems(ctx).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()

Bulk Mark Viewed Access Request Recommendations



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
	accessRequestRecommendationActionItemDto := []openapiclient.AccessRequestRecommendationActionItemDto{*openapiclient.NewAccessRequestRecommendationActionItemDto("2c91808570313110017040b06f344ec9", *openapiclient.NewAccessRequestRecommendationItem())} // []AccessRequestRecommendationActionItemDto | The recommended access items that were viewed for an identity.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItems(context.Background()).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationActionItemDto(accessRequestRecommendationActionItemDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAccessRequestRecommendationsViewedItems`: []AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAccessRequestRecommendationsViewedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **accessRequestRecommendationActionItemDto** | [**[]AccessRequestRecommendationActionItemDto**](AccessRequestRecommendationActionItemDto.md) | The recommended access items that were viewed for an identity. | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendations

> []AccessRequestRecommendationItemDetail GetAccessRequestRecommendations(ctx).XSailPointExperimental(xSailPointExperimental).IdentityId(identityId).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Filters(filters).Sorters(sorters).Execute()

Identity Access Request Recommendations



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
	identityId := "2c91808570313110017040b06f344ec9" // string | Get access request recommendations for an identityId. *me* indicates the current user. (optional) (default to "me")
	limit := int32(15) // int32 | Max number of results to return. (optional) (default to 15)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	includeTranslationMessages := false // bool | If *true* it will populate a list of translation messages in the response. (optional) (default to false)
	filters := "access.name co "admin"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.name**: *co*  **access.type**: *eq, in*  **access.description**: *co, eq, in* (optional)
	sorters := "access.name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name, access.type**  By default the recommendations are sorted by highest confidence first. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendations(context.Background()).XSailPointExperimental(xSailPointExperimental).IdentityId(identityId).Limit(limit).Offset(offset).Count(count).IncludeTranslationMessages(includeTranslationMessages).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestRecommendations`: []AccessRequestRecommendationItemDetail
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **identityId** | **string** | Get access request recommendations for an identityId. *me* indicates the current user. | [default to &quot;me&quot;]
 **limit** | **int32** | Max number of results to return. | [default to 15]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **includeTranslationMessages** | **bool** | If *true* it will populate a list of translation messages in the response. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.name**: *co*  **access.type**: *eq, in*  **access.description**: *co, eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name, access.type**  By default the recommendations are sorted by highest confidence first. | 

### Return type

[**[]AccessRequestRecommendationItemDetail**](AccessRequestRecommendationItemDetail.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsConfig

> AccessRequestRecommendationConfigDto GetAccessRequestRecommendationsConfig(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

Get Access Request Recommendations config



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestRecommendationsConfig`: AccessRequestRecommendationConfigDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**AccessRequestRecommendationConfigDto**](AccessRequestRecommendationConfigDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsIgnoredItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsIgnoredItems(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Ignored Access Request Recommendations



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
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "identityId eq "2c9180846b0a0583016b299f210c1314"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
	sorters := "access.id" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsIgnoredItems(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsIgnoredItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestRecommendationsIgnoredItems`: []AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsIgnoredItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsIgnoredItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsRequestedItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsRequestedItems(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Accepted Access Request Recommendations



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
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "access.id eq "2c9180846b0a0583016b299f210c1314"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
	sorters := "access.id" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsRequestedItems(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsRequestedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestRecommendationsRequestedItems`: []AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsRequestedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsRequestedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessRequestRecommendationsViewedItems

> []AccessRequestRecommendationActionItemResponseDto GetAccessRequestRecommendationsViewedItems(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Viewed Access Request Recommendations



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
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
	filters := "access.id eq "2c9180846b0a0583016b299f210c1314"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* (optional)
	sorters := "access.id" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsViewedItems(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsViewedItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessRequestRecommendationsViewedItems`: []AccessRequestRecommendationActionItemResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsViewedItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestRecommendationsViewedItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **access.id**: *eq, in*  **access.type**: *eq, in*  **identityId**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.id, access.type, identityId, timestamp** | 

### Return type

[**[]AccessRequestRecommendationActionItemResponseDto**](AccessRequestRecommendationActionItemResponseDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAccessRequestRecommendationsConfig

> AccessRequestRecommendationConfigDto SetAccessRequestRecommendationsConfig(ctx).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationConfigDto(accessRequestRecommendationConfigDto).Execute()

Update Access Request Recommendations config



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
	accessRequestRecommendationConfigDto := *openapiclient.NewAccessRequestRecommendationConfigDto(float32(0.5)) // AccessRequestRecommendationConfigDto | The desired configurations for Access Request Recommender for the tenant.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIAccessRequestRecommendationsAPI.SetAccessRequestRecommendationsConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).AccessRequestRecommendationConfigDto(accessRequestRecommendationConfigDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIAccessRequestRecommendationsAPI.SetAccessRequestRecommendationsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetAccessRequestRecommendationsConfig`: AccessRequestRecommendationConfigDto
	fmt.Fprintf(os.Stdout, "Response from `IAIAccessRequestRecommendationsAPI.SetAccessRequestRecommendationsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessRequestRecommendationsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **accessRequestRecommendationConfigDto** | [**AccessRequestRecommendationConfigDto**](AccessRequestRecommendationConfigDto.md) | The desired configurations for Access Request Recommender for the tenant. | 

### Return type

[**AccessRequestRecommendationConfigDto**](AccessRequestRecommendationConfigDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

