---
id: v1-intelligence-package
title: IntelligencePackage
pagination_label: IntelligencePackage
sidebar_label: IntelligencePackage
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IntelligencePackage', 'V1IntelligencePackage'] 
slug: /tools/sdk/go/intelligencepackage/methods/intelligence-package
tags: ['SDK', 'Software Development Kit', 'IntelligencePackage', 'V1IntelligencePackage']
---

# IntelligencePackageAPI
  Read-only HTTP API that returns the Intelligence Package (identity context)
for SecOps enrichment use cases (SIEM/SOAR connectors, MCP, browser
extension). Backed by Atlas internal-REST calls to MICE, Shelby List Accounts,
SDS Search, IDA-outliers, and identity-history.
 
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-intel-identity-access-history-v1**](#get-intel-identity-access-history-v1) | **Get** `/intelligence/v1/identities/{identityID}/access-history` | Return identity access-history events
[**get-intel-identity-access-v1**](#get-intel-identity-access-v1) | **Get** `/intelligence/v1/identities/{identityID}/access` | Accounts merged with privileged data
[**get-intel-identity-risk-outliers-v1**](#get-intel-identity-risk-outliers-v1) | **Get** `/intelligence/v1/identities/{identityID}/risk/outliers` | Risk outliers continuation paging
[**get-intel-identity-risk-v1**](#get-intel-identity-risk-v1) | **Get** `/intelligence/v1/identities/{identityID}/risk` | Identity risk snapshot
[**search-intel-identities-v1**](#search-intel-identities-v1) | **Get** `/intelligence/v1/identities` | Resolve one identity by filter


## get-intel-identity-access-history-v1
Return identity access-history events
Requires tenant license idn:response-and-remediation.

Events are relayed from identity-history; optional limit, offset, and count
are forwarded to the upstream when supplied.


[API Spec](https://developer.sailpoint.com/docs/api/v1/get-intel-identity-access-history-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityID** | **string** | Non-empty identity id path segment for Intelligence Package sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityAccessHistoryV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of historical events to return in this page of results. | 
 **offset** | **int32** | Zero-based index of the first event row to return for pagination. | 
 **count** | **bool** | When true, the service may include total count metadata alongside the result list. | [default to false]

### Return type

[**Intelidentityaccesshistorybody**](../models/intelidentityaccesshistorybody)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_intelligence_package"
)

func main() {
    identityID := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence Package sub-resources. # string | Non-empty identity id path segment for Intelligence Package sub-resources.
    limit := 100 // int32 | Maximum number of historical events to return in this page of results. (optional) # int32 | Maximum number of historical events to return in this page of results. (optional)
    offset := 0 // int32 | Zero-based index of the first event row to return for pagination. (optional) # int32 | Zero-based index of the first event row to return for pagination. (optional)
    count := false // bool | When true, the service may include total count metadata alongside the result list. (optional) (default to false) # bool | When true, the service may include total count metadata alongside the result list. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityAccessHistoryV1(context.Background(), identityID).Execute()
	  //resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityAccessHistoryV1(context.Background(), identityID).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligencePackageAPI.GetIntelIdentityAccessHistoryV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccessHistoryV1`: Intelidentityaccesshistorybody
    fmt.Fprintf(os.Stdout, "Response from `IntelligencePackageAPI.GetIntelIdentityAccessHistoryV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-access-v1
Accounts merged with privileged data
Requires tenant license idn:response-and-remediation.

Client-facing pagination (limit, offset, count) is not supported on this route.
The service issues one Shelby List Accounts request at the upstream maximum page size
and one SDS Search request per identity.


[API Spec](https://developer.sailpoint.com/docs/api/v1/get-intel-identity-access-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityID** | **string** | Non-empty identity id path segment for Intelligence Package sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityAccessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Intelidentityaccessbody**](../models/intelidentityaccessbody)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_intelligence_package"
)

func main() {
    identityID := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence Package sub-resources. # string | Non-empty identity id path segment for Intelligence Package sub-resources.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityAccessV1(context.Background(), identityID).Execute()
	  //resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityAccessV1(context.Background(), identityID).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligencePackageAPI.GetIntelIdentityAccessV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityAccessV1`: Intelidentityaccessbody
    fmt.Fprintf(os.Stdout, "Response from `IntelligencePackageAPI.GetIntelIdentityAccessV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-risk-outliers-v1
Risk outliers continuation paging
Continuation endpoint for risk outlier access-items. Returns one page based on
the supplied limit and offset values and includes an optional continuation link
when more rows remain. Requires tenant license idn:response-and-remediation.


[API Spec](https://developer.sailpoint.com/docs/api/v1/get-intel-identity-risk-outliers-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityID** | **string** | Non-empty identity id path segment for Intelligence Package sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityRiskOutliersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of outlier rows to return for this page. | [default to 250]
 **offset** | **int32** | Zero-based row index for the first returned outlier item. | [default to 0]

### Return type

[**Intelidentityriskbody**](../models/intelidentityriskbody)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_intelligence_package"
)

func main() {
    identityID := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence Package sub-resources. # string | Non-empty identity id path segment for Intelligence Package sub-resources.
    limit := 250 // int32 | Maximum number of outlier rows to return for this page. (optional) (default to 250) # int32 | Maximum number of outlier rows to return for this page. (optional) (default to 250)
    offset := 0 // int32 | Zero-based row index for the first returned outlier item. (optional) (default to 0) # int32 | Zero-based row index for the first returned outlier item. (optional) (default to 0)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityRiskOutliersV1(context.Background(), identityID).Execute()
	  //resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityRiskOutliersV1(context.Background(), identityID).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligencePackageAPI.GetIntelIdentityRiskOutliersV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityRiskOutliersV1`: Intelidentityriskbody
    fmt.Fprintf(os.Stdout, "Response from `IntelligencePackageAPI.GetIntelIdentityRiskOutliersV1`: %v\n", resp)
}
```

[[Back to top]](#)

## get-intel-identity-risk-v1
Identity risk snapshot
Risk snapshot envelope for the identity. The service resolves the first matching
outlier for identityID and returns one page of access-items plus an optional
continuation link for additional pages.

Clients should continue paging using _links.outliers.href when provided.
Requires tenant license idn:response-and-remediation.


[API Spec](https://developer.sailpoint.com/docs/api/v1/get-intel-identity-risk-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityID** | **string** | Non-empty identity id path segment for Intelligence Package sub-resources. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntelIdentityRiskV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Intelidentityriskbody**](../models/intelidentityriskbody)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_intelligence_package"
)

func main() {
    identityID := `ef38f94347e94562b5bb8424a56397d8` // string | Non-empty identity id path segment for Intelligence Package sub-resources. # string | Non-empty identity id path segment for Intelligence Package sub-resources.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityRiskV1(context.Background(), identityID).Execute()
	  //resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.GetIntelIdentityRiskV1(context.Background(), identityID).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligencePackageAPI.GetIntelIdentityRiskV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIntelIdentityRiskV1`: Intelidentityriskbody
    fmt.Fprintf(os.Stdout, "Response from `IntelligencePackageAPI.GetIntelIdentityRiskV1`: %v\n", resp)
}
```

[[Back to top]](#)

## search-intel-identities-v1
Resolve one identity by filter
Requires tenant license idn:response-and-remediation.

Returns the Intelligence Package envelope for the identity that matches the SCIM-style filters expression.
Supported queryable fields are id and email only.

A single match returns HTTP 200 with IntelIdentityResponse.

Zero matches returns HTTP 404 with detailCode IDC_IDENTITY_NOT_FOUND. 

Multiple matches returns HTTP 409 with detailCode IDC_IDENTITY_AMBIGUOUS and candidates listing each match.


[API Spec](https://developer.sailpoint.com/docs/api/v1/search-intel-identities-v1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIntelIdentitiesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq* | 

### Return type

[**Intelidentityresponse**](../models/intelidentityresponse)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3/api_intelligence_package"
)

func main() {
    filters := `id eq "00000000-0000-0000-0000-000000000000"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq* # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq*  **email**: *eq*

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.SearchIntelIdentitiesV1(context.Background()).Filters(filters).Execute()
	  //resp, r, err := apiClient.IntelligencePackage.IntelligencePackageAPI.SearchIntelIdentitiesV1(context.Background()).Filters(filters).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IntelligencePackageAPI.SearchIntelIdentitiesV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIntelIdentitiesV1`: Intelidentityresponse
    fmt.Fprintf(os.Stdout, "Response from `IntelligencePackageAPI.SearchIntelIdentitiesV1`: %v\n", resp)
}
```

[[Back to top]](#)

