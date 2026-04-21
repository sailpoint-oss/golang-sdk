---
id: nerm-workflow-sessions
title: WorkflowSessions
pagination_label: WorkflowSessions
sidebar_label: WorkflowSessions
sidebar_class_name: gosdk
keywords: ['go', 'Golang', 'sdk', 'WorkflowSessions', 'NERMWorkflowSessions'] 
slug: /tools/sdk/go/nerm/methods/workflow-sessions
tags: ['SDK', 'Software Development Kit', 'WorkflowSessions', 'NERMWorkflowSessions']
---

# WorkflowSessionsAPI
   
All URIs are relative to *https://acmeco.nonemployee.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get-workflow-session**](#get-workflow-session) | **Get** `/workflow_sessions/{id}` | Find workflow session
[**get-workflow-session-upload**](#get-workflow-session-upload) | **Get** `/workflow_sessions/{id}/upload/{attribute_id}` | Retrieves workflow session attachment URL
[**get-workflow-sessions**](#get-workflow-sessions) | **Get** `/workflow_sessions` | Get workflow sessions
[**patch-workflow-session**](#patch-workflow-session) | **Patch** `/workflow_sessions/{id}` | Update a workflow session
[**submit-workflow-session**](#submit-workflow-session) | **Post** `/workflow_sessions` | Create a workflow session
[**submit-workflow-session-upload**](#submit-workflow-session-upload) | **Post** `/workflow_sessions/{id}/upload/{attribute_id}` | Uploads workflow session attachment


## get-workflow-session
Find workflow session
Find workflow session by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-workflow-session)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubmitWorkflowSession200Response**](../models/submit-workflow-session200-response)

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
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSession(context.Background(), id).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSession(context.Background(), id).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.GetWorkflowSession``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSession`: SubmitWorkflowSession200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.GetWorkflowSession`: %v\n", resp)
}
```

[[Back to top]](#)

## get-workflow-session-upload
Retrieves workflow session attachment URL
Retrieves the URL of an attachment attribute value from a workflow session

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-workflow-session-upload)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 
**attributeId** | **string** | The id of the attachment attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSessionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Url**](../models/url)

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
    attributeId := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | The id of the attachment attribute # string | The id of the attachment attribute

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSessionUpload(context.Background(), id, attributeId).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSessionUpload(context.Background(), id, attributeId).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.GetWorkflowSessionUpload``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSessionUpload`: Url
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.GetWorkflowSessionUpload`: %v\n", resp)
}
```

[[Back to top]](#)

## get-workflow-sessions
Get workflow sessions
Get workflow sessions

[API Spec](https://developer.sailpoint.com/docs/api/NERM/get-workflow-sessions)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | UID of the object to retrieve, update, or delete.  A UID or \&quot;specified identifier\&quot; is a string typically in \&quot;snake_case\&quot; format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of items to return. | 
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **order** | **string** | The field to order results by. | 
 **profileId** | **string** | Profile ID to filter by | 

 **workflowId** | **string** | Workflow ID for filtering | 
 **requesterId** | **string** | Requester ID for filtering | 
 **status** | **string** | filter by workflow session status | 
 **metadata** | **bool** | Returns batching metadata in the response. This includes &#x60;total&#x60; as the total quantity, &#x60;next&#x60; as the path of the following query url, &#x60;limit&#x60; and &#x60;after_id&#x60; (if requested) with the next following id (null if it is the last \&quot;page\&quot;). | [default to false]

### Return type

[**GetWorkflowSessions200Response**](../models/get-workflow-sessions200-response)

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
    limit := 5 // int32 | The maximum number of items to return. (optional) # int32 | The maximum number of items to return. (optional)
    offset := 5 // int32 | The number of items to skip before starting to collect the result set. (optional) # int32 | The number of items to skip before starting to collect the result set. (optional)
    order := `created_at` // string | The field to order results by. (optional) # string | The field to order results by. (optional)
    profileId := `4e480441-451d-47d9-87c2-9a0f0fe135eb` // string | Profile ID to filter by (optional) # string | Profile ID to filter by (optional)
    uid := `middle_initial_attribute` // string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional) # string | UID of the object to retrieve, update, or delete.  A UID or \"specified identifier\" is a string typically in \"snake_case\" format that provides a human-readable description of the record.  They are commonly used to ensure sandbox, qa, staging and production tenants have the identical configuration items loaded.  Every record has a UID assigned when persisted. When not specified the system assigns one by default.  A default value looks like a 32 character string of random hexadecimal characters. (optional)
    workflowId := `bba9cfb2-96c1-4acb-ac79-a21732527265` // string | Workflow ID for filtering (optional) # string | Workflow ID for filtering (optional)
    requesterId := `c5e1dd38-7e29-464f-a0da-0c0d886d022a` // string | Requester ID for filtering (optional) # string | Requester ID for filtering (optional)
    status := `pending approval` // string | filter by workflow session status (optional) # string | filter by workflow session status (optional)
    metadata := true // bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false) # bool | Returns batching metadata in the response. This includes `total` as the total quantity, `next` as the path of the following query url, `limit` and `after_id` (if requested) with the next following id (null if it is the last \"page\"). (optional) (default to false)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSessions(context.Background(), uid).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.GetWorkflowSessions(context.Background(), uid).Limit(limit).Offset(offset).Order(order).ProfileId(profileId).WorkflowId(workflowId).RequesterId(requesterId).Status(status).Metadata(metadata).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.GetWorkflowSessions``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowSessions`: GetWorkflowSessions200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.GetWorkflowSessions`: %v\n", resp)
}
```

[[Back to top]](#)

## patch-workflow-session
Update a workflow session
Update a workflow session by id

[API Spec](https://developer.sailpoint.com/docs/api/NERM/patch-workflow-session)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkflowSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **submitWorkflowSessionRequest** | [**SubmitWorkflowSessionRequest**](../models/submit-workflow-session-request) |  | 
 **run** | **bool** | Will run the created/updated workflow session if successful | [default to false]

### Return type

[**SubmitWorkflowSession200Response**](../models/submit-workflow-session200-response)

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
    submitworkflowsessionrequest := []byte(``) // SubmitWorkflowSessionRequest | 
    run := false // bool | Will run the created/updated workflow session if successful (optional) (default to false) # bool | Will run the created/updated workflow session if successful (optional) (default to false)

    var submitWorkflowSessionRequest NERM.SubmitWorkflowSessionRequest
    if err := json.Unmarshal(submitworkflowsessionrequest, &submitWorkflowSessionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.PatchWorkflowSession(context.Background(), id).SubmitWorkflowSessionRequest(submitWorkflowSessionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.PatchWorkflowSession(context.Background(), id).SubmitWorkflowSessionRequest(submitWorkflowSessionRequest).Run(run).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.PatchWorkflowSession``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchWorkflowSession`: SubmitWorkflowSession200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.PatchWorkflowSession`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-workflow-session
Create a workflow session
Create a workflow session

[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-workflow-session)

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWorkflowSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitWorkflowSessionRequest** | [**SubmitWorkflowSessionRequest**](../models/submit-workflow-session-request) |  | 
 **run** | **bool** | Will run the created/updated workflow session if successful | [default to false]

### Return type

[**SubmitWorkflowSession200Response**](../models/submit-workflow-session200-response)

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
    submitworkflowsessionrequest := []byte(``) // SubmitWorkflowSessionRequest | 
    run := false // bool | Will run the created/updated workflow session if successful (optional) (default to false) # bool | Will run the created/updated workflow session if successful (optional) (default to false)

    var submitWorkflowSessionRequest NERM.SubmitWorkflowSessionRequest
    if err := json.Unmarshal(submitworkflowsessionrequest, &submitWorkflowSessionRequest); err != nil {
      fmt.Println("Error:", err)
      return
    }
    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.SubmitWorkflowSession(context.Background()).SubmitWorkflowSessionRequest(submitWorkflowSessionRequest).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.SubmitWorkflowSession(context.Background()).SubmitWorkflowSessionRequest(submitWorkflowSessionRequest).Run(run).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.SubmitWorkflowSession``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitWorkflowSession`: SubmitWorkflowSession200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.SubmitWorkflowSession`: %v\n", resp)
}
```

[[Back to top]](#)

## submit-workflow-session-upload
Uploads workflow session attachment
Uploads a new attachment attribute value to a workflow session. The upload must be a FORM data type; this is not a JSON API. The upload must include the binary content of the payload under the 'file' named form element. The upload must not attempt to include the file name or its content type as a other form or JSON as parameters. The upload must not attempt to upload the file body as the POST body payload; it has to arrive as a FORM parameter. Do not use a `File/Binary` payload type for the POST operation in your API client.


[API Spec](https://developer.sailpoint.com/docs/api/NERM/submit-workflow-session-upload)

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the object to retrieve, update, or delete | 
**attributeId** | **string** | The id of the attachment attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWorkflowSessionUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**Url**](../models/url)

### HTTP request headers

- **Content-Type**: multipart/form-data
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
    attributeId := `1246d8b3-ac29-4015-8154-dea4434a73fa` // string | The id of the attachment attribute # string | The id of the attachment attribute
    file := BINARY_DATA_HERE // *os.File |  (optional) # *os.File |  (optional)

    

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.NERM.WorkflowSessionsAPI.SubmitWorkflowSessionUpload(context.Background(), id, attributeId).Execute()
	  //resp, r, err := apiClient.NERM.WorkflowSessionsAPI.SubmitWorkflowSessionUpload(context.Background(), id, attributeId).File(file).Execute()
    if err != nil {
	    fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSessionsAPI.SubmitWorkflowSessionUpload``: %v\n", err)
	    fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitWorkflowSessionUpload`: Url
    fmt.Fprintf(os.Stdout, "Response from `WorkflowSessionsAPI.SubmitWorkflowSessionUpload`: %v\n", resp)
}
```

[[Back to top]](#)

