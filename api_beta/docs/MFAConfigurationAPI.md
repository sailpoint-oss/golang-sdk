# \MFAConfigurationAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMFAConfig**](MFAConfigurationAPI.md#DeleteMFAConfig) | **Delete** /mfa/{method}/delete | Delete MFA method configuration
[**GetMFADuoConfig**](MFAConfigurationAPI.md#GetMFADuoConfig) | **Get** /mfa/duo-web/config | Configuration of Duo MFA method
[**GetMFAOktaConfig**](MFAConfigurationAPI.md#GetMFAOktaConfig) | **Get** /mfa/okta-verify/config | Configuration of Okta MFA method
[**SetMFADuoConfig**](MFAConfigurationAPI.md#SetMFADuoConfig) | **Put** /mfa/duo-web/config | Set Duo MFA configuration
[**SetMFAOktaConfig**](MFAConfigurationAPI.md#SetMFAOktaConfig) | **Put** /mfa/okta-verify/config | Set Okta MFA configuration
[**TestMFAConfig**](MFAConfigurationAPI.md#TestMFAConfig) | **Get** /mfa/{method}/test | MFA method&#39;s test configuration



## DeleteMFAConfig

> MfaOktaConfig DeleteMFAConfig(ctx, method).Execute()

Delete MFA method configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func main() {
    method := "okta-verify" // string | The name of the MFA method. The currently supported method names are 'okta-verify' and 'duo-web'.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAConfigurationAPI.DeleteMFAConfig(context.Background(), method).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAConfigurationAPI.DeleteMFAConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMFAConfig`: MfaOktaConfig
    fmt.Fprintf(os.Stdout, "Response from `MFAConfigurationAPI.DeleteMFAConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**method** | **string** | The name of the MFA method. The currently supported method names are &#39;okta-verify&#39; and &#39;duo-web&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMFAConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
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
    openapiclient "github.com/sailpoint-oss/golang-sdk"
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

