# \ConfigurationHubAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateObjectMapping**](ConfigurationHubAPI.md#CreateObjectMapping) | **Post** /configuration-hub/object-mappings/{sourceOrg} | Creates an object mapping
[**CreateObjectMappings**](ConfigurationHubAPI.md#CreateObjectMappings) | **Post** /configuration-hub/object-mappings/{sourceOrg}/bulk-create | Bulk creates object mappings
[**CreateUploadedConfiguration**](ConfigurationHubAPI.md#CreateUploadedConfiguration) | **Post** /configuration-hub/backups/uploads | Upload a Configuration
[**DeleteObjectMapping**](ConfigurationHubAPI.md#DeleteObjectMapping) | **Delete** /configuration-hub/object-mappings/{sourceOrg}/{objectMappingId} | Deletes an object mapping
[**DeleteUploadedConfiguration**](ConfigurationHubAPI.md#DeleteUploadedConfiguration) | **Delete** /configuration-hub/backups/uploads/{id} | Delete an Uploaded Configuration
[**GetObjectMappings**](ConfigurationHubAPI.md#GetObjectMappings) | **Get** /configuration-hub/object-mappings/{sourceOrg} | Gets list of object mappings
[**GetUploadedConfiguration**](ConfigurationHubAPI.md#GetUploadedConfiguration) | **Get** /configuration-hub/backups/uploads/{id} | Get an Uploaded Configuration
[**ListUploadedConfigurations**](ConfigurationHubAPI.md#ListUploadedConfigurations) | **Get** /configuration-hub/backups/uploads | List Uploaded Configurations
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadedConfiguration

> BackupResponse CreateUploadedConfiguration(ctx).Data(data).Name(name).Execute()

Upload a Configuration



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
	name := "name_example" // string | Name that will be assigned to the uploaded configuration file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationHubAPI.CreateUploadedConfiguration(context.Background()).Data(data).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.CreateUploadedConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadedConfiguration`: BackupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.CreateUploadedConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | ***os.File** | JSON file containing the objects to be imported. | 
 **name** | **string** | Name that will be assigned to the uploaded configuration file. | 

### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUploadedConfiguration

> DeleteUploadedConfiguration(ctx, id).Execute()

Delete an Uploaded Configuration



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
	id := "3d0fe04b-57df-4a46-a83b-8f04b0f9d10b" // string | The id of the uploaded configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationHubAPI.DeleteUploadedConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.DeleteUploadedConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the uploaded configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUploadedConfigurationRequest struct via the builder pattern


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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUploadedConfiguration

> BackupResponse GetUploadedConfiguration(ctx, id).Execute()

Get an Uploaded Configuration



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
	id := "3d0fe04b-57df-4a46-a83b-8f04b0f9d10b" // string | The id of the uploaded configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationHubAPI.GetUploadedConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.GetUploadedConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUploadedConfiguration`: BackupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.GetUploadedConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the uploaded configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUploadedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupResponse**](BackupResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUploadedConfigurations

> []BackupResponse ListUploadedConfigurations(ctx).Filters(filters).Execute()

List Uploaded Configurations



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
	filters := "status eq "COMPLETE"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationHubAPI.ListUploadedConfigurations(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationHubAPI.ListUploadedConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUploadedConfigurations`: []BackupResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationHubAPI.ListUploadedConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUploadedConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **status**: *eq* | 

### Return type

[**[]BackupResponse**](BackupResponse.md)

### Authorization

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: Not defined
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

[userAuth](../README.md#userAuth), [userAuth](../README.md#userAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

