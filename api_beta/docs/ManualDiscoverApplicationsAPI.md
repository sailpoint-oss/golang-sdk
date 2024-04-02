# \ManualDiscoverApplicationsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/beta*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendManualDiscoverApplicationsCsvTemplate**](ManualDiscoverApplicationsAPI.md#SendManualDiscoverApplicationsCsvTemplate) | **Post** /manual-discover-applications | CSV Upload to discover applications



## SendManualDiscoverApplicationsCsvTemplate

> ManualDiscoverApplications SendManualDiscoverApplicationsCsvTemplate(ctx).CsvFile(csvFile).Execute()

CSV Upload to discover applications



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
    csvFile := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ManualDiscoverApplicationsAPI.SendManualDiscoverApplicationsCsvTemplate(context.Background()).CsvFile(csvFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ManualDiscoverApplicationsAPI.SendManualDiscoverApplicationsCsvTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendManualDiscoverApplicationsCsvTemplate`: ManualDiscoverApplications
    fmt.Fprintf(os.Stdout, "Response from `ManualDiscoverApplicationsAPI.SendManualDiscoverApplicationsCsvTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendManualDiscoverApplicationsCsvTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **csvFile** | ***os.File** |  | 

### Return type

[**ManualDiscoverApplications**](ManualDiscoverApplications.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: multipart/form-data, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

