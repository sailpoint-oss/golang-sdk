---
id: v2026-machine-account-deletion-approval-config
title: MachineAccountDeletionApprovalConfig
pagination_label: MachineAccountDeletionApprovalConfig
sidebar_label: MachineAccountDeletionApprovalConfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineAccountDeletionApprovalConfig', 'V2026MachineAccountDeletionApprovalConfig'] 
slug: /tools/sdk/go/v2026/methods/machine-account-deletion-approval-config
tags: ['SDK', 'Software Development Kit', 'MachineAccountDeletionApprovalConfig', 'V2026MachineAccountDeletionApprovalConfig']
---

# MachineAccountDeletionApprovalConfigAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-machine-account-deletion-approval-config-by-source**](#get-machine-account-deletion-approval-config-by-source) | **Get** `/sources/{sourceId}/approval-config/machine-account-delete` | Machine Account Deletion Approval Config
[**update-machine-account-deletion-approval-config**](#update-machine-account-deletion-approval-config) | **Patch** `/sources/{sourceId}/approval-config/machine-account-delete` | Machine Account Deletion Approval Config


## get-machine-account-deletion-approval-config-by-source
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Machine Account Deletion Approval Config
Retrieves the machine account deletion approval configuration for a specific source. This endpoint returns details about the approval requirements, approvers, and comment settings that govern the deletion of machine accounts associated with the given source ID.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-machine-account-deletion-approval-config-by-source)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountDeletionApprovalConfigBySourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]


### Return type

[**AccountDeleteConfigDto**](../models/account-delete-config-dto)

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
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    sourceId := `gt38f94347e94562b5bb8424a56498d8` // string | source id. # string | source id.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineAccountDeletionApprovalConfigAPI.GetMachineAccountDeletionApprovalConfigBySource(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.V2026.MachineAccountDeletionApprovalConfigAPI.GetMachineAccountDeletionApprovalConfigBySource(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountDeletionApprovalConfigAPI.GetMachineAccountDeletionApprovalConfigBySource``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountDeletionApprovalConfigBySource`: AccountDeleteConfigDto
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountDeletionApprovalConfigAPI.GetMachineAccountDeletionApprovalConfigBySource`: %v\n", resp)
}
```

[[Back to top]](#)

## update-machine-account-deletion-approval-config
:::warning experimental 
This API is currently in an experimental state. The API is subject to change based on feedback and further testing. You must include the X-SailPoint-Experimental header and set it to `true` to use this endpoint.
:::
:::tip setting x-sailpoint-experimental header
 on the configuration object you can set the `x-sailpoint-experimental` header to `true' to enable all experimantl endpoints within the SDK.
 Example:
 ```go
   configuration = Configuration()
   configuration.Experimental = true
 ```
:::
Machine Account Deletion Approval Config
Use this endpoint to update the machine account deletion approval configuration for a specific source.
The update is performed using a JSON Patch payload, which allows partial modifications to the approval config.
This operation is typically used to change approval requirements, approvers, or comments settings for machine account deletion.
The endpoint expects the source ID as a path parameter and a valid JSON Patch array in the request body.


[API Spec](https://developer.sailpoint.com/docs/api/v2026/update-machine-account-deletion-approval-config)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | machine account source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMachineAccountDeletionApprovalConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

 **jsonPatchOperation** | [**[]JsonPatchOperation**](../models/json-patch-operation) | The JSONPatch payload used to update the object. | 

### Return type

[**AccountDeleteConfigDto**](../models/account-delete-config-dto)

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
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    sourceId := `00eebcf881994e419d72e757fd30dc0e` // string | machine account source ID. # string | machine account source ID.
    jsonpatchoperation := []byte(``) // []JsonPatchOperation | The JSONPatch payload used to update the object.

    var jsonPatchOperation []v2026.JsonPatchOperation
    if err := json.Unmarshal(jsonpatchoperation, &jsonPatchOperation); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineAccountDeletionApprovalConfigAPI.UpdateMachineAccountDeletionApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	  //resp, r, err := apiClient.V2026.MachineAccountDeletionApprovalConfigAPI.UpdateMachineAccountDeletionApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineAccountDeletionApprovalConfigAPI.UpdateMachineAccountDeletionApprovalConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMachineAccountDeletionApprovalConfig`: AccountDeleteConfigDto
    fmt.Fprintf(os.Stdout, "Response from `MachineAccountDeletionApprovalConfigAPI.UpdateMachineAccountDeletionApprovalConfig`: %v\n", resp)
}
```

[[Back to top]](#)

