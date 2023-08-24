# \IAIRecommendationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessageCatalogs**](IAIRecommendationsApi.md#GetMessageCatalogs) | **Get** /translation-catalogs/{catalog-id} | Get Message catalogs
[**GetRecommendations**](IAIRecommendationsApi.md#GetRecommendations) | **Post** /recommendations/request | Returns a Recommendation Based on Object
[**GetRecommendationsConfig**](IAIRecommendationsApi.md#GetRecommendationsConfig) | **Get** /recommendations/config | Get certification recommendation config values
[**UpdateRecommendationsConfig**](IAIRecommendationsApi.md#UpdateRecommendationsConfig) | **Put** /recommendations/config | Update certification recommendation config values



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
    resp, r, err := apiClient.IAIRecommendationsApi.GetMessageCatalogs(context.Background(), catalogId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsApi.GetMessageCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageCatalogs`: []MessageCatalogDto
    fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsApi.GetMessageCatalogs`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendations

> RecommendationResponseDto GetRecommendations(ctx).RecommendationRequestDto(recommendationRequestDto).Execute()

Returns a Recommendation Based on Object



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
    recommendationRequestDto := *openapiclient.NewRecommendationRequestDto() // RecommendationRequestDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRecommendationsApi.GetRecommendations(context.Background()).RecommendationRequestDto(recommendationRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsApi.GetRecommendations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendations`: RecommendationResponseDto
    fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsApi.GetRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendationRequestDto** | [**RecommendationRequestDto**](RecommendationRequestDto.md) |  | 

### Return type

[**RecommendationResponseDto**](RecommendationResponseDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecommendationsConfig

> RecommendationConfigDto GetRecommendationsConfig(ctx).Execute()

Get certification recommendation config values



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRecommendationsApi.GetRecommendationsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsApi.GetRecommendationsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecommendationsConfig`: RecommendationConfigDto
    fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsApi.GetRecommendationsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsConfigRequest struct via the builder pattern


### Return type

[**RecommendationConfigDto**](RecommendationConfigDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecommendationsConfig

> RecommendationConfigDto UpdateRecommendationsConfig(ctx).RecommendationConfigDto(recommendationConfigDto).Execute()

Update certification recommendation config values



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
    recommendationConfigDto := *openapiclient.NewRecommendationConfigDto() // RecommendationConfigDto | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRecommendationsApi.UpdateRecommendationsConfig(context.Background()).RecommendationConfigDto(recommendationConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsApi.UpdateRecommendationsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecommendationsConfig`: RecommendationConfigDto
    fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsApi.UpdateRecommendationsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecommendationsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendationConfigDto** | [**RecommendationConfigDto**](RecommendationConfigDto.md) |  | 

### Return type

[**RecommendationConfigDto**](RecommendationConfigDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

