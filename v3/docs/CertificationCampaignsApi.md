# \CertificationCampaignsApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CertificationCampaignsApi.md#CreateCampaign) | **Post** /campaigns | Create a campaign
[**GetCampaign**](CertificationCampaignsApi.md#GetCampaign) | **Get** /campaigns/{id} | Get a campaign



## CreateCampaign

> Campaign CreateCampaign(ctx).Campaign(campaign).Execute()

Create a campaign



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
    campaign := *openapiclient.NewCampaign("Manager Campaign", "Everyone needs to be reviewed by their manager", "MANAGER") // Campaign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsApi.CreateCampaign(context.Background()).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsApi.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsApi.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaign** | [**Campaign**](Campaign.md) |  | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> SlimCampaign GetCampaign(ctx, id).Execute()

Get a campaign



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign to be retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsApi.GetCampaign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsApi.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: SlimCampaign
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsApi.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the campaign to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlimCampaign**](SlimCampaign.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

