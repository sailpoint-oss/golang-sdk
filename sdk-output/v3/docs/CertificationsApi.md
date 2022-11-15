# \CertificationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdentityCertification**](CertificationsApi.md#GetIdentityCertification) | **Get** /certifications/{id} | Identity Certification by ID
[**ListIdentityAccessReviewItems**](CertificationsApi.md#ListIdentityAccessReviewItems) | **Get** /certifications/{id}/access-review-items | List of Access Review Items
[**ListIdentityCertifications**](CertificationsApi.md#ListIdentityCertifications) | **Get** /certifications | Identity Campaign Certifications by IDs
[**MakeIdentityDecision**](CertificationsApi.md#MakeIdentityDecision) | **Post** /certifications/{id}/decide | Decide on a Certification Item
[**ReassignIdentityCertifications**](CertificationsApi.md#ReassignIdentityCertifications) | **Post** /certifications/{id}/reassign | Reassign Identities or Items
[**SignOffIdentityCertification**](CertificationsApi.md#SignOffIdentityCertification) | **Post** /certifications/{id}/sign-off | Finalize Identity Certification Decisions



## GetIdentityCertification

> IdentityCertificationDto GetIdentityCertification(ctx, id).Execute()

Identity Certification by ID



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The certification id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.GetIdentityCertification(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.GetIdentityCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityCertification`: IdentityCertificationDto
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.GetIdentityCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The certification id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityCertificationDto**](IdentityCertificationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityAccessReviewItems

> []AccessReviewItem ListIdentityAccessReviewItems(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Entitlements(entitlements).AccessProfiles(accessProfiles).Roles(roles).Execute()

List of Access Review Items



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **type / access.type**: *eq*  **completed**: *eq, ne*  **identitySummary.id**: *eq, in*  **identitySummary.name**: *eq, sw*  **access.id**: *eq, in*  **access.name**: *eq, sw*  **entitlement.sourceName**: *eq, sw*  **accessProfile.sourceName**: *eq, sw* (optional)
    sorters := "access.name,-accessProfile.sourceName" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **identitySummary.name, access.name, access.type, entitlement.sourceName, accessProfile.sourceName** (optional)
    entitlements := "identityEntitlement" // string | Filter results to view access review items that pertain to any of the specified comma-separated entitlement IDs.  An error will occur if this param is used with **access-profiles** or **roles** as only one of these query params can be used at a time. (optional)
    accessProfiles := "accessProfile1" // string | Filter results to view access review items that pertain to any of the specified comma-separated access-profle IDs.  An error will occur if this param is used with **entitlements** or **roles** as only one of these query params can be used at a time. (optional)
    roles := "userRole" // string | Filter results to view access review items that pertain to any of the specified comma-separated role IDs.  An error will occur if this param is used with **entitlements** or **access-profiles** as only one of these query params can be used at a time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.ListIdentityAccessReviewItems(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Entitlements(entitlements).AccessProfiles(accessProfiles).Roles(roles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.ListIdentityAccessReviewItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityAccessReviewItems`: []AccessReviewItem
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.ListIdentityAccessReviewItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityAccessReviewItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **type / access.type**: *eq*  **completed**: *eq, ne*  **identitySummary.id**: *eq, in*  **identitySummary.name**: *eq, sw*  **access.id**: *eq, in*  **access.name**: *eq, sw*  **entitlement.sourceName**: *eq, sw*  **accessProfile.sourceName**: *eq, sw* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **identitySummary.name, access.name, access.type, entitlement.sourceName, accessProfile.sourceName** | 
 **entitlements** | **string** | Filter results to view access review items that pertain to any of the specified comma-separated entitlement IDs.  An error will occur if this param is used with **access-profiles** or **roles** as only one of these query params can be used at a time. | 
 **accessProfiles** | **string** | Filter results to view access review items that pertain to any of the specified comma-separated access-profle IDs.  An error will occur if this param is used with **entitlements** or **roles** as only one of these query params can be used at a time. | 
 **roles** | **string** | Filter results to view access review items that pertain to any of the specified comma-separated role IDs.  An error will occur if this param is used with **entitlements** or **access-profiles** as only one of these query params can be used at a time. | 

### Return type

[**[]AccessReviewItem**](AccessReviewItem.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdentityCertifications

> []IdentityCertificationDto ListIdentityCertifications(ctx).ReviewerIdentity(reviewerIdentity).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Identity Campaign Certifications by IDs



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
    reviewerIdentity := "me" // string | The ID of reviewer identity. *me* indicates the current user. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "id eq "ef38f94347e94562b5bb8424a56397d8"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **id**: *eq, in* **campaign.id**: *eq, in* **phase**: *eq* **completed**: *eq, ne* (optional)
    sorters := "name,due" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, due, signed** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.ListIdentityCertifications(context.Background()).ReviewerIdentity(reviewerIdentity).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.ListIdentityCertifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIdentityCertifications`: []IdentityCertificationDto
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.ListIdentityCertifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdentityCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reviewerIdentity** | **string** | The ID of reviewer identity. *me* indicates the current user. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results) Filtering is supported for the following fields and operators: **id**: *eq, in* **campaign.id**: *eq, in* **phase**: *eq* **completed**: *eq, ne* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results) Sorting is supported for the following fields: **name, due, signed** | 

### Return type

[**[]IdentityCertificationDto**](IdentityCertificationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MakeIdentityDecision

> IdentityCertificationDto MakeIdentityDecision(ctx, id).ReviewDecision(reviewDecision).Execute()

Decide on a Certification Item



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The ID of the identity campaign certification on which to make decisions
    reviewDecision := []openapiclient.ReviewDecision{*openapiclient.NewReviewDecision("ef38f94347e94562b5bb8424a56397d8", openapiclient.CertificationDecision("APPROVE"), true)} // []ReviewDecision | A non-empty array of decisions to be made.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.MakeIdentityDecision(context.Background(), id).ReviewDecision(reviewDecision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.MakeIdentityDecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MakeIdentityDecision`: IdentityCertificationDto
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.MakeIdentityDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the identity campaign certification on which to make decisions | 

### Other Parameters

Other parameters are passed through a pointer to a apiMakeIdentityDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewDecision** | [**[]ReviewDecision**](ReviewDecision.md) | A non-empty array of decisions to be made. | 

### Return type

[**IdentityCertificationDto**](IdentityCertificationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignIdentityCertifications

> IdentityCertificationDto ReassignIdentityCertifications(ctx, id).ReviewReassign(reviewReassign).Execute()

Reassign Identities or Items



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
    id := "id_example" // string | The identity campaign certification ID
    reviewReassign := *openapiclient.NewReviewReassign([]openapiclient.ReassignReference{*openapiclient.NewReassignReference("ef38f94347e94562b5bb8424a56397d8", "ITEM")}, "ef38f94347e94562b5bb8424a56397d8", "reassigned for some reason") // ReviewReassign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.ReassignIdentityCertifications(context.Background(), id).ReviewReassign(reviewReassign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.ReassignIdentityCertifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReassignIdentityCertifications`: IdentityCertificationDto
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.ReassignIdentityCertifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignIdentityCertificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewReassign** | [**ReviewReassign**](ReviewReassign.md) |  | 

### Return type

[**IdentityCertificationDto**](IdentityCertificationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignOffIdentityCertification

> IdentityCertificationDto SignOffIdentityCertification(ctx, id).Execute()

Finalize Identity Certification Decisions



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The identity campaign certification ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationsApi.SignOffIdentityCertification(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationsApi.SignOffIdentityCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignOffIdentityCertification`: IdentityCertificationDto
    fmt.Fprintf(os.Stdout, "Response from `CertificationsApi.SignOffIdentityCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identity campaign certification ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignOffIdentityCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentityCertificationDto**](IdentityCertificationDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

