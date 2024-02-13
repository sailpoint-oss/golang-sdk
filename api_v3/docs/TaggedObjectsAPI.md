# \TaggedObjectsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTaggedObject**](TaggedObjectsAPI.md#DeleteTaggedObject) | **Delete** /tagged-objects/{type}/{id} | Delete Tagged Object
[**DeleteTagsToManyObject**](TaggedObjectsAPI.md#DeleteTagsToManyObject) | **Post** /tagged-objects/bulk-remove | Remove Tags from Multiple Objects
[**GetTaggedObject**](TaggedObjectsAPI.md#GetTaggedObject) | **Get** /tagged-objects/{type}/{id} | Get Tagged Object
[**ListTaggedObjects**](TaggedObjectsAPI.md#ListTaggedObjects) | **Get** /tagged-objects | List Tagged Objects
[**ListTaggedObjectsByType**](TaggedObjectsAPI.md#ListTaggedObjectsByType) | **Get** /tagged-objects/{type} | List Tagged Objects by Type
[**PutTaggedObject**](TaggedObjectsAPI.md#PutTaggedObject) | **Put** /tagged-objects/{type}/{id} | Update Tagged Object
[**SetTagToObject**](TaggedObjectsAPI.md#SetTagToObject) | **Post** /tagged-objects | Add Tag to Object
[**SetTagsToManyObjects**](TaggedObjectsAPI.md#SetTagsToManyObjects) | **Post** /tagged-objects/bulk-add | Tag Multiple Objects



## Delete Tagged Object

> DeleteTaggedObject(ctx, type_, id).Execute()

This deletes a tagged object for the specified type.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    type_ := "ROLE" // string | The type of tagged object to delete.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to delete.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.TaggedObjectsAPI.DeleteTaggedObject(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.DeleteTaggedObject``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Remove Tags from Multiple Objects

> DeleteTagsToManyObject(ctx).BulkTaggedObject(bulkTaggedObject).Execute()

This API removes tags from multiple objects.

A token with API, CERT_ADMIN, ORG_ADMIN, REPORT_ADMIN, ROLE_ADMIN, ROLE_SUBADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    bulkTaggedObject := *sailpoint.NewBulkTaggedObject() // BulkTaggedObject | Supported object types are ACCESS_PROFILE, APPLICATION, CAMPAIGN, ENTITLEMENT, IDENTITY, ROLE, SOD_POLICY, SOURCE.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.TaggedObjectsAPI.DeleteTagsToManyObject(context.Background()).BulkTaggedObject(bulkTaggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.DeleteTagsToManyObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagsToManyObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTaggedObject** | [**BulkTaggedObject**](BulkTaggedObject.md) | Supported object types are ACCESS_PROFILE, APPLICATION, CAMPAIGN, ENTITLEMENT, IDENTITY, ROLE, SOD_POLICY, SOURCE. | 

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


## Get Tagged Object

> TaggedObject GetTaggedObject(ctx, type_, id).Execute()

This gets a tagged object for the specified type.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    type_ := "ROLE" // string | The type of tagged object to retrieve.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to retrieve.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.TaggedObjectsAPI.GetTaggedObject(context.Background(), type_, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.GetTaggedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaggedObject`: TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsAPI.GetTaggedObject`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Tagged Objects

> []TaggedObject ListTaggedObjects(ctx).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

This API returns a list of all tagged objects.

Any authenticated token may be used to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "tagName eq "BU_FINANCE"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq, in*  **objectRef.type**: *eq, in*  **tagName**: *eq, in* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.TaggedObjectsAPI.ListTaggedObjects(context.Background()).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.ListTaggedObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggedObjects`: []TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsAPI.ListTaggedObjects`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Tagged Objects by Type

> []TaggedObject ListTaggedObjectsByType(ctx, type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

This API returns a list of all tagged objects by type.

Any authenticated token may be used to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    type_ := "ROLE" // string | The type of tagged object to retrieve.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "objectRef.id eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **objectRef.id**: *eq*  **objectRef.type**: *eq* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.TaggedObjectsAPI.ListTaggedObjectsByType(context.Background(), type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.ListTaggedObjectsByType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaggedObjectsByType`: []TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsAPI.ListTaggedObjectsByType`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Tagged Object

> TaggedObject PutTaggedObject(ctx, type_, id).TaggedObject(taggedObject).Execute()

This updates a tagged object for the specified type.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    type_ := "ROLE" // string | The type of tagged object to update.
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the object reference to update.
    taggedObject := *sailpoint.NewTaggedObject() // TaggedObject | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.TaggedObjectsAPI.PutTaggedObject(context.Background(), type_, id).TaggedObject(taggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.PutTaggedObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutTaggedObject`: TaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsAPI.PutTaggedObject`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Add Tag to Object

> SetTagToObject(ctx).TaggedObject(taggedObject).Execute()

This adds a tag to an object.

Any authenticated token may be used to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    taggedObject := *sailpoint.NewTaggedObject() // TaggedObject | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.TaggedObjectsAPI.SetTagToObject(context.Background()).TaggedObject(taggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.SetTagToObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTagToObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taggedObject** | [**TaggedObject**](TaggedObject.md) |  | 

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


## Tag Multiple Objects

> BulkTaggedObject SetTagsToManyObjects(ctx).BulkTaggedObject(bulkTaggedObject).Execute()

This API adds tags to multiple objects.

A token with API, CERT_ADMIN, ORG_ADMIN, REPORT_ADMIN, ROLE_ADMIN, ROLE_SUBADMIN, SOURCE_ADMIN, or SOURCE_SUBADMIN authority is required to call this API.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    bulkTaggedObject := *sailpoint.NewBulkTaggedObject() // BulkTaggedObject | Supported object types are ACCESS_PROFILE, APPLICATION, CAMPAIGN, ENTITLEMENT, IDENTITY, ROLE, SOD_POLICY, SOURCE.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.TaggedObjectsAPI.SetTagsToManyObjects(context.Background()).BulkTaggedObject(bulkTaggedObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggedObjectsAPI.SetTagsToManyObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetTagsToManyObjects`: BulkTaggedObject
    fmt.Fprintf(os.Stdout, "Response from `TaggedObjectsAPI.SetTagsToManyObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetTagsToManyObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkTaggedObject** | [**BulkTaggedObject**](BulkTaggedObject.md) | Supported object types are ACCESS_PROFILE, APPLICATION, CAMPAIGN, ENTITLEMENT, IDENTITY, ROLE, SOD_POLICY, SOURCE. | 

### Return type

[**BulkTaggedObject**](BulkTaggedObject.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

