# \OAuthClientsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthClient**](OAuthClientsApi.md#CreateOauthClient) | **Post** /oauth-clients | Create OAuth Client
[**DeleteOauthClient**](OAuthClientsApi.md#DeleteOauthClient) | **Delete** /oauth-clients/{id} | Delete OAuth Client
[**GetOauthClient**](OAuthClientsApi.md#GetOauthClient) | **Get** /oauth-clients/{id} | Get OAuth Client
[**ListOauthClients**](OAuthClientsApi.md#ListOauthClients) | **Get** /oauth-clients | List OAuth Clients
[**PatchOauthClient**](OAuthClientsApi.md#PatchOauthClient) | **Patch** /oauth-clients/{id} | Patch OAuth Client



## CreateOauthClient

> CreateOAuthClientResponse CreateOauthClient(ctx).CreateOAuthClientRequest(createOAuthClientRequest).Execute()

Create OAuth Client



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
    createOAuthClientRequest := *openapiclient.NewCreateOAuthClientRequest("Demo API Client", "An API client used for the authorization_code, refresh_token, and client_credentials flows", int32(750), []openapiclient.GrantType{openapiclient.GrantType("CLIENT_CREDENTIALS")}, openapiclient.AccessType("ONLINE"), true) // CreateOAuthClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthClientsApi.CreateOauthClient(context.Background()).CreateOAuthClientRequest(createOAuthClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsApi.CreateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOauthClient`: CreateOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsApi.CreateOauthClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOAuthClientRequest** | [**CreateOAuthClientRequest**](CreateOAuthClientRequest.md) |  | 

### Return type

[**CreateOAuthClientResponse**](CreateOAuthClientResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOauthClient

> DeleteOauthClient(ctx, id).Execute()

Delete OAuth Client



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthClientsApi.DeleteOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsApi.DeleteOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The OAuth client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOauthClientRequest struct via the builder pattern


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


## GetOauthClient

> GetOAuthClientResponse GetOauthClient(ctx, id).Execute()

Get OAuth Client



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthClientsApi.GetOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsApi.GetOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClient`: GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsApi.GetOauthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The OAuth client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOAuthClientResponse**](GetOAuthClientResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOauthClients

> []GetOAuthClientResponse ListOauthClients(ctx).Execute()

List OAuth Clients



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthClientsApi.ListOauthClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsApi.ListOauthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOauthClients`: []GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsApi.ListOauthClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOauthClientsRequest struct via the builder pattern


### Return type

[**[]GetOAuthClientResponse**](GetOAuthClientResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchOauthClient

> GetOAuthClientResponse PatchOauthClient(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch OAuth Client



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * tenant * businessName * homepageUrl * name * description * accessTokenValiditySeconds * refreshTokenValiditySeconds * redirectUris * grantTypes * accessType * enabled * strongAuthSupported * claimsSupported 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OAuthClientsApi.PatchOauthClient(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsApi.PatchOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchOauthClient`: GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsApi.PatchOauthClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The OAuth client id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchOauthClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * tenant * businessName * homepageUrl * name * description * accessTokenValiditySeconds * refreshTokenValiditySeconds * redirectUris * grantTypes * accessType * enabled * strongAuthSupported * claimsSupported  | 

### Return type

[**GetOAuthClientResponse**](GetOAuthClientResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

