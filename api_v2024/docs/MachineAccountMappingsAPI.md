# \MachineAccountMappingsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMachineAccountMappings**](MachineAccountMappingsAPI.md#CreateMachineAccountMappings) | **Post** /sources/{sourceId}/machine-account-mappings | Create Machine Account Mappings
[**DeleteMachineAccountMappings**](MachineAccountMappingsAPI.md#DeleteMachineAccountMappings) | **Delete** /sources/{sourceId}/machine-account-mappings | Delete Source&#39;s Machine Account Mappings
[**ListMachineAccountMappings**](MachineAccountMappingsAPI.md#ListMachineAccountMappings) | **Get** /sources/{sourceId}/machine-account-mappings | Machine Account Mapping for Source
[**SetMachineAccountMappings**](MachineAccountMappingsAPI.md#SetMachineAccountMappings) | **Put** /sources/{sourceId}/machine-mappings | Update Source&#39;s Machine Account Mappings



## CreateMachineAccountMappings

> []AttributeMappings CreateMachineAccountMappings(ctx, id).AttributeMappings(attributeMappings).Execute()

Create Machine Account Mappings



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID.
	attributeMappings := *openapiclient.NewAttributeMappings() // AttributeMappings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAccountMappingsAPI.CreateMachineAccountMappings(context.Background(), id).AttributeMappings(attributeMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountMappingsAPI.CreateMachineAccountMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMachineAccountMappings`: []AttributeMappings
	fmt.Fprintf(os.Stdout, "Response from `MachineAccountMappingsAPI.CreateMachineAccountMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMachineAccountMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeMappings** | [**AttributeMappings**](AttributeMappings.md) |  | 

### Return type

[**[]AttributeMappings**](AttributeMappings.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMachineAccountMappings

> DeleteMachineAccountMappings(ctx, id).Execute()

Delete Source's Machine Account Mappings



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | source ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachineAccountMappingsAPI.DeleteMachineAccountMappings(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountMappingsAPI.DeleteMachineAccountMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineAccountMappingsRequest struct via the builder pattern


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


## ListMachineAccountMappings

> []AttributeMappings ListMachineAccountMappings(ctx, id).Limit(limit).Offset(offset).Execute()

Machine Account Mapping for Source



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID
	limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
	offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAccountMappingsAPI.ListMachineAccountMappings(context.Background(), id).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountMappingsAPI.ListMachineAccountMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMachineAccountMappings`: []AttributeMappings
	fmt.Fprintf(os.Stdout, "Response from `MachineAccountMappingsAPI.ListMachineAccountMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMachineAccountMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]AttributeMappings**](AttributeMappings.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMachineAccountMappings

> []AttributeMappings SetMachineAccountMappings(ctx, id).AttributeMappings(attributeMappings).Execute()

Update Source's Machine Account Mappings



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID.
	attributeMappings := *openapiclient.NewAttributeMappings() // AttributeMappings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineAccountMappingsAPI.SetMachineAccountMappings(context.Background(), id).AttributeMappings(attributeMappings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountMappingsAPI.SetMachineAccountMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMachineAccountMappings`: []AttributeMappings
	fmt.Fprintf(os.Stdout, "Response from `MachineAccountMappingsAPI.SetMachineAccountMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMachineAccountMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeMappings** | [**AttributeMappings**](AttributeMappings.md) |  | 

### Return type

[**[]AttributeMappings**](AttributeMappings.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

