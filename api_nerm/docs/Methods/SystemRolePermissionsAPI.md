---
id: nerm-system-role-permissions
title: SystemRolePermissions
pagination_label: SystemRolePermissions
sidebar_label: SystemRolePermissions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'SystemRolePermissions', 'NERMSystemRolePermissions'] 
slug: /tools/sdk/go/nerm/methods/system-role-permissions
tags: ['SDK', 'Software Development Kit', 'SystemRolePermissions', 'NERMSystemRolePermissions']
---

# SystemRolePermissionsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-system-role-permission**](#create-system-role-permission) | **Post** `/system_role_permissions` | Create a system role permission


## create-system-role-permission
Create a system role permission
This endpoint can create system role permissions for Lifecycle System Roles

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-system-role-permission)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSystemRolePermissionRequest** | [**CreateSystemRolePermissionRequest**](../models/create-system-role-permission-request) |  | 

### Return type

[**CreateSystemRolePermission200Response**](../models/create-system-role-permission200-response)

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
    createsystemrolepermissionrequest := []byte(``) // CreateSystemRolePermissionRequest | 

    var createSystemRolePermissionRequest NERM.CreateSystemRolePermissionRequest
    if err := json.Unmarshal(createsystemrolepermissionrequest, &createSystemRolePermissionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.SystemRolePermissionsAPI.CreateSystemRolePermission(context.Background()).CreateSystemRolePermissionRequest(createSystemRolePermissionRequest).Execute()
	  //resp, r, err := apiClient.NERM.SystemRolePermissionsAPI.CreateSystemRolePermission(context.Background()).CreateSystemRolePermissionRequest(createSystemRolePermissionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `SystemRolePermissionsAPI.CreateSystemRolePermission``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSystemRolePermission`: CreateSystemRolePermission200Response
    fmt.Fprintf(os.Stdout, "Response from `SystemRolePermissionsAPI.CreateSystemRolePermission`: %v\n", resp)
}
```

[[Back to top]](#)

