# \SODViolationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetArmConfig**](SODViolationsApi.md#GetArmConfig) | **Get** /sod-violations/config | Expose just the ARM config
[**StartPredictSodViolations**](SODViolationsApi.md#StartPredictSodViolations) | **Post** /sod-violations/predict | Predict SOD violations for identity.
[**StartViolationCheck**](SODViolationsApi.md#StartViolationCheck) | **Post** /sod-violations/check | Check SOD violations



## GetArmConfig

> PublicOrgConfigArmData GetArmConfig(ctx).Execute()

Expose just the ARM config



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
    resp, r, err := apiClient.SODViolationsApi.GetArmConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsApi.GetArmConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArmConfig`: PublicOrgConfigArmData
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsApi.GetArmConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetArmConfigRequest struct via the builder pattern


### Return type

[**PublicOrgConfigArmData**](PublicOrgConfigArmData.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPredictSodViolations

> ViolationPrediction StartPredictSodViolations(ctx).IdentityWithNewAccess(identityWithNewAccess).Execute()

Predict SOD violations for identity.



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
    identityWithNewAccess := *openapiclient.NewIdentityWithNewAccess("2c91808568c529c60168cca6f90c1313", []openapiclient.IdentityWithNewAccessAccessRefsInner{*openapiclient.NewIdentityWithNewAccessAccessRefsInner()}) // IdentityWithNewAccess | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODViolationsApi.StartPredictSodViolations(context.Background()).IdentityWithNewAccess(identityWithNewAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsApi.StartPredictSodViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartPredictSodViolations`: ViolationPrediction
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsApi.StartPredictSodViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartPredictSodViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityWithNewAccess** | [**IdentityWithNewAccess**](IdentityWithNewAccess.md) |  | 

### Return type

[**ViolationPrediction**](ViolationPrediction.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartViolationCheck

> SodViolationCheck StartViolationCheck(ctx).IdentityWithNewAccess1(identityWithNewAccess1).Execute()

Check SOD violations



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
    identityWithNewAccess1 := *openapiclient.NewIdentityWithNewAccess1("2c91809050db617d0150e0bf3215385e", []openapiclient.BaseReferenceDto{*openapiclient.NewBaseReferenceDto()}) // IdentityWithNewAccess1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SODViolationsApi.StartViolationCheck(context.Background()).IdentityWithNewAccess1(identityWithNewAccess1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsApi.StartViolationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartViolationCheck`: SodViolationCheck
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsApi.StartViolationCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartViolationCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityWithNewAccess1** | [**IdentityWithNewAccess1**](IdentityWithNewAccess1.md) |  | 

### Return type

[**SodViolationCheck**](SodViolationCheck.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

