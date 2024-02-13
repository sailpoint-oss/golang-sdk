# \ServiceDeskIntegrationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#CreateServiceDeskIntegration) | **Post** /service-desk-integrations | Create new Service Desk integration
[**DeleteServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#DeleteServiceDeskIntegration) | **Delete** /service-desk-integrations/{id} | Delete a Service Desk integration
[**GetServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegration) | **Get** /service-desk-integrations/{id} | Get a Service Desk integration
[**GetServiceDeskIntegrationTemplate**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrationTemplate) | **Get** /service-desk-integrations/templates/{scriptName} | Service Desk integration template by scriptName.
[**GetServiceDeskIntegrationTypes**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrationTypes) | **Get** /service-desk-integrations/types | Service Desk Integration Types List.
[**GetServiceDeskIntegrations**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrations) | **Get** /service-desk-integrations | List existing Service Desk Integrations
[**GetStatusCheckDetails**](ServiceDeskIntegrationAPI.md#GetStatusCheckDetails) | **Get** /service-desk-integrations/status-check-configuration | Get the time check configuration
[**PatchServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#PatchServiceDeskIntegration) | **Patch** /service-desk-integrations/{id} | Service Desk Integration Update PATCH
[**PutServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#PutServiceDeskIntegration) | **Put** /service-desk-integrations/{id} | Update a Service Desk integration
[**UpdateStatusCheckDetails**](ServiceDeskIntegrationAPI.md#UpdateStatusCheckDetails) | **Put** /service-desk-integrations/status-check-configuration | Update the time check configuration



## Create new Service Desk integration

> ServiceDeskIntegrationDto CreateServiceDeskIntegration(ctx).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()

Create a new Service Desk Integrations.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    serviceDeskIntegrationDto := *sailpoint.NewServiceDeskIntegrationDto("Service Desk Integration Name", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of a new integration to create

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.CreateServiceDeskIntegration(context.Background()).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.CreateServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.CreateServiceDeskIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceDeskIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceDeskIntegrationDto** | [**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md) | The specifics of a new integration to create | 

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


## Delete a Service Desk integration

> DeleteServiceDeskIntegration(ctx, id).Execute()

Delete an existing Service Desk integration by ID.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "anId" // string | ID of Service Desk integration to delete

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.ServiceDeskIntegrationAPI.DeleteServiceDeskIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.DeleteServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of Service Desk integration to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceDeskIntegrationRequest struct via the builder pattern


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


## Get a Service Desk integration

> ServiceDeskIntegrationDto GetServiceDeskIntegration(ctx, id).Execute()

Get an existing Service Desk integration by ID.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "anId" // string | ID of the Service Desk integration to get

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.GetServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.GetServiceDeskIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Service Desk integration to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationRequest struct via the builder pattern


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


## Service Desk integration template by scriptName.

> ServiceDeskIntegrationTemplateDto GetServiceDeskIntegrationTemplate(ctx, scriptName).Execute()

This API endpoint returns an existing Service Desk integration template by scriptName.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    scriptName := "aScriptName" // string | The scriptName value of the Service Desk integration template to get

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTemplate(context.Background(), scriptName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrationTemplate`: ServiceDeskIntegrationTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scriptName** | **string** | The scriptName value of the Service Desk integration template to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceDeskIntegrationTemplateDto**](ServiceDeskIntegrationTemplateDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Service Desk Integration Types List.

> []ServiceDeskIntegrationTemplateType GetServiceDeskIntegrationTypes(ctx).Execute()

This API endpoint returns the current list of supported Service Desk integration types.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrationTypes`: []ServiceDeskIntegrationTemplateType
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationTypesRequest struct via the builder pattern


### Return type

[**[]ServiceDeskIntegrationTemplateType**](ServiceDeskIntegrationTemplateType.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List existing Service Desk Integrations

> []ServiceDeskIntegrationDto GetServiceDeskIntegrations(ctx).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()

Get a list of ServiceDeskIntegrationDto for existing Service Desk Integrations.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)
    filters := "name eq "John Doe"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq*  **type**: *eq, in*  **cluster**: *eq, in* (optional)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrations(context.Background()).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.GetServiceDeskIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrations`: []ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.GetServiceDeskIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq*  **type**: *eq, in*  **cluster**: *eq, in* | 
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get the time check configuration

> QueuedCheckConfigDetails GetStatusCheckDetails(ctx).Execute()

Get the time check configuration of queued SDIM tickets.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.GetStatusCheckDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.GetStatusCheckDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatusCheckDetails`: QueuedCheckConfigDetails
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.GetStatusCheckDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCheckDetailsRequest struct via the builder pattern


### Return type

[**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Service Desk Integration Update PATCH

> ServiceDeskIntegrationDto PatchServiceDeskIntegration(ctx, id).PatchServiceDeskIntegrationRequest(patchServiceDeskIntegrationRequest).Execute()

Update an existing ServiceDeskIntegration by ID with a PATCH request.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "anId" // string | ID of the Service Desk integration to update
    patchServiceDeskIntegrationRequest := *sailpoint.NewPatchServiceDeskIntegrationRequest() // PatchServiceDeskIntegrationRequest | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Only `replace` operations are accepted by this endpoint.  A 403 Forbidden Error indicates that you attempted to PATCH a operation that is not allowed. 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.PatchServiceDeskIntegration(context.Background(), id).PatchServiceDeskIntegrationRequest(patchServiceDeskIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.PatchServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.PatchServiceDeskIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Service Desk integration to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchServiceDeskIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchServiceDeskIntegrationRequest** | [**PatchServiceDeskIntegrationRequest**](PatchServiceDeskIntegrationRequest.md) | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Only &#x60;replace&#x60; operations are accepted by this endpoint.  A 403 Forbidden Error indicates that you attempted to PATCH a operation that is not allowed.  | 

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


## Update a Service Desk integration

> ServiceDeskIntegrationDto PutServiceDeskIntegration(ctx, id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()

Update an existing Service Desk integration by ID with updated value in JSON form as the request body.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "anId" // string | ID of the Service Desk integration to update
    serviceDeskIntegrationDto := *sailpoint.NewServiceDeskIntegrationDto("Service Desk Integration Name", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of the integration to update

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.PutServiceDeskIntegration(context.Background(), id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.PutServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.PutServiceDeskIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Service Desk integration to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutServiceDeskIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceDeskIntegrationDto** | [**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md) | The specifics of the integration to update | 

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


## Update the time check configuration

> QueuedCheckConfigDetails UpdateStatusCheckDetails(ctx).QueuedCheckConfigDetails(queuedCheckConfigDetails).Execute()

Update the time check configuration of queued SDIM tickets.  A token with Org Admin or Service Desk Admin authority is required to access this endpoint.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    queuedCheckConfigDetails := *sailpoint.NewQueuedCheckConfigDetails("30", "2") // QueuedCheckConfigDetails | the modified time check configuration

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.ServiceDeskIntegrationAPI.UpdateStatusCheckDetails(context.Background()).QueuedCheckConfigDetails(queuedCheckConfigDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationAPI.UpdateStatusCheckDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStatusCheckDetails`: QueuedCheckConfigDetails
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationAPI.UpdateStatusCheckDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusCheckDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queuedCheckConfigDetails** | [**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md) | the modified time check configuration | 

### Return type

[**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

