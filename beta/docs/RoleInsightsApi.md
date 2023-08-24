# \RoleInsightsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoleInsightRequests**](RoleInsightsApi.md#CreateRoleInsightRequests) | **Post** /role-insights/requests | A request to generate insights for roles
[**DownloadRoleInsightsEntitlementsChanges**](RoleInsightsApi.md#DownloadRoleInsightsEntitlementsChanges) | **Get** /role-insights/{insightId}/entitlement-changes/download | Download entitlement insights for a role
[**GetEntitlementChangesIdentities**](RoleInsightsApi.md#GetEntitlementChangesIdentities) | **Get** /role-insights/{insightId}/entitlement-changes/{entitlementId}/identities | Get identities for a suggested entitlement (for a role)
[**GetRoleInsight**](RoleInsightsApi.md#GetRoleInsight) | **Get** /role-insights/{insightId} | Get a single role insight
[**GetRoleInsights**](RoleInsightsApi.md#GetRoleInsights) | **Get** /role-insights | Get role insights
[**GetRoleInsightsCurrentEntitlements**](RoleInsightsApi.md#GetRoleInsightsCurrentEntitlements) | **Get** /role-insights/{insightId}/current-entitlements | Get current entitlement for a role
[**GetRoleInsightsEntitlementsChanges**](RoleInsightsApi.md#GetRoleInsightsEntitlementsChanges) | **Get** /role-insights/{insightId}/entitlement-changes | Get entitlement insights for a role
[**GetRoleInsightsRequests**](RoleInsightsApi.md#GetRoleInsightsRequests) | **Get** /role-insights/requests/{id} | Returns the metadata for a request in order to generate insights for roles.
[**GetRoleInsightsSummary**](RoleInsightsApi.md#GetRoleInsightsSummary) | **Get** /role-insights/summary | Get role insights summary information



## CreateRoleInsightRequests

> RoleInsightsResponse CreateRoleInsightRequests(ctx).Execute()

A request to generate insights for roles



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
    resp, r, err := apiClient.RoleInsightsApi.CreateRoleInsightRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.CreateRoleInsightRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRoleInsightRequests`: RoleInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.CreateRoleInsightRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleInsightRequestsRequest struct via the builder pattern


### Return type

[**RoleInsightsResponse**](RoleInsightsResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRoleInsightsEntitlementsChanges

> string DownloadRoleInsightsEntitlementsChanges(ctx, insightId).Sorters(sorters).Filters(filters).Execute()

Download entitlement insights for a role



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
    insightId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insight id
    sorters := "sorters_example" // string | sort by identitiesWithAccess, default order descending (optional)
    filters := "filters_example" // string | Filter parameter(s) by \"starts with\" for the name and description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.DownloadRoleInsightsEntitlementsChanges(context.Background(), insightId).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.DownloadRoleInsightsEntitlementsChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadRoleInsightsEntitlementsChanges`: string
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.DownloadRoleInsightsEntitlementsChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** | The role insight id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRoleInsightsEntitlementsChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorters** | **string** | sort by identitiesWithAccess, default order descending | 
 **filters** | **string** | Filter parameter(s) by \&quot;starts with\&quot; for the name and description. | 

### Return type

**string**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementChangesIdentities

> []RoleInsightsIdentities GetEntitlementChangesIdentities(ctx, insightId, entitlementId).HasEntitlement(hasEntitlement).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Execute()

Get identities for a suggested entitlement (for a role)



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
    insightId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insight id
    entitlementId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The entitlement id
    hasEntitlement := true // bool | Identity has this entitlement or not (optional) (default to false)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "sorters_example" // string | sort by name (optional)
    filters := "filters_example" // string | Filter parameter by \"starts with\" for the name. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetEntitlementChangesIdentities(context.Background(), insightId, entitlementId).HasEntitlement(hasEntitlement).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetEntitlementChangesIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntitlementChangesIdentities`: []RoleInsightsIdentities
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetEntitlementChangesIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** | The role insight id | 
**entitlementId** | **string** | The entitlement id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementChangesIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hasEntitlement** | **bool** | Identity has this entitlement or not | [default to false]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | sort by name | 
 **filters** | **string** | Filter parameter by \&quot;starts with\&quot; for the name. | 

### Return type

[**[]RoleInsightsIdentities**](RoleInsightsIdentities.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsight

> RoleInsight GetRoleInsight(ctx, insightId).Execute()

Get a single role insight



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
    insightId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insight id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsight(context.Background(), insightId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsight`: RoleInsight
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** | The role insight id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleInsight**](RoleInsight.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsights

> []RoleInsight GetRoleInsights(ctx).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Execute()

Get role insights



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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "sorters_example" // string | sort by numberOfUpdates, identitiesWithAccess, totalNumberOfIdentities (default- ascending) (optional)
    filters := "filters_example" // string | Filter parameter(s) by \"starts with\" for the name, ownerName and description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsights(context.Background()).Offset(offset).Limit(limit).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsights`: []RoleInsight
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsights`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | sort by numberOfUpdates, identitiesWithAccess, totalNumberOfIdentities (default- ascending) | 
 **filters** | **string** | Filter parameter(s) by \&quot;starts with\&quot; for the name, ownerName and description. | 

### Return type

[**[]RoleInsight**](RoleInsight.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsightsCurrentEntitlements

> []RoleInsightsEntitlement GetRoleInsightsCurrentEntitlements(ctx, insightId).Filters(filters).Execute()

Get current entitlement for a role



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
    insightId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insight id
    filters := "filters_example" // string | Filter parameter(s) by \"starts with\" for the name and description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsightsCurrentEntitlements(context.Background(), insightId).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsightsCurrentEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsightsCurrentEntitlements`: []RoleInsightsEntitlement
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsightsCurrentEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** | The role insight id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightsCurrentEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | Filter parameter(s) by \&quot;starts with\&quot; for the name and description. | 

### Return type

[**[]RoleInsightsEntitlement**](RoleInsightsEntitlement.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsightsEntitlementsChanges

> []RoleInsightsEntitlementChanges GetRoleInsightsEntitlementsChanges(ctx, insightId).Sorters(sorters).Filters(filters).Execute()

Get entitlement insights for a role



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
    insightId := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insight id
    sorters := "sorters_example" // string | sort by identitiesWithAccess or name (optional)
    filters := "filters_example" // string | Filter parameter(s) by \"starts with\" for the name and description. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsightsEntitlementsChanges(context.Background(), insightId).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsightsEntitlementsChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsightsEntitlementsChanges`: []RoleInsightsEntitlementChanges
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsightsEntitlementsChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**insightId** | **string** | The role insight id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightsEntitlementsChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorters** | **string** | sort by identitiesWithAccess or name | 
 **filters** | **string** | Filter parameter(s) by \&quot;starts with\&quot; for the name and description. | 

### Return type

[**[]RoleInsightsEntitlementChanges**](RoleInsightsEntitlementChanges.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsightsRequests

> RoleInsightsResponse GetRoleInsightsRequests(ctx, id).Execute()

Returns the metadata for a request in order to generate insights for roles.



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
    id := "8c190e67-87aa-4ed9-a90b-d9d5344523fb" // string | The role insights request id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsightsRequests(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsightsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsightsRequests`: RoleInsightsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsightsRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The role insights request id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleInsightsResponse**](RoleInsightsResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleInsightsSummary

> RoleInsightsSummary GetRoleInsightsSummary(ctx).Execute()

Get role insights summary information



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
    resp, r, err := apiClient.RoleInsightsApi.GetRoleInsightsSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleInsightsApi.GetRoleInsightsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleInsightsSummary`: RoleInsightsSummary
    fmt.Fprintf(os.Stdout, "Response from `RoleInsightsApi.GetRoleInsightsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleInsightsSummaryRequest struct via the builder pattern


### Return type

[**RoleInsightsSummary**](RoleInsightsSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

