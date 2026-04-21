---
id: nerm-role-profiles
title: RoleProfiles
pagination_label: RoleProfiles
sidebar_label: RoleProfiles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'RoleProfiles', 'NERMRoleProfiles'] 
slug: /tools/sdk/go/nerm/methods/role-profiles
tags: ['SDK', 'Software Development Kit', 'RoleProfiles', 'NERMRoleProfiles']
---

# RoleProfilesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-role-profile**](#delete-role-profile) | **Delete** `/role_profile/{id}` | Delete a role profile assignment
[**get-role-profile**](#get-role-profile) | **Get** `/role_profiles/{id}` | Find role-profile contributor relationship
[**get-role-profiles**](#get-role-profiles) | **Get** `/role_profiles` | Get role-profile contributor relationships
[**patch-role-profile**](#patch-role-profile) | **Patch** `/role_profiles/{id}` | Update a role-profile contributor relationship
[**patch-role-profiles**](#patch-role-profiles) | **Patch** `/role_profiles` | Update multiple role-profile contributor relationships
[**submit-role-profile**](#submit-role-profile) | **Post** `/role_profile` | Create a role-profile contributor relationship
[**submit-role-profiles**](#submit-role-profiles) | **Post** `/role_profiles` | Create multiple role-profile contributor relationships


## delete-role-profile
Delete a role profile assignment
Delete a role profile assignment

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-role-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleProfileRequest struct via the builder pattern


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
    resp, r, err := apiClient.NERM.RoleProfilesAPI.DeleteRoleProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.DeleteRoleProfile(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.DeleteRoleProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoleProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.DeleteRoleProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## get-role-profile
Find role-profile contributor relationship
Find role-profile contributor relationship by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-role-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitRoleProfile200Response**](../models/submit-role-profile200-response)

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
    resp, r, err := apiClient.NERM.RoleProfilesAPI.GetRoleProfile(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.GetRoleProfile(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.GetRoleProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleProfile`: SubmitRoleProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.GetRoleProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## get-role-profiles
Get role-profile contributor relationships
Get role-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-role-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **roleId** | **string** | The ID of a role for filtering | 
 **profileId** | **string** | Profile ID to filter by | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetRoleProfiles200Response**](../models/get-role-profiles200-response)

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
    roleId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | The ID of a role for filtering (optional) # string | The ID of a role for filtering (optional)
    profileId := `4e480441-451d-47d9-87c2-9a0f0fe135eb` // string | Profile ID to filter by (optional) # string | Profile ID to filter by (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RoleProfilesAPI.GetRoleProfiles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.GetRoleProfiles(context.Background()).Limit(limit).Offset(offset).Order(order).RoleId(roleId).ProfileId(profileId).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.GetRoleProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleProfiles`: GetRoleProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.GetRoleProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-role-profile
Update a role-profile contributor relationship
Update a role-profile contributor relationship by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-role-profile)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoleProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitRoleProfileRequest** | [**SubmitRoleProfileRequest**](../models/submit-role-profile-request) |  | 

### Return type

[**SubmitRoleProfile200Response**](../models/submit-role-profile200-response)

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
    submitroleprofilerequest := []byte(``) // SubmitRoleProfileRequest | 

    var submitRoleProfileRequest NERM.SubmitRoleProfileRequest
    if err := json.Unmarshal(submitroleprofilerequest, &submitRoleProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RoleProfilesAPI.PatchRoleProfile(context.Background(), id).SubmitRoleProfileRequest(submitRoleProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.PatchRoleProfile(context.Background(), id).SubmitRoleProfileRequest(submitRoleProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.PatchRoleProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRoleProfile`: SubmitRoleProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.PatchRoleProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-role-profiles
Update multiple role-profile contributor relationships
Update multiple role-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-role-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRoleProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRoleProfilesRequest** | [**SubmitRoleProfilesRequest**](../models/submit-role-profiles-request) |  | 

### Return type

[**SubmitRoleProfiles200Response**](../models/submit-role-profiles200-response)

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
    submitroleprofilesrequest := []byte(``) // SubmitRoleProfilesRequest | 

    var submitRoleProfilesRequest NERM.SubmitRoleProfilesRequest
    if err := json.Unmarshal(submitroleprofilesrequest, &submitRoleProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RoleProfilesAPI.PatchRoleProfiles(context.Background()).SubmitRoleProfilesRequest(submitRoleProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.PatchRoleProfiles(context.Background()).SubmitRoleProfilesRequest(submitRoleProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.PatchRoleProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRoleProfiles`: SubmitRoleProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.PatchRoleProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-role-profile
Create a role-profile contributor relationship
Create a role-profile contributor relationship

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-role-profile)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRoleProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRoleProfileRequest** | [**SubmitRoleProfileRequest**](../models/submit-role-profile-request) |  | 

### Return type

[**SubmitRoleProfile200Response**](../models/submit-role-profile200-response)

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
    submitroleprofilerequest := []byte(``) // SubmitRoleProfileRequest | 

    var submitRoleProfileRequest NERM.SubmitRoleProfileRequest
    if err := json.Unmarshal(submitroleprofilerequest, &submitRoleProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RoleProfilesAPI.SubmitRoleProfile(context.Background()).SubmitRoleProfileRequest(submitRoleProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.SubmitRoleProfile(context.Background()).SubmitRoleProfileRequest(submitRoleProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.SubmitRoleProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRoleProfile`: SubmitRoleProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.SubmitRoleProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-role-profiles
Create multiple role-profile contributor relationships
Create multiple role-profile contributor relationships

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-role-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRoleProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitRoleProfilesRequest** | [**SubmitRoleProfilesRequest**](../models/submit-role-profiles-request) |  | 

### Return type

[**SubmitRoleProfiles200Response**](../models/submit-role-profiles200-response)

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
    submitroleprofilesrequest := []byte(``) // SubmitRoleProfilesRequest | 

    var submitRoleProfilesRequest NERM.SubmitRoleProfilesRequest
    if err := json.Unmarshal(submitroleprofilesrequest, &submitRoleProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.RoleProfilesAPI.SubmitRoleProfiles(context.Background()).SubmitRoleProfilesRequest(submitRoleProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.RoleProfilesAPI.SubmitRoleProfiles(context.Background()).SubmitRoleProfilesRequest(submitRoleProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `RoleProfilesAPI.SubmitRoleProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitRoleProfiles`: SubmitRoleProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RoleProfilesAPI.SubmitRoleProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

