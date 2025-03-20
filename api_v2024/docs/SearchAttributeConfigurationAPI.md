# \SearchAttributeConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#CreateSearchAttributeConfig) | **Post** /accounts/search-attribute-config | Create Extended Search Attributes
[**DeleteSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#DeleteSearchAttributeConfig) | **Delete** /accounts/search-attribute-config/{name} | Delete Extended Search Attribute
[**GetSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#GetSearchAttributeConfig) | **Get** /accounts/search-attribute-config | List Extended Search Attributes
[**GetSingleSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#GetSingleSearchAttributeConfig) | **Get** /accounts/search-attribute-config/{name} | Get Extended Search Attribute
[**PatchSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#PatchSearchAttributeConfig) | **Patch** /accounts/search-attribute-config/{name} | Update Extended Search Attribute



## CreateSearchAttributeConfig

> map[string]interface{} CreateSearchAttributeConfig(ctx).XSailPointExperimental(xSailPointExperimental).SearchAttributeConfig(searchAttributeConfig).Execute()

Create Extended Search Attributes



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
	searchAttributeConfig := *openapiclient.NewSearchAttributeConfig() // SearchAttributeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.CreateSearchAttributeConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).SearchAttributeConfig(searchAttributeConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.CreateSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSearchAttributeConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.CreateSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **searchAttributeConfig** | [**SearchAttributeConfig**](SearchAttributeConfig.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSearchAttributeConfig

> DeleteSearchAttributeConfig(ctx, name).XSailPointExperimental(xSailPointExperimental).Execute()

Delete Extended Search Attribute



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
	name := "newMailAttribute" // string | Name of the extended search attribute configuration to delete.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAttributeConfigurationAPI.DeleteSearchAttributeConfig(context.Background(), name).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.DeleteSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the extended search attribute configuration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchAttributeConfig

> []SearchAttributeConfig GetSearchAttributeConfig(ctx).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Execute()

List Extended Search Attributes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.GetSearchAttributeConfig(context.Background()).XSailPointExperimental(xSailPointExperimental).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.GetSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchAttributeConfig`: []SearchAttributeConfig
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.GetSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleSearchAttributeConfig

> SearchAttributeConfig GetSingleSearchAttributeConfig(ctx, name).XSailPointExperimental(xSailPointExperimental).Execute()

Get Extended Search Attribute



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
	name := "newMailAttribute" // string | Name of the extended search attribute configuration to get.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig(context.Background(), name).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleSearchAttributeConfig`: SearchAttributeConfig
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the extended search attribute configuration to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSearchAttributeConfig

> SearchAttributeConfig PatchSearchAttributeConfig(ctx, name).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()

Update Extended Search Attribute



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
	name := "promotedMailAttribute" // string | Name of the search attribute configuration to patch.
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.PatchSearchAttributeConfig(context.Background(), name).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.PatchSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchSearchAttributeConfig`: SearchAttributeConfig
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.PatchSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the search attribute configuration to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

