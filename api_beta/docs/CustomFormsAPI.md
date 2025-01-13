# \CustomFormsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFormDefinition**](CustomFormsAPI.md#CreateFormDefinition) | **Post** /form-definitions | Creates a form definition.
[**CreateFormDefinitionDynamicSchema**](CustomFormsAPI.md#CreateFormDefinitionDynamicSchema) | **Post** /form-definitions/forms-action-dynamic-schema | Generate JSON Schema dynamically.
[**CreateFormDefinitionFileRequest**](CustomFormsAPI.md#CreateFormDefinitionFileRequest) | **Post** /form-definitions/{formDefinitionID}/upload | Upload new form definition file.
[**CreateFormInstance**](CustomFormsAPI.md#CreateFormInstance) | **Post** /form-instances | Creates a form instance.
[**DeleteFormDefinition**](CustomFormsAPI.md#DeleteFormDefinition) | **Delete** /form-definitions/{formDefinitionID} | Deletes a form definition.
[**ExportFormDefinitionsByTenant**](CustomFormsAPI.md#ExportFormDefinitionsByTenant) | **Get** /form-definitions/export | List form definitions by tenant.
[**GetFileFromS3**](CustomFormsAPI.md#GetFileFromS3) | **Get** /form-definitions/{formDefinitionID}/file/{fileID} | Download definition file by fileId.
[**GetFormDefinitionByKey**](CustomFormsAPI.md#GetFormDefinitionByKey) | **Get** /form-definitions/{formDefinitionID} | Return a form definition.
[**GetFormInstanceByKey**](CustomFormsAPI.md#GetFormInstanceByKey) | **Get** /form-instances/{formInstanceID} | Returns a form instance.
[**GetFormInstanceFile**](CustomFormsAPI.md#GetFormInstanceFile) | **Get** /form-instances/{formInstanceID}/file/{fileID} | Download instance file by fileId.
[**ImportFormDefinitions**](CustomFormsAPI.md#ImportFormDefinitions) | **Post** /form-definitions/import | Import form definitions from export.
[**PatchFormDefinition**](CustomFormsAPI.md#PatchFormDefinition) | **Patch** /form-definitions/{formDefinitionID} | Patch a form definition.
[**PatchFormInstance**](CustomFormsAPI.md#PatchFormInstance) | **Patch** /form-instances/{formInstanceID} | Patch a form instance.
[**SearchFormDefinitionsByTenant**](CustomFormsAPI.md#SearchFormDefinitionsByTenant) | **Get** /form-definitions | Export form definitions by tenant.
[**SearchFormElementDataByElementID**](CustomFormsAPI.md#SearchFormElementDataByElementID) | **Get** /form-instances/{formInstanceID}/data-source/{formElementID} | Retrieves dynamic data by element.
[**SearchFormInstancesByTenant**](CustomFormsAPI.md#SearchFormInstancesByTenant) | **Get** /form-instances | List form instances by tenant.
[**SearchPreDefinedSelectOptions**](CustomFormsAPI.md#SearchPreDefinedSelectOptions) | **Get** /form-definitions/predefined-select-options | List predefined select options.
[**ShowPreviewDataSource**](CustomFormsAPI.md#ShowPreviewDataSource) | **Post** /form-definitions/{formDefinitionID}/data-source | Preview form definition data source.



## CreateFormDefinition

> FormDefinitionResponse CreateFormDefinition(ctx).Body(body).Execute()

Creates a form definition.

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
	body := *openapiclient.NewCreateFormDefinitionRequest("My form", *openapiclient.NewFormOwner()) // CreateFormDefinitionRequest | Body is the request payload to create form definition request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.CreateFormDefinition(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.CreateFormDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormDefinition`: FormDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.CreateFormDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateFormDefinitionRequest**](CreateFormDefinitionRequest.md) | Body is the request payload to create form definition request | 

### Return type

[**FormDefinitionResponse**](FormDefinitionResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFormDefinitionDynamicSchema

> FormDefinitionDynamicSchemaResponse CreateFormDefinitionDynamicSchema(ctx).Body(body).Execute()

Generate JSON Schema dynamically.

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
	body := *openapiclient.NewFormDefinitionDynamicSchemaRequest() // FormDefinitionDynamicSchemaRequest | Body is the request payload to create a form definition dynamic schema (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.CreateFormDefinitionDynamicSchema(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.CreateFormDefinitionDynamicSchema``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormDefinitionDynamicSchema`: FormDefinitionDynamicSchemaResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.CreateFormDefinitionDynamicSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormDefinitionDynamicSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FormDefinitionDynamicSchemaRequest**](FormDefinitionDynamicSchemaRequest.md) | Body is the request payload to create a form definition dynamic schema | 

### Return type

[**FormDefinitionDynamicSchemaResponse**](FormDefinitionDynamicSchemaResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFormDefinitionFileRequest

> FormDefinitionFileUploadResponse CreateFormDefinitionFileRequest(ctx, formDefinitionID).File(file).Execute()

Upload new form definition file.



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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | FormDefinitionID  String specifying FormDefinitionID
	file := os.NewFile(1234, "some_file") // *os.File | File specifying the multipart

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.CreateFormDefinitionFileRequest(context.Background(), formDefinitionID).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.CreateFormDefinitionFileRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormDefinitionFileRequest`: FormDefinitionFileUploadResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.CreateFormDefinitionFileRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | FormDefinitionID  String specifying FormDefinitionID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormDefinitionFileRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | File specifying the multipart | 

### Return type

[**FormDefinitionFileUploadResponse**](FormDefinitionFileUploadResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFormInstance

> FormInstanceResponse CreateFormInstance(ctx).Body(body).Execute()

Creates a form instance.

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
	body := *openapiclient.NewCreateFormInstanceRequest(*openapiclient.NewFormInstanceCreatedBy(), "2023-08-12T20:14:57.74486Z", "00000000-0000-0000-0000-000000000000", []openapiclient.FormInstanceRecipient{*openapiclient.NewFormInstanceRecipient()}) // CreateFormInstanceRequest | Body is the request payload to create a form instance (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.CreateFormInstance(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.CreateFormInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFormInstance`: FormInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.CreateFormInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateFormInstanceRequest**](CreateFormInstanceRequest.md) | Body is the request payload to create a form instance | 

### Return type

[**FormInstanceResponse**](FormInstanceResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFormDefinition

> map[string]interface{} DeleteFormDefinition(ctx, formDefinitionID).Execute()

Deletes a form definition.



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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.DeleteFormDefinition(context.Background(), formDefinitionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.DeleteFormDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFormDefinition`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.DeleteFormDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | Form definition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormDefinitionRequest struct via the builder pattern


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


## ExportFormDefinitionsByTenant

> []ExportFormDefinitionsByTenant200ResponseInner ExportFormDefinitionsByTenant(ctx).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()

List form definitions by tenant.



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
	offset := int64(0) // int64 | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. (optional) (default to 0)
	limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
	filters := "name sw "my form"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, gt, sw, in*  **description**: *eq, gt, sw, in*  **created**: *eq, gt, sw, in*  **modified**: *eq, gt, sw, in* (optional)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, created, modified** (optional) (default to "name")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.ExportFormDefinitionsByTenant(context.Background()).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.ExportFormDefinitionsByTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportFormDefinitionsByTenant`: []ExportFormDefinitionsByTenant200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.ExportFormDefinitionsByTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFormDefinitionsByTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. | [default to 0]
 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 250]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, gt, sw, in*  **description**: *eq, gt, sw, in*  **created**: *eq, gt, sw, in*  **modified**: *eq, gt, sw, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, created, modified** | [default to &quot;name&quot;]

### Return type

[**[]ExportFormDefinitionsByTenant200ResponseInner**](ExportFormDefinitionsByTenant200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileFromS3

> *os.File GetFileFromS3(ctx, formDefinitionID, fileID).Execute()

Download definition file by fileId.

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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | FormDefinitionID  Form definition ID
	fileID := "00000031N0J7R2B57M8YG73J7M.png" // string | FileID  String specifying the hashed name of the uploaded file we are retrieving.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.GetFileFromS3(context.Background(), formDefinitionID, fileID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.GetFileFromS3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileFromS3`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.GetFileFromS3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | FormDefinitionID  Form definition ID | 
**fileID** | **string** | FileID  String specifying the hashed name of the uploaded file we are retrieving. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileFromS3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/jpeg, image/png, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormDefinitionByKey

> FormDefinitionResponse GetFormDefinitionByKey(ctx, formDefinitionID).Execute()

Return a form definition.



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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.GetFormDefinitionByKey(context.Background(), formDefinitionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.GetFormDefinitionByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormDefinitionByKey`: FormDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.GetFormDefinitionByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | Form definition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormDefinitionByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormDefinitionResponse**](FormDefinitionResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormInstanceByKey

> FormInstanceResponse GetFormInstanceByKey(ctx, formInstanceID).Execute()

Returns a form instance.



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
	formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.GetFormInstanceByKey(context.Background(), formInstanceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.GetFormInstanceByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormInstanceByKey`: FormInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.GetFormInstanceByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formInstanceID** | **string** | Form instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormInstanceByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormInstanceResponse**](FormInstanceResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFormInstanceFile

> *os.File GetFormInstanceFile(ctx, formInstanceID, fileID).Execute()

Download instance file by fileId.

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
	formInstanceID := "00000000-0000-0000-0000-000000000000" // string | FormInstanceID  Form instance ID
	fileID := "00000031N0J7R2B57M8YG73J7M.png" // string | FileID  String specifying the hashed name of the uploaded file we are retrieving.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.GetFormInstanceFile(context.Background(), formInstanceID, fileID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.GetFormInstanceFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFormInstanceFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.GetFormInstanceFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formInstanceID** | **string** | FormInstanceID  Form instance ID | 
**fileID** | **string** | FileID  String specifying the hashed name of the uploaded file we are retrieving. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormInstanceFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, image/jpeg, image/png, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFormDefinitions

> ImportFormDefinitions202Response ImportFormDefinitions(ctx).Body(body).Execute()

Import form definitions from export.

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
	body := []openapiclient.ExportFormDefinitionsByTenant200ResponseInner{*openapiclient.NewExportFormDefinitionsByTenant200ResponseInner()} // []ExportFormDefinitionsByTenant200ResponseInner | Body is the request payload to import form definitions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.ImportFormDefinitions(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.ImportFormDefinitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportFormDefinitions`: ImportFormDefinitions202Response
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.ImportFormDefinitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportFormDefinitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]ExportFormDefinitionsByTenant200ResponseInner**](ExportFormDefinitionsByTenant200ResponseInner.md) | Body is the request payload to import form definitions | 

### Return type

[**ImportFormDefinitions202Response**](ImportFormDefinitions202Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFormDefinition

> FormDefinitionResponse PatchFormDefinition(ctx, formDefinitionID).Body(body).Execute()

Patch a form definition.



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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID
	body := []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}} // []map[string]map[string]interface{} | Body is the request payload to patch a form definition, check: https://jsonpatch.com (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.PatchFormDefinition(context.Background(), formDefinitionID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.PatchFormDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFormDefinition`: FormDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.PatchFormDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | Form definition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFormDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]map[string]map[string]interface{}** | Body is the request payload to patch a form definition, check: https://jsonpatch.com | 

### Return type

[**FormDefinitionResponse**](FormDefinitionResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFormInstance

> FormInstanceResponse PatchFormInstance(ctx, formInstanceID).Body(body).Execute()

Patch a form instance.



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
	formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID
	body := []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}} // []map[string]map[string]interface{} | Body is the request payload to patch a form instance, check: https://jsonpatch.com (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.PatchFormInstance(context.Background(), formInstanceID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.PatchFormInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchFormInstance`: FormInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.PatchFormInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formInstanceID** | **string** | Form instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFormInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]map[string]map[string]interface{}** | Body is the request payload to patch a form instance, check: https://jsonpatch.com | 

### Return type

[**FormInstanceResponse**](FormInstanceResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFormDefinitionsByTenant

> ListFormDefinitionsByTenantResponse SearchFormDefinitionsByTenant(ctx).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()

Export form definitions by tenant.



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
	offset := int64(250) // int64 | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. (optional) (default to 0)
	limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
	filters := "name sw "my form"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, gt, sw, in*  **description**: *eq, gt, sw, in*  **created**: *eq, gt, sw, in*  **modified**: *eq, gt, sw, in* (optional)
	sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, created, modified** (optional) (default to "name")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.SearchFormDefinitionsByTenant(context.Background()).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.SearchFormDefinitionsByTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFormDefinitionsByTenant`: ListFormDefinitionsByTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.SearchFormDefinitionsByTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFormDefinitionsByTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. | [default to 0]
 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 250]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, gt, sw, in*  **description**: *eq, gt, sw, in*  **created**: *eq, gt, sw, in*  **modified**: *eq, gt, sw, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, description, created, modified** | [default to &quot;name&quot;]

### Return type

[**ListFormDefinitionsByTenantResponse**](ListFormDefinitionsByTenantResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFormElementDataByElementID

> ListFormElementDataByElementIDResponse SearchFormElementDataByElementID(ctx, formInstanceID, formElementID).Limit(limit).Filters(filters).Execute()

Retrieves dynamic data by element.



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
	formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID
	formElementID := "1" // string | Form element ID
	limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
	filters := "label sw "my label"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **value**: *eq, ne, in*  **label**: *eq, ne, in*  **subLabel**: *eq, ne, in* (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.SearchFormElementDataByElementID(context.Background(), formInstanceID, formElementID).Limit(limit).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.SearchFormElementDataByElementID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFormElementDataByElementID`: ListFormElementDataByElementIDResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.SearchFormElementDataByElementID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formInstanceID** | **string** | Form instance ID | 
**formElementID** | **string** | Form element ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFormElementDataByElementIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 250]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **value**: *eq, ne, in*  **label**: *eq, ne, in*  **subLabel**: *eq, ne, in* | 

### Return type

[**ListFormElementDataByElementIDResponse**](ListFormElementDataByElementIDResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFormInstancesByTenant

> ListFormInstancesByTenantResponse SearchFormInstancesByTenant(ctx).Execute()

List form instances by tenant.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.SearchFormInstancesByTenant(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.SearchFormInstancesByTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFormInstancesByTenant`: ListFormInstancesByTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.SearchFormInstancesByTenant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFormInstancesByTenantRequest struct via the builder pattern


### Return type

[**ListFormInstancesByTenantResponse**](ListFormInstancesByTenantResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPreDefinedSelectOptions

> ListPredefinedSelectOptionsResponse SearchPreDefinedSelectOptions(ctx).Execute()

List predefined select options.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.SearchPreDefinedSelectOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.SearchPreDefinedSelectOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPreDefinedSelectOptions`: ListPredefinedSelectOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.SearchPreDefinedSelectOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchPreDefinedSelectOptionsRequest struct via the builder pattern


### Return type

[**ListPredefinedSelectOptionsResponse**](ListPredefinedSelectOptionsResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPreviewDataSource

> PreviewDataSourceResponse ShowPreviewDataSource(ctx, formDefinitionID).Limit(limit).Filters(filters).Query(query).FormElementPreviewRequest(formElementPreviewRequest).Execute()

Preview form definition data source.

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
	formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID
	limit := int64(10) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 10)
	filters := "label sw "my label"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **value**: *eq, gt, sw, in*  **label**: *eq, gt, sw, in*  **subLabel**: *eq, gt, sw, in* (optional)
	query := "support" // string | Query  String specifying to query against (optional)
	formElementPreviewRequest := *openapiclient.NewFormElementPreviewRequest() // FormElementPreviewRequest | Body is the request payload to create a form definition dynamic schema (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomFormsAPI.ShowPreviewDataSource(context.Background(), formDefinitionID).Limit(limit).Filters(filters).Query(query).FormElementPreviewRequest(formElementPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsAPI.ShowPreviewDataSource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPreviewDataSource`: PreviewDataSourceResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomFormsAPI.ShowPreviewDataSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**formDefinitionID** | **string** | Form definition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPreviewDataSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 10]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **value**: *eq, gt, sw, in*  **label**: *eq, gt, sw, in*  **subLabel**: *eq, gt, sw, in* | 
 **query** | **string** | Query  String specifying to query against | 
 **formElementPreviewRequest** | [**FormElementPreviewRequest**](FormElementPreviewRequest.md) | Body is the request payload to create a form definition dynamic schema | 

### Return type

[**PreviewDataSourceResponse**](PreviewDataSourceResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

