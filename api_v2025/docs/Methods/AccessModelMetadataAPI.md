---
id: v2025-access-model-metadata
title: AccessModelMetadata
pagination_label: AccessModelMetadata
sidebar_label: AccessModelMetadata
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccessModelMetadata', 'V2025AccessModelMetadata'] 
slug: /tools/sdk/go/v2025/methods/access-model-metadata
tags: ['SDK', 'Software Development Kit', 'AccessModelMetadata', 'V2025AccessModelMetadata']
---

# AccessModelMetadataAPI
  Use this API to create and manage metadata attributes for your Access Model.
Access Model Metadata allows you to add contextual information to your ISC Access Model items using pre-defined metadata for risk, regulations, privacy levels, etc., or by creating your own metadata attributes to reflect the unique needs of your organization. This release of the API includes support for entitlement metadata. Support for role and access profile metadata will be introduced in a subsequent release.

Common usages for Access Model metadata include:

- Organizing and categorizing access items to make it easier for your users to search for and find the access rights they want to request, certify, or manage.

- Providing richer information about access that is being acted on to allow stakeholders to make better decisions when approving, certifying, or managing access rights.

- Identifying access that may requires additional approval requirements or be subject to more frequent review.
 
All URIs are relative to *https://sailpoint.api.identitynow.com/v2025*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-access-model-metadata-attribute**](#get-access-model-metadata-attribute) | **Get** `/access-model-metadata/attributes/{key}` | Get access model metadata attribute
[**get-access-model-metadata-attribute-value**](#get-access-model-metadata-attribute-value) | **Get** `/access-model-metadata/attributes/{key}/values/{value}` | Get access model metadata value
[**list-access-model-metadata-attribute**](#list-access-model-metadata-attribute) | **Get** `/access-model-metadata/attributes` | List access model metadata attributes
[**list-access-model-metadata-attribute-value**](#list-access-model-metadata-attribute-value) | **Get** `/access-model-metadata/attributes/{key}/values` | List access model metadata values
[**update-access-model-metadata-by-filter**](#update-access-model-metadata-by-filter) | **Post** `/access-model-metadata/bulk-update/filter` | Metadata Attribute update by filter
[**update-access-model-metadata-by-ids**](#update-access-model-metadata-by-ids) | **Post** `/access-model-metadata/bulk-update/ids` | Metadata Attribute update by ids
[**update-access-model-metadata-by-query**](#update-access-model-metadata-by-query) | **Post** `/access-model-metadata/bulk-update/query` | Metadata Attribute update by query


## get-access-model-metadata-attribute
Get access model metadata attribute
Get single Access Model Metadata Attribute

[API Spec](https://developer.sailpoint.com/docs/api/v2025/get-access-model-metadata-attribute)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessModelMetadataAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttributeDTO**](../models/attribute-dto)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    key := `iscPrivacy` // string | Technical name of the Attribute. # string | Technical name of the Attribute.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.GetAccessModelMetadataAttribute(context.Background(), key).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.GetAccessModelMetadataAttribute(context.Background(), key).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.GetAccessModelMetadataAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessModelMetadataAttribute`: AttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.GetAccessModelMetadataAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## get-access-model-metadata-attribute-value
Get access model metadata value
Get single Access Model Metadata Attribute Value

[API Spec](https://developer.sailpoint.com/docs/api/v2025/get-access-model-metadata-attribute-value)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 
**value** | **string** | Technical name of the Attribute value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessModelMetadataAttributeValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AttributeValueDTO**](../models/attribute-value-dto)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    key := `iscPrivacy` // string | Technical name of the Attribute. # string | Technical name of the Attribute.
    value := `public` // string | Technical name of the Attribute value. # string | Technical name of the Attribute value.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue(context.Background(), key, value).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue(context.Background(), key, value).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessModelMetadataAttributeValue`: AttributeValueDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.GetAccessModelMetadataAttributeValue`: %v\n", resp)
}
```

[[Back to top]](#)

## list-access-model-metadata-attribute
List access model metadata attributes
Get a list of Access Model Metadata Attributes

[API Spec](https://developer.sailpoint.com/docs/api/v2025/list-access-model-metadata-attribute)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessModelMetadataAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **key**: *eq*  **name**: *eq*  **type**: *eq*  **status**: *eq*  **objectTypes**: *eq*  **Supported composite operators**: *and* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, key** | 
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]AttributeDTO**](../models/attribute-dto)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    filters := `name eq "Privacy"` // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **key**: *eq*  **name**: *eq*  **type**: *eq*  **status**: *eq*  **objectTypes**: *eq*  **Supported composite operators**: *and* (optional) # string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **key**: *eq*  **name**: *eq*  **type**: *eq*  **status**: *eq*  **objectTypes**: *eq*  **Supported composite operators**: *and* (optional)
    sorters := `name,-key` // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, key** (optional) # string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, key** (optional)
    offset := 0 // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0) # int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := 250 // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250) # int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.ListAccessModelMetadataAttribute(context.Background()).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.ListAccessModelMetadataAttribute(context.Background()).Filters(filters).Sorters(sorters).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.ListAccessModelMetadataAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessModelMetadataAttribute`: []AttributeDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.ListAccessModelMetadataAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## list-access-model-metadata-attribute-value
List access model metadata values
Get a list of Access Model Metadata Attribute Values

[API Spec](https://developer.sailpoint.com/docs/api/v2025/list-access-model-metadata-attribute-value)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Technical name of the Attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessModelMetadataAttributeValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]AttributeValueDTO**](../models/attribute-value-dto)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    key := `iscPrivacy` // string | Technical name of the Attribute. # string | Technical name of the Attribute.
    offset := 0 // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0) # int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    limit := 250 // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250) # int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false) # bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue(context.Background(), key).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue(context.Background(), key).Offset(offset).Limit(limit).Count(count).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessModelMetadataAttributeValue`: []AttributeValueDTO
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.ListAccessModelMetadataAttributeValue`: %v\n", resp)
}
```

[[Back to top]](#)

## update-access-model-metadata-by-filter
Metadata Attribute update by filter
Bulk update Access Model Metadata Attribute Values using a filter

[API Spec](https://developer.sailpoint.com/docs/api/v2025/update-access-model-metadata-by-filter)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessModelMetadataByFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementAttributeBulkUpdateFilterRequest** | [**EntitlementAttributeBulkUpdateFilterRequest**](../models/entitlement-attribute-bulk-update-filter-request) | Attribute metadata bulk update request body. | 

### Return type

[**AccessModelMetadataBulkUpdateResponse**](../models/access-model-metadata-bulk-update-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v2025 "github.com/sailpoint-oss/golang-sdk/v2/api_v2025"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    entitlementattributebulkupdatefilterrequest := []byte(`{
          "values" : [ {
            "attribute" : "iscFederalClassifications",
            "values" : [ "topSecret" ]
          } ],
          "filters" : "id eq 2c9180867817ac4d017817c491119a20",
          "replaceScope" : "attribute",
          "operation" : "add"
        }`) // EntitlementAttributeBulkUpdateFilterRequest | Attribute metadata bulk update request body.

    var entitlementAttributeBulkUpdateFilterRequest v2025.EntitlementAttributeBulkUpdateFilterRequest
    if err := json.Unmarshal(entitlementattributebulkupdatefilterrequest, &entitlementAttributeBulkUpdateFilterRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByFilter(context.Background()).EntitlementAttributeBulkUpdateFilterRequest(entitlementAttributeBulkUpdateFilterRequest).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByFilter(context.Background()).EntitlementAttributeBulkUpdateFilterRequest(entitlementAttributeBulkUpdateFilterRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.UpdateAccessModelMetadataByFilter``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessModelMetadataByFilter`: AccessModelMetadataBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.UpdateAccessModelMetadataByFilter`: %v\n", resp)
}
```

[[Back to top]](#)

## update-access-model-metadata-by-ids
Metadata Attribute update by ids
Bulk update Access Model Metadata Attribute Values using ids.

[API Spec](https://developer.sailpoint.com/docs/api/v2025/update-access-model-metadata-by-ids)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessModelMetadataByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementAttributeBulkUpdateIdsRequest** | [**EntitlementAttributeBulkUpdateIdsRequest**](../models/entitlement-attribute-bulk-update-ids-request) | Attribute metadata bulk update request body. | 

### Return type

[**AccessModelMetadataBulkUpdateResponse**](../models/access-model-metadata-bulk-update-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v2025 "github.com/sailpoint-oss/golang-sdk/v2/api_v2025"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    entitlementattributebulkupdateidsrequest := []byte(`{
          "entitlements" : [ "2c9180867817ac4d017817c491119a20", "2c9180867817ac4d017817c491119a21" ],
          "values" : [ {
            "attribute" : "iscFederalClassifications",
            "values" : [ "topSecret" ]
          } ],
          "replaceScope" : "attribute",
          "operation" : "add"
        }`) // EntitlementAttributeBulkUpdateIdsRequest | Attribute metadata bulk update request body.

    var entitlementAttributeBulkUpdateIdsRequest v2025.EntitlementAttributeBulkUpdateIdsRequest
    if err := json.Unmarshal(entitlementattributebulkupdateidsrequest, &entitlementAttributeBulkUpdateIdsRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByIds(context.Background()).EntitlementAttributeBulkUpdateIdsRequest(entitlementAttributeBulkUpdateIdsRequest).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByIds(context.Background()).EntitlementAttributeBulkUpdateIdsRequest(entitlementAttributeBulkUpdateIdsRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.UpdateAccessModelMetadataByIds``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessModelMetadataByIds`: AccessModelMetadataBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.UpdateAccessModelMetadataByIds`: %v\n", resp)
}
```

[[Back to top]](#)

## update-access-model-metadata-by-query
Metadata Attribute update by query
Bulk update Access Model Metadata Attribute Values using a query

[API Spec](https://developer.sailpoint.com/docs/api/v2025/update-access-model-metadata-by-query)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessModelMetadataByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementAttributeBulkUpdateQueryRequest** | [**EntitlementAttributeBulkUpdateQueryRequest**](../models/entitlement-attribute-bulk-update-query-request) | Attribute metadata bulk update request body. | 

### Return type

[**AccessModelMetadataBulkUpdateResponse**](../models/access-model-metadata-bulk-update-response)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v2025 "github.com/sailpoint-oss/golang-sdk/v2/api_v2025"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    entitlementattributebulkupdatequeryrequest := []byte(`{
          "query" : {
            "queryDsl" : {
              "match" : {
                "name" : "john.doe"
              }
            },
            "aggregationType" : "DSL",
            "aggregationsVersion" : "",
            "query" : {
              "query" : "name:a*",
              "timeZone" : "America/Chicago",
              "fields" : "[firstName,lastName,email]",
              "innerHit" : {
                "query" : "source.name:\\\"Active Directory\\\"",
                "type" : "access"
              }
            },
            "aggregationsDsl" : { },
            "sort" : [ "displayName", "+id" ],
            "filters" : { },
            "queryVersion" : "",
            "queryType" : "SAILPOINT",
            "includeNested" : true,
            "queryResultFilter" : {
              "excludes" : [ "stacktrace" ],
              "includes" : [ "name", "displayName" ]
            },
            "indices" : [ "identities" ],
            "typeAheadQuery" : {
              "field" : "source.name",
              "size" : 100,
              "query" : "Work",
              "sortByValue" : true,
              "nestedType" : "access",
              "sort" : "asc",
              "maxExpansions" : 10
            },
            "textQuery" : {
              "contains" : true,
              "terms" : [ "The quick brown fox", "3141592", "7" ],
              "matchAny" : false,
              "fields" : [ "displayName", "employeeNumber", "roleCount" ]
            },
            "searchAfter" : [ "John Doe", "2c91808375d8e80a0175e1f88a575221" ],
            "aggregations" : {
              "filter" : {
                "field" : "access.type",
                "name" : "Entitlements",
                "type" : "TERM",
                "value" : "ENTITLEMENT"
              },
              "bucket" : {
                "field" : "attributes.city",
                "size" : 100,
                "minDocCount" : 2,
                "name" : "Identity Locations",
                "type" : "TERMS"
              },
              "metric" : {
                "field" : "@access.name",
                "name" : "Access Name Count",
                "type" : "COUNT"
              },
              "subAggregation" : {
                "filter" : {
                  "field" : "access.type",
                  "name" : "Entitlements",
                  "type" : "TERM",
                  "value" : "ENTITLEMENT"
                },
                "bucket" : {
                  "field" : "attributes.city",
                  "size" : 100,
                  "minDocCount" : 2,
                  "name" : "Identity Locations",
                  "type" : "TERMS"
                },
                "metric" : {
                  "field" : "@access.name",
                  "name" : "Access Name Count",
                  "type" : "COUNT"
                },
                "subAggregation" : {
                  "filter" : {
                    "field" : "access.type",
                    "name" : "Entitlements",
                    "type" : "TERM",
                    "value" : "ENTITLEMENT"
                  },
                  "bucket" : {
                    "field" : "attributes.city",
                    "size" : 100,
                    "minDocCount" : 2,
                    "name" : "Identity Locations",
                    "type" : "TERMS"
                  },
                  "metric" : {
                    "field" : "@access.name",
                    "name" : "Access Name Count",
                    "type" : "COUNT"
                  },
                  "nested" : {
                    "name" : "id",
                    "type" : "access"
                  }
                },
                "nested" : {
                  "name" : "id",
                  "type" : "access"
                }
              },
              "nested" : {
                "name" : "id",
                "type" : "access"
              }
            }
          },
          "values" : [ {
            "attribute" : "iscFederalClassifications",
            "values" : [ "topSecret" ]
          } ],
          "replaceScope" : "attribute",
          "operation" : "add"
        }`) // EntitlementAttributeBulkUpdateQueryRequest | Attribute metadata bulk update request body.

    var entitlementAttributeBulkUpdateQueryRequest v2025.EntitlementAttributeBulkUpdateQueryRequest
    if err := json.Unmarshal(entitlementattributebulkupdatequeryrequest, &entitlementAttributeBulkUpdateQueryRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByQuery(context.Background()).EntitlementAttributeBulkUpdateQueryRequest(entitlementAttributeBulkUpdateQueryRequest).Execute()
	  //resp, r, err := apiClient.V2025.AccessModelMetadataAPI.UpdateAccessModelMetadataByQuery(context.Background()).EntitlementAttributeBulkUpdateQueryRequest(entitlementAttributeBulkUpdateQueryRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessModelMetadataAPI.UpdateAccessModelMetadataByQuery``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessModelMetadataByQuery`: AccessModelMetadataBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessModelMetadataAPI.UpdateAccessModelMetadataByQuery`: %v\n", resp)
}
```

[[Back to top]](#)

