# \SIMIntegrationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSIMIntegration**](SIMIntegrationsAPI.md#CreateSIMIntegration) | **Post** /sim-integrations | Create new SIM integration
[**DeleteSIMIntegration**](SIMIntegrationsAPI.md#DeleteSIMIntegration) | **Delete** /sim-integrations/{id} | Delete a SIM integration
[**GetSIMIntegration**](SIMIntegrationsAPI.md#GetSIMIntegration) | **Get** /sim-integrations/{id} | Get a SIM integration details.
[**GetSIMIntegrations**](SIMIntegrationsAPI.md#GetSIMIntegrations) | **Get** /sim-integrations | List the existing SIM integrations.
[**PatchBeforeProvisioningRule**](SIMIntegrationsAPI.md#PatchBeforeProvisioningRule) | **Patch** /sim-integrations/{id}/beforeProvisioningRule | Patch a SIM beforeProvisioningRule attribute.
[**PatchSIMAttributes**](SIMIntegrationsAPI.md#PatchSIMAttributes) | **Patch** /sim-integrations/{id} | Patch a SIM attribute.
[**PutSIMIntegration**](SIMIntegrationsAPI.md#PutSIMIntegration) | **Put** /sim-integrations/{id} | Update an existing SIM integration



## CreateSIMIntegration

> ServiceDeskIntegrationDto CreateSIMIntegration(ctx).SimIntegrationDetails(simIntegrationDetails).Execute()

Create new SIM integration



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
    simIntegrationDetails := *openapiclient.NewSimIntegrationDetails("aName") // SimIntegrationDetails | DTO containing the details of the SIM integration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SIMIntegrationsAPI.CreateSIMIntegration(context.Background()).SimIntegrationDetails(simIntegrationDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.CreateSIMIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSIMIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.CreateSIMIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSIMIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simIntegrationDetails** | [**SimIntegrationDetails**](SimIntegrationDetails.md) | DTO containing the details of the SIM integration | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSIMIntegration

> DeleteSIMIntegration(ctx, id).Execute()

Delete a SIM integration



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
    id := "12345" // string | The id of the integration to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SIMIntegrationsAPI.DeleteSIMIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.DeleteSIMIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the integration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSIMIntegrationRequest struct via the builder pattern


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


## GetSIMIntegration

> ServiceDeskIntegrationDto GetSIMIntegration(ctx, id).Execute()

Get a SIM integration details.



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
    id := "12345" // string | The id of the integration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SIMIntegrationsAPI.GetSIMIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.GetSIMIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSIMIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.GetSIMIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSIMIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSIMIntegrations

> ServiceDeskIntegrationDto GetSIMIntegrations(ctx).Execute()

List the existing SIM integrations.



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
    resp, r, err := apiClient.SIMIntegrationsAPI.GetSIMIntegrations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.GetSIMIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSIMIntegrations`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.GetSIMIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSIMIntegrationsRequest struct via the builder pattern


### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchBeforeProvisioningRule

> ServiceDeskIntegrationDto PatchBeforeProvisioningRule(ctx, id).JsonPatch(jsonPatch).Execute()

Patch a SIM beforeProvisioningRule attribute.



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
    id := "12345" // string | SIM integration id
    jsonPatch := *openapiclient.NewJsonPatch() // JsonPatch | The JsonPatch object that describes the changes of SIM beforeProvisioningRule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SIMIntegrationsAPI.PatchBeforeProvisioningRule(context.Background(), id).JsonPatch(jsonPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.PatchBeforeProvisioningRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchBeforeProvisioningRule`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.PatchBeforeProvisioningRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SIM integration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBeforeProvisioningRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatch** | [**JsonPatch**](JsonPatch.md) | The JsonPatch object that describes the changes of SIM beforeProvisioningRule. | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSIMAttributes

> ServiceDeskIntegrationDto PatchSIMAttributes(ctx, id).JsonPatch(jsonPatch).Execute()

Patch a SIM attribute.



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
    id := "12345" // string | SIM integration id
    jsonPatch := *openapiclient.NewJsonPatch() // JsonPatch | The JsonPatch object that describes the changes of SIM

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SIMIntegrationsAPI.PatchSIMAttributes(context.Background(), id).JsonPatch(jsonPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.PatchSIMAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSIMAttributes`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.PatchSIMAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | SIM integration id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSIMAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatch** | [**JsonPatch**](JsonPatch.md) | The JsonPatch object that describes the changes of SIM | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSIMIntegration

> ServiceDeskIntegrationDto PutSIMIntegration(ctx, id).SimIntegrationDetails(simIntegrationDetails).Execute()

Update an existing SIM integration



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
    id := "12345" // string | The id of the integration.
    simIntegrationDetails := *openapiclient.NewSimIntegrationDetails("aName") // SimIntegrationDetails | The full DTO of the integration containing the updated model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SIMIntegrationsAPI.PutSIMIntegration(context.Background(), id).SimIntegrationDetails(simIntegrationDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SIMIntegrationsAPI.PutSIMIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSIMIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `SIMIntegrationsAPI.PutSIMIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSIMIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simIntegrationDetails** | [**SimIntegrationDetails**](SimIntegrationDetails.md) | The full DTO of the integration containing the updated model | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

