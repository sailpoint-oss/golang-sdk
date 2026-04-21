---
id: nerm-advanced-search
title: AdvancedSearch
pagination_label: AdvancedSearch
sidebar_label: AdvancedSearch
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AdvancedSearch', 'NERMAdvancedSearch'] 
slug: /tools/sdk/go/nerm/methods/advanced-search
tags: ['SDK', 'Software Development Kit', 'AdvancedSearch', 'NERMAdvancedSearch']
---

# AdvancedSearchAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-advanced-search**](#get-advanced-search) | **Get** `/advanced_search` | Get saved advanced search queries
[**patch-advanced-search**](#patch-advanced-search) | **Patch** `/advanced_search/{id}` | Update a saved advanced search
[**search-advanced-search**](#search-advanced-search) | **Post** `/advanced_search/run` | Run profiles advanced search
[**search-advanced-searchby-id**](#search-advanced-searchby-id) | **Get** `/advanced_search/{id}/run` | Run a saved advanced search
[**submit-advanced-search**](#submit-advanced-search) | **Post** `/advanced_search` | Save an advanced search query


## get-advanced-search
Get saved advanced search queries
Get saved advanced search queries

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-advanced-search)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedSearchRequest struct via the builder pattern


### Return type

[**GetAdvancedSearch200Response**](../models/get-advanced-search200-response)

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

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AdvancedSearchAPI.GetAdvancedSearch(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.AdvancedSearchAPI.GetAdvancedSearch(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.GetAdvancedSearch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdvancedSearch`: GetAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.GetAdvancedSearch`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-advanced-search
Update a saved advanced search
Update a saved advanced search query

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-advanced-search)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAdvancedSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitAdvancedSearchRequest** | [**SubmitAdvancedSearchRequest**](../models/submit-advanced-search-request) |  | 

### Return type

[**SubmitAdvancedSearch200Response**](../models/submit-advanced-search200-response)

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
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete
    submitadvancedsearchrequest := []byte(``) // SubmitAdvancedSearchRequest | 

    var submitAdvancedSearchRequest NERM.SubmitAdvancedSearchRequest
    if err := json.Unmarshal(submitadvancedsearchrequest, &submitAdvancedSearchRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AdvancedSearchAPI.PatchAdvancedSearch(context.Background(), id).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Execute()
	  //resp, r, err := apiClient.NERM.AdvancedSearchAPI.PatchAdvancedSearch(context.Background(), id).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.PatchAdvancedSearch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAdvancedSearch`: SubmitAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.PatchAdvancedSearch`: %v\n", resp)
}
```

[[Back to top]](#)

## search-advanced-search
Run profiles advanced search
Run an advanced search for profiles, without saving the query

[API Spec](https://developer.sailpoint.com/docs/api/NERM/search-advanced-search)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAdvancedSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAdvancedSearchRequest** | [**SubmitAdvancedSearchRequest**](../models/submit-advanced-search-request) |  | 
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 

### Return type

[**SearchAdvancedSearch200Response**](../models/search-advanced-search200-response)

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
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    submitadvancedsearchrequest := []byte(``) // SubmitAdvancedSearchRequest | 
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    order := `created_at` // string | The field to order results by. (optional) # string | The field to order results by. (optional)

    var submitAdvancedSearchRequest NERM.SubmitAdvancedSearchRequest
    if err := json.Unmarshal(submitadvancedsearchrequest, &submitAdvancedSearchRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AdvancedSearchAPI.SearchAdvancedSearch(context.Background()).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Execute()
	  //resp, r, err := apiClient.NERM.AdvancedSearchAPI.SearchAdvancedSearch(context.Background()).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.SearchAdvancedSearch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAdvancedSearch`: SearchAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.SearchAdvancedSearch`: %v\n", resp)
}
```

[[Back to top]](#)

## search-advanced-searchby-id
Run a saved advanced search
Run a saved advanced search query

[API Spec](https://developer.sailpoint.com/docs/api/NERM/search-advanced-searchby-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchAdvancedSearchbyIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 

### Return type

[**SearchAdvancedSearch200Response**](../models/search-advanced-search200-response)

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
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    order := `created_at` // string | The field to order results by. (optional) # string | The field to order results by. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AdvancedSearchAPI.SearchAdvancedSearchbyID(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.AdvancedSearchAPI.SearchAdvancedSearchbyID(context.Background(), id).Limit(limit).Offset(offset).Order(order).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.SearchAdvancedSearchbyID``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchAdvancedSearchbyID`: SearchAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.SearchAdvancedSearchbyID`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-advanced-search
Save an advanced search query
Save an advanced search query for later use

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-advanced-search)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAdvancedSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAdvancedSearchRequest** | [**SubmitAdvancedSearchRequest**](../models/submit-advanced-search-request) |  | 

### Return type

[**SubmitAdvancedSearch200Response**](../models/submit-advanced-search200-response)

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
    NERM "github.com/sailpoint-oss/golang-sdk/v2/api_nerm"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    submitadvancedsearchrequest := []byte(``) // SubmitAdvancedSearchRequest | 

    var submitAdvancedSearchRequest NERM.SubmitAdvancedSearchRequest
    if err := json.Unmarshal(submitadvancedsearchrequest, &submitAdvancedSearchRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AdvancedSearchAPI.SubmitAdvancedSearch(context.Background()).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Execute()
	  //resp, r, err := apiClient.NERM.AdvancedSearchAPI.SubmitAdvancedSearch(context.Background()).SubmitAdvancedSearchRequest(submitAdvancedSearchRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.SubmitAdvancedSearch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAdvancedSearch`: SubmitAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.SubmitAdvancedSearch`: %v\n", resp)
}
```

[[Back to top]](#)

