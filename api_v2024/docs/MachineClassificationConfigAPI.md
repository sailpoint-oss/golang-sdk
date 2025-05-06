# \MachineClassificationConfigAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMachineClassificationConfig**](MachineClassificationConfigAPI.md#DeleteMachineClassificationConfig) | **Delete** /sources/{sourceId}/machine-classification-config | Delete Source&#39;s Classification Config
[**GetMachineClassificationConfig**](MachineClassificationConfigAPI.md#GetMachineClassificationConfig) | **Get** /sources/{sourceId}/machine-classification-config | Machine Classification Config for Source
[**SetMachineClassificationConfig**](MachineClassificationConfigAPI.md#SetMachineClassificationConfig) | **Put** /sources/{sourceId}/machine-classification-config | Update Source&#39;s Classification Config



## DeleteMachineClassificationConfig

> DeleteMachineClassificationConfig(ctx, id).Execute()

Delete Source's Classification Config



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MachineClassificationConfigAPI.DeleteMachineClassificationConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineClassificationConfigAPI.DeleteMachineClassificationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineClassificationConfigRequest struct via the builder pattern


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


## GetMachineClassificationConfig

> MachineClassificationConfig GetMachineClassificationConfig(ctx, id).Execute()

Machine Classification Config for Source



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineClassificationConfigAPI.GetMachineClassificationConfig(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineClassificationConfigAPI.GetMachineClassificationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMachineClassificationConfig`: MachineClassificationConfig
	fmt.Fprintf(os.Stdout, "Response from `MachineClassificationConfigAPI.GetMachineClassificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineClassificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MachineClassificationConfig**](MachineClassificationConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMachineClassificationConfig

> MachineClassificationConfig SetMachineClassificationConfig(ctx, id).MachineClassificationConfig(machineClassificationConfig).Execute()

Update Source's Classification Config



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
	id := "ef38f94347e94562b5bb8424a56397d8" // string | Source ID.
	machineClassificationConfig := *openapiclient.NewMachineClassificationConfig() // MachineClassificationConfig | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MachineClassificationConfigAPI.SetMachineClassificationConfig(context.Background(), id).MachineClassificationConfig(machineClassificationConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MachineClassificationConfigAPI.SetMachineClassificationConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMachineClassificationConfig`: MachineClassificationConfig
	fmt.Fprintf(os.Stdout, "Response from `MachineClassificationConfigAPI.SetMachineClassificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMachineClassificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **machineClassificationConfig** | [**MachineClassificationConfig**](MachineClassificationConfig.md) |  | 

### Return type

[**MachineClassificationConfig**](MachineClassificationConfig.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

