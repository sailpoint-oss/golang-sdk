# \IAIRecommendationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecommendations**](IAIRecommendationsAPI.md#GetRecommendations) | **Post** /recommendations/request | Returns Recommendation Based on Object
[**GetRecommendationsConfig**](IAIRecommendationsAPI.md#GetRecommendationsConfig) | **Get** /recommendations/config | Get certification recommendation config values
[**UpdateRecommendationsConfig**](IAIRecommendationsAPI.md#UpdateRecommendationsConfig) | **Put** /recommendations/config | Update certification recommendation config values



## GetRecommendations

> RecommendationResponseDto GetRecommendations(ctx).RecommendationRequestDto(recommendationRequestDto).Execute()

Returns Recommendation Based on Object



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
	recommendationRequestDto := *openapiclient.NewRecommendationRequestDto() // RecommendationRequestDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIRecommendationsAPI.GetRecommendations(context.Background()).RecommendationRequestDto(recommendationRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsAPI.GetRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendations`: RecommendationResponseDto
	fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsAPI.GetRecommendations`: %v\n", resp)
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIRecommendationsAPI.GetRecommendationsConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsAPI.GetRecommendationsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecommendationsConfig`: RecommendationConfigDto
	fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsAPI.GetRecommendationsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecommendationsConfigRequest struct via the builder pattern


### Return type

[**RecommendationConfigDto**](RecommendationConfigDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	recommendationConfigDto := *openapiclient.NewRecommendationConfigDto() // RecommendationConfigDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IAIRecommendationsAPI.UpdateRecommendationsConfig(context.Background()).RecommendationConfigDto(recommendationConfigDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IAIRecommendationsAPI.UpdateRecommendationsConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRecommendationsConfig`: RecommendationConfigDto
	fmt.Fprintf(os.Stdout, "Response from `IAIRecommendationsAPI.UpdateRecommendationsConfig`: %v\n", resp)
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

