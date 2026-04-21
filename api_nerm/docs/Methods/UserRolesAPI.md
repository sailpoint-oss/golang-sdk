---
id: nerm-user-roles
title: UserRoles
pagination_label: UserRoles
sidebar_label: UserRoles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserRoles', 'NERMUserRoles'] 
slug: /tools/sdk/go/nerm/methods/user-roles
tags: ['SDK', 'Software Development Kit', 'UserRoles', 'NERMUserRoles']
---

# UserRolesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-user-role**](#delete-user-role) | **Delete** `/user_role/{id}` | Delete a user role assignment
[**get-user-role**](#get-user-role) | **Get** `/user_roles/{id}` | Find user role pairing
[**get-user-roles**](#get-user-roles) | **Get** `/user_roles` | Get user role pairings
[**patch-user-role**](#patch-user-role) | **Patch** `/user_roles/{id}` | Update a user role pairing
[**patch-user-roles**](#patch-user-roles) | **Patch** `/user_roles` | Update multiple user role pairings
[**submit-user-role**](#submit-user-role) | **Post** `/user_role` | Assign new role to user
[**submit-user-roles**](#submit-user-roles) | **Post** `/user_roles` | Create new user role pairings


## delete-user-role
Delete a user role assignment
Delete a user role assignment

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-user-role)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

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
    resp, r, err := apiClient.NERM.UserRolesAPI.DeleteUserRole(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.DeleteUserRole(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.DeleteUserRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.DeleteUserRole`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-role
Find user role pairing
Info for a specific user role pairing

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-role)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitUserRole200Response**](../models/submit-user-role200-response)

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
    resp, r, err := apiClient.NERM.UserRolesAPI.GetUserRole(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.GetUserRole(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.GetUserRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRole`: SubmitUserRole200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.GetUserRole`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-roles
Get user role pairings
This endpoint can retrieve user role pairings from Lifecycle or you can search for user role pairings using parameters

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **userId** | **string** | The ID of a user for filtering | 
 **roleId** | **string** | The ID of a role for filtering | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetUserRoles200Response**](../models/get-user-roles200-response)

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
    userId := `bba9cfb2-96c1-4acb-ac79-a21732527265` // string | The ID of a user for filtering (optional) # string | The ID of a user for filtering (optional)
    roleId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | The ID of a role for filtering (optional) # string | The ID of a role for filtering (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserRolesAPI.GetUserRoles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.GetUserRoles(context.Background()).Limit(limit).Offset(offset).Order(order).UserId(userId).RoleId(roleId).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.GetUserRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserRoles`: GetUserRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.GetUserRoles`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-role
Update a user role pairing
Update a user role pairing by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-role)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitUserRoleRequest** | [**SubmitUserRoleRequest**](../models/submit-user-role-request) |  | 

### Return type

[**SubmitUserRole200Response**](../models/submit-user-role200-response)

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
    submituserrolerequest := []byte(``) // SubmitUserRoleRequest | 

    var submitUserRoleRequest NERM.SubmitUserRoleRequest
    if err := json.Unmarshal(submituserrolerequest, &submitUserRoleRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserRolesAPI.PatchUserRole(context.Background(), id).SubmitUserRoleRequest(submitUserRoleRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.PatchUserRole(context.Background(), id).SubmitUserRoleRequest(submitUserRoleRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.PatchUserRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserRole`: SubmitUserRole200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.PatchUserRole`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-roles
Update multiple user role pairings
Update multiple user role pairings

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserRolesRequest** | [**SubmitUserRolesRequest**](../models/submit-user-roles-request) |  | 

### Return type

[**SubmitUserRoles200Response**](../models/submit-user-roles200-response)

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
    submituserrolesrequest := []byte(``) // SubmitUserRolesRequest | 

    var submitUserRolesRequest NERM.SubmitUserRolesRequest
    if err := json.Unmarshal(submituserrolesrequest, &submitUserRolesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserRolesAPI.PatchUserRoles(context.Background()).SubmitUserRolesRequest(submitUserRolesRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.PatchUserRoles(context.Background()).SubmitUserRolesRequest(submitUserRolesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.PatchUserRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserRoles`: SubmitUserRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.PatchUserRoles`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-role
Assign new role to user
Assign a new role to a user

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-role)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserRoleRequest** | [**SubmitUserRoleRequest**](../models/submit-user-role-request) |  | 

### Return type

[**SubmitUserRole200Response**](../models/submit-user-role200-response)

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
    submituserrolerequest := []byte(``) // SubmitUserRoleRequest | 

    var submitUserRoleRequest NERM.SubmitUserRoleRequest
    if err := json.Unmarshal(submituserrolerequest, &submitUserRoleRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserRolesAPI.SubmitUserRole(context.Background()).SubmitUserRoleRequest(submitUserRoleRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.SubmitUserRole(context.Background()).SubmitUserRoleRequest(submitUserRoleRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.SubmitUserRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserRole`: SubmitUserRole200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.SubmitUserRole`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-roles
Create new user role pairings
Create multiple new user role pairings

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserRolesRequest** | [**SubmitUserRolesRequest**](../models/submit-user-roles-request) |  | 

### Return type

[**SubmitUserRoles200Response**](../models/submit-user-roles200-response)

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
    submituserrolesrequest := []byte(``) // SubmitUserRolesRequest | 

    var submitUserRolesRequest NERM.SubmitUserRolesRequest
    if err := json.Unmarshal(submituserrolesrequest, &submitUserRolesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserRolesAPI.SubmitUserRoles(context.Background()).SubmitUserRolesRequest(submitUserRolesRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserRolesAPI.SubmitUserRoles(context.Background()).SubmitUserRolesRequest(submitUserRolesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserRolesAPI.SubmitUserRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserRoles`: SubmitUserRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserRolesAPI.SubmitUserRoles`: %v\n", resp)
}
```

[[Back to top]](#)

