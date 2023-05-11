# \ServiceDeskIntegrationApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDeskIntegration**](ServiceDeskIntegrationApi.md#CreateServiceDeskIntegration) | **Post** /service-desk-integrations | Create new Service Desk integration
[**DeleteServiceDeskIntegration**](ServiceDeskIntegrationApi.md#DeleteServiceDeskIntegration) | **Delete** /service-desk-integrations/{id} | Delete a Service Desk integration
[**GetServiceDeskIntegration**](ServiceDeskIntegrationApi.md#GetServiceDeskIntegration) | **Get** /service-desk-integrations/{id} | Get a Service Desk integration
[**GetServiceDeskIntegrationList**](ServiceDeskIntegrationApi.md#GetServiceDeskIntegrationList) | **Get** /service-desk-integrations | List existing Service Desk Integrations
[**GetServiceDeskIntegrationTemplate**](ServiceDeskIntegrationApi.md#GetServiceDeskIntegrationTemplate) | **Get** /service-desk-integrations/templates/{scriptName} | Service Desk integration template by scriptName.
[**GetServiceDeskIntegrationTypes**](ServiceDeskIntegrationApi.md#GetServiceDeskIntegrationTypes) | **Get** /service-desk-integrations/types | Service Desk Integration Types List.
[**GetStatusCheckDetails**](ServiceDeskIntegrationApi.md#GetStatusCheckDetails) | **Get** /service-desk-integrations/status-check-configuration | Get the time check configuration
[**PatchServiceDeskIntegration**](ServiceDeskIntegrationApi.md#PatchServiceDeskIntegration) | **Patch** /service-desk-integrations/{id} | Service Desk Integration Update PATCH
[**UpdateServiceDeskIntegration**](ServiceDeskIntegrationApi.md#UpdateServiceDeskIntegration) | **Put** /service-desk-integrations/{id} | Update a Service Desk integration
[**UpdateStatusCheckDetails**](ServiceDeskIntegrationApi.md#UpdateStatusCheckDetails) | **Put** /service-desk-integrations/status-check-configuration | Update the time check configuration



## CreateServiceDeskIntegration

> ServiceDeskIntegrationDto CreateServiceDeskIntegration(ctx).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()

Create new Service Desk integration



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
    serviceDeskIntegrationDto := *openapiclient.NewServiceDeskIntegrationDto("aName", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of a new integration to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.CreateServiceDeskIntegration(context.Background()).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.CreateServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.CreateServiceDeskIntegration`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceDeskIntegration

> DeleteServiceDeskIntegration(ctx, id).Execute()

Delete a Service Desk integration



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
    id := "anId" // string | ID of Service Desk integration to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.DeleteServiceDeskIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.DeleteServiceDeskIntegration``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegration

> ServiceDeskIntegrationDto GetServiceDeskIntegration(ctx, id).Execute()

Get a Service Desk integration



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
    id := "anId" // string | ID of the Service Desk integration to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.GetServiceDeskIntegration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.GetServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.GetServiceDeskIntegration`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrationList

> []ServiceDeskIntegrationDto GetServiceDeskIntegrationList(ctx).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()

List existing Service Desk Integrations



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)
    filters := "id eq 2c91808b6ef1d43e016efba0ce470904" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq*  **type**: *eq, in*  **cluster**: *eq, in* (optional)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.GetServiceDeskIntegrationList(context.Background()).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.GetServiceDeskIntegrationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrationList`: []ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.GetServiceDeskIntegrationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationListRequest struct via the builder pattern


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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrationTemplate

> ServiceDeskIntegrationTemplateDto GetServiceDeskIntegrationTemplate(ctx, scriptName).Execute()

Service Desk integration template by scriptName.



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
    scriptName := "aScriptName" // string | The scriptName value of the Service Desk integration template to get

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.GetServiceDeskIntegrationTemplate(context.Background(), scriptName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.GetServiceDeskIntegrationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrationTemplate`: ServiceDeskIntegrationTemplateDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.GetServiceDeskIntegrationTemplate`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrationTypes

> []ServiceDeskIntegrationTemplateType GetServiceDeskIntegrationTypes(ctx).Execute()

Service Desk Integration Types List.



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
    resp, r, err := apiClient.ServiceDeskIntegrationApi.GetServiceDeskIntegrationTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.GetServiceDeskIntegrationTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDeskIntegrationTypes`: []ServiceDeskIntegrationTemplateType
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.GetServiceDeskIntegrationTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDeskIntegrationTypesRequest struct via the builder pattern


### Return type

[**[]ServiceDeskIntegrationTemplateType**](ServiceDeskIntegrationTemplateType.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusCheckDetails

> QueuedCheckConfigDetails GetStatusCheckDetails(ctx).Execute()

Get the time check configuration



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
    resp, r, err := apiClient.ServiceDeskIntegrationApi.GetStatusCheckDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.GetStatusCheckDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatusCheckDetails`: QueuedCheckConfigDetails
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.GetStatusCheckDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCheckDetailsRequest struct via the builder pattern


### Return type

[**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceDeskIntegration

> ServiceDeskIntegrationDto PatchServiceDeskIntegration(ctx, id).JsonPatch(jsonPatch).Execute()

Service Desk Integration Update PATCH



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
    id := "anId" // string | ID of the Service Desk integration to update
    jsonPatch := *openapiclient.NewJsonPatch() // JsonPatch | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  PATCH can only be applied to the following fields:   *   \"beforeProvisioningRule\"  A 403 Forbidden Error indicates that you attempted to PATCH a field that is not allowed. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.PatchServiceDeskIntegration(context.Background(), id).JsonPatch(jsonPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.PatchServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.PatchServiceDeskIntegration`: %v\n", resp)
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

 **jsonPatch** | [**JsonPatch**](JsonPatch.md) | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  PATCH can only be applied to the following fields:   *   \&quot;beforeProvisioningRule\&quot;  A 403 Forbidden Error indicates that you attempted to PATCH a field that is not allowed.  | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceDeskIntegration

> ServiceDeskIntegrationDto UpdateServiceDeskIntegration(ctx, id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()

Update a Service Desk integration



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
    id := "anId" // string | ID of the Service Desk integration to update
    serviceDeskIntegrationDto := *openapiclient.NewServiceDeskIntegrationDto("aName", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of the integration to update

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.UpdateServiceDeskIntegration(context.Background(), id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.UpdateServiceDeskIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceDeskIntegration`: ServiceDeskIntegrationDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.UpdateServiceDeskIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the Service Desk integration to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceDeskIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceDeskIntegrationDto** | [**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md) | The specifics of the integration to update | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatusCheckDetails

> QueuedCheckConfigDetails UpdateStatusCheckDetails(ctx).QueuedCheckConfigDetails(queuedCheckConfigDetails).Execute()

Update the time check configuration



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
    queuedCheckConfigDetails := *openapiclient.NewQueuedCheckConfigDetails("30", "2") // QueuedCheckConfigDetails | the modified time check configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDeskIntegrationApi.UpdateStatusCheckDetails(context.Background()).QueuedCheckConfigDetails(queuedCheckConfigDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDeskIntegrationApi.UpdateStatusCheckDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStatusCheckDetails`: QueuedCheckConfigDetails
    fmt.Fprintf(os.Stdout, "Response from `ServiceDeskIntegrationApi.UpdateStatusCheckDetails`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

