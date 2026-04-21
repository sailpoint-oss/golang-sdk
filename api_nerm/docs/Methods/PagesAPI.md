---
id: nerm-pages
title: Pages
pagination_label: Pages
sidebar_label: Pages
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Pages', 'NERMPages'] 
slug: /tools/sdk/go/nerm/methods/pages
tags: ['SDK', 'Software Development Kit', 'Pages', 'NERMPages']
---

# PagesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-profile-page**](#create-profile-page) | **Post** `/pages/profile_pages` | Create a profile page
[**create-workflow-page**](#create-workflow-page) | **Post** `/pages/workflow_pages` | Create a workflow page


## create-profile-page
Create a profile page
Create a profile page

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-profile-page)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfilePageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfilePageRequest** | [**CreateProfilePageRequest**](../models/create-profile-page-request) |  | 

### Return type

[**CreateProfilePage200Response**](../models/create-profile-page200-response)

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
    createprofilepagerequest := []byte(``) // CreateProfilePageRequest | 

    var createProfilePageRequest NERM.CreateProfilePageRequest
    if err := json.Unmarshal(createprofilepagerequest, &createProfilePageRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PagesAPI.CreateProfilePage(context.Background()).CreateProfilePageRequest(createProfilePageRequest).Execute()
	  //resp, r, err := apiClient.NERM.PagesAPI.CreateProfilePage(context.Background()).CreateProfilePageRequest(createProfilePageRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PagesAPI.CreateProfilePage``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfilePage`: CreateProfilePage200Response
    fmt.Fprintf(os.Stdout, "Response from `PagesAPI.CreateProfilePage`: %v\n", resp)
}
```

[[Back to top]](#)

## create-workflow-page
Create a workflow page
Create a workflow page

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-workflow-page)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowPageRequest** | [**CreateWorkflowPageRequest**](../models/create-workflow-page-request) |  | 

### Return type

[**CreateProfilePage200Response**](../models/create-profile-page200-response)

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
    createworkflowpagerequest := []byte(``) // CreateWorkflowPageRequest | 

    var createWorkflowPageRequest NERM.CreateWorkflowPageRequest
    if err := json.Unmarshal(createworkflowpagerequest, &createWorkflowPageRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PagesAPI.CreateWorkflowPage(context.Background()).CreateWorkflowPageRequest(createWorkflowPageRequest).Execute()
	  //resp, r, err := apiClient.NERM.PagesAPI.CreateWorkflowPage(context.Background()).CreateWorkflowPageRequest(createWorkflowPageRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PagesAPI.CreateWorkflowPage``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowPage`: CreateProfilePage200Response
    fmt.Fprintf(os.Stdout, "Response from `PagesAPI.CreateWorkflowPage`: %v\n", resp)
}
```

[[Back to top]](#)

