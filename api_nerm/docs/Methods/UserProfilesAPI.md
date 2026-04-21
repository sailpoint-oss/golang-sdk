---
id: nerm-user-profiles
title: UserProfiles
pagination_label: UserProfiles
sidebar_label: UserProfiles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'UserProfiles', 'NERMUserProfiles'] 
slug: /tools/sdk/go/nerm/methods/user-profiles
tags: ['SDK', 'Software Development Kit', 'UserProfiles', 'NERMUserProfiles']
---

# UserProfilesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-user-profiles**](#create-user-profiles) | **Post** `/user_profiles` | Create multiple user-profile contributor relationships
[**delete-user-profile**](#delete-user-profile) | **Delete** `/user_profile/{id}` | Delete a user profile assignment
[**delete-user-profiles**](#delete-user-profiles) | **Delete** `/user_profiles` | Delete multiple user-profile contributor relationships
[**get-user-profile**](#get-user-profile) | **Get** `/user_profiles/{id}` | Find user-profile contributor relationship
[**get-user-profiles**](#get-user-profiles) | **Get** `/user_profiles` | Get user-profile contributor relationships
[**patch-user-profile**](#patch-user-profile) | **Patch** `/user_profiles/{id}` | Update a user-profile contributor relationship
[**patch-user-profiles**](#patch-user-profiles) | **Patch** `/user_profiles` | Update multiple user-profile contributor relationships
[**submit-user-profile**](#submit-user-profile) | **Post** `/user_profile` | Create a user-profile contributor relationship


## create-user-profiles
Create multiple user-profile contributor relationships
Create multiple user-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-user-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserProfilesRequest** | [**CreateUserProfilesRequest**](../models/create-user-profiles-request) |  | 

### Return type

[**CreateUserProfiles200Response**](../models/create-user-profiles200-response)

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
    createuserprofilesrequest := []byte(``) // CreateUserProfilesRequest | 

    var createUserProfilesRequest NERM.CreateUserProfilesRequest
    if err := json.Unmarshal(createuserprofilesrequest, &createUserProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.CreateUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.CreateUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.CreateUserProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserProfiles`: CreateUserProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.CreateUserProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-user-profile
Delete a user profile assignment
Delete a user profile assignment

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-user-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserProfileRequest struct via the builder pattern


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
    resp, r, err := apiClient.NERM.UserProfilesAPI.DeleteUserProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.DeleteUserProfile(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.DeleteUserProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.DeleteUserProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-user-profiles
Delete multiple user-profile contributor relationships
Delete multiple user-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-user-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserProfilesRequest** | [**CreateUserProfilesRequest**](../models/create-user-profiles-request) |  | 

### Return type

[**CreateUserProfiles200Response**](../models/create-user-profiles200-response)

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
    createuserprofilesrequest := []byte(``) // CreateUserProfilesRequest | 

    var createUserProfilesRequest NERM.CreateUserProfilesRequest
    if err := json.Unmarshal(createuserprofilesrequest, &createUserProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.DeleteUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.DeleteUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.DeleteUserProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserProfiles`: CreateUserProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.DeleteUserProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-profile
Find user-profile contributor relationship
Find user-profile contributor relationship by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitUserProfile200Response**](../models/submit-user-profile200-response)

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
    resp, r, err := apiClient.NERM.UserProfilesAPI.GetUserProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.GetUserProfile(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.GetUserProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserProfile`: SubmitUserProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.GetUserProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## get-user-profiles
Get user-profile contributor relationships
Get user-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-user-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **userId** | **string** | The ID of a user for filtering | 
 **neAttributeId** | **string** | ID of an attribute for filtering | 
 **profileId** | **string** | Profile ID to filter by | 
 **relationshipType** | **string** | Type of user contributor relationship to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetUserProfiles200Response**](../models/get-user-profiles200-response)

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
    neAttributeId := `33f072dd-13b4-41e1-8ea0-16f2a59b57c8` // string | ID of an attribute for filtering (optional) # string | ID of an attribute for filtering (optional)
    profileId := `4e480441-451d-47d9-87c2-9a0f0fe135eb` // string | Profile ID to filter by (optional) # string | Profile ID to filter by (optional)
    relationshipType := `contributor` // string | Type of user contributor relationship to filter by (optional) # string | Type of user contributor relationship to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.GetUserProfiles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.GetUserProfiles(context.Background()).Limit(limit).Offset(offset).Order(order).UserId(userId).NeAttributeId(neAttributeId).ProfileId(profileId).RelationshipType(relationshipType).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.GetUserProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserProfiles`: GetUserProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.GetUserProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-profile
Update a user-profile contributor relationship
Update a user-profile contributor relationship by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitUserProfileRequest** | [**SubmitUserProfileRequest**](../models/submit-user-profile-request) |  | 

### Return type

[**SubmitUserProfile200Response**](../models/submit-user-profile200-response)

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
    submituserprofilerequest := []byte(``) // SubmitUserProfileRequest | 

    var submitUserProfileRequest NERM.SubmitUserProfileRequest
    if err := json.Unmarshal(submituserprofilerequest, &submitUserProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.PatchUserProfile(context.Background(), id).SubmitUserProfileRequest(submitUserProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.PatchUserProfile(context.Background(), id).SubmitUserProfileRequest(submitUserProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.PatchUserProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserProfile`: SubmitUserProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.PatchUserProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-user-profiles
Update multiple user-profile contributor relationships
Update multiple user-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-user-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchUserProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserProfilesRequest** | [**CreateUserProfilesRequest**](../models/create-user-profiles-request) |  | 

### Return type

[**CreateUserProfiles200Response**](../models/create-user-profiles200-response)

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
    createuserprofilesrequest := []byte(``) // CreateUserProfilesRequest | 

    var createUserProfilesRequest NERM.CreateUserProfilesRequest
    if err := json.Unmarshal(createuserprofilesrequest, &createUserProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.PatchUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.PatchUserProfiles(context.Background()).CreateUserProfilesRequest(createUserProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.PatchUserProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUserProfiles`: CreateUserProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.PatchUserProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-user-profile
Create a user-profile contributor relationship
Create a user-profile contributor relationship

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-user-profile)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitUserProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitUserProfileRequest** | [**SubmitUserProfileRequest**](../models/submit-user-profile-request) |  | 

### Return type

[**SubmitUserProfile200Response**](../models/submit-user-profile200-response)

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
    submituserprofilerequest := []byte(``) // SubmitUserProfileRequest | 

    var submitUserProfileRequest NERM.SubmitUserProfileRequest
    if err := json.Unmarshal(submituserprofilerequest, &submitUserProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.UserProfilesAPI.SubmitUserProfile(context.Background()).SubmitUserProfileRequest(submitUserProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.UserProfilesAPI.SubmitUserProfile(context.Background()).SubmitUserProfileRequest(submitUserProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `UserProfilesAPI.SubmitUserProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitUserProfile`: SubmitUserProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `UserProfilesAPI.SubmitUserProfile`: %v\n", resp)
}
```

[[Back to top]](#)

