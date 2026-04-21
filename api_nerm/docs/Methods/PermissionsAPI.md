---
id: nerm-permissions
title: Permissions
pagination_label: Permissions
sidebar_label: Permissions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Permissions', 'NERMPermissions'] 
slug: /tools/sdk/go/nerm/methods/permissions
tags: ['SDK', 'Software Development Kit', 'Permissions', 'NERMPermissions']
---

# PermissionsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-permission**](#create-permission) | **Post** `/permissions` | Create a permission


## create-permission
Create a permission
This endpoint can create permissions for Lifecycle, Consolidation, and Collaboration

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-permission)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPermissionRequest** | [**CreatePermissionRequest**](../models/create-permission-request) |  | 

### Return type

[**CreatePermission200Response**](../models/create-permission200-response)

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
    createpermissionrequest := []byte(``) // CreatePermissionRequest | 

    var createPermissionRequest NERM.CreatePermissionRequest
    if err := json.Unmarshal(createpermissionrequest, &createPermissionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PermissionsAPI.CreatePermission(context.Background()).CreatePermissionRequest(createPermissionRequest).Execute()
	  //resp, r, err := apiClient.NERM.PermissionsAPI.CreatePermission(context.Background()).CreatePermissionRequest(createPermissionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.CreatePermission``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePermission`: CreatePermission200Response
    fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.CreatePermission`: %v\n", resp)
}
```

[[Back to top]](#)

