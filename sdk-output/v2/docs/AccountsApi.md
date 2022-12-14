# \AccountsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /accounts | Creates a new account on a flat-file source.
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /accounts/{id} | Deletes an existing account from a flat-file source.
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /accounts/{id} | Retrieves the Account by Id.
[**GetAccounts**](AccountsApi.md#GetAccounts) | **Get** /accounts | Retrieves list of Accounts for a given source.
[**PatchAccount**](AccountsApi.md#PatchAccount) | **Patch** /accounts/{id} | Updates an existing account from a flat-file source.
[**PutAccount**](AccountsApi.md#PutAccount) | **Put** /accounts/{id} | Updates an existing account from a flat-file source by replacing all values.



## CreateAccount

> CreateAccount(ctx).SourceId(sourceId).DynamicSchemaEto(dynamicSchemaEto).Execute()

Creates a new account on a flat-file source.



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
    sourceId := "sourceId_example" // string | ID of a flat-file source into which the new account will be created.
    dynamicSchemaEto := *openapiclient.NewDynamicSchemaEto() // DynamicSchemaEto | Attribute values for the new account. The schema and required attributes are dictated by the source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.CreateAccount(context.Background()).SourceId(sourceId).DynamicSchemaEto(dynamicSchemaEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **string** | ID of a flat-file source into which the new account will be created. | 
 **dynamicSchemaEto** | [**DynamicSchemaEto**](DynamicSchemaEto.md) | Attribute values for the new account. The schema and required attributes are dictated by the source. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> DeleteAccount(ctx, id).Execute()

Deletes an existing account from a flat-file source.



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
    id := "id_example" // string | ID of an Account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.DeleteAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, id).Execute()

Retrieves the Account by Id.



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
    id := "id_example" // string | ID of an Account

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> []Account GetAccounts(ctx).SourceId(sourceId).Sort(sort).Offset(offset).Limit(limit).Execute()

Retrieves list of Accounts for a given source.



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
    sourceId := int32(56) // int32 | ID of a source for the Account list
    sort := "sort_example" // string | One or more attributes on which to sort, each separated by a ','. Prefix with a minus sign (ex. -dateCreated) for descending sort. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.GetAccounts(context.Background()).SourceId(sourceId).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccounts`: []Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceId** | **int32** | ID of a source for the Account list | 
 **sort** | **string** | One or more attributes on which to sort, each separated by a &#39;,&#39;. Prefix with a minus sign (ex. -dateCreated) for descending sort. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAccount

> PatchAccount(ctx, id).DynamicSchemaEto(dynamicSchemaEto).Execute()

Updates an existing account from a flat-file source.



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
    id := "id_example" // string | ID of an Account
    dynamicSchemaEto := *openapiclient.NewDynamicSchemaEto() // DynamicSchemaEto | Attribute values for the account. The schema and required attributes are dictated by the source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.PatchAccount(context.Background(), id).DynamicSchemaEto(dynamicSchemaEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PatchAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicSchemaEto** | [**DynamicSchemaEto**](DynamicSchemaEto.md) | Attribute values for the account. The schema and required attributes are dictated by the source. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAccount

> PutAccount(ctx, id).DynamicSchemaEto(dynamicSchemaEto).Execute()

Updates an existing account from a flat-file source by replacing all values.



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
    id := "id_example" // string | ID of an Account
    dynamicSchemaEto := *openapiclient.NewDynamicSchemaEto() // DynamicSchemaEto | Attribute values for the account. The schema and required attributes are dictated by the source.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsApi.PutAccount(context.Background(), id).DynamicSchemaEto(dynamicSchemaEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.PutAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of an Account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicSchemaEto** | [**DynamicSchemaEto**](DynamicSchemaEto.md) | Attribute values for the account. The schema and required attributes are dictated by the source. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

