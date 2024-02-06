# \AccessRequestIdentityMetricsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessRequestIdentityMetrics**](AccessRequestIdentityMetricsAPI.md#GetAccessRequestIdentityMetrics) | **Get** /access-request-identity-metrics/{identityId}/requested-objects/{requestedObjectId}/type/{type} | Return access request identity metrics



## GetAccessRequestIdentityMetrics

> map[string]interface{} GetAccessRequestIdentityMetrics(ctx, identityId, requestedObjectId, type_).Execute()

Return access request identity metrics



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
    identityId := "7025c863-c270-4ba6-beea-edf3cb091573" // string | Identity's ID.
    requestedObjectId := "2db501be-f0fb-4cc5-a695-334133c52891" // string | Requested access item's ID.
    type_ := "ENTITLEMENT" // string | Requested access item's type.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestIdentityMetricsAPI.GetAccessRequestIdentityMetrics(context.Background(), identityId, requestedObjectId, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestIdentityMetricsAPI.GetAccessRequestIdentityMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessRequestIdentityMetrics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestIdentityMetricsAPI.GetAccessRequestIdentityMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | Identity&#39;s ID. | 
**requestedObjectId** | **string** | Requested access item&#39;s ID. | 
**type_** | **string** | Requested access item&#39;s type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessRequestIdentityMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

