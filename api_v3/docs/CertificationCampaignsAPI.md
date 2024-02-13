# \CertificationCampaignsAPI

All URIs are relative to *https://sailpoint.api.identitynow.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteCampaign**](CertificationCampaignsAPI.md#CompleteCampaign) | **Post** /campaigns/{id}/complete | Complete a Campaign
[**CreateCampaign**](CertificationCampaignsAPI.md#CreateCampaign) | **Post** /campaigns | Create a campaign
[**CreateCampaignTemplate**](CertificationCampaignsAPI.md#CreateCampaignTemplate) | **Post** /campaign-templates | Create a Campaign Template
[**DeleteCampaignTemplate**](CertificationCampaignsAPI.md#DeleteCampaignTemplate) | **Delete** /campaign-templates/{id} | Delete a Campaign Template
[**DeleteCampaignTemplateSchedule**](CertificationCampaignsAPI.md#DeleteCampaignTemplateSchedule) | **Delete** /campaign-templates/{id}/schedule | Deletes a Campaign Template&#39;s Schedule
[**DeleteCampaigns**](CertificationCampaignsAPI.md#DeleteCampaigns) | **Post** /campaigns/delete | Deletes Campaigns
[**GetActiveCampaigns**](CertificationCampaignsAPI.md#GetActiveCampaigns) | **Get** /campaigns | List Campaigns
[**GetCampaign**](CertificationCampaignsAPI.md#GetCampaign) | **Get** /campaigns/{id} | Get a campaign
[**GetCampaignReports**](CertificationCampaignsAPI.md#GetCampaignReports) | **Get** /campaigns/{id}/reports | Get Campaign Reports
[**GetCampaignReportsConfig**](CertificationCampaignsAPI.md#GetCampaignReportsConfig) | **Get** /campaigns/reports-configuration | Get Campaign Reports Configuration
[**GetCampaignTemplate**](CertificationCampaignsAPI.md#GetCampaignTemplate) | **Get** /campaign-templates/{id} | Get a Campaign Template
[**GetCampaignTemplateSchedule**](CertificationCampaignsAPI.md#GetCampaignTemplateSchedule) | **Get** /campaign-templates/{id}/schedule | Gets a Campaign Template&#39;s Schedule
[**ListCampaignTemplates**](CertificationCampaignsAPI.md#ListCampaignTemplates) | **Get** /campaign-templates | List Campaign Templates
[**Move**](CertificationCampaignsAPI.md#Move) | **Post** /campaigns/{id}/reassign | Reassign Certifications
[**PatchCampaignTemplate**](CertificationCampaignsAPI.md#PatchCampaignTemplate) | **Patch** /campaign-templates/{id} | Update a Campaign Template
[**SetCampaignReportsConfig**](CertificationCampaignsAPI.md#SetCampaignReportsConfig) | **Put** /campaigns/reports-configuration | Set Campaign Reports Configuration
[**SetCampaignTemplateSchedule**](CertificationCampaignsAPI.md#SetCampaignTemplateSchedule) | **Put** /campaign-templates/{id}/schedule | Sets a Campaign Template&#39;s Schedule
[**StartCampaign**](CertificationCampaignsAPI.md#StartCampaign) | **Post** /campaigns/{id}/activate | Activate a Campaign
[**StartCampaignRemediationScan**](CertificationCampaignsAPI.md#StartCampaignRemediationScan) | **Post** /campaigns/{id}/run-remediation-scan | Run Campaign Remediation Scan
[**StartCampaignReport**](CertificationCampaignsAPI.md#StartCampaignReport) | **Post** /campaigns/{id}/run-report/{type} | Run Campaign Report
[**StartGenerateCampaignTemplate**](CertificationCampaignsAPI.md#StartGenerateCampaignTemplate) | **Post** /campaign-templates/{id}/generate | Generate a Campaign from Template
[**UpdateCampaign**](CertificationCampaignsAPI.md#UpdateCampaign) | **Patch** /campaigns/{id} | Update a Campaign



## Complete a Campaign

> map[string]interface{} CompleteCampaign(ctx, id).CampaignCompleteOptions(campaignCompleteOptions).Execute()

:::caution

This endpoint will run successfully for any campaigns that are **past due**.

This endpoint will return a content error if the campaign is **not past due**.

:::

Completes a certification campaign. This is provided to admins so that they
can complete a certification even if all items have not been completed.

Requires roles of CERT_ADMIN and ORG_ADMIN


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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The campaign id
    campaignCompleteOptions := *sailpoint.NewCampaignCompleteOptions() // CampaignCompleteOptions | Optional. Default behavior is for the campaign to auto-approve upon completion, unless autoCompleteAction=REVOKE (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.CompleteCampaign(context.Background(), id).CampaignCompleteOptions(campaignCompleteOptions).Execute()
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
**id** | **string** | The campaign id | 

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


## Create a campaign

> Campaign CreateCampaign(ctx).Campaign(campaign).Execute()

Creates a new Certification Campaign with the information provided in the request body.

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
    campaign := *sailpoint.NewCampaign("Manager Campaign", "Everyone needs to be reviewed by their manager", "MANAGER") // Campaign | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.CreateCampaign(context.Background()).Campaign(campaign).Execute()
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


## Create a Campaign Template

> CampaignTemplate CreateCampaignTemplate(ctx).CampaignTemplate(campaignTemplate).Execute()

Create a campaign Template based on campaign.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    sailpoint "github.com/sailpoint-oss/golang-sdk/v2"
)

func main() {
    campaignTemplate := *sailpoint.NewCampaignTemplate("Manager Campaign Template", "Template for the annual manager campaign.", time.Now(), time.Now(), *sailpoint.NewCampaign("Manager Campaign", "Everyone needs to be reviewed by their manager", "MANAGER")) // CampaignTemplate | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.CreateCampaignTemplate(context.Background()).CampaignTemplate(campaignTemplate).Execute()
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


## Delete a Campaign Template

> DeleteCampaignTemplate(ctx, id).Execute()

Deletes a campaign template by ID.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The ID of the campaign template being deleted.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.CertificationCampaignsAPI.DeleteCampaignTemplate(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign template being deleted. | 

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


## Deletes a Campaign Template's Schedule

> DeleteCampaignTemplateSchedule(ctx, id).Execute()

Deletes the schedule for a campaign template. Returns a 404 if there is no schedule set.

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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | The ID of the campaign template whose schedule is being deleted.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.CertificationCampaignsAPI.DeleteCampaignTemplateSchedule(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign template whose schedule is being deleted. | 

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


## Deletes Campaigns

> map[string]interface{} DeleteCampaigns(ctx).CampaignsDeleteRequest(campaignsDeleteRequest).Execute()

Deletes campaigns whose Ids are specified in the provided list of campaign Ids. Authorized callers must be an ORG_ADMIN or a CERT_ADMIN.

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
    campaignsDeleteRequest := *sailpoint.NewCampaignsDeleteRequest() // CampaignsDeleteRequest | The ids of the campaigns to delete.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.DeleteCampaigns(context.Background()).CampaignsDeleteRequest(campaignsDeleteRequest).Execute()
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
 **campaignsDeleteRequest** | [**CampaignsDeleteRequest**](CampaignsDeleteRequest.md) | The ids of the campaigns to delete. | 

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


## List Campaigns

> []GetActiveCampaigns200ResponseInner GetActiveCampaigns(ctx).Detail(detail).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()

Gets campaigns and returns them in a list. Can provide increased level of detail for each campaign if provided the correct query.

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
    detail := "FULL" // string | Determines whether slim, or increased level of detail is provided for each campaign in the returned list. Slim is the default behavior. (optional)
    limit := int32(250) // int32 | Max number of results to return. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 250)
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    filters := "name eq "Manager Campaign"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **id**: *eq, in*  **name**: *eq, sw*  **status**: *eq, in* (optional)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created** (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetActiveCampaigns(context.Background()).Detail(detail).Limit(limit).Offset(offset).Count(count).Filters(filters).Sorters(sorters).Execute()
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


## Get a campaign

> SlimCampaign GetCampaign(ctx, id).Execute()

Retrieves information for an existing campaign using the campaign's ID. Authorized callers must be a reviewer for this campaign, an ORG_ADMIN, or a CERT_ADMIN.

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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign to be retrieved

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetCampaign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.GetCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaign`: SlimCampaign
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.GetCampaign`: %v\n", resp)
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

[UserContextAuth](../README.md#UserContextAuth), [UserContextAuth](../README.md#UserContextAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get Campaign Reports

> []CampaignReport GetCampaignReports(ctx, id).Execute()

Fetches all reports for a certification campaign by campaign ID.
Requires roles of CERT_ADMIN, DASHBOARD, ORG_ADMIN and REPORT_ADMIN

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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign for which reports are being fetched.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetCampaignReports(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign for which reports are being fetched. | 

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


## Get Campaign Reports Configuration

> CampaignReportsConfig GetCampaignReportsConfig(ctx).Execute()

Fetches configuration for campaign reports. Currently it includes only one element - identity attributes defined as custom report columns.
Requires roles of CERT_ADMIN and ORG_ADMIN.

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

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetCampaignReportsConfig(context.Background()).Execute()
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


## Get a Campaign Template

> CampaignTemplate GetCampaignTemplate(ctx, id).Execute()

Fetches a campaign template by ID.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The desired campaign template's ID.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetCampaignTemplate(context.Background(), id).Execute()
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
**id** | **string** | The desired campaign template&#39;s ID. | 

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


## Gets a Campaign Template's Schedule

> Schedule GetCampaignTemplateSchedule(ctx, id).Execute()

Gets the schedule for a campaign template. Returns a 404 if there is no schedule set.

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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | The ID of the campaign template whose schedule is being fetched.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.GetCampaignTemplateSchedule(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign template whose schedule is being fetched. | 

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


## List Campaign Templates

> []CampaignTemplate ListCampaignTemplates(ctx).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()

Lists all CampaignTemplates. Scope can be reduced via standard V3 query params.

All CampaignTemplates matching the query params

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
    offset := int32(0) // int32 | Offset into the full result set. Usually specified with *limit* to paginate through the results. See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to 0)
    count := true // bool | If *true* it will populate the *X-Total-Count* response header with the number of results that would be returned if *limit* and *offset* were ignored.  Since requesting a total count can have a performance impact, it is recommended not to send **count=true** if that value will not be used.  See [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters) for more information. (optional) (default to false)
    sorters := "name" // string | Sort results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#sorting-results)  Sorting is supported for the following fields: **name, created, modified** (optional)
    filters := "name eq "manager template"" // string | Filter results using the standard syntax described in [V3 API Standard Collection Parameters](https://developer.sailpoint.com/idn/api/standard-collection-parameters#filtering-results)  Filtering is supported for the following fields and operators:  **name**: *eq, ge, gt, in, le, lt, ne, sw*  **id**: *eq, ge, gt, in, le, lt, ne, sw* (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.ListCampaignTemplates(context.Background()).Limit(limit).Offset(offset).Count(count).Sorters(sorters).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificationCampaignsAPI.ListCampaignTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCampaignTemplates`: []CampaignTemplate
    fmt.Fprintf(os.Stdout, "Response from `CertificationCampaignsAPI.ListCampaignTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignTemplatesRequest struct via the builder pattern


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


## Reassign Certifications

> CertificationTask Move(ctx, id).AdminReviewReassign(adminReviewReassign).Execute()

This API reassigns the specified certifications from one identity to another. A token with ORG_ADMIN or CERT_ADMIN authority is required to call this API.

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The certification campaign ID
    adminReviewReassign := *sailpoint.NewAdminReviewReassign() // AdminReviewReassign | 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.Move(context.Background(), id).AdminReviewReassign(adminReviewReassign).Execute()
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


## Update a Campaign Template

> CampaignTemplate PatchCampaignTemplate(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Allows updating individual fields on a campaign template using the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The ID of the campaign template being modified.
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.  The following fields are patchable: * name * description * deadlineDuration * campaign (all fields that are allowed during create) 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.PatchCampaignTemplate(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
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
**id** | **string** | The ID of the campaign template being modified. | 

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


## Set Campaign Reports Configuration

> CampaignReportsConfig SetCampaignReportsConfig(ctx).CampaignReportsConfig(campaignReportsConfig).Execute()

Overwrites configuration for campaign reports.
Requires roles CERT_ADMIN and ORG_ADMIN.

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
    campaignReportsConfig := *sailpoint.NewCampaignReportsConfig() // CampaignReportsConfig | Campaign Report Configuration

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.SetCampaignReportsConfig(context.Background()).CampaignReportsConfig(campaignReportsConfig).Execute()
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
 **campaignReportsConfig** | [**CampaignReportsConfig**](CampaignReportsConfig.md) | Campaign Report Configuration | 

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


## Sets a Campaign Template's Schedule

> SetCampaignTemplateSchedule(ctx, id).Schedule(schedule).Execute()

Sets the schedule for a campaign template. If a schedule already exists, it will be overwritten with the new one.

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
    id := "04bedce387bd47b2ae1f86eb0bb36dee" // string | The ID of the campaign template being scheduled.
    schedule := *sailpoint.NewSchedule("WEEKLY", *sailpoint.NewScheduleHours("LIST", []string{"Values_example"})) // Schedule |  (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    r, err := apiClient.V3.CertificationCampaignsAPI.SetCampaignTemplateSchedule(context.Background(), id).Schedule(schedule).Execute()
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
**id** | **string** | The ID of the campaign template being scheduled. | 

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


## Activate a Campaign

> map[string]interface{} StartCampaign(ctx, id).ActivateCampaignOptions(activateCampaignOptions).Execute()

Submits a job to activate the campaign with the given Id. The campaign must be staged.
Requires roles of CERT_ADMIN and ORG_ADMIN

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
    id := "ef38f94347e94562b5bb8424a56397d8" // string | The campaign id
    activateCampaignOptions := *sailpoint.NewActivateCampaignOptions() // ActivateCampaignOptions | Optional. If no timezone is specified, the standard UTC timezone is used (i.e. UTC+00:00). Although this can take any timezone, the intended value is the caller's timezone. The activation time calculated from the given timezone may cause the campaign deadline time to be modified, but it will remain within the original date. The timezone must be in a valid ISO 8601 format. (optional)

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.StartCampaign(context.Background(), id).ActivateCampaignOptions(activateCampaignOptions).Execute()
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
**id** | **string** | The campaign id | 

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


## Run Campaign Remediation Scan

> map[string]interface{} StartCampaignRemediationScan(ctx, id).Execute()

Kicks off remediation scan task for a certification campaign.
Requires roles of CERT_ADMIN and ORG_ADMIN

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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign for which remediation scan is being run.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.StartCampaignRemediationScan(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign for which remediation scan is being run. | 

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


## Run Campaign Report

> map[string]interface{} StartCampaignReport(ctx, id, type_).Execute()

Runs a report for a certification campaign.
Requires the following roles: CERT_ADMIN, DASHBOARD, ORG_ADMIN and REPORT_ADMIN.

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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign for which report is being run.
    type_ := sailpoint.ReportType("CAMPAIGN_COMPOSITION_REPORT") // ReportType | The type of the report to run.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.StartCampaignReport(context.Background(), id, type_).Execute()
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
**id** | **string** | The ID of the campaign for which report is being run. | 
**type_** | [**ReportType**](.md) | The type of the report to run. | 

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


## Generate a Campaign from Template

> CampaignReference StartGenerateCampaignTemplate(ctx, id).Execute()

Generates a new campaign from a campaign template.
The campaign object contained in the template has special formatting applied to its name and description fields in order to determine the generated campaign's name/description. Placeholders in those fields are formatted with the current date and time upon generation.
Placeholders consist of a percent sign followed by a letter indicating what should be inserted; for example, "%Y" will insert the current year; a campaign template named "Campaign for %y" would generate a campaign called "Campaign for 2020" (assuming the year at generation time is 2020).
Valid placeholders are the date/time conversion suffix characters supported by [java.util.Formatter](https://docs.oracle.com/javase/8/docs/api/java/util/Formatter.html).
Requires roles ORG_ADMIN.

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
    id := "2c9180835d191a86015d28455b4a2329" // string | The ID of the campaign template to use for generation.

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.StartGenerateCampaignTemplate(context.Background(), id).Execute()
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
**id** | **string** | The ID of the campaign template to use for generation. | 

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


## Update a Campaign

> SlimCampaign UpdateCampaign(ctx, id).JsonPatchOperation(jsonPatchOperation).Execute()

Allows updating individual fields on a campaign using the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard.

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
    id := "2c91808571bcfcf80171c23e4b4221fc" // string | The ID of the campaign template being modified.
    jsonPatchOperation := []sailpoint.JsonPatchOperation{*sailpoint.NewJsonPatchOperation("replace", "/description")} // []JsonPatchOperation | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The fields that can be patched differ based on the status of the campaign.  In the *STAGED* status, the following fields can be patched: * name * description * recommendationsEnabled * deadline * emailNotificationEnabled * autoRevokeAllowed  In the *ACTIVE* status, the following fields can be patched: * deadline 

    configuration := sailpoint.NewDefaultConfiguration()
    apiClient := sailpoint.NewAPIClient(configuration)
    resp, r, err := apiClient.V3.CertificationCampaignsAPI.UpdateCampaign(context.Background(), id).JsonPatchOperation(jsonPatchOperation).Execute()
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
**id** | **string** | The ID of the campaign template being modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonPatchOperation** | [**[]JsonPatchOperation**](JsonPatchOperation.md) | A list of campaign update operations according to the [JSON Patch](https://tools.ietf.org/html/rfc6902) standard. The fields that can be patched differ based on the status of the campaign.  In the *STAGED* status, the following fields can be patched: * name * description * recommendationsEnabled * deadline * emailNotificationEnabled * autoRevokeAllowed  In the *ACTIVE* status, the following fields can be patched: * deadline  | 

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

