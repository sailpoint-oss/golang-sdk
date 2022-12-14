# \AccessRequestsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccessRequests**](AccessRequestsApi.md#ListAccessRequests) | **Get** /access-requests | Lists the access request for an identity.



## ListAccessRequests

> []AccessRequest ListAccessRequests(ctx).Offset(offset).Limit(limit).Execute()

Lists the access request for an identity.



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
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessRequestsApi.ListAccessRequests(context.Background()).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestsApi.ListAccessRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessRequests`: []AccessRequest
    fmt.Fprintf(os.Stdout, "Response from `AccessRequestsApi.ListAccessRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]AccessRequest**](AccessRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

