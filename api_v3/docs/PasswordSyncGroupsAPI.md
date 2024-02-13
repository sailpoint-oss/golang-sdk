# \PasswordSyncGroupsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordSyncGroup**](PasswordSyncGroupsAPI.md#CreatePasswordSyncGroup) | **Post** /password-sync-groups | Create Password Sync Group
[**DeletePasswordSyncGroup**](PasswordSyncGroupsAPI.md#DeletePasswordSyncGroup) | **Delete** /password-sync-groups/{id} | Delete Password Sync Group by ID
[**GetPasswordSyncGroup**](PasswordSyncGroupsAPI.md#GetPasswordSyncGroup) | **Get** /password-sync-groups/{id} | Get Password Sync Group by ID
[**GetPasswordSyncGroups**](PasswordSyncGroupsAPI.md#GetPasswordSyncGroups) | **Get** /password-sync-groups | Get Password Sync Group List
[**UpdatePasswordSyncGroup**](PasswordSyncGroupsAPI.md#UpdatePasswordSyncGroup) | **Put** /password-sync-groups/{id} | Update Password Sync Group by ID



## Create Password Sync Group

> PasswordSyncGroup CreatePasswordSyncGroup(ctx).PasswordSyncGroup(passwordSyncGroup).Execute()

This API creates a password sync group based on the specifications provided. A token with ORG_ADMIN authority is required to call this API.

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
    passwordSyncGroup := *sailpoint.NewPasswordSyncGroup() // PasswordSyncGroup | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordSyncGroupsAPI.CreatePasswordSyncGroup(context.Background()).PasswordSyncGroup(passwordSyncGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsAPI.CreatePasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsAPI.CreatePasswordSyncGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordSyncGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordSyncGroup** | [**PasswordSyncGroup**](PasswordSyncGroup.md) |  | 

### Return type

[**PasswordSyncGroup**](PasswordSyncGroup.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Password Sync Group by ID

> DeletePasswordSyncGroup(ctx, id).Execute()

This API deletes the specified password sync group. A token with ORG_ADMIN authority is required to call this API.

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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to delete.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.PasswordSyncGroupsAPI.DeletePasswordSyncGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsAPI.DeletePasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password sync group to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordSyncGroupRequest struct via the builder pattern


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


## Get Password Sync Group by ID

> PasswordSyncGroup GetPasswordSyncGroup(ctx, id).Execute()

This API returns the sync group for the specified ID. A token with ORG_ADMIN authority is required to call this API.

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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to retrieve.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordSyncGroupsAPI.GetPasswordSyncGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsAPI.GetPasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsAPI.GetPasswordSyncGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password sync group to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordSyncGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PasswordSyncGroup**](PasswordSyncGroup.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Password Sync Group List

> []PasswordSyncGroup GetPasswordSyncGroups(ctx).Limit(limit).Offset(offset).Count(count).Execute()

This API returns a list of password sync groups. A token with ORG_ADMIN authority is required to call this API.

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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordSyncGroupsAPI.GetPasswordSyncGroups(context.Background()).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsAPI.GetPasswordSyncGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordSyncGroups`: []PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsAPI.GetPasswordSyncGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordSyncGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]

### Return type

[**[]PasswordSyncGroup**](PasswordSyncGroup.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update Password Sync Group by ID

> PasswordSyncGroup UpdatePasswordSyncGroup(ctx, id).PasswordSyncGroup(passwordSyncGroup).Execute()

This API updates the specified password sync group. A token with ORG_ADMIN authority is required to call this API.

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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to update.
    passwordSyncGroup := *sailpoint.NewPasswordSyncGroup() // PasswordSyncGroup | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PasswordSyncGroupsAPI.UpdatePasswordSyncGroup(context.Background(), id).PasswordSyncGroup(passwordSyncGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsAPI.UpdatePasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsAPI.UpdatePasswordSyncGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of password sync group to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordSyncGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordSyncGroup** | [**PasswordSyncGroup**](PasswordSyncGroup.md) |  | 

### Return type

[**PasswordSyncGroup**](PasswordSyncGroup.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

