# \CustomFormsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFormDefinition**](CustomFormsApi.md#CreateFormDefinition) | **Post** /form-definitions | Creates a form definition.
[**CreateFormDefinitionDynamicSchema**](CustomFormsApi.md#CreateFormDefinitionDynamicSchema) | **Post** /form-definitions/forms-action-dynamic-schema | Generate JSON Schema dynamically.
[**CreateFormInstance**](CustomFormsApi.md#CreateFormInstance) | **Post** /form-instances | Creates a form instance.
[**DeleteFormDefinition**](CustomFormsApi.md#DeleteFormDefinition) | **Delete** /form-definitions/{formDefinitionID} | Deletes a form definition.
[**ExportFormDefinitionsByTenant**](CustomFormsApi.md#ExportFormDefinitionsByTenant) | **Get** /form-definitions/export | List form definitions by tenant.
[**GetFormDefinitionByKey**](CustomFormsApi.md#GetFormDefinitionByKey) | **Get** /form-definitions/{formDefinitionID} | Return a form definition.
[**GetFormInstanceByKey**](CustomFormsApi.md#GetFormInstanceByKey) | **Get** /form-instances/{formInstanceID} | Returns a form instance.
[**ImportFormDefinitions**](CustomFormsApi.md#ImportFormDefinitions) | **Post** /form-definitions/import | Import form definitions from export.
[**PatchFormDefinition**](CustomFormsApi.md#PatchFormDefinition) | **Patch** /form-definitions/{formDefinitionID} | Patch a form definition.
[**PatchFormInstance**](CustomFormsApi.md#PatchFormInstance) | **Patch** /form-instances/{formInstanceID} | Patch a form instance.
[**SearchFormDefinitionsByTenant**](CustomFormsApi.md#SearchFormDefinitionsByTenant) | **Get** /form-definitions | Export form definitions by tenant.
[**SearchFormElementDataByElementID**](CustomFormsApi.md#SearchFormElementDataByElementID) | **Get** /form-instances/{formInstanceID}/data-source/{formElementID} | Retrieves dynamic data by element.
[**SearchFormInstancesByTenant**](CustomFormsApi.md#SearchFormInstancesByTenant) | **Get** /form-instances | List form instances by tenant.
[**SearchPreDefinedSelectOptions**](CustomFormsApi.md#SearchPreDefinedSelectOptions) | **Get** /predefined-select-options | List predefined select options.
[**ShowPreviewDataSource**](CustomFormsApi.md#ShowPreviewDataSource) | **Post** /form-definitions/{formDefinitionID}/data-source | Preview form definition data source.



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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateFormDefinitionRequest("My form", *openapiclient.NewFormOwner()) // CreateFormDefinitionRequest | Body is the request payload to create form definition request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.CreateFormDefinition(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.CreateFormDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFormDefinition`: FormDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.CreateFormDefinition`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewFormDefinitionDynamicSchemaRequest() // FormDefinitionDynamicSchemaRequest | Body is the request payload to create a form definition dynamic schema (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.CreateFormDefinitionDynamicSchema(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.CreateFormDefinitionDynamicSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFormDefinitionDynamicSchema`: FormDefinitionDynamicSchemaResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.CreateFormDefinitionDynamicSchema`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
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
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateFormInstanceRequest(*openapiclient.NewFormInstanceCreatedBy(), "2023-08-12T20:14:57.74486Z", "00000000-0000-0000-0000-000000000000", []openapiclient.FormInstanceRecipient{*openapiclient.NewFormInstanceRecipient()}) // CreateFormInstanceRequest | Body is the request payload to create a form instance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.CreateFormInstance(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.CreateFormInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFormInstance`: FormInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.CreateFormInstance`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.DeleteFormDefinition(context.Background(), formDefinitionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.DeleteFormDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFormDefinition`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.DeleteFormDefinition`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    offset := int64(0) // int64 | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. (optional) (default to 0)
    limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
    filters := "name sw "my form"" // string | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: <b>name</b>: <i>eq, gt, sw, in</i> <b>description</b>: <i>eq, gt, sw, in</i> <b>created</b>: <i>eq, gt, sw, in</i> <b>modified</b>: <i>eq, gt, sw, in</i> (optional)
    sorters := "name" // string | Sorters  Item will be sorted in the returned array if the sorters expression evaluates to true for that item. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/#sorting-results). Sorting is supported for the following fields: <b>name</b> <b>description</b> <b>created</b> <b>modified</b> (optional) (default to "name")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.ExportFormDefinitionsByTenant(context.Background()).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.ExportFormDefinitionsByTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFormDefinitionsByTenant`: []ExportFormDefinitionsByTenant200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.ExportFormDefinitionsByTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFormDefinitionsByTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. | [default to 0]
 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 250]
 **filters** | **string** | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: &lt;b&gt;name&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;description&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;created&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;modified&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; | 
 **sorters** | **string** | Sorters  Item will be sorted in the returned array if the sorters expression evaluates to true for that item. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/#sorting-results). Sorting is supported for the following fields: &lt;b&gt;name&lt;/b&gt; &lt;b&gt;description&lt;/b&gt; &lt;b&gt;created&lt;/b&gt; &lt;b&gt;modified&lt;/b&gt; | [default to &quot;name&quot;]

### Return type

[**[]ExportFormDefinitionsByTenant200ResponseInner**](ExportFormDefinitionsByTenant200ResponseInner.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "./openapi"
)

func main() {
    formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.GetFormDefinitionByKey(context.Background(), formDefinitionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.GetFormDefinitionByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormDefinitionByKey`: FormDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.GetFormDefinitionByKey`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.GetFormInstanceByKey(context.Background(), formInstanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.GetFormInstanceByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormInstanceByKey`: FormInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.GetFormInstanceByKey`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.ExportFormDefinitionsByTenant200ResponseInner{*openapiclient.NewExportFormDefinitionsByTenant200ResponseInner()} // []ExportFormDefinitionsByTenant200ResponseInner | Body is the request payload to import form definitions (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.ImportFormDefinitions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.ImportFormDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportFormDefinitions`: ImportFormDefinitions202Response
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.ImportFormDefinitions`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID
    body := []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}} // []map[string]map[string]interface{} | Body is the request payload to patch a form definition, check: https://jsonpatch.com (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.PatchFormDefinition(context.Background(), formDefinitionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.PatchFormDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFormDefinition`: FormDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.PatchFormDefinition`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID
    body := []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}} // []map[string]map[string]interface{} | Body is the request payload to patch a form instance, check: https://jsonpatch.com (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.PatchFormInstance(context.Background(), formInstanceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.PatchFormInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchFormInstance`: FormInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.PatchFormInstance`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    offset := int64(250) // int64 | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. (optional) (default to 0)
    limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
    filters := "name sw "my form"" // string | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: <b>name</b>: <i>eq, gt, sw, in</i> <b>description</b>: <i>eq, gt, sw, in</i> <b>created</b>: <i>eq, gt, sw, in</i> <b>modified</b>: <i>eq, gt, sw, in</i> (optional)
    sorters := "name" // string | Sorters  Item will be sorted in the returned array if the sorters expression evaluates to true for that item. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/#sorting-results). Sorting is supported for the following fields: <b>name</b> <b>description</b> <b>created</b> <b>modified</b> (optional) (default to "name")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.SearchFormDefinitionsByTenant(context.Background()).Offset(offset).Limit(limit).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.SearchFormDefinitionsByTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFormDefinitionsByTenant`: ListFormDefinitionsByTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.SearchFormDefinitionsByTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchFormDefinitionsByTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int64** | Offset  Integer specifying the offset of the first result from the beginning of the collection. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). The offset value is record-based, not page-based, and the index starts at 0. | [default to 0]
 **limit** | **int64** | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. | [default to 250]
 **filters** | **string** | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: &lt;b&gt;name&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;description&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;created&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;modified&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; | 
 **sorters** | **string** | Sorters  Item will be sorted in the returned array if the sorters expression evaluates to true for that item. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/#sorting-results). Sorting is supported for the following fields: &lt;b&gt;name&lt;/b&gt; &lt;b&gt;description&lt;/b&gt; &lt;b&gt;created&lt;/b&gt; &lt;b&gt;modified&lt;/b&gt; | [default to &quot;name&quot;]

### Return type

[**ListFormDefinitionsByTenantResponse**](ListFormDefinitionsByTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formInstanceID := "00000000-0000-0000-0000-000000000000" // string | Form instance ID
    formElementID := "1" // string | Form element ID
    limit := int64(250) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 250)
    filters := "label sw "my label"" // string | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: <b>value</b>: <i>eq, ne, in</i> <b>label</b>: <i>eq, ne, in</i> <b>subLabel</b>: <i>eq, ne, in</i> (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.SearchFormElementDataByElementID(context.Background(), formInstanceID, formElementID).Limit(limit).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.SearchFormElementDataByElementID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFormElementDataByElementID`: ListFormElementDataByElementIDResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.SearchFormElementDataByElementID`: %v\n", resp)
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
 **filters** | **string** | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: &lt;b&gt;value&lt;/b&gt;: &lt;i&gt;eq, ne, in&lt;/i&gt; &lt;b&gt;label&lt;/b&gt;: &lt;i&gt;eq, ne, in&lt;/i&gt; &lt;b&gt;subLabel&lt;/b&gt;: &lt;i&gt;eq, ne, in&lt;/i&gt; | 

### Return type

[**ListFormElementDataByElementIDResponse**](ListFormElementDataByElementIDResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.SearchFormInstancesByTenant(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.SearchFormInstancesByTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchFormInstancesByTenant`: ListFormInstancesByTenantResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.SearchFormInstancesByTenant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFormInstancesByTenantRequest struct via the builder pattern


### Return type

[**ListFormInstancesByTenantResponse**](ListFormInstancesByTenantResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.SearchPreDefinedSelectOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.SearchPreDefinedSelectOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchPreDefinedSelectOptions`: ListPredefinedSelectOptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.SearchPreDefinedSelectOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSearchPreDefinedSelectOptionsRequest struct via the builder pattern


### Return type

[**ListPredefinedSelectOptionsResponse**](ListPredefinedSelectOptionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    formDefinitionID := "00000000-0000-0000-0000-000000000000" // string | Form definition ID
    limit := int64(10) // int64 | Limit  Integer specifying the maximum number of records to return in a single API call. The standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#paginating-results). If it is not specified, a default limit is used. (optional) (default to 10)
    filters := "label sw "my label"" // string | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: <b>value</b>: <i>eq, gt, sw, in</i> <b>label</b>: <i>eq, gt, sw, in</i> <b>subLabel</b>: <i>eq, gt, sw, in</i> (optional)
    query := "support" // string | Query  String specifying to query against (optional)
    formElementPreviewRequest := *openapiclient.NewFormElementPreviewRequest() // FormElementPreviewRequest | Body is the request payload to create a form definition dynamic schema (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFormsApi.ShowPreviewDataSource(context.Background(), formDefinitionID).Limit(limit).Filters(filters).Query(query).FormElementPreviewRequest(formElementPreviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFormsApi.ShowPreviewDataSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowPreviewDataSource`: PreviewDataSourceResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomFormsApi.ShowPreviewDataSource`: %v\n", resp)
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
 **filters** | **string** | Filters  Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: &lt;b&gt;value&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;label&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; &lt;b&gt;subLabel&lt;/b&gt;: &lt;i&gt;eq, gt, sw, in&lt;/i&gt; | 
 **query** | **string** | Query  String specifying to query against | 
 **formElementPreviewRequest** | [**FormElementPreviewRequest**](FormElementPreviewRequest.md) | Body is the request payload to create a form definition dynamic schema | 

### Return type

[**PreviewDataSourceResponse**](PreviewDataSourceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

