---
id: nerm-forms
title: Forms
pagination_label: Forms
sidebar_label: Forms
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Forms', 'NERMForms'] 
slug: /tools/sdk/go/nerm/methods/forms
tags: ['SDK', 'Software Development Kit', 'Forms', 'NERMForms']
---

# FormsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-form**](#create-form) | **Post** `/forms` | Create a form
[**delete-form-by-id**](#delete-form-by-id) | **Delete** `/forms/{id}` | Delete form by id
[**delete-form-by-uid**](#delete-form-by-uid) | **Delete** `/forms/{uid}` | Delete form by UID
[**get-form-by-id**](#get-form-by-id) | **Get** `/forms/{id}` | Get form data by Id
[**get-form-by-uid**](#get-form-by-uid) | **Get** `/forms/{uid}` | Get form data by UID
[**get-forms**](#get-forms) | **Get** `/forms` | Get forms
[**update-form-by-id**](#update-form-by-id) | **Patch** `/forms/{id}` | Update form data by id
[**update-form-by-uid**](#update-form-by-uid) | **Patch** `/forms/{uid}` | Update form data by UID


## create-form
Create a form
This endpoint can create a form

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-form)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFormRequest** | [**CreateFormRequest**](../models/create-form-request) |  | 

### Return type

[**GetForms200Response**](../models/get-forms200-response)

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
    createformrequest := []byte(``) // CreateFormRequest | 

    var createFormRequest NERM.CreateFormRequest
    if err := json.Unmarshal(createformrequest, &createFormRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormsAPI.CreateForm(context.Background()).CreateFormRequest(createFormRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.CreateForm(context.Background()).CreateFormRequest(createFormRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.CreateForm``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForm`: GetForms200Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.CreateForm`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-form-by-id
Delete form by id
Delete form by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-form-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
    resp, r, err := apiClient.NERM.FormsAPI.DeleteFormById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.DeleteFormById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.DeleteFormById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFormById`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.DeleteFormById`: %v\n", resp)
}
```

[[Back to top]](#)

## delete-form-by-uid
Delete form by UID
Delete form by UID

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-form-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
    resp, r, err := apiClient.NERM.FormsAPI.DeleteFormByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.DeleteFormByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.DeleteFormByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFormByUid`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.DeleteFormByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-form-by-id
Get form data by Id
Info for a specific form

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-form-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetForms200Response**](../models/get-forms200-response)

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
    resp, r, err := apiClient.NERM.FormsAPI.GetFormById(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.GetFormById(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetFormById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormById`: GetForms200Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetFormById`: %v\n", resp)
}
```

[[Back to top]](#)

## get-form-by-uid
Get form data by UID
Info for a specific form

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-form-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetForms200Response**](../models/get-forms200-response)

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
    resp, r, err := apiClient.NERM.FormsAPI.GetFormByUid(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.GetFormByUid(context.Background(), uid).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetFormByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFormByUid`: GetForms200Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetFormByUid`: %v\n", resp)
}
```

[[Back to top]](#)

## get-forms
Get forms
Get defined forms in the system

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-forms)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFormsRequest struct via the builder pattern


### Return type

[**GetForms200Response**](../models/get-forms200-response)

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
    resp, r, err := apiClient.NERM.FormsAPI.GetForms(context.Background()).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.GetForms(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.GetForms``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetForms`: GetForms200Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.GetForms`: %v\n", resp)
}
```

[[Back to top]](#)

## update-form-by-id
Update form data by id
Update info for a specific form

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-form-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFormByIdRequest** | [**UpdateFormByIdRequest**](../models/update-form-by-id-request) |  | 

### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
    updateformbyidrequest := []byte(``) // UpdateFormByIdRequest | 

    var updateFormByIdRequest NERM.UpdateFormByIdRequest
    if err := json.Unmarshal(updateformbyidrequest, &updateFormByIdRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormsAPI.UpdateFormById(context.Background(), id).UpdateFormByIdRequest(updateFormByIdRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.UpdateFormById(context.Background(), id).UpdateFormByIdRequest(updateFormByIdRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.UpdateFormById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFormById`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.UpdateFormById`: %v\n", resp)
}
```

[[Back to top]](#)

## update-form-by-uid
Update form data by UID
Update info for a specific form

[API Spec](https://developer.sailpoint.com/docs/api/NERM/update-form-by-uid)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormByUidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFormByIdRequest** | [**UpdateFormByIdRequest**](../models/update-form-by-id-request) |  | 


### Return type

[**CreateAttribute201Response**](../models/create-attribute201-response)

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
    updateformbyidrequest := []byte(``) // UpdateFormByIdRequest | 
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)

    var updateFormByIdRequest NERM.UpdateFormByIdRequest
    if err := json.Unmarshal(updateformbyidrequest, &updateFormByIdRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.FormsAPI.UpdateFormByUid(context.Background(), uid).UpdateFormByIdRequest(updateFormByIdRequest).Execute()
	  //resp, r, err := apiClient.NERM.FormsAPI.UpdateFormByUid(context.Background(), uid).UpdateFormByIdRequest(updateFormByIdRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `FormsAPI.UpdateFormByUid``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFormByUid`: CreateAttribute201Response
    fmt.Fprintf(os.Stdout, "Response from `FormsAPI.UpdateFormByUid`: %v\n", resp)
}
```

[[Back to top]](#)

