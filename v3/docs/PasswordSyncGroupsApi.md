# \PasswordSyncGroupsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordSyncGroup**](PasswordSyncGroupsApi.md#CreatePasswordSyncGroup) | **Post** /password-sync-groups | Create Password Sync Group



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

