---
id: v1-machine-account-subtypes
title: MachineAccountSubtypes
pagination_label: MachineAccountSubtypes
sidebar_label: MachineAccountSubtypes
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountSubtypes', 'V1MachineAccountSubtypes'] 
slug: /tools/sdk/go/machineaccountsubtypesv1/methods/machine-account-subtypes
tags: ['SDK', 'Software Development Kit', 'MachineAccountSubtypes', 'V1MachineAccountSubtypes']
---

# MachineAccountSubtypesAPI
  Use this API to get, update, and delete machine account subtype for sources.
 
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-machine-account-subtype-v1**](#delete-machine-account-subtype-v1) | **Delete** `/source-subtypes/v1/{subtypeId}` | Delete subtype by ID
[**get-source-subtype-by-id-v1**](#get-source-subtype-by-id-v1) | **Get** `/source-subtypes/v1/{subtypeId}` | Get subtype by ID
[**patch-machine-account-subtype-v1**](#patch-machine-account-subtype-v1) | **Patch** `/source-subtypes/v1/{subtypeId}` | Patch subtype by ID


## delete-machine-account-subtype-v1
Delete subtype by ID
Delete a machine account subtype by subtype ID.

Note: If subtype has approval settings or entitlement for machine account creation enablement then it'll be also deleted.

[API Spec](https://developer.sailpoint.com/docs/api/v1/delete-machine-account-subtype-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineAccountSubtypeV1Request struct via the builder pattern


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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2/api_machine_account_subtypes_v1"
)

func main() {
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.DeleteMachineAccountSubtypeV1(context.Background(), subtypeId).Execute()
	  //r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.DeleteMachineAccountSubtypeV1(context.Background(), subtypeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.DeleteMachineAccountSubtypeV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## get-source-subtype-by-id-v1
Get subtype by ID
Get a machine account subtype by subtype ID.

[API Spec](https://developer.sailpoint.com/docs/api/v1/get-source-subtype-by-id-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSubtypeByIdV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sourcesubtypewithsource**](../models/sourcesubtypewithsource)

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
  
    
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2/api_machine_account_subtypes_v1"
)

func main() {
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.GetSourceSubtypeByIdV1(context.Background(), subtypeId).Execute()
	  //resp, r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.GetSourceSubtypeByIdV1(context.Background(), subtypeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.GetSourceSubtypeByIdV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSubtypeByIdV1`: Sourcesubtypewithsource
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountSubtypesAPI.GetSourceSubtypeByIdV1`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-machine-account-subtype-v1
Patch subtype by ID
Update fields of a machine account subtype by subtype ID.
Patchable fields only include: `displayName`, `description`.

[API Spec](https://developer.sailpoint.com/docs/api/v1/patch-machine-account-subtype-v1)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | The ID of the subtype. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMachineAccountSubtypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]map[string]interface{}** | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. | 

### Return type

[**Sourcesubtypewithsource**](../models/sourcesubtypewithsource)

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
    v1 "github.com/sailpoint-oss/golang-sdk/v2/api_machine_account_subtypes_v1"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2/api_machine_account_subtypes_v1"
)

func main() {
    subtypeId := `6d28b7c1-620c-49c6-b6d5-cbf81eb4b5fa` // string | The ID of the subtype. # string | The ID of the subtype.
    requestbody := []byte(`[{"op":"replace","path":"/displayName","value":"Test New DisplayName"}]`) // []map[string]interface{} | A JSON of updated values [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

    var requestBody []v1.RequestBody
    if err := json.Unmarshal(requestbody, &requestBody); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.PatchMachineAccountSubtypeV1(context.Background(), subtypeId).RequestBody(requestBody).Execute()
	  //resp, r, err := apiClient.MachineAccountSubtypesV1.MachineAccountSubtypesAPI.PatchMachineAccountSubtypeV1(context.Background(), subtypeId).RequestBody(requestBody).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountSubtypesAPI.PatchMachineAccountSubtypeV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchMachineAccountSubtypeV1`: Sourcesubtypewithsource
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountSubtypesAPI.PatchMachineAccountSubtypeV1`: %v\n", resp)
}
```

[[Back to top]](#)

