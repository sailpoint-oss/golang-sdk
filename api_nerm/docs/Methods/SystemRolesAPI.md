---
id: nerm-system-roles
title: SystemRoles
pagination_label: SystemRoles
sidebar_label: SystemRoles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SystemRoles', 'NERMSystemRoles'] 
slug: /tools/sdk/go/nerm/methods/system-roles
tags: ['SDK', 'Software Development Kit', 'SystemRoles', 'NERMSystemRoles']
---

# SystemRolesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-system-roles**](#get-system-roles) | **Get** `/system_roles` | Get system roles


## get-system-roles
Get system roles
This endpoint can retrieve system roles from NERM. Optionally you can provide parameters to filter results.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-system-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetSystemRoles200Response**](../models/get-system-roles200-response)

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
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.SystemRolesAPI.GetSystemRoles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.SystemRolesAPI.GetSystemRoles(context.Background()).Limit(limit).Offset(offset).Order(order).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `SystemRolesAPI.GetSystemRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemRoles`: GetSystemRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemRolesAPI.GetSystemRoles`: %v\n", resp)
}
```

[[Back to top]](#)

