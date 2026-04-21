---
id: nerm-page-content-translations
title: PageContentTranslations
pagination_label: PageContentTranslations
sidebar_label: PageContentTranslations
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'PageContentTranslations', 'NERMPageContentTranslations'] 
slug: /tools/sdk/go/nerm/methods/page-content-translations
tags: ['SDK', 'Software Development Kit', 'PageContentTranslations', 'NERMPageContentTranslations']
---

# PageContentTranslationsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-page-content-translation**](#create-page-content-translation) | **Post** `/page_content_translations` | Create page content translation
[**delete-page-content-translation-by-id**](#delete-page-content-translation-by-id) | **Delete** `/page_content_translations/{id}` | Delete page content translation
[**delete-page-content-translation-by-uid**](#delete-page-content-translation-by-uid) | **Delete** `/page_content_translations/{uid}` | Delete page content translation
[**get-page-content-translation**](#get-page-content-translation) | **Get** `/page_content_translations` | Get page contents translation
[**get-page-content-translation-by-id**](#get-page-content-translation-by-id) | **Get** `/page_content_translations/{id}` | Find page content translation
[**get-page-content-translation-by-uid**](#get-page-content-translation-by-uid) | **Get** `/page_content_translations/{uid}` | Find page content translation
[**update-page-content-translation-by-id**](#update-page-content-translation-by-id) | **Patch** `/page_content_translations/{id}` | Update page content translation
[**update-page-content-translation-by-uid**](#update-page-content-translation-by-uid) | **Patch** `/page_content_translations/{uid}` | Update page content translation


## create-page-content-translation
Create page content translation
Create a page content translation record.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-page-content-translation)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePageContentTranslationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPageContentTranslationRequest** | [**CreatePageContentTranslationRequest**](../models/create-page-content-translation-request) |  | 

### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    createpagecontenttranslationrequest := []byte(``) // CreatePageContentTranslationRequest | 

    var createPageContentTranslationRequest NERM.CreatePageContentTranslationRequest
    if err := json.Unmarshal(createpagecontenttranslationrequest, &createPageContentTranslationRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.CreatePageContentTranslation(context.Background()).CreatePageContentTranslationRequest(createPageContentTranslationRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.CreatePageContentTranslation(context.Background()).CreatePageContentTranslationRequest(createPageContentTranslationRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.CreatePageContentTranslation``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePageContentTranslation`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.CreatePageContentTranslation`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-content-translation-by-id
Delete page content translation
Delete page content translation by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-content-translation-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageContentTranslationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.DeletePageContentTranslationById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.DeletePageContentTranslationById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.DeletePageContentTranslationById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageContentTranslationById`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.DeletePageContentTranslationById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-page-content-translation-by-uid
Delete page content translation
Delete page content translation by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-page-content-translation-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePageContentTranslationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.DeletePageContentTranslationByUid(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.DeletePageContentTranslationByUid(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.DeletePageContentTranslationByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePageContentTranslationByUid`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.DeletePageContentTranslationByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-content-translation
Get page contents translation
This endpoint can retrieve page content translation data.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-content-translation)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentTranslationRequest struct via the builder pattern


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslation(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslation(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.GetPageContentTranslation``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentTranslation`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.GetPageContentTranslation`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-content-translation-by-id
Find page content translation
Info for a specific page content translation record by Id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-content-translation-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentTranslationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslationById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslationById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.GetPageContentTranslationById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentTranslationById`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.GetPageContentTranslationById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-page-content-translation-by-uid
Find page content translation
Info for a specific page content translation record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-page-content-translation-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageContentTranslationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslationByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.GetPageContentTranslationByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.GetPageContentTranslationByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPageContentTranslationByUid`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.GetPageContentTranslationByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-content-translation-by-id
Update page content translation
Update info for a specific page content translation record by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-content-translation-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageContentTranslationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePageContentTranslationByIdRequest** | [**UpdatePageContentTranslationByIdRequest**](../models/update-page-content-translation-by-id-request) |  | 

### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    updatepagecontenttranslationbyidrequest := []byte(``) // UpdatePageContentTranslationByIdRequest | 

    var updatePageContentTranslationByIdRequest NERM.UpdatePageContentTranslationByIdRequest
    if err := json.Unmarshal(updatepagecontenttranslationbyidrequest, &updatePageContentTranslationByIdRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.UpdatePageContentTranslationById(context.Background(), id).UpdatePageContentTranslationByIdRequest(updatePageContentTranslationByIdRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.UpdatePageContentTranslationById(context.Background(), id).UpdatePageContentTranslationByIdRequest(updatePageContentTranslationByIdRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.UpdatePageContentTranslationById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageContentTranslationById`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.UpdatePageContentTranslationById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-page-content-translation-by-uid
Update page content translation
Update info for a specific page content translation record by UID (user-specified identifier)

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-page-content-translation-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePageContentTranslationByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePageContentTranslationByIdRequest** | [**UpdatePageContentTranslationByIdRequest**](../models/update-page-content-translation-by-id-request) |  | 


### Return type

[**GetPageContentTranslation200Response**](../models/get-page-content-translation200-response)

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
    updatepagecontenttranslationbyidrequest := []byte(``) // UpdatePageContentTranslationByIdRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var updatePageContentTranslationByIdRequest NERM.UpdatePageContentTranslationByIdRequest
    if err := json.Unmarshal(updatepagecontenttranslationbyidrequest, &updatePageContentTranslationByIdRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.PageContentTranslationsAPI.UpdatePageContentTranslationByUid(context.Background(), uid).UpdatePageContentTranslationByIdRequest(updatePageContentTranslationByIdRequest).Execute()
	  //resp, r, err := apiClient.NERM.PageContentTranslationsAPI.UpdatePageContentTranslationByUid(context.Background(), uid).UpdatePageContentTranslationByIdRequest(updatePageContentTranslationByIdRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `PageContentTranslationsAPI.UpdatePageContentTranslationByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePageContentTranslationByUid`: GetPageContentTranslation200Response
    fmt.Fprintf(os.Stdout, "Response from `PageContentTranslationsAPI.UpdatePageContentTranslationByUid`: %v\n", resp)
}
```

[[Back to top]](#)

