# \OrgApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSettings**](OrgApi.md#GetOrgSettings) | **Get** /org | Retrieves your org settings.
[**UpdateOrgSettings**](OrgApi.md#UpdateOrgSettings) | **Patch** /org | Updates one or more org attributes.



## GetOrgSettings

> Org GetOrgSettings(ctx).Execute()

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
    // response from `GetOrgSettings`: Org
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**Org**](Org.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> Org UpdateOrgSettings(ctx).OrgEto(orgEto).Execute()

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
    orgEto := *openapiclient.NewOrgEto() // OrgEto | Org settings to update.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgApi.UpdateOrgSettings(context.Background()).OrgEto(orgEto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgApi.UpdateOrgSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgSettings`: Org
    fmt.Fprintf(os.Stdout, "Response from `OrgApi.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgEto** | [**OrgEto**](OrgEto.md) | Org settings to update. | 

### Return type

[**Org**](Org.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

