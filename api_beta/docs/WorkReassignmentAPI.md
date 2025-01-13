# \WorkReassignmentAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReassignmentConfiguration**](WorkReassignmentAPI.md#CreateReassignmentConfiguration) | **Post** /reassignment-configurations | Create a Reassignment Configuration
[**DeleteReassignmentConfiguration**](WorkReassignmentAPI.md#DeleteReassignmentConfiguration) | **Delete** /reassignment-configurations/{identityId} | Delete Reassignment Configuration
[**GetEvaluateReassignmentConfiguration**](WorkReassignmentAPI.md#GetEvaluateReassignmentConfiguration) | **Get** /reassignment-configurations/{identityId}/evaluate/{configType} | Evaluate Reassignment Configuration
[**GetReassignmentConfigTypes**](WorkReassignmentAPI.md#GetReassignmentConfigTypes) | **Get** /reassignment-configurations/types | List Reassignment Config Types
[**GetReassignmentConfiguration**](WorkReassignmentAPI.md#GetReassignmentConfiguration) | **Get** /reassignment-configurations/{identityId} | Get Reassignment Configuration
[**GetTenantConfigConfiguration**](WorkReassignmentAPI.md#GetTenantConfigConfiguration) | **Get** /reassignment-configurations/tenant-config | Get Tenant-wide Reassignment Configuration settings
[**ListReassignmentConfigurations**](WorkReassignmentAPI.md#ListReassignmentConfigurations) | **Get** /reassignment-configurations | List Reassignment Configurations
[**PutReassignmentConfig**](WorkReassignmentAPI.md#PutReassignmentConfig) | **Put** /reassignment-configurations/{identityId} | Update Reassignment Configuration
[**PutTenantConfiguration**](WorkReassignmentAPI.md#PutTenantConfiguration) | **Put** /reassignment-configurations/tenant-config | Update Tenant-wide Reassignment Configuration settings



## CreateReassignmentConfiguration

> ConfigurationItemResponse CreateReassignmentConfiguration(ctx).ConfigurationItemRequest(configurationItemRequest).Execute()

Create a Reassignment Configuration



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
	configurationItemRequest := *openapiclient.NewConfigurationItemRequest() // ConfigurationItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkReassignmentAPI.CreateReassignmentConfiguration(context.Background()).ConfigurationItemRequest(configurationItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.CreateReassignmentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReassignmentConfiguration`: ConfigurationItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.CreateReassignmentConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReassignmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationItemRequest** | [**ConfigurationItemRequest**](ConfigurationItemRequest.md) |  | 

### Return type

[**ConfigurationItemResponse**](ConfigurationItemResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReassignmentConfiguration

> DeleteReassignmentConfiguration(ctx, identityId).Execute()

Delete Reassignment Configuration



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
	identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkReassignmentAPI.DeleteReassignmentConfiguration(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.DeleteReassignmentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | unique identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReassignmentConfigurationRequest struct via the builder pattern


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


## GetEvaluateReassignmentConfiguration

> []EvaluateResponse GetEvaluateReassignmentConfiguration(ctx, identityId, configType).ExclusionFilters(exclusionFilters).Execute()

Evaluate Reassignment Configuration



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
	identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id
	configType := openapiclient.ConfigTypeEnum("ACCESS_REQUESTS") // ConfigTypeEnum | Reassignment work type
	exclusionFilters := []string{"Inner_example"} // []string | Exclusion filters that disable parts of the reassignment evaluation. Possible values are listed below: - `SELF_REVIEW_DELEGATION`: This will exclude delegations of self-review reassignments (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkReassignmentAPI.GetEvaluateReassignmentConfiguration(context.Background(), identityId, configType).ExclusionFilters(exclusionFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.GetEvaluateReassignmentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvaluateReassignmentConfiguration`: []EvaluateResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.GetEvaluateReassignmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | unique identity id | 
**configType** | [**ConfigTypeEnum**](.md) | Reassignment work type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluateReassignmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionFilters** | **[]string** | Exclusion filters that disable parts of the reassignment evaluation. Possible values are listed below: - &#x60;SELF_REVIEW_DELEGATION&#x60;: This will exclude delegations of self-review reassignments | 

### Return type

[**[]EvaluateResponse**](EvaluateResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReassignmentConfigTypes

> []ConfigType GetReassignmentConfigTypes(ctx).Execute()

List Reassignment Config Types



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
	resp, r, err := apiClient.WorkReassignmentAPI.GetReassignmentConfigTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.GetReassignmentConfigTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReassignmentConfigTypes`: []ConfigType
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.GetReassignmentConfigTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReassignmentConfigTypesRequest struct via the builder pattern


### Return type

[**[]ConfigType**](ConfigType.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReassignmentConfiguration

> ConfigurationResponse GetReassignmentConfiguration(ctx, identityId).Execute()

Get Reassignment Configuration



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
	identityId := "2c91808781a71ddb0181b9090b5c504f" // string | unique identity id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkReassignmentAPI.GetReassignmentConfiguration(context.Background(), identityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.GetReassignmentConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReassignmentConfiguration`: ConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.GetReassignmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | unique identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReassignmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantConfigConfiguration

> TenantConfigurationResponse GetTenantConfigConfiguration(ctx).Execute()

Get Tenant-wide Reassignment Configuration settings



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
	resp, r, err := apiClient.WorkReassignmentAPI.GetTenantConfigConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.GetTenantConfigConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantConfigConfiguration`: TenantConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.GetTenantConfigConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantConfigConfigurationRequest struct via the builder pattern


### Return type

[**TenantConfigurationResponse**](TenantConfigurationResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReassignmentConfigurations

> []ConfigurationResponse ListReassignmentConfigurations(ctx).Execute()

List Reassignment Configurations



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
	resp, r, err := apiClient.WorkReassignmentAPI.ListReassignmentConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.ListReassignmentConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReassignmentConfigurations`: []ConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.ListReassignmentConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListReassignmentConfigurationsRequest struct via the builder pattern


### Return type

[**[]ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReassignmentConfig

> ConfigurationItemResponse PutReassignmentConfig(ctx, identityId).ConfigurationItemRequest(configurationItemRequest).Execute()

Update Reassignment Configuration



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
	identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id
	configurationItemRequest := *openapiclient.NewConfigurationItemRequest() // ConfigurationItemRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkReassignmentAPI.PutReassignmentConfig(context.Background(), identityId).ConfigurationItemRequest(configurationItemRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.PutReassignmentConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutReassignmentConfig`: ConfigurationItemResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.PutReassignmentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | unique identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutReassignmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationItemRequest** | [**ConfigurationItemRequest**](ConfigurationItemRequest.md) |  | 

### Return type

[**ConfigurationItemResponse**](ConfigurationItemResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTenantConfiguration

> TenantConfigurationResponse PutTenantConfiguration(ctx).TenantConfigurationRequest(tenantConfigurationRequest).Execute()

Update Tenant-wide Reassignment Configuration settings



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
	tenantConfigurationRequest := *openapiclient.NewTenantConfigurationRequest() // TenantConfigurationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkReassignmentAPI.PutTenantConfiguration(context.Background()).TenantConfigurationRequest(tenantConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentAPI.PutTenantConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTenantConfiguration`: TenantConfigurationResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentAPI.PutTenantConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutTenantConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantConfigurationRequest** | [**TenantConfigurationRequest**](TenantConfigurationRequest.md) |  | 

### Return type

[**TenantConfigurationResponse**](TenantConfigurationResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

