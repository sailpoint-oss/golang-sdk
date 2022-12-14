# \LaunchersApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLauncher**](LaunchersApi.md#GetLauncher) | **Get** /launchers/{launcherId} | Retrieves the details of the launcher.
[**LogLauncherClick**](LaunchersApi.md#LogLauncherClick) | **Post** /launchers/{launcherId}/click | Records a launcher click.
[**UpdateLauncher**](LaunchersApi.md#UpdateLauncher) | **Patch** /launchers/{launcherId} | Updates one or more attributes of a launcher.



## GetLauncher

> Launcher GetLauncher(ctx, launcherId).Execute()

Retrieves the details of the launcher.



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
    launcherId := "launcherId_example" // string | ID of a launcher.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LaunchersApi.GetLauncher(context.Background(), launcherId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LaunchersApi.GetLauncher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLauncher`: Launcher
    fmt.Fprintf(os.Stdout, "Response from `LaunchersApi.GetLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherId** | **string** | ID of a launcher. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Launcher**](Launcher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogLauncherClick

> LogLauncherClick(ctx, launcherId).Execute()

Records a launcher click.



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
    launcherId := "launcherId_example" // string | ID of a launcher.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LaunchersApi.LogLauncherClick(context.Background(), launcherId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LaunchersApi.LogLauncherClick``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherId** | **string** | ID of a launcher. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogLauncherClickRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLauncher

> Launcher UpdateLauncher(ctx, launcherId).LauncherEto(launcherEto).Execute()

Updates one or more attributes of a launcher.



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
    launcherId := "launcherId_example" // string | ID of a launcher.
    launcherEto := *openapiclient.NewLauncherEto() // LauncherEto | Launcher attributes to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LaunchersApi.UpdateLauncher(context.Background(), launcherId).LauncherEto(launcherEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LaunchersApi.UpdateLauncher``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLauncher`: Launcher
    fmt.Fprintf(os.Stdout, "Response from `LaunchersApi.UpdateLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherId** | **string** | ID of a launcher. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **launcherEto** | [**LauncherEto**](LauncherEto.md) | Launcher attributes to be updated. | 

### Return type

[**Launcher**](Launcher.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

