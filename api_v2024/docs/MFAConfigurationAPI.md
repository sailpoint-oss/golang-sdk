# \MFAConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v2024*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMFADuoConfig**](MFAConfigurationAPI.md#GetMFADuoConfig) | **Get** /mfa/duo-web/config | Configuration of Duo MFA method
[**GetMFAKbaConfig**](MFAConfigurationAPI.md#GetMFAKbaConfig) | **Get** /mfa/kba/config | Configuration of KBA MFA method
[**GetMFAOktaConfig**](MFAConfigurationAPI.md#GetMFAOktaConfig) | **Get** /mfa/okta-verify/config | Configuration of Okta MFA method
[**SetMFADuoConfig**](MFAConfigurationAPI.md#SetMFADuoConfig) | **Put** /mfa/duo-web/config | Set Duo MFA configuration
[**SetMFAKBAConfig**](MFAConfigurationAPI.md#SetMFAKBAConfig) | **Post** /mfa/kba/config/answers | Set MFA KBA configuration
[**SetMFAOktaConfig**](MFAConfigurationAPI.md#SetMFAOktaConfig) | **Put** /mfa/okta-verify/config | Set Okta MFA configuration
[**TestMFAConfig**](MFAConfigurationAPI.md#TestMFAConfig) | **Get** /mfa/{method}/test | MFA method&#39;s test configuration



## GetMFADuoConfig

> MfaDuoConfig GetMFADuoConfig(ctx).Execute()

Configuration of Duo MFA method



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.GetMFADuoConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.GetMFADuoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFADuoConfig`: MfaDuoConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.GetMFADuoConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFADuoConfigRequest struct via the builder pattern


### Return type

[**MfaDuoConfig**](MfaDuoConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAKbaConfig

> []KbaQuestion GetMFAKbaConfig(ctx).AllLanguages(allLanguages).Execute()

Configuration of KBA MFA method



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
    allLanguages := false // bool | Indicator whether the question text should be returned in all configured languages    * If true, the question text is returned in all languages that it is configured in.    * If false, the question text is returned in the user locale if available, else for the default locale.     * If not passed, it behaves the same way as passing this parameter as false (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.GetMFAKbaConfig(context.Background()).AllLanguages(allLanguages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.GetMFAKbaConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAKbaConfig`: []KbaQuestion
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.GetMFAKbaConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAKbaConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allLanguages** | **bool** | Indicator whether the question text should be returned in all configured languages    * If true, the question text is returned in all languages that it is configured in.    * If false, the question text is returned in the user locale if available, else for the default locale.     * If not passed, it behaves the same way as passing this parameter as false | 

### Return type

[**[]KbaQuestion**](KbaQuestion.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAOktaConfig

> MfaOktaConfig GetMFAOktaConfig(ctx).Execute()

Configuration of Okta MFA method



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.GetMFAOktaConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.GetMFAOktaConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAOktaConfig`: MfaOktaConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.GetMFAOktaConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAOktaConfigRequest struct via the builder pattern


### Return type

[**MfaOktaConfig**](MfaOktaConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMFADuoConfig

> MfaDuoConfig SetMFADuoConfig(ctx).MfaDuoConfig(mfaDuoConfig).Execute()

Set Duo MFA configuration



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
    mfaDuoConfig := *openapiclient.NewMfaDuoConfig() // MfaDuoConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.SetMFADuoConfig(context.Background()).MfaDuoConfig(mfaDuoConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.SetMFADuoConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMFADuoConfig`: MfaDuoConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.SetMFADuoConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMFADuoConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaDuoConfig** | [**MfaDuoConfig**](MfaDuoConfig.md) |  | 

### Return type

[**MfaDuoConfig**](MfaDuoConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMFAKBAConfig

> []KbaAnswerResponseItem SetMFAKBAConfig(ctx).KbaAnswerRequestItem(kbaAnswerRequestItem).Execute()

Set MFA KBA configuration



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
    resp, r, err := apiClient.MFAConfigurationAPI.SetMFAKBAConfig(context.Background()).KbaAnswerRequestItem(kbaAnswerRequestItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.SetMFAKBAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMFAKBAConfig`: []KbaAnswerResponseItem
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.SetMFAKBAConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMFAKBAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbaAnswerRequestItem** | [**[]KbaAnswerRequestItem**](KbaAnswerRequestItem.md) |  | 

### Return type

[**[]KbaAnswerResponseItem**](KbaAnswerResponseItem.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMFAOktaConfig

> MfaOktaConfig SetMFAOktaConfig(ctx).MfaOktaConfig(mfaOktaConfig).Execute()

Set Okta MFA configuration



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
    mfaOktaConfig := *openapiclient.NewMfaOktaConfig() // MfaOktaConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.SetMFAOktaConfig(context.Background()).MfaOktaConfig(mfaOktaConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.SetMFAOktaConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMFAOktaConfig`: MfaOktaConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.SetMFAOktaConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetMFAOktaConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfaOktaConfig** | [**MfaOktaConfig**](MfaOktaConfig.md) |  | 

### Return type

[**MfaOktaConfig**](MfaOktaConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestMFAConfig

> MfaConfigTestResponse TestMFAConfig(ctx, method).Execute()

MFA method's test configuration



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
    method := "okta-verify" // string | The name of the MFA method. The currently supported method names are 'okta-verify' and 'duo-web'.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.TestMFAConfig(context.Background(), method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.TestMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMFAConfig`: MfaConfigTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.TestMFAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method names are &#39;okta-verify&#39; and &#39;duo-web&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MfaConfigTestResponse**](MfaConfigTestResponse.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

