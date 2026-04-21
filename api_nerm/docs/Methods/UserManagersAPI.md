---
id: nerm-user-managers
title: UserManagers
pagination_label: UserManagers
sidebar_label: UserManagers
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserManagers', 'NERMUserManagers'] 
slug: /tools/sdk/go/nerm/methods/user-managers
tags: ['SDK', 'Software Development Kit', 'UserManagers', 'NERMUserManagers']
---

# UserManagersAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-user-manager**](#get-user-manager) | **Get** `/user_managers/{id}` | Find user-manager relationship
[**get-user-managers**](#get-user-managers) | **Get** `/user_managers` | Get user-manager relationships
[**patch-user-manager**](#patch-user-manager) | **Patch** `/user_managers/{id}` | Update a user-manager relationship
[**patch-user-managers**](#patch-user-managers) | **Patch** `/user_managers` | Update multiple user-manager relationships
[**submit-user-manager**](#submit-user-manager) | **Post** `/user_manager` | Create a new user-manager relationship
[**submit-user-managers**](#submit-user-managers) | **Post** `/user_managers` | Create multiple new user-manager relationships


## get-user-manager
Find user-manager relationship
Info for a specific user-manager relationship

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-manager)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitUserManager200Response**](../models/submit-user-manager200-response)

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
    resp, r, err := apiClient.NERM.UserManagersAPI.GetUserManager(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.GetUserManager(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.GetUserManager``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserManager`: SubmitUserManager200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.GetUserManager`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-managers
Get user-manager relationships
This endpoint can retrieve user-manager relationships from Lifecycle or you can search for user-manager relationships using parameters

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-managers)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **userId** | **string** | The ID of a user for filtering | 
 **managerId** | **string** | The ID of a user for filtering | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetUserManagers200Response**](../models/get-user-managers200-response)

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
    managerId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | The ID of a user for filtering (optional) # string | The ID of a user for filtering (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserManagersAPI.GetUserManagers(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.GetUserManagers(context.Background()).Limit(limit).Offset(offset).Order(order).UserId(userId).ManagerId(managerId).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.GetUserManagers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserManagers`: GetUserManagers200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.GetUserManagers`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-manager
Update a user-manager relationship
Update a user-manager relationship by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-manager)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitUserManagerRequest** | [**SubmitUserManagerRequest**](../models/submit-user-manager-request) |  | 

### Return type

[**SubmitUserManager200Response**](../models/submit-user-manager200-response)

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
    submitusermanagerrequest := []byte(``) // SubmitUserManagerRequest | 

    var submitUserManagerRequest NERM.SubmitUserManagerRequest
    if err := json.Unmarshal(submitusermanagerrequest, &submitUserManagerRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserManagersAPI.PatchUserManager(context.Background(), id).SubmitUserManagerRequest(submitUserManagerRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.PatchUserManager(context.Background(), id).SubmitUserManagerRequest(submitUserManagerRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.PatchUserManager``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserManager`: SubmitUserManager200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.PatchUserManager`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-managers
Update multiple user-manager relationships
Update multiple user-manager relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-managers)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserManagersRequest** | [**SubmitUserManagersRequest**](../models/submit-user-managers-request) |  | 

### Return type

[**SubmitUserManagers200Response**](../models/submit-user-managers200-response)

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
    submitusermanagersrequest := []byte(``) // SubmitUserManagersRequest | 

    var submitUserManagersRequest NERM.SubmitUserManagersRequest
    if err := json.Unmarshal(submitusermanagersrequest, &submitUserManagersRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserManagersAPI.PatchUserManagers(context.Background()).SubmitUserManagersRequest(submitUserManagersRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.PatchUserManagers(context.Background()).SubmitUserManagersRequest(submitUserManagersRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.PatchUserManagers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserManagers`: SubmitUserManagers200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.PatchUserManagers`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-manager
Create a new user-manager relationship
Create a new user-manager relationship

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-manager)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserManagerRequest** | [**SubmitUserManagerRequest**](../models/submit-user-manager-request) |  | 

### Return type

[**SubmitUserManager200Response**](../models/submit-user-manager200-response)

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
    submitusermanagerrequest := []byte(``) // SubmitUserManagerRequest | 

    var submitUserManagerRequest NERM.SubmitUserManagerRequest
    if err := json.Unmarshal(submitusermanagerrequest, &submitUserManagerRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserManagersAPI.SubmitUserManager(context.Background()).SubmitUserManagerRequest(submitUserManagerRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.SubmitUserManager(context.Background()).SubmitUserManagerRequest(submitUserManagerRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.SubmitUserManager``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserManager`: SubmitUserManager200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.SubmitUserManager`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-managers
Create multiple new user-manager relationships
Create multiple new user-manager relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-managers)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserManagersRequest** | [**SubmitUserManagersRequest**](../models/submit-user-managers-request) |  | 

### Return type

[**SubmitUserManagers200Response**](../models/submit-user-managers200-response)

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
    submitusermanagersrequest := []byte(``) // SubmitUserManagersRequest | 

    var submitUserManagersRequest NERM.SubmitUserManagersRequest
    if err := json.Unmarshal(submitusermanagersrequest, &submitUserManagersRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserManagersAPI.SubmitUserManagers(context.Background()).SubmitUserManagersRequest(submitUserManagersRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserManagersAPI.SubmitUserManagers(context.Background()).SubmitUserManagersRequest(submitUserManagersRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserManagersAPI.SubmitUserManagers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserManagers`: SubmitUserManagers200Response
    fmt.Fprintf(os.Stdout, "Response from `UserManagersAPI.SubmitUserManagers`: %v\n", resp)
}
```

[[Back to top]](#)

