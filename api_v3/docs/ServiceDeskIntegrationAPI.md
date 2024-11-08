# \ServiceDeskIntegrationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#CreateServiceDeskIntegration) | **Post** /service-desk-integrations | Create new Service Desk integration
[**DeleteServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#DeleteServiceDeskIntegration) | **Delete** /service-desk-integrations/{id} | Delete a Service Desk integration
[**GetServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegration) | **Get** /service-desk-integrations/{id} | Get a Service Desk integration
[**GetServiceDeskIntegrationTemplate**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrationTemplate) | **Get** /service-desk-integrations/templates/{scriptName} | Service Desk integration template by scriptName
[**GetServiceDeskIntegrationTypes**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrationTypes) | **Get** /service-desk-integrations/types | List Service Desk integration types
[**GetServiceDeskIntegrations**](ServiceDeskIntegrationAPI.md#GetServiceDeskIntegrations) | **Get** /service-desk-integrations | List existing Service Desk integrations
[**GetStatusCheckDetails**](ServiceDeskIntegrationAPI.md#GetStatusCheckDetails) | **Get** /service-desk-integrations/status-check-configuration | Get the time check configuration
[**PatchServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#PatchServiceDeskIntegration) | **Patch** /service-desk-integrations/{id} | Patch a Service Desk Integration
[**PutServiceDeskIntegration**](ServiceDeskIntegrationAPI.md#PutServiceDeskIntegration) | **Put** /service-desk-integrations/{id} | Update a Service Desk integration
[**UpdateStatusCheckDetails**](ServiceDeskIntegrationAPI.md#UpdateStatusCheckDetails) | **Put** /service-desk-integrations/status-check-configuration | Update the time check configuration



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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	serviceDeskIntegrationDto := *openapiclient.NewServiceDeskIntegrationDto("Service Desk Integration Name", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of a new integration to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.CreateServiceDeskIntegration(context.Background()).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	id := "anId" // string | ID of Service Desk integration to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ServiceDeskIntegrationAPI.DeleteServiceDeskIntegration(context.Background(), id).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	id := "anId" // string | ID of the Service Desk integration to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.GetServiceDeskIntegration(context.Background(), id).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrationTemplate

> ServiceDeskIntegrationTemplateDto GetServiceDeskIntegrationTemplate(ctx, scriptName).Execute()

Service Desk integration template by scriptName



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
	scriptName := "aScriptName" // string | The scriptName value of the Service Desk integration template to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTemplate(context.Background(), scriptName).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrationTypes

> []ServiceDeskIntegrationTemplateType GetServiceDeskIntegrationTypes(ctx).Execute()

List Service Desk integration types



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
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTypes(context.Background()).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDeskIntegrations

> []ServiceDeskIntegrationDto GetServiceDeskIntegrations(ctx).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()

List existing Service Desk integrations



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
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)
	filters := "name eq "John Doe"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq*  **type**: *eq, in*  **cluster**: *eq, in* (optional)
	count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.GetServiceDeskIntegrations(context.Background()).Offset(offset).Limit(limit).Sorters(sorters).Filters(filters).Count(count).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.GetStatusCheckDetails(context.Background()).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchServiceDeskIntegration

> ServiceDeskIntegrationDto PatchServiceDeskIntegration(ctx, id).PatchServiceDeskIntegrationRequest(patchServiceDeskIntegrationRequest).Execute()

Patch a Service Desk Integration



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
	id := "anId" // string | ID of the Service Desk integration to update
	patchServiceDeskIntegrationRequest := *openapiclient.NewPatchServiceDeskIntegrationRequest() // PatchServiceDeskIntegrationRequest | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Only `replace` operations are accepted by this endpoint.  A 403 Forbidden Error indicates that a PATCH operation was attempted that is not allowed. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.PatchServiceDeskIntegration(context.Background(), id).PatchServiceDeskIntegrationRequest(patchServiceDeskIntegrationRequest).Execute()
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

 **patchServiceDeskIntegrationRequest** | [**PatchServiceDeskIntegrationRequest**](PatchServiceDeskIntegrationRequest.md) | A list of SDIM update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  Only &#x60;replace&#x60; operations are accepted by this endpoint.  A 403 Forbidden Error indicates that a PATCH operation was attempted that is not allowed.  | 

### Return type

[**ServiceDeskIntegrationDto**](ServiceDeskIntegrationDto.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServiceDeskIntegration

> ServiceDeskIntegrationDto PutServiceDeskIntegration(ctx, id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()

Update a Service Desk integration



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
	id := "anId" // string | ID of the Service Desk integration to update
	serviceDeskIntegrationDto := *openapiclient.NewServiceDeskIntegrationDto("Service Desk Integration Name", "A very nice Service Desk integration", "ServiceNowSDIM", map[string]interface{}{"key": interface{}(123)}) // ServiceDeskIntegrationDto | The specifics of the integration to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.PutServiceDeskIntegration(context.Background(), id).ServiceDeskIntegrationDto(serviceDeskIntegrationDto).Execute()
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
	queuedCheckConfigDetails := *openapiclient.NewQueuedCheckConfigDetails("30", "2") // QueuedCheckConfigDetails | The modified time check configuration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceDeskIntegrationAPI.UpdateStatusCheckDetails(context.Background()).QueuedCheckConfigDetails(queuedCheckConfigDetails).Execute()
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
 **queuedCheckConfigDetails** | [**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md) | The modified time check configuration | 

### Return type

[**QueuedCheckConfigDetails**](QueuedCheckConfigDetails.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

