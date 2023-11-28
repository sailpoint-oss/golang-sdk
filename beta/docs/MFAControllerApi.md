# \MFAControllerApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSendToken**](MFAControllerApi.md#CreateSendToken) | **Post** /mfa/token/send | Create and send user token
[**PingVerificationStatus**](MFAControllerApi.md#PingVerificationStatus) | **Post** /mfa/{method}/poll | Polling MFA method by VerificationPollRequest
[**SendKbaAnswers**](MFAControllerApi.md#SendKbaAnswers) | **Post** /mfa/kba/authenticate | Authenticate KBA provided MFA method
[**SendTokenAuthRequest**](MFAControllerApi.md#SendTokenAuthRequest) | **Post** /mfa/token/authenticate | Authenticate Token provided MFA method



## CreateSendToken

> SendTokenResponse CreateSendToken(ctx).SendTokenRequest(sendTokenRequest).Execute()

Create and send user token



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
    sendTokenRequest := *openapiclient.NewSendTokenRequest("will.albin", "EMAIL_WORK") // SendTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerApi.CreateSendToken(context.Background()).SendTokenRequest(sendTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerApi.CreateSendToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSendToken`: SendTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerApi.CreateSendToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSendTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendTokenRequest** | [**SendTokenRequest**](SendTokenRequest.md) |  | 

### Return type

[**SendTokenResponse**](SendTokenResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingVerificationStatus

> VerificationResponse PingVerificationStatus(ctx, method).VerificationPollRequest(verificationPollRequest).Execute()

Polling MFA method by VerificationPollRequest



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
    method := "okta-verify" // string | The name of the MFA method. The currently supported method names are 'okta-verify', 'duo-web', 'kba','token', 'rsa'
    verificationPollRequest := *openapiclient.NewVerificationPollRequest("089899f13a8f4da7824996191587bab9") // VerificationPollRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerApi.PingVerificationStatus(context.Background(), method).VerificationPollRequest(verificationPollRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerApi.PingVerificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingVerificationStatus`: VerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerApi.PingVerificationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method names are &#39;okta-verify&#39;, &#39;duo-web&#39;, &#39;kba&#39;,&#39;token&#39;, &#39;rsa&#39; | 

### Other Parameters

Other parameters are passed through a pointer to a apiPingVerificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verificationPollRequest** | [**VerificationPollRequest**](VerificationPollRequest.md) |  | 

### Return type

[**VerificationResponse**](VerificationResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendKbaAnswers

> KbaAuthResponse SendKbaAnswers(ctx).KbaAnswerRequest(kbaAnswerRequest).Execute()

Authenticate KBA provided MFA method



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
    kbaAnswerRequest := *openapiclient.NewKbaAnswerRequest([]openapiclient.KbaAnswerRequestItem{*openapiclient.NewKbaAnswerRequestItem("089899f13a8f4da7824996191587bab9", "Your answer")}) // KbaAnswerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerApi.SendKbaAnswers(context.Background()).KbaAnswerRequest(kbaAnswerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerApi.SendKbaAnswers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendKbaAnswers`: KbaAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerApi.SendKbaAnswers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendKbaAnswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbaAnswerRequest** | [**KbaAnswerRequest**](KbaAnswerRequest.md) |  | 

### Return type

[**KbaAuthResponse**](KbaAuthResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTokenAuthRequest

> TokenAuthResponse SendTokenAuthRequest(ctx).TokenAuthRequest(tokenAuthRequest).Execute()

Authenticate Token provided MFA method



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
    tokenAuthRequest := *openapiclient.NewTokenAuthRequest("12345", "will.albin", "EMAIL_WORK") // TokenAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerApi.SendTokenAuthRequest(context.Background()).TokenAuthRequest(tokenAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerApi.SendTokenAuthRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTokenAuthRequest`: TokenAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerApi.SendTokenAuthRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendTokenAuthRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAuthRequest** | [**TokenAuthRequest**](TokenAuthRequest.md) |  | 

### Return type

[**TokenAuthResponse**](TokenAuthResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

