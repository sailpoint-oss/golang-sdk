# \NonEmployeeLifecycleManagementApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveNonEmployeeRequest**](NonEmployeeLifecycleManagementApi.md#ApproveNonEmployeeRequest) | **Post** /non-employee-approvals/{id}/approve | Approve a Non-Employee Request
[**CreateNonEmployeeRecord**](NonEmployeeLifecycleManagementApi.md#CreateNonEmployeeRecord) | **Post** /non-employee-records | Create Non-Employee Record
[**CreateNonEmployeeRequest**](NonEmployeeLifecycleManagementApi.md#CreateNonEmployeeRequest) | **Post** /non-employee-requests | Create Non-Employee Request
[**CreateNonEmployeeSource**](NonEmployeeLifecycleManagementApi.md#CreateNonEmployeeSource) | **Post** /non-employee-sources | Create Non-Employee Source
[**CreateNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementApi.md#CreateNonEmployeeSourceSchemaAttributes) | **Post** /non-employee-sources/{sourceId}/schema-attributes | Create Non-Employee Source Schema Attribute
[**DeleteNonEmployeeRecord**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeRecord) | **Delete** /non-employee-records/{id} | Delete Non-Employee Record
[**DeleteNonEmployeeRecordInBulk**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeRecordInBulk) | **Post** /non-employee-records/bulk-delete | Delete Multiple Non-Employee Records
[**DeleteNonEmployeeRequest**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeRequest) | **Delete** /non-employee-requests/{id} | Delete Non-Employee Request
[**DeleteNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeSchemaAttribute) | **Delete** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Delete Non-Employee Source&#39;s Schema Attribute
[**DeleteNonEmployeeSource**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeSource) | **Delete** /non-employee-sources/{sourceId} | Delete Non-Employee Source
[**DeleteNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementApi.md#DeleteNonEmployeeSourceSchemaAttributes) | **Delete** /non-employee-sources/{sourceId}/schema-attributes | Delete all custom schema attributes
[**ExportNonEmployeeRecords**](NonEmployeeLifecycleManagementApi.md#ExportNonEmployeeRecords) | **Get** /non-employee-sources/{id}/non-employees/download | Exports Non-Employee Records to CSV
[**ExportNonEmployeeSourceSchemaTemplate**](NonEmployeeLifecycleManagementApi.md#ExportNonEmployeeSourceSchemaTemplate) | **Get** /non-employee-sources/{id}/schema-attributes-template/download | Exports Source Schema Template
[**GetNonEmployeeApproval**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeApproval) | **Get** /non-employee-approvals/{id} | Get a non-employee approval item detail
[**GetNonEmployeeApprovalSummary**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeApprovalSummary) | **Get** /non-employee-approvals/summary/{requested-for} | Get Summary of Non-Employee Approval Requests
[**GetNonEmployeeBulkUploadStatus**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeBulkUploadStatus) | **Get** /non-employee-sources/{id}/non-employee-bulk-upload/status | Bulk upload status on source
[**GetNonEmployeeRecord**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeRecord) | **Get** /non-employee-records/{id} | Get a Non-Employee Record
[**GetNonEmployeeRequest**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeRequest) | **Get** /non-employee-requests/{id} | Get a Non-Employee Request
[**GetNonEmployeeRequestSummary**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeRequestSummary) | **Get** /non-employee-requests/summary/{requested-for} | Get Summary of Non-Employee Requests
[**GetNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeSchemaAttribute) | **Get** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Get Schema Attribute Non-Employee Source
[**GetNonEmployeeSource**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeSource) | **Get** /non-employee-sources/{sourceId} | Get a Non-Employee Source
[**GetNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementApi.md#GetNonEmployeeSourceSchemaAttributes) | **Get** /non-employee-sources/{sourceId}/schema-attributes | List Schema Attributes Non-Employee Source
[**ImportNonEmployeeRecordsInBulk**](NonEmployeeLifecycleManagementApi.md#ImportNonEmployeeRecordsInBulk) | **Post** /non-employee-sources/{id}/non-employee-bulk-upload | Imports, or Updates, Non-Employee Records
[**ListNonEmployeeApproval**](NonEmployeeLifecycleManagementApi.md#ListNonEmployeeApproval) | **Get** /non-employee-approvals | Get List of Non-Employee Approval Requests
[**ListNonEmployeeRecords**](NonEmployeeLifecycleManagementApi.md#ListNonEmployeeRecords) | **Get** /non-employee-records | List Non-Employee Records
[**ListNonEmployeeRequests**](NonEmployeeLifecycleManagementApi.md#ListNonEmployeeRequests) | **Get** /non-employee-requests | List Non-Employee Requests
[**ListNonEmployeeSources**](NonEmployeeLifecycleManagementApi.md#ListNonEmployeeSources) | **Get** /non-employee-sources | List Non-Employee Sources
[**PatchNonEmployeeRecord**](NonEmployeeLifecycleManagementApi.md#PatchNonEmployeeRecord) | **Patch** /non-employee-records/{id} | Patch Non-Employee Record
[**PatchNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#PatchNonEmployeeSchemaAttribute) | **Patch** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Patch Non-Employee Source&#39;s Schema Attribute
[**PatchNonEmployeeSource**](NonEmployeeLifecycleManagementApi.md#PatchNonEmployeeSource) | **Patch** /non-employee-sources/{sourceId} | Patch a Non-Employee Source
[**RejectNonEmployeeRequest**](NonEmployeeLifecycleManagementApi.md#RejectNonEmployeeRequest) | **Post** /non-employee-approvals/{id}/reject | Reject a Non-Employee Request
[**UpdateNonEmployeeRecord**](NonEmployeeLifecycleManagementApi.md#UpdateNonEmployeeRecord) | **Put** /non-employee-records/{id} | Update Non-Employee Record



## ApproveNonEmployeeRequest

> NonEmployeeApprovalItem ApproveNonEmployeeRequest(ctx, id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()

Approve a Non-Employee Request



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
    id := "id_example" // string | Non-Employee approval item id (UUID)
    nonEmployeeApprovalDecision := *openapiclient.NewNonEmployeeApprovalDecision() // NonEmployeeApprovalDecision | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ApproveNonEmployeeRequest(context.Background(), id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ApproveNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveNonEmployeeRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ApproveNonEmployeeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveNonEmployeeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonEmployeeApprovalDecision** | [**NonEmployeeApprovalDecision**](NonEmployeeApprovalDecision.md) |  | 

### Return type

[**NonEmployeeApprovalItem**](NonEmployeeApprovalItem.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonEmployeeRecord

> NonEmployeeRecord CreateNonEmployeeRecord(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

Create Non-Employee Record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    nonEmployeeRequestBody := *openapiclient.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-Employee record creation request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeRecord(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.CreateNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.CreateNonEmployeeRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonEmployeeRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonEmployeeRequestBody** | [**NonEmployeeRequestBody**](NonEmployeeRequestBody.md) | Non-Employee record creation request body. | 

### Return type

[**NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonEmployeeRequest

> NonEmployeeRequest CreateNonEmployeeRequest(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

Create Non-Employee Request



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    nonEmployeeRequestBody := *openapiclient.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-Employee creation request body

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeRequest(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.CreateNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeRequest`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.CreateNonEmployeeRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonEmployeeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonEmployeeRequestBody** | [**NonEmployeeRequestBody**](NonEmployeeRequestBody.md) | Non-Employee creation request body | 

### Return type

[**NonEmployeeRequest**](NonEmployeeRequest.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonEmployeeSource

> NonEmployeeSourceWithCloudExternalId CreateNonEmployeeSource(ctx).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()

Create Non-Employee Source



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
    nonEmployeeSourceRequestBody := *openapiclient.NewNonEmployeeSourceRequestBody("Retail", "Source description", *openapiclient.NewNonEmployeeIdnUserRequest("2c91808570313110017040b06f344ec9")) // NonEmployeeSourceRequestBody | Non-Employee source creation request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeSource(context.Background()).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.CreateNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeSource`: NonEmployeeSourceWithCloudExternalId
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.CreateNonEmployeeSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonEmployeeSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonEmployeeSourceRequestBody** | [**NonEmployeeSourceRequestBody**](NonEmployeeSourceRequestBody.md) | Non-Employee source creation request body. | 

### Return type

[**NonEmployeeSourceWithCloudExternalId**](NonEmployeeSourceWithCloudExternalId.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNonEmployeeSourceSchemaAttributes

> NonEmployeeSchemaAttribute CreateNonEmployeeSourceSchemaAttributes(ctx, sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()

Create Non-Employee Source Schema Attribute



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
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | The Source id
    nonEmployeeSchemaAttributeBody := *openapiclient.NewNonEmployeeSchemaAttributeBody("TEXT", "Account Name", "account.name") // NonEmployeeSchemaAttributeBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.CreateNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.CreateNonEmployeeSourceSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeSourceSchemaAttributes`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.CreateNonEmployeeSourceSchemaAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNonEmployeeSourceSchemaAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonEmployeeSchemaAttributeBody** | [**NonEmployeeSchemaAttributeBody**](NonEmployeeSchemaAttributeBody.md) |  | 

### Return type

[**NonEmployeeSchemaAttribute**](NonEmployeeSchemaAttribute.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeRecord

> DeleteNonEmployeeRecord(ctx, id).Execute()

Delete Non-Employee Record



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
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-Employee record id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecord(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeRecordInBulk

> DeleteNonEmployeeRecordInBulk(ctx).DeleteNonEmployeeRecordInBulkRequest(deleteNonEmployeeRecordInBulkRequest).Execute()

Delete Multiple Non-Employee Records



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
    deleteNonEmployeeRecordInBulkRequest := *openapiclient.NewDeleteNonEmployeeRecordInBulkRequest([]string{"Ids_example"}) // DeleteNonEmployeeRecordInBulkRequest | Non-Employee bulk delete request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecordInBulk(context.Background()).DeleteNonEmployeeRecordInBulkRequest(deleteNonEmployeeRecordInBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRecordInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeRecordInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNonEmployeeRecordInBulkRequest** | [**DeleteNonEmployeeRecordInBulkRequest**](DeleteNonEmployeeRecordInBulkRequest.md) | Non-Employee bulk delete request body. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeRequest

> DeleteNonEmployeeRequest(ctx, id).Execute()

Delete Non-Employee Request



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
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-Employee request id in the UUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee request id in the UUID format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeSchemaAttribute

> DeleteNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).Execute()

Delete Non-Employee Source's Schema Attribute



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
    attributeId := "2c91808b6ef1d43e016efba0ce470904" // string | The Schema Attribute Id (UUID)
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **string** | The Schema Attribute Id (UUID) | 
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeSource

> DeleteNonEmployeeSource(ctx, sourceId).Execute()

Delete Non-Employee Source



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
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | Source Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNonEmployeeSourceSchemaAttributes

> DeleteNonEmployeeSourceSchemaAttributes(ctx, sourceId).Execute()

Delete all custom schema attributes



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
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteNonEmployeeSourceSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeSourceSchemaAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportNonEmployeeRecords

> ExportNonEmployeeRecords(ctx, id).Execute()

Exports Non-Employee Records to CSV



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
    id := "2c918085842e69ae018432d22ccb212f" // string | Source Id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ExportNonEmployeeRecords(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ExportNonEmployeeRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportNonEmployeeRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportNonEmployeeSourceSchemaTemplate

> ExportNonEmployeeSourceSchemaTemplate(ctx, id).Execute()

Exports Source Schema Template



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
    id := "2c918085842e69ae018432d22ccb212f" // string | Source Id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ExportNonEmployeeSourceSchemaTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ExportNonEmployeeSourceSchemaTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportNonEmployeeSourceSchemaTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeApproval

> NonEmployeeApprovalItemDetail GetNonEmployeeApproval(ctx, id).IncludeDetail(includeDetail).Execute()

Get a non-employee approval item detail



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
    id := "ac10d20a-841e-1e7d-8184-32d2e22c0179" // string | Non-Employee approval item id (UUID)
    includeDetail := "include-detail=false" // string | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeApproval(context.Background(), id).IncludeDetail(includeDetail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeApproval`: NonEmployeeApprovalItemDetail
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeApproval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDetail** | **string** | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* | 

### Return type

[**NonEmployeeApprovalItemDetail**](NonEmployeeApprovalItemDetail.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeApprovalSummary

> NonEmployeeApprovalSummary GetNonEmployeeApprovalSummary(ctx, requestedFor).Execute()

Get Summary of Non-Employee Approval Requests



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
    requestedFor := "ac10d20a-841e-1e7d-8184-32d2e22c0179" // string | The identity (UUID) of the approver for whom for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeApprovalSummary(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeApprovalSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeApprovalSummary`: NonEmployeeApprovalSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeApprovalSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestedFor** | **string** | The identity (UUID) of the approver for whom for whom the summary is being retrieved. Use \&quot;me\&quot; instead to indicate the current user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeApprovalSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeApprovalSummary**](NonEmployeeApprovalSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeBulkUploadStatus

> NonEmployeeBulkUploadStatus GetNonEmployeeBulkUploadStatus(ctx, id).Execute()

Bulk upload status on source



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
    id := "2c918085842e69ae018432d22ccb212f" // string | Source ID (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeBulkUploadStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeBulkUploadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeBulkUploadStatus`: NonEmployeeBulkUploadStatus
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeBulkUploadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeBulkUploadStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeBulkUploadStatus**](NonEmployeeBulkUploadStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeRecord

> NonEmployeeRecord GetNonEmployeeRecord(ctx, id).Execute()

Get a Non-Employee Record



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
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-Employee record id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRecord(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeRequest

> NonEmployeeRequest GetNonEmployeeRequest(ctx, id).Execute()

Get a Non-Employee Request



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
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-Employee request id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRequest`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee request id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeRequest**](NonEmployeeRequest.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeRequestSummary

> NonEmployeeRequestSummary GetNonEmployeeRequestSummary(ctx, requestedFor).Execute()

Get Summary of Non-Employee Requests



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
    requestedFor := "ac10d20a-841e-1e7d-8184-32d2e22c0179" // string | The identity (UUID) of the non-employee account manager for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeRequestSummary(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeRequestSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRequestSummary`: NonEmployeeRequestSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeRequestSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestedFor** | **string** | The identity (UUID) of the non-employee account manager for whom the summary is being retrieved. Use \&quot;me\&quot; instead to indicate the current user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeRequestSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeRequestSummary**](NonEmployeeRequestSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeSchemaAttribute

> NonEmployeeSchemaAttribute GetNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).Execute()

Get Schema Attribute Non-Employee Source



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
    attributeId := "2c918085842e69ae018432d22ccb212f" // string | The Schema Attribute Id (UUID)
    sourceId := "2c918085842e69ae018432d22ccb212f" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **string** | The Schema Attribute Id (UUID) | 
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NonEmployeeSchemaAttribute**](NonEmployeeSchemaAttribute.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeSource

> NonEmployeeSource GetNonEmployeeSource(ctx, sourceId).Execute()

Get a Non-Employee Source



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
    sourceId := "2c91808b7c28b350017c2a2ec5790aa1" // string | Source Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSource`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NonEmployeeSource**](NonEmployeeSource.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNonEmployeeSourceSchemaAttributes

> []NonEmployeeSchemaAttribute GetNonEmployeeSourceSchemaAttributes(ctx, sourceId).Execute()

List Schema Attributes Non-Employee Source



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
    sourceId := "2c918085842e69ae018432d22ccb212f" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetNonEmployeeSourceSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSourceSchemaAttributes`: []NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetNonEmployeeSourceSchemaAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNonEmployeeSourceSchemaAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NonEmployeeSchemaAttribute**](NonEmployeeSchemaAttribute.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportNonEmployeeRecordsInBulk

> NonEmployeeBulkUploadJob ImportNonEmployeeRecordsInBulk(ctx, id).Data(data).Execute()

Imports, or Updates, Non-Employee Records



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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id (UUID)
    data := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ImportNonEmployeeRecordsInBulk(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ImportNonEmployeeRecordsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportNonEmployeeRecordsInBulk`: NonEmployeeBulkUploadJob
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ImportNonEmployeeRecordsInBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportNonEmployeeRecordsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | ***os.File** |  | 

### Return type

[**NonEmployeeBulkUploadJob**](NonEmployeeBulkUploadJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNonEmployeeApproval

> []NonEmployeeApprovalItem ListNonEmployeeApproval(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Get List of Non-Employee Approval Requests



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
    requestedFor := "ac10d20a-841e-1e7d-8184-32d2e22c0179" // string | The identity for whom the request was made. *me* indicates the current user. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "filters=approvalStatus eq 'PENDING'" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **approvalStatus**: *eq*  *Example:* approvalStatus eq \"PENDING\" (optional)
    sorters := "sorters=-created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeApproval(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ListNonEmployeeApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeApproval`: []NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ListNonEmployeeApproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeApprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **approvalStatus**: *eq*  *Example:* approvalStatus eq \&quot;PENDING\&quot; | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]NonEmployeeApprovalItem**](NonEmployeeApprovalItem.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNonEmployeeRecords

> []NonEmployeeRecord ListNonEmployeeRecords(ctx).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

List Non-Employee Records



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "accountName,sourceId" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **sourceId**: *eq*  *Example:* sourceId eq \"2c91808568c529c60168cca6f90c1313\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeRecords(context.Background()).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ListNonEmployeeRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeRecords`: []NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ListNonEmployeeRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **sourceId**: *eq*  *Example:* sourceId eq \&quot;2c91808568c529c60168cca6f90c1313\&quot; | 

### Return type

[**[]NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNonEmployeeRequests

> []NonEmployeeRequest ListNonEmployeeRequests(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

List Non-Employee Requests



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
    requestedFor := "me" // string | The identity for whom the request was made. *me* indicates the current user.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "approvalStatus,firstName" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **sourceId**: *eq*  *Example:* sourceId eq \"2c91808568c529c60168cca6f90c1313\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeRequests(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ListNonEmployeeRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeRequests`: []NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ListNonEmployeeRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **sourceId**: *eq*  *Example:* sourceId eq \&quot;2c91808568c529c60168cca6f90c1313\&quot; | 

### Return type

[**[]NonEmployeeRequest**](NonEmployeeRequest.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNonEmployeeSources

> []NonEmployeeSourceWithNECount ListNonEmployeeSources(ctx).RequestedFor(requestedFor).NonEmployeeCount(nonEmployeeCount).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()

List Non-Employee Sources



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
    requestedFor := "me" // string | The identity for whom the request was made. *me* indicates the current user.
    nonEmployeeCount := false // bool | The flag to determine whether return a non-employee count associate with source.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name,created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.ListNonEmployeeSources(context.Background()).RequestedFor(requestedFor).NonEmployeeCount(nonEmployeeCount).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.ListNonEmployeeSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeSources`: []NonEmployeeSourceWithNECount
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.ListNonEmployeeSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **nonEmployeeCount** | **bool** | The flag to determine whether return a non-employee count associate with source. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, created** | 

### Return type

[**[]NonEmployeeSourceWithNECount**](NonEmployeeSourceWithNECount.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNonEmployeeRecord

> NonEmployeeRecord PatchNonEmployeeRecord(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch Non-Employee Record



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
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-employee record id (UUID)
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeRecord(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.PatchNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.PatchNonEmployeeRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNonEmployeeRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of non-employee update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields. | 

### Return type

[**NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNonEmployeeSchemaAttribute

> NonEmployeeSchemaAttribute PatchNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

Patch Non-Employee Source's Schema Attribute



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
    attributeId := "2c91808b6ef1d43e016efba0ce470904" // string | The Schema Attribute Id (UUID)
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | The Source id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following properties are allowed for update ':' 'label', 'helpText', 'placeholder', 'required'.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.PatchNonEmployeeSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.PatchNonEmployeeSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **string** | The Schema Attribute Id (UUID) | 
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNonEmployeeSchemaAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following properties are allowed for update &#39;:&#39; &#39;label&#39;, &#39;helpText&#39;, &#39;placeholder&#39;, &#39;required&#39;. | 

### Return type

[**NonEmployeeSchemaAttribute**](NonEmployeeSchemaAttribute.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNonEmployeeSource

> NonEmployeeSource PatchNonEmployeeSource(ctx, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a Non-Employee Source



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
    sourceId := "2c91808b6ef1d43e016efba0ce470904" // string | Source Id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee source update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.PatchNonEmployeeSource(context.Background(), sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.PatchNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeSource`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.PatchNonEmployeeSource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNonEmployeeSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of non-employee source update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**NonEmployeeSource**](NonEmployeeSource.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectNonEmployeeRequest

> NonEmployeeApprovalItem RejectNonEmployeeRequest(ctx, id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()

Reject a Non-Employee Request



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
    id := "id_example" // string | Non-Employee approval item id (UUID)
    nonEmployeeRejectApprovalDecision := *openapiclient.NewNonEmployeeRejectApprovalDecision("Comment_example") // NonEmployeeRejectApprovalDecision | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.RejectNonEmployeeRequest(context.Background(), id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.RejectNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectNonEmployeeRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.RejectNonEmployeeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectNonEmployeeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonEmployeeRejectApprovalDecision** | [**NonEmployeeRejectApprovalDecision**](NonEmployeeRejectApprovalDecision.md) |  | 

### Return type

[**NonEmployeeApprovalItem**](NonEmployeeApprovalItem.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNonEmployeeRecord

> NonEmployeeRecord UpdateNonEmployeeRecord(ctx, id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

Update Non-Employee Record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "2c91808b6ef1d43e016efba0ce470904" // string | Non-employee record id (UUID)
    nonEmployeeRequestBody := *openapiclient.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-employee record creation request body. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.UpdateNonEmployeeRecord(context.Background(), id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.UpdateNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.UpdateNonEmployeeRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNonEmployeeRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonEmployeeRequestBody** | [**NonEmployeeRequestBody**](NonEmployeeRequestBody.md) | Non-employee record creation request body. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields. | 

### Return type

[**NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

