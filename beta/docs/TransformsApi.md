# \TransformsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransform**](TransformsApi.md#CreateTransform) | **Post** /transforms | Create transform
[**DeleteTransform**](TransformsApi.md#DeleteTransform) | **Delete** /transforms/{id} | Delete a transform
[**GetTransform**](TransformsApi.md#GetTransform) | **Get** /transforms/{id} | Transform by ID
[**ListTransforms**](TransformsApi.md#ListTransforms) | **Get** /transforms | List transforms
[**UpdateTransform**](TransformsApi.md#UpdateTransform) | **Put** /transforms/{id} | Update a transform



## CreateTransform

> Transform CreateTransform(ctx).Transform(transform).Execute()

Create transform



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
    transform := *openapiclient.NewTransform("Timestamp To Date", "concat", map[string]interface{}({inputFormat=MMM dd yyyy, HH:mm:ss.SSS, outputFormat=yyyy/dd/MM})) // Transform | The transform to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsApi.CreateTransform(context.Background()).Transform(transform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.CreateTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransform`: Transform
    fmt.Fprintf(os.Stdout, "Response from `TransformsApi.CreateTransform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transform** | [**Transform**](Transform.md) | The transform to be created. | 

### Return type

[**Transform**](Transform.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTransform

> DeleteTransform(ctx, id).Execute()

Delete a transform



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
    id := "2c9180835d2e5168015d32f890ca1581" // string | ID of the transform to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsApi.DeleteTransform(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.DeleteTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the transform to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransformRequest struct via the builder pattern


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


## GetTransform

> Transform GetTransform(ctx, id).Execute()

Transform by ID



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
    id := "2c9180835d2e5168015d32f890ca1581" // string | ID of the transform to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsApi.GetTransform(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.GetTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransform`: Transform
    fmt.Fprintf(os.Stdout, "Response from `TransformsApi.GetTransform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the transform to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transform**](Transform.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransforms

> []Transform ListTransforms(ctx).Offset(offset).Limit(limit).Count(count).Name(name).Filters(filters).Execute()

List transforms



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
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    name := "ExampleTransformName123" // string | Name of the transform to retrieve from the list. (optional)
    filters := "name eq ExampleTransformName123" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **internal**: *eq* **name**: *eq*, *sw* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsApi.ListTransforms(context.Background()).Offset(offset).Limit(limit).Count(count).Name(name).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.ListTransforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransforms`: []Transform
    fmt.Fprintf(os.Stdout, "Response from `TransformsApi.ListTransforms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransformsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **name** | **string** | Name of the transform to retrieve from the list. | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **internal**: *eq* **name**: *eq*, *sw* | 

### Return type

[**[]Transform**](Transform.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransform

> Transform UpdateTransform(ctx, id).Transform(transform).Execute()

Update a transform



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
    id := "2c9180835d2e5168015d32f890ca1581" // string | ID of the transform to update
    transform := *openapiclient.NewTransform("Timestamp To Date", "concat", map[string]interface{}({inputFormat=MMM dd yyyy, HH:mm:ss.SSS, outputFormat=yyyy/dd/MM})) // Transform | The updated transform object (must include \"name\", \"type\", and \"attributes\" fields). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsApi.UpdateTransform(context.Background(), id).Transform(transform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsApi.UpdateTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransform`: Transform
    fmt.Fprintf(os.Stdout, "Response from `TransformsApi.UpdateTransform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the transform to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transform** | [**Transform**](Transform.md) | The updated transform object (must include \&quot;name\&quot;, \&quot;type\&quot;, and \&quot;attributes\&quot; fields). | 

### Return type

[**Transform**](Transform.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

