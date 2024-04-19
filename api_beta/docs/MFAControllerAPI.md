# \MFAControllerAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSendToken**](MFAControllerAPI.md#CreateSendToken) | **Post** /mfa/token/send | Create and send user token
[**PingVerificationStatus**](MFAControllerAPI.md#PingVerificationStatus) | **Post** /mfa/{method}/poll | Polling MFA method by VerificationPollRequest
[**SendDuoVerifyRequest**](MFAControllerAPI.md#SendDuoVerifyRequest) | **Post** /mfa/duo-web/verify | Verifying authentication via Duo method
[**SendKbaAnswers**](MFAControllerAPI.md#SendKbaAnswers) | **Post** /mfa/kba/authenticate | Authenticate KBA provided MFA method
[**SendOktaVerifyRequest**](MFAControllerAPI.md#SendOktaVerifyRequest) | **Post** /mfa/okta-verify/verify | Verifying authentication via Okta method
[**SendTokenAuthRequest**](MFAControllerAPI.md#SendTokenAuthRequest) | **Post** /mfa/token/authenticate | Authenticate Token provided MFA method



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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    sendTokenRequest := *openapiclient.NewSendTokenRequest("will.albin", "EMAIL_WORK") // SendTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.CreateSendToken(context.Background()).SendTokenRequest(sendTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.CreateSendToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSendToken`: SendTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.CreateSendToken`: %v\n", resp)
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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    method := "okta-verify" // string | The name of the MFA method. The currently supported method names are 'okta-verify', 'duo-web', 'kba','token', 'rsa'
    verificationPollRequest := *openapiclient.NewVerificationPollRequest("089899f13a8f4da7824996191587bab9") // VerificationPollRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.PingVerificationStatus(context.Background(), method).VerificationPollRequest(verificationPollRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.PingVerificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PingVerificationStatus`: VerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.PingVerificationStatus`: %v\n", resp)
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


## SendDuoVerifyRequest

> VerificationResponse SendDuoVerifyRequest(ctx).DuoVerificationRequest(duoVerificationRequest).Execute()

Verifying authentication via Duo method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    duoVerificationRequest := *openapiclient.NewDuoVerificationRequest("2c9180947f0ef465017f215cbcfd004b", "AUTH|d2lsbC5hbGJpbnxESTZNMFpHSThKQVRWTVpZN0M5VXwxNzAxMjUzMDg5|f1f5f8ced5b340f3d303b05d0efa0e43b6a8f970:APP|d2lsbC5hbGJpbnxESTZNMFpHSThKQVRWTVpZN0M5VXwxNzAxMjU2NjE5|cb44cf44353f5127edcae31b1da0355f87357db2") // DuoVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.SendDuoVerifyRequest(context.Background()).DuoVerificationRequest(duoVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.SendDuoVerifyRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendDuoVerifyRequest`: VerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.SendDuoVerifyRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDuoVerifyRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duoVerificationRequest** | [**DuoVerificationRequest**](DuoVerificationRequest.md) |  | 

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

> KbaAuthResponse SendKbaAnswers(ctx).KbaAnswerRequestItem(kbaAnswerRequestItem).Execute()

Authenticate KBA provided MFA method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    kbaAnswerRequestItem := []openapiclient.KbaAnswerRequestItem{*openapiclient.NewKbaAnswerRequestItem("c54fee53-2d63-4fc5-9259-3e93b9994135", "Your answer")} // []KbaAnswerRequestItem | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.SendKbaAnswers(context.Background()).KbaAnswerRequestItem(kbaAnswerRequestItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.SendKbaAnswers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendKbaAnswers`: KbaAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.SendKbaAnswers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendKbaAnswersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbaAnswerRequestItem** | [**[]KbaAnswerRequestItem**](KbaAnswerRequestItem.md) |  | 

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


## SendOktaVerifyRequest

> VerificationResponse SendOktaVerifyRequest(ctx).OktaVerificationRequest(oktaVerificationRequest).Execute()

Verifying authentication via Okta method



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    oktaVerificationRequest := *openapiclient.NewOktaVerificationRequest("example@mail.com") // OktaVerificationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.SendOktaVerifyRequest(context.Background()).OktaVerificationRequest(oktaVerificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.SendOktaVerifyRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendOktaVerifyRequest`: VerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.SendOktaVerifyRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendOktaVerifyRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oktaVerificationRequest** | [**OktaVerificationRequest**](OktaVerificationRequest.md) |  | 

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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    tokenAuthRequest := *openapiclient.NewTokenAuthRequest("12345", "will.albin", "EMAIL_WORK") // TokenAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAControllerAPI.SendTokenAuthRequest(context.Background()).TokenAuthRequest(tokenAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAControllerAPI.SendTokenAuthRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendTokenAuthRequest`: TokenAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAControllerAPI.SendTokenAuthRequest`: %v\n", resp)
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

