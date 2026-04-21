---
id: nerm-audits
title: Audits
pagination_label: Audits
sidebar_label: Audits
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Audits', 'NERMAudits'] 
slug: /tools/sdk/go/nerm/methods/audits
tags: ['SDK', 'Software Development Kit', 'Audits', 'NERMAudits']
---

# AuditsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**search**](#search) | **Post** `/audit_events/query` | Query for Audit events


## search
Query for Audit events
This endpoint provides a search engine for Audit Events by optionally combining subject_type, type, and subject_id to narrow down the audit events. A Subject Type of Profile links up to the AuditableProfile types. An Subject Type of WorkflowSession links up to the AuditableWorkflow types. An Subject Type of Get/Post/Patch/Delete links up to the AuditableApi types. The remaining Subject Types link up to the ActiveRecord types (configuration changes).

- Any workflow audit event created as of 10/11/2024 will be able to be queried by workflow name, workflow uid, or workflow profile type.
- Any profile audit event created as of 10/11/2024 will be able to be queried by profile type.
- The entity_type parameter has been updated to subject_type, which now matches what is in the response object.
- With the additional query filters added, there is a max of 5 filter parameters at one time (aside from pagination parameters)

To accommodate these changes, an API contract change was required.  Please read the updated API documentation for the new request syntax.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/search)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequest** | [**SearchRequest**](../models/search-request) |  | 

### Return type

[**Search200Response**](../models/search200-response)

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
    searchrequest := []byte(``) // SearchRequest | 

    var searchRequest NERM.SearchRequest
    if err := json.Unmarshal(searchrequest, &searchRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.AuditsAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
	  //resp, r, err := apiClient.NERM.AuditsAPI.Search(context.Background()).SearchRequest(searchRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `AuditsAPI.Search``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: Search200Response
    fmt.Fprintf(os.Stdout, "Response from `AuditsAPI.Search`: %v\n", resp)
}
```

[[Back to top]](#)

