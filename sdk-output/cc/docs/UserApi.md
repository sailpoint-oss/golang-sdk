# \UserApi

All URIs are relative to *https://sailpoint.api.identitynow.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateUserPermissions**](UserApi.md#UpdateUserPermissions) | **Post** /cc/api/user/updatePermissions | Update User Permissions



## UpdateUserPermissions

> UpdateUserPermissions(ctx).UpdateUserPermissionsRequest(updateUserPermissionsRequest).Execute()

Update User Permissions

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
    updateUserPermissionsRequest := *openapiclient.NewUpdateUserPermissionsRequest() // UpdateUserPermissionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UpdateUserPermissions(context.Background()).UpdateUserPermissionsRequest(updateUserPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUserPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserPermissionsRequest** | [**UpdateUserPermissionsRequest**](UpdateUserPermissionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

