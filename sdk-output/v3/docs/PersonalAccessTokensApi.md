# \PersonalAccessTokensApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePersonalAccessToken**](PersonalAccessTokensApi.md#CreatePersonalAccessToken) | **Post** /personal-access-tokens | Create Personal Access Token
[**DeletePersonalAccessToken**](PersonalAccessTokensApi.md#DeletePersonalAccessToken) | **Delete** /personal-access-tokens/{id} | Delete Personal Access Token
[**ListPersonalAccessTokens**](PersonalAccessTokensApi.md#ListPersonalAccessTokens) | **Get** /personal-access-tokens | List Personal Access Tokens



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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersonalAccessTokens

> []GetPersonalAccessTokenResponse ListPersonalAccessTokens(ctx).OwnerId(ownerId).Execute()

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
    ownerId := "2c9180867b50d088017b554662fb281e" // string | The identity ID of the owner whose personal access tokens should be listed. </br> If \"me\", the caller should have the following right: 'idn:my-personal-access-tokens:read'</br> </br> If an actual owner ID or if the <code>owner-id</code> parameter is omitted in the request, </br> the caller should have the following right: 'idn:all-personal-access-tokens:read'. </br> </br> If the caller has the following right, then managed personal access tokens associated with <code>owner-id</code> </br> will be retrieved: 'idn:managed-personal-access-tokens:read' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PersonalAccessTokensApi.ListPersonalAccessTokens(context.Background()).OwnerId(ownerId).Execute()
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
 **ownerId** | **string** | The identity ID of the owner whose personal access tokens should be listed. &lt;/br&gt; If \&quot;me\&quot;, the caller should have the following right: &#39;idn:my-personal-access-tokens:read&#39;&lt;/br&gt; &lt;/br&gt; If an actual owner ID or if the &lt;code&gt;owner-id&lt;/code&gt; parameter is omitted in the request, &lt;/br&gt; the caller should have the following right: &#39;idn:all-personal-access-tokens:read&#39;. &lt;/br&gt; &lt;/br&gt; If the caller has the following right, then managed personal access tokens associated with &lt;code&gt;owner-id&lt;/code&gt; &lt;/br&gt; will be retrieved: &#39;idn:managed-personal-access-tokens:read&#39; | 

### Return type

[**[]GetPersonalAccessTokenResponse**](GetPersonalAccessTokenResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

