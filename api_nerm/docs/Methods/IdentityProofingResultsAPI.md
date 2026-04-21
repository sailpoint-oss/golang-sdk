---
id: nerm-identity-proofing-results
title: IdentityProofingResults
pagination_label: IdentityProofingResults
sidebar_label: IdentityProofingResults
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'IdentityProofingResults', 'NERMIdentityProofingResults'] 
slug: /tools/sdk/go/nerm/methods/identity-proofing-results
tags: ['SDK', 'Software Development Kit', 'IdentityProofingResults', 'NERMIdentityProofingResults']
---

# IdentityProofingResultsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-identity-proofing-results**](#get-identity-proofing-results) | **Get** `/identity_proofing_results` | Get identity proofing result data


## get-identity-proofing-results
Get identity proofing result data
Retrieves identity proofing result data in bulk from Lifecycle

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-identity-proofing-results)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProofingResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **profileId** | **string** | Profile ID to filter by | 
 **workflowSessionId** | **string** | Workflow Session ID to filter by | 
 **result** | **string** | ID Proofing Result to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetIdentityProofingResults200Response**](../models/get-identity-proofing-results200-response)

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
    profileId := `4e480441-451d-47d9-87c2-9a0f0fe135eb` // string | Profile ID to filter by (optional) # string | Profile ID to filter by (optional)
    workflowSessionId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | Workflow Session ID to filter by (optional) # string | Workflow Session ID to filter by (optional)
    result := `pass` // string | ID Proofing Result to filter by (optional) # string | ID Proofing Result to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.IdentityProofingResultsAPI.GetIdentityProofingResults(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.IdentityProofingResultsAPI.GetIdentityProofingResults(context.Background()).Limit(limit).Offset(offset).Order(order).ProfileId(profileId).WorkflowSessionId(workflowSessionId).Result(result).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `IdentityProofingResultsAPI.GetIdentityProofingResults``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProofingResults`: GetIdentityProofingResults200Response
    fmt.Fprintf(os.Stdout, "Response from `IdentityProofingResultsAPI.GetIdentityProofingResults`: %v\n", resp)
}
```

[[Back to top]](#)

