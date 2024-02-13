# \LifecycleStatesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLifecycleState**](LifecycleStatesAPI.md#CreateLifecycleState) | **Post** /identity-profiles/{identity-profile-id}/lifecycle-states | Create Lifecycle State
[**DeleteLifecycleState**](LifecycleStatesAPI.md#DeleteLifecycleState) | **Delete** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Delete Lifecycle State by ID
[**GetLifecycleState**](LifecycleStatesAPI.md#GetLifecycleState) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Retrieves Lifecycle State
[**ListLifecycleStates**](LifecycleStatesAPI.md#ListLifecycleStates) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states | Lists LifecycleStates
[**SetLifecycleState**](LifecycleStatesAPI.md#SetLifecycleState) | **Post** /identities/{identity-id}/set-lifecycle-state | Set Lifecycle State
[**UpdateLifecycleStates**](LifecycleStatesAPI.md#UpdateLifecycleStates) | **Patch** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Update Lifecycle State



## Create Lifecycle State

> LifecycleState CreateLifecycleState(ctx, identityProfileId).LifecycleState(lifecycleState).Execute()

This API creates a new Lifecycle State.
A token with ORG_ADMIN or API authority is required to call this API.

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
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | Identity Profile ID
    lifecycleState := *sailpoint.NewLifecycleState("aName", "Technical Name") // LifecycleState | Lifecycle State

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.CreateLifecycleState(context.Background(), identityProfileId).LifecycleState(lifecycleState).Execute()
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
**identityProfileId** | **string** | Identity Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLifecycleStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycleState** | [**LifecycleState**](LifecycleState.md) | Lifecycle State | 

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


## Delete Lifecycle State by ID

> LifecyclestateDeleted DeleteLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

This endpoint deletes the Lifecycle State using its ID.
A token with API, or ORG_ADMIN authority is required to call this API.

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
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity Profile ID
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle State ID

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
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
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

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


## Retrieves Lifecycle State

> LifecycleState GetLifecycleState(ctx, identityProfileId, lifecycleStateId).Execute()

This endpoint retrieves a Lifecycle State.
A token with ORG_ADMIN or API authority is required to call this API.

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
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity Profile ID
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle State ID

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()
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
**identityProfileId** | **string** | Identity Profile ID | 
**lifecycleStateId** | **string** | Lifecycle State ID | 

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


## Lists LifecycleStates

> []LifecycleState ListLifecycleStates(ctx, identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()

This end-point lists all the LifecycleStates associated with IdentityProfiles.
A token with API, or ORG_ADMIN authority is required to call this API.

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
    identityProfileId := "ef38f94347e94562b5bb8424a56397d8" // string | The IdentityProfile id
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "created,modified" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.ListLifecycleStates(context.Background(), identityProfileId).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.ListLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLifecycleStates`: []LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.ListLifecycleStates`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set Lifecycle State

> SetLifecycleState200Response SetLifecycleState(ctx, identityId).SetLifecycleStateRequest(setLifecycleStateRequest).Execute()

This endpoint will set/update an identity's lifecycle state to the one provided and updates the corresponding Identity Profile.
A token with ORG_ADMIN or API authority is required to call this API.

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
    identityId := "2c9180857893f1290178944561990364" // string | The ID of the identity to update
    setLifecycleStateRequest := *sailpoint.NewSetLifecycleStateRequest() // SetLifecycleStateRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.SetLifecycleState(context.Background(), identityId).SetLifecycleStateRequest(setLifecycleStateRequest).Execute()
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
**identityId** | **string** | The ID of the identity to update | 

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


## Update Lifecycle State

> LifecycleState UpdateLifecycleStates(ctx, identityProfileId, lifecycleStateId).JsonPatchOperation(jsonPatchOperation).Execute()

This endpoint updates individual Lifecycle State fields using the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.
A token with ORG_ADMIN or API authority is required to call this API.

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
    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238" // string | Identity Profile ID
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8" // string | Lifecycle State ID
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).JsonPatchOperation(jsonPatchOperation).Execute()
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

