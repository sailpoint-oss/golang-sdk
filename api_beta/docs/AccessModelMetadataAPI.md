# \AccessModelMetadataAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessModelMetadataAttribute**](AccessModelMetadataAPI.md#GetAccessModelMetadataAttribute) | **Get** /access-model-metadata/attributes/{key} | Get Access Model Metadata Attribute
[**GetAccessModelMetadataAttributeValue**](AccessModelMetadataAPI.md#GetAccessModelMetadataAttributeValue) | **Get** /access-model-metadata/attributes/{key}/values/{value} | Get Access Model Metadata Value
[**ListAccessModelMetadataAttribute**](AccessModelMetadataAPI.md#ListAccessModelMetadataAttribute) | **Get** /access-model-metadata/attributes | List Access Model Metadata Attributes
[**ListAccessModelMetadataAttributeValue**](AccessModelMetadataAPI.md#ListAccessModelMetadataAttributeValue) | **Get** /access-model-metadata/attributes/{key}/values | List Access Model Metadata Values



## GetAccessModelMetadataAttribute

> AttributeDTO GetAccessModelMetadataAttribute(ctx, key).Execute()

Get Access Model Metadata Attribute



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
    key := "iscPrivacy" // string | Technical name of the Attribute.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessModelMetadataAPI.GetAccessModelMetadataAttribute(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.GetAccessModelMetadataAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessModelMetadataAttribute`: AttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.GetAccessModelMetadataAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessModelMetadataAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttributeDTO**](AttributeDTO.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessModelMetadataAttributeValue

> AttributeValueDTO GetAccessModelMetadataAttributeValue(ctx, key, value).Execute()

Get Access Model Metadata Value



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
    key := "iscPrivacy" // string | Technical name of the Attribute.
    value := "public" // string | Technical name of the Attribute value.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue(context.Background(), key, value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessModelMetadataAttributeValue`: AttributeValueDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 
**value** | **string** | Technical name of the Attribute value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessModelMetadataAttributeValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttributeValueDTO**](AttributeValueDTO.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessModelMetadataAttribute

> []AttributeDTO ListAccessModelMetadataAttribute(ctx).Filters(filters).Execute()

List Access Model Metadata Attributes



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
    filters := "name eq "Privacy"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq*  **type**: *eq*  **status**: *eq*  **objectTypes**: *eq*  Supported composite operators: *and* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessModelMetadataAPI.ListAccessModelMetadataAttribute(context.Background()).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.ListAccessModelMetadataAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessModelMetadataAttribute`: []AttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.ListAccessModelMetadataAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessModelMetadataAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq*  **type**: *eq*  **status**: *eq*  **objectTypes**: *eq*  Supported composite operators: *and* | 

### Return type

[**[]AttributeDTO**](AttributeDTO.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessModelMetadataAttributeValue

> []AttributeValueDTO ListAccessModelMetadataAttributeValue(ctx, key).Execute()

List Access Model Metadata Values



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
    key := "iscPrivacy" // string | Technical name of the Attribute.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue(context.Background(), key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessModelMetadataAttributeValue`: []AttributeValueDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessModelMetadataAttributeValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AttributeValueDTO**](AttributeValueDTO.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

