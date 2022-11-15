# \SourcesAccountSchemaApi

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountSchemaAttribute**](SourcesAccountSchemaApi.md#CreateAccountSchemaAttribute) | **Post** /cc/api/source/createSchemaAttribute/{id} | Create Account Schema Attribute
[**DeleteAccountSchemaAttribute**](SourcesAccountSchemaApi.md#DeleteAccountSchemaAttribute) | **Post** /cc/api/source/deleteSchemaAttribute/{id} | Delete Account Schema Attribute
[**GetSourceAccountSchema**](SourcesAccountSchemaApi.md#GetSourceAccountSchema) | **Get** /cc/api/source/getAccountSchema/{id} | Get Account Schema
[**UpdateAccountSchemaAttribute**](SourcesAccountSchemaApi.md#UpdateAccountSchemaAttribute) | **Post** /cc/api/source/updateSchemaAttributes/{id} | Update Account Schema Attribute



## CreateAccountSchemaAttribute

> CreateAccountSchemaAttribute(ctx, id).ObjectType(objectType).Entitlement(entitlement).Multi(multi).Names(names).Type_(type_).Description(description).Execute()

Create Account Schema Attribute



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
    id := "id_example" // string | 
    objectType := "objectType_example" // string |  (optional)
    entitlement := true // bool |  (optional)
    multi := true // bool |  (optional)
    names := "names_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    description := "description_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAccountSchemaApi.CreateAccountSchemaAttribute(context.Background(), id).ObjectType(objectType).Entitlement(entitlement).Multi(multi).Names(names).Type_(type_).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAccountSchemaApi.CreateAccountSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** |  | 
 **entitlement** | **bool** |  | 
 **multi** | **bool** |  | 
 **names** | **string** |  | 
 **type_** | **string** |  | 
 **description** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountSchemaAttribute

> DeleteAccountSchemaAttribute(ctx, id).ObjectType(objectType).Names(names).Execute()

Delete Account Schema Attribute



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
    id := "id_example" // string | 
    objectType := "objectType_example" // string |  (optional)
    names := "names_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAccountSchemaApi.DeleteAccountSchemaAttribute(context.Background(), id).ObjectType(objectType).Names(names).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAccountSchemaApi.DeleteAccountSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** |  | 
 **names** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceAccountSchema

> GetSourceAccountSchema(ctx, id).Page(page).Start(start).Limit(limit).Execute()

Get Account Schema

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
    id := "id_example" // string | 
    page := int32(1) // int32 |  (optional)
    start := int32(0) // int32 |  (optional)
    limit := int32(25) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAccountSchemaApi.GetSourceAccountSchema(context.Background(), id).Page(page).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAccountSchemaApi.GetSourceAccountSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceAccountSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountSchemaAttribute

> string UpdateAccountSchemaAttribute(ctx, id).ObjectType(objectType).FieldName(fieldName).FieldValue(fieldValue).Names(names).Execute()

Update Account Schema Attribute



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
    id := "id_example" // string | 
    objectType := "objectType_example" // string |  (optional)
    fieldName := "fieldName_example" // string |  (optional)
    fieldValue := true // bool |  (optional)
    names := "names_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAccountSchemaApi.UpdateAccountSchemaAttribute(context.Background(), id).ObjectType(objectType).FieldName(fieldName).FieldValue(fieldValue).Names(names).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAccountSchemaApi.UpdateAccountSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountSchemaAttribute`: string
    fmt.Fprintf(os.Stdout, "Response from `SourcesAccountSchemaApi.UpdateAccountSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** |  | 
 **fieldName** | **string** |  | 
 **fieldValue** | **bool** |  | 
 **names** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

