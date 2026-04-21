---
id: nerm-form-attributes
title: FormAttributes
pagination_label: FormAttributes
sidebar_label: FormAttributes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'FormAttributes', 'NERMFormAttributes'] 
slug: /tools/sdk/go/nerm/methods/form-attributes
tags: ['SDK', 'Software Development Kit', 'FormAttributes', 'NERMFormAttributes']
---

# FormAttributesAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-form-attribute**](#create-form-attribute) | **Post** `/form_attributes` | Create a form attribute
[**delete-form-attribute-by-id**](#delete-form-attribute-by-id) | **Delete** `/form_attributes/{id}` | Delete form attribute
[**delete-form-attribute-by-uid**](#delete-form-attribute-by-uid) | **Delete** `/form_attributes/{uid}` | Delete form attribute
[**get-form-attribute-by-id**](#get-form-attribute-by-id) | **Get** `/form_attributes/{id}` | Get form attribute data
[**get-form-attribute-by-uid**](#get-form-attribute-by-uid) | **Get** `/form_attributes/{uid}` | Get form attribute data
[**get-form-attributes**](#get-form-attributes) | **Get** `/form_attributes` | Get form attributes
[**update-form-attribute-by-id**](#update-form-attribute-by-id) | **Patch** `/form_attributes/{id}` | Update form attribute data
[**update-form-attribute-by-uid**](#update-form-attribute-by-uid) | **Patch** `/form_attributes/{uid}` | Update form attribute data


## create-form-attribute
Create a form attribute
This endpoint can create a form attribute

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-form-attribute)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFormAttributeRequest** | [**CreateFormAttributeRequest**](../models/create-form-attribute-request) |  | 

### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    createformattributerequest := []byte(``) // CreateFormAttributeRequest | 

    var createFormAttributeRequest NERM.CreateFormAttributeRequest
    if err := json.Unmarshal(createformattributerequest, &createFormAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormAttributesAPI.CreateFormAttribute(context.Background()).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.CreateFormAttribute(context.Background()).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.CreateFormAttribute``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFormAttribute`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.CreateFormAttribute`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-form-attribute-by-id
Delete form attribute
Delete form attribute by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-form-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    resp, r, err := apiClient.NERM.FormAttributesAPI.DeleteFormAttributeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.DeleteFormAttributeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.DeleteFormAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFormAttributeById`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.DeleteFormAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-form-attribute-by-uid
Delete form attribute
Delete form attribute by UID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-form-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    resp, r, err := apiClient.NERM.FormAttributesAPI.DeleteFormAttributeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.DeleteFormAttributeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.DeleteFormAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFormAttributeByUid`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.DeleteFormAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-form-attribute-by-id
Get form attribute data
Gets info for a specific form attribute by ID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-form-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributeById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributeById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.GetFormAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormAttributeById`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.GetFormAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-form-attribute-by-uid
Get form attribute data
Get info for a specific form attribute by UID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-form-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributeByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributeByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.GetFormAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormAttributeByUid`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.GetFormAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-form-attributes
Get form attributes
Get defined form attribute in the system

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-form-attributes)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormAttributesRequest struct via the builder pattern


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributes(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.GetFormAttributes(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.GetFormAttributes``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormAttributes`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.GetFormAttributes`: %v\n", resp)
}
```

[[Back to top]](#)

## update-form-attribute-by-id
Update form attribute data
Update info for a specific form attribute by ID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-form-attribute-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormAttributeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFormAttributeRequest** | [**CreateFormAttributeRequest**](../models/create-form-attribute-request) |  | 

### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    createformattributerequest := []byte(``) // CreateFormAttributeRequest | 

    var createFormAttributeRequest NERM.CreateFormAttributeRequest
    if err := json.Unmarshal(createformattributerequest, &createFormAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormAttributesAPI.UpdateFormAttributeById(context.Background(), id).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.UpdateFormAttributeById(context.Background(), id).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.UpdateFormAttributeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFormAttributeById`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.UpdateFormAttributeById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-form-attribute-by-uid
Update form attribute data
Update info for a specific form attribute by UID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-form-attribute-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormAttributeByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFormAttributeRequest** | [**CreateFormAttributeRequest**](../models/create-form-attribute-request) |  | 


### Return type

[**GetFormAttributes200Response**](../models/get-form-attributes200-response)

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
    createformattributerequest := []byte(``) // CreateFormAttributeRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var createFormAttributeRequest NERM.CreateFormAttributeRequest
    if err := json.Unmarshal(createformattributerequest, &createFormAttributeRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormAttributesAPI.UpdateFormAttributeByUid(context.Background(), uid).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormAttributesAPI.UpdateFormAttributeByUid(context.Background(), uid).CreateFormAttributeRequest(createFormAttributeRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormAttributesAPI.UpdateFormAttributeByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFormAttributeByUid`: GetFormAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `FormAttributesAPI.UpdateFormAttributeByUid`: %v\n", resp)
}
```

[[Back to top]](#)

