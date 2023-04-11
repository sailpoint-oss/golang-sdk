# \PasswordSyncGroupsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordSyncGroup**](PasswordSyncGroupsApi.md#CreatePasswordSyncGroup) | **Post** /password-sync-groups | Create Password Sync Group
[**DeletePasswordSyncGroup**](PasswordSyncGroupsApi.md#DeletePasswordSyncGroup) | **Delete** /password-sync-groups/{id} | Delete Password Sync Group by ID
[**GetPasswordSyncGroup**](PasswordSyncGroupsApi.md#GetPasswordSyncGroup) | **Get** /password-sync-groups/{id} | Get Password Sync Group by ID
[**GetPasswordSyncGroups**](PasswordSyncGroupsApi.md#GetPasswordSyncGroups) | **Get** /password-sync-groups | Get Password Sync Group List
[**UpdatePasswordSyncGroup**](PasswordSyncGroupsApi.md#UpdatePasswordSyncGroup) | **Put** /password-sync-groups/{id} | Update Password Sync Group by ID



## CreatePasswordSyncGroup

> PasswordSyncGroup CreatePasswordSyncGroup(ctx).PasswordSyncGroup(passwordSyncGroup).Execute()

Create Password Sync Group



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
    passwordSyncGroup := *openapiclient.NewPasswordSyncGroup() // PasswordSyncGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordSyncGroupsApi.CreatePasswordSyncGroup(context.Background()).PasswordSyncGroup(passwordSyncGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsApi.CreatePasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsApi.CreatePasswordSyncGroup`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordSyncGroup

> DeletePasswordSyncGroup(ctx, id).Execute()

Delete Password Sync Group by ID



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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordSyncGroupsApi.DeletePasswordSyncGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsApi.DeletePasswordSyncGroup``: %v\n", err)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordSyncGroup

> PasswordSyncGroup GetPasswordSyncGroup(ctx, id).Execute()

Get Password Sync Group by ID



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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordSyncGroupsApi.GetPasswordSyncGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsApi.GetPasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsApi.GetPasswordSyncGroup`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordSyncGroups

> []PasswordSyncGroup GetPasswordSyncGroups(ctx).Limit(limit).Offset(offset).Count(count).Execute()

Get Password Sync Group List



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
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordSyncGroupsApi.GetPasswordSyncGroups(context.Background()).Limit(limit).Offset(offset).Count(count).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsApi.GetPasswordSyncGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordSyncGroups`: []PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsApi.GetPasswordSyncGroups`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordSyncGroup

> PasswordSyncGroup UpdatePasswordSyncGroup(ctx, id).PasswordSyncGroup(passwordSyncGroup).Execute()

Update Password Sync Group by ID



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
    id := "6881f631-3bd5-4213-9c75-8e05cc3e35dd" // string | The ID of password sync group to update.
    passwordSyncGroup := *openapiclient.NewPasswordSyncGroup() // PasswordSyncGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordSyncGroupsApi.UpdatePasswordSyncGroup(context.Background(), id).PasswordSyncGroup(passwordSyncGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordSyncGroupsApi.UpdatePasswordSyncGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordSyncGroup`: PasswordSyncGroup
    fmt.Fprintf(os.Stdout, "Response from `PasswordSyncGroupsApi.UpdatePasswordSyncGroup`: %v\n", resp)
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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

