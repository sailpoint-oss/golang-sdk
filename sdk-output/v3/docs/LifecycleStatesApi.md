# \LifecycleStatesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLifecycleState**](LifecycleStatesApi.md#CreateLifecycleState) | **Post** /identity-profiles/{identity-profile-id}/lifecycle-states | Create Lifecycle State
[**DeleteLifecycleState**](LifecycleStatesApi.md#DeleteLifecycleState) | **Delete** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Delete Lifecycle State by ID
[**GetLifecycleState**](LifecycleStatesApi.md#GetLifecycleState) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Retrieves Lifecycle State
[**ListLifecycleStates**](LifecycleStatesApi.md#ListLifecycleStates) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states | Lists LifecycleStates
[**SetLifecycleState**](LifecycleStatesApi.md#SetLifecycleState) | **Post** /identities/{identity-id}/set-lifecycle-state | Set Lifecycle State
[**UpdateLifecycleStates**](LifecycleStatesApi.md#UpdateLifecycleStates) | **Patch** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Update Lifecycle State



## CreateLifecycleState

> LifecycleState CreateLifecycleState(ctx, identityProfileId).LifecycleState(lifecycleState).Execute()

Create Lifecycle State



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
    lifecycleState := *openapiclient.NewLifecycleState("aName", "Technical Name") // LifecycleState | Lifecycle State

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.CreateLifecycleState(context.Background(), identityProfileId).LifecycleState(lifecycleState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.CreateLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLifecycleState`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.CreateLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycleState** | [**LifecycleState**](LifecycleState.md) | Lifecycle State | 

### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLifecycleState

> BaseReferenceDto DeleteLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

Delete Lifecycle State by ID



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
    resp, r, err := apiClient.LifecycleStatesApi.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.DeleteLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLifecycleState`: BaseReferenceDto
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.DeleteLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseReferenceDto**](BaseReferenceDto.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLifecycleState

> LifecycleState GetLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

Retrieves Lifecycle State



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
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Profile ID
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle State ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.GetLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLifecycleState`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.GetLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLifecycleStateRequest struct via the builder pattern


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


## ListLifecycleStates

> []LifecycleState ListLifecycleStates(ctx, identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()

Lists LifecycleStates



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
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The IdentityProfile id
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "created,modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.ListLifecycleStates(context.Background(), identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.ListLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLifecycleStates`: []LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.ListLifecycleStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | The IdentityProfile id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLifecycleStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]LifecycleState**](LifecycleState.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLifecycleState

> SetLifecycleState200Response SetLifecycleState(ctx, identityId).SetLifecycleStateRequest(setLifecycleStateRequest).Execute()

Set Lifecycle State



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
    identityId := "2c9180857893f1290178944561990364" // string | The ID of the identity to update
    setLifecycleStateRequest := *openapiclient.NewSetLifecycleStateRequest() // SetLifecycleStateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesApi.SetLifecycleState(context.Background(), identityId).SetLifecycleStateRequest(setLifecycleStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesApi.SetLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLifecycleState`: SetLifecycleState200Response
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesApi.SetLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The ID of the identity to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLifecycleStateRequest** | [**SetLifecycleStateRequest**](SetLifecycleStateRequest.md) |  | 

### Return type

[**SetLifecycleState200Response**](SetLifecycleState200Response.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
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

