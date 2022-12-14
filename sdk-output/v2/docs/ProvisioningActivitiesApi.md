# \ProvisioningActivitiesApi

All URIs are relative to *https://sailpoint.api.identitynow.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProvisioningActivity**](ProvisioningActivitiesApi.md#GetProvisioningActivity) | **Get** /provisioning-activities/{provisioningActivityId} | Retrieves a provisioning activity.
[**ListProvisioningActivities**](ProvisioningActivitiesApi.md#ListProvisioningActivities) | **Get** /provisioning-activities | Lists the provisioning activities.



## GetProvisioningActivity

> ProvisioningActivity GetProvisioningActivity(ctx, provisioningActivityId).Execute()

Retrieves a provisioning activity.



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
    provisioningActivityId := "provisioningActivityId_example" // string | ID of a provisioning activity.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningActivitiesApi.GetProvisioningActivity(context.Background(), provisioningActivityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningActivitiesApi.GetProvisioningActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProvisioningActivity`: ProvisioningActivity
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningActivitiesApi.GetProvisioningActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provisioningActivityId** | **string** | ID of a provisioning activity. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvisioningActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProvisioningActivity**](ProvisioningActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProvisioningActivities

> []ProvisioningActivity ListProvisioningActivities(ctx).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()

Lists the provisioning activities.



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
    filters := "filters_example" // string | Search filters. Supported operators are 'co', 'sw', 'eq', 'or', 'and', 'gt', 'lt', 'ge', 'le', 'ne', 'not' and 'pr' which are SCIM-compliant. Supported attributes for filtering are ‘operation’, ’name’, ‘sourceName’ and ‘status’. Example: 'property eq \"value\"'. (optional)
    sort := "sort_example" // string | One or more attributes on which to sort each separated by a ','. Prefix the attribute name with a minus sign (ex. -dateCreated) for descending sort. Supported attributes for sorting are 'operation', 'name', 'sourceName' ,'status', 'dateCreated' and 'lastUpdated'. (optional)
    offset := int32(56) // int32 | Paging offset. (optional) (default to 0)
    limit := int32(56) // int32 | Paging limit. (optional) (default to 250)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningActivitiesApi.ListProvisioningActivities(context.Background()).Filters(filters).Sort(sort).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningActivitiesApi.ListProvisioningActivities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProvisioningActivities`: []ProvisioningActivity
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningActivitiesApi.ListProvisioningActivities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListProvisioningActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | Search filters. Supported operators are &#39;co&#39;, &#39;sw&#39;, &#39;eq&#39;, &#39;or&#39;, &#39;and&#39;, &#39;gt&#39;, &#39;lt&#39;, &#39;ge&#39;, &#39;le&#39;, &#39;ne&#39;, &#39;not&#39; and &#39;pr&#39; which are SCIM-compliant. Supported attributes for filtering are ‘operation’, ’name’, ‘sourceName’ and ‘status’. Example: &#39;property eq \&quot;value\&quot;&#39;. | 
 **sort** | **string** | One or more attributes on which to sort each separated by a &#39;,&#39;. Prefix the attribute name with a minus sign (ex. -dateCreated) for descending sort. Supported attributes for sorting are &#39;operation&#39;, &#39;name&#39;, &#39;sourceName&#39; ,&#39;status&#39;, &#39;dateCreated&#39; and &#39;lastUpdated&#39;. | 
 **offset** | **int32** | Paging offset. | [default to 0]
 **limit** | **int32** | Paging limit. | [default to 250]

### Return type

[**[]ProvisioningActivity**](ProvisioningActivity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

