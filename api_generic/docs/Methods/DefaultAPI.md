---
id: default
title: Default
pagination_label: Default
sidebar_label: Default
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Default', 'Default'] 
slug: /tools/sdk/go//methods/default
tags: ['SDK', 'Software Development Kit', 'Default', 'Default']
---

# DefaultAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**generic-delete**](#generic-delete) | **Delete** `/{path}` | Generic DELETE request
[**generic-get**](#generic-get) | **Get** `/{path}` | Generic GET request
[**generic-patch**](#generic-patch) | **Patch** `/{path}` | Generic PATCH request
[**generic-post**](#generic-post) | **Post** `/{path}` | Generic POST request
[**generic-put**](#generic-put) | **Put** `/{path}` | Generic PUT request


## generic-delete
Generic DELETE request


[API Spec](https://developer.sailpoint.com/docs/api//generic-delete)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericResponse**](../models/generic-response)

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
    path := `path_example` // string |  # string | 

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient..DefaultAPI.GenericDelete(context.Background(), path).Execute()
	  //resp, r, err := apiClient..DefaultAPI.GenericDelete(context.Background(), path).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericDelete``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDelete`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericDelete`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-get
Generic GET request


[API Spec](https://developer.sailpoint.com/docs/api//generic-get)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericResponse**](../models/generic-response)

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
    path := `path_example` // string |  # string | 

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient..DefaultAPI.GenericGet(context.Background(), path).Execute()
	  //resp, r, err := apiClient..DefaultAPI.GenericGet(context.Background(), path).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericGet``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericGet`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericGet`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-patch
Generic PATCH request


[API Spec](https://developer.sailpoint.com/docs/api//generic-patch)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**GenericResponse**](../models/generic-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    path := `path_example` // string |  # string | 
    requestBody := Object // map[string]interface{} |  (optional) # map[string]interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient..DefaultAPI.GenericPatch(context.Background(), path).Execute()
	  //resp, r, err := apiClient..DefaultAPI.GenericPatch(context.Background(), path).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPatch``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPatch`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPatch`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-post
Generic POST request


[API Spec](https://developer.sailpoint.com/docs/api//generic-post)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**GenericResponse**](../models/generic-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    path := `path_example` // string |  # string | 
    requestBody := Object // map[string]interface{} |  (optional) # map[string]interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient..DefaultAPI.GenericPost(context.Background(), path).Execute()
	  //resp, r, err := apiClient..DefaultAPI.GenericPost(context.Background(), path).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPost``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPost`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPost`: %v\n", resp)
}
```

[[Back to top]](#)

## generic-put
Generic PUT request


[API Spec](https://developer.sailpoint.com/docs/api//generic-put)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenericPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**GenericResponse**](../models/generic-response)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    path := `path_example` // string |  # string | 
    requestBody := Object // map[string]interface{} |  (optional) # map[string]interface{} |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient..DefaultAPI.GenericPut(context.Background(), path).Execute()
	  //resp, r, err := apiClient..DefaultAPI.GenericPut(context.Background(), path).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenericPut``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericPut`: GenericResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenericPut`: %v\n", resp)
}
```

[[Back to top]](#)

