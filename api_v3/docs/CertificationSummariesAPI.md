# \CertificationSummariesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentityAccessSummaries**](CertificationSummariesAPI.md#GetIdentityAccessSummaries) | **Get** /certifications/{id}/access-summaries/{type} | Access Summaries
[**GetIdentityDecisionSummary**](CertificationSummariesAPI.md#GetIdentityDecisionSummary) | **Get** /certifications/{id}/decision-summary | Summary of Certification Decisions
[**GetIdentitySummaries**](CertificationSummariesAPI.md#GetIdentitySummaries) | **Get** /certifications/{id}/identity-summaries | Identity Summaries for Campaign Certification
[**GetIdentitySummary**](CertificationSummariesAPI.md#GetIdentitySummary) | **Get** /certifications/{id}/identity-summaries/{identitySummaryId} | Summary for Identity



## Access Summaries

> []AccessSummary GetIdentityAccessSummaries(ctx, id, type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

This API returns a list of access summaries for the specified identity campaign certification and type. A token with ORG_ADMIN or CERT_ADMIN authority is required to call this API. Reviewers for this certification can also call this API.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID
    type_ := "ACCESS_PROFILE" // string | The type of access review item to retrieve summaries for
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "access.id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **completed**: *eq, ne*  **access.id**: *eq, in*  **access.name**: *eq, sw*  **entitlement.sourceName**: *eq, sw*  **accessProfile.sourceName**: *eq, sw* (optional)
    sorters := "access.name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationSummariesAPI.GetIdentityAccessSummaries(context.Background(), id, type_).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationSummariesAPI.GetIdentityAccessSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityAccessSummaries`: []AccessSummary
    fmt.Fprintf(os.Stdout, "Response from `CertificationSummariesAPI.GetIdentityAccessSummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 
**type_** | **string** | The type of access review item to retrieve summaries for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityAccessSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **completed**: *eq, ne*  **access.id**: *eq, in*  **access.name**: *eq, sw*  **entitlement.sourceName**: *eq, sw*  **accessProfile.sourceName**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **access.name** | 

### Return type

[**[]AccessSummary**](AccessSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Summary of Certification Decisions

> IdentityCertDecisionSummary GetIdentityDecisionSummary(ctx, id).Filters(filters).Execute()

This API returns a summary of the decisions made on an identity campaign certification. The decisions are summarized by type. A token with ORG_ADMIN or CERT_ADMIN authority is required to call this API. Reviewers for this certification can also call this API.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The certification ID
    filters := "identitySummary.id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **identitySummary.id**: *eq, in* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationSummariesAPI.GetIdentityDecisionSummary(context.Background(), id).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationSummariesAPI.GetIdentityDecisionSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityDecisionSummary`: IdentityCertDecisionSummary
    fmt.Fprintf(os.Stdout, "Response from `CertificationSummariesAPI.GetIdentityDecisionSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityDecisionSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **identitySummary.id**: *eq, in* | 

### Return type

[**IdentityCertDecisionSummary**](IdentityCertDecisionSummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Identity Summaries for Campaign Certification

> []CertificationIdentitySummary GetIdentitySummaries(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

This API returns a list of the identity summaries for a specific identity campaign certification. A token with ORG_ADMIN or CERT_ADMIN authority is required to call this API. Reviewers for this certification can also call this API.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **completed**: *eq, ne*  **name**: *eq, sw* (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationSummariesAPI.GetIdentitySummaries(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationSummariesAPI.GetIdentitySummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitySummaries`: []CertificationIdentitySummary
    fmt.Fprintf(os.Stdout, "Response from `CertificationSummariesAPI.GetIdentitySummaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **completed**: *eq, ne*  **name**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name** | 

### Return type

[**[]CertificationIdentitySummary**](CertificationIdentitySummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Summary for Identity

> CertificationIdentitySummary GetIdentitySummary(ctx, id, identitySummaryId).Execute()

This API returns the summary for an identity on a specified identity campaign certification. A token with ORG_ADMIN or CERT_ADMIN authority is required to call this API. Reviewers for this certification can also call this API.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID
    identitySummaryId := "2c91808772a504f50172a9540e501ba8" // string | The identity summary ID

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationSummariesAPI.GetIdentitySummary(context.Background(), id, identitySummaryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationSummariesAPI.GetIdentitySummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentitySummary`: CertificationIdentitySummary
    fmt.Fprintf(os.Stdout, "Response from `CertificationSummariesAPI.GetIdentitySummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 
**identitySummaryId** | **string** | The identity summary ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentitySummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CertificationIdentitySummary**](CertificationIdentitySummary.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

