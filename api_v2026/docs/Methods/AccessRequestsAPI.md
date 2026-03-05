---
id: v2026-access-requests
title: AccessRequests
pagination_label: AccessRequests
sidebar_label: AccessRequests
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'AccessRequests', 'V2026AccessRequests'] 
slug: /tools/sdk/go/v2026/methods/access-requests
tags: ['SDK', 'Software Development Kit', 'AccessRequests', 'V2026AccessRequests']
---

# AccessRequestsAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-access-request-config**](#get-access-request-config) | **Get** `/access-request-config` | Get access request configuration
[**set-access-request-config**](#set-access-request-config) | **Put** `/access-request-config` | Update access request configuration


## get-access-request-config
Get access request configuration
This endpoint returns the current access-request configuration.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-access-request-config)

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestConfigRequest struct via the builder pattern


### Return type

[**AccessRequestConfig**](../models/access-request-config)

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
    resp, r, err := apiClient.V2026.AccessRequestsAPI.GetAccessRequestConfig(context.Background()).Execute()
	  //resp, r, err := apiClient.V2026.AccessRequestsAPI.GetAccessRequestConfig(context.Background()).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.GetAccessRequestConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestConfig`: AccessRequestConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.GetAccessRequestConfig`: %v\n", resp)
}
```

[[Back to top]](#)

## set-access-request-config
Update access request configuration
This endpoint replaces the current access-request configuration.

[API Spec](https://developer.sailpoint.com/docs/api/v2026/set-access-request-config)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetAccessRequestConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequestConfig** | [**AccessRequestConfig**](../models/access-request-config) |  | 

### Return type

[**AccessRequestConfig**](../models/access-request-config)

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
    v2026 "github.com/sailpoint-oss/golang-sdk/v2/api_v2026"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    accessrequestconfig := []byte(`{
          "govGroupVisibilityEnabled" : true,
          "requestOnBehalfOfConfig" : {
            "allowRequestOnBehalfOfEmployeeByManager" : true,
            "allowRequestOnBehalfOfAnyoneByAnyone" : true
          },
          "autoApprovalEnabled" : true,
          "entitlementRequestConfig" : {
            "accessRequestConfig" : {
              "denialCommentRequired" : false,
              "approvalSchemes" : [ {
                "approverId" : "e3eab852-8315-467f-9de7-70eda97f63c8",
                "approverType" : "GOVERNANCE_GROUP"
              }, {
                "approverId" : "e3eab852-8315-467f-9de7-70eda97f63c8",
                "approverType" : "GOVERNANCE_GROUP"
              } ],
              "reauthorizationRequired" : false,
              "requestCommentRequired" : true,
              "requireEndDate" : true,
              "maxPermittedAccessDuration" : {
                "value" : 5,
                "timeUnit" : "DAYS"
              }
            },
            "revocationRequestConfig" : {
              "approvalSchemes" : [ {
                "approverId" : "e3eab852-8315-467f-9de7-70eda97f63c8",
                "approverType" : "GOVERNANCE_GROUP"
              }, {
                "approverId" : "e3eab852-8315-467f-9de7-70eda97f63c8",
                "approverType" : "GOVERNANCE_GROUP"
              } ]
            }
          },
          "reauthorizationEnabled" : true,
          "approvalsMustBeExternal" : true
        }`) // AccessRequestConfig | 

    var accessRequestConfig v2026.AccessRequestConfig
    if err := json.Unmarshal(accessrequestconfig, &accessRequestConfig); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.AccessRequestsAPI.SetAccessRequestConfig(context.Background()).AccessRequestConfig(accessRequestConfig).Execute()
	  //resp, r, err := apiClient.V2026.AccessRequestsAPI.SetAccessRequestConfig(context.Background()).AccessRequestConfig(accessRequestConfig).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsAPI.SetAccessRequestConfig``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAccessRequestConfig`: AccessRequestConfig
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsAPI.SetAccessRequestConfig`: %v\n", resp)
}
```

[[Back to top]](#)

