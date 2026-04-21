---
id: nerm-delegations
title: Delegations
pagination_label: Delegations
sidebar_label: Delegations
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Delegations', 'NERMDelegations'] 
slug: /tools/sdk/go/nerm/methods/delegations
tags: ['SDK', 'Software Development Kit', 'Delegations', 'NERMDelegations']
---

# DelegationsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delegations-get**](#delegations-get) | **Get** `/delegations` | List delegations
[**delegations-id-delete**](#delegations-id-delete) | **Delete** `/delegations/{id}` | Delete a delegation
[**delegations-id-get**](#delegations-id-get) | **Get** `/delegations/{id}` | Get a single delegation
[**delegations-id-patch**](#delegations-id-patch) | **Patch** `/delegations/{id}` | Update a delegation
[**delegations-post**](#delegations-post) | **Post** `/delegations` | Create a delegation


## delegations-get
List delegations
Returns a list of delegation records, optionally filtered by delegate, delegator, or expiration status.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delegations-get)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelegationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delegateId** | **string** | Filter by delegate ID | [default to &quot;false&quot;]
 **delegatorId** | **string** | Filter by delegator ID | [default to &quot;false&quot;]
 **expired** | **bool** | Filter by expiration status (true for expired, false for not expired) | [default to false]
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 

### Return type

[**DelegationsGet200Response**](../models/delegations-get200-response)

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
    delegateId := `ac4aae0b-4140-49a4-a84c-126762fd0c8f` // string | Filter by delegate ID (optional) (default to "false") # string | Filter by delegate ID (optional) (default to "false")
    delegatorId := `ac4aae0b-4140-49a4-a84c-126762fd0c8f` // string | Filter by delegator ID (optional) (default to "false") # string | Filter by delegator ID (optional) (default to "false")
    expired := true // bool | Filter by expiration status (true for expired, false for not expired) (optional) (default to false) # bool | Filter by expiration status (true for expired, false for not expired) (optional) (default to false)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsGet(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsGet(context.Background()).DelegateId(delegateId).DelegatorId(delegatorId).Expired(expired).Metadata(metadata).Limit(limit).Offset(offset).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DelegationsGet``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelegationsGet`: DelegationsGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.DelegationsGet`: %v\n", resp)
}
```

[[Back to top]](#)

## delegations-id-delete
Delete a delegation
Delete an existing delegation record.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delegations-id-delete)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegationsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
    r, err := apiClient.NERM.DelegationsAPI.DelegationsIdDelete(context.Background(), id).Execute()
	  //r, err := apiClient.NERM.DelegationsAPI.DelegationsIdDelete(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DelegationsIdDelete``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## delegations-id-get
Get a single delegation
Returns a single delegation record by its ID.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delegations-id-get)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegationsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DelegationsPost201Response**](../models/delegations-post201-response)

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
    resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsIdGet(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsIdGet(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DelegationsIdGet``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelegationsIdGet`: DelegationsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.DelegationsIdGet`: %v\n", resp)
}
```

[[Back to top]](#)

## delegations-id-patch
Update a delegation
Update an existing delegation record.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delegations-id-patch)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelegationsIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delegationsPostRequest** | [**DelegationsPostRequest**](../models/delegations-post-request) |  | 

### Return type

[**DelegationsPost201Response**](../models/delegations-post201-response)

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
    delegationspostrequest := []byte(``) // DelegationsPostRequest | 

    var delegationsPostRequest NERM.DelegationsPostRequest
    if err := json.Unmarshal(delegationspostrequest, &delegationsPostRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsIdPatch(context.Background(), id).DelegationsPostRequest(delegationsPostRequest).Execute()
	  //resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsIdPatch(context.Background(), id).DelegationsPostRequest(delegationsPostRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DelegationsIdPatch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelegationsIdPatch`: DelegationsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.DelegationsIdPatch`: %v\n", resp)
}
```

[[Back to top]](#)

## delegations-post
Create a delegation
Create a new delegation record.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delegations-post)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelegationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delegationsPostRequest** | [**DelegationsPostRequest**](../models/delegations-post-request) |  | 

### Return type

[**DelegationsPost201Response**](../models/delegations-post201-response)

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
    delegationspostrequest := []byte(``) // DelegationsPostRequest | 

    var delegationsPostRequest NERM.DelegationsPostRequest
    if err := json.Unmarshal(delegationspostrequest, &delegationsPostRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsPost(context.Background()).DelegationsPostRequest(delegationsPostRequest).Execute()
	  //resp, r, err := apiClient.NERM.DelegationsAPI.DelegationsPost(context.Background()).DelegationsPostRequest(delegationsPostRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DelegationsAPI.DelegationsPost``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DelegationsPost`: DelegationsPost201Response
    fmt.Fprintf(os.Stdout, "Response from `DelegationsAPI.DelegationsPost`: %v\n", resp)
}
```

[[Back to top]](#)

