# \SourcesAttributeSyncApi

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceAttributeSyncConfig**](SourcesAttributeSyncApi.md#GetSourceAttributeSyncConfig) | **Get** /cc/api/source/getAttributeSyncConfig/{id} | Get Source Attribute Sync Config
[**SetSourceAttributeSyncConfig**](SourcesAttributeSyncApi.md#SetSourceAttributeSyncConfig) | **Post** /cc/api/source/setAttributeSyncConfig/{id} | Set Source Attribute Sync Config



## GetSourceAttributeSyncConfig

> GetSourceAttributeSyncConfig(ctx, id).Page(page).Start(start).Limit(limit).Execute()

Get Source Attribute Sync Config

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
    resp, r, err := apiClient.SourcesAttributeSyncApi.GetSourceAttributeSyncConfig(context.Background(), id).Page(page).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAttributeSyncApi.GetSourceAttributeSyncConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetSourceAttributeSyncConfigRequest struct via the builder pattern


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


## SetSourceAttributeSyncConfig

> SetSourceAttributeSyncConfig(ctx, id).ContentType(contentType).Body(body).Execute()

Set Source Attribute Sync Config

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
    contentType := "application/json" // string |  (optional)
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcesAttributeSyncApi.SetSourceAttributeSyncConfig(context.Background(), id).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcesAttributeSyncApi.SetSourceAttributeSyncConfig``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetSourceAttributeSyncConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **body** | **string** |  | 

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

