# \SPConfigApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpConfigExport**](SPConfigApi.md#SpConfigExport) | **Post** /sp-config/export | Initiates Configuration Objects Export Job.
[**SpConfigExportDownload**](SPConfigApi.md#SpConfigExportDownload) | **Get** /sp-config/export/{id}/download | Download Result of Export Job
[**SpConfigExportJobStatus**](SPConfigApi.md#SpConfigExportJobStatus) | **Get** /sp-config/export/{id} | Get Status of Export Job
[**SpConfigImport**](SPConfigApi.md#SpConfigImport) | **Post** /sp-config/import | Initiates Configuration Objects Import Job.
[**SpConfigImportDownload**](SPConfigApi.md#SpConfigImportDownload) | **Get** /sp-config/import/{id}/download | Download Result of Import Job
[**SpConfigImportJobStatus**](SPConfigApi.md#SpConfigImportJobStatus) | **Get** /sp-config/import/{id} | Get Status of Import Job
[**SpConfigObjects**](SPConfigApi.md#SpConfigObjects) | **Get** /sp-config/config-objects | Get Config Object details



## SpConfigExport

> SpConfigJob SpConfigExport(ctx).ExportPayload(exportPayload).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigExport(context.Background()).ExportPayload(exportPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigExport`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigExportRequest struct via the builder pattern


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


## SpConfigExportDownload

> SpConfigExportResults SpConfigExportDownload(ctx, id).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigExportDownload(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigExportDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigExportDownload`: SpConfigExportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigExportDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job for which the results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigExportDownloadRequest struct via the builder pattern


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


## SpConfigExportJobStatus

> SpConfigJob SpConfigExportJobStatus(ctx, id).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigExportJobStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigExportJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigExportJobStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigExportJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job for which status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigExportJobStatusRequest struct via the builder pattern


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


## SpConfigImport

> SpConfigJob SpConfigImport(ctx).Data(data).Preview(preview).Options(options).Execute()

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
    data := "data_example" // string | Name of JSON file containing the objects to be imported.
    preview := true // bool | This option is intended to give the user information about how an import operation would proceed, without having any affect on the target tenant. If true, no objects will be imported. Instead, the import process will pre-process the import file and attempt to resolve references within imported objects. The import result file will contain messages pertaining to how specific references were resolved, any errors associated with the preprocessing, and messages indicating which objects would be imported. (optional) (default to false)
    options := *openapiclient.NewImportOptions() // ImportOptions |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.SpConfigImport(context.Background()).Data(data).Preview(preview).Options(options).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigImport`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | **string** | Name of JSON file containing the objects to be imported. | 
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


## SpConfigImportDownload

> SpConfigImportResults SpConfigImportDownload(ctx, id).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigImportDownload(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigImportDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigImportDownload`: SpConfigImportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigImportDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job for which the results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigImportDownloadRequest struct via the builder pattern


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


## SpConfigImportJobStatus

> SpConfigJob SpConfigImportJobStatus(ctx, id).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigImportJobStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigImportJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigImportJobStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigImportJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job for which status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigImportJobStatusRequest struct via the builder pattern


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


## SpConfigObjects

> []SpConfigObject SpConfigObjects(ctx).Execute()

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
    resp, r, err := apiClient.SPConfigApi.SpConfigObjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.SpConfigObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SpConfigObjects`: []SpConfigObject
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.SpConfigObjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpConfigObjectsRequest struct via the builder pattern


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

