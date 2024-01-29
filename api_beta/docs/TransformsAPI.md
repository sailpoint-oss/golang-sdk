# \TransformsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransform**](TransformsAPI.md#CreateTransform) | **Post** /transforms | Create transform
[**DeleteTransform**](TransformsAPI.md#DeleteTransform) | **Delete** /transforms/{id} | Delete a transform
[**GetTransform**](TransformsAPI.md#GetTransform) | **Get** /transforms/{id} | Transform by ID
[**ListTransforms**](TransformsAPI.md#ListTransforms) | **Get** /transforms | List transforms
[**UpdateTransform**](TransformsAPI.md#UpdateTransform) | **Put** /transforms/{id} | Update a transform



## CreateTransform

> TransformRead CreateTransform(ctx).Transform(transform).Execute()

Create transform



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    transform := *openapiclient.NewTransform("Timestamp To Date", "dateFormat", map[string]interface{}(123)) // Transform | The transform to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsAPI.CreateTransform(context.Background()).Transform(transform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsAPI.CreateTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransform`: TransformRead
    fmt.Fprintf(os.Stdout, "Response from `TransformsAPI.CreateTransform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transform** | [**Transform**](Transform.md) | The transform to be created. | 

### Return type

[**TransformRead**](TransformRead.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "2cd78adghjkja34jh2b1hkjhasuecd" // string | ID of the transform to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TransformsAPI.DeleteTransform(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsAPI.DeleteTransform``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransform

> TransformRead GetTransform(ctx, id).Execute()

Transform by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "2cd78adghjkja34jh2b1hkjhasuecd" // string | ID of the transform to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsAPI.GetTransform(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsAPI.GetTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransform`: TransformRead
    fmt.Fprintf(os.Stdout, "Response from `TransformsAPI.GetTransform`: %v\n", resp)
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

[**TransformRead**](TransformRead.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransforms

> []TransformRead ListTransforms(ctx).Offset(offset).Limit(limit).Count(count).Name(name).Filters(filters).Execute()

List transforms



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    name := "ExampleTransformName123" // string | Name of the transform to retrieve from the list. (optional)
    filters := "name eq "Uppercase"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **internal**: *eq*  **name**: *eq, sw* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsAPI.ListTransforms(context.Background()).Offset(offset).Limit(limit).Count(count).Name(name).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsAPI.ListTransforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransforms`: []TransformRead
    fmt.Fprintf(os.Stdout, "Response from `TransformsAPI.ListTransforms`: %v\n", resp)
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
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **internal**: *eq*  **name**: *eq, sw* | 

### Return type

[**[]TransformRead**](TransformRead.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransform

> TransformRead UpdateTransform(ctx, id).Transform(transform).Execute()

Update a transform



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    id := "2cd78adghjkja34jh2b1hkjhasuecd" // string | ID of the transform to update
    transform := *openapiclient.NewTransform("Timestamp To Date", "dateFormat", map[string]interface{}(123)) // Transform | The updated transform object. Must include \"name\", \"type\", and \"attributes\" fields, but \"name\" and \"type\" must not be modified. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransformsAPI.UpdateTransform(context.Background(), id).Transform(transform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformsAPI.UpdateTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransform`: TransformRead
    fmt.Fprintf(os.Stdout, "Response from `TransformsAPI.UpdateTransform`: %v\n", resp)
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

 **transform** | [**Transform**](Transform.md) | The updated transform object. Must include \&quot;name\&quot;, \&quot;type\&quot;, and \&quot;attributes\&quot; fields, but \&quot;name\&quot; and \&quot;type\&quot; must not be modified. | 

### Return type

[**TransformRead**](TransformRead.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

