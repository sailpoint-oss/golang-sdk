---
id: nerm-profile-types
title: ProfileTypes
pagination_label: ProfileTypes
sidebar_label: ProfileTypes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'ProfileTypes', 'NERMProfileTypes'] 
slug: /tools/sdk/go/nerm/methods/profile-types
tags: ['SDK', 'Software Development Kit', 'ProfileTypes', 'NERMProfileTypes']
---

# ProfileTypesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-profile-type-by-id**](#delete-profile-type-by-id) | **Delete** `/profile_types/{id}` | Delete profile type
[**delete-profile-type-by-uid**](#delete-profile-type-by-uid) | **Delete** `/profile_types/{uid}` | Delete profile type
[**get-profile-type-by-id**](#get-profile-type-by-id) | **Get** `/profile_types/{id}` | Find profile type
[**get-profile-type-by-uid**](#get-profile-type-by-uid) | **Get** `/profile_types/{uid}` | Find profile type
[**get-profile-types**](#get-profile-types) | **Get** `/profile_types` | Get profile types
[**patch-profile-type-by-id**](#patch-profile-type-by-id) | **Patch** `/profile_types/{id}` | Update a profile type
[**patch-profile-type-by-uid**](#patch-profile-type-by-uid) | **Patch** `/profile_types/{uid}` | Update a profile type
[**submit-profile-type**](#submit-profile-type) | **Post** `/profile_type` | Create a profile type


## delete-profile-type-by-id
Delete profile type
Delete a profile type. All profiles of that type must first be destroyed before the profile type can be destroyed.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-profile-type-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileTypeByIdRequest struct via the builder pattern


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
    resp, r, err := apiClient.NERM.ProfileTypesAPI.DeleteProfileTypeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.DeleteProfileTypeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.DeleteProfileTypeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProfileTypeById`: DeleteProfileTypeById200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.DeleteProfileTypeById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-profile-type-by-uid
Delete profile type
Delete a profile type by UID (user-specified identifier). All profiles of that type must first be destroyed before the profile type can be destroyed.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-profile-type-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileTypeByUidRequest struct via the builder pattern


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
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.DeleteProfileTypeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.DeleteProfileTypeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.DeleteProfileTypeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProfileTypeByUid`: DeleteProfileTypeById200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.DeleteProfileTypeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-type-by-id
Find profile type
Find profile type by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-type-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitProfileType200Response**](../models/submit-profile-type200-response)

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
    resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.GetProfileTypeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileTypeById`: SubmitProfileType200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.GetProfileTypeById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-type-by-uid
Find profile type
Find profile type by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-type-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileTypeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitProfileType200Response**](../models/submit-profile-type200-response)

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
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.GetProfileTypeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileTypeByUid`: SubmitProfileType200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.GetProfileTypeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-profile-types
Get profile types
Get option based attribute values

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-profile-types)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **name** | **string** | object name for filtering | 
 **archived** | **bool** | Filter by archive status | [default to false]
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetProfileTypes200Response**](../models/get-profile-types200-response)

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
    archived := false // bool | Filter by archive status (optional) (default to false) # bool | Filter by archive status (optional) (default to false)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypes(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.GetProfileTypes(context.Background()).Limit(limit).Offset(offset).Order(order).Name(name).Archived(archived).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.GetProfileTypes``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileTypes`: GetProfileTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.GetProfileTypes`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-profile-type-by-id
Update a profile type
Update a profile type by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-profile-type-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProfileTypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitProfileTypeRequest** | [**SubmitProfileTypeRequest**](../models/submit-profile-type-request) |  | 

### Return type

[**SubmitProfileType200Response**](../models/submit-profile-type200-response)

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
    submitprofiletyperequest := []byte(``) // SubmitProfileTypeRequest | 

    var submitProfileTypeRequest NERM.SubmitProfileTypeRequest
    if err := json.Unmarshal(submitprofiletyperequest, &submitProfileTypeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.PatchProfileTypeById(context.Background(), id).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.PatchProfileTypeById(context.Background(), id).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.PatchProfileTypeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProfileTypeById`: SubmitProfileType200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.PatchProfileTypeById`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-profile-type-by-uid
Update a profile type
Update a profile type by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-profile-type-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchProfileTypeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitProfileTypeRequest** | [**SubmitProfileTypeRequest**](../models/submit-profile-type-request) |  | 


### Return type

[**SubmitProfileType200Response**](../models/submit-profile-type200-response)

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
    submitprofiletyperequest := []byte(``) // SubmitProfileTypeRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var submitProfileTypeRequest NERM.SubmitProfileTypeRequest
    if err := json.Unmarshal(submitprofiletyperequest, &submitProfileTypeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.PatchProfileTypeByUid(context.Background(), uid).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.PatchProfileTypeByUid(context.Background(), uid).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.PatchProfileTypeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchProfileTypeByUid`: SubmitProfileType200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.PatchProfileTypeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-profile-type
Create a profile type
Create a profile type

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-profile-type)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitProfileTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitProfileTypeRequest** | [**SubmitProfileTypeRequest**](../models/submit-profile-type-request) |  | 

### Return type

[**SubmitProfileType200Response**](../models/submit-profile-type200-response)

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
    submitprofiletyperequest := []byte(``) // SubmitProfileTypeRequest | 

    var submitProfileTypeRequest NERM.SubmitProfileTypeRequest
    if err := json.Unmarshal(submitprofiletyperequest, &submitProfileTypeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.ProfileTypesAPI.SubmitProfileType(context.Background()).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
	  //resp, r, err := apiClient.NERM.ProfileTypesAPI.SubmitProfileType(context.Background()).SubmitProfileTypeRequest(submitProfileTypeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ProfileTypesAPI.SubmitProfileType``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitProfileType`: SubmitProfileType200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfileTypesAPI.SubmitProfileType`: %v\n", resp)
}
```

[[Back to top]](#)

