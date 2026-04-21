---
id: nerm-page-contents
title: PageContents
pagination_label: PageContents
sidebar_label: PageContents
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContents', 'NERMPageContents'] 
slug: /tools/sdk/go/nerm/methods/page-contents
tags: ['SDK', 'Software Development Kit', 'PageContents', 'NERMPageContents']
---

# PageContentsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-page-content**](#create-page-content) | **Post** `/page_contents` | Create a page content entry
[**delete-page-content-by-id**](#delete-page-content-by-id) | **Delete** `/page_contents/{id}` | Delete page content record
[**delete-page-content-by-uid**](#delete-page-content-by-uid) | **Delete** `/page_contents/{uid}` | Delete page content record
[**get-page-content-by-id**](#get-page-content-by-id) | **Get** `/page_contents/{id}` | Find page content record
[**get-page-content-by-uid**](#get-page-content-by-uid) | **Get** `/page_contents/{uid}` | Find a page content record
[**get-page-contents**](#get-page-contents) | **Get** `/page_contents` | Get page contents data
[**update-page-content-by-id**](#update-page-content-by-id) | **Patch** `/page_contents/{id}` | Update page content record
[**update-page-content-by-uid**](#update-page-content-by-uid) | **Patch** `/page_contents/{uid}` | Update page content record


## create-page-content
Create a page content entry
This endpoint can create page content

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-page-content)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePageContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageContentRequest** | [**CreatePageContentRequest**](../models/create-page-content-request) |  | 

### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    createpagecontentrequest := []byte(``) // CreatePageContentRequest | 

    var createPageContentRequest NERM.CreatePageContentRequest
    if err := json.Unmarshal(createpagecontentrequest, &createPageContentRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentsAPI.CreatePageContent(context.Background()).CreatePageContentRequest(createPageContentRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.CreatePageContent(context.Background()).CreatePageContentRequest(createPageContentRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.CreatePageContent``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePageContent`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.CreatePageContent`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-content-by-id
Delete page content record
Delete page content by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-content-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    resp, r, err := apiClient.NERM.PageContentsAPI.DeletePageContentById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.DeletePageContentById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.DeletePageContentById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageContentById`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.DeletePageContentById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-content-by-uid
Delete page content record
Delete page content by by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-content-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageContentByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    resp, r, err := apiClient.NERM.PageContentsAPI.DeletePageContentByUid(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.DeletePageContentByUid(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.DeletePageContentByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageContentByUid`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.DeletePageContentByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-content-by-id
Find page content record
Info for a specific page content record

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-content-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContentById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContentById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.GetPageContentById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentById`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.GetPageContentById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-content-by-uid
Find a page content record
Info for a specific page content record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-content-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContentByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContentByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.GetPageContentByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentByUid`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.GetPageContentByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-contents
Get page contents data
This endpoint can retrieve page content data.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-contents)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentsRequest struct via the builder pattern


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContents(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.GetPageContents(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.GetPageContents``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContents`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.GetPageContents`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-content-by-id
Update page content record
Update info for a specific page content record by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-content-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageContentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPageContentRequest** | [**CreatePageContentRequest**](../models/create-page-content-request) |  | 

### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    createpagecontentrequest := []byte(``) // CreatePageContentRequest | 

    var createPageContentRequest NERM.CreatePageContentRequest
    if err := json.Unmarshal(createpagecontentrequest, &createPageContentRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentsAPI.UpdatePageContentById(context.Background(), id).CreatePageContentRequest(createPageContentRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.UpdatePageContentById(context.Background(), id).CreatePageContentRequest(createPageContentRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.UpdatePageContentById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageContentById`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.UpdatePageContentById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-content-by-uid
Update page content record
Update info for a specific page content record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-content-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageContentByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageContentRequest** | [**CreatePageContentRequest**](../models/create-page-content-request) |  | 


### Return type

[**GetPageContents200Response**](../models/get-page-contents200-response)

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
    createpagecontentrequest := []byte(``) // CreatePageContentRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var createPageContentRequest NERM.CreatePageContentRequest
    if err := json.Unmarshal(createpagecontentrequest, &createPageContentRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentsAPI.UpdatePageContentByUid(context.Background(), uid).CreatePageContentRequest(createPageContentRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentsAPI.UpdatePageContentByUid(context.Background(), uid).CreatePageContentRequest(createPageContentRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentsAPI.UpdatePageContentByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageContentByUid`: GetPageContents200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentsAPI.UpdatePageContentByUid`: %v\n", resp)
}
```

[[Back to top]](#)

