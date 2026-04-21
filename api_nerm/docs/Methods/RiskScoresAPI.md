---
id: nerm-risk-scores
title: RiskScores
pagination_label: RiskScores
sidebar_label: RiskScores
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RiskScores', 'NERMRiskScores'] 
slug: /tools/sdk/go/nerm/methods/risk-scores
tags: ['SDK', 'Software Development Kit', 'RiskScores', 'NERMRiskScores']
---

# RiskScoresAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-risk-score**](#get-risk-score) | **Get** `/risk_scores/{id}` | Find risk score data
[**get-risk-scores**](#get-risk-scores) | **Get** `/risk_scores` | Get risk score data


## get-risk-score
Find risk score data
Find risk score data by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-risk-score)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRiskScore200Response**](../models/get-risk-score200-response)

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
    resp, r, err := apiClient.NERM.RiskScoresAPI.GetRiskScore(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.RiskScoresAPI.GetRiskScore(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RiskScoresAPI.GetRiskScore``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskScore`: GetRiskScore200Response
    fmt.Fprintf(os.Stdout, "Response from `RiskScoresAPI.GetRiskScore`: %v\n", resp)
}
```

[[Back to top]](#)

## get-risk-scores
Get risk score data
This endpoint can retrieve risk score data in bulk from Lifecycle

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-risk-scores)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskScoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **objectId** | **string** | ID of an object for filtering. Used along with object_type | 
 **objectType** | **string** | Type of object that object_id represents | 
 **overallRiskLevelId** | **string** | Overall risk level to filter by | 
 **impactRiskLevelId** | **string** | Impact risk level to filter by | 
 **probabilityRiskLevelId** | **string** | Probability risk level to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetRiskScores200Response**](../models/get-risk-scores200-response)

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
    objectId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | ID of an object for filtering. Used along with object_type (optional) # string | ID of an object for filtering. Used along with object_type (optional)
    objectType := `Profile` // string | Type of object that object_id represents (optional) # string | Type of object that object_id represents (optional)
    overallRiskLevelId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | Overall risk level to filter by (optional) # string | Overall risk level to filter by (optional)
    impactRiskLevelId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | Impact risk level to filter by (optional) # string | Impact risk level to filter by (optional)
    probabilityRiskLevelId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | Probability risk level to filter by (optional) # string | Probability risk level to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RiskScoresAPI.GetRiskScores(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.RiskScoresAPI.GetRiskScores(context.Background()).Limit(limit).Offset(offset).Order(order).ObjectId(objectId).ObjectType(objectType).OverallRiskLevelId(overallRiskLevelId).ImpactRiskLevelId(impactRiskLevelId).ProbabilityRiskLevelId(probabilityRiskLevelId).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RiskScoresAPI.GetRiskScores``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRiskScores`: GetRiskScores200Response
    fmt.Fprintf(os.Stdout, "Response from `RiskScoresAPI.GetRiskScores`: %v\n", resp)
}
```

[[Back to top]](#)

