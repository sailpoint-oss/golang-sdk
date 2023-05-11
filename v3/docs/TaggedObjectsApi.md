# \TaggedObjectsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaggedObject**](TaggedObjectsApi.md#DeleteTaggedObject) | **Delete** /tagged-objects/{type}/{id} | Delete Tagged Object
[**DeleteTagsToManyObject**](TaggedObjectsApi.md#DeleteTagsToManyObject) | **Post** /tagged-objects/bulk-remove | Remove Tags from Multiple Objects
[**GetTaggedObject**](TaggedObjectsApi.md#GetTaggedObject) | **Get** /tagged-objects/{type}/{id} | Get Tagged Object
[**ListTaggedObjects**](TaggedObjectsApi.md#ListTaggedObjects) | **Get** /tagged-objects | List Tagged Objects
[**ListTaggedObjectsByType**](TaggedObjectsApi.md#ListTaggedObjectsByType) | **Get** /tagged-objects/{type} | List Tagged Objects
[**PutTaggedObject**](TaggedObjectsApi.md#PutTaggedObject) | **Put** /tagged-objects/{type}/{id} | Update Tagged Object
[**SetTagToObject**](TaggedObjectsApi.md#SetTagToObject) | **Post** /tagged-objects | Add Tag to Object
[**SetTagsToManyObjects**](TaggedObjectsApi.md#SetTagsToManyObjects) | **Post** /tagged-objects/bulk-add | Tag Multiple Objects



## DeleteTaggedObject

> DeleteTaggedObject(ctx, type_, id).Execute()

Delete Tagged Object



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
    type_ := "ROLE" // string | The type of tagged object to delete.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.DeleteTaggedObject(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.DeleteTaggedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of tagged object to delete. | 
**id** | **string** | The ID of the object reference to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaggedObjectRequest struct via the builder pattern


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


## DeleteTagsToManyObject

> DeleteTagsToManyObject(ctx).BulkTaggedObject(bulkTaggedObject).Execute()

Remove Tags from Multiple Objects



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
    bulkTaggedObject := *openapiclient.NewBulkTaggedObject() // BulkTaggedObject | Supported object types are ROLE, IDENTITY and SOD_POLICY.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.DeleteTagsToManyObject(context.Background()).BulkTaggedObject(bulkTaggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.DeleteTagsToManyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsToManyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTaggedObject** | [**BulkTaggedObject**](BulkTaggedObject.md) | Supported object types are ROLE, IDENTITY and SOD_POLICY. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaggedObject

> TaggedObject GetTaggedObject(ctx, type_, id).Execute()

Get Tagged Object



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
    type_ := "ROLE" // string | The type of tagged object to retrieve.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.GetTaggedObject(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.GetTaggedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaggedObject`: TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.GetTaggedObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of tagged object to retrieve. | 
**id** | **string** | The ID of the object reference to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaggedObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaggedObject**](TaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaggedObjects

> []TaggedObject ListTaggedObjects(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

List Tagged Objects



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "tagName eq "BU_FINANCE"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq, in*  **objectRef.type**: *eq, in*  **tagName**: *eq, in* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.ListTaggedObjects(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.ListTaggedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggedObjects`: []TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.ListTaggedObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaggedObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq, in*  **objectRef.type**: *eq, in*  **tagName**: *eq, in* | 

### Return type

[**[]TaggedObject**](TaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaggedObjectsByType

> []TaggedObject ListTaggedObjectsByType(ctx, type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

List Tagged Objects



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
    type_ := "ROLE" // string | The type of tagged object to retrieve.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "objectRef.id eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq*  **objectRef.type**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.ListTaggedObjectsByType(context.Background(), type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.ListTaggedObjectsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggedObjectsByType`: []TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.ListTaggedObjectsByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of tagged object to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaggedObjectsByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq*  **objectRef.type**: *eq* | 

### Return type

[**[]TaggedObject**](TaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTaggedObject

> TaggedObject PutTaggedObject(ctx, type_, id).TaggedObject(taggedObject).Execute()

Update Tagged Object



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
    type_ := "ROLE" // string | The type of tagged object to update.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to update.
    taggedObject := *openapiclient.NewTaggedObject() // TaggedObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.PutTaggedObject(context.Background(), type_, id).TaggedObject(taggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.PutTaggedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTaggedObject`: TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.PutTaggedObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The type of tagged object to update. | 
**id** | **string** | The ID of the object reference to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTaggedObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **taggedObject** | [**TaggedObject**](TaggedObject.md) |  | 

### Return type

[**TaggedObject**](TaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTagToObject

> TaggedObject SetTagToObject(ctx).TaggedObject(taggedObject).Execute()

Add Tag to Object



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
    taggedObject := *openapiclient.NewTaggedObject() // TaggedObject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.SetTagToObject(context.Background()).TaggedObject(taggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.SetTagToObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTagToObject`: TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.SetTagToObject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTagToObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taggedObject** | [**TaggedObject**](TaggedObject.md) |  | 

### Return type

[**TaggedObject**](TaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTagsToManyObjects

> BulkTaggedObject SetTagsToManyObjects(ctx).BulkTaggedObject(bulkTaggedObject).Execute()

Tag Multiple Objects



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
    bulkTaggedObject := *openapiclient.NewBulkTaggedObject() // BulkTaggedObject | Supported object types are ROLE, IDENTITY and SOD_POLICY.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaggedObjectsApi.SetTagsToManyObjects(context.Background()).BulkTaggedObject(bulkTaggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsApi.SetTagsToManyObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTagsToManyObjects`: BulkTaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsApi.SetTagsToManyObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTagsToManyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTaggedObject** | [**BulkTaggedObject**](BulkTaggedObject.md) | Supported object types are ROLE, IDENTITY and SOD_POLICY. | 

### Return type

[**BulkTaggedObject**](BulkTaggedObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

