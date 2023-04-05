# \PasswordDictionaryApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordDictionary**](PasswordDictionaryApi.md#GetPasswordDictionary) | **Get** /password-dictionary | Get Password Dictionary



## GetPasswordDictionary

> string GetPasswordDictionary(ctx).Execute()

Get Password Dictionary



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
    resp, r, err := apiClient.PasswordDictionaryApi.GetPasswordDictionary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordDictionaryApi.GetPasswordDictionary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordDictionary`: string
    fmt.Fprintf(os.Stdout, "Response from `PasswordDictionaryApi.GetPasswordDictionary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordDictionaryRequest struct via the builder pattern


### Return type

**string**

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

