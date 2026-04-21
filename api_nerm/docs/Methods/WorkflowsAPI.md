---
id: nerm-workflows
title: Workflows
pagination_label: Workflows
sidebar_label: Workflows
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'Workflows', 'NERMWorkflows'] 
slug: /tools/sdk/go/nerm/methods/workflows
tags: ['SDK', 'Software Development Kit', 'Workflows', 'NERMWorkflows']
---

# WorkflowsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create-automated-workflow**](#create-automated-workflow) | **Post** `/workflows/automated_workflows` | Create an automated workflow
[**create-batch-workflow**](#create-batch-workflow) | **Post** `/workflows/batch_workflows` | Create a batch workflow
[**create-create-workflow**](#create-create-workflow) | **Post** `/workflows/create_workflows` | Create a create workflow
[**create-login-workflow**](#create-login-workflow) | **Post** `/workflows/login_workflows` | Create a login workflow
[**create-password-update-workflow**](#create-password-update-workflow) | **Post** `/workflows/password_reset_workflows` | Create a password reset workflow
[**create-registration-workflow**](#create-registration-workflow) | **Post** `/workflows/registration_workflows` | Create a registration workflow
[**create-update-workflow**](#create-update-workflow) | **Post** `/workflows/update_workflows` | Create an update workflow


## create-automated-workflow
Create an automated workflow
Create an automated workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-automated-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomatedWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAutomatedWorkflowRequest** | [**CreateAutomatedWorkflowRequest**](../models/create-automated-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createautomatedworkflowrequest := []byte(``) // CreateAutomatedWorkflowRequest | 

    var createAutomatedWorkflowRequest NERM.CreateAutomatedWorkflowRequest
    if err := json.Unmarshal(createautomatedworkflowrequest, &createAutomatedWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateAutomatedWorkflow(context.Background()).CreateAutomatedWorkflowRequest(createAutomatedWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateAutomatedWorkflow(context.Background()).CreateAutomatedWorkflowRequest(createAutomatedWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateAutomatedWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomatedWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateAutomatedWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-batch-workflow
Create a batch workflow
Create a batch workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-batch-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBatchWorkflowRequest** | [**CreateBatchWorkflowRequest**](../models/create-batch-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createbatchworkflowrequest := []byte(``) // CreateBatchWorkflowRequest | 

    var createBatchWorkflowRequest NERM.CreateBatchWorkflowRequest
    if err := json.Unmarshal(createbatchworkflowrequest, &createBatchWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateBatchWorkflow(context.Background()).CreateBatchWorkflowRequest(createBatchWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateBatchWorkflow(context.Background()).CreateBatchWorkflowRequest(createBatchWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateBatchWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBatchWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateBatchWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-create-workflow
Create a create workflow
Create a create workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-create-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreateWorkflowRequest** | [**CreateCreateWorkflowRequest**](../models/create-create-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createcreateworkflowrequest := []byte(``) // CreateCreateWorkflowRequest | 

    var createCreateWorkflowRequest NERM.CreateCreateWorkflowRequest
    if err := json.Unmarshal(createcreateworkflowrequest, &createCreateWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateCreateWorkflow(context.Background()).CreateCreateWorkflowRequest(createCreateWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateCreateWorkflow(context.Background()).CreateCreateWorkflowRequest(createCreateWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateCreateWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreateWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateCreateWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-login-workflow
Create a login workflow
Create a login workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-login-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoginWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoginWorkflowRequest** | [**CreateLoginWorkflowRequest**](../models/create-login-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createloginworkflowrequest := []byte(``) // CreateLoginWorkflowRequest | 

    var createLoginWorkflowRequest NERM.CreateLoginWorkflowRequest
    if err := json.Unmarshal(createloginworkflowrequest, &createLoginWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateLoginWorkflow(context.Background()).CreateLoginWorkflowRequest(createLoginWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateLoginWorkflow(context.Background()).CreateLoginWorkflowRequest(createLoginWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateLoginWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoginWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateLoginWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-password-update-workflow
Create a password reset workflow
Create a password reset workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-password-update-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLoginWorkflowRequest** | [**CreateLoginWorkflowRequest**](../models/create-login-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createloginworkflowrequest := []byte(``) // CreateLoginWorkflowRequest | 

    var createLoginWorkflowRequest NERM.CreateLoginWorkflowRequest
    if err := json.Unmarshal(createloginworkflowrequest, &createLoginWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreatePasswordUpdateWorkflow(context.Background()).CreateLoginWorkflowRequest(createLoginWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreatePasswordUpdateWorkflow(context.Background()).CreateLoginWorkflowRequest(createLoginWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreatePasswordUpdateWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordUpdateWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreatePasswordUpdateWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-registration-workflow
Create a registration workflow
Create a registration workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-registration-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistrationWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCreateWorkflowRequest** | [**CreateCreateWorkflowRequest**](../models/create-create-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createcreateworkflowrequest := []byte(``) // CreateCreateWorkflowRequest | 

    var createCreateWorkflowRequest NERM.CreateCreateWorkflowRequest
    if err := json.Unmarshal(createcreateworkflowrequest, &createCreateWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateRegistrationWorkflow(context.Background()).CreateCreateWorkflowRequest(createCreateWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateRegistrationWorkflow(context.Background()).CreateCreateWorkflowRequest(createCreateWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateRegistrationWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistrationWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateRegistrationWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

## create-update-workflow
Create an update workflow
Create an update workflow

[API Spec](https://developer.sailpoint.com/docs/api/NERM/create-update-workflow)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateWorkflowRequest** | [**CreateUpdateWorkflowRequest**](../models/create-update-workflow-request) |  | 

### Return type

[**CreateCreateWorkflow200Response**](../models/create-create-workflow200-response)

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
    createupdateworkflowrequest := []byte(``) // CreateUpdateWorkflowRequest | 

    var createUpdateWorkflowRequest NERM.CreateUpdateWorkflowRequest
    if err := json.Unmarshal(createupdateworkflowrequest, &createUpdateWorkflowRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowsAPI.CreateUpdateWorkflow(context.Background()).CreateUpdateWorkflowRequest(createUpdateWorkflowRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowsAPI.CreateUpdateWorkflow(context.Background()).CreateUpdateWorkflowRequest(createUpdateWorkflowRequest).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateUpdateWorkflow``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateWorkflow`: CreateCreateWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateUpdateWorkflow`: %v\n", resp)
}
```

[[Back to top]](#)

