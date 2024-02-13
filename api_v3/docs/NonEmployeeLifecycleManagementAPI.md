# \NonEmployeeLifecycleManagementAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveNonEmployeeRequest**](NonEmployeeLifecycleManagementAPI.md#ApproveNonEmployeeRequest) | **Post** /non-employee-approvals/{id}/approve | Approve a Non-Employee Request
[**CreateNonEmployeeRecord**](NonEmployeeLifecycleManagementAPI.md#CreateNonEmployeeRecord) | **Post** /non-employee-records | Create Non-Employee Record
[**CreateNonEmployeeRequest**](NonEmployeeLifecycleManagementAPI.md#CreateNonEmployeeRequest) | **Post** /non-employee-requests | Create Non-Employee Request
[**CreateNonEmployeeSource**](NonEmployeeLifecycleManagementAPI.md#CreateNonEmployeeSource) | **Post** /non-employee-sources | Create Non-Employee Source
[**CreateNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementAPI.md#CreateNonEmployeeSourceSchemaAttributes) | **Post** /non-employee-sources/{sourceId}/schema-attributes | Create a new Schema Attribute for Non-Employee Source
[**DeleteNonEmployeeRecord**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeRecord) | **Delete** /non-employee-records/{id} | Delete Non-Employee Record
[**DeleteNonEmployeeRecordsInBulk**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeRecordsInBulk) | **Post** /non-employee-records/bulk-delete | Delete Multiple Non-Employee Records
[**DeleteNonEmployeeRequest**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeRequest) | **Delete** /non-employee-requests/{id} | Delete Non-Employee Request
[**DeleteNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeSchemaAttribute) | **Delete** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Delete a Schema Attribute for Non-Employee Source
[**DeleteNonEmployeeSource**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeSource) | **Delete** /non-employee-sources/{sourceId} | Delete Non-Employee Source
[**DeleteNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementAPI.md#DeleteNonEmployeeSourceSchemaAttributes) | **Delete** /non-employee-sources/{sourceId}/schema-attributes | Delete all custom schema attributes for Non-Employee Source
[**ExportNonEmployeeRecords**](NonEmployeeLifecycleManagementAPI.md#ExportNonEmployeeRecords) | **Get** /non-employee-sources/{id}/non-employees/download | Exports Non-Employee Records to CSV
[**ExportNonEmployeeSourceSchemaTemplate**](NonEmployeeLifecycleManagementAPI.md#ExportNonEmployeeSourceSchemaTemplate) | **Get** /non-employee-sources/{id}/schema-attributes-template/download | Exports Source Schema Template
[**GetNonEmployeeApproval**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeApproval) | **Get** /non-employee-approvals/{id} | Get a non-employee approval item detail
[**GetNonEmployeeApprovalSummary**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeApprovalSummary) | **Get** /non-employee-approvals/summary/{requested-for} | Get Summary of Non-Employee Approval Requests
[**GetNonEmployeeBulkUploadStatus**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeBulkUploadStatus) | **Get** /non-employee-sources/{id}/non-employee-bulk-upload/status | Obtain the status of bulk upload on the source
[**GetNonEmployeeRecord**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeRecord) | **Get** /non-employee-records/{id} | Get a Non-Employee Record
[**GetNonEmployeeRequest**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeRequest) | **Get** /non-employee-requests/{id} | Get a Non-Employee Request
[**GetNonEmployeeRequestSummary**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeRequestSummary) | **Get** /non-employee-requests/summary/{requested-for} | Get Summary of Non-Employee Requests
[**GetNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeSchemaAttribute) | **Get** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Get Schema Attribute Non-Employee Source
[**GetNonEmployeeSource**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeSource) | **Get** /non-employee-sources/{sourceId} | Get a Non-Employee Source
[**GetNonEmployeeSourceSchemaAttributes**](NonEmployeeLifecycleManagementAPI.md#GetNonEmployeeSourceSchemaAttributes) | **Get** /non-employee-sources/{sourceId}/schema-attributes | List Schema Attributes Non-Employee Source
[**ImportNonEmployeeRecordsInBulk**](NonEmployeeLifecycleManagementAPI.md#ImportNonEmployeeRecordsInBulk) | **Post** /non-employee-sources/{id}/non-employee-bulk-upload | Imports, or Updates, Non-Employee Records
[**ListNonEmployeeApprovals**](NonEmployeeLifecycleManagementAPI.md#ListNonEmployeeApprovals) | **Get** /non-employee-approvals | Get List of Non-Employee Approval Requests
[**ListNonEmployeeRecords**](NonEmployeeLifecycleManagementAPI.md#ListNonEmployeeRecords) | **Get** /non-employee-records | List Non-Employee Records
[**ListNonEmployeeRequests**](NonEmployeeLifecycleManagementAPI.md#ListNonEmployeeRequests) | **Get** /non-employee-requests | List Non-Employee Requests
[**ListNonEmployeeSources**](NonEmployeeLifecycleManagementAPI.md#ListNonEmployeeSources) | **Get** /non-employee-sources | List Non-Employee Sources
[**PatchNonEmployeeRecord**](NonEmployeeLifecycleManagementAPI.md#PatchNonEmployeeRecord) | **Patch** /non-employee-records/{id} | Patch Non-Employee Record
[**PatchNonEmployeeSchemaAttribute**](NonEmployeeLifecycleManagementAPI.md#PatchNonEmployeeSchemaAttribute) | **Patch** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Patch a Schema Attribute for Non-Employee Source
[**PatchNonEmployeeSource**](NonEmployeeLifecycleManagementAPI.md#PatchNonEmployeeSource) | **Patch** /non-employee-sources/{sourceId} | Patch a Non-Employee Source
[**RejectNonEmployeeRequest**](NonEmployeeLifecycleManagementAPI.md#RejectNonEmployeeRequest) | **Post** /non-employee-approvals/{id}/reject | Reject a Non-Employee Request
[**UpdateNonEmployeeRecord**](NonEmployeeLifecycleManagementAPI.md#UpdateNonEmployeeRecord) | **Put** /non-employee-records/{id} | Update Non-Employee Record



## Approve a Non-Employee Request

> NonEmployeeApprovalItem ApproveNonEmployeeRequest(ctx, id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()

Approves a non-employee approval request and notifies the next approver. The current user must be the requested approver.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    nonEmployeeApprovalDecision := *sailpoint.NewNonEmployeeApprovalDecision() // NonEmployeeApprovalDecision | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ApproveNonEmployeeRequest(context.Background(), id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ApproveNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApproveNonEmployeeRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ApproveNonEmployeeRequest`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create Non-Employee Record

> NonEmployeeRecord CreateNonEmployeeRecord(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

This request will create a non-employee record.
Requires role context of `idn:nesr:create`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    nonEmployeeRequestBody := *sailpoint.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-Employee record creation request body.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRecord(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRecord`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create Non-Employee Request

> NonEmployeeRequest CreateNonEmployeeRequest(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

This request will create a non-employee request and notify the approver. Requires role context of `idn:nesr:create` or the user must own the source.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    nonEmployeeRequestBody := *sailpoint.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-Employee creation request body

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRequest(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeRequest`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeRequest`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create Non-Employee Source

> NonEmployeeSourceWithCloudExternalId CreateNonEmployeeSource(ctx).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()

This request will create a non-employee source. Requires role context of `idn:nesr:create`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    nonEmployeeSourceRequestBody := *sailpoint.NewNonEmployeeSourceRequestBody("Retail", "Source description", *sailpoint.NewNonEmployeeIdnUserRequest("2c91808570313110017040b06f344ec9")) // NonEmployeeSourceRequestBody | Non-Employee source creation request body.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSource(context.Background()).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeSource`: NonEmployeeSourceWithCloudExternalId
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSource`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create a new Schema Attribute for Non-Employee Source

> NonEmployeeSchemaAttribute CreateNonEmployeeSourceSchemaAttributes(ctx, sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()

This API creates a new schema attribute for Non-Employee Source. The schema technical name must be unique in the source. Attempts to create a schema attribute with an existing name will result in a "400.1.409 Reference conflict" response. At most, 10 custom attributes can be created per schema. Attempts to create more than 10 will result in a "400.1.4 Limit violation" response.
Requires role context of `idn:nesr:create`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id
    nonEmployeeSchemaAttributeBody := *sailpoint.NewNonEmployeeSchemaAttributeBody("TEXT", "Account Name", "account.name") // NonEmployeeSchemaAttributeBody | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSourceSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNonEmployeeSourceSchemaAttributes`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.CreateNonEmployeeSourceSchemaAttributes`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Non-Employee Record

> DeleteNonEmployeeRecord(ctx, id).Execute()

This request will delete a non-employee record.
Requires role context of `idn:nesr:delete`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-Employee record id (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRecord(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRecord``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Multiple Non-Employee Records

> DeleteNonEmployeeRecordsInBulk(ctx).DeleteNonEmployeeRecordsInBulkRequest(deleteNonEmployeeRecordsInBulkRequest).Execute()

This request will delete multiple non-employee records based on the non-employee ids provided. Requires role context of `idn:nesr:delete`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    deleteNonEmployeeRecordsInBulkRequest := *sailpoint.NewDeleteNonEmployeeRecordsInBulkRequest([]string{"Ids_example"}) // DeleteNonEmployeeRecordsInBulkRequest | Non-Employee bulk delete request body.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRecordsInBulk(context.Background()).DeleteNonEmployeeRecordsInBulkRequest(deleteNonEmployeeRecordsInBulkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRecordsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNonEmployeeRecordsInBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteNonEmployeeRecordsInBulkRequest** | [**DeleteNonEmployeeRecordsInBulkRequest**](DeleteNonEmployeeRecordsInBulkRequest.md) | Non-Employee bulk delete request body. | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Non-Employee Request

> DeleteNonEmployeeRequest(ctx, id).Execute()

This request will delete a non-employee request. 
Requires role context of `idn:nesr:delete`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ac110005-7156-1150-8171-5b292e3e0084" // string | Non-Employee request id in the UUID format

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeRequest``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete a Schema Attribute for Non-Employee Source

> DeleteNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).Execute()

This end-point deletes a specific schema attribute for a non-employee source.
Requires role context of `idn:nesr:delete`


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    attributeId := "ef38f94347e94562b5bb8424a56397d8" // string | The Schema Attribute Id (UUID)
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSchemaAttribute``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Non-Employee Source

> DeleteNonEmployeeSource(ctx, sourceId).Execute()

This request will delete a non-employee source. Requires role context of `idn:nesr:delete`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSource``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete all custom schema attributes for Non-Employee Source

> DeleteNonEmployeeSourceSchemaAttributes(ctx, sourceId).Execute()

This end-point deletes all custom schema attributes for a non-employee source. Requires role context of `idn:nesr:delete`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.DeleteNonEmployeeSourceSchemaAttributes``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Exports Non-Employee Records to CSV

> ExportNonEmployeeRecords(ctx, id).Execute()

This requests a CSV download for all non-employees from a provided source. Requires role context of `idn:nesr:read`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ExportNonEmployeeRecords(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ExportNonEmployeeRecords``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Exports Source Schema Template

> ExportNonEmployeeSourceSchemaTemplate(ctx, id).Execute()

This requests a download for the Source Schema Template for a provided source. Requires role context of `idn:nesr:read`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Source Id (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ExportNonEmployeeSourceSchemaTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ExportNonEmployeeSourceSchemaTemplate``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get a non-employee approval item detail

> NonEmployeeApprovalItemDetail GetNonEmployeeApproval(ctx, id).IncludeDetail(includeDetail).Execute()

Gets a non-employee approval item detail. There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:read`, in which case they
can get any approval.
  2. The user owns the requested approval.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    includeDetail := true // bool | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeApproval(context.Background(), id).IncludeDetail(includeDetail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeApproval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeApproval`: NonEmployeeApprovalItemDetail
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeApproval`: %v\n", resp)
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

 **includeDetail** | **bool** | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* | 

### Return type

[**NonEmployeeApprovalItemDetail**](NonEmployeeApprovalItemDetail.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Summary of Non-Employee Approval Requests

> NonEmployeeApprovalSummary GetNonEmployeeApprovalSummary(ctx, requestedFor).Execute()

This request will retrieve a summary of non-employee approval requests. There are two contextual uses for the `requested-for` path parameter:
  1. The user has the role context of `idn:nesr:read`, in which case he or
she may request a summary of all non-employee approval requests assigned to a particular approver by passing in that approver's id.
  2. The current user is an approver, in which case "me" should be provided
as the `requested-for` value. This will provide the approver with a summary of the approval items assigned to him or her.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity (UUID) of the approver for whom for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeApprovalSummary(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeApprovalSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeApprovalSummary`: NonEmployeeApprovalSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeApprovalSummary`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Obtain the status of bulk upload on the source

> NonEmployeeBulkUploadStatus GetNonEmployeeBulkUploadStatus(ctx, id).Execute()

The nonEmployeeBulkUploadStatus API returns the status of the newest bulk upload job for the specified source.
Requires role context of `idn:nesr:read`


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source ID (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeBulkUploadStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeBulkUploadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeBulkUploadStatus`: NonEmployeeBulkUploadStatus
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeBulkUploadStatus`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get a Non-Employee Record

> NonEmployeeRecord GetNonEmployeeRecord(ctx, id).Execute()

This gets a non-employee record.
Requires role context of `idn:nesr:read`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-Employee record id (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeRecord(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRecord`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get a Non-Employee Request

> NonEmployeeRequest GetNonEmployeeRequest(ctx, id).Execute()

This gets a non-employee request.
There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:read`, in this case the user
can get the non-employee request for any user.
  2. The user must be the owner of the non-employee request.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ac110005-7156-1150-8171-5b292e3e0084" // string | Non-Employee request id (UUID)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRequest`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequest`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Summary of Non-Employee Requests

> NonEmployeeRequestSummary GetNonEmployeeRequestSummary(ctx, requestedFor).Execute()

This request will retrieve a summary of non-employee requests. There are two contextual uses for the `requested-for` path parameter:
  1. The user has the role context of `idn:nesr:read`, in which case he or
she may request a summary of all non-employee approval requests assigned to a particular account manager by passing in that manager's id.
  2. The current user is an account manager, in which case "me" should be
provided as the `requested-for` value. This will provide the user with a summary of the non-employee requests in the source(s) he or she manages.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity (UUID) of the non-employee account manager for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequestSummary(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequestSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeRequestSummary`: NonEmployeeRequestSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeRequestSummary`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Schema Attribute Non-Employee Source

> NonEmployeeSchemaAttribute GetNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).Execute()

This API gets a schema attribute by Id for the specified Non-Employee SourceId. Requires role context of `idn:nesr:read` or the user must be an account manager of the source.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    attributeId := "ef38f94347e94562b5bb8424a56397d8" // string | The Schema Attribute Id (UUID)
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSchemaAttribute`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get a Non-Employee Source

> NonEmployeeSource GetNonEmployeeSource(ctx, sourceId).Execute()

This gets a non-employee source. There are two contextual uses for the requested-for path parameter: 
  1. The user has the role context of `idn:nesr:read`, in which case he or
she may request any source.
  2. The current user is an account manager, in which case the user can only
request sources that they own.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "2c91808b7c28b350017c2a2ec5790aa1" // string | Source Id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeSource(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSource`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSource`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Schema Attributes Non-Employee Source

> []NonEmployeeSchemaAttribute GetNonEmployeeSourceSchemaAttributes(ctx, sourceId).Execute()

This API gets the list of schema attributes for the specified Non-Employee SourceId. There are 8 mandatory attributes added to each new Non-Employee Source automatically. Additionaly, user can add up to 10 custom attributes. This interface returns all the mandatory attributes followed by any custom attributes. At most, a total of 18 attributes will be returned.
Requires role context of `idn:nesr:read` or the user must be an account manager of the source.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.GetNonEmployeeSourceSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSourceSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNonEmployeeSourceSchemaAttributes`: []NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.GetNonEmployeeSourceSchemaAttributes`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Imports, or Updates, Non-Employee Records

> NonEmployeeBulkUploadJob ImportNonEmployeeRecordsInBulk(ctx, id).Data(data).Execute()

This post will import, or update, Non-Employee records found in the CSV. Requires role context of `idn:nesr:create`

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id (UUID)
    data := os.NewFile(1234, "some_file") // *os.File | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ImportNonEmployeeRecordsInBulk(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ImportNonEmployeeRecordsInBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportNonEmployeeRecordsInBulk`: NonEmployeeBulkUploadJob
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ImportNonEmployeeRecordsInBulk`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get List of Non-Employee Approval Requests

> []NonEmployeeApprovalItem ListNonEmployeeApprovals(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

This gets a list of non-employee approval requests.
There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:read`, in which case they
can list the approvals for any approver.
  2. The user owns the requested approval.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity for whom the request was made. *me* indicates the current user. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "approvalStatus eq "Pending"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **approvalStatus**: *eq* (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ListNonEmployeeApprovals(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ListNonEmployeeApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeApprovals`: []NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ListNonEmployeeApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **approvalStatus**: *eq* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]NonEmployeeApprovalItem**](NonEmployeeApprovalItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Non-Employee Records

> []NonEmployeeRecord ListNonEmployeeRecords(ctx).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

This gets a list of non-employee records. There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:read`, in which case they can get a list of all of the non-employees.
  2. The user is an account manager, in which case they can get a list of the non-employees that they manage.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "accountName,sourceId" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **sourceId**: *eq* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ListNonEmployeeRecords(context.Background()).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ListNonEmployeeRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeRecords`: []NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ListNonEmployeeRecords`: %v\n", resp)
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
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **sourceId**: *eq* | 

### Return type

[**[]NonEmployeeRecord**](NonEmployeeRecord.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Non-Employee Requests

> []NonEmployeeRequest ListNonEmployeeRequests(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

This gets a list of non-employee requests. There are two contextual uses for the `requested-for` path parameter:
  1. The user has the role context of `idn:nesr:read`, in which case he or
she may request a list non-employee requests assigned to a particular account manager by passing in that manager's id.
  2. The current user is an account manager, in which case "me" should be
provided as the `requested-for` value. This will provide the user with a list of the non-employee requests in the source(s) he or she manages.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "e136567de87e4d029e60b3c3c55db56d" // string | The identity for whom the request was made. *me* indicates the current user.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "created,approvalStatus" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **sourceId**: *eq*  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ListNonEmployeeRequests(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ListNonEmployeeRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeRequests`: []NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ListNonEmployeeRequests`: %v\n", resp)
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
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **sourceId**: *eq*  | 

### Return type

[**[]NonEmployeeRequest**](NonEmployeeRequest.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Non-Employee Sources

> []NonEmployeeSourceWithNECount ListNonEmployeeSources(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).NonEmployeeCount(nonEmployeeCount).Sorters(sorters).Execute()

This gets a list of non-employee sources. There are two contextual uses for the requested-for path parameter: 
  1. The user has the role context of `idn:nesr:read`, in which case he or
she may request a list sources assigned to a particular account manager by passing in that manager's id.
  2. The current user is an account manager, in which case "me" should be
provided as the `requested-for` value. This will provide the user with a list of the sources that he or she owns.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestedFor := "me" // string | The identity for whom the request was made. *me* indicates the current user.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    nonEmployeeCount := true // bool | The flag to determine whether return a non-employee count associate with source. (optional)
    sorters := "name,created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, sourceId** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.ListNonEmployeeSources(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).NonEmployeeCount(nonEmployeeCount).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.ListNonEmployeeSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNonEmployeeSources`: []NonEmployeeSourceWithNECount
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.ListNonEmployeeSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNonEmployeeSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **nonEmployeeCount** | **bool** | The flag to determine whether return a non-employee count associate with source. | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, sourceId** | 

### Return type

[**[]NonEmployeeSourceWithNECount**](NonEmployeeSourceWithNECount.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch Non-Employee Record

> NonEmployeeRecord PatchNonEmployeeRecord(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

This request will patch a non-employee record. There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:update`, in which case they
update all available fields.
  2. The user is owner of the source, in this case they can only update the
end date.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-employee record id (UUID)
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.PatchNonEmployeeRecord(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeRecord`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch a Schema Attribute for Non-Employee Source

> NonEmployeeSchemaAttribute PatchNonEmployeeSchemaAttribute(ctx, attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

This end-point patches a specific schema attribute for a non-employee SourceId.
Requires role context of `idn:nesr:update`


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    attributeId := "ef38f94347e94562b5bb8424a56397d8" // string | The Schema Attribute Id (UUID)
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following properties are allowed for update ':' 'label', 'helpText', 'placeholder', 'required'.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSchemaAttribute(context.Background(), attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSchemaAttribute`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch a Non-Employee Source

> NonEmployeeSource PatchNonEmployeeSource(ctx, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

patch a non-employee source. (partial update) <br/> Patchable field: **name, description, approvers, accountManagers** Requires role context of `idn:nesr:update`.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sourceId := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee source update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSource(context.Background(), sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNonEmployeeSource`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.PatchNonEmployeeSource`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Reject a Non-Employee Request

> NonEmployeeApprovalItem RejectNonEmployeeRequest(ctx, id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()

This endpoint will reject an approval item request and notify user. The current user must be the requested approver.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    nonEmployeeRejectApprovalDecision := *sailpoint.NewNonEmployeeRejectApprovalDecision("approved") // NonEmployeeRejectApprovalDecision | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.RejectNonEmployeeRequest(context.Background(), id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.RejectNonEmployeeRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RejectNonEmployeeRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.RejectNonEmployeeRequest`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Non-Employee Record

> NonEmployeeRecord UpdateNonEmployeeRecord(ctx, id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

This request will update a non-employee record. There are two contextual uses for this endpoint:
  1. The user has the role context of `idn:nesr:update`, in which case they
update all available fields.
  2. The user is owner of the source, in this case they can only update the
end date.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-employee record id (UUID)
    nonEmployeeRequestBody := *sailpoint.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-employee record creation request body. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.NonEmployeeLifecycleManagementAPI.UpdateNonEmployeeRecord(context.Background(), id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementAPI.UpdateNonEmployeeRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNonEmployeeRecord`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementAPI.UpdateNonEmployeeRecord`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

