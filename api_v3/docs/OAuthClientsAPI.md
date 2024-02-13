# \OAuthClientsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOauthClient**](OAuthClientsAPI.md#CreateOauthClient) | **Post** /oauth-clients | Create OAuth Client
[**DeleteOauthClient**](OAuthClientsAPI.md#DeleteOauthClient) | **Delete** /oauth-clients/{id} | Delete OAuth Client
[**GetOauthClient**](OAuthClientsAPI.md#GetOauthClient) | **Get** /oauth-clients/{id} | Get OAuth Client
[**ListOauthClients**](OAuthClientsAPI.md#ListOauthClients) | **Get** /oauth-clients | List OAuth Clients
[**PatchOauthClient**](OAuthClientsAPI.md#PatchOauthClient) | **Patch** /oauth-clients/{id} | Patch OAuth Client



## Create OAuth Client

> CreateOAuthClientResponse CreateOauthClient(ctx).CreateOAuthClientRequest(createOAuthClientRequest).Execute()

This creates an OAuth client.

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
    createOAuthClientRequest := *sailpoint.NewCreateOAuthClientRequest("Demo API Client", "An API client used for the authorization_code, refresh_token, and client_credentials flows", int32(750), []sailpoint.GrantType{sailpoint.GrantType("CLIENT_CREDENTIALS")}, sailpoint.AccessType("ONLINE"), true) // CreateOAuthClientRequest | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.OAuthClientsAPI.CreateOauthClient(context.Background()).CreateOAuthClientRequest(createOAuthClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsAPI.CreateOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOauthClient`: CreateOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsAPI.CreateOauthClient`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete OAuth Client

> DeleteOauthClient(ctx, id).Execute()

This deletes an OAuth client.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.OAuthClientsAPI.DeleteOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsAPI.DeleteOauthClient``: %v\n", err)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get OAuth Client

> GetOAuthClientResponse GetOauthClient(ctx, id).Execute()

This gets details of an OAuth client.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.OAuthClientsAPI.GetOauthClient(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsAPI.GetOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOauthClient`: GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsAPI.GetOauthClient`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List OAuth Clients

> []GetOAuthClientResponse ListOauthClients(ctx).Filters(filters).Execute()

This gets a list of OAuth clients.

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
    filters := "lastUsed le 2023-02-05T10:59:27.214Z" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **lastUsed**: *le, isnull* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.OAuthClientsAPI.ListOauthClients(context.Background()).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsAPI.ListOauthClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOauthClients`: []GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsAPI.ListOauthClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOauthClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **lastUsed**: *le, isnull* | 

### Return type

[**[]GetOAuthClientResponse**](GetOAuthClientResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch OAuth Client

> GetOAuthClientResponse PatchOauthClient(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

This performs a targeted update to the field(s) of an OAuth client.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The OAuth client id
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * tenant * businessName * homepageUrl * name * description * accessTokenValiditySeconds * refreshTokenValiditySeconds * redirectUris * grantTypes * accessType * enabled * strongAuthSupported * claimsSupported 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.OAuthClientsAPI.PatchOauthClient(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OAuthClientsAPI.PatchOauthClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchOauthClient`: GetOAuthClientResponse
    fmt.Fprintf(os.Stdout, "Response from `OAuthClientsAPI.PatchOauthClient`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

