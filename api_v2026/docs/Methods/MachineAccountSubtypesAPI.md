---
id: v2026-machine-account-subtypes
title: MachineAccountSubtypes
pagination_label: MachineAccountSubtypes
sidebar_label: MachineAccountSubtypes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountSubtypes', 'V2026MachineAccountSubtypes'] 
slug: /tools/sdk/go/v2026/methods/machine-account-subtypes
tags: ['SDK', 'Software Development Kit', 'MachineAccountSubtypes', 'V2026MachineAccountSubtypes']
---

# MachineAccountSubtypesAPI
  Use this API to get, update, and delete machine account subtype for sources.
 
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-machine-account-subtype**](#delete-machine-account-subtype) | **Delete** `/source-subtypes/{subtypeId}` | Delete subtype by ID
[**get-source-subtype-by-id**](#get-source-subtype-by-id) | **Get** `/source-subtypes/{subtypeId}` | Get subtype by ID
[**patch-machine-account-subtype**](#patch-machine-account-subtype) | **Patch** `/source-subtypes/{subtypeId}` | Patch subtype by ID


## delete-machine-account-subtype
Delete subtype by ID
Delete a machine account subtype by subtype ID.

Note: If subtype has approval settings or entitlement for machine account creation enablement then it'll be also deleted.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/delete-machine-account-subtype)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineAccountSubtypeRequest struct via the builder pattern


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
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V2026.MachineAccountSubtypesAPI.DeleteMachineAccountSubtype(context.Background(), subtypeId).Execute()
	  //r, err := apiClient.V2026.MachineAccountSubtypesAPI.DeleteMachineAccountSubtype(context.Background(), subtypeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.DeleteMachineAccountSubtype``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## get-source-subtype-by-id
Get subtype by ID
Get a machine account subtype by subtype ID.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-source-subtype-by-id)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSubtypeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceSubtypeWithSource**](../models/source-subtype-with-source)

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
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineAccountSubtypesAPI.GetSourceSubtypeById(context.Background(), subtypeId).Execute()
	  //resp, r, err := apiClient.V2026.MachineAccountSubtypesAPI.GetSourceSubtypeById(context.Background(), subtypeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.GetSourceSubtypeById``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSubtypeById`: SourceSubtypeWithSource
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountSubtypesAPI.GetSourceSubtypeById`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-machine-account-subtype
Patch subtype by ID
Update fields of a machine account subtype by subtype ID.
Patchable fields only include: `displayName`, `description`.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/patch-machine-account-subtype)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMachineAccountSubtypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**SourceSubtypeWithSource**](../models/source-subtype-with-source)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
  "encoding/json"
    v2026 "github.com/sailpoint-oss/golang-sdk/v2/api_v2026"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.
    requestbody := []byte(`[{op=replace, path=/displayName, value=Test New DisplayName}]`) // []map[string]interface{} | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    var requestBody []v2026.RequestBody
    if err := json.Unmarshal(requestbody, &requestBody); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineAccountSubtypesAPI.PatchMachineAccountSubtype(context.Background(), subtypeId).RequestBody(requestBody).Execute()
	  //resp, r, err := apiClient.V2026.MachineAccountSubtypesAPI.PatchMachineAccountSubtype(context.Background(), subtypeId).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.PatchMachineAccountSubtype``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMachineAccountSubtype`: SourceSubtypeWithSource
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountSubtypesAPI.PatchMachineAccountSubtype`: %v\n", resp)
}
```

[[Back to top]](#)

