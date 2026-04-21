---
id: nerm-page-elements
title: PageElements
pagination_label: PageElements
sidebar_label: PageElements
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageElements', 'NERMPageElements'] 
slug: /tools/sdk/go/nerm/methods/page-elements
tags: ['SDK', 'Software Development Kit', 'PageElements', 'NERMPageElements']
---

# PageElementsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-page-element**](#create-page-element) | **Post** `/page_elements` | Create a page element entry
[**delete-page-element-by-id**](#delete-page-element-by-id) | **Delete** `/page_elements/{id}` | Delete page element
[**delete-page-element-by-uid**](#delete-page-element-by-uid) | **Delete** `/page_elements/{uid}` | Delete page element
[**get-page-element-by-id**](#get-page-element-by-id) | **Get** `/page_elements/{id}` | Find a page element
[**get-page-element-by-uid**](#get-page-element-by-uid) | **Get** `/page_elements/{uid}` | Find page element
[**get-page-elements**](#get-page-elements) | **Get** `/page_elements` | Get page element data
[**update-page-element-by-id**](#update-page-element-by-id) | **Patch** `/page_elements/{id}` | Update page element
[**update-page-element-by-uid**](#update-page-element-by-uid) | **Patch** `/page_elements/{uid}` | Update page element


## create-page-element
Create a page element entry
Creates a page element.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-page-element)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePageElementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageElementRequest** | [**CreatePageElementRequest**](../models/create-page-element-request) |  | 

### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    createpageelementrequest := []byte(``) // CreatePageElementRequest | 

    var createPageElementRequest NERM.CreatePageElementRequest
    if err := json.Unmarshal(createpageelementrequest, &createPageElementRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageElementsAPI.CreatePageElement(context.Background()).CreatePageElementRequest(createPageElementRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.CreatePageElement(context.Background()).CreatePageElementRequest(createPageElementRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.CreatePageElement``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePageElement`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.CreatePageElement`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-element-by-id
Delete page element
Delete page element by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-element-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageElementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    resp, r, err := apiClient.NERM.PageElementsAPI.DeletePageElementById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.DeletePageElementById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.DeletePageElementById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageElementById`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.DeletePageElementById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-element-by-uid
Delete page element
Delete page element by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-element-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageElementByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    resp, r, err := apiClient.NERM.PageElementsAPI.DeletePageElementByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.DeletePageElementByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.DeletePageElementByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageElementByUid`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.DeletePageElementByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-element-by-id
Find a page element
Info for a specific page element record

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-element-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageElementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElementById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElementById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.GetPageElementById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageElementById`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.GetPageElementById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-element-by-uid
Find page element
Info for a specific page element record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-element-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageElementByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElementByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElementByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.GetPageElementByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageElementByUid`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.GetPageElementByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-elements
Get page element data
Retrieves page elements data.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-elements)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageElementsRequest struct via the builder pattern


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElements(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.GetPageElements(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.GetPageElements``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageElements`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.GetPageElements`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-element-by-id
Update page element
Update info for a specific page element record by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-element-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageElementByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPageElementRequest** | [**CreatePageElementRequest**](../models/create-page-element-request) |  | 

### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    createpageelementrequest := []byte(``) // CreatePageElementRequest | 

    var createPageElementRequest NERM.CreatePageElementRequest
    if err := json.Unmarshal(createpageelementrequest, &createPageElementRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageElementsAPI.UpdatePageElementById(context.Background(), id).CreatePageElementRequest(createPageElementRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.UpdatePageElementById(context.Background(), id).CreatePageElementRequest(createPageElementRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.UpdatePageElementById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageElementById`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.UpdatePageElementById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-element-by-uid
Update page element
Update info for a specific page element record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-element-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageElementByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageElementRequest** | [**CreatePageElementRequest**](../models/create-page-element-request) |  | 


### Return type

[**GetPageElements200Response**](../models/get-page-elements200-response)

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
    createpageelementrequest := []byte(``) // CreatePageElementRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var createPageElementRequest NERM.CreatePageElementRequest
    if err := json.Unmarshal(createpageelementrequest, &createPageElementRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageElementsAPI.UpdatePageElementByUid(context.Background(), uid).CreatePageElementRequest(createPageElementRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageElementsAPI.UpdatePageElementByUid(context.Background(), uid).CreatePageElementRequest(createPageElementRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageElementsAPI.UpdatePageElementByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageElementByUid`: GetPageElements200Response
    fmt.Fprintf(os.Stdout, "Response from `PageElementsAPI.UpdatePageElementByUid`: %v\n", resp)
}
```

[[Back to top]](#)

