# \SODViolationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PredictViolations**](SODViolationsApi.md#PredictViolations) | **Post** /sod-violations/predict | Predict SOD violations for the given identity if they were granted the given access.



## PredictViolations

> ViolationPrediction PredictViolations(ctx).IdentityWithNewAccess(identityWithNewAccess).Execute()

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
    resp, r, err := apiClient.SODViolationsApi.PredictViolations(context.Background()).IdentityWithNewAccess(identityWithNewAccess).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SODViolationsApi.PredictViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PredictViolations`: ViolationPrediction
    fmt.Fprintf(os.Stdout, "Response from `SODViolationsApi.PredictViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredictViolationsRequest struct via the builder pattern


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

