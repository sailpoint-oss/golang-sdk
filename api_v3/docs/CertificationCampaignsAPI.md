# \CertificationCampaignsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteCampaign**](CertificationCampaignsAPI.md#CompleteCampaign) | **Post** /campaigns/{id}/complete | Complete a Campaign
[**CreateCampaign**](CertificationCampaignsAPI.md#CreateCampaign) | **Post** /campaigns | Create a campaign
[**CreateCampaignTemplate**](CertificationCampaignsAPI.md#CreateCampaignTemplate) | **Post** /campaign-templates | Create a Campaign Template
[**DeleteCampaignTemplate**](CertificationCampaignsAPI.md#DeleteCampaignTemplate) | **Delete** /campaign-templates/{id} | Delete a Campaign Template
[**DeleteCampaignTemplateSchedule**](CertificationCampaignsAPI.md#DeleteCampaignTemplateSchedule) | **Delete** /campaign-templates/{id}/schedule | Delete Campaign Template Schedule
[**DeleteCampaigns**](CertificationCampaignsAPI.md#DeleteCampaigns) | **Post** /campaigns/delete | Delete Campaigns
[**GetActiveCampaigns**](CertificationCampaignsAPI.md#GetActiveCampaigns) | **Get** /campaigns | List Campaigns
[**GetCampaign**](CertificationCampaignsAPI.md#GetCampaign) | **Get** /campaigns/{id} | Get Campaign
[**GetCampaignReports**](CertificationCampaignsAPI.md#GetCampaignReports) | **Get** /campaigns/{id}/reports | Get Campaign Reports
[**GetCampaignReportsConfig**](CertificationCampaignsAPI.md#GetCampaignReportsConfig) | **Get** /campaigns/reports-configuration | Get Campaign Reports Configuration
[**GetCampaignTemplate**](CertificationCampaignsAPI.md#GetCampaignTemplate) | **Get** /campaign-templates/{id} | Get a Campaign Template
[**GetCampaignTemplateSchedule**](CertificationCampaignsAPI.md#GetCampaignTemplateSchedule) | **Get** /campaign-templates/{id}/schedule | Get Campaign Template Schedule
[**GetCampaignTemplates**](CertificationCampaignsAPI.md#GetCampaignTemplates) | **Get** /campaign-templates | List Campaign Templates
[**Move**](CertificationCampaignsAPI.md#Move) | **Post** /campaigns/{id}/reassign | Reassign Certifications
[**PatchCampaignTemplate**](CertificationCampaignsAPI.md#PatchCampaignTemplate) | **Patch** /campaign-templates/{id} | Update a Campaign Template
[**SetCampaignReportsConfig**](CertificationCampaignsAPI.md#SetCampaignReportsConfig) | **Put** /campaigns/reports-configuration | Set Campaign Reports Configuration
[**SetCampaignTemplateSchedule**](CertificationCampaignsAPI.md#SetCampaignTemplateSchedule) | **Put** /campaign-templates/{id}/schedule | Set Campaign Template Schedule
[**StartCampaign**](CertificationCampaignsAPI.md#StartCampaign) | **Post** /campaigns/{id}/activate | Activate a Campaign
[**StartCampaignRemediationScan**](CertificationCampaignsAPI.md#StartCampaignRemediationScan) | **Post** /campaigns/{id}/run-remediation-scan | Run Campaign Remediation Scan
[**StartCampaignReport**](CertificationCampaignsAPI.md#StartCampaignReport) | **Post** /campaigns/{id}/run-report/{type} | Run Campaign Report
[**StartGenerateCampaignTemplate**](CertificationCampaignsAPI.md#StartGenerateCampaignTemplate) | **Post** /campaign-templates/{id}/generate | Generate a Campaign from Template
[**UpdateCampaign**](CertificationCampaignsAPI.md#UpdateCampaign) | **Patch** /campaigns/{id} | Update a Campaign



## CompleteCampaign

> map[string]interface{} CompleteCampaign(ctx, id).CampaignCompleteOptions(campaignCompleteOptions).Execute()

Complete a Campaign



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Campaign ID.
    campaignCompleteOptions := *openapiclient.NewCampaignCompleteOptions() // CampaignCompleteOptions | Optional. Default behavior is for the campaign to auto-approve upon completion, unless autoCompleteAction=REVOKE (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.CompleteCampaign(context.Background(), id).CampaignCompleteOptions(campaignCompleteOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.CompleteCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteCampaign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.CompleteCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Campaign ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignCompleteOptions** | [**CampaignCompleteOptions**](CampaignCompleteOptions.md) | Optional. Default behavior is for the campaign to auto-approve upon completion, unless autoCompleteAction&#x3D;REVOKE | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    campaign := *openapiclient.NewCampaign("Manager Campaign", "Everyone needs to be reviewed by their manager", "MANAGER") // Campaign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.CreateCampaign(context.Background()).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.CreateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaign`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.CreateCampaign`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignTemplate

> CampaignTemplate CreateCampaignTemplate(ctx).CampaignTemplate(campaignTemplate).Execute()

Create a Campaign Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    campaignTemplate := *openapiclient.NewCampaignTemplate("Manager Campaign Template", "Template for the annual manager campaign.", time.Now(), time.Now(), *openapiclient.NewCampaign("Manager Campaign", "Everyone needs to be reviewed by their manager", "MANAGER")) // CampaignTemplate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.CreateCampaignTemplate(context.Background()).CampaignTemplate(campaignTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.CreateCampaignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaignTemplate`: CampaignTemplate
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.CreateCampaignTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignTemplate** | [**CampaignTemplate**](CampaignTemplate.md) |  | 

### Return type

[**CampaignTemplate**](CampaignTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignTemplate

> DeleteCampaignTemplate(ctx, id).Execute()

Delete a Campaign Template



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
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the campaign template being deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificationCampaignsAPI.DeleteCampaignTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.DeleteCampaignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignTemplateSchedule

> DeleteCampaignTemplateSchedule(ctx, id).Execute()

Delete Campaign Template Schedule



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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | ID of the campaign template whose schedule is being deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificationCampaignsAPI.DeleteCampaignTemplateSchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.DeleteCampaignTemplateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template whose schedule is being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignTemplateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaigns

> map[string]interface{} DeleteCampaigns(ctx).CampaignsDeleteRequest(campaignsDeleteRequest).Execute()

Delete Campaigns



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
    campaignsDeleteRequest := *openapiclient.NewCampaignsDeleteRequest() // CampaignsDeleteRequest | IDs of the campaigns to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.DeleteCampaigns(context.Background()).CampaignsDeleteRequest(campaignsDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.DeleteCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaigns`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.DeleteCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignsDeleteRequest** | [**CampaignsDeleteRequest**](CampaignsDeleteRequest.md) | IDs of the campaigns to delete. | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveCampaigns

> []GetActiveCampaigns200ResponseInner GetActiveCampaigns(ctx).Detail(detail).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

List Campaigns



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
    detail := "FULL" // string | Determines whether slim, or increased level of detail is provided for each campaign in the returned list. Slim is the default behavior. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name eq "Manager Campaign"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **status**: *eq, in* (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created** (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetActiveCampaigns(context.Background()).Detail(detail).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetActiveCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveCampaigns`: []GetActiveCampaigns200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetActiveCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detail** | **string** | Determines whether slim, or increased level of detail is provided for each campaign in the returned list. Slim is the default behavior. | 
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **status**: *eq, in* | 
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created** | 

### Return type

[**[]GetActiveCampaigns200ResponseInner**](GetActiveCampaigns200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> GetActiveCampaigns200ResponseInner GetCampaign(ctx, id).Detail(detail).Execute()

Get Campaign



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | ID of the campaign to be retrieved.
    detail := "FULL" // string | Determines whether slim, or increased level of detail is provided for each campaign in the returned list. Slim is the default behavior. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaign(context.Background(), id).Detail(detail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: GetActiveCampaigns200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detail** | **string** | Determines whether slim, or increased level of detail is provided for each campaign in the returned list. Slim is the default behavior. | 

### Return type

[**GetActiveCampaigns200ResponseInner**](GetActiveCampaigns200ResponseInner.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignReports

> []CampaignReport GetCampaignReports(ctx, id).Execute()

Get Campaign Reports



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | ID of the campaign whose reports are being fetched.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaignReports(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaignReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignReports`: []CampaignReport
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaignReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign whose reports are being fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]CampaignReport**](CampaignReport.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignReportsConfig

> CampaignReportsConfig GetCampaignReportsConfig(ctx).Execute()

Get Campaign Reports Configuration



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
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaignReportsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaignReportsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignReportsConfig`: CampaignReportsConfig
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaignReportsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignReportsConfigRequest struct via the builder pattern


### Return type

[**CampaignReportsConfig**](CampaignReportsConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplate

> CampaignTemplate GetCampaignTemplate(ctx, id).Execute()

Get a Campaign Template



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
    id := "2c9180835d191a86015d28455b4a2329" // string | Requested campaign template's ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaignTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignTemplate`: CampaignTemplate
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaignTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Requested campaign template&#39;s ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignTemplate**](CampaignTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplateSchedule

> Schedule GetCampaignTemplateSchedule(ctx, id).Execute()

Get Campaign Template Schedule



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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | ID of the campaign template whose schedule is being fetched.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaignTemplateSchedule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaignTemplateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignTemplateSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaignTemplateSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template whose schedule is being fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTemplateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Schedule**](Schedule.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplates

> []CampaignTemplate GetCampaignTemplates(ctx).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

List Campaign Templates



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
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)
    filters := "name eq "manager template"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, ge, gt, in, le, lt, ne, sw*  **id**: *eq, ge, gt, in, le, lt, ne, sw* (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.GetCampaignTemplates(context.Background()).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaignTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignTemplates`: []CampaignTemplate
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaignTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 250]
 **offset** | **int32** | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to 0]
 **count** | **bool** | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count&#x3D;true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. | [default to false]
 **sorters** | **string** | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** | 
 **filters** | **string** | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, ge, gt, in, le, lt, ne, sw*  **id**: *eq, ge, gt, in, le, lt, ne, sw* | 

### Return type

[**[]CampaignTemplate**](CampaignTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Move

> CertificationTask Move(ctx, id).AdminReviewReassign(adminReviewReassign).Execute()

Reassign Certifications



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The certification campaign ID
    adminReviewReassign := *openapiclient.NewAdminReviewReassign() // AdminReviewReassign | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.Move(context.Background(), id).AdminReviewReassign(adminReviewReassign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.Move``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Move`: CertificationTask
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.Move`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The certification campaign ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminReviewReassign** | [**AdminReviewReassign**](AdminReviewReassign.md) |  | 

### Return type

[**CertificationTask**](CertificationTask.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCampaignTemplate

> CampaignTemplate PatchCampaignTemplate(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Update a Campaign Template



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
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the campaign template being modified.
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * deadlineDuration * campaign (all fields that are allowed during create) 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.PatchCampaignTemplate(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.PatchCampaignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchCampaignTemplate`: CampaignTemplate
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.PatchCampaignTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCampaignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * deadlineDuration * campaign (all fields that are allowed during create)  | 

### Return type

[**CampaignTemplate**](CampaignTemplate.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCampaignReportsConfig

> CampaignReportsConfig SetCampaignReportsConfig(ctx).CampaignReportsConfig(campaignReportsConfig).Execute()

Set Campaign Reports Configuration



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
    campaignReportsConfig := *openapiclient.NewCampaignReportsConfig() // CampaignReportsConfig | Campaign report configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.SetCampaignReportsConfig(context.Background()).CampaignReportsConfig(campaignReportsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.SetCampaignReportsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetCampaignReportsConfig`: CampaignReportsConfig
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.SetCampaignReportsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetCampaignReportsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignReportsConfig** | [**CampaignReportsConfig**](CampaignReportsConfig.md) | Campaign report configuration. | 

### Return type

[**CampaignReportsConfig**](CampaignReportsConfig.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCampaignTemplateSchedule

> SetCampaignTemplateSchedule(ctx, id).Schedule(schedule).Execute()

Set Campaign Template Schedule



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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | ID of the campaign template being scheduled.
    schedule := *openapiclient.NewSchedule("WEEKLY", *openapiclient.NewScheduleHours("LIST", []string{"Values_example"})) // Schedule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CertificationCampaignsAPI.SetCampaignTemplateSchedule(context.Background(), id).Schedule(schedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.SetCampaignTemplateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template being scheduled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCampaignTemplateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schedule** | [**Schedule**](Schedule.md) |  | 

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


## StartCampaign

> map[string]interface{} StartCampaign(ctx, id).ActivateCampaignOptions(activateCampaignOptions).Execute()

Activate a Campaign



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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | Campaign ID.
    activateCampaignOptions := *openapiclient.NewActivateCampaignOptions() // ActivateCampaignOptions | Optional. If no timezone is specified, the standard UTC timezone is used (i.e. UTC+00:00). Although this can take any timezone, the intended value is the caller's timezone. The activation time calculated from the given timezone may cause the campaign deadline time to be modified, but it will remain within the original date. The timezone must be in a valid ISO 8601 format. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.StartCampaign(context.Background(), id).ActivateCampaignOptions(activateCampaignOptions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.StartCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCampaign`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.StartCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Campaign ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activateCampaignOptions** | [**ActivateCampaignOptions**](ActivateCampaignOptions.md) | Optional. If no timezone is specified, the standard UTC timezone is used (i.e. UTC+00:00). Although this can take any timezone, the intended value is the caller&#39;s timezone. The activation time calculated from the given timezone may cause the campaign deadline time to be modified, but it will remain within the original date. The timezone must be in a valid ISO 8601 format. | 

### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCampaignRemediationScan

> map[string]interface{} StartCampaignRemediationScan(ctx, id).Execute()

Run Campaign Remediation Scan



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | ID of the campaign the remediation scan is being run for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.StartCampaignRemediationScan(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.StartCampaignRemediationScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCampaignRemediationScan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.StartCampaignRemediationScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign the remediation scan is being run for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCampaignRemediationScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartCampaignReport

> map[string]interface{} StartCampaignReport(ctx, id, type_).Execute()

Run Campaign Report



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | ID of the campaign the report is being run for.
    type_ := openapiclient.ReportType("CAMPAIGN_COMPOSITION_REPORT") // ReportType | Type of the report to run.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.StartCampaignReport(context.Background(), id, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.StartCampaignReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartCampaignReport`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.StartCampaignReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign the report is being run for. | 
**type_** | [**ReportType**](.md) | Type of the report to run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartCampaignReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartGenerateCampaignTemplate

> CampaignReference StartGenerateCampaignTemplate(ctx, id).Execute()

Generate a Campaign from Template



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
    id := "2c9180835d191a86015d28455b4a2329" // string | ID of the campaign template to use for generation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.StartGenerateCampaignTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.StartGenerateCampaignTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartGenerateCampaignTemplate`: CampaignReference
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.StartGenerateCampaignTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template to use for generation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartGenerateCampaignTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignReference**](CampaignReference.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> SlimCampaign UpdateCampaign(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Update a Campaign



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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | ID of the campaign template being modified.
    jsonPatchOperation := []openapiclient.JsonPatchOperation{*openapiclient.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The fields that can be patched differ based on the status of the campaign.  When the campaign is in the *STAGED* status, you can patch these fields: * name * description * recommendationsEnabled * deadline * emailNotificationEnabled * autoRevokeAllowed  When the campaign is in the *ACTIVE* status, you can patch these fields: * deadline 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificationCampaignsAPI.UpdateCampaign(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.UpdateCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaign`: SlimCampaign
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the campaign template being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The fields that can be patched differ based on the status of the campaign.  When the campaign is in the *STAGED* status, you can patch these fields: * name * description * recommendationsEnabled * deadline * emailNotificationEnabled * autoRevokeAllowed  When the campaign is in the *ACTIVE* status, you can patch these fields: * deadline  | 

### Return type

[**SlimCampaign**](SlimCampaign.md)

### Authorization

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

