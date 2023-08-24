# \PersonalAccessTokensApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalAccessToken**](PersonalAccessTokensApi.md#CreatePersonalAccessToken) | **Post** /personal-access-tokens | Create Personal Access Token
[**DeletePersonalAccessToken**](PersonalAccessTokensApi.md#DeletePersonalAccessToken) | **Delete** /personal-access-tokens/{id} | Delete Personal Access Token
[**ListPersonalAccessTokens**](PersonalAccessTokensApi.md#ListPersonalAccessTokens) | **Get** /personal-access-tokens | List Personal Access Tokens
[**PatchPersonalAccessToken**](PersonalAccessTokensApi.md#PatchPersonalAccessToken) | **Patch** /personal-access-tokens/{id} | Patch Personal Access Token



## CreatePersonalAccessToken

> CreatePersonalAccessTokenResponse CreatePersonalAccessToken(ctx).CreatePersonalAccessTokenRequest(createPersonalAccessTokenRequest).Execute()

Create Personal Access Token



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
    createPersonalAccessTokenRequest := *openapiclient.NewCreatePersonalAccessTokenRequest("NodeJS Integration") // CreatePersonalAccessTokenRequest | Name and scope of personal access token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.CreatePersonalAccessToken(context.Background()).CreatePersonalAccessTokenRequest(createPersonalAccessTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.CreatePersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePersonalAccessToken`: CreatePersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.CreatePersonalAccessToken`: %v\n", resp)
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


## DeletePersonalAccessToken

> DeletePersonalAccessToken(ctx, id).Execute()

Delete Personal Access Token



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The personal access token id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.DeletePersonalAccessToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.DeletePersonalAccessToken``: %v\n", err)
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


## ListPersonalAccessTokens

> []GetPersonalAccessTokenResponse ListPersonalAccessTokens(ctx).OwnerId(ownerId).Filters(filters).Execute()

List Personal Access Tokens



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
    ownerId := "2c9180867b50d088017b554662fb281e" // string | The identity ID of the owner whose personal access tokens should be listed.  If \"me\", the caller should have the following right: 'idn:my-personal-access-tokens:read' If an actual owner ID or if the `owner-id` parameter is omitted in the request,  the caller should have the following right: 'idn:all-personal-access-tokens:read'.  If the caller has the following right, then managed personal access tokens associated with `owner-id`  will be retrieved: 'idn:managed-personal-access-tokens:read' (optional)
    filters := "lastUsed le 2023-02-05T10:59:27.214Z" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **lastUsed**: *le, isnull* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.ListPersonalAccessTokens(context.Background()).OwnerId(ownerId).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.ListPersonalAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersonalAccessTokens`: []GetPersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.ListPersonalAccessTokens`: %v\n", resp)
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


## PatchPersonalAccessToken

> GetPersonalAccessTokenResponse PatchPersonalAccessToken(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Patch Personal Access Token



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The Personal Access Token id
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of OAuth client update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * scope 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.PatchPersonalAccessToken(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PersonalAccessTokensApi.PatchPersonalAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchPersonalAccessToken`: GetPersonalAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `PersonalAccessTokensApi.PatchPersonalAccessToken`: %v\n", resp)
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

