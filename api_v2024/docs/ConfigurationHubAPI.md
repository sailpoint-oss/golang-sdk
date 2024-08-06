# \ConfigurationHubAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectMapping**](ConfigurationHubAPI.md#CreateObjectMapping) | **Post** /configuration-hub/object-mappings/{sourceOrg} | Creates an object mapping
[**CreateObjectMappings**](ConfigurationHubAPI.md#CreateObjectMappings) | **Post** /configuration-hub/object-mappings/{sourceOrg}/bulk-create | Bulk creates object mappings
[**DeleteObjectMapping**](ConfigurationHubAPI.md#DeleteObjectMapping) | **Delete** /configuration-hub/object-mappings/{sourceOrg}/{objectMappingId} | Deletes an object mapping
[**DeleteUploadedBackup**](ConfigurationHubAPI.md#DeleteUploadedBackup) | **Delete** /configuration-hub/backups/uploads/{id} | Deletes an uploaded backup file
[**GetObjectMappings**](ConfigurationHubAPI.md#GetObjectMappings) | **Get** /configuration-hub/object-mappings/{sourceOrg} | Gets list of object mappings
[**GetUploadedBackup**](ConfigurationHubAPI.md#GetUploadedBackup) | **Get** /configuration-hub/backups/uploads/{id} | Get an uploaded backup&#39;s information
[**GetUploadedBackups**](ConfigurationHubAPI.md#GetUploadedBackups) | **Get** /configuration-hub/backups/uploads | Gets list of Uploaded backups
[**ImportUploadedBackup**](ConfigurationHubAPI.md#ImportUploadedBackup) | **Post** /configuration-hub/backups/uploads | Uploads a backup file
[**UpdateObjectMappings**](ConfigurationHubAPI.md#UpdateObjectMappings) | **Post** /configuration-hub/object-mappings/{sourceOrg}/bulk-patch | Bulk updates object mappings



## CreateObjectMapping

> ObjectMappingResponse CreateObjectMapping(ctx, sourceOrg).ObjectMappingRequest(objectMappingRequest).Execute()

Creates an object mapping



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
    sourceOrg := "source-org" // string | The name of the source org.
    objectMappingRequest := *openapiclient.NewObjectMappingRequest("IDENTITY", "$.name", "My Governance Group Name", "My New Governance Group Name") // ObjectMappingRequest | The object mapping request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.CreateObjectMapping(context.Background(), sourceOrg).ObjectMappingRequest(objectMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.CreateObjectMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectMapping`: ObjectMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.CreateObjectMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrg** | **string** | The name of the source org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectMappingRequest** | [**ObjectMappingRequest**](ObjectMappingRequest.md) | The object mapping request body. | 

### Return type

[**ObjectMappingResponse**](ObjectMappingResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateObjectMappings

> ObjectMappingBulkCreateResponse CreateObjectMappings(ctx, sourceOrg).ObjectMappingBulkCreateRequest(objectMappingBulkCreateRequest).Execute()

Bulk creates object mappings



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
    sourceOrg := "source-org" // string | The name of the source org.
    objectMappingBulkCreateRequest := *openapiclient.NewObjectMappingBulkCreateRequest([]openapiclient.ObjectMappingRequest{*openapiclient.NewObjectMappingRequest("IDENTITY", "$.name", "My Governance Group Name", "My New Governance Group Name")}) // ObjectMappingBulkCreateRequest | The bulk create object mapping request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.CreateObjectMappings(context.Background(), sourceOrg).ObjectMappingBulkCreateRequest(objectMappingBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.CreateObjectMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateObjectMappings`: ObjectMappingBulkCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.CreateObjectMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrg** | **string** | The name of the source org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateObjectMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectMappingBulkCreateRequest** | [**ObjectMappingBulkCreateRequest**](ObjectMappingBulkCreateRequest.md) | The bulk create object mapping request body. | 

### Return type

[**ObjectMappingBulkCreateResponse**](ObjectMappingBulkCreateResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteObjectMapping

> DeleteObjectMapping(ctx, sourceOrg, objectMappingId).Execute()

Deletes an object mapping



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
    sourceOrg := "source-org" // string | The name of the source org.
    objectMappingId := "3d6e0144-963f-4bd6-8d8d-d77b4e507ce4" // string | The id of the object mapping to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationHubAPI.DeleteObjectMapping(context.Background(), sourceOrg, objectMappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.DeleteObjectMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrg** | **string** | The name of the source org. | 
**objectMappingId** | **string** | The id of the object mapping to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectMappingRequest struct via the builder pattern


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


## DeleteUploadedBackup

> DeleteUploadedBackup(ctx, id).Execute()

Deletes an uploaded backup file



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
    id := "3d0fe04b-57df-4a46-a83b-8f04b0f9d10b" // string | The id of the uploaded backup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigurationHubAPI.DeleteUploadedBackup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.DeleteUploadedBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the uploaded backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUploadedBackupRequest struct via the builder pattern


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


## GetObjectMappings

> []ObjectMappingResponse GetObjectMappings(ctx, sourceOrg).Execute()

Gets list of object mappings



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
    sourceOrg := "source-org" // string | The name of the source org.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.GetObjectMappings(context.Background(), sourceOrg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.GetObjectMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectMappings`: []ObjectMappingResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.GetObjectMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrg** | **string** | The name of the source org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ObjectMappingResponse**](ObjectMappingResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadedBackup

> map[string]interface{} GetUploadedBackup(ctx, id).Execute()

Get an uploaded backup's information



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
    id := "3d0fe04b-57df-4a46-a83b-8f04b0f9d10b" // string | The id of the uploaded backup.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.GetUploadedBackup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.GetUploadedBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadedBackup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.GetUploadedBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the uploaded backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadedBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadedBackups

> []UploadsResponse GetUploadedBackups(ctx).Status(status).Execute()

Gets list of Uploaded backups



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
    status := "COMPLETE" // string | Filter listed uploaded backups by status of operation (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.GetUploadedBackups(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.GetUploadedBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUploadedBackups`: []UploadsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.GetUploadedBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadedBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter listed uploaded backups by status of operation | 

### Return type

[**[]UploadsResponse**](UploadsResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportUploadedBackup

> UploadsRequest ImportUploadedBackup(ctx).Data(data).Name(name).Execute()

Uploads a backup file



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
    data := os.NewFile(1234, "some_file") // *os.File | JSON file containing the objects to be imported.
    name := "name_example" // string | Name that will be assigned to the uploaded file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.ImportUploadedBackup(context.Background()).Data(data).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.ImportUploadedBackup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportUploadedBackup`: UploadsRequest
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.ImportUploadedBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportUploadedBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | ***os.File** | JSON file containing the objects to be imported. | 
 **name** | **string** | Name that will be assigned to the uploaded file. | 

### Return type

[**UploadsRequest**](UploadsRequest.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateObjectMappings

> ObjectMappingBulkPatchResponse UpdateObjectMappings(ctx, sourceOrg).ObjectMappingBulkPatchRequest(objectMappingBulkPatchRequest).Execute()

Bulk updates object mappings



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
    sourceOrg := "source-org" // string | The name of the source org.
    objectMappingBulkPatchRequest := *openapiclient.NewObjectMappingBulkPatchRequest(map[string][]openapiclient.JsonPatchOperation{"key": []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")}}) // ObjectMappingBulkPatchRequest | The object mapping request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationHubAPI.UpdateObjectMappings(context.Background(), sourceOrg).ObjectMappingBulkPatchRequest(objectMappingBulkPatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.UpdateObjectMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateObjectMappings`: ObjectMappingBulkPatchResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.UpdateObjectMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceOrg** | **string** | The name of the source org. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateObjectMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectMappingBulkPatchRequest** | [**ObjectMappingBulkPatchRequest**](ObjectMappingBulkPatchRequest.md) | The object mapping request body. | 

### Return type

[**ObjectMappingBulkPatchResponse**](ObjectMappingBulkPatchResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

