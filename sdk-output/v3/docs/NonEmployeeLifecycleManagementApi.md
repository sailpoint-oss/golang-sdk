# \NonEmployeeLifecycleManagementApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#CreateSchemaAttribute) | **Post** /non-employee-sources/{sourceId}/schema-attributes | Create a new Schema Attribute for Non-Employee Source
[**DeleteSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#DeleteSchemaAttribute) | **Delete** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Delete a Schema Attribute for Non-Employee Source
[**DeleteSchemaAttributes**](NonEmployeeLifecycleManagementApi.md#DeleteSchemaAttributes) | **Delete** /non-employee-sources/{sourceId}/schema-attributes | Delete all custom schema attributes for Non-Employee Source
[**GetSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#GetSchemaAttribute) | **Get** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Get Schema Attribute Non-Employee Source
[**GetSchemaAttributes**](NonEmployeeLifecycleManagementApi.md#GetSchemaAttributes) | **Get** /non-employee-sources/{sourceId}/schema-attributes | List Schema Attributes Non-Employee Source
[**NonEmployeeApprovalGet**](NonEmployeeLifecycleManagementApi.md#NonEmployeeApprovalGet) | **Get** /non-employee-approvals/{id} | Get a non-employee approval item detail
[**NonEmployeeApprovalList**](NonEmployeeLifecycleManagementApi.md#NonEmployeeApprovalList) | **Get** /non-employee-approvals | Get List of Non-Employee Approval Requests
[**NonEmployeeApprovalSummary**](NonEmployeeLifecycleManagementApi.md#NonEmployeeApprovalSummary) | **Get** /non-employee-approvals/summary/{requested-for} | Get Summary of Non-Employee Approval Requests
[**NonEmployeeApproveRequest**](NonEmployeeLifecycleManagementApi.md#NonEmployeeApproveRequest) | **Post** /non-employee-approvals/{id}/approve | Approve a Non-Employee Request
[**NonEmployeeBulkUploadStatus**](NonEmployeeLifecycleManagementApi.md#NonEmployeeBulkUploadStatus) | **Get** /non-employee-sources/{id}/non-employee-bulk-upload/status | Obtain the status of bulk upload on the source
[**NonEmployeeExportSourceSchemaTemplate**](NonEmployeeLifecycleManagementApi.md#NonEmployeeExportSourceSchemaTemplate) | **Get** /non-employee-sources/{id}/schema-attributes-template/download | Exports Source Schema Template
[**NonEmployeeRecordBulkDelete**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordBulkDelete) | **Post** /non-employee-records/bulk-delete | Delete Multiple Non-Employee Records
[**NonEmployeeRecordCreation**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordCreation) | **Post** /non-employee-records | Create Non-Employee Record
[**NonEmployeeRecordDelete**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordDelete) | **Delete** /non-employee-records/{id} | Delete Non-Employee Record
[**NonEmployeeRecordGet**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordGet) | **Get** /non-employee-records/{id} | Get a Non-Employee Record
[**NonEmployeeRecordList**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordList) | **Get** /non-employee-records | List Non-Employee Records
[**NonEmployeeRecordPatch**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordPatch) | **Patch** /non-employee-records/{id} | Patch Non-Employee Record
[**NonEmployeeRecordUpdate**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordUpdate) | **Put** /non-employee-records/{id} | Update Non-Employee Record
[**NonEmployeeRecordsBulkUpload**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordsBulkUpload) | **Post** /non-employee-sources/{id}/non-employee-bulk-upload | Imports, or Updates, Non-Employee Records
[**NonEmployeeRecordsExport**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRecordsExport) | **Get** /non-employee-sources/{id}/non-employees/download | Exports Non-Employee Records to CSV
[**NonEmployeeRejectRequest**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRejectRequest) | **Post** /non-employee-approvals/{id}/reject | Reject a Non-Employee Request
[**NonEmployeeRequestCreation**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRequestCreation) | **Post** /non-employee-requests | Create Non-Employee Request
[**NonEmployeeRequestDeletion**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRequestDeletion) | **Delete** /non-employee-requests/{id} | Delete Non-Employee Request
[**NonEmployeeRequestGet**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRequestGet) | **Get** /non-employee-requests/{id} | Get a Non-Employee Request
[**NonEmployeeRequestList**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRequestList) | **Get** /non-employee-requests | List Non-Employee Requests
[**NonEmployeeRequestSummaryGet**](NonEmployeeLifecycleManagementApi.md#NonEmployeeRequestSummaryGet) | **Get** /non-employee-requests/summary/{requested-for} | Get Summary of Non-Employee Requests
[**NonEmployeeSourceDelete**](NonEmployeeLifecycleManagementApi.md#NonEmployeeSourceDelete) | **Delete** /non-employee-sources/{sourceId} | Delete Non-Employee Source
[**NonEmployeeSourceGet**](NonEmployeeLifecycleManagementApi.md#NonEmployeeSourceGet) | **Get** /non-employee-sources/{sourceId} | Get a Non-Employee Source
[**NonEmployeeSourcePatch**](NonEmployeeLifecycleManagementApi.md#NonEmployeeSourcePatch) | **Patch** /non-employee-sources/{sourceId} | Patch a Non-Employee Source
[**NonEmployeeSourcesCreation**](NonEmployeeLifecycleManagementApi.md#NonEmployeeSourcesCreation) | **Post** /non-employee-sources | Create Non-Employee Source
[**NonEmployeeSourcesList**](NonEmployeeLifecycleManagementApi.md#NonEmployeeSourcesList) | **Get** /non-employee-sources | List Non-Employee Sources
[**PatchSchemaAttribute**](NonEmployeeLifecycleManagementApi.md#PatchSchemaAttribute) | **Patch** /non-employee-sources/{sourceId}/schema-attributes/{attributeId} | Patch a Schema Attribute for Non-Employee Source



## CreateSchemaAttribute

> NonEmployeeSchemaAttribute CreateSchemaAttribute(ctx, sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()

Create a new Schema Attribute for Non-Employee Source



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
    sourceId := "sourceId_example" // string | The Source id
    nonEmployeeSchemaAttributeBody := *openapiclient.NewNonEmployeeSchemaAttributeBody("TEXT", "Account Name", "account.name") // NonEmployeeSchemaAttributeBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.CreateSchemaAttribute(context.Background(), sourceId).NonEmployeeSchemaAttributeBody(nonEmployeeSchemaAttributeBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.CreateSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.CreateSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSchemaAttributeRequest struct via the builder pattern


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


## DeleteSchemaAttribute

> DeleteSchemaAttribute(ctx, attributeId, sourceId).Execute()

Delete a Schema Attribute for Non-Employee Source



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
    attributeId := "attributeId_example" // string | The Schema Attribute Id (UUID)
    sourceId := "sourceId_example" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteSchemaAttribute``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSchemaAttributeRequest struct via the builder pattern


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


## DeleteSchemaAttributes

> DeleteSchemaAttributes(ctx, sourceId).Execute()

Delete all custom schema attributes for Non-Employee Source



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
    sourceId := "sourceId_example" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.DeleteSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.DeleteSchemaAttributes``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSchemaAttributesRequest struct via the builder pattern


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


## GetSchemaAttribute

> NonEmployeeSchemaAttribute GetSchemaAttribute(ctx, attributeId, sourceId).Execute()

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
    attributeId := "ef38f94347e94562b5bb8424a56397d8" // string | The Schema Attribute Id (UUID)
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetSchemaAttribute(context.Background(), attributeId, sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **string** | The Schema Attribute Id (UUID) | 
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaAttributeRequest struct via the builder pattern


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


## GetSchemaAttributes

> []NonEmployeeSchemaAttribute GetSchemaAttributes(ctx, sourceId).Execute()

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
    sourceId := "ef38f94347e94562b5bb8424a56397d8" // string | The Source id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.GetSchemaAttributes(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.GetSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaAttributes`: []NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.GetSchemaAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaAttributesRequest struct via the builder pattern


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


## NonEmployeeApprovalGet

> NonEmployeeApprovalItemDetail NonEmployeeApprovalGet(ctx, id).IncludeDetail(includeDetail).Execute()

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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    includeDetail := true // bool | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalGet(context.Background(), id).IncludeDetail(includeDetail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeApprovalGet`: NonEmployeeApprovalItemDetail
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeApprovalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDetail** | **bool** | The object nonEmployeeRequest will not be included detail when set to false. *Default value is true* | 

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


## NonEmployeeApprovalList

> []NonEmployeeApprovalItem NonEmployeeApprovalList(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

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
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity for whom the request was made. *me* indicates the current user. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "approvalStatus eq "Pending"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)<br/><br/> Filtering is supported for the following fields and operators:<br/><br/> **approvalStatus**: *eq* <br/><br/> *Example:* approvalStatus eq \"PENDING\" (optional)
    sorters := "created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)<br/><br/> Sorting is supported for the following fields: **created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalList(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeApprovalList`: []NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeApprovalListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)&lt;br/&gt;&lt;br/&gt; Filtering is supported for the following fields and operators:&lt;br/&gt;&lt;br/&gt; **approvalStatus**: *eq* &lt;br/&gt;&lt;br/&gt; *Example:* approvalStatus eq \&quot;PENDING\&quot; | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)&lt;br/&gt;&lt;br/&gt; Sorting is supported for the following fields: **created, modified** | 

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


## NonEmployeeApprovalSummary

> NonEmployeeApprovalSummary NonEmployeeApprovalSummary(ctx, requestedFor).Execute()

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
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity (UUID) of the approver for whom for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApprovalSummary(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeApprovalSummary`: NonEmployeeApprovalSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeApprovalSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestedFor** | **string** | The identity (UUID) of the approver for whom for whom the summary is being retrieved. Use \&quot;me\&quot; instead to indicate the current user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeApprovalSummaryRequest struct via the builder pattern


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


## NonEmployeeApproveRequest

> NonEmployeeApprovalItem NonEmployeeApproveRequest(ctx, id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()

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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    nonEmployeeApprovalDecision := *openapiclient.NewNonEmployeeApprovalDecision() // NonEmployeeApprovalDecision | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeApproveRequest(context.Background(), id).NonEmployeeApprovalDecision(nonEmployeeApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeApproveRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeApproveRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeApproveRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeApproveRequestRequest struct via the builder pattern


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


## NonEmployeeBulkUploadStatus

> NonEmployeeBulkUploadStatus NonEmployeeBulkUploadStatus(ctx, id).Execute()

Obtain the status of bulk upload on the source



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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source ID (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeBulkUploadStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeBulkUploadStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeBulkUploadStatus`: NonEmployeeBulkUploadStatus
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeBulkUploadStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source ID (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeBulkUploadStatusRequest struct via the builder pattern


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


## NonEmployeeExportSourceSchemaTemplate

> NonEmployeeExportSourceSchemaTemplate(ctx, id).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Source Id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeExportSourceSchemaTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeExportSourceSchemaTemplate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNonEmployeeExportSourceSchemaTemplateRequest struct via the builder pattern


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


## NonEmployeeRecordBulkDelete

> NonEmployeeRecordBulkDelete(ctx).NonEmployeeRecordBulkDeleteRequest(nonEmployeeRecordBulkDeleteRequest).Execute()

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
    nonEmployeeRecordBulkDeleteRequest := *openapiclient.NewNonEmployeeRecordBulkDeleteRequest([]string{"Ids_example"}) // NonEmployeeRecordBulkDeleteRequest | Non-Employee bulk delete request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordBulkDelete(context.Background()).NonEmployeeRecordBulkDeleteRequest(nonEmployeeRecordBulkDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nonEmployeeRecordBulkDeleteRequest** | [**NonEmployeeRecordBulkDeleteRequest**](NonEmployeeRecordBulkDeleteRequest.md) | Non-Employee bulk delete request body. | 

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


## NonEmployeeRecordCreation

> NonEmployeeRecord NonEmployeeRecordCreation(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

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
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordCreation(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordCreation`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordCreation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordCreationRequest struct via the builder pattern


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


## NonEmployeeRecordDelete

> NonEmployeeRecordDelete(ctx, id).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-Employee record id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNonEmployeeRecordDeleteRequest struct via the builder pattern


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


## NonEmployeeRecordGet

> NonEmployeeRecord NonEmployeeRecordGet(ctx, id).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-Employee record id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordGet`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordGetRequest struct via the builder pattern


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


## NonEmployeeRecordList

> []NonEmployeeRecord NonEmployeeRecordList(ctx).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

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
    sorters := "accountName,sourceId" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)<br/><br/> Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)<br/><br/> Filtering is supported for the following fields and operators:<br/><br/> **sourceId**: *eq* <br/><br/> *Example:* sourceId eq \"2c91808568c529c60168cca6f90c1313\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordList(context.Background()).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordList`: []NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)&lt;br/&gt;&lt;br/&gt; Sorting is supported for the following fields: **id, accountName, sourceId, manager, firstName, lastName, email, phone, startDate, endDate, created, modified** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)&lt;br/&gt;&lt;br/&gt; Filtering is supported for the following fields and operators:&lt;br/&gt;&lt;br/&gt; **sourceId**: *eq* &lt;br/&gt;&lt;br/&gt; *Example:* sourceId eq \&quot;2c91808568c529c60168cca6f90c1313\&quot; | 

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


## NonEmployeeRecordPatch

> NonEmployeeRecord NonEmployeeRecordPatch(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-employee record id (UUID)
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordPatch(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordPatch`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordPatchRequest struct via the builder pattern


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


## NonEmployeeRecordUpdate

> NonEmployeeRecord NonEmployeeRecordUpdate(ctx, id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-employee record id (UUID)
    nonEmployeeRequestBody := *openapiclient.NewNonEmployeeRequestBody("william.smith", "William", "Smith", "william.smith@example.com", "5555555555", "jane.doe", "2c91808568c529c60168cca6f90c1313", time.Now(), time.Now()) // NonEmployeeRequestBody | Non-employee record creation request body. Attributes are restricted by user type. Owner of source can update end date. Organization admins can update all available fields.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordUpdate(context.Background(), id).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordUpdate`: NonEmployeeRecord
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-employee record id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordUpdateRequest struct via the builder pattern


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


## NonEmployeeRecordsBulkUpload

> NonEmployeeBulkUploadJob NonEmployeeRecordsBulkUpload(ctx, id).Data(data).Execute()

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
    data := "data_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordsBulkUpload(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordsBulkUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRecordsBulkUpload`: NonEmployeeBulkUploadJob
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRecordsBulkUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Source Id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRecordsBulkUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | **string** |  | 

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


## NonEmployeeRecordsExport

> NonEmployeeRecordsExport(ctx, id).Execute()

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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRecordsExport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRecordsExport``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNonEmployeeRecordsExportRequest struct via the builder pattern


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


## NonEmployeeRejectRequest

> NonEmployeeApprovalItem NonEmployeeRejectRequest(ctx, id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()

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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee approval item id (UUID)
    nonEmployeeRejectApprovalDecision := *openapiclient.NewNonEmployeeRejectApprovalDecision("approved") // NonEmployeeRejectApprovalDecision | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRejectRequest(context.Background(), id).NonEmployeeRejectApprovalDecision(nonEmployeeRejectApprovalDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRejectRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRejectRequest`: NonEmployeeApprovalItem
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRejectRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee approval item id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRejectRequestRequest struct via the builder pattern


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


## NonEmployeeRequestCreation

> NonEmployeeRequest NonEmployeeRequestCreation(ctx).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()

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
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestCreation(context.Background()).NonEmployeeRequestBody(nonEmployeeRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRequestCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRequestCreation`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRequestCreation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRequestCreationRequest struct via the builder pattern


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


## NonEmployeeRequestDeletion

> NonEmployeeRequestDeletion(ctx, id).Execute()

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
    id := "e136567de87e4d029e60b3c3c55db56d" // string | Non-Employee request id in the UUID format

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestDeletion(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRequestDeletion``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNonEmployeeRequestDeletionRequest struct via the builder pattern


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


## NonEmployeeRequestGet

> NonEmployeeRequest NonEmployeeRequestGet(ctx, id).Execute()

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Non-Employee request id (UUID)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRequestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRequestGet`: NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRequestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-Employee request id (UUID) | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRequestGetRequest struct via the builder pattern


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


## NonEmployeeRequestList

> []NonEmployeeRequest NonEmployeeRequestList(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

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
    requestedFor := "e136567de87e4d029e60b3c3c55db56d" // string | The identity for whom the request was made. *me* indicates the current user.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "created,approvalStatus" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)<br/><br/> Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** (optional)
    filters := "sourceId eq "2c91808568c529c60168cca6f90c1313"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)<br/><br/> Filtering is supported for the following fields and operators:<br/><br/> **sourceId**: *eq* <br/><br/> *Example:* sourceId eq \"2c91808568c529c60168cca6f90c1313\" (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestList(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRequestList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRequestList`: []NonEmployeeRequest
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRequestList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRequestListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)&lt;br/&gt;&lt;br/&gt; Sorting is supported for the following fields: **created, approvalStatus, firstName, lastName, email, phone, accountName, startDate, endDate** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407)&lt;br/&gt;&lt;br/&gt; Filtering is supported for the following fields and operators:&lt;br/&gt;&lt;br/&gt; **sourceId**: *eq* &lt;br/&gt;&lt;br/&gt; *Example:* sourceId eq \&quot;2c91808568c529c60168cca6f90c1313\&quot; | 

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


## NonEmployeeRequestSummaryGet

> NonEmployeeRequestSummary NonEmployeeRequestSummaryGet(ctx, requestedFor).Execute()

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
    requestedFor := "2c91808280430dfb0180431a59440460" // string | The identity (UUID) of the non-employee account manager for whom the summary is being retrieved. Use \"me\" instead to indicate the current user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeRequestSummaryGet(context.Background(), requestedFor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeRequestSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeRequestSummaryGet`: NonEmployeeRequestSummary
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeRequestSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestedFor** | **string** | The identity (UUID) of the non-employee account manager for whom the summary is being retrieved. Use \&quot;me\&quot; instead to indicate the current user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeRequestSummaryGetRequest struct via the builder pattern


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


## NonEmployeeSourceDelete

> NonEmployeeSourceDelete(ctx, sourceId).Execute()

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
    sourceId := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourceDelete(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeSourceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiNonEmployeeSourceDeleteRequest struct via the builder pattern


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


## NonEmployeeSourceGet

> NonEmployeeSource NonEmployeeSourceGet(ctx, sourceId).Execute()

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
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourceGet(context.Background(), sourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeSourceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeSourceGet`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeSourceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeSourceGetRequest struct via the builder pattern


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


## NonEmployeeSourcePatch

> NonEmployeeSource NonEmployeeSourcePatch(ctx, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

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
    sourceId := "e136567de87e4d029e60b3c3c55db56d" // string | Source Id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of non-employee source update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcePatch(context.Background(), sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeSourcePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeSourcePatch`: NonEmployeeSource
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeSourcePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Source Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeSourcePatchRequest struct via the builder pattern


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


## NonEmployeeSourcesCreation

> NonEmployeeSourceWithCloudExternalId NonEmployeeSourcesCreation(ctx).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()

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
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcesCreation(context.Background()).NonEmployeeSourceRequestBody(nonEmployeeSourceRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeSourcesCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeSourcesCreation`: NonEmployeeSourceWithCloudExternalId
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeSourcesCreation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeSourcesCreationRequest struct via the builder pattern


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


## NonEmployeeSourcesList

> []NonEmployeeSourceWithNECount NonEmployeeSourcesList(ctx).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).NonEmployeeCount(nonEmployeeCount).Sorters(sorters).Execute()

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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    nonEmployeeCount := true // bool | The flag to determine whether return a non-employee count associate with source. (optional)
    sorters := "name,created" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)<br/><br/> Sorting is supported for the following fields: **name, created** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.NonEmployeeSourcesList(context.Background()).RequestedFor(requestedFor).Limit(limit).Offset(offset).Count(count).NonEmployeeCount(nonEmployeeCount).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.NonEmployeeSourcesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NonEmployeeSourcesList`: []NonEmployeeSourceWithNECount
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.NonEmployeeSourcesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNonEmployeeSourcesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestedFor** | **string** | The identity for whom the request was made. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **nonEmployeeCount** | **bool** | The flag to determine whether return a non-employee count associate with source. | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://community.sailpoint.com/t5/IdentityNow-Wiki/V3-API-Standard-Collection-Parameters/ta-p/156407#toc-hId-2058949)&lt;br/&gt;&lt;br/&gt; Sorting is supported for the following fields: **name, created** | 

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


## PatchSchemaAttribute

> NonEmployeeSchemaAttribute PatchSchemaAttribute(ctx, attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a Schema Attribute for Non-Employee Source



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
    attributeId := "attributeId_example" // string | The Schema Attribute Id (UUID)
    sourceId := "sourceId_example" // string | The Source id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of schema attribute update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The following properties are allowed for update ':' 'label', 'helpText', 'placeholder', 'required'.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NonEmployeeLifecycleManagementApi.PatchSchemaAttribute(context.Background(), attributeId, sourceId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NonEmployeeLifecycleManagementApi.PatchSchemaAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchSchemaAttribute`: NonEmployeeSchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `NonEmployeeLifecycleManagementApi.PatchSchemaAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **string** | The Schema Attribute Id (UUID) | 
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSchemaAttributeRequest struct via the builder pattern


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

