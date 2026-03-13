---
id: v2026-notifications
title: Notifications
pagination_label: Notifications
sidebar_label: Notifications
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Notifications', 'V2026Notifications'] 
slug: /tools/sdk/go/v2026/methods/notifications
tags: ['SDK', 'Software Development Kit', 'Notifications', 'V2026Notifications']
---

# NotificationsAPI
   
All URIs are relative to *https://sailpoint.api.identitynow.com/v2026*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-notification-template-variables**](#get-notification-template-variables) | **Get** `/notification-template-variables/{key}/{medium}` | Get notification template variables


## get-notification-template-variables
Get notification template variables
Returns global variables and template-specific variables for a given notification template key and medium.
Use these variable names in template content; they are replaced at send time with the corresponding values.
Variable lists can be sorted by key, type, or description via the sorters query parameter (default ascending by key).


[API Spec](https://developer.sailpoint.com/docs/api/v2026/get-notification-template-variables)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The notification template key. Valid keys (and key/medium pairs) are available from the list notification templates operation.  | 
**medium** | **string** | The notification template medium (e.g. EMAIL, SLACK, TEAMS). Valid key/medium pairs are available from the list notification templates operation.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationTemplateVariablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **key, type, description** | 

### Return type

[**TemplateVariablesDto**](../models/template-variables-dto)

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
    key := `approval_request_notification` // string | The notification template key. Valid keys (and key/medium pairs) are available from the list notification templates operation.  # string | The notification template key. Valid keys (and key/medium pairs) are available from the list notification templates operation. 
    medium := `EMAIL` // string | The notification template medium (e.g. EMAIL, SLACK, TEAMS). Valid key/medium pairs are available from the list notification templates operation.  # string | The notification template medium (e.g. EMAIL, SLACK, TEAMS). Valid key/medium pairs are available from the list notification templates operation. 
    sorters := `-description` // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **key, type, description** (optional) # string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **key, type, description** (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V2026.NotificationsAPI.GetNotificationTemplateVariables(context.Background(), key, medium).Execute()
	  //resp, r, err := apiClient.V2026.NotificationsAPI.GetNotificationTemplateVariables(context.Background(), key, medium).Sorters(sorters).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationTemplateVariables``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationTemplateVariables`: TemplateVariablesDto
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationTemplateVariables`: %v\n", resp)
}
```

[[Back to top]](#)

