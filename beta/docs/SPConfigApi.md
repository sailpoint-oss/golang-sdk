# \SPConfigApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSpConfig**](SPConfigApi.md#ExportSpConfig) | **Post** /sp-config/export | Initiates Configuration Objects Export Job.
[**ExportSpConfigDownload**](SPConfigApi.md#ExportSpConfigDownload) | **Get** /sp-config/export/{id}/download | Download Result of Export Job
[**ExportSpConfigJobStatus**](SPConfigApi.md#ExportSpConfigJobStatus) | **Get** /sp-config/export/{id} | Get Status of Export Job
[**ImportSpConfig**](SPConfigApi.md#ImportSpConfig) | **Post** /sp-config/import | Initiates Configuration Objects Import Job.
[**ImportSpConfigDownload**](SPConfigApi.md#ImportSpConfigDownload) | **Get** /sp-config/import/{id}/download | Download Result of Import Job
[**ImportSpConfigJobStatus**](SPConfigApi.md#ImportSpConfigJobStatus) | **Get** /sp-config/import/{id} | Get Status of Import Job
[**ListSpConfigObjects**](SPConfigApi.md#ListSpConfigObjects) | **Get** /sp-config/config-objects | Get Config Object details



## ExportSpConfig

> SpConfigJob ExportSpConfig(ctx).ExportPayload(exportPayload).Execute()

Initiates Configuration Objects Export Job.



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
    exportPayload := *openapiclient.NewExportPayload() // ExportPayload | Export options control what will be included in the export.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ExportSpConfig(context.Background()).ExportPayload(exportPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ExportSpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSpConfig`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ExportSpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportPayload** | [**ExportPayload**](ExportPayload.md) | Export options control what will be included in the export. | 

### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSpConfigDownload

> SpConfigExportResults ExportSpConfigDownload(ctx, id).Execute()

Download Result of Export Job



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the export job for which the results will be downloaded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ExportSpConfigDownload(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ExportSpConfigDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSpConfigDownload`: SpConfigExportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ExportSpConfigDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job for which the results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSpConfigDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigExportResults**](SpConfigExportResults.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSpConfigJobStatus

> SpConfigJob ExportSpConfigJobStatus(ctx, id).Execute()

Get Status of Export Job



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the export job for which status will be returned.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ExportSpConfigJobStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ExportSpConfigJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSpConfigJobStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ExportSpConfigJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job for which status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSpConfigJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSpConfig

> SpConfigJob ImportSpConfig(ctx).Data(data).Preview(preview).Options(options).Execute()

Initiates Configuration Objects Import Job.



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
    data := os.NewFile(1234, "some_file") // *os.File | JSON file containing the objects to be imported.
    preview := true // bool | This option is intended to give the user information about how an import operation would proceed, without having any affect on the target tenant. If true, no objects will be imported. Instead, the import process will pre-process the import file and attempt to resolve references within imported objects. The import result file will contain messages pertaining to how specific references were resolved, any errors associated with the preprocessing, and messages indicating which objects would be imported. (optional) (default to false)
    options := *openapiclient.NewImportOptions() // ImportOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ImportSpConfig(context.Background()).Data(data).Preview(preview).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ImportSpConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSpConfig`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ImportSpConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSpConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | ***os.File** | JSON file containing the objects to be imported. | 
 **preview** | **bool** | This option is intended to give the user information about how an import operation would proceed, without having any affect on the target tenant. If true, no objects will be imported. Instead, the import process will pre-process the import file and attempt to resolve references within imported objects. The import result file will contain messages pertaining to how specific references were resolved, any errors associated with the preprocessing, and messages indicating which objects would be imported. | [default to false]
 **options** | [**ImportOptions**](ImportOptions.md) |  | 

### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSpConfigDownload

> SpConfigImportResults ImportSpConfigDownload(ctx, id).Execute()

Download Result of Import Job



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the import job for which the results will be downloaded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ImportSpConfigDownload(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ImportSpConfigDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSpConfigDownload`: SpConfigImportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ImportSpConfigDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job for which the results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSpConfigDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigImportResults**](SpConfigImportResults.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSpConfigJobStatus

> SpConfigJob ImportSpConfigJobStatus(ctx, id).Execute()

Get Status of Import Job



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the import job for which status will be returned.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ImportSpConfigJobStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ImportSpConfigJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportSpConfigJobStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ImportSpConfigJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job for which status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportSpConfigJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpConfigObjects

> []SpConfigObject ListSpConfigObjects(ctx).Execute()

Get Config Object details



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.ListSpConfigObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.ListSpConfigObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSpConfigObjects`: []SpConfigObject
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.ListSpConfigObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSpConfigObjectsRequest struct via the builder pattern


### Return type

[**[]SpConfigObject**](SpConfigObject.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

