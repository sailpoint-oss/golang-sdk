---
id: nerm-users
title: Users
pagination_label: Users
sidebar_label: Users
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Users', 'NERMUsers'] 
slug: /tools/sdk/go/nerm/methods/users
tags: ['SDK', 'Software Development Kit', 'Users', 'NERMUsers']
---

# UsersAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-user**](#delete-user) | **Delete** `/users/{id}` | Delete a user
[**get-user**](#get-user) | **Get** `/users/{id}` | Find user by id
[**get-user-avatar**](#get-user-avatar) | **Get** `/users/{id}/avatar` | Retrieves  URL user avatar
[**get-users**](#get-users) | **Get** `/users` | Get users
[**patch-user**](#patch-user) | **Patch** `/users/{id}` | Update a user by id
[**patch-users**](#patch-users) | **Patch** `/users` | Update multiple users
[**submit-user**](#submit-user) | **Post** `/user` | Create a new user
[**submit-user-avatar**](#submit-user-avatar) | **Post** `/users/{id}/avatar` | Uploads new user avatar
[**submit-users**](#submit-users) | **Post** `/users` | Create multiple new users


## delete-user
Delete a user
Delete a user

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-user)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.DeleteUser(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.DeleteUser(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: DeleteProfileTypeById200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUser`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user
Find user by id
Info for a specific user

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitUser200Response**](../models/submit-user200-response)

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
    resp, r, err := apiClient.NERM.UsersAPI.GetUser(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.GetUser(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: SubmitUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-avatar
Retrieves  URL user avatar
Retrieves the URL of the user avatar

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-avatar)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Url**](../models/url)

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
    resp, r, err := apiClient.NERM.UsersAPI.GetUserAvatar(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.GetUserAvatar(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserAvatar``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAvatar`: Url
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserAvatar`: %v\n", resp)
}
```

[[Back to top]](#)

## get-users
Get users
This endpoint can retrieve users from Lifecycle or you can search for users using parameters

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-users)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **name** | **string** | object name for filtering | 
 **login** | **string** | The user login to search by | 
 **title** | **string** | The user title to search by | 
 **userStatus** | **string** | user status value for filtering | 
 **type_** | **string** | user type | 
 **email** | **string** | The user email to search by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]
 **sailpointIdentityId** | **string** | SailPoint identity ID | 

### Return type

[**GetUsers200Response**](../models/get-users200-response)

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
    name := `name` // string | object name for filtering (optional) # string | object name for filtering (optional)
    login := `jane.doe` // string | The user login to search by (optional) # string | The user login to search by (optional)
    title := `Marketing Director` // string | The user title to search by (optional) # string | The user title to search by (optional)
    userStatus := `Active` // string | user status value for filtering (optional) # string | user status value for filtering (optional)
    type_ := `NeprofileUser` // string | user type (optional) # string | user type (optional)
    email := `support@sailpoint.com` // string | The user email to search by (optional) # string | The user email to search by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)
    sailpointIdentityId := `9496f8d6ddab49c0bef1e9ee6f1b835a` // string | SailPoint identity ID (optional) # string | SailPoint identity ID (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.GetUsers(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.GetUsers(context.Background()).Limit(limit).Offset(offset).Order(order).Name(name).Login(login).Title(title).UserStatus(userStatus).Type_(type_).Email(email).Metadata(metadata).SailpointIdentityId(sailpointIdentityId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUsers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: GetUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUsers`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user
Update a user by id
Update a user by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitUserRequest** | [**SubmitUserRequest**](../models/submit-user-request) |  | 

### Return type

[**SubmitUser200Response**](../models/submit-user200-response)

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
    submituserrequest := []byte(``) // SubmitUserRequest | 

    var submitUserRequest NERM.SubmitUserRequest
    if err := json.Unmarshal(submituserrequest, &submitUserRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.PatchUser(context.Background(), id).SubmitUserRequest(submitUserRequest).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.PatchUser(context.Background(), id).SubmitUserRequest(submitUserRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PatchUser``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUser`: SubmitUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PatchUser`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-users
Update multiple users
Update multiple users

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-users)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUsersRequest** | [**SubmitUsersRequest**](../models/submit-users-request) |  | 

### Return type

[**SubmitUsers200Response**](../models/submit-users200-response)

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
    submitusersrequest := []byte(``) // SubmitUsersRequest | 

    var submitUsersRequest NERM.SubmitUsersRequest
    if err := json.Unmarshal(submitusersrequest, &submitUsersRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.PatchUsers(context.Background()).SubmitUsersRequest(submitUsersRequest).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.PatchUsers(context.Background()).SubmitUsersRequest(submitUsersRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PatchUsers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUsers`: SubmitUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PatchUsers`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user
Create a new user
Create a new user

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserRequest** | [**SubmitUserRequest**](../models/submit-user-request) |  | 

### Return type

[**SubmitUser200Response**](../models/submit-user200-response)

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
    submituserrequest := []byte(``) // SubmitUserRequest | 

    var submitUserRequest NERM.SubmitUserRequest
    if err := json.Unmarshal(submituserrequest, &submitUserRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.SubmitUser(context.Background()).SubmitUserRequest(submitUserRequest).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.SubmitUser(context.Background()).SubmitUserRequest(submitUserRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SubmitUser``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUser`: SubmitUser200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SubmitUser`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-avatar
Uploads new user avatar
Uploads a new user avatar

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-avatar)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**Url**](../models/url)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
    file := BINARY_DATA_HERE // *os.File |  (optional) # *os.File |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.SubmitUserAvatar(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.SubmitUserAvatar(context.Background(), id).File(file).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SubmitUserAvatar``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserAvatar`: Url
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SubmitUserAvatar`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-users
Create multiple new users
Create multiple new users

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-users)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUsersRequest** | [**SubmitUsersRequest**](../models/submit-users-request) |  | 

### Return type

[**SubmitUsers200Response**](../models/submit-users200-response)

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
    submitusersrequest := []byte(``) // SubmitUsersRequest | 

    var submitUsersRequest NERM.SubmitUsersRequest
    if err := json.Unmarshal(submitusersrequest, &submitUsersRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UsersAPI.SubmitUsers(context.Background()).SubmitUsersRequest(submitUsersRequest).Execute()
	  //resp, r, err := apiClient.NERM.UsersAPI.SubmitUsers(context.Background()).SubmitUsersRequest(submitUsersRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SubmitUsers``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUsers`: SubmitUsers200Response
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SubmitUsers`: %v\n", resp)
}
```

[[Back to top]](#)

