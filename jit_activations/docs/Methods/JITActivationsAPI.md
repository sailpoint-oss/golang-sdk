---
id: v1-jit-activations
title: JITActivations
pagination_label: JITActivations
sidebar_label: JITActivations
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'JITActivations', 'V1JITActivations'] 
slug: /tools/sdk/go/jitactivations/methods/jit-activations
tags: ['SDK', 'Software Development Kit', 'JITActivations', 'V1JITActivations']
---

# JITActivationsAPI
  Use this API to start and manage Just-In-Time (JIT) Privileged activation workflows for entitlement connections.
 
All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**start-activate-workflow-v1**](#start-activate-workflow-v1) | **Post** `/jit-activations/v1/activate` | Start JIT activation workflow
[**start-deactivate-workflow-v1**](#start-deactivate-workflow-v1) | **Post** `/jit-activations/v1/deactivate` | Deactivate JIT activation workflow
[**start-extend-workflow-v1**](#start-extend-workflow-v1) | **Post** `/jit-activations/v1/extend` | Extend JIT activation workflow


## start-activate-workflow-v1
Start JIT activation workflow
Starts a JIT Privileged (JIT P) activation workflow for the given entitlement connection and duration.
The service performs quick validation; the workflow performs additional validation.

The response is returned with HTTP 202 Accepted while the workflow initializes.


[API Spec](https://developer.sailpoint.com/docs/api/start-activate-workflow-v-1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartActivateWorkflowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jitActivationActivateRequest** | [**JitActivationActivateRequest**](../models/jit-activation-activate-request) |  | 

### Return type

[**JitActivationActivateResponse**](../models/jit-activation-activate-response)

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
    jit_activations "github.com/sailpoint-oss/golang-sdk/v3/jit_activations"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    jitactivationactivaterequestJson := []byte(`{
          "activationPeriodMins" : 120,
          "connectionId" : "757fb803-9024-5861-e510-83a56e4c5bd3"
        }`) // JitActivationActivateRequest | 

    var jitActivationActivateRequest jit_activations.JitActivationActivateRequest
    if err := json.Unmarshal(jitactivationactivaterequestJson, &jitActivationActivateRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.JITActivationsAPI.StartActivateWorkflowV1(context.Background()).JitActivationActivateRequest(jitActivationActivateRequest).Execute()
	  //resp, r, err := apiClient.JITActivationsAPI.StartActivateWorkflowV1(context.Background()).JitActivationActivateRequest(jitActivationActivateRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `JITActivationsAPI.StartActivateWorkflowV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartActivateWorkflowV1`: JitActivationActivateResponse
    fmt.Fprintf(os.Stdout, "Response from `JITActivationsAPI.StartActivateWorkflowV1`: %v\n", resp)
}
```

[[Back to top]](#)

## start-deactivate-workflow-v1
Deactivate JIT activation workflow
Sends a signal to a running JIT Privileged (JIT P) activation workflow to deactivate.

This request cannot be applied to a workflow that does not exist or whose execution has already completed.
The client receives an error response in those cases.

The response is returned with HTTP 202 Accepted after the signal is sent.


[API Spec](https://developer.sailpoint.com/docs/api/start-deactivate-workflow-v-1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartDeactivateWorkflowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jitActivationDeactivateRequest** | [**JitActivationDeactivateRequest**](../models/jit-activation-deactivate-request) |  | 

### Return type

[**JitActivationDeactivateResponse**](../models/jit-activation-deactivate-response)

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
    jit_activations "github.com/sailpoint-oss/golang-sdk/v3/jit_activations"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    jitactivationdeactivaterequestJson := []byte(`{
          "connectionId" : "757fb803-9024-5861-e510-83a56e4c5bd3"
        }`) // JitActivationDeactivateRequest | 

    var jitActivationDeactivateRequest jit_activations.JitActivationDeactivateRequest
    if err := json.Unmarshal(jitactivationdeactivaterequestJson, &jitActivationDeactivateRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.JITActivationsAPI.StartDeactivateWorkflowV1(context.Background()).JitActivationDeactivateRequest(jitActivationDeactivateRequest).Execute()
	  //resp, r, err := apiClient.JITActivationsAPI.StartDeactivateWorkflowV1(context.Background()).JitActivationDeactivateRequest(jitActivationDeactivateRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `JITActivationsAPI.StartDeactivateWorkflowV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartDeactivateWorkflowV1`: JitActivationDeactivateResponse
    fmt.Fprintf(os.Stdout, "Response from `JITActivationsAPI.StartDeactivateWorkflowV1`: %v\n", resp)
}
```

[[Back to top]](#)

## start-extend-workflow-v1
Extend JIT activation workflow
Sends a signal to a running JIT Privileged (JIT P) activation workflow to extend the activation period
by the requested number of minutes.

This request cannot be applied to a workflow that does not exist or whose execution has already completed.
The client receives an error response in those cases.

The response is returned with HTTP 202 Accepted after the signal is sent.


[API Spec](https://developer.sailpoint.com/docs/api/start-extend-workflow-v-1)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartExtendWorkflowV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jitActivationExtendRequest** | [**JitActivationExtendRequest**](../models/jit-activation-extend-request) |  | 

### Return type

[**JitActivationExtendResponse**](../models/jit-activation-extend-response)

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
    jit_activations "github.com/sailpoint-oss/golang-sdk/v3/jit_activations"
	sailpoint "github.com/sailpoint-oss/golang-sdk/v3"
)

func main() {
    jitactivationextendrequestJson := []byte(`{
          "activationPeriodExtensionMins" : 120,
          "connectionId" : "757fb803-9024-5861-e510-83a56e4c5bd3"
        }`) // JitActivationExtendRequest | 

    var jitActivationExtendRequest jit_activations.JitActivationExtendRequest
    if err := json.Unmarshal(jitactivationextendrequestJson, &jitActivationExtendRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.JITActivationsAPI.StartExtendWorkflowV1(context.Background()).JitActivationExtendRequest(jitActivationExtendRequest).Execute()
	  //resp, r, err := apiClient.JITActivationsAPI.StartExtendWorkflowV1(context.Background()).JitActivationExtendRequest(jitActivationExtendRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `JITActivationsAPI.StartExtendWorkflowV1``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartExtendWorkflowV1`: JitActivationExtendResponse
    fmt.Fprintf(os.Stdout, "Response from `JITActivationsAPI.StartExtendWorkflowV1`: %v\n", resp)
}
```

[[Back to top]](#)

