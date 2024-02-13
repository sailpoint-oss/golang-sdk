# \SODViolationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StartPredictSodViolations**](SODViolationsAPI.md#StartPredictSodViolations) | **Post** /sod-violations/predict | Predict SOD violations for identity.
[**StartViolationCheck**](SODViolationsAPI.md#StartViolationCheck) | **Post** /sod-violations/check | Check SOD violations



## Predict SOD violations for identity.

> ViolationPrediction StartPredictSodViolations(ctx).IdentityWithNewAccess(identityWithNewAccess).Execute()

This API is used to check if granting some additional accesses would cause the subject to be in violation of any SOD policies. Returns the violations that would be caused.

A token with ORG_ADMIN or API authority is required to call this API.

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
    identityWithNewAccess := *sailpoint.NewIdentityWithNewAccess("2c91808568c529c60168cca6f90c1313", []sailpoint.IdentityWithNewAccessAccessRefsInner{*sailpoint.NewIdentityWithNewAccessAccessRefsInner()}) // IdentityWithNewAccess | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODViolationsAPI.StartPredictSodViolations(context.Background()).IdentityWithNewAccess(identityWithNewAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsAPI.StartPredictSodViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartPredictSodViolations`: ViolationPrediction
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsAPI.StartPredictSodViolations`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Check SOD violations

> SodViolationCheck StartViolationCheck(ctx).IdentityWithNewAccess1(identityWithNewAccess1).Execute()

This API initiates a SOD policy verification asynchronously.

A token with ORG_ADMIN authority is required to call this API.

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
    identityWithNewAccess1 := *sailpoint.NewIdentityWithNewAccess1("2c91809050db617d0150e0bf3215385e", []sailpoint.EntitlementRef{*sailpoint.NewEntitlementRef()}) // IdentityWithNewAccess1 | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.SODViolationsAPI.StartViolationCheck(context.Background()).IdentityWithNewAccess1(identityWithNewAccess1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsAPI.StartViolationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartViolationCheck`: SodViolationCheck
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsAPI.StartViolationCheck`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

