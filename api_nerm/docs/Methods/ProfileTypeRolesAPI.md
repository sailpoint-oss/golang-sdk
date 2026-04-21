---
id: nerm-profile-type-roles
title: ProfileTypeRoles
pagination_label: ProfileTypeRoles
sidebar_label: ProfileTypeRoles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypeRoles', 'NERMProfileTypeRoles'] 
slug: /tools/sdk/go/nerm/methods/profile-type-roles
tags: ['SDK', 'Software Development Kit', 'ProfileTypeRoles', 'NERMProfileTypeRoles']
---

# ProfileTypeRolesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-profile-type-role**](#create-profile-type-role) | **Post** `/profile_type_roles` | Create a profile type role


## create-profile-type-role
Create a profile type role
This endpoint can create a profile type role. NOTE- The ability to toggle Allow/Block is done through the Profile Type

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-profile-type-role)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileTypeRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfileTypeRoleRequest** | [**CreateProfileTypeRoleRequest**](../models/create-profile-type-role-request) |  | 

### Return type

[**CreateProfileTypeRole200Response**](../models/create-profile-type-role200-response)

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
    createprofiletyperolerequest := []byte(``) // CreateProfileTypeRoleRequest | 

    var createProfileTypeRoleRequest NERM.CreateProfileTypeRoleRequest
    if err := json.Unmarshal(createprofiletyperolerequest, &createProfileTypeRoleRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypeRolesAPI.CreateProfileTypeRole(context.Background()).CreateProfileTypeRoleRequest(createProfileTypeRoleRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypeRolesAPI.CreateProfileTypeRole(context.Background()).CreateProfileTypeRoleRequest(createProfileTypeRoleRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypeRolesAPI.CreateProfileTypeRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfileTypeRole`: CreateProfileTypeRole200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypeRolesAPI.CreateProfileTypeRole`: %v\n", resp)
}
```

[[Back to top]](#)

