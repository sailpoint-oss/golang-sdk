# \WorkReassignmentApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReassignmentConfiguration**](WorkReassignmentApi.md#CreateReassignmentConfiguration) | **Post** /reassignment-configurations | Create a Reassignment Configuration
[**DeleteReassignmentConfiguration**](WorkReassignmentApi.md#DeleteReassignmentConfiguration) | **Delete** /reassignment-configurations/{identityId} | Delete Reassignment Configuration
[**GetEvaluateReassignmentConfiguration**](WorkReassignmentApi.md#GetEvaluateReassignmentConfiguration) | **Get** /reassignment-configurations/{identityId}/evaluate/{configType} | Evaluate Reassignment Configuration
[**GetReassignmentConfigTypes**](WorkReassignmentApi.md#GetReassignmentConfigTypes) | **Get** /reassignment-configurations/types | List Reassignment Config Types
[**GetReassignmentConfiguration**](WorkReassignmentApi.md#GetReassignmentConfiguration) | **Get** /reassignment-configurations/{identityId} | Get Reassignment Configuration
[**GetTenantConfigConfiguration**](WorkReassignmentApi.md#GetTenantConfigConfiguration) | **Get** /reassignment-configurations/tenant-config | Get Tenant-wide Reassignment Configuration settings
[**ListReassignmentConfigurations**](WorkReassignmentApi.md#ListReassignmentConfigurations) | **Get** /reassignment-configurations | List Reassignment Configurations
[**PutReassignmentConfig**](WorkReassignmentApi.md#PutReassignmentConfig) | **Put** /reassignment-configurations/{identityId} | Update Reassignment Configuration
[**PutTenantConfiguration**](WorkReassignmentApi.md#PutTenantConfiguration) | **Put** /reassignment-configurations/tenant-config | Update Tenant-wide Reassignment Configuration settings



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
    openapiclient "./openapi"
)

func main() {
    configurationItemRequest := *openapiclient.NewConfigurationItemRequest() // ConfigurationItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.CreateReassignmentConfiguration(context.Background()).ConfigurationItemRequest(configurationItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.CreateReassignmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReassignmentConfiguration`: ConfigurationItemResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.CreateReassignmentConfiguration`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.DeleteReassignmentConfiguration(context.Background(), identityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.DeleteReassignmentConfiguration``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id
    configType := openapiclient.WorkTypeEnum("accessRequests") // WorkTypeEnum | Reassignment work type
    exclusionFilters := []string{"Inner_example"} // []string | Exclusion filters that disable parts of the reassignment evaluation. Possible values are listed below: - `SELF_REVIEW_DELEGATION`: This will exclude delegations of self-review reassignments (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.GetEvaluateReassignmentConfiguration(context.Background(), identityId, configType).ExclusionFilters(exclusionFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.GetEvaluateReassignmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvaluateReassignmentConfiguration`: []EvaluateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.GetEvaluateReassignmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | unique identity id | 
**configType** | [**WorkTypeEnum**](.md) | Reassignment work type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEvaluateReassignmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exclusionFilters** | **[]string** | Exclusion filters that disable parts of the reassignment evaluation. Possible values are listed below: - &#x60;SELF_REVIEW_DELEGATION&#x60;: This will exclude delegations of self-review reassignments | 

### Return type

[**[]EvaluateResponse**](EvaluateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.GetReassignmentConfigTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.GetReassignmentConfigTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReassignmentConfigTypes`: []ConfigType
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.GetReassignmentConfigTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetReassignmentConfigTypesRequest struct via the builder pattern


### Return type

[**[]ConfigType**](ConfigType.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    identityId := "2c91808781a71ddb0181b9090b5c504f" // string | unique identity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.GetReassignmentConfiguration(context.Background(), identityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.GetReassignmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReassignmentConfiguration`: ConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.GetReassignmentConfiguration`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.GetTenantConfigConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.GetTenantConfigConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantConfigConfiguration`: TenantConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.GetTenantConfigConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantConfigConfigurationRequest struct via the builder pattern


### Return type

[**TenantConfigurationResponse**](TenantConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.ListReassignmentConfigurations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.ListReassignmentConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReassignmentConfigurations`: []ConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.ListReassignmentConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListReassignmentConfigurationsRequest struct via the builder pattern


### Return type

[**[]ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    identityId := "2c91808781a71ddb0181b9090b5c504e" // string | unique identity id
    configurationItemRequest := *openapiclient.NewConfigurationItemRequest() // ConfigurationItemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.PutReassignmentConfig(context.Background(), identityId).ConfigurationItemRequest(configurationItemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.PutReassignmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutReassignmentConfig`: ConfigurationItemResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.PutReassignmentConfig`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    tenantConfigurationRequest := *openapiclient.NewTenantConfigurationRequest() // TenantConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkReassignmentApi.PutTenantConfiguration(context.Background()).TenantConfigurationRequest(tenantConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkReassignmentApi.PutTenantConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTenantConfiguration`: TenantConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkReassignmentApi.PutTenantConfiguration`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

