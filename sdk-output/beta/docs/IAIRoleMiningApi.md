# \IAIRoleMiningApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePotentialRoleProvisionRequest**](IAIRoleMiningApi.md#CreatePotentialRoleProvisionRequest) | **Post** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/provision | Create request to provision a potential role into an actual role.
[**DownloadRoleMiningPotentialRoleZip**](IAIRoleMiningApi.md#DownloadRoleMiningPotentialRoleZip) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/export-async/{exportId}/download | Export (download) details for a potential role in a role mining session
[**EditEntitlementsPotentialRole**](IAIRoleMiningApi.md#EditEntitlementsPotentialRole) | **Post** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/edit-entitlements | Edit entitlements for a potential role to exclude some entitlements
[**ExportRoleMiningPotentialRole**](IAIRoleMiningApi.md#ExportRoleMiningPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/export | Export (download) details for a potential role in a role mining session
[**ExportRoleMiningPotentialRoleAsync**](IAIRoleMiningApi.md#ExportRoleMiningPotentialRoleAsync) | **Post** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/export-async | Asynchronously export details for a potential role in a role mining session and upload to S3
[**ExportRoleMiningPotentialRoleStatus**](IAIRoleMiningApi.md#ExportRoleMiningPotentialRoleStatus) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/export-async/{exportId} | Retrieve status of a potential role export job
[**GetEntitlementDistributionPotentialRole**](IAIRoleMiningApi.md#GetEntitlementDistributionPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/entitlement-popularity-distribution | Retrieves entitlement popularity distribution for a potential role in a role mining session
[**GetEntitlementsPotentialRole**](IAIRoleMiningApi.md#GetEntitlementsPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/entitlement-popularities | Retrieves entitlements for a potential role in a role mining session
[**GetExcludedEntitlementsPotentialRole**](IAIRoleMiningApi.md#GetExcludedEntitlementsPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/excluded-entitlements | Retrieves excluded entitlements for a potential role in a role mining session
[**GetIdentitiesPotentialRole**](IAIRoleMiningApi.md#GetIdentitiesPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-roles/{potentialRoleId}/identities | Retrieves identities for a potential role in a role mining session
[**GetPotentialRole**](IAIRoleMiningApi.md#GetPotentialRole) | **Get** /role-mining-sessions/{sessionId}/potential-role-summaries/{potentialRoleId} | Retrieves a specific potential role
[**GetPotentialRoleApplications**](IAIRoleMiningApi.md#GetPotentialRoleApplications) | **Get** /role-mining-sessions/{sessionId}/potential-role-summaries/{potentialRoleId}/applications | Retrieves the applications of a potential role for a role mining session
[**GetPotentialRoleSummaries**](IAIRoleMiningApi.md#GetPotentialRoleSummaries) | **Get** /role-mining-sessions/{sessionId}/potential-role-summaries | Retrieves all potential role summaries
[**GetRoleMiningSession**](IAIRoleMiningApi.md#GetRoleMiningSession) | **Get** /role-mining-sessions/{sessionId} | Get a role mining session
[**GetRoleMiningSessionStatus**](IAIRoleMiningApi.md#GetRoleMiningSessionStatus) | **Get** /role-mining-sessions/{sessionId}/status | Get role mining session status state
[**GetRoleMiningSessions**](IAIRoleMiningApi.md#GetRoleMiningSessions) | **Get** /role-mining-sessions | Retrieves all role mining sessions
[**PatchPotentialRole**](IAIRoleMiningApi.md#PatchPotentialRole) | **Patch** /role-mining-sessions/{sessionId}/potential-role-summaries/{potentialRoleId} | Update a potential role
[**PatchRoleMiningSession**](IAIRoleMiningApi.md#PatchRoleMiningSession) | **Patch** /role-mining-sessions/{sessionId} | Patch a role mining session
[**RoleMiningSessions**](IAIRoleMiningApi.md#RoleMiningSessions) | **Post** /role-mining-sessions | Create a role mining session



## CreatePotentialRoleProvisionRequest

> RoleMiningPotentialRoleSummary CreatePotentialRoleProvisionRequest(ctx, sessionId, potentialRoleId).MinEntitlementPopularity(minEntitlementPopularity).IncludeCommonAccess(includeCommonAccess).RoleMiningPotentialRoleProvisionRequest(roleMiningPotentialRoleProvisionRequest).Execute()

Create request to provision a potential role into an actual role.



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    minEntitlementPopularity := int32(56) // int32 | Minimum popularity required for an entitlement to be included in the provisioned role. (optional) (default to 0)
    includeCommonAccess := true // bool | Boolean determining whether common access entitlements will be included in the provisioned role. (optional) (default to true)
    roleMiningPotentialRoleProvisionRequest := *openapiclient.NewRoleMiningPotentialRoleProvisionRequest() // RoleMiningPotentialRoleProvisionRequest | Required information to create a new role (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.CreatePotentialRoleProvisionRequest(context.Background(), sessionId, potentialRoleId).MinEntitlementPopularity(minEntitlementPopularity).IncludeCommonAccess(includeCommonAccess).RoleMiningPotentialRoleProvisionRequest(roleMiningPotentialRoleProvisionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.CreatePotentialRoleProvisionRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePotentialRoleProvisionRequest`: RoleMiningPotentialRoleSummary
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.CreatePotentialRoleProvisionRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePotentialRoleProvisionRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **minEntitlementPopularity** | **int32** | Minimum popularity required for an entitlement to be included in the provisioned role. | [default to 0]
 **includeCommonAccess** | **bool** | Boolean determining whether common access entitlements will be included in the provisioned role. | [default to true]
 **roleMiningPotentialRoleProvisionRequest** | [**RoleMiningPotentialRoleProvisionRequest**](RoleMiningPotentialRoleProvisionRequest.md) | Required information to create a new role | 

### Return type

[**RoleMiningPotentialRoleSummary**](RoleMiningPotentialRoleSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRoleMiningPotentialRoleZip

> *os.File DownloadRoleMiningPotentialRoleZip(ctx, sessionId, potentialRoleId, exportId).Execute()

Export (download) details for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "278359a6-04b7-4669-9468-924cf580964a" // string | A potential role id in a role mining session
    exportId := "4940ffd4-836f-48a3-b2b0-6d498c3fdf40" // string | The id of a previously run export job for this potential role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.DownloadRoleMiningPotentialRoleZip(context.Background(), sessionId, potentialRoleId, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.DownloadRoleMiningPotentialRoleZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadRoleMiningPotentialRoleZip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.DownloadRoleMiningPotentialRoleZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 
**exportId** | **string** | The id of a previously run export job for this potential role | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRoleMiningPotentialRoleZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditEntitlementsPotentialRole

> RoleMiningPotentialRole EditEntitlementsPotentialRole(ctx, sessionId, potentialRoleId).RoleMiningPotentialRoleEditEntitlements(roleMiningPotentialRoleEditEntitlements).Execute()

Edit entitlements for a potential role to exclude some entitlements



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    roleMiningPotentialRoleEditEntitlements := *openapiclient.NewRoleMiningPotentialRoleEditEntitlements() // RoleMiningPotentialRoleEditEntitlements | Role mining session parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.EditEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).RoleMiningPotentialRoleEditEntitlements(roleMiningPotentialRoleEditEntitlements).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.EditEntitlementsPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEntitlementsPotentialRole`: RoleMiningPotentialRole
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.EditEntitlementsPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEntitlementsPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleMiningPotentialRoleEditEntitlements** | [**RoleMiningPotentialRoleEditEntitlements**](RoleMiningPotentialRoleEditEntitlements.md) | Role mining session parameters | 

### Return type

[**RoleMiningPotentialRole**](RoleMiningPotentialRole.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRoleMiningPotentialRole

> *os.File ExportRoleMiningPotentialRole(ctx, sessionId, potentialRoleId).Execute()

Export (download) details for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.ExportRoleMiningPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRoleMiningPotentialRole`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.ExportRoleMiningPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRoleMiningPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRoleMiningPotentialRoleAsync

> RoleMiningPotentialRoleExportResponse ExportRoleMiningPotentialRoleAsync(ctx, sessionId, potentialRoleId).RoleMiningPotentialRoleExportRequest(roleMiningPotentialRoleExportRequest).Execute()

Asynchronously export details for a potential role in a role mining session and upload to S3



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "278359a6-04b7-4669-9468-924cf580964a" // string | A potential role id in a role mining session
    roleMiningPotentialRoleExportRequest := *openapiclient.NewRoleMiningPotentialRoleExportRequest() // RoleMiningPotentialRoleExportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRoleAsync(context.Background(), sessionId, potentialRoleId).RoleMiningPotentialRoleExportRequest(roleMiningPotentialRoleExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.ExportRoleMiningPotentialRoleAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRoleMiningPotentialRoleAsync`: RoleMiningPotentialRoleExportResponse
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.ExportRoleMiningPotentialRoleAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRoleMiningPotentialRoleAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleMiningPotentialRoleExportRequest** | [**RoleMiningPotentialRoleExportRequest**](RoleMiningPotentialRoleExportRequest.md) |  | 

### Return type

[**RoleMiningPotentialRoleExportResponse**](RoleMiningPotentialRoleExportResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportRoleMiningPotentialRoleStatus

> RoleMiningPotentialRoleExportResponse ExportRoleMiningPotentialRoleStatus(ctx, sessionId, potentialRoleId, exportId).Execute()

Retrieve status of a potential role export job



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "278359a6-04b7-4669-9468-924cf580964a" // string | A potential role id in a role mining session
    exportId := "4940ffd4-836f-48a3-b2b0-6d498c3fdf40" // string | The id of a previously run export job for this potential role

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRoleStatus(context.Background(), sessionId, potentialRoleId, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.ExportRoleMiningPotentialRoleStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportRoleMiningPotentialRoleStatus`: RoleMiningPotentialRoleExportResponse
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.ExportRoleMiningPotentialRoleStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 
**exportId** | **string** | The id of a previously run export job for this potential role | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportRoleMiningPotentialRoleStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleMiningPotentialRoleExportResponse**](RoleMiningPotentialRoleExportResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementDistributionPotentialRole

> map[string]int32 GetEntitlementDistributionPotentialRole(ctx, sessionId, potentialRoleId).IncludeCommonAccess(includeCommonAccess).Execute()

Retrieves entitlement popularity distribution for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    includeCommonAccess := true // bool | Boolean determining whether common access entitlements will be included or not (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetEntitlementDistributionPotentialRole(context.Background(), sessionId, potentialRoleId).IncludeCommonAccess(includeCommonAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetEntitlementDistributionPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlementDistributionPotentialRole`: map[string]int32
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetEntitlementDistributionPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementDistributionPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCommonAccess** | **bool** | Boolean determining whether common access entitlements will be included or not | 

### Return type

**map[string]int32**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementsPotentialRole

> []RoleMiningEntitlement GetEntitlementsPotentialRole(ctx, sessionId, potentialRoleId).IncludeCommonAccess(includeCommonAccess).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves entitlements for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    includeCommonAccess := true // bool | Boolean determining whether common access entitlements will be included or not (optional)
    sorters := "sorters_example" // string | sort by popularity, default order descending (optional)
    filters := "filters_example" // string | Filter parameter by \"starts with\" for the applicationName and entitlementRef.name. (optional)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).IncludeCommonAccess(includeCommonAccess).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetEntitlementsPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlementsPotentialRole`: []RoleMiningEntitlement
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetEntitlementsPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementsPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCommonAccess** | **bool** | Boolean determining whether common access entitlements will be included or not | 
 **sorters** | **string** | sort by popularity, default order descending | 
 **filters** | **string** | Filter parameter by \&quot;starts with\&quot; for the applicationName and entitlementRef.name. | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningEntitlement**](RoleMiningEntitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExcludedEntitlementsPotentialRole

> []RoleMiningEntitlement GetExcludedEntitlementsPotentialRole(ctx, sessionId, potentialRoleId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves excluded entitlements for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    sorters := "sorters_example" // string | sort by popularity, default order descending (optional)
    filters := "filters_example" // string | Filter parameter by \"starts with\" for the applicationName and entitlementRef.name. (optional)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetExcludedEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetExcludedEntitlementsPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExcludedEntitlementsPotentialRole`: []RoleMiningEntitlement
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetExcludedEntitlementsPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExcludedEntitlementsPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sorters** | **string** | sort by popularity, default order descending | 
 **filters** | **string** | Filter parameter by \&quot;starts with\&quot; for the applicationName and entitlementRef.name. | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningEntitlement**](RoleMiningEntitlement.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentitiesPotentialRole

> []RoleMiningIdentity GetIdentitiesPotentialRole(ctx, sessionId, potentialRoleId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves identities for a potential role in a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    sorters := "sorters_example" // string | sort by name, default order ascending (optional)
    filters := "filters_example" // string | Filter parameter by \"starts with\" for the name. (optional)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetIdentitiesPotentialRole(context.Background(), sessionId, potentialRoleId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetIdentitiesPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitiesPotentialRole`: []RoleMiningIdentity
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetIdentitiesPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitiesPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sorters** | **string** | sort by name, default order ascending | 
 **filters** | **string** | Filter parameter by \&quot;starts with\&quot; for the name. | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningIdentity**](RoleMiningIdentity.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPotentialRole

> RoleMiningPotentialRole GetPotentialRole(ctx, sessionId, potentialRoleId).Execute()

Retrieves a specific potential role



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPotentialRole`: RoleMiningPotentialRole
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleMiningPotentialRole**](RoleMiningPotentialRole.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPotentialRoleApplications

> []RoleMiningPotentialRoleApplication GetPotentialRoleApplications(ctx, sessionId, potentialRoleId).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves the applications of a potential role for a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | A potential role id in a role mining session
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetPotentialRoleApplications(context.Background(), sessionId, potentialRoleId).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetPotentialRoleApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPotentialRoleApplications`: []RoleMiningPotentialRoleApplication
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetPotentialRoleApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | A potential role id in a role mining session | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPotentialRoleApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningPotentialRoleApplication**](RoleMiningPotentialRoleApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPotentialRoleSummaries

> []RoleMiningPotentialRoleSummary GetPotentialRoleSummaries(ctx, sessionId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves all potential role summaries



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    sorters := "createdDate" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/) Sorting is supported for the following fields: **createdDate** (optional)
    filters := "(createdByName co "int")and (createdById sw "2c9180907")and (type eq "COMMON")and ((name co "entt")or (saved eq true))" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/) Filtering is supported for the following fields and operators: **createdById**: *eq, sw, co* **createdByName**: *eq, sw, co* **description**: *sw, co* **endDate**: *le, lt* **freshness**: *eq, ge, gt, le, lt* **name**: *eq, sw, co* **quality**: *eq, ge, gt, le, lt* **startDate**: *ge, gt* **saved**: *eq* **type**: *eq* (optional)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetPotentialRoleSummaries(context.Background(), sessionId).Sorters(sorters).Filters(filters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetPotentialRoleSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPotentialRoleSummaries`: []RoleMiningPotentialRoleSummary
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetPotentialRoleSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPotentialRoleSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/) Sorting is supported for the following fields: **createdDate** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters/) Filtering is supported for the following fields and operators: **createdById**: *eq, sw, co* **createdByName**: *eq, sw, co* **description**: *sw, co* **endDate**: *le, lt* **freshness**: *eq, ge, gt, le, lt* **name**: *eq, sw, co* **quality**: *eq, ge, gt, le, lt* **startDate**: *ge, gt* **saved**: *eq* **type**: *eq* | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningPotentialRoleSummary**](RoleMiningPotentialRoleSummary.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMiningSession

> RoleMiningSessionResponse GetRoleMiningSession(ctx, sessionId).Execute()

Get a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id to be retrieved.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetRoleMiningSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetRoleMiningSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMiningSession`: RoleMiningSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetRoleMiningSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMiningSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleMiningSessionResponse**](RoleMiningSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMiningSessionStatus

> RoleMiningSessionStatus GetRoleMiningSessionStatus(ctx, sessionId).Execute()

Get role mining session status state



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetRoleMiningSessionStatus(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetRoleMiningSessionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMiningSessionStatus`: RoleMiningSessionStatus
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetRoleMiningSessionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMiningSessionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleMiningSessionStatus**](RoleMiningSessionStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMiningSessions

> []RoleMiningSessionDto GetRoleMiningSessions(ctx).Filters(filters).Sorters(sorters).Offset(offset).Limit(limit).Count(count).Execute()

Retrieves all role mining sessions



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
    filters := "saved eq "true" and name sw "RM Session"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **saved**: *eq* \"true\" or \"false\" **name**: *eq, sw* (optional)
    sorters := "createdBy,createdDate" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **createdBy, createdDate** (optional)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.GetRoleMiningSessions(context.Background()).Filters(filters).Sorters(sorters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.GetRoleMiningSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMiningSessions`: []RoleMiningSessionDto
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.GetRoleMiningSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMiningSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **saved**: *eq* \&quot;true\&quot; or \&quot;false\&quot; **name**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **createdBy, createdDate** | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]RoleMiningSessionDto**](RoleMiningSessionDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPotentialRole

> map[string]interface{} PatchPotentialRole(ctx, sessionId, potentialRoleId).PatchPotentialRoleRequestInner(patchPotentialRoleRequestInner).Execute()

Update a potential role



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id
    potentialRoleId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The potential role summary id
    patchPotentialRoleRequestInner := []openapiclient.PatchPotentialRoleRequestInner{*openapiclient.NewPatchPotentialRoleRequestInner("replace", "/description")} // []PatchPotentialRoleRequestInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.PatchPotentialRole(context.Background(), sessionId, potentialRoleId).PatchPotentialRoleRequestInner(patchPotentialRoleRequestInner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.PatchPotentialRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPotentialRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.PatchPotentialRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id | 
**potentialRoleId** | **string** | The potential role summary id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPotentialRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **patchPotentialRoleRequestInner** | [**[]PatchPotentialRoleRequestInner**](PatchPotentialRoleRequestInner.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRoleMiningSession

> map[string]interface{} PatchRoleMiningSession(ctx, sessionId).JsonPatchOperation(jsonPatchOperation).Execute()

Patch a role mining session



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
    sessionId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role mining session id to be patched
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | Replace pruneThreshold and/or minNumIdentitiesInPotentialRole in role mining session. Update saved status or saved name for a role mining session.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.PatchRoleMiningSession(context.Background(), sessionId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.PatchRoleMiningSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRoleMiningSession`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.PatchRoleMiningSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The role mining session id to be patched | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoleMiningSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | Replace pruneThreshold and/or minNumIdentitiesInPotentialRole in role mining session. Update saved status or saved name for a role mining session. | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RoleMiningSessions

> RoleMiningSessionResponse RoleMiningSessions(ctx).RoleMiningSessionDto(roleMiningSessionDto).Execute()

Create a role mining session



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
    roleMiningSessionDto := *openapiclient.NewRoleMiningSessionDto() // RoleMiningSessionDto | Role mining session parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IAIRoleMiningApi.RoleMiningSessions(context.Background()).RoleMiningSessionDto(roleMiningSessionDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IAIRoleMiningApi.RoleMiningSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoleMiningSessions`: RoleMiningSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `IAIRoleMiningApi.RoleMiningSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRoleMiningSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleMiningSessionDto** | [**RoleMiningSessionDto**](RoleMiningSessionDto.md) | Role mining session parameters | 

### Return type

[**RoleMiningSessionResponse**](RoleMiningSessionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

