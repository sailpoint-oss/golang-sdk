---
id: nerm-workflow-action-performer
title: WorkflowActionPerformer
pagination_label: WorkflowActionPerformer
sidebar_label: WorkflowActionPerformer
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowActionPerformer', 'NERMWorkflowActionPerformer'] 
slug: /tools/sdk/go/nerm/methods/workflow-action-performer
tags: ['SDK', 'Software Development Kit', 'WorkflowActionPerformer', 'NERMWorkflowActionPerformer']
---

# WorkflowActionPerformerAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-workflow-action-performer**](#create-workflow-action-performer) | **Post** `/workflow_action_performers` | Create a workflow action performer


## create-workflow-action-performer
Create a workflow action performer
Create a workflow action performer for an existing workflow action

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-workflow-action-performer)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowActionPerformerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowActionPerformerRequest** | [**CreateWorkflowActionPerformerRequest**](../models/create-workflow-action-performer-request) |  | 

### Return type

[**CreateWorkflowActionPerformer200Response**](../models/create-workflow-action-performer200-response)

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
    createworkflowactionperformerrequest := []byte(``) // CreateWorkflowActionPerformerRequest | 

    var createWorkflowActionPerformerRequest NERM.CreateWorkflowActionPerformerRequest
    if err := json.Unmarshal(createworkflowactionperformerrequest, &createWorkflowActionPerformerRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowActionPerformerAPI.CreateWorkflowActionPerformer(context.Background()).CreateWorkflowActionPerformerRequest(createWorkflowActionPerformerRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowActionPerformerAPI.CreateWorkflowActionPerformer(context.Background()).CreateWorkflowActionPerformerRequest(createWorkflowActionPerformerRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowActionPerformerAPI.CreateWorkflowActionPerformer``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowActionPerformer`: CreateWorkflowActionPerformer200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowActionPerformerAPI.CreateWorkflowActionPerformer`: %v\n", resp)
}
```

[[Back to top]](#)

