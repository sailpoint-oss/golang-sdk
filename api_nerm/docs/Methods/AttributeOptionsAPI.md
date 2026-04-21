---
id: nerm-attribute-options
title: AttributeOptions
pagination_label: AttributeOptions
sidebar_label: AttributeOptions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AttributeOptions', 'NERMAttributeOptions'] 
slug: /tools/sdk/go/nerm/methods/attribute-options
tags: ['SDK', 'Software Development Kit', 'AttributeOptions', 'NERMAttributeOptions']
---

# AttributeOptionsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-attribute-option-by-id**](#delete-attribute-option-by-id) | **Delete** `/ne_attribute_options/{id}` | Delete option based attribute value
[**delete-attribute-option-by-uid**](#delete-attribute-option-by-uid) | **Delete** `/ne_attribute_options/{uid}` | Delete option value
[**get-attribute-option-by-id**](#get-attribute-option-by-id) | **Get** `/ne_attribute_options/{id}` | Find option based attribute value
[**get-attribute-option-by-uid**](#get-attribute-option-by-uid) | **Get** `/ne_attribute_options/{uid}` | Find option attribute value
[**get-attribute-options**](#get-attribute-options) | **Get** `/ne_attribute_options` | Get option based attribute values
[**patch-attribute-option-by-id**](#patch-attribute-option-by-id) | **Patch** `/ne_attribute_options/{id}` | Update option based attribute value
[**patch-attribute-option-by-uid**](#patch-attribute-option-by-uid) | **Patch** `/ne_attribute_options/{uid}` | Update option value
[**patch-attribute-options**](#patch-attribute-options) | **Patch** `/ne_attribute_options` | Update multiple option values
[**submit-attribute-option**](#submit-attribute-option) | **Post** `/ne_attribute_option` | Add value to option
[**submit-attribute-options**](#submit-attribute-options) | **Post** `/ne_attribute_options` | Create multiple option values


## delete-attribute-option-by-id
Delete option based attribute value
Delete a option based attribute value by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-attribute-option-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAttributeOptionById200Response**](../models/delete-attribute-option-by-id200-response)

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
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.DeleteAttributeOptionById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.DeleteAttributeOptionById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.DeleteAttributeOptionById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAttributeOptionById`: DeleteAttributeOptionById200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.DeleteAttributeOptionById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-attribute-option-by-uid
Delete option value
Delete a option based attribute value by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-attribute-option-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeOptionByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteAttributeOptionById200Response**](../models/delete-attribute-option-by-id200-response)

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
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.DeleteAttributeOptionByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.DeleteAttributeOptionByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.DeleteAttributeOptionByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAttributeOptionByUid`: DeleteAttributeOptionById200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.DeleteAttributeOptionByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attribute-option-by-id
Find option based attribute value
Info for a specific option based attribute value by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attribute-option-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitAttributeOption200Response**](../models/submit-attribute-option200-response)

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
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptionById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptionById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.GetAttributeOptionById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOptionById`: SubmitAttributeOption200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.GetAttributeOptionById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attribute-option-by-uid
Find option attribute value
Get a specific option based attribute value by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attribute-option-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitAttributeOption200Response**](../models/submit-attribute-option200-response)

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
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptionByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptionByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.GetAttributeOptionByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOptionByUid`: SubmitAttributeOption200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.GetAttributeOptionByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-attribute-options
Get option based attribute values
Get option based attribute values

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-attribute-options)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **neAttributeId** | **string** | ID of an attribute for filtering | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetAttributeOptions200Response**](../models/get-attribute-options200-response)

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
    neAttributeId := `33f072dd-13b4-41e1-8ea0-16f2a59b57c8` // string | ID of an attribute for filtering (optional) # string | ID of an attribute for filtering (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptions(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.GetAttributeOptions(context.Background()).Limit(limit).Offset(offset).Order(order).NeAttributeId(neAttributeId).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.GetAttributeOptions``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributeOptions`: GetAttributeOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.GetAttributeOptions`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-attribute-option-by-id
Update option based attribute value
Update a option based attribute value by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-attribute-option-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAttributeOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitAttributeOptionRequest** | [**SubmitAttributeOptionRequest**](../models/submit-attribute-option-request) |  | 

### Return type

[**SubmitAttributeOption200Response**](../models/submit-attribute-option200-response)

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
    submitattributeoptionrequest := []byte(``) // SubmitAttributeOptionRequest | 

    var submitAttributeOptionRequest NERM.SubmitAttributeOptionRequest
    if err := json.Unmarshal(submitattributeoptionrequest, &submitAttributeOptionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptionById(context.Background(), id).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptionById(context.Background(), id).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.PatchAttributeOptionById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAttributeOptionById`: SubmitAttributeOption200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.PatchAttributeOptionById`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-attribute-option-by-uid
Update option value
Update a option based attribute value by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-attribute-option-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAttributeOptionByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAttributeOptionRequest** | [**SubmitAttributeOptionRequest**](../models/submit-attribute-option-request) |  | 


### Return type

[**SubmitAttributeOption200Response**](../models/submit-attribute-option200-response)

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
    submitattributeoptionrequest := []byte(``) // SubmitAttributeOptionRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var submitAttributeOptionRequest NERM.SubmitAttributeOptionRequest
    if err := json.Unmarshal(submitattributeoptionrequest, &submitAttributeOptionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptionByUid(context.Background(), uid).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptionByUid(context.Background(), uid).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.PatchAttributeOptionByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAttributeOptionByUid`: SubmitAttributeOption200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.PatchAttributeOptionByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-attribute-options
Update multiple option values
Update multiple option based attribute values

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-attribute-options)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAttributeOptionsRequest** | [**SubmitAttributeOptionsRequest**](../models/submit-attribute-options-request) |  | 

### Return type

[**SubmitAttributeOptions200Response**](../models/submit-attribute-options200-response)

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
    submitattributeoptionsrequest := []byte(``) // SubmitAttributeOptionsRequest | 

    var submitAttributeOptionsRequest NERM.SubmitAttributeOptionsRequest
    if err := json.Unmarshal(submitattributeoptionsrequest, &submitAttributeOptionsRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptions(context.Background()).SubmitAttributeOptionsRequest(submitAttributeOptionsRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.PatchAttributeOptions(context.Background()).SubmitAttributeOptionsRequest(submitAttributeOptionsRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.PatchAttributeOptions``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAttributeOptions`: SubmitAttributeOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.PatchAttributeOptions`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-attribute-option
Add value to option
Adds a value to an option based attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-attribute-option)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAttributeOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAttributeOptionRequest** | [**SubmitAttributeOptionRequest**](../models/submit-attribute-option-request) |  | 

### Return type

[**SubmitAttributeOption200Response**](../models/submit-attribute-option200-response)

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
    submitattributeoptionrequest := []byte(``) // SubmitAttributeOptionRequest | 

    var submitAttributeOptionRequest NERM.SubmitAttributeOptionRequest
    if err := json.Unmarshal(submitattributeoptionrequest, &submitAttributeOptionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.SubmitAttributeOption(context.Background()).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.SubmitAttributeOption(context.Background()).SubmitAttributeOptionRequest(submitAttributeOptionRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.SubmitAttributeOption``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAttributeOption`: SubmitAttributeOption200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.SubmitAttributeOption`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-attribute-options
Create multiple option values
Create multiple new option based attribute values

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-attribute-options)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitAttributeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitAttributeOptionsRequest** | [**SubmitAttributeOptionsRequest**](../models/submit-attribute-options-request) |  | 

### Return type

[**SubmitAttributeOptions200Response**](../models/submit-attribute-options200-response)

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
    submitattributeoptionsrequest := []byte(``) // SubmitAttributeOptionsRequest | 

    var submitAttributeOptionsRequest NERM.SubmitAttributeOptionsRequest
    if err := json.Unmarshal(submitattributeoptionsrequest, &submitAttributeOptionsRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AttributeOptionsAPI.SubmitAttributeOptions(context.Background()).SubmitAttributeOptionsRequest(submitAttributeOptionsRequest).Execute()
	  //resp, r, err := apiClient.NERM.AttributeOptionsAPI.SubmitAttributeOptions(context.Background()).SubmitAttributeOptionsRequest(submitAttributeOptionsRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AttributeOptionsAPI.SubmitAttributeOptions``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitAttributeOptions`: SubmitAttributeOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `AttributeOptionsAPI.SubmitAttributeOptions`: %v\n", resp)
}
```

[[Back to top]](#)

