# \LifecycleStatesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLifecycleStates**](LifecycleStatesApi.md#ListLifecycleStates) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Lifecycle State
[**UpdateLifecycleStates**](LifecycleStatesApi.md#UpdateLifecycleStates) | **Patch** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Update Lifecycle State



## ListLifecycleStates

> LifecycleState ListLifecycleStates(ctx, identityProfileId, lifecycleStateId).Execute()

Lifecycle State



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
    identityProfileId := "identityProfileId_example" // string | Identity Profile ID
    lifecycleStateId := "lifecycleStateId_example" // string | Lifecycle State ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.ListLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.ListLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLifecycleStates`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.ListLifecycleStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLifecycleStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLifecycleStates

> LifecycleState UpdateLifecycleStates(ctx, identityProfileId, lifecycleStateId).JsonPatchOperation(jsonPatchOperation).Execute()

Update Lifecycle State



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
    identityProfileId := "identityProfileId_example" // string | Identity Profile ID
    lifecycleStateId := "lifecycleStateId_example" // string | Lifecycle State ID
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.UpdateLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLifecycleStates`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.UpdateLifecycleStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLifecycleStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption  | 

### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

