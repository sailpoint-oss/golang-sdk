# \AuthUserApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthUser**](AuthUserApi.md#GetAuthUser) | **Get** /auth-users/{id} | Auth User Details
[**PatchAuthUser**](AuthUserApi.md#PatchAuthUser) | **Patch** /auth-users/{id} | Auth User Update



## GetAuthUser

> AuthUser GetAuthUser(ctx, id).Execute()

Auth User Details



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthUserApi.GetAuthUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthUserApi.GetAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthUserApi.GetAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthUser**](AuthUser.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAuthUser

> AuthUser PatchAuthUser(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Auth User Update



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Identity ID
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of auth user update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  PATCH can only be applied to the following fields:   *   \"capabilities\"  A list of valid capabilities can be found using the GET ams/v3/authorization/authorization-capabilities/ endpoint. Capabilities can only be patched if they are administrator assignable, as indicated by the 'adminAssignable' field from the output of list authorization-capabilities. Capabilities that have a legacy group ('legacyGroup' field) need to be patched using the legacyGroup name (e.g. 'ORG_ADMIN'). Capabilities that are adminAssignable but do not have a legacyGroup can be patched using the ams id (e.g. 'cam:new-role').  A 400.1.1 Illegal update attempt detail code indicates that you attempted to PATCH a field that is not allowed.  Requires security scope of 'sp:auth-user:update' 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthUserApi.PatchAuthUser(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthUserApi.PatchAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchAuthUser`: AuthUser
    fmt.Fprintf(os.Stdout, "Response from `AuthUserApi.PatchAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identity ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of auth user update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  PATCH can only be applied to the following fields:   *   \&quot;capabilities\&quot;  A list of valid capabilities can be found using the GET ams/v3/authorization/authorization-capabilities/ endpoint. Capabilities can only be patched if they are administrator assignable, as indicated by the &#39;adminAssignable&#39; field from the output of list authorization-capabilities. Capabilities that have a legacy group (&#39;legacyGroup&#39; field) need to be patched using the legacyGroup name (e.g. &#39;ORG_ADMIN&#39;). Capabilities that are adminAssignable but do not have a legacyGroup can be patched using the ams id (e.g. &#39;cam:new-role&#39;).  A 400.1.1 Illegal update attempt detail code indicates that you attempted to PATCH a field that is not allowed.  Requires security scope of &#39;sp:auth-user:update&#39;  | 

### Return type

[**AuthUser**](AuthUser.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

