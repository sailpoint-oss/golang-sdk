# \LifecycleStatesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLifecycleState**](LifecycleStatesAPI.md#CreateLifecycleState) | **Post** /identity-profiles/{identity-profile-id}/lifecycle-states | Create Lifecycle State
[**DeleteLifecycleState**](LifecycleStatesAPI.md#DeleteLifecycleState) | **Delete** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Delete Lifecycle State
[**GetLifecycleState**](LifecycleStatesAPI.md#GetLifecycleState) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Get Lifecycle State
[**GetLifecycleStates**](LifecycleStatesAPI.md#GetLifecycleStates) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states | Lists LifecycleStates
[**SetLifecycleState**](LifecycleStatesAPI.md#SetLifecycleState) | **Post** /identities/{identity-id}/set-lifecycle-state | Set Lifecycle State
[**UpdateLifecycleStates**](LifecycleStatesAPI.md#UpdateLifecycleStates) | **Patch** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Update Lifecycle State



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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity profile ID.
    lifecycleState := *openapiclient.NewLifecycleState("aName", "Technical Name") // LifecycleState | Lifecycle state to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.CreateLifecycleState(context.Background(), identityProfileId).LifecycleState(lifecycleState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.CreateLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLifecycleState`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.CreateLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycleState** | [**LifecycleState**](LifecycleState.md) | Lifecycle state to be created. | 

### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLifecycleState

> LifecyclestateDeleted DeleteLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

Delete Lifecycle State



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity profile ID.
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle state ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.DeleteLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLifecycleState`: LifecyclestateDeleted
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.DeleteLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity profile ID. | 
**lifecycleStateId** | **string** | Lifecycle state ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LifecyclestateDeleted**](LifecyclestateDeleted.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLifecycleState

> LifecycleState GetLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

Get Lifecycle State



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity profile ID.
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle state ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.GetLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLifecycleState`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.GetLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity profile ID. | 
**lifecycleStateId** | **string** | Lifecycle state ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLifecycleStates

> []LifecycleState GetLifecycleStates(ctx, identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()

Lists LifecycleStates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity profile ID.
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "created,modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.GetLifecycleStates(context.Background(), identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.GetLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLifecycleStates`: []LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.GetLifecycleStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity profile ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLifecycleStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** | 

### Return type

[**[]LifecycleState**](LifecycleState.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityId := "2c9180857893f1290178944561990364" // string | ID of the identity to update.
    setLifecycleStateRequest := *openapiclient.NewSetLifecycleStateRequest() // SetLifecycleStateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.SetLifecycleState(context.Background(), identityId).SetLifecycleStateRequest(setLifecycleStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.SetLifecycleState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLifecycleState`: SetLifecycleState200Response
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.SetLifecycleState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | ID of the identity to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setLifecycleStateRequest** | [**SetLifecycleStateRequest**](SetLifecycleStateRequest.md) |  | 

### Return type

[**SetLifecycleState200Response**](SetLifecycleState200Response.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity profile ID.
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle state ID.
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LifecycleStatesAPI.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.UpdateLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLifecycleStates`: LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.UpdateLifecycleStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProfileId** | **string** | Identity profile ID. | 
**lifecycleStateId** | **string** | Lifecycle state ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLifecycleStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption  | 

### Return type

[**LifecycleState**](LifecycleState.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

