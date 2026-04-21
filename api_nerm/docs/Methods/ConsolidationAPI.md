---
id: nerm-consolidation
title: Consolidation
pagination_label: Consolidation
sidebar_label: Consolidation
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Consolidation', 'NERMConsolidation'] 
slug: /tools/sdk/go/nerm/methods/consolidation
tags: ['SDK', 'Software Development Kit', 'Consolidation', 'NERMConsolidation']
---

# ConsolidationAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete-master-record**](#delete-master-record) | **Delete** `/idproxy/identities/{id}` | Delete a master record
[**patch-data-record**](#patch-data-record) | **Patch** `/idproxy/data_records/{id}/reassign` | Reassign data record


## delete-master-record
Delete a master record
Consolidation is a deprecated feature, this endpoint provides a mechanism to delete a master record to assist customers.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/delete-master-record)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMasterRecordRequest struct via the builder pattern


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
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.NERM.ConsolidationAPI.DeleteMasterRecord(context.Background(), id).Execute()
	  //r, err := apiClient.NERM.ConsolidationAPI.DeleteMasterRecord(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ConsolidationAPI.DeleteMasterRecord``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

## patch-data-record
Reassign data record
Consolidation is a deprecated feature, this endpoint provides a mechanism to reassign a data record to a new master record to assist customers.

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-data-record)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDataRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataRecords** | [**DataRecords**](../models/data-records) |  | 

### Return type

 (empty response body)

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
    id := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | ID of the object to retrieve, update, or delete # string | ID of the object to retrieve, update, or delete
    datarecords := []byte(`{
          "master_record_id" : "456738c9ba999a0076cf8a9b"
        }`) // DataRecords | 

    var dataRecords NERM.DataRecords
    if err := json.Unmarshal(datarecords, &dataRecords); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.NERM.ConsolidationAPI.PatchDataRecord(context.Background(), id).DataRecords(dataRecords).Execute()
	  //r, err := apiClient.NERM.ConsolidationAPI.PatchDataRecord(context.Background(), id).DataRecords(dataRecords).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `ConsolidationAPI.PatchDataRecord``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    
}
```

[[Back to top]](#)

