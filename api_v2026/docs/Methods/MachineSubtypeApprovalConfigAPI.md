---
id: v2026-machine-subtype-approval-config
title: MachineSubtypeApprovalConfig
pagination_label: MachineSubtypeApprovalConfig
sidebar_label: MachineSubtypeApprovalConfig
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'MachineSubtypeApprovalConfig', 'V2026MachineSubtypeApprovalConfig'] 
slug: /tools/sdk/go/v2026/methods/machine-subtype-approval-config
tags: ['SDK', 'Software Development Kit', 'MachineSubtypeApprovalConfig', 'V2026MachineSubtypeApprovalConfig']
---

# MachineSubtypeApprovalConfigAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-machine-account-deletion-sub-type-approval-config**](#get-machine-account-deletion-sub-type-approval-config) | **Get** `/source-subtypes/{subtypeId}/machine-config` | Machine Subtype Approval Config
[**update-machine-account-deletion-by-sub-type-approval-config**](#update-machine-account-deletion-by-sub-type-approval-config) | **Patch** `/source-subtypes/{subtypeId}/machine-config` | Machine Subtype Approval Config


## get-machine-account-deletion-sub-type-approval-config
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
Machine Subtype Approval Config
This endpoint retrieves the approval configuration for machine account deletion at the machine subtype level. By providing a specific subtypeId in the path, clients can fetch the approval rules and settings (such as required approvers and comments policy) that govern account deletion for that particular machine subtype. The response includes a MachineAccountSubtypeConfigDto object detailing these configurations, enabling clients to understand or display the approval workflow required for deleting machine accounts of the given subtype. Use this endpoint to get machine subtype level approval config for account deletion.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-machine-account-deletion-sub-type-approval-config)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | machine subtype id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineAccountDeletionSubTypeApprovalConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]


### Return type

[**MachineAccountSubTypeConfigDto**](../models/machine-account-sub-type-config-dto)

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
    subtypeId := `ef38f94347e94562b5bb8424a56498d8` // string | machine subtype id. # string | machine subtype id.

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineSubtypeApprovalConfigAPI.GetMachineAccountDeletionSubTypeApprovalConfig(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.V2026.MachineSubtypeApprovalConfigAPI.GetMachineAccountDeletionSubTypeApprovalConfig(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineSubtypeApprovalConfigAPI.GetMachineAccountDeletionSubTypeApprovalConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineAccountDeletionSubTypeApprovalConfig`: MachineAccountSubTypeConfigDto
    fmt.Fprintf(os.Stdout, "Response from `MachineSubtypeApprovalConfigAPI.GetMachineAccountDeletionSubTypeApprovalConfig`: %v\n", resp)
}
```

[[Back to top]](#)

## update-machine-account-deletion-by-sub-type-approval-config
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
Machine Subtype Approval Config
Updates the approval configuration for machine account deletion at the specified machine subtype level. This endpoint allows clients to modify approval rules and settings (such as required approvers and comments policy) for account deletion workflows associated with a given subtypeId. Use this to customize or enforce approval requirements for deleting machine accounts of a particular subtype.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/update-machine-account-deletion-by-sub-type-approval-config)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtypeId** | **string** | machine account subtype ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMachineAccountDeletionBySubTypeApprovalConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

 **jsonPatchOperation** | [**[]JsonPatchOperation**](../models/json-patch-operation) | The JSONPatch payload used to update the object. | 

### Return type

[**MachineAccountSubTypeConfigDto**](../models/machine-account-sub-type-config-dto)

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
    subtypeId := `00eebcf881994e419d72e757fd30dc0e` // string | machine account subtype ID. # string | machine account subtype ID.
    jsonpatchoperation := []byte(``) // []JsonPatchOperation | The JSONPatch payload used to update the object.

    var jsonPatchOperation []v2026.JsonPatchOperation
    if err := json.Unmarshal(jsonpatchoperation, &jsonPatchOperation); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.MachineSubtypeApprovalConfigAPI.UpdateMachineAccountDeletionBySubTypeApprovalConfig(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
	  //resp, r, err := apiClient.V2026.MachineSubtypeApprovalConfigAPI.UpdateMachineAccountDeletionBySubTypeApprovalConfig(context.Background(), subtypeId).XSailPointExperimental(xSailPointExperimental).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `MachineSubtypeApprovalConfigAPI.UpdateMachineAccountDeletionBySubTypeApprovalConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMachineAccountDeletionBySubTypeApprovalConfig`: MachineAccountSubTypeConfigDto
    fmt.Fprintf(os.Stdout, "Response from `MachineSubtypeApprovalConfigAPI.UpdateMachineAccountDeletionBySubTypeApprovalConfig`: %v\n", resp)
}
```

[[Back to top]](#)

