# \DefaultAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLauncher**](DefaultAPI.md#CreateLauncher) | **Post** /launchers | Create launcher
[**DeleteLauncher**](DefaultAPI.md#DeleteLauncher) | **Delete** /launchers/{launcherID} | Delete Launcher
[**PutLauncher**](DefaultAPI.md#PutLauncher) | **Put** /launchers/{launcherID} | Replace Launcher



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
	resp, r, err := apiClient.DefaultAPI.CreateLauncher(context.Background()).LauncherRequest(launcherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLauncher`: Launcher
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateLauncher`: %v\n", resp)
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
	r, err := apiClient.DefaultAPI.DeleteLauncher(context.Background(), launcherID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteLauncher``: %v\n", err)
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
	resp, r, err := apiClient.DefaultAPI.PutLauncher(context.Background(), launcherID).LauncherRequest(launcherRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PutLauncher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLauncher`: Launcher
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PutLauncher`: %v\n", resp)
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

