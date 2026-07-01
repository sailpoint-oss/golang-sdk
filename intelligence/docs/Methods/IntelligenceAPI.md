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

Resolves exactly one identity by SCIM-style filters expression and returns the Intelligence envelope.
Supported queryable fields are id and email only.
The response embeds the first page of accounts, rare access, access-history access items, and
access-history certifications. Paged slices include a next link only when more results exist.
The privilegedAccess slice contains the full result and is not paged.
The outliers slice is omitted when the tenant lacks the IDA-outliers license.

A single match returns HTTP 200 with IntelIdentityAggregate.

Zero matches returns HTTP 404 with detailCode IDC_IDENTITY_NOT_FOUND.

Multiple matches returns HTTP 409 with detailCode IDC_IDENTITY_AMBIGUOUS and candidates listing each match.


[API Spec](https://developer.sailpoint.com/docs/api/get-identity-intelligence-v-1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityIntelligenceV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq* | 

### Return type

[**Intelidentityaggregate**](../models/intelidentityaggregate)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/intelligence"
)

func main() {
    filters := `id eq "ef38f94347e94562b5bb8424a56397d8"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq* # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq*

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIdentityIntelligenceV1(context.Background()).Filters(filters).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIdentityIntelligenceV1(context.Background()).Filters(filters).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIdentityIntelligenceV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityIntelligenceV1`: Intelidentityaggregate
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIdentityIntelligenceV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-access-item-history-v1
List identity access item history
Continuation endpoint for the parent response's `accessHistory.accessItems.next` link.
Returns one page of access-item history events for the supplied limit and offset values.
Unsupported event types and per-record decode failures are dropped server-side.
Requires tenant license idn:response-and-remediation.


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

### Return type

[**[]Intelaccessitemhistoryevent**](../models/intelaccessitemhistoryevent)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/intelligence"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccessItemHistoryV1`: []Intelaccessitemhistoryevent
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityAccessItemHistoryV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-accounts-v1
List identity accounts
Continuation endpoint for the parent response's `accounts.next` link.
Returns one page of account rows for the supplied limit and offset values.
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

### Return type

[**[]Intelaccessaccountwire**](../models/intelaccessaccountwire)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/intelligence"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccountsV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityAccountsV1(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityAccountsV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccountsV1`: []Intelaccessaccountwire
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityAccountsV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-certification-history-v1
List identity certification history
Continuation endpoint for the parent response's `accessHistory.certifications.next` link.
Returns one page of certification history events for the supplied limit and offset values.
Per-record decode failures are dropped server-side.
Requires tenant license idn:response-and-remediation.


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

### Return type

[**[]Intelcertificationhistoryevent**](../models/intelcertificationhistoryevent)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/intelligence"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityCertificationHistoryV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityCertificationHistoryV1(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityCertificationHistoryV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityCertificationHistoryV1`: []Intelcertificationhistoryevent
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityCertificationHistoryV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-rare-access-v1
List identity rare access
Continuation endpoint for the parent response's `outliers.rareAccess.next` link.
Resolves the identity's first outlier, then returns one page of rare access
items for the supplied limit and offset values. An identity with no outlier
returns an empty array. Requires tenant license idn:response-and-remediation
and the IDA-outliers license.


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

### Return type

[**[]Inteloutlieraccessitem**](../models/inteloutlieraccessitem)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/intelligence"
)

func main() {
    id := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence sub-resources. # string | Non-empty identity id path segment for Intelligence sub-resources.
    limit := 250 // int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250) # int32 | Page size. Defaults to 250; values above 250 are rejected with 400. (optional) (default to 250)
    offset := 0 // int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0) # int32 | Zero-based page offset. Defaults to 0. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityRareAccessV1(context.Background(), id).Execute()
	  //resp, r, err := apiClient.IntelligenceAPI.GetIntelIdentityRareAccessV1(context.Background(), id).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligenceAPI.GetIntelIdentityRareAccessV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityRareAccessV1`: []Inteloutlieraccessitem
    fmt.Fprintf(os.Stdout, "Response from `IntelligenceAPI.GetIntelIdentityRareAccessV1`: %v\n", resp)
}
```

[[Back to top]](#)

