---
id: v2026-human-account-deletion-approval-config
title: HumanAccountDeletionApprovalConfig
pagination_label: HumanAccountDeletionApprovalConfig
sidebar_label: HumanAccountDeletionApprovalConfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'HumanAccountDeletionApprovalConfig', 'V2026HumanAccountDeletionApprovalConfig'] 
slug: /tools/sdk/go/v2026/methods/human-account-deletion-approval-config
tags: ['SDK', 'Software Development Kit', 'HumanAccountDeletionApprovalConfig', 'V2026HumanAccountDeletionApprovalConfig']
---

# HumanAccountDeletionApprovalConfigAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-account-delete-approval-config**](#get-account-delete-approval-config) | **Get** `/sources/{sourceId}/approval-config/account-delete` | Human Account Deletion Approval Config
[**update-account-deletion-approval-config**](#update-account-deletion-approval-config) | **Patch** `/sources/{sourceId}/approval-config/account-delete` | Human Account Deletion Approval Config


## get-account-delete-approval-config
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
Human Account Deletion Approval Config
The endpoint retrieves the approval configuration for deleting human accounts from a specified source. It returns details such as whether approval is required, who the approvers are, and any additional approval settings. This helps administrators understand and manage the approval workflow for human account deletions in their organization. The response is provided as an AccountDeleteConfigDto object.


[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-account-delete-approval-config)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The Source id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDeleteApprovalConfigRequest struct via the builder pattern


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
    sourceId := `ha38f94347e94562b5bb8424a56498d8` // string | The Source id # string | The Source id

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.HumanAccountDeletionApprovalConfigAPI.GetAccountDeleteApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.V2026.HumanAccountDeletionApprovalConfigAPI.GetAccountDeleteApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `HumanAccountDeletionApprovalConfigAPI.GetAccountDeleteApprovalConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDeleteApprovalConfig`: AccountDeleteConfigDto
    fmt.Fprintf(os.Stdout, "Response from `HumanAccountDeletionApprovalConfigAPI.GetAccountDeleteApprovalConfig`: %v\n", resp)
}
```

[[Back to top]](#)

## update-account-deletion-approval-config
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
Human Account Deletion Approval Config
Updates the approval configuration for deleting human accounts for a specific source, identified by source ID. This endpoint allows administrators to modify settings such as whether approval is required, who the approvers are, and other approval-related options. The update is performed using a JSON Patch payload, and the response returns the updated AccountDeleteConfigDto object reflecting the new approval workflow configuration.


[API Spec](https://developer.sailpoint.com/docs/api/v2026/update-account-deletion-approval-config)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | Human account source ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountDeletionApprovalConfigRequest struct via the builder pattern


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
    sourceId := `00eebcf881994e419d72e757fd30dc0e` // string | Human account source ID. # string | Human account source ID.
    jsonpatchoperation := []byte(``) // []JsonPatchOperation | The JSONPatch payload used to update the object.

    var jsonPatchOperation []v2026.JsonPatchOperation
    if err := json.Unmarshal(jsonpatchoperation, &jsonPatchOperation); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.HumanAccountDeletionApprovalConfigAPI.UpdateAccountDeletionApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	  //resp, r, err := apiClient.V2026.HumanAccountDeletionApprovalConfigAPI.UpdateAccountDeletionApprovalConfig(context.Background(), sourceId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `HumanAccountDeletionApprovalConfigAPI.UpdateAccountDeletionApprovalConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountDeletionApprovalConfig`: AccountDeleteConfigDto
    fmt.Fprintf(os.Stdout, "Response from `HumanAccountDeletionApprovalConfigAPI.UpdateAccountDeletionApprovalConfig`: %v\n", resp)
}
```

[[Back to top]](#)

