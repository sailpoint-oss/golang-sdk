---
id: nerm-job-status
title: JobStatus
pagination_label: JobStatus
sidebar_label: JobStatus
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'JobStatus', 'NERMJobStatus'] 
slug: /tools/sdk/go/nerm/methods/job-status
tags: ['SDK', 'Software Development Kit', 'JobStatus', 'NERMJobStatus']
---

# JobStatusAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-job-status**](#get-job-status) | **Get** `/job_status` | Get bulk job status


## get-job-status
Get bulk job status
Retrieve the status of a bulk job

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-job-status)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | The id of the job returned from a POST or PATCH endpoint that resulted in a job being created | 

### Return type

[**GetJobStatus200Response**](../models/get-job-status200-response)

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
    jobId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | The id of the job returned from a POST or PATCH endpoint that resulted in a job being created # string | The id of the job returned from a POST or PATCH endpoint that resulted in a job being created

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.JobStatusAPI.GetJobStatus(context.Background()).JobId(jobId).Execute()
	  //resp, r, err := apiClient.NERM.JobStatusAPI.GetJobStatus(context.Background()).JobId(jobId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `JobStatusAPI.GetJobStatus``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobStatus`: GetJobStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `JobStatusAPI.GetJobStatus`: %v\n", resp)
}
```

[[Back to top]](#)

