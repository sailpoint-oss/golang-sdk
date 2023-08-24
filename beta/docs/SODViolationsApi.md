# \SODViolationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PredictSodViolations**](SODViolationsApi.md#PredictSodViolations) | **Post** /sod-violations/predict | Predict SOD violations for the given identity if they were granted the given access.



## PredictSodViolations

> ViolationPrediction PredictSodViolations(ctx).IdentityWithNewAccess(identityWithNewAccess).Execute()

Predict SOD violations for the given identity if they were granted the given access.



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
    resp, r, err := apiClient.SODViolationsApi.PredictSodViolations(context.Background()).IdentityWithNewAccess(identityWithNewAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsApi.PredictSodViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PredictSodViolations`: ViolationPrediction
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsApi.PredictSodViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredictSodViolationsRequest struct via the builder pattern


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

