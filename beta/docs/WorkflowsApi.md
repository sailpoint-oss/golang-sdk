# \WorkflowsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelWorkflowExecution**](WorkflowsApi.md#CancelWorkflowExecution) | **Post** /workflow-executions/{id}/cancel | Cancel Workflow Execution by ID
[**CreateWorkflow**](WorkflowsApi.md#CreateWorkflow) | **Post** /workflows | Create Workflow
[**DeleteWorkflow**](WorkflowsApi.md#DeleteWorkflow) | **Delete** /workflows/{id} | Delete Workflow By Id
[**GetWorkflow**](WorkflowsApi.md#GetWorkflow) | **Get** /workflows/{id} | Get Workflow By Id
[**GetWorkflowExecution**](WorkflowsApi.md#GetWorkflowExecution) | **Get** /workflow-executions/{id} | Get a Workflow Execution
[**GetWorkflowExecutionHistory**](WorkflowsApi.md#GetWorkflowExecutionHistory) | **Get** /workflow-executions/{id}/history | Get Workflow Execution History
[**ListCompleteWorkflowLibrary**](WorkflowsApi.md#ListCompleteWorkflowLibrary) | **Get** /workflow-library | List Complete Workflow Library
[**ListWorkflowExecutions**](WorkflowsApi.md#ListWorkflowExecutions) | **Get** /workflows/{id}/executions | List Workflow Executions
[**ListWorkflowLibraryActions**](WorkflowsApi.md#ListWorkflowLibraryActions) | **Get** /workflow-library/actions | List Workflow Library Actions
[**ListWorkflowLibraryOperators**](WorkflowsApi.md#ListWorkflowLibraryOperators) | **Get** /workflow-library/operators | List Workflow Library Operators
[**ListWorkflowLibraryTriggers**](WorkflowsApi.md#ListWorkflowLibraryTriggers) | **Get** /workflow-library/triggers | List Workflow Library Triggers
[**ListWorkflows**](WorkflowsApi.md#ListWorkflows) | **Get** /workflows | List Workflows
[**PatchWorkflow**](WorkflowsApi.md#PatchWorkflow) | **Patch** /workflows/{id} | Patch Workflow
[**PostExternalExecuteWorkflow**](WorkflowsApi.md#PostExternalExecuteWorkflow) | **Post** /workflows/execute/external/{id} | Execute Workflow via External Trigger
[**PostWorkflowExternalTrigger**](WorkflowsApi.md#PostWorkflowExternalTrigger) | **Post** /workflows/{id}/external/oauth-clients | Generate External Trigger OAuth Client
[**TestExternalExecuteWorkflow**](WorkflowsApi.md#TestExternalExecuteWorkflow) | **Post** /workflows/execute/external/{id}/test | Test Workflow via External Trigger
[**TestWorkflow**](WorkflowsApi.md#TestWorkflow) | **Post** /workflows/{id}/test | Test Workflow By Id
[**UpdateWorkflow**](WorkflowsApi.md#UpdateWorkflow) | **Put** /workflows/{id} | Update Workflow



## CancelWorkflowExecution

> CancelWorkflowExecution(ctx, id).Execute()

Cancel Workflow Execution by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | The workflow execution ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CancelWorkflowExecution(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CancelWorkflowExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The workflow execution ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelWorkflowExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflow

> Workflow CreateWorkflow(ctx).CreateWorkflowRequest(createWorkflowRequest).Execute()

Create Workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createWorkflowRequest := *openapiclient.NewCreateWorkflowRequest("Send Email", *openapiclient.NewWorkflowBodyOwner()) // CreateWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflow(context.Background()).CreateWorkflowRequest(createWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRequest** | [**CreateWorkflowRequest**](CreateWorkflowRequest.md) |  | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflow(ctx, id).Execute()

Delete Workflow By Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the Workflow

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.DeleteWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.DeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the Workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> Workflow GetWorkflow(ctx, id).Execute()

Get Workflow By Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workflow**](Workflow.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecution

> map[string]interface{} GetWorkflowExecution(ctx, id).Execute()

Get a Workflow Execution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow execution

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowExecution(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowExecution`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow execution | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowExecutionHistory

> []WorkflowExecutionEvent GetWorkflowExecutionHistory(ctx, id).Execute()

Get Workflow Execution History



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow execution

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowExecutionHistory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowExecutionHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowExecutionHistory`: []WorkflowExecutionEvent
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowExecutionHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow execution | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExecutionHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]WorkflowExecutionEvent**](WorkflowExecutionEvent.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompleteWorkflowLibrary

> []ListCompleteWorkflowLibrary200ResponseInner ListCompleteWorkflowLibrary(ctx).Limit(limit).Offset(offset).Execute()

List Complete Workflow Library



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListCompleteWorkflowLibrary(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListCompleteWorkflowLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompleteWorkflowLibrary`: []ListCompleteWorkflowLibrary200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListCompleteWorkflowLibrary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompleteWorkflowLibraryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]

### Return type

[**[]ListCompleteWorkflowLibrary200ResponseInner**](ListCompleteWorkflowLibrary200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowExecutions

> []WorkflowExecution ListWorkflowExecutions(ctx, id).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()

List Workflow Executions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "status eq "Failed"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **startTime**: *eq, lt, le, gt, ge*  **status**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowExecutions(context.Background(), id).Limit(limit).Offset(offset).Count(count).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowExecutions`: []WorkflowExecution
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowExecutions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **startTime**: *eq, lt, le, gt, ge*  **status**: *eq* | 

### Return type

[**[]WorkflowExecution**](WorkflowExecution.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowLibraryActions

> []WorkflowLibraryAction ListWorkflowLibraryActions(ctx).Limit(limit).Offset(offset).Filters(filters).Execute()

List Workflow Library Actions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    filters := "id eq "sp:create-campaign"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowLibraryActions(context.Background()).Limit(limit).Offset(offset).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowLibraryActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowLibraryActions`: []WorkflowLibraryAction
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowLibraryActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowLibraryActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* | 

### Return type

[**[]WorkflowLibraryAction**](WorkflowLibraryAction.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowLibraryOperators

> []WorkflowLibraryOperator ListWorkflowLibraryOperators(ctx).Execute()

List Workflow Library Operators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowLibraryOperators(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowLibraryOperators``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowLibraryOperators`: []WorkflowLibraryOperator
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowLibraryOperators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowLibraryOperatorsRequest struct via the builder pattern


### Return type

[**[]WorkflowLibraryOperator**](WorkflowLibraryOperator.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowLibraryTriggers

> []WorkflowLibraryTrigger ListWorkflowLibraryTriggers(ctx).Limit(limit).Offset(offset).Filters(filters).Execute()

List Workflow Library Triggers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    filters := "id eq "idn:identity-attributes-changed"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowLibraryTriggers(context.Background()).Limit(limit).Offset(offset).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowLibraryTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowLibraryTriggers`: []WorkflowLibraryTrigger
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowLibraryTriggers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowLibraryTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq* | 

### Return type

[**[]WorkflowLibraryTrigger**](WorkflowLibraryTrigger.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> []Workflow ListWorkflows(ctx).Execute()

List Workflows



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflows`: []Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


### Return type

[**[]Workflow**](Workflow.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchWorkflow

> Workflow PatchWorkflow(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch Workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the Workflow
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.PatchWorkflow(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.PatchWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.PatchWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the Workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) |  | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostExternalExecuteWorkflow

> PostExternalExecuteWorkflow200Response PostExternalExecuteWorkflow(ctx, id).PostExternalExecuteWorkflowRequest(postExternalExecuteWorkflowRequest).Execute()

Execute Workflow via External Trigger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow
    postExternalExecuteWorkflowRequest := *openapiclient.NewPostExternalExecuteWorkflowRequest() // PostExternalExecuteWorkflowRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.PostExternalExecuteWorkflow(context.Background(), id).PostExternalExecuteWorkflowRequest(postExternalExecuteWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.PostExternalExecuteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostExternalExecuteWorkflow`: PostExternalExecuteWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.PostExternalExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExternalExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postExternalExecuteWorkflowRequest** | [**PostExternalExecuteWorkflowRequest**](PostExternalExecuteWorkflowRequest.md) |  | 

### Return type

[**PostExternalExecuteWorkflow200Response**](PostExternalExecuteWorkflow200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostWorkflowExternalTrigger

> WorkflowOAuthClient PostWorkflowExternalTrigger(ctx, id).Execute()

Generate External Trigger OAuth Client



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.PostWorkflowExternalTrigger(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.PostWorkflowExternalTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostWorkflowExternalTrigger`: WorkflowOAuthClient
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.PostWorkflowExternalTrigger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostWorkflowExternalTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowOAuthClient**](WorkflowOAuthClient.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestExternalExecuteWorkflow

> TestExternalExecuteWorkflow200Response TestExternalExecuteWorkflow(ctx, id).TestExternalExecuteWorkflowRequest(testExternalExecuteWorkflowRequest).Execute()

Test Workflow via External Trigger



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow
    testExternalExecuteWorkflowRequest := *openapiclient.NewTestExternalExecuteWorkflowRequest() // TestExternalExecuteWorkflowRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.TestExternalExecuteWorkflow(context.Background(), id).TestExternalExecuteWorkflowRequest(testExternalExecuteWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.TestExternalExecuteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestExternalExecuteWorkflow`: TestExternalExecuteWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.TestExternalExecuteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestExternalExecuteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testExternalExecuteWorkflowRequest** | [**TestExternalExecuteWorkflowRequest**](TestExternalExecuteWorkflowRequest.md) |  | 

### Return type

[**TestExternalExecuteWorkflow200Response**](TestExternalExecuteWorkflow200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestWorkflow

> TestWorkflow200Response TestWorkflow(ctx, id).TestWorkflowRequest(testWorkflowRequest).Execute()

Test Workflow By Id



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the workflow
    testWorkflowRequest := *openapiclient.NewTestWorkflowRequest(map[string]interface{}(123)) // TestWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.TestWorkflow(context.Background(), id).TestWorkflowRequest(testWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.TestWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestWorkflow`: TestWorkflow200Response
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.TestWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testWorkflowRequest** | [**TestWorkflowRequest**](TestWorkflowRequest.md) |  | 

### Return type

[**TestWorkflow200Response**](TestWorkflow200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflow

> Workflow UpdateWorkflow(ctx, id).WorkflowBody(workflowBody).Execute()

Update Workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "c17bea3a-574d-453c-9e04-4365fbf5af0b" // string | Id of the Workflow
    workflowBody := *openapiclient.NewWorkflowBody() // WorkflowBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.UpdateWorkflow(context.Background(), id).WorkflowBody(workflowBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.UpdateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.UpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Id of the Workflow | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowBody** | [**WorkflowBody**](WorkflowBody.md) |  | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

