---
id: nerm-risk-levels
title: RiskLevels
pagination_label: RiskLevels
sidebar_label: RiskLevels
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskLevels', 'NERMRiskLevels'] 
slug: /tools/sdk/go/nerm/methods/risk-levels
tags: ['SDK', 'Software Development Kit', 'RiskLevels', 'NERMRiskLevels']
---

# RiskLevelsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-risk-level**](#get-risk-level) | **Get** `/risk_levels/{id}` | Find risk level data
[**get-risk-levels**](#get-risk-levels) | **Get** `/risk_levels` | Get risk level data


## get-risk-level
Find risk level data
Find risk level data by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-risk-level)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRiskLevel200Response**](../models/get-risk-level200-response)

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

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RiskLevelsAPI.GetRiskLevel(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.RiskLevelsAPI.GetRiskLevel(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RiskLevelsAPI.GetRiskLevel``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskLevel`: GetRiskLevel200Response
    fmt.Fprintf(os.Stdout, "Response from `RiskLevelsAPI.GetRiskLevel`: %v\n", resp)
}
```

[[Back to top]](#)

## get-risk-levels
Get risk level data
This endpoint can retrieve risk level data in bulk from Lifecycle

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-risk-levels)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **label** | **string** | The attribute label to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetRiskLevels200Response**](../models/get-risk-levels200-response)

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
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    order := `created_at` // string | The field to order results by. (optional) # string | The field to order results by. (optional)
    label := `birthday` // string | The attribute label to filter by (optional) # string | The attribute label to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RiskLevelsAPI.GetRiskLevels(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.RiskLevelsAPI.GetRiskLevels(context.Background()).Limit(limit).Offset(offset).Order(order).Label(label).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RiskLevelsAPI.GetRiskLevels``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskLevels`: GetRiskLevels200Response
    fmt.Fprintf(os.Stdout, "Response from `RiskLevelsAPI.GetRiskLevels`: %v\n", resp)
}
```

[[Back to top]](#)

