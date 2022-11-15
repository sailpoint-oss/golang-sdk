# \AccountAggregationsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountAggregationStatus**](AccountAggregationsApi.md#GetAccountAggregationStatus) | **Get** /account-aggregations/{id}/status | In-progress Account Aggregation status



## GetAccountAggregationStatus

> AccountAggregationStatus GetAccountAggregationStatus(ctx, id).Execute()

In-progress Account Aggregation status



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
    id := "2c91808477a6b0c60177a81146b8110b" // string | The account aggregation id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountAggregationsApi.GetAccountAggregationStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountAggregationsApi.GetAccountAggregationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountAggregationStatus`: AccountAggregationStatus
    fmt.Fprintf(os.Stdout, "Response from `AccountAggregationsApi.GetAccountAggregationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The account aggregation id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAggregationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAggregationStatus**](AccountAggregationStatus.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

