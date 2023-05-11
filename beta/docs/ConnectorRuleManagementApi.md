# \ConnectorRuleManagementApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorRule**](ConnectorRuleManagementApi.md#CreateConnectorRule) | **Post** /connector-rules | Create Connector Rule
[**DeleteConnectorRule**](ConnectorRuleManagementApi.md#DeleteConnectorRule) | **Delete** /connector-rules/{id} | Delete a Connector-Rule
[**GetConnectorRule**](ConnectorRuleManagementApi.md#GetConnectorRule) | **Get** /connector-rules/{id} | Connector-Rule by ID
[**GetConnectorRuleList**](ConnectorRuleManagementApi.md#GetConnectorRuleList) | **Get** /connector-rules | List Connector Rules
[**UpdateConnectorRule**](ConnectorRuleManagementApi.md#UpdateConnectorRule) | **Put** /connector-rules/{id} | Update a Connector Rule
[**ValidateConnectorRule**](ConnectorRuleManagementApi.md#ValidateConnectorRule) | **Post** /connector-rules/validate | Validate Connector Rule



## CreateConnectorRule

> ConnectorRuleResponse CreateConnectorRule(ctx).ConnectorRuleCreateRequest(connectorRuleCreateRequest).Execute()

Create Connector Rule



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
    connectorRuleCreateRequest := *openapiclient.NewConnectorRuleCreateRequest("WebServiceBeforeOperationRule", "BuildMap", *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;")) // ConnectorRuleCreateRequest | The connector rule to create

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorRuleManagementApi.CreateConnectorRule(context.Background()).ConnectorRuleCreateRequest(connectorRuleCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.CreateConnectorRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnectorRule`: ConnectorRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementApi.CreateConnectorRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectorRuleCreateRequest** | [**ConnectorRuleCreateRequest**](ConnectorRuleCreateRequest.md) | The connector rule to create | 

### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorRule

> DeleteConnectorRule(ctx, id).Execute()

Delete a Connector-Rule



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorRuleManagementApi.DeleteConnectorRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.DeleteConnectorRule``: %v\n", err)
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


## GetConnectorRule

> ConnectorRuleResponse GetConnectorRule(ctx, id).Execute()

Connector-Rule by ID



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorRuleManagementApi.GetConnectorRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.GetConnectorRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorRule`: ConnectorRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementApi.GetConnectorRule`: %v\n", resp)
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


### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorRuleList

> []ConnectorRuleResponse GetConnectorRuleList(ctx).Execute()

List Connector Rules



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
    resp, r, err := apiClient.ConnectorRuleManagementApi.GetConnectorRuleList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.GetConnectorRuleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorRuleList`: []ConnectorRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementApi.GetConnectorRuleList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorRuleListRequest struct via the builder pattern


### Return type

[**[]ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnectorRule

> ConnectorRuleResponse UpdateConnectorRule(ctx, id).ConnectorRuleUpdateRequest(connectorRuleUpdateRequest).Execute()

Update a Connector Rule



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
    id := "8c190e6787aa4ed9a90bd9d5344523fb" // string | ID of the connector rule to update
    connectorRuleUpdateRequest := *openapiclient.NewConnectorRuleUpdateRequest("8113d48c0b914f17b4c6072d4dcb9dfe", "WebServiceBeforeOperationRule", "BuildMap", *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;")) // ConnectorRuleUpdateRequest | The connector rule with updated data (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorRuleManagementApi.UpdateConnectorRule(context.Background(), id).ConnectorRuleUpdateRequest(connectorRuleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.UpdateConnectorRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnectorRule`: ConnectorRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementApi.UpdateConnectorRule`: %v\n", resp)
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

 **connectorRuleUpdateRequest** | [**ConnectorRuleUpdateRequest**](ConnectorRuleUpdateRequest.md) | The connector rule with updated data | 

### Return type

[**ConnectorRuleResponse**](ConnectorRuleResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateConnectorRule

> ConnectorRuleValidationResponse ValidateConnectorRule(ctx).SourceCode(sourceCode).Execute()

Validate Connector Rule



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
    sourceCode := *openapiclient.NewSourceCode("1.0", "return "Mr. " + firstName;") // SourceCode | The code to validate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConnectorRuleManagementApi.ValidateConnectorRule(context.Background()).SourceCode(sourceCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectorRuleManagementApi.ValidateConnectorRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateConnectorRule`: ConnectorRuleValidationResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectorRuleManagementApi.ValidateConnectorRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateConnectorRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceCode** | [**SourceCode**](SourceCode.md) | The code to validate | 

### Return type

[**ConnectorRuleValidationResponse**](ConnectorRuleValidationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

