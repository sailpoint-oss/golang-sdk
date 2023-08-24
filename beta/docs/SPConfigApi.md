# \SPConfigApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSpConfig**](SPConfigApi.md#ExportSpConfig) | **Post** /sp-config/export | Initiates configuration objects export job
[**GetSpConfigExport**](SPConfigApi.md#GetSpConfigExport) | **Get** /sp-config/export/{id}/download | Download export job result.
[**GetSpConfigExportStatus**](SPConfigApi.md#GetSpConfigExportStatus) | **Get** /sp-config/export/{id} | Get export job status
[**GetSpConfigImport**](SPConfigApi.md#GetSpConfigImport) | **Get** /sp-config/import/{id}/download | Download import job result
[**GetSpConfigImportStatus**](SPConfigApi.md#GetSpConfigImportStatus) | **Get** /sp-config/import/{id} | Get import job status
[**ImportSpConfig**](SPConfigApi.md#ImportSpConfig) | **Post** /sp-config/import | Initiates configuration objects import job
[**ListSpConfigObjects**](SPConfigApi.md#ListSpConfigObjects) | **Get** /sp-config/config-objects | Get config object details



## ExportSpConfig

> SpConfigJob ExportSpConfig(ctx).ExportPayload(exportPayload).Execute()

Initiates configuration objects export job



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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConfigExport

> SpConfigExportResults GetSpConfigExport(ctx, id).Execute()

Download export job result.



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the export job whose results will be downloaded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.GetSpConfigExport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.GetSpConfigExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConfigExport`: SpConfigExportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.GetSpConfigExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job whose results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConfigExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigExportResults**](SpConfigExportResults.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConfigExportStatus

> SpConfigJob GetSpConfigExportStatus(ctx, id).Execute()

Get export job status



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the export job whose status will be returned.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.GetSpConfigExportStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.GetSpConfigExportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConfigExportStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.GetSpConfigExportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the export job whose status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConfigExportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConfigImport

> SpConfigImportResults GetSpConfigImport(ctx, id).Execute()

Download import job result



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the import job whose results will be downloaded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.GetSpConfigImport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.GetSpConfigImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConfigImport`: SpConfigImportResults
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.GetSpConfigImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job whose results will be downloaded. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConfigImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigImportResults**](SpConfigImportResults.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpConfigImportStatus

> SpConfigJob GetSpConfigImportStatus(ctx, id).Execute()

Get import job status



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the import job whose status will be returned.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SPConfigApi.GetSpConfigImportStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SPConfigApi.GetSpConfigImportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSpConfigImportStatus`: SpConfigJob
    fmt.Fprintf(os.Stdout, "Response from `SPConfigApi.GetSpConfigImportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the import job whose status will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpConfigImportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSpConfig

> SpConfigJob ImportSpConfig(ctx).Data(data).Preview(preview).Options(options).Execute()

Initiates configuration objects import job



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
    preview := true // bool | This option is intended to give the user information about how an import operation would proceed, without having any effect on the target tenant. If this parameter is \"true\", no objects will be imported. Instead, the import process will pre-process the import file and attempt to resolve references within imported objects. The import result file will contain messages pertaining to how specific references were resolved, any errors associated with the preprocessing, and messages indicating which objects would be imported. (optional) (default to false)
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
 **preview** | **bool** | This option is intended to give the user information about how an import operation would proceed, without having any effect on the target tenant. If this parameter is \&quot;true\&quot;, no objects will be imported. Instead, the import process will pre-process the import file and attempt to resolve references within imported objects. The import result file will contain messages pertaining to how specific references were resolved, any errors associated with the preprocessing, and messages indicating which objects would be imported. | [default to false]
 **options** | [**ImportOptions**](ImportOptions.md) |  | 

### Return type

[**SpConfigJob**](SpConfigJob.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSpConfigObjects

> []SpConfigObject ListSpConfigObjects(ctx).Execute()

Get config object details



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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

