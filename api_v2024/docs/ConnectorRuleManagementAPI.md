# \ConnectorRuleManagementAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorRule**](ConnectorRuleManagementAPI.md#CreateConnectorRule) | **Post** /connector-rules | Create Connector Rule
[**DeleteConnectorRule**](ConnectorRuleManagementAPI.md#DeleteConnectorRule) | **Delete** /connector-rules/{id} | Delete a Connector-Rule
[**GetConnectorRule**](ConnectorRuleManagementAPI.md#GetConnectorRule) | **Get** /connector-rules/{id} | Connector-Rule by ID
[**GetConnectorRuleList**](ConnectorRuleManagementAPI.md#GetConnectorRuleList) | **Get** /connector-rules | List Connector Rules
[**UpdateConnectorRule**](ConnectorRuleManagementAPI.md#UpdateConnectorRule) | **Put** /connector-rules/{id} | Update a Connector Rule
[**ValidateConnectorRule**](ConnectorRuleManagementAPI.md#ValidateConnectorRule) | **Post** /connector-rules/validate | Validate Connector Rule



## CreateConnectorRule

> ConnectorRuleResponse CreateConnectorRule(ctx).XSailPointExperimental(xSailPointExperimental).ConnectorRuleCreateRequest(connectorRuleCreateRequest).Execute()

Create Connector Rule



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
	connectorRuleCreateRequest := *openapiclient.NewConnectorRuleCreateRequest("WebServiceBeforeOperationRule", "BuildMap", *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;")) // ConnectorRuleCreateRequest | The connector rule to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorRuleManagementAPI.CreateConnectorRule(context.Background()).XSailPointExperimental(xSailPointExperimental).ConnectorRuleCreateRequest(connectorRuleCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.CreateConnectorRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorRule`: ConnectorRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementAPI.CreateConnectorRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **connectorRuleCreateRequest** | [**ConnectorRuleCreateRequest**](ConnectorRuleCreateRequest.md) | The connector rule to create | 

### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorRule

> DeleteConnectorRule(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Delete a Connector-Rule



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to delete
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectorRuleManagementAPI.DeleteConnectorRule(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.DeleteConnectorRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector rule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorRuleRequest struct via the builder pattern


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


## GetConnectorRule

> ConnectorRuleResponse GetConnectorRule(ctx, id).XSailPointExperimental(xSailPointExperimental).Execute()

Connector-Rule by ID



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to retrieve
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorRuleManagementAPI.GetConnectorRule(context.Background(), id).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.GetConnectorRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorRule`: ConnectorRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementAPI.GetConnectorRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector rule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorRuleList

> []ConnectorRuleResponse GetConnectorRuleList(ctx).XSailPointExperimental(xSailPointExperimental).Execute()

List Connector Rules



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorRuleManagementAPI.GetConnectorRuleList(context.Background()).XSailPointExperimental(xSailPointExperimental).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.GetConnectorRuleList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorRuleList`: []ConnectorRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementAPI.GetConnectorRuleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

### Return type

[**[]ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectorRule

> ConnectorRuleResponse UpdateConnectorRule(ctx, id).XSailPointExperimental(xSailPointExperimental).ConnectorRuleUpdateRequest(connectorRuleUpdateRequest).Execute()

Update a Connector Rule



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
	id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to update
	xSailPointExperimental := "true" // string | Use this header to enable this experimental API. (default to "true")
	connectorRuleUpdateRequest := *openapiclient.NewConnectorRuleUpdateRequest("WebServiceBeforeOperationRule", "BuildMap", *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;"), "8113d48c0b914f17b4c6072d4dcb9dfe") // ConnectorRuleUpdateRequest | The connector rule with updated data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorRuleManagementAPI.UpdateConnectorRule(context.Background(), id).XSailPointExperimental(xSailPointExperimental).ConnectorRuleUpdateRequest(connectorRuleUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.UpdateConnectorRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConnectorRule`: ConnectorRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementAPI.UpdateConnectorRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the connector rule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **connectorRuleUpdateRequest** | [**ConnectorRuleUpdateRequest**](ConnectorRuleUpdateRequest.md) | The connector rule with updated data | 

### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnectorRule

> ConnectorRuleValidationResponse ValidateConnectorRule(ctx).XSailPointExperimental(xSailPointExperimental).SourceCode(sourceCode).Execute()

Validate Connector Rule



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
	sourceCode := *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;") // SourceCode | The code to validate

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectorRuleManagementAPI.ValidateConnectorRule(context.Background()).XSailPointExperimental(xSailPointExperimental).SourceCode(sourceCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementAPI.ValidateConnectorRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateConnectorRule`: ConnectorRuleValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementAPI.ValidateConnectorRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]
 **sourceCode** | [**SourceCode**](SourceCode.md) | The code to validate | 

### Return type

[**ConnectorRuleValidationResponse**](ConnectorRuleValidationResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

