---
id: nerm-roles
title: Roles
pagination_label: Roles
sidebar_label: Roles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Roles', 'NERMRoles'] 
slug: /tools/sdk/go/nerm/methods/roles
tags: ['SDK', 'Software Development Kit', 'Roles', 'NERMRoles']
---

# RolesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-role**](#get-role) | **Get** `/roles/{id}` | Find role by id
[**get-roles**](#get-roles) | **Get** `/roles` | Get roles
[**patch-role**](#patch-role) | **Patch** `/roles/{id}` | Update an existing role
[**patch-roles**](#patch-roles) | **Patch** `/roles` | Update multiple roles
[**submit-role**](#submit-role) | **Post** `/role` | Create a new role
[**submit-roles**](#submit-roles) | **Post** `/roles` | Create multiple new roles


## get-role
Find role by id
Info for a specific user role

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-role)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitRole200Response**](../models/submit-role200-response)

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
    resp, r, err := apiClient.NERM.RolesAPI.GetRole(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.GetRole(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: SubmitRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRole`: %v\n", resp)
}
```

[[Back to top]](#)

## get-roles
Get roles
This endpoint can retrieve roles from NERM. Optionally you can provide parameters to filter results.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]
 **type_** | **string** | Filter roles by type. | 

### Return type

[**GetRoles200Response**](../models/get-roles200-response)

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
    type_ := `NeprofileRole` // string | Filter roles by type. (optional) # string | Filter roles by type. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RolesAPI.GetRoles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.GetRoles(context.Background()).Limit(limit).Offset(offset).Order(order).Metadata(metadata).Type_(type_).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: GetRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoles`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-role
Update an existing role
Update an existing role

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-role)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitRoleRequest** | [**SubmitRoleRequest**](../models/submit-role-request) |  | 

### Return type

[**SubmitRole200Response**](../models/submit-role200-response)

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
    submitrolerequest := []byte(``) // SubmitRoleRequest | 

    var submitRoleRequest NERM.SubmitRoleRequest
    if err := json.Unmarshal(submitrolerequest, &submitRoleRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RolesAPI.PatchRole(context.Background(), id).SubmitRoleRequest(submitRoleRequest).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.PatchRole(context.Background(), id).SubmitRoleRequest(submitRoleRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PatchRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRole`: SubmitRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.PatchRole`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-roles
Update multiple roles
Update multiple users

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRolesRequest** | [**SubmitRolesRequest**](../models/submit-roles-request) |  | 

### Return type

[**SubmitRoles200Response**](../models/submit-roles200-response)

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
    submitrolesrequest := []byte(``) // SubmitRolesRequest | 

    var submitRolesRequest NERM.SubmitRolesRequest
    if err := json.Unmarshal(submitrolesrequest, &submitRolesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RolesAPI.PatchRoles(context.Background()).SubmitRolesRequest(submitRolesRequest).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.PatchRoles(context.Background()).SubmitRolesRequest(submitRolesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.PatchRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRoles`: SubmitRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.PatchRoles`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-role
Create a new role
Create a new role

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-role)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRoleRequest** | [**SubmitRoleRequest**](../models/submit-role-request) |  | 

### Return type

[**SubmitRole200Response**](../models/submit-role200-response)

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
    submitrolerequest := []byte(``) // SubmitRoleRequest | 

    var submitRoleRequest NERM.SubmitRoleRequest
    if err := json.Unmarshal(submitrolerequest, &submitRoleRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RolesAPI.SubmitRole(context.Background()).SubmitRoleRequest(submitRoleRequest).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.SubmitRole(context.Background()).SubmitRoleRequest(submitRoleRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.SubmitRole``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRole`: SubmitRole200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.SubmitRole`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-roles
Create multiple new roles
Create multiple new users

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-roles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRolesRequest** | [**SubmitRolesRequest**](../models/submit-roles-request) |  | 

### Return type

[**SubmitRoles200Response**](../models/submit-roles200-response)

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
    submitrolesrequest := []byte(``) // SubmitRolesRequest | 

    var submitRolesRequest NERM.SubmitRolesRequest
    if err := json.Unmarshal(submitrolesrequest, &submitRolesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RolesAPI.SubmitRoles(context.Background()).SubmitRolesRequest(submitRolesRequest).Execute()
	  //resp, r, err := apiClient.NERM.RolesAPI.SubmitRoles(context.Background()).SubmitRolesRequest(submitRolesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.SubmitRoles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRoles`: SubmitRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.SubmitRoles`: %v\n", resp)
}
```

[[Back to top]](#)

