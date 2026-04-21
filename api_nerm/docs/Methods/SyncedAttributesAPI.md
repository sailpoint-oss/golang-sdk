---
id: nerm-synced-attributes
title: SyncedAttributes
pagination_label: SyncedAttributes
sidebar_label: SyncedAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SyncedAttributes', 'NERMSyncedAttributes'] 
slug: /tools/sdk/go/nerm/methods/synced-attributes
tags: ['SDK', 'Software Development Kit', 'SyncedAttributes', 'NERMSyncedAttributes']
---

# SyncedAttributesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-synced-attribute**](#create-synced-attribute) | **Post** `/profile_types/{profile_type_id}/synced_attributes` | Create a synced attribute
[**delete-synced-attribute**](#delete-synced-attribute) | **Delete** `/profile_types/{profile_type_id}/synced_attributes/{ne_attribute_id}` | Delete synced attribute
[**get-profile-type-attributes**](#get-profile-type-attributes) | **Get** `/profile_types/{profile_type_id}/ne_attributes` | profile_types/ne_attributes synced status


## create-synced-attribute
Create a synced attribute
Create synced attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-synced-attribute)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSyncedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncedAttribute1** | [**SyncedAttribute1**](../models/synced-attribute1) |  | 

### Return type

[**CreateSyncedAttribute201Response**](../models/create-synced-attribute201-response)

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
    syncedattribute1 := []byte(`Object`) // SyncedAttribute1 | 

    var syncedAttribute1 NERM.SyncedAttribute1
    if err := json.Unmarshal(syncedattribute1, &syncedAttribute1); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.SyncedAttributesAPI.CreateSyncedAttribute(context.Background()).SyncedAttribute1(syncedAttribute1).Execute()
	  //resp, r, err := apiClient.NERM.SyncedAttributesAPI.CreateSyncedAttribute(context.Background()).SyncedAttribute1(syncedAttribute1).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `SyncedAttributesAPI.CreateSyncedAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSyncedAttribute`: CreateSyncedAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `SyncedAttributesAPI.CreateSyncedAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-synced-attribute
Delete synced attribute
Delete a synced attribute.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-synced-attribute)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSyncedAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileTypeId** | **string** | Profile Type ID for filtering | 
 **neAttributeId** | **string** | ID of an attribute for filtering | 

### Return type

[**DeleteProfileTypeById200Response**](../models/delete-profile-type-by-id200-response)

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
    profileTypeId := `79ed1cb6-9977-4965-9bfe-f2bcc242523e` // string | Profile Type ID for filtering (optional) # string | Profile Type ID for filtering (optional)
    neAttributeId := `33f072dd-13b4-41e1-8ea0-16f2a59b57c8` // string | ID of an attribute for filtering (optional) # string | ID of an attribute for filtering (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.SyncedAttributesAPI.DeleteSyncedAttribute(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.SyncedAttributesAPI.DeleteSyncedAttribute(context.Background()).ProfileTypeId(profileTypeId).NeAttributeId(neAttributeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `SyncedAttributesAPI.DeleteSyncedAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSyncedAttribute`: DeleteProfileTypeById200Response
    fmt.Fprintf(os.Stdout, "Response from `SyncedAttributesAPI.DeleteSyncedAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-type-attributes
profile_types/ne_attributes synced status
Get ne attributes and synced attribute relationship to profile type.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-type-attributes)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileTypeAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileTypeId** | **string** | Profile Type ID for filtering | 
 **activeFilter** | **string** | Filter for profile type synced attributes | 
 **search** | **string** | Filter by string | 
 **page** | **int32** | Pagination items per page | 
 **sort** | [**GetProfileTypeAttributesSortParameter**](../models/get-profile-type-attributes-sort-parameter) | How records should be sorted | 

### Return type

[**GetProfileTypeAttributes200Response**](../models/get-profile-type-attributes200-response)

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
    profileTypeId := `79ed1cb6-9977-4965-9bfe-f2bcc242523e` // string | Profile Type ID for filtering (optional) # string | Profile Type ID for filtering (optional)
    activeFilter := `all` // string | Filter for profile type synced attributes (optional) # string | Filter for profile type synced attributes (optional)
    search := `search` // string | Filter by string (optional) # string | Filter by string (optional)
    page := 5 // int32 | Pagination items per page (optional) # int32 | Pagination items per page (optional)
    sort := []byte(``) // GetProfileTypeAttributesSortParameter | How records should be sorted (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.SyncedAttributesAPI.GetProfileTypeAttributes(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.SyncedAttributesAPI.GetProfileTypeAttributes(context.Background()).ProfileTypeId(profileTypeId).ActiveFilter(activeFilter).Search(search).Page(page).Sort(sort).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `SyncedAttributesAPI.GetProfileTypeAttributes``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileTypeAttributes`: GetProfileTypeAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `SyncedAttributesAPI.GetProfileTypeAttributes`: %v\n", resp)
}
```

[[Back to top]](#)

