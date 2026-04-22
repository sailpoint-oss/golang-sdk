---
id: nerm-profiles
title: Profiles
pagination_label: Profiles
sidebar_label: Profiles
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Profiles', 'NERMProfiles'] 
slug: /tools/sdk/go/nerm/methods/profiles
tags: ['SDK', 'Software Development Kit', 'Profiles', 'NERMProfiles']
---

# ProfilesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-profiles**](#create-profiles) | **Post** `/profiles` | Create multiple profiles
[**delete-profile-by-id**](#delete-profile-by-id) | **Delete** `/profiles/{id}` | Delete a single profile
[**delete-profiles**](#delete-profiles) | **Delete** `/profiles` | Delete multiple profiles
[**get-profile-avatar**](#get-profile-avatar) | **Get** `/profiles/{id}/avatar` | Retrieves profile avatar URL
[**get-profile-by-id**](#get-profile-by-id) | **Get** `/profiles/{id}` | Find profile by id
[**get-profile-upload**](#get-profile-upload) | **Get** `/profiles/{id}/upload/{attribute_id}` | Retrieves profile attribute attachment URL
[**get-profiles**](#get-profiles) | **Get** `/profiles` | Get profiles
[**patch-profile-by-id**](#patch-profile-by-id) | **Patch** `/profiles/{id}` | Update a profile by id
[**patch-profiles**](#patch-profiles) | **Patch** `/profiles` | Update multiple profiles
[**submit-profile**](#submit-profile) | **Post** `/profile` | Create a profile
[**submit-profile-avatar**](#submit-profile-avatar) | **Post** `/profiles/{id}/avatar` | Uploads new profile avatar
[**submit-profile-upload**](#submit-profile-upload) | **Post** `/profiles/{id}/upload/{attribute_id}` | Uploads profile attachment attribute


## create-profiles
Create multiple profiles
Create multiple profiles

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfilesRequest** | [**CreateProfilesRequest**](../models/create-profiles-request) |  | 

### Return type

[**SearchAdvancedSearch200Response**](../models/search-advanced-search200-response)

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
    createprofilesrequest := []byte(``) // CreateProfilesRequest | 

    var createProfilesRequest NERM.CreateProfilesRequest
    if err := json.Unmarshal(createprofilesrequest, &createProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.CreateProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.CreateProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.CreateProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProfiles`: SearchAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.CreateProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-profile-by-id
Delete a single profile
Delete a single profile

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-profile-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

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
    r, err := apiClient.NERM.ProfilesAPI.DeleteProfileById(context.Background(), id).Execute()
	  //r, err := apiClient.NERM.ProfilesAPI.DeleteProfileById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteProfileById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## delete-profiles
Delete multiple profiles
Delete multiple profiles

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfilesRequest** | [**CreateProfilesRequest**](../models/create-profiles-request) |  | 

### Return type

[**DeleteProfiles200Response**](../models/delete-profiles200-response)

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
    createprofilesrequest := []byte(``) // CreateProfilesRequest | 

    var createProfilesRequest NERM.CreateProfilesRequest
    if err := json.Unmarshal(createprofilesrequest, &createProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.DeleteProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.DeleteProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.DeleteProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProfiles`: DeleteProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.DeleteProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-avatar
Retrieves profile avatar URL
Retrieves the URL of the profile avatar

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-avatar)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileAvatarRequest struct via the builder pattern


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
    resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileAvatar(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileAvatar(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileAvatar``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileAvatar`: Url
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileAvatar`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-by-id
Find profile by id
Find profile by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSingleSchemaMappedProfile200Response**](../models/get-single-schema-mapped-profile200-response)

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
    resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileById`: GetSingleSchemaMappedProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-upload
Retrieves profile attribute attachment URL
Retrieves the URL of an attachment attribute value from a profile

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-upload)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 
**attributeId** | **string** | The id of the attachment attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileUploadRequest struct via the builder pattern


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
    attributeId := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | The id of the attachment attribute # string | The id of the attachment attribute

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileUpload(context.Background(), id, attributeId).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.GetProfileUpload(context.Background(), id, attributeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfileUpload``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileUpload`: Url
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfileUpload`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profiles
Get profiles
Get profiles

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **excludeAttributes** | **bool** | Allows for optimization by not returning the associated attribute data for the returned profiles | [default to false]
 **name** | **string** | object name for filtering | 
 **profileTypeId** | **string** | Profile Type ID for filtering | 
 **status** | **string** | status value for filtering | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]
 **afterId** | **string** | Represents the ID where the query should begin from. If blank, it represents the first ID. When used, forces sorting by ID ascending and does not allow use of &#x60;offset&#x60;. When &#x60;after_id&#x60; is specified it changes the mode of the API such that any filter parameters other than &#x60;profile_type_id&#x60;, &#x60;limit&#x60;, and &#x60;offset&#x60; are not supported and will be either silently ignored or result in an HTTP 400 error. For example you can not include an &#x60;after_id&#x60; along with an &#x60;archived&#x3D;false&#x60; in the same request. Can be used alongside &#x60;metadata&#x60; parameter. | 
 **updatedAfter** | **string** | Adds support for filtering profiles based on the date of the latest modification made on them. Can be used alongside the after_id parameter. | 

### Return type

[**GetSchemaMappedProfilesCollection200Response**](../models/get-schema-mapped-profiles-collection200-response)

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
    excludeAttributes := false // bool | Allows for optimization by not returning the associated attribute data for the returned profiles (optional) (default to false) # bool | Allows for optimization by not returning the associated attribute data for the returned profiles (optional) (default to false)
    name := `name` // string | object name for filtering (optional) # string | object name for filtering (optional)
    profileTypeId := `79ed1cb6-9977-4965-9bfe-f2bcc242523e` // string | Profile Type ID for filtering (optional) # string | Profile Type ID for filtering (optional)
    status := `Active` // string | status value for filtering (optional) # string | status value for filtering (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)
    afterId := `4eaa719f-4312-4c5b-9264-d0eb04d4a02a` // string | Represents the ID where the query should begin from. If blank, it represents the first ID. When used, forces sorting by ID ascending and does not allow use of `offset`. When `after_id` is specified it changes the mode of the API such that any filter parameters other than `profile_type_id`, `limit`, and `offset` are not supported and will be either silently ignored or result in an HTTP 400 error. For example you can not include an `after_id` along with an `archived=false` in the same request. Can be used alongside `metadata` parameter. (optional) # string | Represents the ID where the query should begin from. If blank, it represents the first ID. When used, forces sorting by ID ascending and does not allow use of `offset`. When `after_id` is specified it changes the mode of the API such that any filter parameters other than `profile_type_id`, `limit`, and `offset` are not supported and will be either silently ignored or result in an HTTP 400 error. For example you can not include an `after_id` along with an `archived=false` in the same request. Can be used alongside `metadata` parameter. (optional)
    updatedAfter := Mon May 05 00:00:00 UTC 2025 // string | Adds support for filtering profiles based on the date of the latest modification made on them. Can be used alongside the after_id parameter. (optional) # string | Adds support for filtering profiles based on the date of the latest modification made on them. Can be used alongside the after_id parameter. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.GetProfiles(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.GetProfiles(context.Background()).Limit(limit).Offset(offset).Order(order).ExcludeAttributes(excludeAttributes).Name(name).ProfileTypeId(profileTypeId).Status(status).Metadata(metadata).AfterId(afterId).UpdatedAfter(updatedAfter).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.GetProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfiles`: GetSchemaMappedProfilesCollection200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.GetProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-profile-by-id
Update a profile by id
Update a profile by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-profile-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProfileByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitProfileRequest** | [**SubmitProfileRequest**](../models/submit-profile-request) |  | 

### Return type

[**GetSingleSchemaMappedProfile200Response**](../models/get-single-schema-mapped-profile200-response)

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
    submitprofilerequest := []byte(``) // SubmitProfileRequest | 

    var submitProfileRequest NERM.SubmitProfileRequest
    if err := json.Unmarshal(submitprofilerequest, &submitProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.PatchProfileById(context.Background(), id).SubmitProfileRequest(submitProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.PatchProfileById(context.Background(), id).SubmitProfileRequest(submitProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.PatchProfileById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProfileById`: GetSingleSchemaMappedProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.PatchProfileById`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-profiles
Update multiple profiles
Update multiple profiles

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-profiles)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProfilesRequest** | [**CreateProfilesRequest**](../models/create-profiles-request) |  | 

### Return type

[**SearchAdvancedSearch200Response**](../models/search-advanced-search200-response)

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
    createprofilesrequest := []byte(``) // CreateProfilesRequest | 

    var createProfilesRequest NERM.CreateProfilesRequest
    if err := json.Unmarshal(createprofilesrequest, &createProfilesRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.PatchProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.PatchProfiles(context.Background()).CreateProfilesRequest(createProfilesRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.PatchProfiles``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProfiles`: SearchAdvancedSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.PatchProfiles`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-profile
Create a profile
Create a profile

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-profile)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitProfileRequest** | [**SubmitProfileRequest**](../models/submit-profile-request) |  | 

### Return type

[**GetSingleSchemaMappedProfile200Response**](../models/get-single-schema-mapped-profile200-response)

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
    submitprofilerequest := []byte(``) // SubmitProfileRequest | 

    var submitProfileRequest NERM.SubmitProfileRequest
    if err := json.Unmarshal(submitprofilerequest, &submitProfileRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfile(context.Background()).SubmitProfileRequest(submitProfileRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfile(context.Background()).SubmitProfileRequest(submitProfileRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SubmitProfile``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitProfile`: GetSingleSchemaMappedProfile200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.SubmitProfile`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-profile-avatar
Uploads new profile avatar
Uploads a new profile avatar

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-profile-avatar)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitProfileAvatarRequest struct via the builder pattern


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
    resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfileAvatar(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfileAvatar(context.Background(), id).File(file).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SubmitProfileAvatar``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitProfileAvatar`: Url
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.SubmitProfileAvatar`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-profile-upload
Uploads profile attachment attribute
Uploads a new attachment attribute value to a profile. The upload must be a FORM data type; this is not a JSON API. The upload must include the binary content of the payload under the 'file' named form element. The upload must not attempt to include the file name or its content type as a other form or JSON as parameters. The upload must not attempt to upload the file body as the POST body payload; it has to arrive as a FORM parameter. Do not use a `File/Binary` payload type for the POST operation in your API client.


[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-profile-upload)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 
**attributeId** | **string** | The id of the attachment attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitProfileUploadRequest struct via the builder pattern


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
    attributeId := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | The id of the attachment attribute # string | The id of the attachment attribute
    file := BINARY_DATA_HERE // *os.File |  (optional) # *os.File |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfileUpload(context.Background(), id, attributeId).Execute()
	  //resp, r, err := apiClient.NERM.ProfilesAPI.SubmitProfileUpload(context.Background(), id, attributeId).File(file).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.SubmitProfileUpload``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitProfileUpload`: Url
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.SubmitProfileUpload`: %v\n", resp)
}
```

[[Back to top]](#)

