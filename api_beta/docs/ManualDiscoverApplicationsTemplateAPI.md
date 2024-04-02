# \ManualDiscoverApplicationsTemplateAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetManualDiscoverApplicationsCsvTemplate**](ManualDiscoverApplicationsTemplateAPI.md#GetManualDiscoverApplicationsCsvTemplate) | **Get** /manual-discover-applications-template | CSV template download for discovery



## GetManualDiscoverApplicationsCsvTemplate

> ManualDiscoverApplicationsTemplate GetManualDiscoverApplicationsCsvTemplate(ctx).Execute()

CSV template download for discovery



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualDiscoverApplicationsTemplateAPI.GetManualDiscoverApplicationsCsvTemplate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualDiscoverApplicationsTemplateAPI.GetManualDiscoverApplicationsCsvTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManualDiscoverApplicationsCsvTemplate`: ManualDiscoverApplicationsTemplate
    fmt.Fprintf(os.Stdout, "Response from `ManualDiscoverApplicationsTemplateAPI.GetManualDiscoverApplicationsCsvTemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualDiscoverApplicationsCsvTemplateRequest struct via the builder pattern


### Return type

[**ManualDiscoverApplicationsTemplate**](ManualDiscoverApplicationsTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

