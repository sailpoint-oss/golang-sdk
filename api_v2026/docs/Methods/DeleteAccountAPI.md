---
id: v2026-delete-account
title: DeleteAccount
pagination_label: DeleteAccount
sidebar_label: DeleteAccount
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'DeleteAccount', 'V2026DeleteAccount'] 
slug: /tools/sdk/go/v2026/methods/delete-account
tags: ['SDK', 'Software Development Kit', 'DeleteAccount', 'V2026DeleteAccount']
---

# DeleteAccountAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-account-request**](#delete-account-request) | **Post** `/account-requests/account/{accountId}/delete` | Delete account


## delete-account-request
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
Delete account
Initiates an account deletion request for the specified account.
This method validates the input data, processes the deletion request,
and generates an asynchronous result containing a tracking ID. 
>**NOTE: You can only delete accounts from sources of the "Connected" type. which supports account deletion**

[API Spec](https://developer.sailpoint.com/docs/api/v2026/delete-account-request)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | Account ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xSailPointExperimental** | **string** | Use this header to enable this experimental API. | [default to &quot;true&quot;]

 **accountDeleteRequestInput** | [**AccountDeleteRequestInput**](../models/account-delete-request-input) |  | 

### Return type

[**AccountDeleteAsyncResult**](../models/account-delete-async-result)

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
    xSailPointExperimental := `true` // string | Use this header to enable this experimental API. (default to "true") # string | Use this header to enable this experimental API. (default to "true")
    accountId := `ef38f94347e94562b5bb8424a56498d8` // string | Account ID. # string | Account ID.
    accountdeleterequestinput := []byte(`{
          "comments" : "Requesting account deletion request"
        }`) // AccountDeleteRequestInput |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.DeleteAccountAPI.DeleteAccountRequest(context.Background(), accountId).XSailPointExperimental(xSailPointExperimental).Execute()
	  //resp, r, err := apiClient.V2026.DeleteAccountAPI.DeleteAccountRequest(context.Background(), accountId).XSailPointExperimental(xSailPointExperimental).AccountDeleteRequestInput(accountDeleteRequestInput).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `DeleteAccountAPI.DeleteAccountRequest``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccountRequest`: AccountDeleteAsyncResult
    fmt.Fprintf(os.Stdout, "Response from `DeleteAccountAPI.DeleteAccountRequest`: %v\n", resp)
}
```

[[Back to top]](#)

