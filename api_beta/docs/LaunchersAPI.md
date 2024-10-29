# \LaunchersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLauncher**](LaunchersAPI.md#CreateLauncher) | **Post** /launchers | Create launcher
[**DeleteLauncher**](LaunchersAPI.md#DeleteLauncher) | **Delete** /launchers/{launcherID} | Delete Launcher
[**GetLauncher**](LaunchersAPI.md#GetLauncher) | **Get** /launchers/{launcherID} | Get Launcher by ID
[**GetLaunchers**](LaunchersAPI.md#GetLaunchers) | **Get** /launchers | List all Launchers for tenant
[**PutLauncher**](LaunchersAPI.md#PutLauncher) | **Put** /launchers/{launcherID} | Replace Launcher
[**StartLauncher**](LaunchersAPI.md#StartLauncher) | **Post** /beta/launchers/{launcherID}/launch | Launch a Launcher



## CreateLauncher

> Launcher CreateLauncher(ctx).LauncherRequest(launcherRequest).Execute()

Create launcher



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
	launcherRequest := *openapiclient.NewLauncherRequest("Group Create", "Create a new Active Directory Group", "INTERACTIVE_PROCESS", false, "{"workflowId" : "6b42d9be-61b6-46af-827e-ea29ba8aa3d9"}") // LauncherRequest | Payload to create a Launcher

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchersAPI.CreateLauncher(context.Background()).LauncherRequest(launcherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.CreateLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLauncher`: Launcher
	fmt.Fprintf(os.Stdout, "Response from `LaunchersAPI.CreateLauncher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **launcherRequest** | [**LauncherRequest**](LauncherRequest.md) | Payload to create a Launcher | 

### Return type

[**Launcher**](Launcher.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLauncher

> DeleteLauncher(ctx, launcherID).Execute()

Delete Launcher



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
	launcherID := "e3012408-8b61-4564-ad41-c5ec131c325b" // string | ID of the Launcher to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LaunchersAPI.DeleteLauncher(context.Background(), launcherID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.DeleteLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherID** | **string** | ID of the Launcher to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLauncher

> Launcher GetLauncher(ctx, launcherID).Execute()

Get Launcher by ID



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
	launcherID := "e3012408-8b61-4564-ad41-c5ec131c325b" // string | ID of the Launcher to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchersAPI.GetLauncher(context.Background(), launcherID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.GetLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLauncher`: Launcher
	fmt.Fprintf(os.Stdout, "Response from `LaunchersAPI.GetLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherID** | **string** | ID of the Launcher to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Launcher**](Launcher.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLaunchers

> GetLaunchers200Response GetLaunchers(ctx).Filters(filters).Next(next).Limit(limit).Execute()

List all Launchers for tenant



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
	filters := "disabled eq "true"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **description**: *sw*  **disabled**: *eq*  **name**: *sw* (optional)
	next := "eyJuZXh0IjoxMjN9Cg==" // string | Pagination marker (optional)
	limit := int32(42) // int32 | Number of Launchers to return (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchersAPI.GetLaunchers(context.Background()).Filters(filters).Next(next).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.GetLaunchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLaunchers`: GetLaunchers200Response
	fmt.Fprintf(os.Stdout, "Response from `LaunchersAPI.GetLaunchers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLaunchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **description**: *sw*  **disabled**: *eq*  **name**: *sw* | 
 **next** | **string** | Pagination marker | 
 **limit** | **int32** | Number of Launchers to return | [default to 10]

### Return type

[**GetLaunchers200Response**](GetLaunchers200Response.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLauncher

> Launcher PutLauncher(ctx, launcherID).LauncherRequest(launcherRequest).Execute()

Replace Launcher



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
	launcherID := "e3012408-8b61-4564-ad41-c5ec131c325b" // string | ID of the Launcher to be replaced
	launcherRequest := *openapiclient.NewLauncherRequest("Group Create", "Create a new Active Directory Group", "INTERACTIVE_PROCESS", false, "{"workflowId" : "6b42d9be-61b6-46af-827e-ea29ba8aa3d9"}") // LauncherRequest | Payload to replace Launcher

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchersAPI.PutLauncher(context.Background(), launcherID).LauncherRequest(launcherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.PutLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLauncher`: Launcher
	fmt.Fprintf(os.Stdout, "Response from `LaunchersAPI.PutLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherID** | **string** | ID of the Launcher to be replaced | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **launcherRequest** | [**LauncherRequest**](LauncherRequest.md) | Payload to replace Launcher | 

### Return type

[**Launcher**](Launcher.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartLauncher

> StartLauncher200Response StartLauncher(ctx, launcherID).Execute()

Launch a Launcher



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
	launcherID := "e3012408-8b61-4564-ad41-c5ec131c325b" // string | ID of the Launcher to be launched

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LaunchersAPI.StartLauncher(context.Background(), launcherID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LaunchersAPI.StartLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartLauncher`: StartLauncher200Response
	fmt.Fprintf(os.Stdout, "Response from `LaunchersAPI.StartLauncher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**launcherID** | **string** | ID of the Launcher to be launched | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartLauncherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StartLauncher200Response**](StartLauncher200Response.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

