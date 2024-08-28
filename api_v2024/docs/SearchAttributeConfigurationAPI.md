# \SearchAttributeConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#CreateSearchAttributeConfig) | **Post** /accounts/search-attribute-config | Configure/create search attributes in IdentityNow.
[**DeleteSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#DeleteSearchAttributeConfig) | **Delete** /accounts/search-attribute-config/{name} | Delete search attribute in IdentityNow.
[**GetSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#GetSearchAttributeConfig) | **Get** /accounts/search-attribute-config | Retrieve attribute list in IdentityNow.
[**GetSingleSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#GetSingleSearchAttributeConfig) | **Get** /accounts/search-attribute-config/{name} | Get specific attribute in IdentityNow.
[**PatchSearchAttributeConfig**](SearchAttributeConfigurationAPI.md#PatchSearchAttributeConfig) | **Patch** /accounts/search-attribute-config/{name} | Update search attribute in IdentityNow.



## CreateSearchAttributeConfig

> map[string]interface{} CreateSearchAttributeConfig(ctx).SearchAttributeConfig(searchAttributeConfig).Execute()

Configure/create search attributes in IdentityNow.



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
	searchAttributeConfig := *openapiclient.NewSearchAttributeConfig() // SearchAttributeConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.CreateSearchAttributeConfig(context.Background()).SearchAttributeConfig(searchAttributeConfig).Execute()
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
 **searchAttributeConfig** | [**SearchAttributeConfig**](SearchAttributeConfig.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSearchAttributeConfig

> DeleteSearchAttributeConfig(ctx, name).Execute()

Delete search attribute in IdentityNow.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SearchAttributeConfigurationAPI.DeleteSearchAttributeConfig(context.Background(), name).Execute()
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


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSearchAttributeConfig

> []SearchAttributeConfig GetSearchAttributeConfig(ctx).Execute()

Retrieve attribute list in IdentityNow.



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
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.GetSearchAttributeConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.GetSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchAttributeConfig`: []SearchAttributeConfig
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.GetSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchAttributeConfigRequest struct via the builder pattern


### Return type

[**[]SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleSearchAttributeConfig

> []SearchAttributeConfig GetSingleSearchAttributeConfig(ctx, name).Execute()

Get specific attribute in IdentityNow.



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
	name := "newMailAttribute" // string | Name of the extended search attribute configuration to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleSearchAttributeConfig`: []SearchAttributeConfig
	fmt.Fprintf(os.Stdout, "Response from `SearchAttributeConfigurationAPI.GetSingleSearchAttributeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the extended search attribute configuration to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSearchAttributeConfig

> SearchAttributeConfig PatchSearchAttributeConfig(ctx, name).JsonPatchOperation(jsonPatchOperation).Execute()

Update search attribute in IdentityNow.



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
	name := "promotedMailAttribute" // string | Name of the Search Attribute Configuration to patch.
	jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAttributeConfigurationAPI.PatchSearchAttributeConfig(context.Background(), name).JsonPatchOperation(jsonPatchOperation).Execute()
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
**name** | **string** | Name of the Search Attribute Configuration to patch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSearchAttributeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**SearchAttributeConfig**](SearchAttributeConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

