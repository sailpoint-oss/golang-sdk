# \IdentitiesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SynchronizeAttributesForIdentity**](IdentitiesApi.md#SynchronizeAttributesForIdentity) | **Post** /identities/{identityId}/synchronize-attributes | Attribute synchronization for single identity.



## SynchronizeAttributesForIdentity

> IdentitySyncJob SynchronizeAttributesForIdentity(ctx, identityId).Execute()

Attribute synchronization for single identity.



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
    identityId := "identityId_example" // string | The Identity id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentitiesApi.SynchronizeAttributesForIdentity(context.Background(), identityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentitiesApi.SynchronizeAttributesForIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SynchronizeAttributesForIdentity`: IdentitySyncJob
    fmt.Fprintf(os.Stdout, "Response from `IdentitiesApi.SynchronizeAttributesForIdentity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityId** | **string** | The Identity id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSynchronizeAttributesForIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdentitySyncJob**](IdentitySyncJob.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

