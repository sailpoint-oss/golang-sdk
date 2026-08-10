---
id: v1-intelligence
title: Intelligence
pagination_label: Intelligence
sidebar_label: Intelligence
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Intelligence', 'V1Intelligence'] 
slug: /tools/sdk/go/intelligence/methods/intelligence
tags: ['SDK', 'Software Development Kit', 'Intelligence', 'V1Intelligence']
---

# IntelligenceAPI
  Read-only HTTP API that returns the Intelligence (identity context)
for SecOps enrichment use cases (SIEM/SOAR connectors, MCP, browser
extension). Backed by Atlas internal-REST calls to MICE, Shelby List Accounts,
SDS Search, IDA-outliers, and identity-history.

## Pagination

The aggregated Human GET embeds the first **10** items per paged slice. Each upstream paged call
sends &#x60;count&#x3D;true&#x60; and reads &#x60;X-Total-Count&#x60;. Parent slices expose &#x60;totalCount&#x60; when &#x60;items&#x60; is
non-empty and set &#x60;next&#x60; when &#x60;totalCount &gt; offset + len(items)&#x60; (aggregate offset is always 0).
Empty slices render as &#x60;items: []&#x60; with no &#x60;totalCount&#x60;. &#x60;privilegedAccess&#x60; is never paged and
carries no &#x60;totalCount&#x60;.

Human child routes (&#x60;/accounts&#x60;, &#x60;/outliers/rare-access&#x60;, &#x60;/access-history/*&#x60;) follow the
SailPoint V3 pattern: pass &#x60;count&#x3D;true&#x60; to receive &#x60;X-Total-Count&#x60; (including &#x60;0&#x60; on empty
pages). When &#x60;count&#x60; is omitted, upstream count work is skipped and the header is omitted.
 
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-identity-intelligence-v1**](#get-identity-intelligence-v1) | **Get** `/intelligence/v1/identities` | Get identity by filter
[**get-intel-identity-access-item-history-v1**](#get-intel-identity-access-item-history-v1) | **Get** `/intelligence/v1/identities/{id}/access-history/access-items` | List identity access item history
[**get-intel-identity-accounts-v1**](#get-intel-identity-accounts-v1) | **Get** `/intelligence/v1/identities/{id}/accounts` | List identity accounts
[**get-intel-identity-certification-history-v1**](#get-intel-identity-certification-history-v1) | **Get** `/intelligence/v1/identities/{id}/access-history/certifications` | List identity certification history
[**get-intel-identity-rare-access-v1**](#get-intel-identity-rare-access-v1) | **Get** `/intelligence/v1/identities/{id}/outliers/rare-access` | List identity rare access


## get-identity-intelligence-v1
Get identity by filter
Requires tenant license idn:response-and-remediation.

**Authentication and data segmentation**

Intelligence forwards the caller JWT to downstream identity and search services (context client).
Enriched results, including non-human identity resolution, are filtered to the caller's Data
Segmentation visibility.

**Caution:** Generic API Management API keys are not tied to a user identity. When Data
Segmentation is enabled, API key authentication may fail or return incomplete data because
downstream calls require a user context. Use a [personal access token](https://developer.sailpoint.com/docs/api/authentication/#generate-a-personal-access-token)
or other user-scoped OAuth token. See [API keys](https://documentation.sailpoint.com/saas/help/common/api_keys.html)
and [Data Segmentation](https://documentation.sailpoint.com/saas/help/segmentation/index.html).

Resolves exactly one identity using a single SCIM-style filters expression.

**Supported filters**

| Filter field | Lookup mode | Notes |
|---|---|---|
| id eq | Human (+ optional non-human identity when feature-flagged) | Resolves human identities by id; when non-human resolution is enabled, a parallel non-human lookup runs. If both match different identities, returns HTTP 409. |
| email eq | Human only | Human identity lookup by email only. |
| opaqueIdentifier eq | Non-human identity only | Parallel nativeIdentity eq on machine-identities and machine-accounts, then name-prefix fallback on machine-accounts. Requires feature flag ISCRR-1905_NHI_TYPE_MACHINE_FILTER_ENABLED; when disabled, returns HTTP 400. |

Single-clause filters only; composite and or expressions are rejected with HTTP 400.

**Human envelope (type Human)**

Embeds the first page (10 items) of each enrichment slice. Each paged slice includes totalCount
from upstream X-Total-Count when items is non-empty, and carries a next continuation URL when
totalCount exceeds the items returned on this page. Slices are always present (empty uses
items [] with no totalCount). privilegedAccess returns the full privileged-access result and never carries
next or totalCount. If any enrichment upstream fails, the whole request fails with HTTP 500,
except outliers, which is omitted (not an error) when the tenant lacks the IDA-outliers license
(upstream 401 or 403). identityGraph is omitted when the tenant lacks the idg:base license.

**Non-human identity envelope (type NHI)**

Returns flat non-human identity fields at the top level plus correlated machine accounts on the
aggregate and a derived block (isOrphaned, authorizedHumanIdentities, blastRadiusSummary).
Omits Human-only slices (privilegedAccess, outliers, accessHistory). Account paging via child
routes is not yet released. Opaque prefix resolution that deduplicates to one parent identity
returns HTTP 200 with matchConfidence partial; multiple distinct parent identities return HTTP 409
with IDC_IDENTITY_AMBIGUOUS and candidate id and displayName values. identityGraph is omitted
when the tenant lacks the idg:base license.


[API Spec](https://developer.sailpoint.com/docs/api/get-identity-intelligence-v-1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityIntelligenceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq*  **opaqueIdentifier**: *eq* | 

### Return type

[**Intelidentityenvelope**](../models/intelidentityenvelope)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    filters := `id eq "ef38f94347e94562b5bb8424a56397d8"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq*  **opaqueIdentifier**: *eq* # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq*  **opaqueIdentifier**: *eq*

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIdentityIntelligenceV1(context.Background()).Filters(filters).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIdentityIntelligenceV1(context.Background()).Filters(filters).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIdentityIntelligenceV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityIntelligenceV1`: Intelidentityenvelope
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIdentityIntelligenceV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-access-item-history-v1
List identity access item history
Continuation endpoint for the parent response's `accessHistory.accessItems.next` link.
Returns one page of access-item history events for the supplied limit and offset values.
Pass `count=true` to receive `X-Total-Count` (including `0` on empty pages).
Unsupported event types and per-record decode failures are dropped server-side.
Requires tenant license idn:response-and-remediation.

Not applicable to non-human identities.


[API Spec](https://developer.sailpoint.com/docs/api/get-intel-identity-access-item-history-v-1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-empty identity id path segment for Intelligence sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityAccessItemHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size. Defaults to 250; values above 250 are rejected with 400. | [default to 250]
 **offset** | **int32** | Zero-based page offset. Defaults to 0. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IntelAccessItemHistoryEvent**](../models/intel-access-item-history-event)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccessItemHistoryV1`: []IntelAccessItemHistoryEvent
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-accounts-v1
List identity accounts
Continuation endpoint for a Human identity's `accounts.next` link.
Returns one page of account rows for the supplied limit and offset values.
Pass `count=true` to receive `X-Total-Count` (including `0` on empty pages).
Not applicable to non-human identities (NHI accounts are returned on the NHI aggregate only).
Requires tenant license idn:response-and-remediation.


[API Spec](https://developer.sailpoint.com/docs/api/get-intel-identity-accounts-v-1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-empty identity id path segment for Intelligence sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityAccountsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size. Defaults to 250; values above 250 are rejected with 400. | [default to 250]
 **offset** | **int32** | Zero-based page offset. Defaults to 0. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IntelAccessAccountWire**](../models/intel-access-account-wire)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccountsV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccountsV1(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityAccountsV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccountsV1`: []IntelAccessAccountWire
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityAccountsV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-certification-history-v1
List identity certification history
Continuation endpoint for the parent response's `accessHistory.certifications.next` link.
Returns one page of certification history events for the supplied limit and offset values.
Pass `count=true` to receive `X-Total-Count` (including `0` on empty pages).
Per-record decode failures are dropped server-side.
Requires tenant license idn:response-and-remediation.

Not applicable to non-human identities.


[API Spec](https://developer.sailpoint.com/docs/api/get-intel-identity-certification-history-v-1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-empty identity id path segment for Intelligence sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityCertificationHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size. Defaults to 250; values above 250 are rejected with 400. | [default to 250]
 **offset** | **int32** | Zero-based page offset. Defaults to 0. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IntelCertificationHistoryEvent**](../models/intel-certification-history-event)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityCertificationHistoryV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityCertificationHistoryV1(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityCertificationHistoryV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityCertificationHistoryV1`: []IntelCertificationHistoryEvent
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityCertificationHistoryV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-rare-access-v1
List identity rare access
Continuation endpoint for the parent response's `outliers.rareAccess.next` link.
Resolves the identity's first outlier, then returns one page of rare access
items for the supplied limit and offset values. Pass `count=true` to receive
`X-Total-Count` (including `0` on empty pages). An identity with no outlier
returns an empty array with `X-Total-Count: 0` when `count=true`. Requires
tenant license idn:response-and-remediation and the IDA-outliers license.

Not applicable to non-human identities (no outliers slice on the NHI envelope).


[API Spec](https://developer.sailpoint.com/docs/api/get-intel-identity-rare-access-v-1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Non-empty identity id path segment for Intelligence sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityRareAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Page size. Defaults to 250; values above 250 are rejected with 400. | [default to 250]
 **offset** | **int32** | Zero-based page offset. Defaults to 0. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]IntelOutlierAccessItem**](../models/intel-outlier-access-item)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityRareAccessV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityRareAccessV1(context.Background(), id).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityRareAccessV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityRareAccessV1`: []IntelOutlierAccessItem
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityRareAccessV1`: %v\n", resp)
}
```

[[Back to top]](#)

