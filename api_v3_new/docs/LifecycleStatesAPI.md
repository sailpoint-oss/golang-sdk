# \LifecycleStatesAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLifecycleState**](LifecycleStatesAPI.md#create-lifecycle-state) | **Post** /identity-profiles/{identity-profile-id}/lifecycle-states | Create Lifecycle State
[**DeleteLifecycleState**](LifecycleStatesAPI.md#delete-lifecycle-state) | **Delete** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Delete Lifecycle State by ID
[**GetLifecycleState**](LifecycleStatesAPI.md#get-lifecycle-state) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Retrieves Lifecycle State
[**ListLifecycleStates**](LifecycleStatesAPI.md#list-lifecycle-states) | **Get** /identity-profiles/{identity-profile-id}/lifecycle-states | Lists LifecycleStates
[**SetLifecycleState**](LifecycleStatesAPI.md#set-lifecycle-state) | **Post** /identities/{identity-id}/set-lifecycle-state | Set Lifecycle State
[**UpdateLifecycleStates**](LifecycleStatesAPI.md#update-lifecycle-states) | **Patch** /identity-profiles/{identity-profile-id}/lifecycle-states/{lifecycle-state-id} | Update Lifecycle State



## create-lifecycle-state


This API creates a new Lifecycle State.
A token with ORG_ADMIN or API authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityProfileId | **string** | True  | Identity Profile ID
 Body  | lifecycleState | [**LifecycleState**](LifecycleState.md) | True  | Lifecycle State


### Return type

[**LifecycleState**](LifecycleState.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
201 | Created LifecycleState object. | LifecycleState
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //CreateLifecycleState

    identityProfileId := "ef38f94347e94562b5bb8424a56397d8"
    lifecycleState := *sailpoint.NewLifecycleState("aName", "Technical Name")



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




## delete-lifecycle-state


This endpoint deletes the Lifecycle State using its ID.
A token with API, or ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityProfileId | **string** | True  | Identity Profile ID
Path   | lifecycleStateId | **string** | True  | Lifecycle State ID


### Return type

[**LifecyclestateDeleted**](LifecyclestateDeleted.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
202 | The request was successfully accepted into the system. | LifecyclestateDeleted
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //DeleteLifecycleState

    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238"
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8"



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




## get-lifecycle-state


This endpoint retrieves a Lifecycle State.
A token with ORG_ADMIN or API authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityProfileId | **string** | True  | Identity Profile ID
Path   | lifecycleStateId | **string** | True  | Lifecycle State ID


### Return type

[**LifecycleState**](LifecycleState.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The requested LifecycleState was successfully retrieved. | LifecycleState
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //GetLifecycleState

    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238"
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8"



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




## list-lifecycle-states


This end-point lists all the LifecycleStates associated with IdentityProfiles.
A token with API, or ORG_ADMIN authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityProfileId | **string** | True  | The IdentityProfile id
  Query | limit | **int32** |   (optional) (default to 250) | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | offset | **int32** |   (optional) (default to 0) | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | count | **bool** |   (optional) (default to false) | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information.
  Query | sorters | **string** |   (optional) | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **created, modified**


### Return type

[**[]LifecycleState**](LifecycleState.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | List of LifecycleState objects | []LifecycleState
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //ListLifecycleStates

    identityProfileId := "ef38f94347e94562b5bb8424a56397d8"
    //limit := int32(250)
    //offset := int32(0)
    //count := true
    //sorters := "created,modified"



    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.LifecycleStatesAPI.ListLifecycleStates(context.Background(), identityProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LifecycleStatesAPI.ListLifecycleStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLifecycleStates`: []LifecycleState
    fmt.Fprintf(os.Stdout, "Response from `LifecycleStatesAPI.ListLifecycleStates`: %v\n", resp)
}
```




## set-lifecycle-state


This endpoint will set/update an identity's lifecycle state to the one provided and updates the corresponding Identity Profile.
A token with ORG_ADMIN or API authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityId | **string** | True  | The ID of the identity to update
 Body  | setLifecycleStateRequest | [**SetLifecycleStateRequest**](SetLifecycleStateRequest.md) | True  | 


### Return type

[**SetLifecycleState200Response**](SetLifecycleState200Response.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The request was successfully accepted into the system. | SetLifecycleState200Response
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //SetLifecycleState

    identityId := "2c9180857893f1290178944561990364"
    setLifecycleStateRequest := *sailpoint.NewSetLifecycleStateRequest()



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




## update-lifecycle-states


This endpoint updates individual Lifecycle State fields using the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.
A token with ORG_ADMIN or API authority is required to call this API.

### Parameters 
Param Type | Name | Data Type | Required  | Description
------------- | ------------- | ------------- | ------------- | ------------- 
Path   | identityProfileId | **string** | True  | Identity Profile ID
Path   | lifecycleStateId | **string** | True  | Lifecycle State ID
 Body  | jsonPatchOperation | [**[]JsonPatchOperation**](JsonPatchOperation.md) | True  | A list of lifecycle state update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields can be updated: * enabled * description * accountActions * accessProfileIds * emailNotificationOption 


### Return type

[**LifecycleState**](LifecycleState.md)

### Responses
Code | Description  | Data Type
------------- | ------------- | -------------
200 | The LifecycleState was successfully updated. | LifecycleState
400 | Client Error - Returned if the request body is invalid. | ErrorResponseDto
401 | Unauthorized - Returned if there is no authorization header, or if the JWT token is expired. | ListAccessProfiles401Response
403 | Forbidden - Returned if the user you are running as, doesn&#39;t have access to this end-point. | ErrorResponseDto
404 | Not Found - returned if the request URL refers to a resource or object that does not exist | ErrorResponseDto
429 | Too Many Requests - Returned in response to too many requests in a given period of time - rate limited. The Retry-After header in the response includes how long to wait before trying again. | ListAccessProfiles429Response
500 | Internal Server Error - Returned if there is an unexpected error. | ErrorResponseDto


### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    v3 "github.com/sailpoint-oss/golang-sdk/v2/api_v3"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    //UpdateLifecycleStates

    identityProfileId := "2b838de9-db9b-abcf-e646-d4f274ad4238"
    lifecycleStateId := "ef38f94347e94562b5bb8424a56397d8"
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")}



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



