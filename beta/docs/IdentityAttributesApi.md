# \IdentityAttributesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityAttribute**](IdentityAttributesApi.md#CreateIdentityAttribute) | **Post** /identity-attributes | Create Identity Attribute
[**DeleteIdentityAttribute**](IdentityAttributesApi.md#DeleteIdentityAttribute) | **Delete** /identity-attributes/{name} | Delete Identity Attribute
[**DeleteIdentityAttributesInBulk**](IdentityAttributesApi.md#DeleteIdentityAttributesInBulk) | **Post** /identity-attributes/bulk-delete | Bulk delete Identity Attributes
[**GetIdentityAttribute**](IdentityAttributesApi.md#GetIdentityAttribute) | **Get** /identity-attributes/{name} | Get Identity Attribute
[**ListIdentityAttributes**](IdentityAttributesApi.md#ListIdentityAttributes) | **Get** /identity-attributes | List Identity Attributes
[**PutIdentityAttribute**](IdentityAttributesApi.md#PutIdentityAttribute) | **Put** /identity-attributes/{name} | Update Identity Attribute



## CreateIdentityAttribute

> IdentityAttribute CreateIdentityAttribute(ctx).IdentityAttribute(identityAttribute).Execute()

Create Identity Attribute



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
    identityAttribute := *openapiclient.NewIdentityAttribute() // IdentityAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.CreateIdentityAttribute(context.Background()).IdentityAttribute(identityAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.CreateIdentityAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityAttribute`: IdentityAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesApi.CreateIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityAttribute** | [**IdentityAttribute**](IdentityAttribute.md) |  | 

### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityAttribute

> DeleteIdentityAttribute(ctx, name).Execute()

Delete Identity Attribute



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
    name := "displayName" // string | The attribute's technical name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.DeleteIdentityAttribute(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.DeleteIdentityAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityAttributeRequest struct via the builder pattern


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


## DeleteIdentityAttributesInBulk

> DeleteIdentityAttributesInBulk(ctx).IdentityAttributeNames(identityAttributeNames).Execute()

Bulk delete Identity Attributes



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
    identityAttributeNames := *openapiclient.NewIdentityAttributeNames() // IdentityAttributeNames | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.DeleteIdentityAttributesInBulk(context.Background()).IdentityAttributeNames(identityAttributeNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.DeleteIdentityAttributesInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityAttributesInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityAttributeNames** | [**IdentityAttributeNames**](IdentityAttributeNames.md) |  | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityAttribute

> IdentityAttribute GetIdentityAttribute(ctx, name).Execute()

Get Identity Attribute



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
    name := "displayName" // string | The attribute's technical name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.GetIdentityAttribute(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.GetIdentityAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityAttribute`: IdentityAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesApi.GetIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityAttributes

> []IdentityAttribute ListIdentityAttributes(ctx).IncludeSystem(includeSystem).IncludeSilent(includeSilent).SearchableOnly(searchableOnly).Count(count).Execute()

List Identity Attributes



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
    includeSystem := false // bool | Include \"system\" attributes in the response. (optional) (default to false)
    includeSilent := false // bool | Include \"silent\" attributes in the response. (optional) (default to false)
    searchableOnly := false // bool | Include only \"searchable\" attributes in the response. (optional) (default to false)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.ListIdentityAttributes(context.Background()).IncludeSystem(includeSystem).IncludeSilent(includeSilent).SearchableOnly(searchableOnly).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.ListIdentityAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityAttributes`: []IdentityAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesApi.ListIdentityAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeSystem** | **bool** | Include \&quot;system\&quot; attributes in the response. | [default to false]
 **includeSilent** | **bool** | Include \&quot;silent\&quot; attributes in the response. | [default to false]
 **searchableOnly** | **bool** | Include only \&quot;searchable\&quot; attributes in the response. | [default to false]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IdentityAttribute**](IdentityAttribute.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIdentityAttribute

> IdentityAttribute PutIdentityAttribute(ctx, name).IdentityAttribute(identityAttribute).Execute()

Update Identity Attribute



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
    name := "displayName" // string | The attribute's technical name.
    identityAttribute := *openapiclient.NewIdentityAttribute() // IdentityAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityAttributesApi.PutIdentityAttribute(context.Background(), name).IdentityAttribute(identityAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityAttributesApi.PutIdentityAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIdentityAttribute`: IdentityAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityAttributesApi.PutIdentityAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The attribute&#39;s technical name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIdentityAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityAttribute** | [**IdentityAttribute**](IdentityAttribute.md) |  | 

### Return type

[**IdentityAttribute**](IdentityAttribute.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

