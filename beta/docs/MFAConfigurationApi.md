# \MFAConfigurationApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMFAConfig**](MFAConfigurationApi.md#GetMFAConfig) | **Get** /mfa/{method}/config | Configuration of a MFA method
[**SetMFAConfig**](MFAConfigurationApi.md#SetMFAConfig) | **Put** /mfa/{method}/config | Set MFA method configuration
[**TestMFAConfig**](MFAConfigurationApi.md#TestMFAConfig) | **Get** /mfa/{method}/test | MFA method&#39;s test configuration



## GetMFAConfig

> MfaConfig GetMFAConfig(ctx, method).Execute()

Configuration of a MFA method



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
    method := "okta-verify" // string | The name of the MFA method. The currently supported method name is okta-verify.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationApi.GetMFAConfig(context.Background(), method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationApi.GetMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAConfig`: MfaConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationApi.GetMFAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method name is okta-verify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MfaConfig**](MfaConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMFAConfig

> MfaConfig SetMFAConfig(ctx, method).MfaConfig(mfaConfig).Execute()

Set MFA method configuration



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
    method := "okta-verify" // string | The name of the MFA method. The currently supported method name is okta-verify.
    mfaConfig := *openapiclient.NewMfaConfig() // MfaConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationApi.SetMFAConfig(context.Background(), method).MfaConfig(mfaConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationApi.SetMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetMFAConfig`: MfaConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationApi.SetMFAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method name is okta-verify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfaConfig** | [**MfaConfig**](MfaConfig.md) |  | 

### Return type

[**MfaConfig**](MfaConfig.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

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
    openapiclient "./openapi"
)

func main() {
    method := "okta-verify" // string | The name of the MFA method. The currently supported method name is okta-verify.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationApi.TestMFAConfig(context.Background(), method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationApi.TestMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestMFAConfig`: MfaConfigTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationApi.TestMFAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method name is okta-verify. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MfaConfigTestResponse**](MfaConfigTestResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

