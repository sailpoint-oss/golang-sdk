# \CertificationCampaignFiltersAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaignFilter**](CertificationCampaignFiltersAPI.md#CreateCampaignFilter) | **Post** /campaign-filters | Create a Campaign Filter
[**DeleteCampaignFilters**](CertificationCampaignFiltersAPI.md#DeleteCampaignFilters) | **Post** /campaign-filters/delete | Deletes Campaign Filters
[**GetCampaignFilterById**](CertificationCampaignFiltersAPI.md#GetCampaignFilterById) | **Get** /campaign-filters/{id} | Get Campaign Filter by ID
[**ListCampaignFilters**](CertificationCampaignFiltersAPI.md#ListCampaignFilters) | **Get** /campaign-filters | List Campaign Filters
[**UpdateCampaignFilter**](CertificationCampaignFiltersAPI.md#UpdateCampaignFilter) | **Post** /campaign-filters/{id} | Updates a Campaign Filter



## Create a Campaign Filter

> CampaignFilterDetails CreateCampaignFilter(ctx).CampaignFilterDetails(campaignFilterDetails).Execute()

Create a campaign Filter based on filter details and criteria.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    campaignFilterDetails := *sailpoint.NewCampaignFilterDetails("Identity Attribute Campaign Filter", "Campaign filter to certify data based on specified property of Identity Attribute.", "SailPoint Support", map[string]interface{}(INCLUSION)) // CampaignFilterDetails | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignFiltersAPI.CreateCampaignFilter(context.Background()).CampaignFilterDetails(campaignFilterDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignFiltersAPI.CreateCampaignFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaignFilter`: CampaignFilterDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignFiltersAPI.CreateCampaignFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignFilterDetails** | [**CampaignFilterDetails**](CampaignFilterDetails.md) |  | 

### Return type

[**CampaignFilterDetails**](CampaignFilterDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deletes Campaign Filters

> DeleteCampaignFilters(ctx).RequestBody(requestBody).Execute()

Deletes campaign filters whose Ids are specified in the provided list of campaign filter Ids. Authorized callers must be an ORG_ADMIN or a CERT_ADMIN.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    requestBody := []string{"Property_example"} // []string | A json list of IDs of campaign filters to delete.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.CertificationCampaignFiltersAPI.DeleteCampaignFilters(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignFiltersAPI.DeleteCampaignFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | A json list of IDs of campaign filters to delete. | 

### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Campaign Filter by ID

> []CampaignFilterDetails GetCampaignFilterById(ctx, filterId).Execute()

Retrieves information for an existing campaign filter using the filter's ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    filterId := "e9f9a1397b842fd5a65842087040d3ac" // string | The ID of the campaign filter to be retrieved.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignFiltersAPI.GetCampaignFilterById(context.Background(), filterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignFiltersAPI.GetCampaignFilterById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignFilterById`: []CampaignFilterDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignFiltersAPI.GetCampaignFilterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The ID of the campaign filter to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CampaignFilterDetails**](CampaignFilterDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List Campaign Filters

> []CampaignFilterDetails ListCampaignFilters(ctx).Limit(limit).Start(start).IncludeSystemFilters(includeSystemFilters).Execute()

Lists all Campaign Filters. Scope can be reduced via standard V3 query params.

All Campaign Filters matching the query params

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    start := int32(0) // int32 | Start/Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    includeSystemFilters := true // bool | If true, include system filters in the count and results, exclude them otherwise. If not provided any value for it then by default it is true. (optional) (default to true)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignFiltersAPI.ListCampaignFilters(context.Background()).Limit(limit).Start(start).IncludeSystemFilters(includeSystemFilters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignFiltersAPI.ListCampaignFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCampaignFilters`: []CampaignFilterDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignFiltersAPI.ListCampaignFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **start** | **int32** | Start/Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **includeSystemFilters** | **bool** | If true, include system filters in the count and results, exclude them otherwise. If not provided any value for it then by default it is true. | [default to true]

### Return type

[**[]CampaignFilterDetails**](CampaignFilterDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Updates a Campaign Filter

> CampaignFilterDetails UpdateCampaignFilter(ctx, filterId).CampaignFilterDetails(campaignFilterDetails).Execute()

Updates an existing campaign filter using the filter's ID.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    filterId := "e9f9a1397b842fd5a65842087040d3ac" // string | The ID of the campaign filter being modified.
    campaignFilterDetails := *sailpoint.NewCampaignFilterDetails("Identity Attribute Campaign Filter", "Campaign filter to certify data based on specified property of Identity Attribute.", "SailPoint Support", map[string]interface{}(INCLUSION)) // CampaignFilterDetails | A campaign filter details with updated field values.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignFiltersAPI.UpdateCampaignFilter(context.Background(), filterId).CampaignFilterDetails(campaignFilterDetails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignFiltersAPI.UpdateCampaignFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignFilter`: CampaignFilterDetails
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignFiltersAPI.UpdateCampaignFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **string** | The ID of the campaign filter being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignFilterDetails** | [**CampaignFilterDetails**](CampaignFilterDetails.md) | A campaign filter details with updated field values. | 

### Return type

[**CampaignFilterDetails**](CampaignFilterDetails.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

