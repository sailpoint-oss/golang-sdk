# \OrgApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSettings**](OrgApi.md#GetOrgSettings) | **Get** /org | Retrieves your org settings.
[**UpdateOrgSettings**](OrgApi.md#UpdateOrgSettings) | **Patch** /org | Updates one or more org attributes.



## GetOrgSettings

> GetOrgSettings200Response GetOrgSettings(ctx).Execute()

Retrieves your org settings.



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
    resp, r, err := apiClient.OrgApi.GetOrgSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.GetOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgSettings`: GetOrgSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**GetOrgSettings200Response**](GetOrgSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> GetOrgSettings200Response UpdateOrgSettings(ctx).UpdateOrgSettingsRequest(updateOrgSettingsRequest).Execute()

Updates one or more org attributes.



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
    updateOrgSettingsRequest := *openapiclient.NewUpdateOrgSettingsRequest() // UpdateOrgSettingsRequest | Org settings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.UpdateOrgSettings(context.Background()).UpdateOrgSettingsRequest(updateOrgSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.UpdateOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgSettings`: GetOrgSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateOrgSettingsRequest** | [**UpdateOrgSettingsRequest**](UpdateOrgSettingsRequest.md) | Org settings to update. | 

### Return type

[**GetOrgSettings200Response**](GetOrgSettings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

