# \ConnectorCustomizersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorCustomizer**](ConnectorCustomizersAPI.md#CreateConnectorCustomizer) | **Post** /connector-customizers | Create Connector Customizer
[**CreateConnectorCustomizerVersion**](ConnectorCustomizersAPI.md#CreateConnectorCustomizerVersion) | **Post** /connector-customizers/{id}/versions | Creates a connector customizer version
[**DeleteConnectorCustomizer**](ConnectorCustomizersAPI.md#DeleteConnectorCustomizer) | **Delete** /connector-customizers/{id} | Delete Connector Customizer
[**GetConnectorCustomizer**](ConnectorCustomizersAPI.md#GetConnectorCustomizer) | **Get** /connector-customizers/{id} | Get connector customizer
[**ListConnectorCustomizers**](ConnectorCustomizersAPI.md#ListConnectorCustomizers) | **Get** /connector-customizers | List All Connector Customizers
[**PutConnectorCustomizer**](ConnectorCustomizersAPI.md#PutConnectorCustomizer) | **Put** /connector-customizers/{id} | Update Connector Customizer



## CreateConnectorCustomizer

> ConnectorCustomizerCreateResponse CreateConnectorCustomizer(ctx).ConnectorCustomizerCreateRequest(connectorCustomizerCreateRequest).Execute()

Create Connector Customizer



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
	connectorCustomizerCreateRequest := *openapiclient.NewConnectorCustomizerCreateRequest() // ConnectorCustomizerCreateRequest | Connector customizer to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorCustomizersAPI.CreateConnectorCustomizer(context.Background()).ConnectorCustomizerCreateRequest(connectorCustomizerCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.CreateConnectorCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorCustomizer`: ConnectorCustomizerCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorCustomizersAPI.CreateConnectorCustomizer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorCustomizerCreateRequest** | [**ConnectorCustomizerCreateRequest**](ConnectorCustomizerCreateRequest.md) | Connector customizer to create. | 

### Return type

[**ConnectorCustomizerCreateResponse**](ConnectorCustomizerCreateResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectorCustomizerVersion

> ConnectorCustomizerVersionCreateResponse CreateConnectorCustomizerVersion(ctx, id).Execute()

Creates a connector customizer version



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
	id := "b07dc46a-1498-4de8-bfbb-259a68e70c8a" // string | The id of the connector customizer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorCustomizersAPI.CreateConnectorCustomizerVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.CreateConnectorCustomizerVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorCustomizerVersion`: ConnectorCustomizerVersionCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorCustomizersAPI.CreateConnectorCustomizerVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the connector customizer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorCustomizerVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorCustomizerVersionCreateResponse**](ConnectorCustomizerVersionCreateResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorCustomizer

> DeleteConnectorCustomizer(ctx, id).Execute()

Delete Connector Customizer



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
	id := "b07dc46a-1498-4de8-bfbb-259a68e70c8a" // string | ID of the connector customizer to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorCustomizersAPI.DeleteConnectorCustomizer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.DeleteConnectorCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector customizer to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorCustomizerRequest struct via the builder pattern


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


## GetConnectorCustomizer

> ConnectorCustomizersResponse GetConnectorCustomizer(ctx, id).Execute()

Get connector customizer



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
	id := "b07dc46a-1498-4de8-bfbb-259a68e70c8a" // string | ID of the connector customizer to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorCustomizersAPI.GetConnectorCustomizer(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.GetConnectorCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorCustomizer`: ConnectorCustomizersResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorCustomizersAPI.GetConnectorCustomizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector customizer to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConnectorCustomizersResponse**](ConnectorCustomizersResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorCustomizers

> []ConnectorCustomizersResponse ListConnectorCustomizers(ctx).Offset(offset).Limit(limit).Execute()

List All Connector Customizers



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorCustomizersAPI.ListConnectorCustomizers(context.Background()).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.ListConnectorCustomizers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConnectorCustomizers`: []ConnectorCustomizersResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorCustomizersAPI.ListConnectorCustomizers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorCustomizersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]

### Return type

[**[]ConnectorCustomizersResponse**](ConnectorCustomizersResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConnectorCustomizer

> ConnectorCustomizerUpdateResponse PutConnectorCustomizer(ctx, id).ConnectorCustomizerUpdateRequest(connectorCustomizerUpdateRequest).Execute()

Update Connector Customizer



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
	id := "b07dc46a-1498-4de8-bfbb-259a68e70c8a" // string | ID of the connector customizer to update.
	connectorCustomizerUpdateRequest := *openapiclient.NewConnectorCustomizerUpdateRequest() // ConnectorCustomizerUpdateRequest | Connector rule with updated data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorCustomizersAPI.PutConnectorCustomizer(context.Background(), id).ConnectorCustomizerUpdateRequest(connectorCustomizerUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorCustomizersAPI.PutConnectorCustomizer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutConnectorCustomizer`: ConnectorCustomizerUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorCustomizersAPI.PutConnectorCustomizer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector customizer to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutConnectorCustomizerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectorCustomizerUpdateRequest** | [**ConnectorCustomizerUpdateRequest**](ConnectorCustomizerUpdateRequest.md) | Connector rule with updated data. | 

### Return type

[**ConnectorCustomizerUpdateResponse**](ConnectorCustomizerUpdateResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

