# \PersonalAccessTokensAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalAccessToken**](PersonalAccessTokensAPI.md#CreatePersonalAccessToken) | **Post** /personal-access-tokens | Create Personal Access Token
[**DeletePersonalAccessToken**](PersonalAccessTokensAPI.md#DeletePersonalAccessToken) | **Delete** /personal-access-tokens/{id} | Delete Personal Access Token
[**ListPersonalAccessTokens**](PersonalAccessTokensAPI.md#ListPersonalAccessTokens) | **Get** /personal-access-tokens | List Personal Access Tokens
[**PatchPersonalAccessToken**](PersonalAccessTokensAPI.md#PatchPersonalAccessToken) | **Patch** /personal-access-tokens/{id} | Patch Personal Access Token



## Create Personal Access Token

> CreatePersonalAccessTokenResponse CreatePersonalAccessToken(ctx).CreatePersonalAccessTokenRequest(createPersonalAccessTokenRequest).Execute()

This creates a personal access token.

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
    createPersonalAccessTokenRequest := *sailpoint.NewCreatePersonalAccessTokenRequest("NodeJS Integration") // CreatePersonalAccessTokenRequest | Name and scope of personal access token.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PersonalAccessTokensAPI.CreatePersonalAccessToken(context.Background()).CreatePersonalAccessTokenRequest(createPersonalAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensAPI.CreatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersonalAccessToken`: CreatePersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensAPI.CreatePersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPersonalAccessTokenRequest** | [**CreatePersonalAccessTokenRequest**](CreatePersonalAccessTokenRequest.md) | Name and scope of personal access token. | 

### Return type

[**CreatePersonalAccessTokenResponse**](CreatePersonalAccessTokenResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete Personal Access Token

> DeletePersonalAccessToken(ctx, id).Execute()

This deletes a personal access token.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The personal access token id

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.PersonalAccessTokensAPI.DeletePersonalAccessToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensAPI.DeletePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The personal access token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePersonalAccessTokenRequest struct via the builder pattern


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


## List Personal Access Tokens

> []GetPersonalAccessTokenResponse ListPersonalAccessTokens(ctx).OwnerId(ownerId).Filters(filters).Execute()

This gets a collection of personal access tokens associated with the optional `owner-id`.  query parameter. If the `owner-id` query parameter is omitted, all personal access tokens  for a tenant will be retrieved, but the caller must have the 'idn:all-personal-access-tokens:read' right.

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
    ownerId := "2c9180867b50d088017b554662fb281e" // string | The identity ID of the owner whose personal access tokens should be listed.  If \"me\", the caller should have the following right: 'idn:my-personal-access-tokens:read' If an actual owner ID or if the `owner-id` parameter is omitted in the request,  the caller should have the following right: 'idn:all-personal-access-tokens:read'.  If the caller has the following right, then managed personal access tokens associated with `owner-id`  will be retrieved: 'idn:managed-personal-access-tokens:read' (optional)
    filters := "lastUsed le 2023-02-05T10:59:27.214Z" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **lastUsed**: *le, isnull* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PersonalAccessTokensAPI.ListPersonalAccessTokens(context.Background()).OwnerId(ownerId).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensAPI.ListPersonalAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersonalAccessTokens`: []GetPersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensAPI.ListPersonalAccessTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalAccessTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | The identity ID of the owner whose personal access tokens should be listed.  If \&quot;me\&quot;, the caller should have the following right: &#39;idn:my-personal-access-tokens:read&#39; If an actual owner ID or if the &#x60;owner-id&#x60; parameter is omitted in the request,  the caller should have the following right: &#39;idn:all-personal-access-tokens:read&#39;.  If the caller has the following right, then managed personal access tokens associated with &#x60;owner-id&#x60;  will be retrieved: &#39;idn:managed-personal-access-tokens:read&#39; | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **lastUsed**: *le, isnull* | 

### Return type

[**[]GetPersonalAccessTokenResponse**](GetPersonalAccessTokenResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Patch Personal Access Token

> GetPersonalAccessTokenResponse PatchPersonalAccessToken(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

This performs a targeted update to the field(s) of a Personal Access Token.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The Personal Access Token id
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * scope 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.PersonalAccessTokensAPI.PatchPersonalAccessToken(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensAPI.PatchPersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPersonalAccessToken`: GetPersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensAPI.PatchPersonalAccessToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Personal Access Token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPersonalAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * scope  | 

### Return type

[**GetPersonalAccessTokenResponse**](GetPersonalAccessTokenResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

